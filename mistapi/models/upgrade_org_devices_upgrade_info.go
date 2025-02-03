package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// UpgradeOrgDevicesUpgradeInfo represents a UpgradeOrgDevicesUpgradeInfo struct.
type UpgradeOrgDevicesUpgradeInfo struct {
    // Unique ID of the object instance in the Mist Organnization
    Id                   *uuid.UUID               `json:"id,omitempty"`
    StartTime            *int                     `json:"start_time,omitempty"`
    // status upgrade is in. enum: `cancelled`, `completed`, `created`, `downloaded`, `downloading`, `failed`, `upgrading`, `queued`
    Status               *UpgradeDeviceStatusEnum `json:"status,omitempty"`
    Targets              *UpgradeDevicesTargetIds `json:"targets,omitempty"`
    AdditionalProperties map[string]interface{}   `json:"_"`
}

// String implements the fmt.Stringer interface for UpgradeOrgDevicesUpgradeInfo,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UpgradeOrgDevicesUpgradeInfo) String() string {
    return fmt.Sprintf(
    	"UpgradeOrgDevicesUpgradeInfo[Id=%v, StartTime=%v, Status=%v, Targets=%v, AdditionalProperties=%v]",
    	u.Id, u.StartTime, u.Status, u.Targets, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UpgradeOrgDevicesUpgradeInfo.
// It customizes the JSON marshaling process for UpgradeOrgDevicesUpgradeInfo objects.
func (u UpgradeOrgDevicesUpgradeInfo) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "id", "start_time", "status", "targets"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UpgradeOrgDevicesUpgradeInfo object to a map representation for JSON marshaling.
func (u UpgradeOrgDevicesUpgradeInfo) toMap() map[string]any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for UpgradeOrgDevicesUpgradeInfo.
// It customizes the JSON unmarshaling process for UpgradeOrgDevicesUpgradeInfo objects.
func (u *UpgradeOrgDevicesUpgradeInfo) UnmarshalJSON(input []byte) error {
    var temp tempUpgradeOrgDevicesUpgradeInfo
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

// tempUpgradeOrgDevicesUpgradeInfo is a temporary struct used for validating the fields of UpgradeOrgDevicesUpgradeInfo.
type tempUpgradeOrgDevicesUpgradeInfo  struct {
    Id        *uuid.UUID               `json:"id,omitempty"`
    StartTime *int                     `json:"start_time,omitempty"`
    Status    *UpgradeDeviceStatusEnum `json:"status,omitempty"`
    Targets   *UpgradeDevicesTargetIds `json:"targets,omitempty"`
}
