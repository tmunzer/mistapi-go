/*
Mist API

> Version: **2406.1.14** > > Date: **July 3, 2024**  ---  ### Additional Documentation * [Mist Automation Guide](https://www.juniper.net/documentation/us/en/software/mist/automation-integration/index.html) * [Mist Location SDK](https://www.juniper.net/documentation/us/en/software/mist/location_services/topics/concept/mist-how-get-mist-sdk.html) * [Mist Product Updates](https://www.mist.com/documentation/category/product-updates/)  ---  ### Helpful Resources * [API Sandbox and Exercises](https://api-class.mist.com/) * [Postman Collection, Runners and Webhook Samples](https://www.postman.com/juniper-mist/workspace/mist-systems-s-public-workspace) * [API Demo Apps](https://apps.mist-lab.fr/) * [Juniper Blog](https://blogs.juniper.net/)  --- 

API version: 2406.1.14
Contact: tmunzer@juniper.net
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mistapigo

import (
	"encoding/json"
)

// checks if the UpgradeOrgDevices type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpgradeOrgDevices{}

// UpgradeOrgDevices struct for UpgradeOrgDevices
type UpgradeOrgDevices struct {
	// phases for canary deployment. Each phase represents percentage of devices that need to be upgraded in that phase.
	CanaryPhases []int32 `json:"canary_phases,omitempty"`
	// whether to allow local AP-to-AP FW upgrade
	EnableP2p *bool `json:"enable_p2p,omitempty"`
	// true will force upgrade when requested version is same as running version
	Force *bool `json:"force,omitempty"`
	// percentage of failures allowed across the entire upgrade(not applicable for `big_bang`)
	MaxFailurePercentage *float32 `json:"max_failure_percentage,omitempty"`
	// number of failures allowed within each phase. Only applicable for canary. Array length should be same as `canary_phases``. Will be used if provided, else max_failure_percentage will be used.
	MaxFailures []int32 `json:"max_failures,omitempty"`
	Models []string `json:"models,omitempty"`
	P2pClusterSize *int32 `json:"p2p_cluster_size,omitempty"`
	// number of parallel p2p download batches to creat
	P2pParallelism *int32 `json:"p2p_parallelism,omitempty"`
	// Reboot device immediately after upgrade is completed (Available on Junos OS devices)
	Reboot *bool `json:"reboot,omitempty"`
	// reboot start time in epoch seconds, default is `start_time`
	RebootAt *float32 `json:"reboot_at,omitempty"`
	// percentage of AP’s that need to be present in the first rrm batch
	RrmFirstBatchPercentage *int32 `json:"rrm_first_batch_percentage,omitempty"`
	// max percentage of AP’s that need to be present in each rrm batch
	RrmMaxBatchPercentage *int32 `json:"rrm_max_batch_percentage,omitempty"`
	RrmMeshUpgrade *DeviceUpgradeRrmMeshUpgrade `json:"rrm_mesh_upgrade,omitempty"`
	RrmNodeOrder *DeviceUpgradeRrmNodeOrder `json:"rrm_node_order,omitempty"`
	// true will make rrm batch sizes slowly ramp up
	RrmSlowRamp *bool `json:"rrm_slow_ramp,omitempty"`
	SiteIds []string `json:"site_ids,omitempty"`
	// Perform recovery snapshot after device is rebooted (Available on Junos OS devices)
	Snapshot *bool `json:"snapshot,omitempty"`
	// upgrade start time in epoch seconds, default is now
	StartTime *float32 `json:"start_time,omitempty"`
	Strategy *DeviceUpgradeStrategy `json:"strategy,omitempty"`
	// specific version / stable
	Version *string `json:"version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpgradeOrgDevices UpgradeOrgDevices

// NewUpgradeOrgDevices instantiates a new UpgradeOrgDevices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpgradeOrgDevices() *UpgradeOrgDevices {
	this := UpgradeOrgDevices{}
	var force bool = false
	this.Force = &force
	var maxFailurePercentage float32 = 5
	this.MaxFailurePercentage = &maxFailurePercentage
	var p2pClusterSize int32 = 10
	this.P2pClusterSize = &p2pClusterSize
	var reboot bool = false
	this.Reboot = &reboot
	var rrmMeshUpgrade DeviceUpgradeRrmMeshUpgrade = DEVICEUPGRADERRMMESHUPGRADE_SEQUENTIAL
	this.RrmMeshUpgrade = &rrmMeshUpgrade
	var rrmNodeOrder DeviceUpgradeRrmNodeOrder = DEVICEUPGRADERRMNODEORDER_FRINGE_TO_CENTER
	this.RrmNodeOrder = &rrmNodeOrder
	var snapshot bool = false
	this.Snapshot = &snapshot
	var strategy DeviceUpgradeStrategy = DEVICEUPGRADESTRATEGY_BIG_BANG
	this.Strategy = &strategy
	var version string = "latest"
	this.Version = &version
	return &this
}

// NewUpgradeOrgDevicesWithDefaults instantiates a new UpgradeOrgDevices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpgradeOrgDevicesWithDefaults() *UpgradeOrgDevices {
	this := UpgradeOrgDevices{}
	var force bool = false
	this.Force = &force
	var maxFailurePercentage float32 = 5
	this.MaxFailurePercentage = &maxFailurePercentage
	var p2pClusterSize int32 = 10
	this.P2pClusterSize = &p2pClusterSize
	var reboot bool = false
	this.Reboot = &reboot
	var rrmMeshUpgrade DeviceUpgradeRrmMeshUpgrade = DEVICEUPGRADERRMMESHUPGRADE_SEQUENTIAL
	this.RrmMeshUpgrade = &rrmMeshUpgrade
	var rrmNodeOrder DeviceUpgradeRrmNodeOrder = DEVICEUPGRADERRMNODEORDER_FRINGE_TO_CENTER
	this.RrmNodeOrder = &rrmNodeOrder
	var snapshot bool = false
	this.Snapshot = &snapshot
	var strategy DeviceUpgradeStrategy = DEVICEUPGRADESTRATEGY_BIG_BANG
	this.Strategy = &strategy
	var version string = "latest"
	this.Version = &version
	return &this
}

// GetCanaryPhases returns the CanaryPhases field value if set, zero value otherwise.
func (o *UpgradeOrgDevices) GetCanaryPhases() []int32 {
	if o == nil || IsNil(o.CanaryPhases) {
		var ret []int32
		return ret
	}
	return o.CanaryPhases
}

// GetCanaryPhasesOk returns a tuple with the CanaryPhases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradeOrgDevices) GetCanaryPhasesOk() ([]int32, bool) {
	if o == nil || IsNil(o.CanaryPhases) {
		return nil, false
	}
	return o.CanaryPhases, true
}

// HasCanaryPhases returns a boolean if a field has been set.
func (o *UpgradeOrgDevices) HasCanaryPhases() bool {
	if o != nil && !IsNil(o.CanaryPhases) {
		return true
	}

	return false
}

// SetCanaryPhases gets a reference to the given []int32 and assigns it to the CanaryPhases field.
func (o *UpgradeOrgDevices) SetCanaryPhases(v []int32) {
	o.CanaryPhases = v
}

// GetEnableP2p returns the EnableP2p field value if set, zero value otherwise.
func (o *UpgradeOrgDevices) GetEnableP2p() bool {
	if o == nil || IsNil(o.EnableP2p) {
		var ret bool
		return ret
	}
	return *o.EnableP2p
}

// GetEnableP2pOk returns a tuple with the EnableP2p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradeOrgDevices) GetEnableP2pOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableP2p) {
		return nil, false
	}
	return o.EnableP2p, true
}

// HasEnableP2p returns a boolean if a field has been set.
func (o *UpgradeOrgDevices) HasEnableP2p() bool {
	if o != nil && !IsNil(o.EnableP2p) {
		return true
	}

	return false
}

// SetEnableP2p gets a reference to the given bool and assigns it to the EnableP2p field.
func (o *UpgradeOrgDevices) SetEnableP2p(v bool) {
	o.EnableP2p = &v
}

// GetForce returns the Force field value if set, zero value otherwise.
func (o *UpgradeOrgDevices) GetForce() bool {
	if o == nil || IsNil(o.Force) {
		var ret bool
		return ret
	}
	return *o.Force
}

// GetForceOk returns a tuple with the Force field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradeOrgDevices) GetForceOk() (*bool, bool) {
	if o == nil || IsNil(o.Force) {
		return nil, false
	}
	return o.Force, true
}

// HasForce returns a boolean if a field has been set.
func (o *UpgradeOrgDevices) HasForce() bool {
	if o != nil && !IsNil(o.Force) {
		return true
	}

	return false
}

// SetForce gets a reference to the given bool and assigns it to the Force field.
func (o *UpgradeOrgDevices) SetForce(v bool) {
	o.Force = &v
}

// GetMaxFailurePercentage returns the MaxFailurePercentage field value if set, zero value otherwise.
func (o *UpgradeOrgDevices) GetMaxFailurePercentage() float32 {
	if o == nil || IsNil(o.MaxFailurePercentage) {
		var ret float32
		return ret
	}
	return *o.MaxFailurePercentage
}

// GetMaxFailurePercentageOk returns a tuple with the MaxFailurePercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradeOrgDevices) GetMaxFailurePercentageOk() (*float32, bool) {
	if o == nil || IsNil(o.MaxFailurePercentage) {
		return nil, false
	}
	return o.MaxFailurePercentage, true
}

// HasMaxFailurePercentage returns a boolean if a field has been set.
func (o *UpgradeOrgDevices) HasMaxFailurePercentage() bool {
	if o != nil && !IsNil(o.MaxFailurePercentage) {
		return true
	}

	return false
}

// SetMaxFailurePercentage gets a reference to the given float32 and assigns it to the MaxFailurePercentage field.
func (o *UpgradeOrgDevices) SetMaxFailurePercentage(v float32) {
	o.MaxFailurePercentage = &v
}

// GetMaxFailures returns the MaxFailures field value if set, zero value otherwise.
func (o *UpgradeOrgDevices) GetMaxFailures() []int32 {
	if o == nil || IsNil(o.MaxFailures) {
		var ret []int32
		return ret
	}
	return o.MaxFailures
}

// GetMaxFailuresOk returns a tuple with the MaxFailures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradeOrgDevices) GetMaxFailuresOk() ([]int32, bool) {
	if o == nil || IsNil(o.MaxFailures) {
		return nil, false
	}
	return o.MaxFailures, true
}

// HasMaxFailures returns a boolean if a field has been set.
func (o *UpgradeOrgDevices) HasMaxFailures() bool {
	if o != nil && !IsNil(o.MaxFailures) {
		return true
	}

	return false
}

// SetMaxFailures gets a reference to the given []int32 and assigns it to the MaxFailures field.
func (o *UpgradeOrgDevices) SetMaxFailures(v []int32) {
	o.MaxFailures = v
}

// GetModels returns the Models field value if set, zero value otherwise.
func (o *UpgradeOrgDevices) GetModels() []string {
	if o == nil || IsNil(o.Models) {
		var ret []string
		return ret
	}
	return o.Models
}

// GetModelsOk returns a tuple with the Models field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradeOrgDevices) GetModelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Models) {
		return nil, false
	}
	return o.Models, true
}

// HasModels returns a boolean if a field has been set.
func (o *UpgradeOrgDevices) HasModels() bool {
	if o != nil && !IsNil(o.Models) {
		return true
	}

	return false
}

// SetModels gets a reference to the given []string and assigns it to the Models field.
func (o *UpgradeOrgDevices) SetModels(v []string) {
	o.Models = v
}

// GetP2pClusterSize returns the P2pClusterSize field value if set, zero value otherwise.
func (o *UpgradeOrgDevices) GetP2pClusterSize() int32 {
	if o == nil || IsNil(o.P2pClusterSize) {
		var ret int32
		return ret
	}
	return *o.P2pClusterSize
}

// GetP2pClusterSizeOk returns a tuple with the P2pClusterSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradeOrgDevices) GetP2pClusterSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.P2pClusterSize) {
		return nil, false
	}
	return o.P2pClusterSize, true
}

// HasP2pClusterSize returns a boolean if a field has been set.
func (o *UpgradeOrgDevices) HasP2pClusterSize() bool {
	if o != nil && !IsNil(o.P2pClusterSize) {
		return true
	}

	return false
}

// SetP2pClusterSize gets a reference to the given int32 and assigns it to the P2pClusterSize field.
func (o *UpgradeOrgDevices) SetP2pClusterSize(v int32) {
	o.P2pClusterSize = &v
}

// GetP2pParallelism returns the P2pParallelism field value if set, zero value otherwise.
func (o *UpgradeOrgDevices) GetP2pParallelism() int32 {
	if o == nil || IsNil(o.P2pParallelism) {
		var ret int32
		return ret
	}
	return *o.P2pParallelism
}

// GetP2pParallelismOk returns a tuple with the P2pParallelism field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradeOrgDevices) GetP2pParallelismOk() (*int32, bool) {
	if o == nil || IsNil(o.P2pParallelism) {
		return nil, false
	}
	return o.P2pParallelism, true
}

// HasP2pParallelism returns a boolean if a field has been set.
func (o *UpgradeOrgDevices) HasP2pParallelism() bool {
	if o != nil && !IsNil(o.P2pParallelism) {
		return true
	}

	return false
}

// SetP2pParallelism gets a reference to the given int32 and assigns it to the P2pParallelism field.
func (o *UpgradeOrgDevices) SetP2pParallelism(v int32) {
	o.P2pParallelism = &v
}

// GetReboot returns the Reboot field value if set, zero value otherwise.
func (o *UpgradeOrgDevices) GetReboot() bool {
	if o == nil || IsNil(o.Reboot) {
		var ret bool
		return ret
	}
	return *o.Reboot
}

// GetRebootOk returns a tuple with the Reboot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradeOrgDevices) GetRebootOk() (*bool, bool) {
	if o == nil || IsNil(o.Reboot) {
		return nil, false
	}
	return o.Reboot, true
}

// HasReboot returns a boolean if a field has been set.
func (o *UpgradeOrgDevices) HasReboot() bool {
	if o != nil && !IsNil(o.Reboot) {
		return true
	}

	return false
}

// SetReboot gets a reference to the given bool and assigns it to the Reboot field.
func (o *UpgradeOrgDevices) SetReboot(v bool) {
	o.Reboot = &v
}

// GetRebootAt returns the RebootAt field value if set, zero value otherwise.
func (o *UpgradeOrgDevices) GetRebootAt() float32 {
	if o == nil || IsNil(o.RebootAt) {
		var ret float32
		return ret
	}
	return *o.RebootAt
}

// GetRebootAtOk returns a tuple with the RebootAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradeOrgDevices) GetRebootAtOk() (*float32, bool) {
	if o == nil || IsNil(o.RebootAt) {
		return nil, false
	}
	return o.RebootAt, true
}

// HasRebootAt returns a boolean if a field has been set.
func (o *UpgradeOrgDevices) HasRebootAt() bool {
	if o != nil && !IsNil(o.RebootAt) {
		return true
	}

	return false
}

// SetRebootAt gets a reference to the given float32 and assigns it to the RebootAt field.
func (o *UpgradeOrgDevices) SetRebootAt(v float32) {
	o.RebootAt = &v
}

// GetRrmFirstBatchPercentage returns the RrmFirstBatchPercentage field value if set, zero value otherwise.
func (o *UpgradeOrgDevices) GetRrmFirstBatchPercentage() int32 {
	if o == nil || IsNil(o.RrmFirstBatchPercentage) {
		var ret int32
		return ret
	}
	return *o.RrmFirstBatchPercentage
}

// GetRrmFirstBatchPercentageOk returns a tuple with the RrmFirstBatchPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradeOrgDevices) GetRrmFirstBatchPercentageOk() (*int32, bool) {
	if o == nil || IsNil(o.RrmFirstBatchPercentage) {
		return nil, false
	}
	return o.RrmFirstBatchPercentage, true
}

// HasRrmFirstBatchPercentage returns a boolean if a field has been set.
func (o *UpgradeOrgDevices) HasRrmFirstBatchPercentage() bool {
	if o != nil && !IsNil(o.RrmFirstBatchPercentage) {
		return true
	}

	return false
}

// SetRrmFirstBatchPercentage gets a reference to the given int32 and assigns it to the RrmFirstBatchPercentage field.
func (o *UpgradeOrgDevices) SetRrmFirstBatchPercentage(v int32) {
	o.RrmFirstBatchPercentage = &v
}

// GetRrmMaxBatchPercentage returns the RrmMaxBatchPercentage field value if set, zero value otherwise.
func (o *UpgradeOrgDevices) GetRrmMaxBatchPercentage() int32 {
	if o == nil || IsNil(o.RrmMaxBatchPercentage) {
		var ret int32
		return ret
	}
	return *o.RrmMaxBatchPercentage
}

// GetRrmMaxBatchPercentageOk returns a tuple with the RrmMaxBatchPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradeOrgDevices) GetRrmMaxBatchPercentageOk() (*int32, bool) {
	if o == nil || IsNil(o.RrmMaxBatchPercentage) {
		return nil, false
	}
	return o.RrmMaxBatchPercentage, true
}

// HasRrmMaxBatchPercentage returns a boolean if a field has been set.
func (o *UpgradeOrgDevices) HasRrmMaxBatchPercentage() bool {
	if o != nil && !IsNil(o.RrmMaxBatchPercentage) {
		return true
	}

	return false
}

// SetRrmMaxBatchPercentage gets a reference to the given int32 and assigns it to the RrmMaxBatchPercentage field.
func (o *UpgradeOrgDevices) SetRrmMaxBatchPercentage(v int32) {
	o.RrmMaxBatchPercentage = &v
}

// GetRrmMeshUpgrade returns the RrmMeshUpgrade field value if set, zero value otherwise.
func (o *UpgradeOrgDevices) GetRrmMeshUpgrade() DeviceUpgradeRrmMeshUpgrade {
	if o == nil || IsNil(o.RrmMeshUpgrade) {
		var ret DeviceUpgradeRrmMeshUpgrade
		return ret
	}
	return *o.RrmMeshUpgrade
}

// GetRrmMeshUpgradeOk returns a tuple with the RrmMeshUpgrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradeOrgDevices) GetRrmMeshUpgradeOk() (*DeviceUpgradeRrmMeshUpgrade, bool) {
	if o == nil || IsNil(o.RrmMeshUpgrade) {
		return nil, false
	}
	return o.RrmMeshUpgrade, true
}

// HasRrmMeshUpgrade returns a boolean if a field has been set.
func (o *UpgradeOrgDevices) HasRrmMeshUpgrade() bool {
	if o != nil && !IsNil(o.RrmMeshUpgrade) {
		return true
	}

	return false
}

// SetRrmMeshUpgrade gets a reference to the given DeviceUpgradeRrmMeshUpgrade and assigns it to the RrmMeshUpgrade field.
func (o *UpgradeOrgDevices) SetRrmMeshUpgrade(v DeviceUpgradeRrmMeshUpgrade) {
	o.RrmMeshUpgrade = &v
}

// GetRrmNodeOrder returns the RrmNodeOrder field value if set, zero value otherwise.
func (o *UpgradeOrgDevices) GetRrmNodeOrder() DeviceUpgradeRrmNodeOrder {
	if o == nil || IsNil(o.RrmNodeOrder) {
		var ret DeviceUpgradeRrmNodeOrder
		return ret
	}
	return *o.RrmNodeOrder
}

// GetRrmNodeOrderOk returns a tuple with the RrmNodeOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradeOrgDevices) GetRrmNodeOrderOk() (*DeviceUpgradeRrmNodeOrder, bool) {
	if o == nil || IsNil(o.RrmNodeOrder) {
		return nil, false
	}
	return o.RrmNodeOrder, true
}

// HasRrmNodeOrder returns a boolean if a field has been set.
func (o *UpgradeOrgDevices) HasRrmNodeOrder() bool {
	if o != nil && !IsNil(o.RrmNodeOrder) {
		return true
	}

	return false
}

// SetRrmNodeOrder gets a reference to the given DeviceUpgradeRrmNodeOrder and assigns it to the RrmNodeOrder field.
func (o *UpgradeOrgDevices) SetRrmNodeOrder(v DeviceUpgradeRrmNodeOrder) {
	o.RrmNodeOrder = &v
}

// GetRrmSlowRamp returns the RrmSlowRamp field value if set, zero value otherwise.
func (o *UpgradeOrgDevices) GetRrmSlowRamp() bool {
	if o == nil || IsNil(o.RrmSlowRamp) {
		var ret bool
		return ret
	}
	return *o.RrmSlowRamp
}

// GetRrmSlowRampOk returns a tuple with the RrmSlowRamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradeOrgDevices) GetRrmSlowRampOk() (*bool, bool) {
	if o == nil || IsNil(o.RrmSlowRamp) {
		return nil, false
	}
	return o.RrmSlowRamp, true
}

// HasRrmSlowRamp returns a boolean if a field has been set.
func (o *UpgradeOrgDevices) HasRrmSlowRamp() bool {
	if o != nil && !IsNil(o.RrmSlowRamp) {
		return true
	}

	return false
}

// SetRrmSlowRamp gets a reference to the given bool and assigns it to the RrmSlowRamp field.
func (o *UpgradeOrgDevices) SetRrmSlowRamp(v bool) {
	o.RrmSlowRamp = &v
}

// GetSiteIds returns the SiteIds field value if set, zero value otherwise.
func (o *UpgradeOrgDevices) GetSiteIds() []string {
	if o == nil || IsNil(o.SiteIds) {
		var ret []string
		return ret
	}
	return o.SiteIds
}

// GetSiteIdsOk returns a tuple with the SiteIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradeOrgDevices) GetSiteIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.SiteIds) {
		return nil, false
	}
	return o.SiteIds, true
}

// HasSiteIds returns a boolean if a field has been set.
func (o *UpgradeOrgDevices) HasSiteIds() bool {
	if o != nil && !IsNil(o.SiteIds) {
		return true
	}

	return false
}

// SetSiteIds gets a reference to the given []string and assigns it to the SiteIds field.
func (o *UpgradeOrgDevices) SetSiteIds(v []string) {
	o.SiteIds = v
}

// GetSnapshot returns the Snapshot field value if set, zero value otherwise.
func (o *UpgradeOrgDevices) GetSnapshot() bool {
	if o == nil || IsNil(o.Snapshot) {
		var ret bool
		return ret
	}
	return *o.Snapshot
}

// GetSnapshotOk returns a tuple with the Snapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradeOrgDevices) GetSnapshotOk() (*bool, bool) {
	if o == nil || IsNil(o.Snapshot) {
		return nil, false
	}
	return o.Snapshot, true
}

// HasSnapshot returns a boolean if a field has been set.
func (o *UpgradeOrgDevices) HasSnapshot() bool {
	if o != nil && !IsNil(o.Snapshot) {
		return true
	}

	return false
}

// SetSnapshot gets a reference to the given bool and assigns it to the Snapshot field.
func (o *UpgradeOrgDevices) SetSnapshot(v bool) {
	o.Snapshot = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *UpgradeOrgDevices) GetStartTime() float32 {
	if o == nil || IsNil(o.StartTime) {
		var ret float32
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradeOrgDevices) GetStartTimeOk() (*float32, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *UpgradeOrgDevices) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given float32 and assigns it to the StartTime field.
func (o *UpgradeOrgDevices) SetStartTime(v float32) {
	o.StartTime = &v
}

// GetStrategy returns the Strategy field value if set, zero value otherwise.
func (o *UpgradeOrgDevices) GetStrategy() DeviceUpgradeStrategy {
	if o == nil || IsNil(o.Strategy) {
		var ret DeviceUpgradeStrategy
		return ret
	}
	return *o.Strategy
}

// GetStrategyOk returns a tuple with the Strategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradeOrgDevices) GetStrategyOk() (*DeviceUpgradeStrategy, bool) {
	if o == nil || IsNil(o.Strategy) {
		return nil, false
	}
	return o.Strategy, true
}

// HasStrategy returns a boolean if a field has been set.
func (o *UpgradeOrgDevices) HasStrategy() bool {
	if o != nil && !IsNil(o.Strategy) {
		return true
	}

	return false
}

// SetStrategy gets a reference to the given DeviceUpgradeStrategy and assigns it to the Strategy field.
func (o *UpgradeOrgDevices) SetStrategy(v DeviceUpgradeStrategy) {
	o.Strategy = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *UpgradeOrgDevices) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradeOrgDevices) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *UpgradeOrgDevices) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *UpgradeOrgDevices) SetVersion(v string) {
	o.Version = &v
}

func (o UpgradeOrgDevices) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpgradeOrgDevices) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CanaryPhases) {
		toSerialize["canary_phases"] = o.CanaryPhases
	}
	if !IsNil(o.EnableP2p) {
		toSerialize["enable_p2p"] = o.EnableP2p
	}
	if !IsNil(o.Force) {
		toSerialize["force"] = o.Force
	}
	if !IsNil(o.MaxFailurePercentage) {
		toSerialize["max_failure_percentage"] = o.MaxFailurePercentage
	}
	if !IsNil(o.MaxFailures) {
		toSerialize["max_failures"] = o.MaxFailures
	}
	if !IsNil(o.Models) {
		toSerialize["models"] = o.Models
	}
	if !IsNil(o.P2pClusterSize) {
		toSerialize["p2p_cluster_size"] = o.P2pClusterSize
	}
	if !IsNil(o.P2pParallelism) {
		toSerialize["p2p_parallelism"] = o.P2pParallelism
	}
	if !IsNil(o.Reboot) {
		toSerialize["reboot"] = o.Reboot
	}
	if !IsNil(o.RebootAt) {
		toSerialize["reboot_at"] = o.RebootAt
	}
	if !IsNil(o.RrmFirstBatchPercentage) {
		toSerialize["rrm_first_batch_percentage"] = o.RrmFirstBatchPercentage
	}
	if !IsNil(o.RrmMaxBatchPercentage) {
		toSerialize["rrm_max_batch_percentage"] = o.RrmMaxBatchPercentage
	}
	if !IsNil(o.RrmMeshUpgrade) {
		toSerialize["rrm_mesh_upgrade"] = o.RrmMeshUpgrade
	}
	if !IsNil(o.RrmNodeOrder) {
		toSerialize["rrm_node_order"] = o.RrmNodeOrder
	}
	if !IsNil(o.RrmSlowRamp) {
		toSerialize["rrm_slow_ramp"] = o.RrmSlowRamp
	}
	if !IsNil(o.SiteIds) {
		toSerialize["site_ids"] = o.SiteIds
	}
	if !IsNil(o.Snapshot) {
		toSerialize["snapshot"] = o.Snapshot
	}
	if !IsNil(o.StartTime) {
		toSerialize["start_time"] = o.StartTime
	}
	if !IsNil(o.Strategy) {
		toSerialize["strategy"] = o.Strategy
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpgradeOrgDevices) UnmarshalJSON(data []byte) (err error) {
	varUpgradeOrgDevices := _UpgradeOrgDevices{}

	err = json.Unmarshal(data, &varUpgradeOrgDevices)

	if err != nil {
		return err
	}

	*o = UpgradeOrgDevices(varUpgradeOrgDevices)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "canary_phases")
		delete(additionalProperties, "enable_p2p")
		delete(additionalProperties, "force")
		delete(additionalProperties, "max_failure_percentage")
		delete(additionalProperties, "max_failures")
		delete(additionalProperties, "models")
		delete(additionalProperties, "p2p_cluster_size")
		delete(additionalProperties, "p2p_parallelism")
		delete(additionalProperties, "reboot")
		delete(additionalProperties, "reboot_at")
		delete(additionalProperties, "rrm_first_batch_percentage")
		delete(additionalProperties, "rrm_max_batch_percentage")
		delete(additionalProperties, "rrm_mesh_upgrade")
		delete(additionalProperties, "rrm_node_order")
		delete(additionalProperties, "rrm_slow_ramp")
		delete(additionalProperties, "site_ids")
		delete(additionalProperties, "snapshot")
		delete(additionalProperties, "start_time")
		delete(additionalProperties, "strategy")
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpgradeOrgDevices struct {
	value *UpgradeOrgDevices
	isSet bool
}

func (v NullableUpgradeOrgDevices) Get() *UpgradeOrgDevices {
	return v.value
}

func (v *NullableUpgradeOrgDevices) Set(val *UpgradeOrgDevices) {
	v.value = val
	v.isSet = true
}

func (v NullableUpgradeOrgDevices) IsSet() bool {
	return v.isSet
}

func (v *NullableUpgradeOrgDevices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpgradeOrgDevices(val *UpgradeOrgDevices) *NullableUpgradeOrgDevices {
	return &NullableUpgradeOrgDevices{value: val, isSet: true}
}

func (v NullableUpgradeOrgDevices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpgradeOrgDevices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

