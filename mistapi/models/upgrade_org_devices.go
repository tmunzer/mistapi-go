// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// UpgradeOrgDevices represents a UpgradeOrgDevices struct.
type UpgradeOrgDevices struct {
    // If `true`, will upgrade all sites in this org
    AllSites                *bool                                  `json:"all_sites,omitempty"`
    // Only if `strategy`==`canary`. Phases for canary deployment. Each phase represents percentage of devices that need to be upgraded in that phase. default is [1, 10, 50, 100]
    CanaryPhases            []int                                  `json:"canary_phases,omitempty"`
    // enum: `ap`, `gateway`, `switch`
    DeviceType              *DeviceTypeEnum                        `json:"device_type,omitempty"`
    // enum:
    // * `big_bang`: download all at once, no orchestration
    // * `serial`: one at a time'
    // * `canary`: upgrade in phases
    DownloadStrategy        *UpgradeOrgDevicesDownloadStrategyEnum `json:"download_strategy,omitempty"`
    // If `strategy`!=`big_bang`. percentage of failures allowed across the entire upgrade
    MaxFailurePercentage    *int                                   `json:"max_failure_percentage,omitempty"`
    // If `strategy`==`canary`. Number of failures allowed within each phase. Only applicable for `canary`. Array length should be same as `canary_phases`. Will be used if provided, else `max_failure_percentage` will be used
    MaxFailures             []int                                  `json:"max_failures,omitempty"`
    // Only devices of these model types will be selected for upgrade
    Models                  [][]string                             `json:"models,omitempty"`
    // For APs only and if `enable_p2p`==`true`.
    P2pClusterSize          *int                                   `json:"p2p_cluster_size,omitempty"`
    // For APs only and if `enable_p2p`==`true`. Number of parallel p2p download batches to create
    P2pParallelism          *int                                   `json:"p2p_parallelism,omitempty"`
    // For Switches and Gateways only and if `reboot`==`true`. Reboot start time in epoch seconds, default is `start_time`
    RebootAt                *int                                   `json:"reboot_at,omitempty"`                  // Deprecated
    // Process start date and time, ISO8601 format. Exclude timezone component if site local timezone needs to be used
    RebootDatetime          *string                                `json:"reboot_datetime,omitempty"`
    // enum: `big_bang` (upgrade all at once), `canary`, `rrm` (APs only), `serial` (one at a time)
    RebootStrategy          *UpgradeDeviceStrategyEnum             `json:"reboot_strategy,omitempty"`
    // For APs only and if `strategy`==`rrm`. Percentage of APs that need to be present in the first RRM batch
    RrmFirstBatchPercentage *int                                   `json:"rrm_first_batch_percentage,omitempty"`
    // For APs only and if `strategy`==`rrm`. Max percentage of APs that need to be present in each RRM batch
    RrmMaxBatchPercentage   *int                                   `json:"rrm_max_batch_percentage,omitempty"`
    // For APs only and if `strategy`==`rrm`. Whether to upgrade mesh AP’s parallelly or sequentially at the end of the upgrade. enum: `parallel`, `sequential`
    RrmMeshUpgrade          *UpgradeDeviceRrmMeshUpgradeEnum       `json:"rrm_mesh_upgrade,omitempty"`
    // For APs only and if `strategy`==`rrm`. Used in rrm to determine whether to start upgrade from fringe or center AP’s. enum: `center_to_fringe`, `fringe_to_center`
    RrmNodeOrder            *UpgradeDeviceRrmNodeOrderEnum         `json:"rrm_node_order,omitempty"`
    // For APs only and if `strategy`==`rrm`. True will make rrm batch sizes slowly ramp up
    RrmSlowRamp             *bool                                  `json:"rrm_slow_ramp,omitempty"`
    // Rules used to identify devices which will be selected for upgrade. Device will be selected as long as it satisfies any one rule
    // Property key defines the type of matching, value is the string to match. e.g:
    // * `match_name`: Device name must match the property value
    // * `match_name[0:3]`: Device name must match the first 3 letters of the property value
    // * `match_name[2:6]`: Device name must match the property value from the 2nd to the 6th letter
    // * `match_model`: Device model must match the property value
    // * `match_model[1:3]`: Device model must match the property value from the 1nd to the 3rd letter
    Rules                   []map[string]string                    `json:"rules,omitempty"`
    // Only devices belonging to these sites will be selected for upgrade. Will be ignored if `all_sites`==`true`
    SiteIds                 []uuid.UUID                            `json:"site_ids,omitempty"`
    // For Junos devices only. Perform recovery snapshot after device is rebooted
    Snapshot                *bool                                  `json:"snapshot,omitempty"`
    // Process start date and time, ISO8601 format
    StartDatetime           *string                                `json:"start_datetime,omitempty"`
    // Upgrade start time in epoch seconds, default is now
    StartTime               *int                                   `json:"start_time,omitempty"`                 // Deprecated
    // enum: `big_bang` (upgrade all at once), `canary`, `rrm` (APs only), `serial` (one at a time)
    Strategy                *UpgradeDeviceStrategyEnum             `json:"strategy,omitempty"`
    Versions                []UpgradeOrgDevicesVersion             `json:"versions,omitempty"`
    AdditionalProperties    map[string]interface{}                 `json:"_"`
}

// String implements the fmt.Stringer interface for UpgradeOrgDevices,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UpgradeOrgDevices) String() string {
    return fmt.Sprintf(
    	"UpgradeOrgDevices[AllSites=%v, CanaryPhases=%v, DeviceType=%v, DownloadStrategy=%v, MaxFailurePercentage=%v, MaxFailures=%v, Models=%v, P2pClusterSize=%v, P2pParallelism=%v, RebootAt=%v, RebootDatetime=%v, RebootStrategy=%v, RrmFirstBatchPercentage=%v, RrmMaxBatchPercentage=%v, RrmMeshUpgrade=%v, RrmNodeOrder=%v, RrmSlowRamp=%v, Rules=%v, SiteIds=%v, Snapshot=%v, StartDatetime=%v, StartTime=%v, Strategy=%v, Versions=%v, AdditionalProperties=%v]",
    	u.AllSites, u.CanaryPhases, u.DeviceType, u.DownloadStrategy, u.MaxFailurePercentage, u.MaxFailures, u.Models, u.P2pClusterSize, u.P2pParallelism, u.RebootAt, u.RebootDatetime, u.RebootStrategy, u.RrmFirstBatchPercentage, u.RrmMaxBatchPercentage, u.RrmMeshUpgrade, u.RrmNodeOrder, u.RrmSlowRamp, u.Rules, u.SiteIds, u.Snapshot, u.StartDatetime, u.StartTime, u.Strategy, u.Versions, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UpgradeOrgDevices.
// It customizes the JSON marshaling process for UpgradeOrgDevices objects.
func (u UpgradeOrgDevices) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "all_sites", "canary_phases", "device_type", "download_strategy", "max_failure_percentage", "max_failures", "models", "p2p_cluster_size", "p2p_parallelism", "reboot_at", "reboot_datetime", "reboot_strategy", "rrm_first_batch_percentage", "rrm_max_batch_percentage", "rrm_mesh_upgrade", "rrm_node_order", "rrm_slow_ramp", "rules", "site_ids", "snapshot", "start_datetime", "start_time", "strategy", "versions"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UpgradeOrgDevices object to a map representation for JSON marshaling.
func (u UpgradeOrgDevices) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
    if u.AllSites != nil {
        structMap["all_sites"] = u.AllSites
    }
    if u.CanaryPhases != nil {
        structMap["canary_phases"] = u.CanaryPhases
    }
    if u.DeviceType != nil {
        structMap["device_type"] = u.DeviceType
    }
    if u.DownloadStrategy != nil {
        structMap["download_strategy"] = u.DownloadStrategy
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
    if u.RebootAt != nil {
        structMap["reboot_at"] = u.RebootAt
    }
    if u.RebootDatetime != nil {
        structMap["reboot_datetime"] = u.RebootDatetime
    }
    if u.RebootStrategy != nil {
        structMap["reboot_strategy"] = u.RebootStrategy
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
    if u.StartDatetime != nil {
        structMap["start_datetime"] = u.StartDatetime
    }
    if u.StartTime != nil {
        structMap["start_time"] = u.StartTime
    }
    if u.Strategy != nil {
        structMap["strategy"] = u.Strategy
    }
    if u.Versions != nil {
        structMap["versions"] = u.Versions
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "all_sites", "canary_phases", "device_type", "download_strategy", "max_failure_percentage", "max_failures", "models", "p2p_cluster_size", "p2p_parallelism", "reboot_at", "reboot_datetime", "reboot_strategy", "rrm_first_batch_percentage", "rrm_max_batch_percentage", "rrm_mesh_upgrade", "rrm_node_order", "rrm_slow_ramp", "rules", "site_ids", "snapshot", "start_datetime", "start_time", "strategy", "versions")
    if err != nil {
    	return err
    }
    u.AdditionalProperties = additionalProperties
    
    u.AllSites = temp.AllSites
    u.CanaryPhases = temp.CanaryPhases
    u.DeviceType = temp.DeviceType
    u.DownloadStrategy = temp.DownloadStrategy
    u.MaxFailurePercentage = temp.MaxFailurePercentage
    u.MaxFailures = temp.MaxFailures
    u.Models = temp.Models
    u.P2pClusterSize = temp.P2pClusterSize
    u.P2pParallelism = temp.P2pParallelism
    u.RebootAt = temp.RebootAt
    u.RebootDatetime = temp.RebootDatetime
    u.RebootStrategy = temp.RebootStrategy
    u.RrmFirstBatchPercentage = temp.RrmFirstBatchPercentage
    u.RrmMaxBatchPercentage = temp.RrmMaxBatchPercentage
    u.RrmMeshUpgrade = temp.RrmMeshUpgrade
    u.RrmNodeOrder = temp.RrmNodeOrder
    u.RrmSlowRamp = temp.RrmSlowRamp
    u.Rules = temp.Rules
    u.SiteIds = temp.SiteIds
    u.Snapshot = temp.Snapshot
    u.StartDatetime = temp.StartDatetime
    u.StartTime = temp.StartTime
    u.Strategy = temp.Strategy
    u.Versions = temp.Versions
    return nil
}

// tempUpgradeOrgDevices is a temporary struct used for validating the fields of UpgradeOrgDevices.
type tempUpgradeOrgDevices  struct {
    AllSites                *bool                                  `json:"all_sites,omitempty"`
    CanaryPhases            []int                                  `json:"canary_phases,omitempty"`
    DeviceType              *DeviceTypeEnum                        `json:"device_type,omitempty"`
    DownloadStrategy        *UpgradeOrgDevicesDownloadStrategyEnum `json:"download_strategy,omitempty"`
    MaxFailurePercentage    *int                                   `json:"max_failure_percentage,omitempty"`
    MaxFailures             []int                                  `json:"max_failures,omitempty"`
    Models                  [][]string                             `json:"models,omitempty"`
    P2pClusterSize          *int                                   `json:"p2p_cluster_size,omitempty"`
    P2pParallelism          *int                                   `json:"p2p_parallelism,omitempty"`
    RebootAt                *int                                   `json:"reboot_at,omitempty"`
    RebootDatetime          *string                                `json:"reboot_datetime,omitempty"`
    RebootStrategy          *UpgradeDeviceStrategyEnum             `json:"reboot_strategy,omitempty"`
    RrmFirstBatchPercentage *int                                   `json:"rrm_first_batch_percentage,omitempty"`
    RrmMaxBatchPercentage   *int                                   `json:"rrm_max_batch_percentage,omitempty"`
    RrmMeshUpgrade          *UpgradeDeviceRrmMeshUpgradeEnum       `json:"rrm_mesh_upgrade,omitempty"`
    RrmNodeOrder            *UpgradeDeviceRrmNodeOrderEnum         `json:"rrm_node_order,omitempty"`
    RrmSlowRamp             *bool                                  `json:"rrm_slow_ramp,omitempty"`
    Rules                   []map[string]string                    `json:"rules,omitempty"`
    SiteIds                 []uuid.UUID                            `json:"site_ids,omitempty"`
    Snapshot                *bool                                  `json:"snapshot,omitempty"`
    StartDatetime           *string                                `json:"start_datetime,omitempty"`
    StartTime               *int                                   `json:"start_time,omitempty"`
    Strategy                *UpgradeDeviceStrategyEnum             `json:"strategy,omitempty"`
    Versions                []UpgradeOrgDevicesVersion             `json:"versions,omitempty"`
}
