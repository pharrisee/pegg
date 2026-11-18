package agents

import (
	"github.com/pharrisee/pegg/internal/builtin/agents/developer"
	"github.com/pharrisee/pegg/internal/builtin/agents/explorer"
	"github.com/pharrisee/pegg/internal/builtin/agents/orchestrator"
	"github.com/pharrisee/pegg/internal/vfs"
)

func Create(fs *vfs.VFS) {
	if err := explorer.Register(); err != nil {
		fs.Warn("agents: explorer is disabled: %v", err)
	}
	if err := developer.Register(fs); err != nil {
		fs.Warn("agents: developer is disabled: %v", err)
	}
	if err := orchestrator.Register(fs); err != nil {
		fs.Warn("agents: orchestrator is disabled: %v", err)
	}
}
