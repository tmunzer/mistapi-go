package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// StatsSite represents a StatsSite struct.
// Site statistics
type StatsSite struct {
    Address              string         `json:"address"`
    AlarmtemplateId      uuid.UUID      `json:"alarmtemplate_id"`
    CountryCode          string         `json:"country_code"`
    CreatedTime          float64        `json:"created_time"`
    Id                   uuid.UUID      `json:"id"`
    Lat                  float64        `json:"lat"`
    Latlng               LatLng         `json:"latlng"`
    Lng                  float64        `json:"lng"`
    ModifiedTime         float64        `json:"modified_time"`
    MspId                string         `json:"msp_id"`
    Name                 string         `json:"name"`
    NetworktemplateId    uuid.UUID      `json:"networktemplate_id"`
    NumAp                int            `json:"num_ap"`
    NumApConnected       int            `json:"num_ap_connected"`
    NumClients           int            `json:"num_clients"`
    NumDevices           int            `json:"num_devices"`
    NumDevicesConnected  int            `json:"num_devices_connected"`
    NumGateway           int            `json:"num_gateway"`
    NumGatewayConnected  int            `json:"num_gateway_connected"`
    NumSwitch            int            `json:"num_switch"`
    NumSwitchConnected   int            `json:"num_switch_connected"`
    OrgId                uuid.UUID      `json:"org_id"`
    RftemplateId         uuid.UUID      `json:"rftemplate_id"`
    SecpolicyId          *interface{}   `json:"secpolicy_id,omitempty"`
    SitegroupIds         []uuid.UUID    `json:"sitegroup_ids"`
    Timezone             string         `json:"timezone"`
    Tzoffset             int            `json:"tzoffset"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for StatsSite.
// It customizes the JSON marshaling process for StatsSite objects.
func (s StatsSite) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the StatsSite object to a map representation for JSON marshaling.
func (s StatsSite) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["address"] = s.Address
    structMap["alarmtemplate_id"] = s.AlarmtemplateId
    structMap["country_code"] = s.CountryCode
    structMap["created_time"] = s.CreatedTime
    structMap["id"] = s.Id
    structMap["lat"] = s.Lat
    structMap["latlng"] = s.Latlng.toMap()
    structMap["lng"] = s.Lng
    structMap["modified_time"] = s.ModifiedTime
    structMap["msp_id"] = s.MspId
    structMap["name"] = s.Name
    structMap["networktemplate_id"] = s.NetworktemplateId
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
    structMap["rftemplate_id"] = s.RftemplateId
    if s.SecpolicyId != nil {
        structMap["secpolicy_id"] = s.SecpolicyId
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
    additionalProperties, err := UnmarshalAdditionalProperties(input, "address", "alarmtemplate_id", "country_code", "created_time", "id", "lat", "latlng", "lng", "modified_time", "msp_id", "name", "networktemplate_id", "num_ap", "num_ap_connected", "num_clients", "num_devices", "num_devices_connected", "num_gateway", "num_gateway_connected", "num_switch", "num_switch_connected", "org_id", "rftemplate_id", "secpolicy_id", "sitegroup_ids", "timezone", "tzoffset")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Address = *temp.Address
    s.AlarmtemplateId = *temp.AlarmtemplateId
    s.CountryCode = *temp.CountryCode
    s.CreatedTime = *temp.CreatedTime
    s.Id = *temp.Id
    s.Lat = *temp.Lat
    s.Latlng = *temp.Latlng
    s.Lng = *temp.Lng
    s.ModifiedTime = *temp.ModifiedTime
    s.MspId = *temp.MspId
    s.Name = *temp.Name
    s.NetworktemplateId = *temp.NetworktemplateId
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
    s.RftemplateId = *temp.RftemplateId
    s.SecpolicyId = temp.SecpolicyId
    s.SitegroupIds = *temp.SitegroupIds
    s.Timezone = *temp.Timezone
    s.Tzoffset = *temp.Tzoffset
    return nil
}

// tempStatsSite is a temporary struct used for validating the fields of StatsSite.
type tempStatsSite  struct {
    Address             *string      `json:"address"`
    AlarmtemplateId     *uuid.UUID   `json:"alarmtemplate_id"`
    CountryCode         *string      `json:"country_code"`
    CreatedTime         *float64     `json:"created_time"`
    Id                  *uuid.UUID   `json:"id"`
    Lat                 *float64     `json:"lat"`
    Latlng              *LatLng      `json:"latlng"`
    Lng                 *float64     `json:"lng"`
    ModifiedTime        *float64     `json:"modified_time"`
    MspId               *string      `json:"msp_id"`
    Name                *string      `json:"name"`
    NetworktemplateId   *uuid.UUID   `json:"networktemplate_id"`
    NumAp               *int         `json:"num_ap"`
    NumApConnected      *int         `json:"num_ap_connected"`
    NumClients          *int         `json:"num_clients"`
    NumDevices          *int         `json:"num_devices"`
    NumDevicesConnected *int         `json:"num_devices_connected"`
    NumGateway          *int         `json:"num_gateway"`
    NumGatewayConnected *int         `json:"num_gateway_connected"`
    NumSwitch           *int         `json:"num_switch"`
    NumSwitchConnected  *int         `json:"num_switch_connected"`
    OrgId               *uuid.UUID   `json:"org_id"`
    RftemplateId        *uuid.UUID   `json:"rftemplate_id"`
    SecpolicyId         *interface{} `json:"secpolicy_id,omitempty"`
    SitegroupIds        *[]uuid.UUID `json:"sitegroup_ids"`
    Timezone            *string      `json:"timezone"`
    Tzoffset            *int         `json:"tzoffset"`
}

func (s *tempStatsSite) validate() error {
    var errs []string
    if s.Address == nil {
        errs = append(errs, "required field `address` is missing for type `stats_site`")
    }
    if s.AlarmtemplateId == nil {
        errs = append(errs, "required field `alarmtemplate_id` is missing for type `stats_site`")
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
    if s.NetworktemplateId == nil {
        errs = append(errs, "required field `networktemplate_id` is missing for type `stats_site`")
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
    if s.RftemplateId == nil {
        errs = append(errs, "required field `rftemplate_id` is missing for type `stats_site`")
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
    return errors.New(strings.Join(errs, "\n"))
}
