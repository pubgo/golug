package golug_cmd

import (
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/pubgo/dix/dix_run"
	"github.com/pubgo/golug/golug_app"
	"github.com/pubgo/golug/golug_config"
	"github.com/pubgo/golug/golug_entry"
	"github.com/pubgo/golug/golug_plugin"
	"github.com/pubgo/golug/golug_watcher"
	"github.com/pubgo/golug/version"
	"github.com/pubgo/xerror"
	"github.com/pubgo/xlog"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: golug_app.Domain, Version: version.Version}

func Init() (err error) {
	defer xerror.RespErr(&err)

	// 处理所有的配置,环境变量和flag
	// 配置顺序, 默认值->环境变量->配置文件->flag->配置文件
	// 配置文件中可以设置环境变量
	// flag可以指定配置文件位置
	// 始化配置文件
	xerror.Panic(golug_config.Init())

	// 初始化框架, 加载环境变量, 加载本地配置
	// 初始化完毕所有的配置以及外部配置以及相关的参数和变量
	// 剩下的就是获取配置了
	if !golug_config.IsExist() {
		xerror.Panic(golug_config.InitProject())
	}

	// 指定配置文件
	if !golug_config.IsExist() && golug_config.CfgPath != "" {
		xerror.Panic(golug_config.InitWithCfgPath())
	}

	xerror.Assert(golug_config.GetCfg().ConfigFileUsed() == "", "config file not found")

	xerror.ExitF(golug_config.GetCfg().ReadInConfig(), "read config failed")
	golug_config.CfgPath = golug_config.GetCfg().ConfigFileUsed()
	golug_app.Home = filepath.Dir(filepath.Dir(golug_config.CfgPath))

	xerror.Panic(golug_config.InitOtherConfig())

	xerror.Panic(golug_app.CheckMod())

	xerror.Panic(golug_config.Trigger())
	return nil
}

func handleSignal() {
	if golug_app.CatchSigpipe {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGPIPE)
		go func() {
			<-sigChan
			xlog.Warn("Caught SIGPIPE (ignoring all future SIGPIPEs)")
			signal.Ignore(syscall.SIGPIPE)
		}()
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL, syscall.SIGHUP)
	golug_app.Signal = <-ch
}

func start(ent golug_entry.RunEntry) (err error) {
	defer xerror.RespErr(&err)

	xerror.Panic(dix_run.BeforeStart())
	xerror.Panic(ent.Start())
	xerror.Panic(dix_run.AfterStart())

	return
}

func stop(ent golug_entry.RunEntry) (err error) {
	defer xerror.RespErr(&err)

	xerror.Panic(dix_run.BeforeStop())
	xerror.Panic(ent.Stop())
	xerror.Panic(dix_run.AfterStop())

	return nil
}

func Run(entries ...golug_entry.Entry) (err error) {
	defer xerror.RespErr(&err)

	xerror.Assert(len(entries) == 0, "[entries] should not be zero")

	for _, ent := range entries {
		xerror.Assert(ent == nil, "[ent] should not be nil")

		entRun := ent.(golug_entry.RunEntry)
		opt := entRun.Options()
		xerror.Assert(opt.Name == "" || opt.Version == "", "[name], [version] should not be empty")
	}

	rootCmd.PersistentFlags().AddFlagSet(golug_app.DefaultFlags())
	rootCmd.PersistentFlags().AddFlagSet(golug_config.DefaultFlags())
	rootCmd.RunE = func(cmd *cobra.Command, args []string) error { return xerror.Wrap(cmd.Help()) }

	for _, ent := range entries {
		entRun := ent.(golug_entry.RunEntry)
		cmd := entRun.Options().Command

		// 检查Command是否注册
		for _, c := range rootCmd.Commands() {
			xerror.Assert(c.Name() == cmd.Name(), "command(%s) already exists", cmd.Name())
		}

		// 注册plugin的command和flags
		entPlugins := golug_plugin.List(golug_plugin.Module(entRun.Options().Name))
		for _, pl := range append(golug_plugin.List(), entPlugins...) {
			cmd.PersistentFlags().AddFlagSet(pl.Flags())
			ent.Commands(pl.Commands())
		}

		cmd.RunE = func(cmd *cobra.Command, args []string) (err error) {
			defer xerror.RespErr(&err)

			golug_app.Project = entRun.Options().Name
			xerror.Panic(Init())
			xerror.Panic(entRun.Init())

			// 初始化组件, 初始化插件
			plugins := golug_plugin.List(golug_plugin.Module(entRun.Options().Name))
			plugins = append(golug_plugin.List(), plugins...)
			for _, pg := range plugins {
				key := pg.String()
				xerror.PanicF(pg.Init(ent), "plugin [%s] init error", key)
				golug_watcher.Watch(key, pg.Watch)
			}

			xerror.Panic(start(entRun))

			if golug_app.IsBlock {
				handleSignal()
			}

			xerror.Panic(stop(entRun))
			return nil
		}

		rootCmd.AddCommand(cmd)
	}

	return xerror.Wrap(rootCmd.Execute())
}
