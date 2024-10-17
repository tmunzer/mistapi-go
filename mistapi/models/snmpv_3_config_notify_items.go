package models

import (
	"encoding/json"
)

// Snmpv3ConfigNotifyItems represents a Snmpv3ConfigNotifyItems struct.
type Snmpv3ConfigNotifyItems struct {
	Name *string `json:"name,omitempty"`
	Tag  *string `json:"tag,omitempty"`
	// enum: `inform`, `trap`
	Type                 *Snmpv3ConfigNotifyTypeEnum `json:"type,omitempty"`
	AdditionalProperties map[string]any              `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for Snmpv3ConfigNotifyItems.
// It customizes the JSON marshaling process for Snmpv3ConfigNotifyItems objects.
func (s Snmpv3ConfigNotifyItems) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the Snmpv3ConfigNotifyItems object to a map representation for JSON marshaling.
func (s Snmpv3ConfigNotifyItems) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Name != nil {
		structMap["name"] = s.Name
	}
	if s.Tag != nil {
		structMap["tag"] = s.Tag
	}
	if s.Type != nil {
		structMap["type"] = s.Type
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Snmpv3ConfigNotifyItems.
// It customizes the JSON unmarshaling process for Snmpv3ConfigNotifyItems objects.
func (s *Snmpv3ConfigNotifyItems) UnmarshalJSON(input []byte) error {
	var temp tempSnmpv3ConfigNotifyItems
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "name", "tag", "type")
	if err != nil {
		return err
	}

	s.AdditionalProperties = additionalProperties
	s.Name = temp.Name
	s.Tag = temp.Tag
	s.Type = temp.Type
	return nil
}

// tempSnmpv3ConfigNotifyItems is a temporary struct used for validating the fields of Snmpv3ConfigNotifyItems.
type tempSnmpv3ConfigNotifyItems struct {
	Name *string                     `json:"name,omitempty"`
	Tag  *string                     `json:"tag,omitempty"`
	Type *Snmpv3ConfigNotifyTypeEnum `json:"type,omitempty"`
}
