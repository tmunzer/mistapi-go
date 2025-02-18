package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/google/uuid"
    "strings"
)

// StatsOrg represents a StatsOrg struct.
// Org statistics
type StatsOrg struct {
    AlarmtemplateId        uuid.UUID              `json:"alarmtemplate_id"`
    AllowMist              bool                   `json:"allow_mist"`
    // When the object has been created, in epoch
    CreatedTime            float64                `json:"created_time"`
    // Unique ID of the object instance in the Mist Organnization
    Id                     uuid.UUID              `json:"id"`
    // When the object has been modified for the last time, in epoch
    ModifiedTime           float64                `json:"modified_time"`
    MspId                  uuid.UUID              `json:"msp_id"`
    Name                   string                 `json:"name"`
    NumDevices             int                    `json:"num_devices"`
    NumDevicesConnected    int                    `json:"num_devices_connected"`
    NumDevicesDisconnected int                    `json:"num_devices_disconnected"`
    NumInventory           int                    `json:"num_inventory"`
    NumSites               int                    `json:"num_sites"`
    OrggroupIds            []uuid.UUID            `json:"orggroup_ids"`
    SessionExpiry          int64                  `json:"session_expiry"`
    Sle                    []StatsOrgSle          `json:"sle"`
    AdditionalProperties   map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsOrg,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsOrg) String() string {
    return fmt.Sprintf(
    	"StatsOrg[AlarmtemplateId=%v, AllowMist=%v, CreatedTime=%v, Id=%v, ModifiedTime=%v, MspId=%v, Name=%v, NumDevices=%v, NumDevicesConnected=%v, NumDevicesDisconnected=%v, NumInventory=%v, NumSites=%v, OrggroupIds=%v, SessionExpiry=%v, Sle=%v, AdditionalProperties=%v]",
    	s.AlarmtemplateId, s.AllowMist, s.CreatedTime, s.Id, s.ModifiedTime, s.MspId, s.Name, s.NumDevices, s.NumDevicesConnected, s.NumDevicesDisconnected, s.NumInventory, s.NumSites, s.OrggroupIds, s.SessionExpiry, s.Sle, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsOrg.
// It customizes the JSON marshaling process for StatsOrg objects.
func (s StatsOrg) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "alarmtemplate_id", "allow_mist", "created_time", "id", "modified_time", "msp_id", "name", "num_devices", "num_devices_connected", "num_devices_disconnected", "num_inventory", "num_sites", "orggroup_ids", "session_expiry", "sle"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsOrg object to a map representation for JSON marshaling.
func (s StatsOrg) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["alarmtemplate_id"] = s.AlarmtemplateId
    structMap["allow_mist"] = s.AllowMist
    structMap["created_time"] = s.CreatedTime
    structMap["id"] = s.Id
    structMap["modified_time"] = s.ModifiedTime
    structMap["msp_id"] = s.MspId
    structMap["name"] = s.Name
    structMap["num_devices"] = s.NumDevices
    structMap["num_devices_connected"] = s.NumDevicesConnected
    structMap["num_devices_disconnected"] = s.NumDevicesDisconnected
    structMap["num_inventory"] = s.NumInventory
    structMap["num_sites"] = s.NumSites
    structMap["orggroup_ids"] = s.OrggroupIds
    structMap["session_expiry"] = s.SessionExpiry
    structMap["sle"] = s.Sle
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsOrg.
// It customizes the JSON unmarshaling process for StatsOrg objects.
func (s *StatsOrg) UnmarshalJSON(input []byte) error {
    var temp tempStatsOrg
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "alarmtemplate_id", "allow_mist", "created_time", "id", "modified_time", "msp_id", "name", "num_devices", "num_devices_connected", "num_devices_disconnected", "num_inventory", "num_sites", "orggroup_ids", "session_expiry", "sle")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.AlarmtemplateId = *temp.AlarmtemplateId
    s.AllowMist = *temp.AllowMist
    s.CreatedTime = *temp.CreatedTime
    s.Id = *temp.Id
    s.ModifiedTime = *temp.ModifiedTime
    s.MspId = *temp.MspId
    s.Name = *temp.Name
    s.NumDevices = *temp.NumDevices
    s.NumDevicesConnected = *temp.NumDevicesConnected
    s.NumDevicesDisconnected = *temp.NumDevicesDisconnected
    s.NumInventory = *temp.NumInventory
    s.NumSites = *temp.NumSites
    s.OrggroupIds = *temp.OrggroupIds
    s.SessionExpiry = *temp.SessionExpiry
    s.Sle = *temp.Sle
    return nil
}

// tempStatsOrg is a temporary struct used for validating the fields of StatsOrg.
type tempStatsOrg  struct {
    AlarmtemplateId        *uuid.UUID     `json:"alarmtemplate_id"`
    AllowMist              *bool          `json:"allow_mist"`
    CreatedTime            *float64       `json:"created_time"`
    Id                     *uuid.UUID     `json:"id"`
    ModifiedTime           *float64       `json:"modified_time"`
    MspId                  *uuid.UUID     `json:"msp_id"`
    Name                   *string        `json:"name"`
    NumDevices             *int           `json:"num_devices"`
    NumDevicesConnected    *int           `json:"num_devices_connected"`
    NumDevicesDisconnected *int           `json:"num_devices_disconnected"`
    NumInventory           *int           `json:"num_inventory"`
    NumSites               *int           `json:"num_sites"`
    OrggroupIds            *[]uuid.UUID   `json:"orggroup_ids"`
    SessionExpiry          *int64         `json:"session_expiry"`
    Sle                    *[]StatsOrgSle `json:"sle"`
}

func (s *tempStatsOrg) validate() error {
    var errs []string
    if s.AlarmtemplateId == nil {
        errs = append(errs, "required field `alarmtemplate_id` is missing for type `stats_org`")
    }
    if s.AllowMist == nil {
        errs = append(errs, "required field `allow_mist` is missing for type `stats_org`")
    }
    if s.CreatedTime == nil {
        errs = append(errs, "required field `created_time` is missing for type `stats_org`")
    }
    if s.Id == nil {
        errs = append(errs, "required field `id` is missing for type `stats_org`")
    }
    if s.ModifiedTime == nil {
        errs = append(errs, "required field `modified_time` is missing for type `stats_org`")
    }
    if s.MspId == nil {
        errs = append(errs, "required field `msp_id` is missing for type `stats_org`")
    }
    if s.Name == nil {
        errs = append(errs, "required field `name` is missing for type `stats_org`")
    }
    if s.NumDevices == nil {
        errs = append(errs, "required field `num_devices` is missing for type `stats_org`")
    }
    if s.NumDevicesConnected == nil {
        errs = append(errs, "required field `num_devices_connected` is missing for type `stats_org`")
    }
    if s.NumDevicesDisconnected == nil {
        errs = append(errs, "required field `num_devices_disconnected` is missing for type `stats_org`")
    }
    if s.NumInventory == nil {
        errs = append(errs, "required field `num_inventory` is missing for type `stats_org`")
    }
    if s.NumSites == nil {
        errs = append(errs, "required field `num_sites` is missing for type `stats_org`")
    }
    if s.OrggroupIds == nil {
        errs = append(errs, "required field `orggroup_ids` is missing for type `stats_org`")
    }
    if s.SessionExpiry == nil {
        errs = append(errs, "required field `session_expiry` is missing for type `stats_org`")
    }
    if s.Sle == nil {
        errs = append(errs, "required field `sle` is missing for type `stats_org`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
