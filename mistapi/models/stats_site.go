package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/google/uuid"
    "strings"
)

// StatsSite represents a StatsSite struct.
// Site statistics
type StatsSite struct {
    Address              string                 `json:"address"`
    AlarmtemplateId      *uuid.UUID             `json:"alarmtemplate_id"`
    CountryCode          string                 `json:"country_code"`
    // When the object has been created, in epoch
    CreatedTime          float64                `json:"created_time"`
    // Unique ID of the object instance in the Mist Organnization
    Id                   uuid.UUID              `json:"id"`
    Lat                  float64                `json:"lat"`
    Latlng               LatLng                 `json:"latlng"`
    Lng                  float64                `json:"lng"`
    // When the object has been modified for the last time, in epoch
    ModifiedTime         float64                `json:"modified_time"`
    MspId                uuid.UUID              `json:"msp_id"`
    Name                 string                 `json:"name"`
    NetworktemplateId    *uuid.UUID             `json:"networktemplate_id"`
    NumAp                int                    `json:"num_ap"`
    NumApConnected       int                    `json:"num_ap_connected"`
    NumClients           int                    `json:"num_clients"`
    NumDevices           int                    `json:"num_devices"`
    NumDevicesConnected  int                    `json:"num_devices_connected"`
    NumGateway           int                    `json:"num_gateway"`
    NumGatewayConnected  int                    `json:"num_gateway_connected"`
    NumSwitch            int                    `json:"num_switch"`
    NumSwitchConnected   int                    `json:"num_switch_connected"`
    OrgId                uuid.UUID              `json:"org_id"`
    RftemplateId         *uuid.UUID             `json:"rftemplate_id"`
    SecpolicyId          Optional[uuid.UUID]    `json:"secpolicy_id"`
    SitegroupIds         []uuid.UUID            `json:"sitegroup_ids"`
    Timezone             string                 `json:"timezone"`
    Tzoffset             int                    `json:"tzoffset"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsSite,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsSite) String() string {
    return fmt.Sprintf(
    	"StatsSite[Address=%v, AlarmtemplateId=%v, CountryCode=%v, CreatedTime=%v, Id=%v, Lat=%v, Latlng=%v, Lng=%v, ModifiedTime=%v, MspId=%v, Name=%v, NetworktemplateId=%v, NumAp=%v, NumApConnected=%v, NumClients=%v, NumDevices=%v, NumDevicesConnected=%v, NumGateway=%v, NumGatewayConnected=%v, NumSwitch=%v, NumSwitchConnected=%v, OrgId=%v, RftemplateId=%v, SecpolicyId=%v, SitegroupIds=%v, Timezone=%v, Tzoffset=%v, AdditionalProperties=%v]",
    	s.Address, s.AlarmtemplateId, s.CountryCode, s.CreatedTime, s.Id, s.Lat, s.Latlng, s.Lng, s.ModifiedTime, s.MspId, s.Name, s.NetworktemplateId, s.NumAp, s.NumApConnected, s.NumClients, s.NumDevices, s.NumDevicesConnected, s.NumGateway, s.NumGatewayConnected, s.NumSwitch, s.NumSwitchConnected, s.OrgId, s.RftemplateId, s.SecpolicyId, s.SitegroupIds, s.Timezone, s.Tzoffset, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsSite.
// It customizes the JSON marshaling process for StatsSite objects.
func (s StatsSite) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "address", "alarmtemplate_id", "country_code", "created_time", "id", "lat", "latlng", "lng", "modified_time", "msp_id", "name", "networktemplate_id", "num_ap", "num_ap_connected", "num_clients", "num_devices", "num_devices_connected", "num_gateway", "num_gateway_connected", "num_switch", "num_switch_connected", "org_id", "rftemplate_id", "secpolicy_id", "sitegroup_ids", "timezone", "tzoffset"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsSite object to a map representation for JSON marshaling.
func (s StatsSite) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["address"] = s.Address
    if s.AlarmtemplateId != nil {
        structMap["alarmtemplate_id"] = s.AlarmtemplateId
    } else {
        structMap["alarmtemplate_id"] = nil
    }
    structMap["country_code"] = s.CountryCode
    structMap["created_time"] = s.CreatedTime
    structMap["id"] = s.Id
    structMap["lat"] = s.Lat
    structMap["latlng"] = s.Latlng.toMap()
    structMap["lng"] = s.Lng
    structMap["modified_time"] = s.ModifiedTime
    structMap["msp_id"] = s.MspId
    structMap["name"] = s.Name
    if s.NetworktemplateId != nil {
        structMap["networktemplate_id"] = s.NetworktemplateId
    } else {
        structMap["networktemplate_id"] = nil
    }
    structMap["num_ap"] = s.NumAp
    structMap["num_ap_connected"] = s.NumApConnected
    structMap["num_clients"] = s.NumClients
    structMap["num_devices"] = s.NumDevices
    structMap["num_devices_connected"] = s.NumDevicesConnected
    structMap["num_gateway"] = s.NumGateway
    structMap["num_gateway_connected"] = s.NumGatewayConnected
    structMap["num_switch"] = s.NumSwitch
    structMap["num_switch_connected"] = s.NumSwitchConnected
    structMap["org_id"] = s.OrgId
    if s.RftemplateId != nil {
        structMap["rftemplate_id"] = s.RftemplateId
    } else {
        structMap["rftemplate_id"] = nil
    }
    if s.SecpolicyId.IsValueSet() {
        if s.SecpolicyId.Value() != nil {
            structMap["secpolicy_id"] = s.SecpolicyId.Value()
        } else {
            structMap["secpolicy_id"] = nil
        }
    }
    structMap["sitegroup_ids"] = s.SitegroupIds
    structMap["timezone"] = s.Timezone
    structMap["tzoffset"] = s.Tzoffset
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsSite.
// It customizes the JSON unmarshaling process for StatsSite objects.
func (s *StatsSite) UnmarshalJSON(input []byte) error {
    var temp tempStatsSite
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "address", "alarmtemplate_id", "country_code", "created_time", "id", "lat", "latlng", "lng", "modified_time", "msp_id", "name", "networktemplate_id", "num_ap", "num_ap_connected", "num_clients", "num_devices", "num_devices_connected", "num_gateway", "num_gateway_connected", "num_switch", "num_switch_connected", "org_id", "rftemplate_id", "secpolicy_id", "sitegroup_ids", "timezone", "tzoffset")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Address = *temp.Address
    s.AlarmtemplateId = temp.AlarmtemplateId
    s.CountryCode = *temp.CountryCode
    s.CreatedTime = *temp.CreatedTime
    s.Id = *temp.Id
    s.Lat = *temp.Lat
    s.Latlng = *temp.Latlng
    s.Lng = *temp.Lng
    s.ModifiedTime = *temp.ModifiedTime
    s.MspId = *temp.MspId
    s.Name = *temp.Name
    s.NetworktemplateId = temp.NetworktemplateId
    s.NumAp = *temp.NumAp
    s.NumApConnected = *temp.NumApConnected
    s.NumClients = *temp.NumClients
    s.NumDevices = *temp.NumDevices
    s.NumDevicesConnected = *temp.NumDevicesConnected
    s.NumGateway = *temp.NumGateway
    s.NumGatewayConnected = *temp.NumGatewayConnected
    s.NumSwitch = *temp.NumSwitch
    s.NumSwitchConnected = *temp.NumSwitchConnected
    s.OrgId = *temp.OrgId
    s.RftemplateId = temp.RftemplateId
    s.SecpolicyId = temp.SecpolicyId
    s.SitegroupIds = *temp.SitegroupIds
    s.Timezone = *temp.Timezone
    s.Tzoffset = *temp.Tzoffset
    return nil
}

// tempStatsSite is a temporary struct used for validating the fields of StatsSite.
type tempStatsSite  struct {
    Address             *string             `json:"address"`
    AlarmtemplateId     *uuid.UUID          `json:"alarmtemplate_id"`
    CountryCode         *string             `json:"country_code"`
    CreatedTime         *float64            `json:"created_time"`
    Id                  *uuid.UUID          `json:"id"`
    Lat                 *float64            `json:"lat"`
    Latlng              *LatLng             `json:"latlng"`
    Lng                 *float64            `json:"lng"`
    ModifiedTime        *float64            `json:"modified_time"`
    MspId               *uuid.UUID          `json:"msp_id"`
    Name                *string             `json:"name"`
    NetworktemplateId   *uuid.UUID          `json:"networktemplate_id"`
    NumAp               *int                `json:"num_ap"`
    NumApConnected      *int                `json:"num_ap_connected"`
    NumClients          *int                `json:"num_clients"`
    NumDevices          *int                `json:"num_devices"`
    NumDevicesConnected *int                `json:"num_devices_connected"`
    NumGateway          *int                `json:"num_gateway"`
    NumGatewayConnected *int                `json:"num_gateway_connected"`
    NumSwitch           *int                `json:"num_switch"`
    NumSwitchConnected  *int                `json:"num_switch_connected"`
    OrgId               *uuid.UUID          `json:"org_id"`
    RftemplateId        *uuid.UUID          `json:"rftemplate_id"`
    SecpolicyId         Optional[uuid.UUID] `json:"secpolicy_id"`
    SitegroupIds        *[]uuid.UUID        `json:"sitegroup_ids"`
    Timezone            *string             `json:"timezone"`
    Tzoffset            *int                `json:"tzoffset"`
}

func (s *tempStatsSite) validate() error {
    var errs []string
    if s.Address == nil {
        errs = append(errs, "required field `address` is missing for type `stats_site`")
    }
    if s.CountryCode == nil {
        errs = append(errs, "required field `country_code` is missing for type `stats_site`")
    }
    if s.CreatedTime == nil {
        errs = append(errs, "required field `created_time` is missing for type `stats_site`")
    }
    if s.Id == nil {
        errs = append(errs, "required field `id` is missing for type `stats_site`")
    }
    if s.Lat == nil {
        errs = append(errs, "required field `lat` is missing for type `stats_site`")
    }
    if s.Latlng == nil {
        errs = append(errs, "required field `latlng` is missing for type `stats_site`")
    }
    if s.Lng == nil {
        errs = append(errs, "required field `lng` is missing for type `stats_site`")
    }
    if s.ModifiedTime == nil {
        errs = append(errs, "required field `modified_time` is missing for type `stats_site`")
    }
    if s.MspId == nil {
        errs = append(errs, "required field `msp_id` is missing for type `stats_site`")
    }
    if s.Name == nil {
        errs = append(errs, "required field `name` is missing for type `stats_site`")
    }
    if s.NumAp == nil {
        errs = append(errs, "required field `num_ap` is missing for type `stats_site`")
    }
    if s.NumApConnected == nil {
        errs = append(errs, "required field `num_ap_connected` is missing for type `stats_site`")
    }
    if s.NumClients == nil {
        errs = append(errs, "required field `num_clients` is missing for type `stats_site`")
    }
    if s.NumDevices == nil {
        errs = append(errs, "required field `num_devices` is missing for type `stats_site`")
    }
    if s.NumDevicesConnected == nil {
        errs = append(errs, "required field `num_devices_connected` is missing for type `stats_site`")
    }
    if s.NumGateway == nil {
        errs = append(errs, "required field `num_gateway` is missing for type `stats_site`")
    }
    if s.NumGatewayConnected == nil {
        errs = append(errs, "required field `num_gateway_connected` is missing for type `stats_site`")
    }
    if s.NumSwitch == nil {
        errs = append(errs, "required field `num_switch` is missing for type `stats_site`")
    }
    if s.NumSwitchConnected == nil {
        errs = append(errs, "required field `num_switch_connected` is missing for type `stats_site`")
    }
    if s.OrgId == nil {
        errs = append(errs, "required field `org_id` is missing for type `stats_site`")
    }
    if s.SitegroupIds == nil {
        errs = append(errs, "required field `sitegroup_ids` is missing for type `stats_site`")
    }
    if s.Timezone == nil {
        errs = append(errs, "required field `timezone` is missing for type `stats_site`")
    }
    if s.Tzoffset == nil {
        errs = append(errs, "required field `tzoffset` is missing for type `stats_site`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
