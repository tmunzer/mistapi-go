package models

import (
    "encoding/json"
)

// SwitchPortMirroring represents a SwitchPortMirroring struct.
// Property key is the port mirroring instance name
// port_mirroring can be added under device/site settings. It takes interface and ports as input for ingress, interface as input for egress and can take interface and port as output.
type SwitchPortMirroring struct {
    PortMirror           *GatewayPortMirroringPortMirror `json:"port_mirror,omitempty"`
    AdditionalProperties map[string]any                  `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SwitchPortMirroring.
// It customizes the JSON marshaling process for SwitchPortMirroring objects.
func (s SwitchPortMirroring) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SwitchPortMirroring object to a map representation for JSON marshaling.
func (s SwitchPortMirroring) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.PortMirror != nil {
        structMap["port_mirror"] = s.PortMirror.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchPortMirroring.
// It customizes the JSON unmarshaling process for SwitchPortMirroring objects.
func (s *SwitchPortMirroring) UnmarshalJSON(input []byte) error {
    var temp switchPortMirroring
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "port_mirror")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.PortMirror = temp.PortMirror
    return nil
}

// switchPortMirroring is a temporary struct used for validating the fields of SwitchPortMirroring.
type switchPortMirroring  struct {
    PortMirror *GatewayPortMirroringPortMirror `json:"port_mirror,omitempty"`
}
