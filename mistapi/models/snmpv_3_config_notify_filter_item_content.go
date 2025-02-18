package models

import (
    "encoding/json"
    "fmt"
)

// Snmpv3ConfigNotifyFilterItemContent represents a Snmpv3ConfigNotifyFilterItemContent struct.
type Snmpv3ConfigNotifyFilterItemContent struct {
    Include              *bool                  `json:"include,omitempty"`
    Oid                  *string                `json:"oid,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for Snmpv3ConfigNotifyFilterItemContent,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s Snmpv3ConfigNotifyFilterItemContent) String() string {
    return fmt.Sprintf(
    	"Snmpv3ConfigNotifyFilterItemContent[Include=%v, Oid=%v, AdditionalProperties=%v]",
    	s.Include, s.Oid, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Snmpv3ConfigNotifyFilterItemContent.
// It customizes the JSON marshaling process for Snmpv3ConfigNotifyFilterItemContent objects.
func (s Snmpv3ConfigNotifyFilterItemContent) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "include", "oid"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the Snmpv3ConfigNotifyFilterItemContent object to a map representation for JSON marshaling.
func (s Snmpv3ConfigNotifyFilterItemContent) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "include", "oid")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Include = temp.Include
    s.Oid = temp.Oid
    return nil
}

// tempSnmpv3ConfigNotifyFilterItemContent is a temporary struct used for validating the fields of Snmpv3ConfigNotifyFilterItemContent.
type tempSnmpv3ConfigNotifyFilterItemContent  struct {
    Include *bool   `json:"include,omitempty"`
    Oid     *string `json:"oid,omitempty"`
}
