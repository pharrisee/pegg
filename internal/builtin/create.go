package builtin

import (
	"github.com/pharrisee/pegg/internal/builtin/agents"
	"github.com/pharrisee/pegg/internal/builtin/middlewares"
	"github.com/pharrisee/pegg/internal/builtin/reminders"
	"github.com/pharrisee/pegg/internal/builtin/tools"
	"github.com/pharrisee/pegg/internal/core/config"
	"github.com/pharrisee/pegg/internal/core/event"
	"github.com/pharrisee/pegg/internal/decision"
	"github.com/pharrisee/pegg/internal/llm"
	"github.com/pharrisee/pegg/internal/memory"
	"github.com/pharrisee/pegg/internal/schedule"
	"github.com/pharrisee/pegg/internal/session"
	"github.com/pharrisee/pegg/internal/vfs"
)

type Options struct {
	LLM       *llm.Manager
	Decision  *decision.Manager
	Config    *config.Config
	Bus       event.Bus
	Memory    *memory.Manager
	Scheduler *schedule.Manager
}

func Create(fs *vfs.VFS, sess *session.Manager, opts Options) error {
	agents.Create(fs)
	middlewares.Create(fs, middlewares.Deps{
		Config:   opts.Config,
		LLM:      opts.LLM,
		Decision: opts.Decision,
		Sessions: sess,
	})
	tools.Create(fs, sess)
	if opts.Memory != nil {
		if err := opts.Memory.Start(); err != nil {
			return err
		}
		if err := memory.RegisterTools(opts.Memory); err != nil {
			return err
		}
	}
	if err := reminders.Create(opts.Bus); err != nil {
		return err
	}
	if opts.Scheduler != nil {
		schedule.ScheduleTools(opts.Scheduler)
	}
	return nil
}
