package models

import (
    "encoding/json"
)

// DeviceOtherStatsVendorSpecificPort represents a DeviceOtherStatsVendorSpecificPort struct.
type DeviceOtherStatsVendorSpecificPort struct {
    BytesIn              *int           `json:"bytes_in,omitempty"`
    BytesOut             *int           `json:"bytes_out,omitempty"`
    HealthCategory       *string        `json:"health_category,omitempty"`
    HealthScore          *int           `json:"health_score,omitempty"`
    Id                   *string        `json:"id,omitempty"`
    Mode                 *string        `json:"mode,omitempty"`
    Model                *string        `json:"model,omitempty"`
    State                *string        `json:"state,omitempty"`
    Type                 *string        `json:"type,omitempty"`
    Uptime               *float64       `json:"uptime,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for DeviceOtherStatsVendorSpecificPort.
// It customizes the JSON marshaling process for DeviceOtherStatsVendorSpecificPort objects.
func (d DeviceOtherStatsVendorSpecificPort) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(d.toMap())
}

// toMap converts the DeviceOtherStatsVendorSpecificPort object to a map representation for JSON marshaling.
func (d DeviceOtherStatsVendorSpecificPort) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, d.AdditionalProperties)
    if d.BytesIn != nil {
        structMap["bytes_in"] = d.BytesIn
    }
    if d.BytesOut != nil {
        structMap["bytes_out"] = d.BytesOut
    }
    if d.HealthCategory != nil {
        structMap["health_category"] = d.HealthCategory
    }
    if d.HealthScore != nil {
        structMap["health_score"] = d.HealthScore
    }
    if d.Id != nil {
        structMap["id"] = d.Id
    }
    if d.Mode != nil {
        structMap["mode"] = d.Mode
    }
    if d.Model != nil {
        structMap["model"] = d.Model
    }
    if d.State != nil {
        structMap["state"] = d.State
    }
    if d.Type != nil {
        structMap["type"] = d.Type
    }
    if d.Uptime != nil {
        structMap["uptime"] = d.Uptime
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DeviceOtherStatsVendorSpecificPort.
// It customizes the JSON unmarshaling process for DeviceOtherStatsVendorSpecificPort objects.
func (d *DeviceOtherStatsVendorSpecificPort) UnmarshalJSON(input []byte) error {
    var temp deviceOtherStatsVendorSpecificPort
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "bytes_in", "bytes_out", "health_category", "health_score", "id", "mode", "model", "state", "type", "uptime")
    if err != nil {
    	return err
    }
    
    d.AdditionalProperties = additionalProperties
    d.BytesIn = temp.BytesIn
    d.BytesOut = temp.BytesOut
    d.HealthCategory = temp.HealthCategory
    d.HealthScore = temp.HealthScore
    d.Id = temp.Id
    d.Mode = temp.Mode
    d.Model = temp.Model
    d.State = temp.State
    d.Type = temp.Type
    d.Uptime = temp.Uptime
    return nil
}

// deviceOtherStatsVendorSpecificPort is a temporary struct used for validating the fields of DeviceOtherStatsVendorSpecificPort.
type deviceOtherStatsVendorSpecificPort  struct {
    BytesIn        *int     `json:"bytes_in,omitempty"`
    BytesOut       *int     `json:"bytes_out,omitempty"`
    HealthCategory *string  `json:"health_category,omitempty"`
    HealthScore    *int     `json:"health_score,omitempty"`
    Id             *string  `json:"id,omitempty"`
    Mode           *string  `json:"mode,omitempty"`
    Model          *string  `json:"model,omitempty"`
    State          *string  `json:"state,omitempty"`
    Type           *string  `json:"type,omitempty"`
    Uptime         *float64 `json:"uptime,omitempty"`
}
