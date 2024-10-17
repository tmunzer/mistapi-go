package models

import (
	"encoding/json"
	"github.com/google/uuid"
)

// DiscoveredSwitchMetric represents a DiscoveredSwitchMetric struct.
type DiscoveredSwitchMetric struct {
	Adopted              *bool                      `json:"adopted,omitempty"`
	Aps                  []DiscoveredSwitchMetricAp `json:"aps,omitempty"`
	ChassisId            []string                   `json:"chassis_id,omitempty"`
	Hostname             *string                    `json:"hostname,omitempty"`
	MgmtAddr             *string                    `json:"mgmt_addr,omitempty"`
	Model                *string                    `json:"model,omitempty"`
	OrgId                *uuid.UUID                 `json:"org_id,omitempty"`
	Scope                *string                    `json:"scope,omitempty"`
	Score                *int                       `json:"score,omitempty"`
	SiteId               *uuid.UUID                 `json:"site_id,omitempty"`
	SystemDesc           *string                    `json:"system_desc,omitempty"`
	SystemName           *string                    `json:"system_name,omitempty"`
	Timestamp            *int                       `json:"timestamp,omitempty"`
	Type                 *string                    `json:"type,omitempty"`
	Vendor               *string                    `json:"vendor,omitempty"`
	Version              *string                    `json:"version,omitempty"`
	AdditionalProperties map[string]any             `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for DiscoveredSwitchMetric.
// It customizes the JSON marshaling process for DiscoveredSwitchMetric objects.
func (d DiscoveredSwitchMetric) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(d.toMap())
}

// toMap converts the DiscoveredSwitchMetric object to a map representation for JSON marshaling.
func (d DiscoveredSwitchMetric) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, d.AdditionalProperties)
	if d.Adopted != nil {
		structMap["adopted"] = d.Adopted
	}
	if d.Aps != nil {
		structMap["aps"] = d.Aps
	}
	if d.ChassisId != nil {
		structMap["chassis_id"] = d.ChassisId
	}
	if d.Hostname != nil {
		structMap["hostname"] = d.Hostname
	}
	if d.MgmtAddr != nil {
		structMap["mgmt_addr"] = d.MgmtAddr
	}
	if d.Model != nil {
		structMap["model"] = d.Model
	}
	if d.OrgId != nil {
		structMap["org_id"] = d.OrgId
	}
	if d.Scope != nil {
		structMap["scope"] = d.Scope
	}
	if d.Score != nil {
		structMap["score"] = d.Score
	}
	if d.SiteId != nil {
		structMap["site_id"] = d.SiteId
	}
	if d.SystemDesc != nil {
		structMap["system_desc"] = d.SystemDesc
	}
	if d.SystemName != nil {
		structMap["system_name"] = d.SystemName
	}
	if d.Timestamp != nil {
		structMap["timestamp"] = d.Timestamp
	}
	if d.Type != nil {
		structMap["type"] = d.Type
	}
	if d.Vendor != nil {
		structMap["vendor"] = d.Vendor
	}
	if d.Version != nil {
		structMap["version"] = d.Version
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DiscoveredSwitchMetric.
// It customizes the JSON unmarshaling process for DiscoveredSwitchMetric objects.
func (d *DiscoveredSwitchMetric) UnmarshalJSON(input []byte) error {
	var temp tempDiscoveredSwitchMetric
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "adopted", "aps", "chassis_id", "hostname", "mgmt_addr", "model", "org_id", "scope", "score", "site_id", "system_desc", "system_name", "timestamp", "type", "vendor", "version")
	if err != nil {
		return err
	}

	d.AdditionalProperties = additionalProperties
	d.Adopted = temp.Adopted
	d.Aps = temp.Aps
	d.ChassisId = temp.ChassisId
	d.Hostname = temp.Hostname
	d.MgmtAddr = temp.MgmtAddr
	d.Model = temp.Model
	d.OrgId = temp.OrgId
	d.Scope = temp.Scope
	d.Score = temp.Score
	d.SiteId = temp.SiteId
	d.SystemDesc = temp.SystemDesc
	d.SystemName = temp.SystemName
	d.Timestamp = temp.Timestamp
	d.Type = temp.Type
	d.Vendor = temp.Vendor
	d.Version = temp.Version
	return nil
}

// tempDiscoveredSwitchMetric is a temporary struct used for validating the fields of DiscoveredSwitchMetric.
type tempDiscoveredSwitchMetric struct {
	Adopted    *bool                      `json:"adopted,omitempty"`
	Aps        []DiscoveredSwitchMetricAp `json:"aps,omitempty"`
	ChassisId  []string                   `json:"chassis_id,omitempty"`
	Hostname   *string                    `json:"hostname,omitempty"`
	MgmtAddr   *string                    `json:"mgmt_addr,omitempty"`
	Model      *string                    `json:"model,omitempty"`
	OrgId      *uuid.UUID                 `json:"org_id,omitempty"`
	Scope      *string                    `json:"scope,omitempty"`
	Score      *int                       `json:"score,omitempty"`
	SiteId     *uuid.UUID                 `json:"site_id,omitempty"`
	SystemDesc *string                    `json:"system_desc,omitempty"`
	SystemName *string                    `json:"system_name,omitempty"`
	Timestamp  *int                       `json:"timestamp,omitempty"`
	Type       *string                    `json:"type,omitempty"`
	Vendor     *string                    `json:"vendor,omitempty"`
	Version    *string                    `json:"version,omitempty"`
}
