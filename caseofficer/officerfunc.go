package caseofficer1

import (
	"github.com/advanced-go/common/core"
	"github.com/advanced-go/common/messaging"
	"github.com/advanced-go/resiliency/common"
	"github.com/advanced-go/resiliency/service"
)

// A nod to Linus Torvalds and plain C
type caseOfficerFunc struct {
	startup func(c *caseOfficer, guide *common.Guidance, newAgent newServiceAgent) *core.Status
	update  func(c *caseOfficer, guide *common.Guidance, newAgent newServiceAgent) *core.Status
}

var (
	officer = func() *caseOfficerFunc {
		return &caseOfficerFunc{
			startup: func(c *caseOfficer, guide *common.Guidance, newAgent newServiceAgent) *core.Status {
				entry, status := guide.Assignments(c.handler, c.origin)
				if status.OK() {
					updateExchange(c, entry, newAgent)
				}
				c.startup()
				return core.StatusOK()
			},
			update: func(c *caseOfficer, guide *common.Guidance, newAgent newServiceAgent) *core.Status {
				entry, status := guide.NewAssignments(c.handler, c.origin)
				if status.OK() {
					updateExchange(c, entry, newAgent)
				}
				return core.StatusOK()
			},
		}
	}()
)

func updateExchange(c *caseOfficer, entries []common.HostEntry, newAgent newServiceAgent) {
	for _, e := range entries {
		newAgent(e.Origin, c)
		newAgent(e.Origin, c)
	}
}

func initAgent(origin core.Origin, c *caseOfficer) {
	var agent messaging.Agent
	var err error

	agent = service.NewAgent(origin, c)
	err = c.serviceAgents.Register(agent)
	if err != nil {
		c.handler.Handle(core.NewStatusError(core.StatusInvalidArgument, err))
	} else {
		agent.Run()
	}
}
