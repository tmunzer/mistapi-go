package models

import (
	"encoding/json"
)

// Snmpv3ConfigNotifyFilterItemContent represents a Snmpv3ConfigNotifyFilterItemContent struct.
type Snmpv3ConfigNotifyFilterItemContent struct {
	Include              *bool          `json:"include,omitempty"`
	Oid                  *string        `json:"oid,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for Snmpv3ConfigNotifyFilterItemContent.
// It customizes the JSON marshaling process for Snmpv3ConfigNotifyFilterItemContent objects.
func (s Snmpv3ConfigNotifyFilterItemContent) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the Snmpv3ConfigNotifyFilterItemContent object to a map representation for JSON marshaling.
func (s Snmpv3ConfigNotifyFilterItemContent) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Include != nil {
		structMap["include"] = s.Include
	}
	if s.Oid != nil {
		structMap["oid"] = s.Oid
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Snmpv3ConfigNotifyFilterItemContent.
// It customizes the JSON unmarshaling process for Snmpv3ConfigNotifyFilterItemContent objects.
func (s *Snmpv3ConfigNotifyFilterItemContent) UnmarshalJSON(input []byte) error {
	var temp tempSnmpv3ConfigNotifyFilterItemContent
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "include", "oid")
	if err != nil {
		return err
	}

	s.AdditionalProperties = additionalProperties
	s.Include = temp.Include
	s.Oid = temp.Oid
	return nil
}

// tempSnmpv3ConfigNotifyFilterItemContent is a temporary struct used for validating the fields of Snmpv3ConfigNotifyFilterItemContent.
type tempSnmpv3ConfigNotifyFilterItemContent struct {
	Include *bool   `json:"include,omitempty"`
	Oid     *string `json:"oid,omitempty"`
}
