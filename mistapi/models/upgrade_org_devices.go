package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// UpgradeOrgDevices represents a UpgradeOrgDevices struct.
type UpgradeOrgDevices struct {
    // Only if `strategy`==`canary`. Phases for canary deployment. Each phase represents percentage of devices that need to be upgraded in that phase. default is [1, 10, 50, 100]
    CanaryPhases            []int                            `json:"canary_phases,omitempty"`
    // for APs only. whether to allow local AP-to-AP FW upgrade
    EnableP2p               *bool                            `json:"enable_p2p,omitempty"`
    // true will force upgrade when requested version is same as running version
    Force                   *bool                            `json:"force,omitempty"`
    // for APs only and if `strategy`!=`big_bang`. percentage of failures allowed across the entire upgrade
    MaxFailurePercentage    *int                             `json:"max_failure_percentage,omitempty"`
    // if `strategy`==`canary`. Number of failures allowed within each phase. Only applicable for `canary`. Array length should be same as `canary_phases`. Will be used if provided, else `max_failure_percentage` will be used
    MaxFailures             []int                            `json:"max_failures,omitempty"`
    // models which will be selected for upgrade
    Models                  []string                         `json:"models,omitempty"`
    // For APs only and if `enable_p2p`==`true`.
    P2pClusterSize          *int                             `json:"p2p_cluster_size,omitempty"`
    // For APs only and if `enable_p2p`==`true`. Number of parallel p2p download batches to create
    P2pParallelism          *int                             `json:"p2p_parallelism,omitempty"`
    // For Switches and Gateways only (APs are automatically rebooted). Reboot device immediately after upgrade is completed
    Reboot                  *bool                            `json:"reboot,omitempty"`
    // For Switches and Gateways only and if `reboot`==`true`. Reboot start time in epoch seconds, default is `start_time`
    RebootAt                *int                             `json:"reboot_at,omitempty"`
    // For APs only and if `strategy`==`rrm`. Percentage of APs that need to be present in the first RRM batch
    RrmFirstBatchPercentage *int                             `json:"rrm_first_batch_percentage,omitempty"`
    // For APs only and if `strategy`==`rrm`. Max percentage of APs that need to be present in each RRM batch
    RrmMaxBatchPercentage   *int                             `json:"rrm_max_batch_percentage,omitempty"`
    // For APs only and if `strategy`==`rrm`. Whether to upgrade mesh AP’s parallelly or sequentially at the end of the upgrade. enum: `parallel`, `sequential`
    RrmMeshUpgrade          *UpgradeDeviceRrmMeshUpgradeEnum `json:"rrm_mesh_upgrade,omitempty"`
    // For APs only and if `strategy`==`rrm`. Used in rrm to determine whether to start upgrade from fringe or center AP’s. enum: `center_to_fringe`, `fringe_to_center`
    RrmNodeOrder            *UpgradeDeviceRrmNodeOrderEnum   `json:"rrm_node_order,omitempty"`
    // For APs only and if `strategy`==`rrm`. True will make rrm batch sizes slowly ramp up
    RrmSlowRamp             *bool                            `json:"rrm_slow_ramp,omitempty"`
    // Rules used to identify devices which will be selected for upgrade. Device will be selected as long as it satisfies any one rule
    // Property key defines the type of matching, value is the string to match. e.g:
    // * `match_name`: Device name must match the property value
    // * `match_name[0:3]`: Device name must match the first 3 letters of the property value
    // * `match_name[2:6]`: Device name must match the property value from the 2nd to the 6th letter
    // * `match_model`: Device model must match the property value
    // * `match_model[1:3]`: Device model must match the property value from the 1nd to the 3rd letter
    Rules                   []map[string]string              `json:"rules,omitempty"`
    SiteIds                 []uuid.UUID                      `json:"site_ids,omitempty"`
    // For Junos devices only. Perform recovery snapshot after device is rebooted
    Snapshot                *bool                            `json:"snapshot,omitempty"`
    // upgrade start time in epoch seconds, default is now
    StartTime               *int                             `json:"start_time,omitempty"`
    // For APs only. enum: `big_bang` (upgrade all at once), `canary`, `rrm`, `serial` (one at a time)
    Strategy                *UpgradeDeviceStrategyEnum       `json:"strategy,omitempty"`
    // specific version / stable, default is to use the lastest available version
    Version                 *string                          `json:"version,omitempty"`
    AdditionalProperties    map[string]interface{}           `json:"_"`
}

// String implements the fmt.Stringer interface for UpgradeOrgDevices,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UpgradeOrgDevices) String() string {
    return fmt.Sprintf(
    	"UpgradeOrgDevices[CanaryPhases=%v, EnableP2p=%v, Force=%v, MaxFailurePercentage=%v, MaxFailures=%v, Models=%v, P2pClusterSize=%v, P2pParallelism=%v, Reboot=%v, RebootAt=%v, RrmFirstBatchPercentage=%v, RrmMaxBatchPercentage=%v, RrmMeshUpgrade=%v, RrmNodeOrder=%v, RrmSlowRamp=%v, Rules=%v, SiteIds=%v, Snapshot=%v, StartTime=%v, Strategy=%v, Version=%v, AdditionalProperties=%v]",
    	u.CanaryPhases, u.EnableP2p, u.Force, u.MaxFailurePercentage, u.MaxFailures, u.Models, u.P2pClusterSize, u.P2pParallelism, u.Reboot, u.RebootAt, u.RrmFirstBatchPercentage, u.RrmMaxBatchPercentage, u.RrmMeshUpgrade, u.RrmNodeOrder, u.RrmSlowRamp, u.Rules, u.SiteIds, u.Snapshot, u.StartTime, u.Strategy, u.Version, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UpgradeOrgDevices.
// It customizes the JSON marshaling process for UpgradeOrgDevices objects.
func (u UpgradeOrgDevices) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "canary_phases", "enable_p2p", "force", "max_failure_percentage", "max_failures", "models", "p2p_cluster_size", "p2p_parallelism", "reboot", "reboot_at", "rrm_first_batch_percentage", "rrm_max_batch_percentage", "rrm_mesh_upgrade", "rrm_node_order", "rrm_slow_ramp", "rules", "site_ids", "snapshot", "start_time", "strategy", "version"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UpgradeOrgDevices object to a map representation for JSON marshaling.
func (u UpgradeOrgDevices) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
    if u.CanaryPhases != nil {
        structMap["canary_phases"] = u.CanaryPhases
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
    if u.SiteIds != nil {
        structMap["site_ids"] = u.SiteIds
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

// UnmarshalJSON implements the json.Unmarshaler interface for UpgradeOrgDevices.
// It customizes the JSON unmarshaling process for UpgradeOrgDevices objects.
func (u *UpgradeOrgDevices) UnmarshalJSON(input []byte) error {
    var temp tempUpgradeOrgDevices
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "canary_phases", "enable_p2p", "force", "max_failure_percentage", "max_failures", "models", "p2p_cluster_size", "p2p_parallelism", "reboot", "reboot_at", "rrm_first_batch_percentage", "rrm_max_batch_percentage", "rrm_mesh_upgrade", "rrm_node_order", "rrm_slow_ramp", "rules", "site_ids", "snapshot", "start_time", "strategy", "version")
    if err != nil {
    	return err
    }
    u.AdditionalProperties = additionalProperties
    
    u.CanaryPhases = temp.CanaryPhases
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
    u.SiteIds = temp.SiteIds
    u.Snapshot = temp.Snapshot
    u.StartTime = temp.StartTime
    u.Strategy = temp.Strategy
    u.Version = temp.Version
    return nil
}

// tempUpgradeOrgDevices is a temporary struct used for validating the fields of UpgradeOrgDevices.
type tempUpgradeOrgDevices  struct {
    CanaryPhases            []int                            `json:"canary_phases,omitempty"`
    EnableP2p               *bool                            `json:"enable_p2p,omitempty"`
    Force                   *bool                            `json:"force,omitempty"`
    MaxFailurePercentage    *int                             `json:"max_failure_percentage,omitempty"`
    MaxFailures             []int                            `json:"max_failures,omitempty"`
    Models                  []string                         `json:"models,omitempty"`
    P2pClusterSize          *int                             `json:"p2p_cluster_size,omitempty"`
    P2pParallelism          *int                             `json:"p2p_parallelism,omitempty"`
    Reboot                  *bool                            `json:"reboot,omitempty"`
    RebootAt                *int                             `json:"reboot_at,omitempty"`
    RrmFirstBatchPercentage *int                             `json:"rrm_first_batch_percentage,omitempty"`
    RrmMaxBatchPercentage   *int                             `json:"rrm_max_batch_percentage,omitempty"`
    RrmMeshUpgrade          *UpgradeDeviceRrmMeshUpgradeEnum `json:"rrm_mesh_upgrade,omitempty"`
    RrmNodeOrder            *UpgradeDeviceRrmNodeOrderEnum   `json:"rrm_node_order,omitempty"`
    RrmSlowRamp             *bool                            `json:"rrm_slow_ramp,omitempty"`
    Rules                   []map[string]string              `json:"rules,omitempty"`
    SiteIds                 []uuid.UUID                      `json:"site_ids,omitempty"`
    Snapshot                *bool                            `json:"snapshot,omitempty"`
    StartTime               *int                             `json:"start_time,omitempty"`
    Strategy                *UpgradeDeviceStrategyEnum       `json:"strategy,omitempty"`
    Version                 *string                          `json:"version,omitempty"`
}
