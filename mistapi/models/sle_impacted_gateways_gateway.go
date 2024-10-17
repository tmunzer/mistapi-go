package models

import (
	"encoding/json"
)

// SleImpactedGatewaysGateway represents a SleImpactedGatewaysGateway struct.
type SleImpactedGatewaysGateway struct {
	Degraded             *float64       `json:"degraded,omitempty"`
	Duration             *int           `json:"duration,omitempty"`
	GatewayMac           *string        `json:"gateway_mac,omitempty"`
	GatewayModel         *string        `json:"gateway_model,omitempty"`
	GatewayVersion       *string        `json:"gateway_version,omitempty"`
	Name                 *string        `json:"name,omitempty"`
	Total                *int           `json:"total,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SleImpactedGatewaysGateway.
// It customizes the JSON marshaling process for SleImpactedGatewaysGateway objects.
func (s SleImpactedGatewaysGateway) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the SleImpactedGatewaysGateway object to a map representation for JSON marshaling.
func (s SleImpactedGatewaysGateway) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Degraded != nil {
		structMap["degraded"] = s.Degraded
	}
	if s.Duration != nil {
		structMap["duration"] = s.Duration
	}
	if s.GatewayMac != nil {
		structMap["gateway_mac"] = s.GatewayMac
	}
	if s.GatewayModel != nil {
		structMap["gateway_model"] = s.GatewayModel
	}
	if s.GatewayVersion != nil {
		structMap["gateway_version"] = s.GatewayVersion
	}
	if s.Name != nil {
		structMap["name"] = s.Name
	}
	if s.Total != nil {
		structMap["total"] = s.Total
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SleImpactedGatewaysGateway.
// It customizes the JSON unmarshaling process for SleImpactedGatewaysGateway objects.
func (s *SleImpactedGatewaysGateway) UnmarshalJSON(input []byte) error {
	var temp tempSleImpactedGatewaysGateway
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "degraded", "duration", "gateway_mac", "gateway_model", "gateway_version", "name", "total")
	if err != nil {
		return err
	}

	s.AdditionalProperties = additionalProperties
	s.Degraded = temp.Degraded
	s.Duration = temp.Duration
	s.GatewayMac = temp.GatewayMac
	s.GatewayModel = temp.GatewayModel
	s.GatewayVersion = temp.GatewayVersion
	s.Name = temp.Name
	s.Total = temp.Total
	return nil
}

// tempSleImpactedGatewaysGateway is a temporary struct used for validating the fields of SleImpactedGatewaysGateway.
type tempSleImpactedGatewaysGateway struct {
	Degraded       *float64 `json:"degraded,omitempty"`
	Duration       *int     `json:"duration,omitempty"`
	GatewayMac     *string  `json:"gateway_mac,omitempty"`
	GatewayModel   *string  `json:"gateway_model,omitempty"`
	GatewayVersion *string  `json:"gateway_version,omitempty"`
	Name           *string  `json:"name,omitempty"`
	Total          *int     `json:"total,omitempty"`
}
