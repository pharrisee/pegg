package reminders

import (
	"github.com/pharrisee/pegg/internal/builtin/reminders/commentline"
	"github.com/pharrisee/pegg/internal/builtin/reminders/todowrite"
	"github.com/pharrisee/pegg/internal/builtin/reminders/usage"
	"github.com/pharrisee/pegg/internal/core/event"
)

func Create(bus event.Bus) error {
	if bus == nil {
		return nil
	}
	if err := usage.New().Start(bus); err != nil {
		return err
	}
	if err := todowrite.New().Start(bus); err != nil {
		return err
	}
	if err := commentline.New().Start(bus); err != nil {
		return err
	}
	return nil
}
