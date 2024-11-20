package guidance

import (
	"github.com/advanced-go/common/messaging"
)

type ProcessingCalendar struct {
	week [7][24]string
}

func NewProcessingCalendar() *ProcessingCalendar {
	c := new(ProcessingCalendar)
	return c
}

func GetCalendar(h messaging.Notifier, agentId string, msg *messaging.Message) *ProcessingCalendar {
	if !msg.IsContentType(ContentTypeCalendar) {
		return nil
	}
	if p, ok := msg.Body.(*ProcessingCalendar); ok {
		return p
	}
	h.Notify(nil, CalendarTypeErrorStatus(agentId, msg.Body))
	return nil
}
