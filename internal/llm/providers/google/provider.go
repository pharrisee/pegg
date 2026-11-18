package google

import (
	"time"

	"github.com/pharrisee/pegg/internal/core/config"
	"github.com/pharrisee/pegg/internal/llm"
	geminidriver "github.com/pharrisee/pegg/internal/llm/drivers/gemini"
)

const (
	ProviderName   = "google"
	DefaultBaseURL = "https://generativelanguage.googleapis.com/v1beta"
)

func init() {
	llm.RegisterProvider(ProviderName, NewFromConfig)
}

func NewFromConfig(cfg config.LLMConfig) (llm.Provider, error) {
	return geminidriver.NewService(ProviderName, geminidriver.ServiceConfig{
		BaseURL: DefaultBaseURL,
		APIKey:  cfg.APIKey,
		Headers: cfg.Headers,
		Timeout: time.Duration(cfg.Timeout) * time.Second,
	}), nil
}
