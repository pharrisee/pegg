package boundless

import (
	"time"

	"github.com/pharrisee/pegg/internal/core/config"
	"github.com/pharrisee/pegg/internal/llm"
	openaidriver "github.com/pharrisee/pegg/internal/llm/drivers/openai"
)

const (
	ProviderName   = "boundless"
	DefaultBaseURL = "https://api.inference.boundless.network/v1"
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
