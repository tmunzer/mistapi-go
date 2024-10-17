package models

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"strings"
)

// ResponseDeviceUpgrade represents a ResponseDeviceUpgrade struct.
type ResponseDeviceUpgrade struct {
	Counts *ResponseDeviceUpgradeCounts `json:"counts,omitempty"`
	// current canary or rrm phase in progress
	CurrentPhase *int `json:"current_phase,omitempty"`
	// whether to allow local AP-to-AP FW upgrade
	EnableP2p *bool `json:"enable_p2p,omitempty"`
	// whether to force upgrade when requested version is same as running version
	Force *bool `json:"force,omitempty"`
	// unique id for the upgrade
	Id uuid.UUID `json:"id"`
	// percentage of failures allowed
	MaxFailurePercentage *int `json:"max_failure_percentage,omitempty"`
	// number of failures allowed within a canary phase or serial rollout
	MaxFailures []int `json:"max_failures,omitempty"`
	// reboot start time in epoch
	RebootAt *int `json:"reboot_at,omitempty"`
	// firmware download start time in epoch
	StartTime *float64 `json:"start_time,omitempty"`
	// status upgrade is in. enum: `cancelled`, `completed`, `created`, `downloaded`, `downloading`, `failed`, `upgrading`
	Status *DeviceUpgradeStatusEnum `json:"status,omitempty"`
	// enum: `big_bang` (upgrade all at once), `canary`, `rrm`, `serial` (one at a time)
	Strategy *DeviceUpgradeStrategyEnum `json:"strategy,omitempty"`
	// version to upgrade to
	TargetVersion *string `json:"target_version,omitempty"`
	// a dictionary of rrm phase number to devices part of that phase
	UpgradePlan          *interface{}   `json:"upgrade_plan,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseDeviceUpgrade.
// It customizes the JSON marshaling process for ResponseDeviceUpgrade objects.
func (r ResponseDeviceUpgrade) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseDeviceUpgrade object to a map representation for JSON marshaling.
func (r ResponseDeviceUpgrade) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, r.AdditionalProperties)
	if r.Counts != nil {
		structMap["counts"] = r.Counts.toMap()
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
	if r.UpgradePlan != nil {
		structMap["upgrade_plan"] = r.UpgradePlan
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseDeviceUpgrade.
// It customizes the JSON unmarshaling process for ResponseDeviceUpgrade objects.
func (r *ResponseDeviceUpgrade) UnmarshalJSON(input []byte) error {
	var temp tempResponseDeviceUpgrade
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "counts", "current_phase", "enable_p2p", "force", "id", "max_failure_percentage", "max_failures", "reboot_at", "start_time", "status", "strategy", "target_version", "upgrade_plan")
	if err != nil {
		return err
	}

	r.AdditionalProperties = additionalProperties
	r.Counts = temp.Counts
	r.CurrentPhase = temp.CurrentPhase
	r.EnableP2p = temp.EnableP2p
	r.Force = temp.Force
	r.Id = *temp.Id
	r.MaxFailurePercentage = temp.MaxFailurePercentage
	r.MaxFailures = temp.MaxFailures
	r.RebootAt = temp.RebootAt
	r.StartTime = temp.StartTime
	r.Status = temp.Status
	r.Strategy = temp.Strategy
	r.TargetVersion = temp.TargetVersion
	r.UpgradePlan = temp.UpgradePlan
	return nil
}

// tempResponseDeviceUpgrade is a temporary struct used for validating the fields of ResponseDeviceUpgrade.
type tempResponseDeviceUpgrade struct {
	Counts               *ResponseDeviceUpgradeCounts `json:"counts,omitempty"`
	CurrentPhase         *int                         `json:"current_phase,omitempty"`
	EnableP2p            *bool                        `json:"enable_p2p,omitempty"`
	Force                *bool                        `json:"force,omitempty"`
	Id                   *uuid.UUID                   `json:"id"`
	MaxFailurePercentage *int                         `json:"max_failure_percentage,omitempty"`
	MaxFailures          []int                        `json:"max_failures,omitempty"`
	RebootAt             *int                         `json:"reboot_at,omitempty"`
	StartTime            *float64                     `json:"start_time,omitempty"`
	Status               *DeviceUpgradeStatusEnum     `json:"status,omitempty"`
	Strategy             *DeviceUpgradeStrategyEnum   `json:"strategy,omitempty"`
	TargetVersion        *string                      `json:"target_version,omitempty"`
	UpgradePlan          *interface{}                 `json:"upgrade_plan,omitempty"`
}

func (r *tempResponseDeviceUpgrade) validate() error {
	var errs []string
	if r.Id == nil {
		errs = append(errs, "required field `id` is missing for type `response_device_upgrade`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
