package bee

import (
	"time"

	"github.com/pharrisee/pegg/internal/core/config"
	"github.com/pharrisee/pegg/internal/llm"
	openaidriver "github.com/pharrisee/pegg/internal/llm/drivers/openai"
)

const (
	ProviderName   = "bee"
	DefaultBaseURL = "https://api.bee.heossi.com/bee"
)

func init() {
	llm.RegisterProvider(ProviderName, NewFromConfig)
}

func NewFromConfig(cfg config.LLMConfig) (llm.Provider, error) {
	return openaidriver.NewService(ProviderName, openaidriver.ServiceConfig{
		BaseURL: DefaultBaseURL,
		APIKey:  cfg.APIKey,
		Headers: cfg.Headers,
		Timeout: time.Duration(cfg.Timeout) * time.Second,
	}), nil
}
