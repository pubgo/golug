package golug_base

import (
	"fmt"
	"strings"

	ver "github.com/hashicorp/go-version"
	"github.com/pubgo/dix"
	"github.com/pubgo/golug/golug_app"
	"github.com/pubgo/golug/golug_config"
	"github.com/pubgo/golug/golug_entry"
	"github.com/pubgo/xerror"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var _ golug_entry.Entry = (*baseEntry)(nil)

type baseEntry struct {
	opts golug_entry.Options
}

func (t *baseEntry) OnCfg(fn interface{}) {
	xerror.Assert(fn == nil, "[fn] is null")

	golug_config.On(func(cfg *golug_config.Config) { golug_config.Decode(t.opts.Name, fn) })
}

func (t *baseEntry) Init() (err error) {
	defer xerror.RespErr(&err)

	xerror.Assert(golug_app.Project != t.Options().Name, "project name not match(%s, %s)", golug_app.Project, t.Options().Name)

	t.opts.Initialized = true
	return
}
func (t *baseEntry) Dix(data ...interface{})      { xerror.Next().Panic(dix.Dix(data...)) }
func (t *baseEntry) Run() golug_entry.RunEntry    { return t }
func (t *baseEntry) Start() error                 { return nil }
func (t *baseEntry) Stop() error                  { return nil }
func (t *baseEntry) UnWrap(fn interface{})        { panic("implement me") }
func (t *baseEntry) Options() golug_entry.Options { return t.opts }
func (t *baseEntry) Flags(fn func(flags *pflag.FlagSet)) {
	defer xerror.RespExit()
	fn(t.opts.Command.PersistentFlags())
}

func (t *baseEntry) Description(description ...string) {
	t.opts.Command.Short = fmt.Sprintf("This is a %s service", t.opts.Name)

	if len(description) > 0 {
		t.opts.Command.Short = description[0]
	}
	if len(description) > 1 {
		t.opts.Command.Long = description[1]
	}
	if len(description) > 2 {
		t.opts.Command.Example = description[2]
	}

	return
}

func (t *baseEntry) Version(v string) {
	t.opts.Version = strings.TrimSpace(v)
	if t.opts.Version == "" {
		return
	}

	t.opts.Command.Version = v
	_, err := ver.NewVersion(v)
	xerror.Next().Panic(err)
	return
}

func (t *baseEntry) Commands(commands ...*cobra.Command) {
	rootCmd := t.opts.Command
	for _, cmd := range commands {
		if cmd == nil {
			continue
		}

		if rootCmd.Name() == cmd.Name() {
			return
		}

		rootCmd.AddCommand(cmd)
	}
}

func (t *baseEntry) initFlags() {
	t.Flags(func(flags *pflag.FlagSet) {
		flags.UintVar(&t.opts.Port, "port", t.opts.Port, "the server port")
	})
}

func handleCmdName(name string) string {
	if !strings.Contains(name, ".") {
		return name
	}

	names := strings.Split(name, ".")
	return names[len(names)-1]
}

func newEntry(name string) *baseEntry {
	name = strings.TrimSpace(name)
	xerror.Assert(name == "", "the [name] parameter should not be empty")
	xerror.Assert(strings.Contains(name, " "), "[name] should not contain blank")

	ent := &baseEntry{
		opts: golug_entry.Options{
			Name:    name,
			Port:    8080,
			Command: &cobra.Command{Use: handleCmdName(name)},
		},
	}

	ent.initFlags()

	return ent
}

func New(name string) *baseEntry {
	return newEntry(name)
}
