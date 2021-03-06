/*
 * HCS API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package hcsschema

type Properties struct {

	Id string `json:"Id,omitempty"`

	SystemType string `json:"SystemType,omitempty"`

	RuntimeOsType string `json:"RuntimeOsType,omitempty"`

	Name string `json:"Name,omitempty"`

	Owner string `json:"Owner,omitempty"`

	ObRoot string `json:"ObRoot,omitempty"`

	IsDummy bool `json:"IsDummy,omitempty"`

	RuntimeId string `json:"RuntimeId,omitempty"`

	RuntimeImagePath string `json:"RuntimeImagePath,omitempty"`

	IsRuntimeTemplate bool `json:"IsRuntimeTemplate,omitempty"`

	RuntimeTemplateId string `json:"RuntimeTemplateId,omitempty"`

	State string `json:"State,omitempty"`

	Stopped bool `json:"Stopped,omitempty"`

	ExitType string `json:"ExitType,omitempty"`

	Memory *MemoryInformationForVm `json:"Memory,omitempty"`

	ContainerReportedMemory *ContainerMemoryInformation `json:"ContainerReportedMemory,omitempty"`

	CpuGroupId string `json:"CpuGroupId,omitempty"`

	Statistics *Statistics `json:"Statistics,omitempty"`

	ProcessList []ProcessDetails `json:"ProcessList,omitempty"`

	TerminateOnLastHandleClosed bool `json:"TerminateOnLastHandleClosed,omitempty"`

	SystemGUID string `json:"SystemGUID,omitempty"`

	HostingSystemId string `json:"HostingSystemId,omitempty"`

	SharedMemoryRegionInfo []SharedMemoryRegionInfo `json:"SharedMemoryRegionInfo,omitempty"`

	GuestInterfaceInfo *GuestInterfaceInfo `json:"GuestInterfaceInfo,omitempty"`

	Silo *SiloProperties `json:"Silo,omitempty"`

	CosIndex int32 `json:"CosIndex,omitempty"`

	Rmid int32 `json:"Rmid,omitempty"`

	CacheStats *CacheQueryStatsResponse `json:"CacheStats,omitempty"`

	//   Controller # -> list of disks on that controller. For Argons, the controller number is synthetic and has no bearing. 
	MappedVirtualDiskControllers map[string]MappedVirtualDiskController `json:"MappedVirtualDiskControllers,omitempty"`
}
