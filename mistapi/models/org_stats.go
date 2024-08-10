package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// OrgStats represents a OrgStats struct.
// Org statistics
type OrgStats struct {
    AlarmtemplateId        uuid.UUID      `json:"alarmtemplate_id"`
    AllowMist              bool           `json:"allow_mist"`
    CreatedTime            float64        `json:"created_time"`
    Id                     uuid.UUID      `json:"id"`
    ModifiedTime           float64        `json:"modified_time"`
    MspId                  uuid.UUID      `json:"msp_id"`
    Name                   string         `json:"name"`
    NumDevices             int            `json:"num_devices"`
    NumDevicesConnected    int            `json:"num_devices_connected"`
    NumDevicesDisconnected int            `json:"num_devices_disconnected"`
    NumInventory           int            `json:"num_inventory"`
    NumSites               int            `json:"num_sites"`
    OrggroupIds            []uuid.UUID    `json:"orggroup_ids"`
    SessionExpiry          int64          `json:"session_expiry"`
    Sle                    []OrgSleStat   `json:"sle"`
    AdditionalProperties   map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for OrgStats.
// It customizes the JSON marshaling process for OrgStats objects.
func (o OrgStats) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(o.toMap())
}

// toMap converts the OrgStats object to a map representation for JSON marshaling.
func (o OrgStats) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, o.AdditionalProperties)
    structMap["alarmtemplate_id"] = o.AlarmtemplateId
    structMap["allow_mist"] = o.AllowMist
    structMap["created_time"] = o.CreatedTime
    structMap["id"] = o.Id
    structMap["modified_time"] = o.ModifiedTime
    structMap["msp_id"] = o.MspId
    structMap["name"] = o.Name
    structMap["num_devices"] = o.NumDevices
    structMap["num_devices_connected"] = o.NumDevicesConnected
    structMap["num_devices_disconnected"] = o.NumDevicesDisconnected
    structMap["num_inventory"] = o.NumInventory
    structMap["num_sites"] = o.NumSites
    structMap["orggroup_ids"] = o.OrggroupIds
    structMap["session_expiry"] = o.SessionExpiry
    structMap["sle"] = o.Sle
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgStats.
// It customizes the JSON unmarshaling process for OrgStats objects.
func (o *OrgStats) UnmarshalJSON(input []byte) error {
    var temp tempOrgStats
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "alarmtemplate_id", "allow_mist", "created_time", "id", "modified_time", "msp_id", "name", "num_devices", "num_devices_connected", "num_devices_disconnected", "num_inventory", "num_sites", "orggroup_ids", "session_expiry", "sle")
    if err != nil {
    	return err
    }
    
    o.AdditionalProperties = additionalProperties
    o.AlarmtemplateId = *temp.AlarmtemplateId
    o.AllowMist = *temp.AllowMist
    o.CreatedTime = *temp.CreatedTime
    o.Id = *temp.Id
    o.ModifiedTime = *temp.ModifiedTime
    o.MspId = *temp.MspId
    o.Name = *temp.Name
    o.NumDevices = *temp.NumDevices
    o.NumDevicesConnected = *temp.NumDevicesConnected
    o.NumDevicesDisconnected = *temp.NumDevicesDisconnected
    o.NumInventory = *temp.NumInventory
    o.NumSites = *temp.NumSites
    o.OrggroupIds = *temp.OrggroupIds
    o.SessionExpiry = *temp.SessionExpiry
    o.Sle = *temp.Sle
    return nil
}

// tempOrgStats is a temporary struct used for validating the fields of OrgStats.
type tempOrgStats  struct {
    AlarmtemplateId        *uuid.UUID    `json:"alarmtemplate_id"`
    AllowMist              *bool         `json:"allow_mist"`
    CreatedTime            *float64      `json:"created_time"`
    Id                     *uuid.UUID    `json:"id"`
    ModifiedTime           *float64      `json:"modified_time"`
    MspId                  *uuid.UUID    `json:"msp_id"`
    Name                   *string       `json:"name"`
    NumDevices             *int          `json:"num_devices"`
    NumDevicesConnected    *int          `json:"num_devices_connected"`
    NumDevicesDisconnected *int          `json:"num_devices_disconnected"`
    NumInventory           *int          `json:"num_inventory"`
    NumSites               *int          `json:"num_sites"`
    OrggroupIds            *[]uuid.UUID  `json:"orggroup_ids"`
    SessionExpiry          *int64        `json:"session_expiry"`
    Sle                    *[]OrgSleStat `json:"sle"`
}

func (o *tempOrgStats) validate() error {
    var errs []string
    if o.AlarmtemplateId == nil {
        errs = append(errs, "required field `alarmtemplate_id` is missing for type `org_stats`")
    }
    if o.AllowMist == nil {
        errs = append(errs, "required field `allow_mist` is missing for type `org_stats`")
    }
    if o.CreatedTime == nil {
        errs = append(errs, "required field `created_time` is missing for type `org_stats`")
    }
    if o.Id == nil {
        errs = append(errs, "required field `id` is missing for type `org_stats`")
    }
    if o.ModifiedTime == nil {
        errs = append(errs, "required field `modified_time` is missing for type `org_stats`")
    }
    if o.MspId == nil {
        errs = append(errs, "required field `msp_id` is missing for type `org_stats`")
    }
    if o.Name == nil {
        errs = append(errs, "required field `name` is missing for type `org_stats`")
    }
    if o.NumDevices == nil {
        errs = append(errs, "required field `num_devices` is missing for type `org_stats`")
    }
    if o.NumDevicesConnected == nil {
        errs = append(errs, "required field `num_devices_connected` is missing for type `org_stats`")
    }
    if o.NumDevicesDisconnected == nil {
        errs = append(errs, "required field `num_devices_disconnected` is missing for type `org_stats`")
    }
    if o.NumInventory == nil {
        errs = append(errs, "required field `num_inventory` is missing for type `org_stats`")
    }
    if o.NumSites == nil {
        errs = append(errs, "required field `num_sites` is missing for type `org_stats`")
    }
    if o.OrggroupIds == nil {
        errs = append(errs, "required field `orggroup_ids` is missing for type `org_stats`")
    }
    if o.SessionExpiry == nil {
        errs = append(errs, "required field `session_expiry` is missing for type `org_stats`")
    }
    if o.Sle == nil {
        errs = append(errs, "required field `sle` is missing for type `org_stats`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
