package guidance

import (
	"github.com/advanced-go/common/core"
	"github.com/advanced-go/common/messaging"
	"github.com/advanced-go/resiliency/common"
)

type ProcessingCalendar struct {
	week [7][24]string
}

func NewProcessingCalendar() *ProcessingCalendar {
	c := new(ProcessingCalendar)
	return c
}

func GetCalendar(h core.ErrorHandler, agentId string, msg *messaging.Message) *ProcessingCalendar {
	if !msg.IsContentType(common.ContentTypeCalendar) {
		return nil
	}
	if p, ok := msg.Body.(*ProcessingCalendar); ok {
		return p
	}
	h.Handle(common.ProfileTypeErrorStatus(agentId, msg.Body))
	return nil
}
