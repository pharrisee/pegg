package live

import (
	"github.com/pharrisee/pegg/internal/session"
)

type EventType string

const (
	EventRunStarted  EventType = "run.started"
	EventRunEvent    EventType = "run.event"
	EventRunFinished EventType = "run.finished"
)

type Event struct {
	Type    EventType
	Run     session.Run
	Seq     int64
	Journal session.RunEvent
	Local   bool
	State   session.RunState
	Error   string
}
