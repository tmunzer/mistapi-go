// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"strings"
)

// ResponseSiteDeviceUpgrade represents a ResponseSiteDeviceUpgrade struct.
type ResponseSiteDeviceUpgrade struct {
	// phases for canary deployment. Each phase represents percentage of devices that need to be upgraded in that phase.
	CanaryPhases []int `json:"canary_phases,omitempty"`
	// Current canary or rrm phase in progress
	CurrentPhase *int `json:"current_phase,omitempty"`
	// Whether to allow local AP-to-AP FW upgrade
	EnableP2p *bool `json:"enable_p2p,omitempty"`
	// Whether to force upgrade when requested version is same as running version
	Force *bool `json:"force,omitempty"`
	// Unique ID of the object instance in the Mist Organization
	Id uuid.UUID `json:"id"`
	// Percentage of failures allowed
	MaxFailurePercentage *int `json:"max_failure_percentage,omitempty"`
	// If `strategy`==`canary`. Number of failures allowed within each phase. Only applicable for `canary`. Array length should be same as `canary_phases`. Will be used if provided, else `max_failure_percentage` will be used
	MaxFailures []int `json:"max_failures,omitempty"`
	// size to split the devices for p2p
	P2pClusterSize *int `json:"p2p_cluster_size,omitempty"`
	// number of parallel p2p download batches to create
	P2pParallelism *int `json:"p2p_parallelism,omitempty"`
	// reboot start time in epoch
	RebootAt *int `json:"reboot_at,omitempty"`
	// Firmware download start time in epoch
	StartTime *int `json:"start_time,omitempty"`
	// status upgrade is in. enum: `cancelled`, `completed`, `created`, `downloaded`, `downloading`, `failed`, `upgrading`, `queued`
	Status *UpgradeDeviceStatusEnum `json:"status,omitempty"`
	// enum: `big_bang` (upgrade all at once), `canary`, `rrm` (APs only), `serial` (one at a time)
	Strategy *UpgradeDeviceStrategyEnum `json:"strategy,omitempty"`
	// Version to upgrade to
	TargetVersion *string                `json:"target_version,omitempty"`
	Targets       *UpgradeDevicesTargets `json:"targets,omitempty"`
	// If `strategy`!=`big_bang`, a dictionary of phase number to devices part of that phase
	UpgradePlan          map[string][]string    `json:"upgrade_plan,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseSiteDeviceUpgrade,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseSiteDeviceUpgrade) String() string {
	return fmt.Sprintf(
		"ResponseSiteDeviceUpgrade[CanaryPhases=%v, CurrentPhase=%v, EnableP2p=%v, Force=%v, Id=%v, MaxFailurePercentage=%v, MaxFailures=%v, P2pClusterSize=%v, P2pParallelism=%v, RebootAt=%v, StartTime=%v, Status=%v, Strategy=%v, TargetVersion=%v, Targets=%v, UpgradePlan=%v, AdditionalProperties=%v]",
		r.CanaryPhases, r.CurrentPhase, r.EnableP2p, r.Force, r.Id, r.MaxFailurePercentage, r.MaxFailures, r.P2pClusterSize, r.P2pParallelism, r.RebootAt, r.StartTime, r.Status, r.Strategy, r.TargetVersion, r.Targets, r.UpgradePlan, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseSiteDeviceUpgrade.
// It customizes the JSON marshaling process for ResponseSiteDeviceUpgrade objects.
func (r ResponseSiteDeviceUpgrade) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"canary_phases", "current_phase", "enable_p2p", "force", "id", "max_failure_percentage", "max_failures", "p2p_cluster_size", "p2p_parallelism", "reboot_at", "start_time", "status", "strategy", "target_version", "targets", "upgrade_plan"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseSiteDeviceUpgrade object to a map representation for JSON marshaling.
func (r ResponseSiteDeviceUpgrade) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, r.AdditionalProperties)
	if r.CanaryPhases != nil {
		structMap["canary_phases"] = r.CanaryPhases
	}
	if r.CurrentPhase != nil {
		structMap["current_phase"] = r.CurrentPhase
	}
	if r.EnableP2p != nil {
		structMap["enable_p2p"] = r.EnableP2p
	}
	if r.Force != nil {
		structMap["force"] = r.Force
	}
	structMap["id"] = r.Id
	if r.MaxFailurePercentage != nil {
		structMap["max_failure_percentage"] = r.MaxFailurePercentage
	}
	if r.MaxFailures != nil {
		structMap["max_failures"] = r.MaxFailures
	}
	if r.P2pClusterSize != nil {
		structMap["p2p_cluster_size"] = r.P2pClusterSize
	}
	if r.P2pParallelism != nil {
		structMap["p2p_parallelism"] = r.P2pParallelism
	}
	if r.RebootAt != nil {
		structMap["reboot_at"] = r.RebootAt
	}
	if r.StartTime != nil {
		structMap["start_time"] = r.StartTime
	}
	if r.Status != nil {
		structMap["status"] = r.Status
	}
	if r.Strategy != nil {
		structMap["strategy"] = r.Strategy
	}
	if r.TargetVersion != nil {
		structMap["target_version"] = r.TargetVersion
	}
	if r.Targets != nil {
		structMap["targets"] = r.Targets.toMap()
	}
	if r.UpgradePlan != nil {
		structMap["upgrade_plan"] = r.UpgradePlan
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseSiteDeviceUpgrade.
// It customizes the JSON unmarshaling process for ResponseSiteDeviceUpgrade objects.
func (r *ResponseSiteDeviceUpgrade) UnmarshalJSON(input []byte) error {
	var temp tempResponseSiteDeviceUpgrade
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "canary_phases", "current_phase", "enable_p2p", "force", "id", "max_failure_percentage", "max_failures", "p2p_cluster_size", "p2p_parallelism", "reboot_at", "start_time", "status", "strategy", "target_version", "targets", "upgrade_plan")
	if err != nil {
		return err
	}
	r.AdditionalProperties = additionalProperties

	r.CanaryPhases = temp.CanaryPhases
	r.CurrentPhase = temp.CurrentPhase
	r.EnableP2p = temp.EnableP2p
	r.Force = temp.Force
	r.Id = *temp.Id
	r.MaxFailurePercentage = temp.MaxFailurePercentage
	r.MaxFailures = temp.MaxFailures
	r.P2pClusterSize = temp.P2pClusterSize
	r.P2pParallelism = temp.P2pParallelism
	r.RebootAt = temp.RebootAt
	r.StartTime = temp.StartTime
	r.Status = temp.Status
	r.Strategy = temp.Strategy
	r.TargetVersion = temp.TargetVersion
	r.Targets = temp.Targets
	r.UpgradePlan = temp.UpgradePlan
	return nil
}

// tempResponseSiteDeviceUpgrade is a temporary struct used for validating the fields of ResponseSiteDeviceUpgrade.
type tempResponseSiteDeviceUpgrade struct {
	CanaryPhases         []int                      `json:"canary_phases,omitempty"`
	CurrentPhase         *int                       `json:"current_phase,omitempty"`
	EnableP2p            *bool                      `json:"enable_p2p,omitempty"`
	Force                *bool                      `json:"force,omitempty"`
	Id                   *uuid.UUID                 `json:"id"`
	MaxFailurePercentage *int                       `json:"max_failure_percentage,omitempty"`
	MaxFailures          []int                      `json:"max_failures,omitempty"`
	P2pClusterSize       *int                       `json:"p2p_cluster_size,omitempty"`
	P2pParallelism       *int                       `json:"p2p_parallelism,omitempty"`
	RebootAt             *int                       `json:"reboot_at,omitempty"`
	StartTime            *int                       `json:"start_time,omitempty"`
	Status               *UpgradeDeviceStatusEnum   `json:"status,omitempty"`
	Strategy             *UpgradeDeviceStrategyEnum `json:"strategy,omitempty"`
	TargetVersion        *string                    `json:"target_version,omitempty"`
	Targets              *UpgradeDevicesTargets     `json:"targets,omitempty"`
	UpgradePlan          map[string][]string        `json:"upgrade_plan,omitempty"`
}

func (r *tempResponseSiteDeviceUpgrade) validate() error {
	var errs []string
	if r.Id == nil {
		errs = append(errs, "required field `id` is missing for type `response_site_device_upgrade`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
