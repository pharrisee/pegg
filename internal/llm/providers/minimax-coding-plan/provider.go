package minimaxcodingplan

import (
	"time"

	"github.com/pharrisee/pegg/internal/core/config"
	"github.com/pharrisee/pegg/internal/llm"
	claudedriver "github.com/pharrisee/pegg/internal/llm/drivers/claude"
)

const (
	ProviderName   = "minimax-coding-plan"
	DefaultBaseURL = "https://api.minimax.io/anthropic"
)

func init() {
	llm.RegisterProvider(ProviderName, NewFromConfig)
}

func NewFromConfig(cfg config.LLMConfig) (llm.Provider, error) {
	return claudedriver.NewService(ProviderName, claudedriver.ServiceConfig{
		BaseURL: DefaultBaseURL,
		APIKey:  cfg.APIKey,
		Headers: cfg.Headers,
		Timeout: time.Duration(cfg.Timeout) * time.Second,
	}), nil
}
