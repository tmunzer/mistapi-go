package models

import (
    "encoding/json"
)

// TunnelConfigsAutoProvision represents a TunnelConfigsAutoProvision struct.
type TunnelConfigsAutoProvision struct {
    Enable               *bool                           `json:"enable,omitempty"`
    Latlng               *LatLng                         `json:"latlng,omitempty"`
    Primary              *TunnelConfigsAutoProvisionNode `json:"primary,omitempty"`
    Secondary            *TunnelConfigsAutoProvisionNode `json:"secondary,omitempty"`
    AdditionalProperties map[string]any                  `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for TunnelConfigsAutoProvision.
// It customizes the JSON marshaling process for TunnelConfigsAutoProvision objects.
func (t TunnelConfigsAutoProvision) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(t.toMap())
}

// toMap converts the TunnelConfigsAutoProvision object to a map representation for JSON marshaling.
func (t TunnelConfigsAutoProvision) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, t.AdditionalProperties)
    if t.Enable != nil {
        structMap["enable"] = t.Enable
    }
    if t.Latlng != nil {
        structMap["latlng"] = t.Latlng.toMap()
    }
    if t.Primary != nil {
        structMap["primary"] = t.Primary.toMap()
    }
    if t.Secondary != nil {
        structMap["secondary"] = t.Secondary.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TunnelConfigsAutoProvision.
// It customizes the JSON unmarshaling process for TunnelConfigsAutoProvision objects.
func (t *TunnelConfigsAutoProvision) UnmarshalJSON(input []byte) error {
    var temp tempTunnelConfigsAutoProvision
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "enable", "latlng", "primary", "secondary")
    if err != nil {
    	return err
    }
    
    t.AdditionalProperties = additionalProperties
    t.Enable = temp.Enable
    t.Latlng = temp.Latlng
    t.Primary = temp.Primary
    t.Secondary = temp.Secondary
    return nil
}

// tempTunnelConfigsAutoProvision is a temporary struct used for validating the fields of TunnelConfigsAutoProvision.
type tempTunnelConfigsAutoProvision  struct {
    Enable    *bool                           `json:"enable,omitempty"`
    Latlng    *LatLng                         `json:"latlng,omitempty"`
    Primary   *TunnelConfigsAutoProvisionNode `json:"primary,omitempty"`
    Secondary *TunnelConfigsAutoProvisionNode `json:"secondary,omitempty"`
}
