package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// UpgradeOrgDeviceUpgrade represents a UpgradeOrgDeviceUpgrade struct.
type UpgradeOrgDeviceUpgrade struct {
    // Unique ID of the object instance in the Mist Organnization
    Id                   *uuid.UUID               `json:"id,omitempty"`
    StartTime            *float64                 `json:"start_time,omitempty"`
    // status upgrade is in. enum: `cancelled`, `completed`, `created`, `downloaded`, `downloading`, `failed`, `upgrading`
    Status               *DeviceUpgradeStatusEnum `json:"status,omitempty"`
    Targets              *UpgradeOrgDeviceTargets `json:"targets,omitempty"`
    AdditionalProperties map[string]interface{}   `json:"_"`
}

// String implements the fmt.Stringer interface for UpgradeOrgDeviceUpgrade,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UpgradeOrgDeviceUpgrade) String() string {
    return fmt.Sprintf(
    	"UpgradeOrgDeviceUpgrade[Id=%v, StartTime=%v, Status=%v, Targets=%v, AdditionalProperties=%v]",
    	u.Id, u.StartTime, u.Status, u.Targets, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UpgradeOrgDeviceUpgrade.
// It customizes the JSON marshaling process for UpgradeOrgDeviceUpgrade objects.
func (u UpgradeOrgDeviceUpgrade) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "id", "start_time", "status", "targets"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UpgradeOrgDeviceUpgrade object to a map representation for JSON marshaling.
func (u UpgradeOrgDeviceUpgrade) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
    if u.Id != nil {
        structMap["id"] = u.Id
    }
    if u.StartTime != nil {
        structMap["start_time"] = u.StartTime
    }
    if u.Status != nil {
        structMap["status"] = u.Status
    }
    if u.Targets != nil {
        structMap["targets"] = u.Targets.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpgradeOrgDeviceUpgrade.
// It customizes the JSON unmarshaling process for UpgradeOrgDeviceUpgrade objects.
func (u *UpgradeOrgDeviceUpgrade) UnmarshalJSON(input []byte) error {
    var temp tempUpgradeOrgDeviceUpgrade
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "start_time", "status", "targets")
    if err != nil {
    	return err
    }
    u.AdditionalProperties = additionalProperties
    
    u.Id = temp.Id
    u.StartTime = temp.StartTime
    u.Status = temp.Status
    u.Targets = temp.Targets
    return nil
}

// tempUpgradeOrgDeviceUpgrade is a temporary struct used for validating the fields of UpgradeOrgDeviceUpgrade.
type tempUpgradeOrgDeviceUpgrade  struct {
    Id        *uuid.UUID               `json:"id,omitempty"`
    StartTime *float64                 `json:"start_time,omitempty"`
    Status    *DeviceUpgradeStatusEnum `json:"status,omitempty"`
    Targets   *UpgradeOrgDeviceTargets `json:"targets,omitempty"`
}
