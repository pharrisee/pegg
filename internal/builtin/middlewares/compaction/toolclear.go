package compaction

import (
	"context"

	"github.com/pharrisee/pegg/internal/agent/middleware"
	"github.com/pharrisee/pegg/internal/llm"
)

func (m *Middleware) InvokeTool(ctx context.Context, call llm.ToolCall, next middleware.ToolInvoker) (string, error) {
	output, err := next(ctx, call)

	if m.cfg == nil || !m.cfg.Enabled || !m.hasStrategy("tool-clearing") {
		return output, err
	}

	maxChars := m.cfg.MaxToolOutputChars
	if maxChars <= 0 {
		maxChars = 4000
	}

	if len(output) > maxChars {
		output = output[:maxChars] + "\n... [truncated]"
	}

	return output, err
}
