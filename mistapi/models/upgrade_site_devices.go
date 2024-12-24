package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// UpgradeSiteDevices represents a UpgradeSiteDevices struct.
type UpgradeSiteDevices struct {
    // phases for canary deployment. Each phase represents percentage of devices that need to be upgraded in that phase. default is [1, 10, 50, 100]
    CanaryPhases            []int                               `json:"canary_phases,omitempty"`
    // id’s of devices which will be selected for upgrade
    DeviceIds               []uuid.UUID                         `json:"device_ids,omitempty"`
    // whether to allow local AP-to-AP FW upgrade
    EnableP2p               *bool                               `json:"enable_p2p,omitempty"`
    // true will force upgrade when requested version is same as running version
    Force                   *bool                               `json:"force,omitempty"`
    // percentage of failures allowed across the entire upgrade(not applicable for `big_bang`)
    MaxFailurePercentage    *float64                            `json:"max_failure_percentage,omitempty"`
    // number of failures allowed within each phase(applicable for `canary` or `rrm`). Will be used if provided, else max_failure_percentage will be used
    MaxFailures             []int                               `json:"max_failures,omitempty"`
    // models which will be selected for upgrade
    Models                  []string                            `json:"models,omitempty"`
    P2pClusterSize          *int                                `json:"p2p_cluster_size,omitempty"`
    // number of parallel p2p download batches to creat
    P2pParallelism          *int                                `json:"p2p_parallelism,omitempty"`
    // Reboot device immediately after upgrade is completed (Available on Junos OS devices)
    Reboot                  *bool                               `json:"reboot,omitempty"`
    // reboot start time in epoch seconds, default is `start_time`
    RebootAt                *float64                            `json:"reboot_at,omitempty"`
    // percentage of AP’s that need to be present in the first rrm batch
    RrmFirstBatchPercentage *int                                `json:"rrm_first_batch_percentage,omitempty"`
    // max percentage of AP’s that need to be present in each rrm batch
    RrmMaxBatchPercentage   *int                                `json:"rrm_max_batch_percentage,omitempty"`
    // sequential or parallel (default parallel). Whether to upgrade mesh AP’s parallelly or sequentially at the end of the upgrade
    RrmMeshUpgrade          *string                             `json:"rrm_mesh_upgrade,omitempty"`
    // Used in rrm to determine whether to start upgrade from fringe or center AP’s. enum: `center_to_fringe`, `fringe_to_center`
    RrmNodeOrder            *UpgradeSiteDevicesRrmNodeOrderEnum `json:"rrm_node_order,omitempty"`
    // true will make rrm batch sizes slowly ramp up
    RrmSlowRamp             *bool                               `json:"rrm_slow_ramp,omitempty"`
    // rules used to identify devices which will be selected for upgrade. Device will be selected as long as it satisfies any one rule
    Rules                   []string                            `json:"rules,omitempty"`
    // Perform recovery snapshot after device is rebooted (Available on Junos OS devices)
    Snapshot                *bool                               `json:"snapshot,omitempty"`
    // upgrade start time in epoch seconds, default is now
    StartTime               *float64                            `json:"start_time,omitempty"`
    // enum: `big_bang` (upgrade all at once), `canary`, `rrm`, `serial` (one at a time)
    Strategy                *DeviceUpgradeStrategyEnum          `json:"strategy,omitempty"`
    // specific version / stable
    Version                 *string                             `json:"version,omitempty"`
    AdditionalProperties    map[string]interface{}              `json:"_"`
}

// String implements the fmt.Stringer interface for UpgradeSiteDevices,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UpgradeSiteDevices) String() string {
    return fmt.Sprintf(
    	"UpgradeSiteDevices[CanaryPhases=%v, DeviceIds=%v, EnableP2p=%v, Force=%v, MaxFailurePercentage=%v, MaxFailures=%v, Models=%v, P2pClusterSize=%v, P2pParallelism=%v, Reboot=%v, RebootAt=%v, RrmFirstBatchPercentage=%v, RrmMaxBatchPercentage=%v, RrmMeshUpgrade=%v, RrmNodeOrder=%v, RrmSlowRamp=%v, Rules=%v, Snapshot=%v, StartTime=%v, Strategy=%v, Version=%v, AdditionalProperties=%v]",
    	u.CanaryPhases, u.DeviceIds, u.EnableP2p, u.Force, u.MaxFailurePercentage, u.MaxFailures, u.Models, u.P2pClusterSize, u.P2pParallelism, u.Reboot, u.RebootAt, u.RrmFirstBatchPercentage, u.RrmMaxBatchPercentage, u.RrmMeshUpgrade, u.RrmNodeOrder, u.RrmSlowRamp, u.Rules, u.Snapshot, u.StartTime, u.Strategy, u.Version, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UpgradeSiteDevices.
// It customizes the JSON marshaling process for UpgradeSiteDevices objects.
func (u UpgradeSiteDevices) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "canary_phases", "device_ids", "enable_p2p", "force", "max_failure_percentage", "max_failures", "models", "p2p_cluster_size", "p2p_parallelism", "reboot", "reboot_at", "rrm_first_batch_percentage", "rrm_max_batch_percentage", "rrm_mesh_upgrade", "rrm_node_order", "rrm_slow_ramp", "rules", "snapshot", "start_time", "strategy", "version"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UpgradeSiteDevices object to a map representation for JSON marshaling.
func (u UpgradeSiteDevices) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
    if u.CanaryPhases != nil {
        structMap["canary_phases"] = u.CanaryPhases
    }
    if u.DeviceIds != nil {
        structMap["device_ids"] = u.DeviceIds
    }
    if u.EnableP2p != nil {
        structMap["enable_p2p"] = u.EnableP2p
    }
    if u.Force != nil {
        structMap["force"] = u.Force
    }
    if u.MaxFailurePercentage != nil {
        structMap["max_failure_percentage"] = u.MaxFailurePercentage
    }
    if u.MaxFailures != nil {
        structMap["max_failures"] = u.MaxFailures
    }
    if u.Models != nil {
        structMap["models"] = u.Models
    }
    if u.P2pClusterSize != nil {
        structMap["p2p_cluster_size"] = u.P2pClusterSize
    }
    if u.P2pParallelism != nil {
        structMap["p2p_parallelism"] = u.P2pParallelism
    }
    if u.Reboot != nil {
        structMap["reboot"] = u.Reboot
    }
    if u.RebootAt != nil {
        structMap["reboot_at"] = u.RebootAt
    }
    if u.RrmFirstBatchPercentage != nil {
        structMap["rrm_first_batch_percentage"] = u.RrmFirstBatchPercentage
    }
    if u.RrmMaxBatchPercentage != nil {
        structMap["rrm_max_batch_percentage"] = u.RrmMaxBatchPercentage
    }
    if u.RrmMeshUpgrade != nil {
        structMap["rrm_mesh_upgrade"] = u.RrmMeshUpgrade
    }
    if u.RrmNodeOrder != nil {
        structMap["rrm_node_order"] = u.RrmNodeOrder
    }
    if u.RrmSlowRamp != nil {
        structMap["rrm_slow_ramp"] = u.RrmSlowRamp
    }
    if u.Rules != nil {
        structMap["rules"] = u.Rules
    }
    if u.Snapshot != nil {
        structMap["snapshot"] = u.Snapshot
    }
    if u.StartTime != nil {
        structMap["start_time"] = u.StartTime
    }
    if u.Strategy != nil {
        structMap["strategy"] = u.Strategy
    }
    if u.Version != nil {
        structMap["version"] = u.Version
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpgradeSiteDevices.
// It customizes the JSON unmarshaling process for UpgradeSiteDevices objects.
func (u *UpgradeSiteDevices) UnmarshalJSON(input []byte) error {
    var temp tempUpgradeSiteDevices
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "canary_phases", "device_ids", "enable_p2p", "force", "max_failure_percentage", "max_failures", "models", "p2p_cluster_size", "p2p_parallelism", "reboot", "reboot_at", "rrm_first_batch_percentage", "rrm_max_batch_percentage", "rrm_mesh_upgrade", "rrm_node_order", "rrm_slow_ramp", "rules", "snapshot", "start_time", "strategy", "version")
    if err != nil {
    	return err
    }
    u.AdditionalProperties = additionalProperties
    
    u.CanaryPhases = temp.CanaryPhases
    u.DeviceIds = temp.DeviceIds
    u.EnableP2p = temp.EnableP2p
    u.Force = temp.Force
    u.MaxFailurePercentage = temp.MaxFailurePercentage
    u.MaxFailures = temp.MaxFailures
    u.Models = temp.Models
    u.P2pClusterSize = temp.P2pClusterSize
    u.P2pParallelism = temp.P2pParallelism
    u.Reboot = temp.Reboot
    u.RebootAt = temp.RebootAt
    u.RrmFirstBatchPercentage = temp.RrmFirstBatchPercentage
    u.RrmMaxBatchPercentage = temp.RrmMaxBatchPercentage
    u.RrmMeshUpgrade = temp.RrmMeshUpgrade
    u.RrmNodeOrder = temp.RrmNodeOrder
    u.RrmSlowRamp = temp.RrmSlowRamp
    u.Rules = temp.Rules
    u.Snapshot = temp.Snapshot
    u.StartTime = temp.StartTime
    u.Strategy = temp.Strategy
    u.Version = temp.Version
    return nil
}

// tempUpgradeSiteDevices is a temporary struct used for validating the fields of UpgradeSiteDevices.
type tempUpgradeSiteDevices  struct {
    CanaryPhases            []int                               `json:"canary_phases,omitempty"`
    DeviceIds               []uuid.UUID                         `json:"device_ids,omitempty"`
    EnableP2p               *bool                               `json:"enable_p2p,omitempty"`
    Force                   *bool                               `json:"force,omitempty"`
    MaxFailurePercentage    *float64                            `json:"max_failure_percentage,omitempty"`
    MaxFailures             []int                               `json:"max_failures,omitempty"`
    Models                  []string                            `json:"models,omitempty"`
    P2pClusterSize          *int                                `json:"p2p_cluster_size,omitempty"`
    P2pParallelism          *int                                `json:"p2p_parallelism,omitempty"`
    Reboot                  *bool                               `json:"reboot,omitempty"`
    RebootAt                *float64                            `json:"reboot_at,omitempty"`
    RrmFirstBatchPercentage *int                                `json:"rrm_first_batch_percentage,omitempty"`
    RrmMaxBatchPercentage   *int                                `json:"rrm_max_batch_percentage,omitempty"`
    RrmMeshUpgrade          *string                             `json:"rrm_mesh_upgrade,omitempty"`
    RrmNodeOrder            *UpgradeSiteDevicesRrmNodeOrderEnum `json:"rrm_node_order,omitempty"`
    RrmSlowRamp             *bool                               `json:"rrm_slow_ramp,omitempty"`
    Rules                   []string                            `json:"rules,omitempty"`
    Snapshot                *bool                               `json:"snapshot,omitempty"`
    StartTime               *float64                            `json:"start_time,omitempty"`
    Strategy                *DeviceUpgradeStrategyEnum          `json:"strategy,omitempty"`
    Version                 *string                             `json:"version,omitempty"`
}
