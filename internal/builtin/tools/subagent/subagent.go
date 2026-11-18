package subagent

import (
	"github.com/pharrisee/pegg/internal/agent/tools"
	"github.com/pharrisee/pegg/internal/core/event"
	"github.com/pharrisee/pegg/internal/session"
)

type HistoryReader interface {
	Messages(sessionID string) ([]session.Message, error)
	Bus() event.Bus
}

var sessionReader HistoryReader

func SubAgentTools(reader HistoryReader) {
	sessionReader = reader
	if reader != nil {
		store.subscribe(reader.Bus())
	}
	tools.Register(subAgentTool())
	tools.Register(subAgentsStatusTool())
}
