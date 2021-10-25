package runtime

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/pubgo/dix"
	"github.com/pubgo/x/stack"
	"github.com/pubgo/xerror"
	"github.com/spf13/cobra"

	"github.com/pubgo/lava/config"
	"github.com/pubgo/lava/entry"
	"github.com/pubgo/lava/healthy"
	v "github.com/pubgo/lava/internal/cmds/version"
	"github.com/pubgo/lava/internal/logz"
	"github.com/pubgo/lava/pkg/syncx"
	"github.com/pubgo/lava/plugin"
	"github.com/pubgo/lava/plugins/watcher"
	"github.com/pubgo/lava/runenv"
	"github.com/pubgo/lava/vars"
	"github.com/pubgo/lava/version"
)

const name = "runtime"

var logs = logz.New(name)
var rootCmd = &cobra.Command{Use: runenv.Domain, Version: version.Version}

func init() {
	rootCmd.AddCommand(v.Cmd)
	rootCmd.AddCommand(healthy.Cmd)
}

func handleSignal() {
	if runenv.CatchSigpipe {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGPIPE)
		syncx.GoSafe(func() {
			<-sigChan
			logs.Warn("Caught SIGPIPE (ignoring all future SIGPIPE)")
			signal.Ignore(syscall.SIGPIPE)
		})
	}

	if !runenv.Block {
		return
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL, syscall.SIGHUP)
	runenv.Signal = <-ch
	logs.Infof("signal [%s] trigger", runenv.Signal.String())
}

func start(ent entry.Runtime) (err error) {
	defer xerror.RespErr(&err)

	logs.Info("before-start running")
	beforeList := append(entry.GetBeforeStartsList(), ent.Options().BeforeStarts...)
	for i := range beforeList {
		xerror.TryThrow(beforeList[i], "before-start error", stack.Func(beforeList[i]))
	}
	logs.Info("before-start ok")

	xerror.Panic(ent.Start())

	logs.Info("after-start running")
	afterList := append(entry.GetAfterStartsList(), ent.Options().AfterStarts...)
	for i := range afterList {
		xerror.TryThrow(afterList[i], "after-start error", stack.Func(afterList[i]))
	}
	logs.Info("after-start ok")
	return
}

func stop(ent entry.Runtime) (err error) {
	defer xerror.RespErr(&err)

	logs.Info("before-stop running")
	beforeList := append(entry.GetBeforeStopsList(), ent.Options().BeforeStops...)
	for i := range beforeList {
		logs.TryWith(beforeList[i]).Errorf("before-stop error: %s", stack.Func(beforeList[i]))
	}
	logs.Info("before-stop ok")

	xerror.Panic(ent.Stop())

	logs.Info("after-stop running")
	afterList := append(entry.GetAfterStopsList(), ent.Options().AfterStops...)
	for i := range afterList {
		logs.TryWith(afterList[i]).Errorf("after-stop error: %s", stack.Func(afterList[i]))
	}
	logs.Info("after-stop ok")
	return nil
}

func Run(description string, entries ...entry.Entry) {
	defer xerror.RespExit()

	xerror.Assert(len(entries) == 0, "[entries] should not be zero")

	for _, ent := range entries {
		xerror.Assert(ent == nil, "[ent] should not be nil")

		_, ok := ent.(entry.Runtime)
		if !ok {
			panic(fmt.Sprintf("[ent] not implement runtime, ent:%#v", ent))
		}

	}

	rootCmd.Short = description
	rootCmd.Long = description
	rootCmd.PersistentFlags().AddFlagSet(config.DefaultFlags())
	rootCmd.RunE = func(cmd *cobra.Command, args []string) error { return xerror.Wrap(cmd.Help()) }

	for i := range entries {
		ent := entries[i]
		entRT := ent.(entry.Runtime)
		cmd := entRT.Options().Command

		// 检查Command是否注册
		for _, c := range rootCmd.Commands() {
			xerror.Assert(c.Name() == cmd.Name(), "command(%s) already exists", cmd.Name())
		}

		// 注册plugin的command和flags
		// 先注册全局, 后注册项目相关
		xerror.TryThrow(func() {
			entPlugins := plugin.ListWithDefault(plugin.Module(entRT.Options().Name))
			for _, plg := range entPlugins {
				cmd.PersistentFlags().AddFlagSet(plg.Flags())
				ent.Commands(plg.Commands())
				entRT.MiddlewareInter(plg.Middleware())
			}
		})

		cmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
			defer xerror.RespExit()

			// 项目名初始化
			runenv.Project = entRT.Options().Name

			// config初始化
			xerror.Panic(config.Init())

			// 配置依赖注入
			xerror.Exit(dix.Provider(config.GetCfg()))

			xerror.TryThrow(func() {
				// 获取本项目所有plugin
				plugins := plugin.ListWithDefault(plugin.Module(entRT.Options().Name))
				for _, plg := range plugins {

					// 注册watcher
					logs.Infof("plugin [%s] watch register", plg.Id())
					watcher.Watch("watcher/"+plg.Id(), plg.Watch)

					// 注册健康检查
					healthy.Register(plg.Id(), plg.Health())

					// 注册vars
					xerror.Panic(plg.Vars(vars.Watch))
				}
			})

			// plugin初始化
			plugins := plugin.ListWithDefault(plugin.Module(runenv.Project))
			for _, plg := range plugins {
				logs.Infof("plugin [%s] init", plg.Id())
				xerror.PanicF(plg.Init(ent), "plugin [%s] init error", plg.String())
			}

			// watcher初始化
			xerror.Panic(watcher.Init())

			// entry初始化
			entRT.InitRT()
		}

		cmd.Run = func(cmd *cobra.Command, args []string) {
			defer xerror.RespExit()
			xerror.Panic(start(entRT))
			handleSignal()
			xerror.Panic(stop(entRT))
		}

		rootCmd.AddCommand(cmd)
	}

	xerror.Panic(rootCmd.Execute())
}

//func Start(ent entry.Entry) {
//	defer xerror.RespExit()
//
//	xerror.Assert(ent == nil, "[ent] should not be nil")
//
//	entRun, ok := ent.(entry.Runtime)
//	xerror.Assert(!ok, "[ent] not implement runtime")
//
//	opt := entRun.Options()
//	xerror.Assert(opt.Name == "", "[name] should not be empty")
//
//	plugins := plugin.ListWithDefault(plugin.Module(entRun.Options().Name))
//	for _, pl := range plugins {
//		// 加载flag
//		_ = pl.Flags()
//	}
//
//	// config初始化
//	runenv.Project = entRun.Options().Name
//	xerror.Panic(config.Init())
//
//	// plugin初始化
//	for _, pg := range plugins {
//		key := pg.String()
//		xerror.PanicF(pg.Init(ent), "plugin [%s] init error", key)
//
//		// watch key
//		watcher.Watch(key, pg.Watch)
//	}
//
//	xerror.Panic(watcher.Init())
//
//	// entry初始化
//	entRun.InitRT()
//
//	xerror.Panic(start(entRun))
//}
//
//func Stop(ent entry.Entry) {
//	defer xerror.RespExit()
//
//	xerror.Assert(ent == nil, "[ent] should not be nil")
//
//	entRun, ok := ent.(entry.Runtime)
//	xerror.Assert(!ok, "[ent] not implement runtime")
//
//	xerror.Panic(stop(entRun))
//}
