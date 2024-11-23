package guidance

import (
	"github.com/advanced-go/common/core"
	"time"
)

var (
	//safeEntry = common.NewSafe()
	westData = []HostEntry{
		{Origin: core.Origin{Region: WestRegion, Zone: WestZoneA, SubZone: "", Host: "www.west-host1.com"}, CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
		{Origin: core.Origin{Region: WestRegion, Zone: WestZoneB, SubZone: "", Host: "www.west-host2.com"}, CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
	}

	centralData = []HostEntry{
		{Origin: core.Origin{Region: CentralRegion, Zone: CentralZoneA, SubZone: "", Host: "www.central-host1.com"}, CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
		{Origin: core.Origin{Region: CentralRegion, Zone: CentralZoneB, SubZone: "", Host: "www.central-host2.com"}, CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
	}

	eastData = []HostEntry{
		{Origin: core.Origin{Region: EastRegion, Zone: EastZoneA, SubZone: "", Host: "www.east-host1.com"}, CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
	}
)

// HostEntry - host
type HostEntry struct {
	EntryId   int         `json:"entry-id"`
	CreatedTS time.Time   `json:"created-ts"`
	Origin    core.Origin `json:"origin"`
}
