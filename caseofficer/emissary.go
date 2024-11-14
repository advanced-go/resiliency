package caseofficer1

import (
	"github.com/advanced-go/common/messaging"
	"github.com/advanced-go/resiliency/common"
)

func emissaryAttend(c *caseOfficer, fn *caseOfficerFunc) {
	//processMsg := messaging.NewControlMessage("", "", messaging.ProcessEvent)
	fn.startup(c, nil)

	for {
		select {
		case <-c.ticker.C():
			c.handler.AddActivity(c.agentId, "onTick()")
			//c.failoverAgent.Message(processMsg)
			//c.redirectAgent.Message(processMsg)
			fn.update(c, nil)
		default:
		}
		// control channel processing
		select {
		case msg := <-c.ctrlC:
			switch msg.Event() {
			case messaging.ShutdownEvent:
				c.shutdown()
				c.handler.AddActivity(c.agentId, messaging.ShutdownEvent)
				return
			case messaging.DataChangeEvent:
				if msg.IsContentType(common.ContentTypeProfile) {
					c.serviceAgents.Broadcast(msg)
				}
			default:
				c.handler.Handle(common.MessageEventErrorStatus(c.agentId, msg))
			}
		default:
		}
	}
}
