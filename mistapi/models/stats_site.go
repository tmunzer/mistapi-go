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

// StatsSite represents a StatsSite struct.
// Site statistics
type StatsSite struct {
	Address         *string             `json:"address,omitempty"`
	AlarmtemplateId Optional[uuid.UUID] `json:"alarmtemplate_id"`
	AnalyticEnabled *bool               `json:"analyticEnabled,omitempty"`
	AptemplateId    Optional[uuid.UUID] `json:"aptemplate_id"`
	CountryCode     string              `json:"country_code"`
	// When the object has been created, in epoch
	CreatedTime       float64             `json:"created_time"`
	EngagementEnabled *bool               `json:"engagementEnabled,omitempty"`
	GatewaytemplateId Optional[uuid.UUID] `json:"gatewaytemplate_id"`
	// Unique ID of the object instance in the Mist Organization
	Id     uuid.UUID `json:"id"`
	Lat    *float64  `json:"lat,omitempty"`
	Latlng LatLng    `json:"latlng"`
	Lng    *float64  `json:"lng,omitempty"`
	// When the object has been modified for the last time, in epoch
	ModifiedTime         float64                `json:"modified_time"`
	MspId                *uuid.UUID             `json:"msp_id,omitempty"`
	Name                 string                 `json:"name"`
	NetworktemplateId    Optional[uuid.UUID]    `json:"networktemplate_id"`
	Notes                *string                `json:"notes,omitempty"`
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
	RftemplateId         Optional[uuid.UUID]    `json:"rftemplate_id"`
	SecpolicyId          Optional[uuid.UUID]    `json:"secpolicy_id"`
	SitegroupIds         []uuid.UUID            `json:"sitegroup_ids,omitempty"`
	SitetemplateId       Optional[uuid.UUID]    `json:"sitetemplate_id"`
	Timezone             string                 `json:"timezone"`
	Tzoffset             int                    `json:"tzoffset"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsSite,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsSite) String() string {
	return fmt.Sprintf(
		"StatsSite[Address=%v, AlarmtemplateId=%v, AnalyticEnabled=%v, AptemplateId=%v, CountryCode=%v, CreatedTime=%v, EngagementEnabled=%v, GatewaytemplateId=%v, Id=%v, Lat=%v, Latlng=%v, Lng=%v, ModifiedTime=%v, MspId=%v, Name=%v, NetworktemplateId=%v, Notes=%v, NumAp=%v, NumApConnected=%v, NumClients=%v, NumDevices=%v, NumDevicesConnected=%v, NumGateway=%v, NumGatewayConnected=%v, NumSwitch=%v, NumSwitchConnected=%v, OrgId=%v, RftemplateId=%v, SecpolicyId=%v, SitegroupIds=%v, SitetemplateId=%v, Timezone=%v, Tzoffset=%v, AdditionalProperties=%v]",
		s.Address, s.AlarmtemplateId, s.AnalyticEnabled, s.AptemplateId, s.CountryCode, s.CreatedTime, s.EngagementEnabled, s.GatewaytemplateId, s.Id, s.Lat, s.Latlng, s.Lng, s.ModifiedTime, s.MspId, s.Name, s.NetworktemplateId, s.Notes, s.NumAp, s.NumApConnected, s.NumClients, s.NumDevices, s.NumDevicesConnected, s.NumGateway, s.NumGatewayConnected, s.NumSwitch, s.NumSwitchConnected, s.OrgId, s.RftemplateId, s.SecpolicyId, s.SitegroupIds, s.SitetemplateId, s.Timezone, s.Tzoffset, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsSite.
// It customizes the JSON marshaling process for StatsSite objects.
func (s StatsSite) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"address", "alarmtemplate_id", "analyticEnabled", "aptemplate_id", "country_code", "created_time", "engagementEnabled", "gatewaytemplate_id", "id", "lat", "latlng", "lng", "modified_time", "msp_id", "name", "networktemplate_id", "notes", "num_ap", "num_ap_connected", "num_clients", "num_devices", "num_devices_connected", "num_gateway", "num_gateway_connected", "num_switch", "num_switch_connected", "org_id", "rftemplate_id", "secpolicy_id", "sitegroup_ids", "sitetemplate_id", "timezone", "tzoffset"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the StatsSite object to a map representation for JSON marshaling.
func (s StatsSite) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Address != nil {
		structMap["address"] = s.Address
	}
	if s.AlarmtemplateId.IsValueSet() {
		if s.AlarmtemplateId.Value() != nil {
			structMap["alarmtemplate_id"] = s.AlarmtemplateId.Value()
		} else {
			structMap["alarmtemplate_id"] = nil
		}
	}
	if s.AnalyticEnabled != nil {
		structMap["analyticEnabled"] = s.AnalyticEnabled
	}
	if s.AptemplateId.IsValueSet() {
		if s.AptemplateId.Value() != nil {
			structMap["aptemplate_id"] = s.AptemplateId.Value()
		} else {
			structMap["aptemplate_id"] = nil
		}
	}
	structMap["country_code"] = s.CountryCode
	structMap["created_time"] = s.CreatedTime
	if s.EngagementEnabled != nil {
		structMap["engagementEnabled"] = s.EngagementEnabled
	}
	if s.GatewaytemplateId.IsValueSet() {
		if s.GatewaytemplateId.Value() != nil {
			structMap["gatewaytemplate_id"] = s.GatewaytemplateId.Value()
		} else {
			structMap["gatewaytemplate_id"] = nil
		}
	}
	structMap["id"] = s.Id
	if s.Lat != nil {
		structMap["lat"] = s.Lat
	}
	structMap["latlng"] = s.Latlng.toMap()
	if s.Lng != nil {
		structMap["lng"] = s.Lng
	}
	structMap["modified_time"] = s.ModifiedTime
	if s.MspId != nil {
		structMap["msp_id"] = s.MspId
	}
	structMap["name"] = s.Name
	if s.NetworktemplateId.IsValueSet() {
		if s.NetworktemplateId.Value() != nil {
			structMap["networktemplate_id"] = s.NetworktemplateId.Value()
		} else {
			structMap["networktemplate_id"] = nil
		}
	}
	if s.Notes != nil {
		structMap["notes"] = s.Notes
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
	if s.RftemplateId.IsValueSet() {
		if s.RftemplateId.Value() != nil {
			structMap["rftemplate_id"] = s.RftemplateId.Value()
		} else {
			structMap["rftemplate_id"] = nil
		}
	}
	if s.SecpolicyId.IsValueSet() {
		if s.SecpolicyId.Value() != nil {
			structMap["secpolicy_id"] = s.SecpolicyId.Value()
		} else {
			structMap["secpolicy_id"] = nil
		}
	}
	if s.SitegroupIds != nil {
		structMap["sitegroup_ids"] = s.SitegroupIds
	}
	if s.SitetemplateId.IsValueSet() {
		if s.SitetemplateId.Value() != nil {
			structMap["sitetemplate_id"] = s.SitetemplateId.Value()
		} else {
			structMap["sitetemplate_id"] = nil
		}
	}
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
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "address", "alarmtemplate_id", "analyticEnabled", "aptemplate_id", "country_code", "created_time", "engagementEnabled", "gatewaytemplate_id", "id", "lat", "latlng", "lng", "modified_time", "msp_id", "name", "networktemplate_id", "notes", "num_ap", "num_ap_connected", "num_clients", "num_devices", "num_devices_connected", "num_gateway", "num_gateway_connected", "num_switch", "num_switch_connected", "org_id", "rftemplate_id", "secpolicy_id", "sitegroup_ids", "sitetemplate_id", "timezone", "tzoffset")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.Address = temp.Address
	s.AlarmtemplateId = temp.AlarmtemplateId
	s.AnalyticEnabled = temp.AnalyticEnabled
	s.AptemplateId = temp.AptemplateId
	s.CountryCode = *temp.CountryCode
	s.CreatedTime = *temp.CreatedTime
	s.EngagementEnabled = temp.EngagementEnabled
	s.GatewaytemplateId = temp.GatewaytemplateId
	s.Id = *temp.Id
	s.Lat = temp.Lat
	s.Latlng = *temp.Latlng
	s.Lng = temp.Lng
	s.ModifiedTime = *temp.ModifiedTime
	s.MspId = temp.MspId
	s.Name = *temp.Name
	s.NetworktemplateId = temp.NetworktemplateId
	s.Notes = temp.Notes
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
	s.SitegroupIds = temp.SitegroupIds
	s.SitetemplateId = temp.SitetemplateId
	s.Timezone = *temp.Timezone
	s.Tzoffset = *temp.Tzoffset
	return nil
}

// tempStatsSite is a temporary struct used for validating the fields of StatsSite.
type tempStatsSite struct {
	Address             *string             `json:"address,omitempty"`
	AlarmtemplateId     Optional[uuid.UUID] `json:"alarmtemplate_id"`
	AnalyticEnabled     *bool               `json:"analyticEnabled,omitempty"`
	AptemplateId        Optional[uuid.UUID] `json:"aptemplate_id"`
	CountryCode         *string             `json:"country_code"`
	CreatedTime         *float64            `json:"created_time"`
	EngagementEnabled   *bool               `json:"engagementEnabled,omitempty"`
	GatewaytemplateId   Optional[uuid.UUID] `json:"gatewaytemplate_id"`
	Id                  *uuid.UUID          `json:"id"`
	Lat                 *float64            `json:"lat,omitempty"`
	Latlng              *LatLng             `json:"latlng"`
	Lng                 *float64            `json:"lng,omitempty"`
	ModifiedTime        *float64            `json:"modified_time"`
	MspId               *uuid.UUID          `json:"msp_id,omitempty"`
	Name                *string             `json:"name"`
	NetworktemplateId   Optional[uuid.UUID] `json:"networktemplate_id"`
	Notes               *string             `json:"notes,omitempty"`
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
	RftemplateId        Optional[uuid.UUID] `json:"rftemplate_id"`
	SecpolicyId         Optional[uuid.UUID] `json:"secpolicy_id"`
	SitegroupIds        []uuid.UUID         `json:"sitegroup_ids,omitempty"`
	SitetemplateId      Optional[uuid.UUID] `json:"sitetemplate_id"`
	Timezone            *string             `json:"timezone"`
	Tzoffset            *int                `json:"tzoffset"`
}

func (s *tempStatsSite) validate() error {
	var errs []string
	if s.CountryCode == nil {
		errs = append(errs, "required field `country_code` is missing for type `stats_site`")
	}
	if s.CreatedTime == nil {
		errs = append(errs, "required field `created_time` is missing for type `stats_site`")
	}
	if s.Id == nil {
		errs = append(errs, "required field `id` is missing for type `stats_site`")
	}
	if s.Latlng == nil {
		errs = append(errs, "required field `latlng` is missing for type `stats_site`")
	}
	if s.ModifiedTime == nil {
		errs = append(errs, "required field `modified_time` is missing for type `stats_site`")
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
