package settings

import (
	"github.com/pharrisee/pegg/internal/agent"
	"github.com/pharrisee/pegg/internal/core/cache"
	"github.com/pharrisee/pegg/internal/core/config"
	"github.com/pharrisee/pegg/internal/core/event"
	"github.com/pharrisee/pegg/internal/decision"
	"github.com/pharrisee/pegg/internal/live"
	"github.com/pharrisee/pegg/internal/llm"
	"github.com/pharrisee/pegg/internal/mcp"
	"github.com/pharrisee/pegg/internal/memory"
	"github.com/pharrisee/pegg/internal/plugin"
	"github.com/pharrisee/pegg/internal/run"
	"github.com/pharrisee/pegg/internal/session"
	"github.com/pharrisee/pegg/internal/vfs"
)

type Deps struct {
	Config       *config.Config
	LLM          *llm.Manager
	Decision     *decision.Manager
	Memory       *memory.Manager
	MCP          *mcp.Manager
	Sessions     *session.Manager
	Agent        *agent.Agent
	AgentFactory func() (*agent.Agent, error)
	Bus          event.Bus
	VFS          *vfs.VFS
	Cache        cache.Cache
	Plugin       *plugin.Manager
	Runs         *run.Coordinator
	Live         *live.Hub
}
