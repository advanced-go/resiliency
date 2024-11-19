package guidance

import (
	"github.com/advanced-go/common/core"
	"github.com/advanced-go/common/messaging"
)

// Guidance - guidance functions struct, a nod to Linus Torvalds and plain C
type Guidance struct {
	Assignments    func(h messaging.Notifier, origin core.Origin) ([]HostEntry, *core.Status)
	NewAssignments func(h messaging.Notifier, origin core.Origin) ([]HostEntry, *core.Status)
}

var Guide = func() *Guidance {
	return &Guidance{
		Assignments: func(h messaging.Notifier, origin core.Origin) ([]HostEntry, *core.Status) {
			e, status := GetEntry(origin)
			if !status.OK() {
				h.Notify(status)
			}
			return []HostEntry{e}, status
		},
		NewAssignments: func(h messaging.Notifier, origin core.Origin) ([]HostEntry, *core.Status) {
			return nil, core.StatusNotFound()
		},
	}
}()
