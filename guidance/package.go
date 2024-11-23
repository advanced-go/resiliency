package guidance

import "github.com/advanced-go/common/core"

const (
	PkgPath    = "github/advanced-go/resiliency/guidance"
	WestRegion = "us-west1"
	WestZoneA  = "us-west1-a"
	WestZoneB  = "us-west1-b"

	CentralRegion = "us-central1"
	CentralZoneA  = "us-central1-a"
	CentralZoneB  = "us-central1-b"

	EastRegion = "us-east1"
	EastZoneA  = "us-east1-a"
	EastZoneB  = "us-east1-b"
)

func GetRegion(origin core.Origin) ([]HostEntry, *core.Status) {
	if origin.Region == WestRegion {
		return westData, core.StatusOK()
	}
	if origin.Region == CentralRegion {
		return centralData, core.StatusOK()
	}
	return []HostEntry{}, core.StatusNotFound()
}
