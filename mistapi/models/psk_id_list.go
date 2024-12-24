package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// PskIdList represents a PskIdList struct.
type PskIdList struct {
    PskIds               []uuid.UUID            `json:"psk_ids,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for PskIdList,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (p PskIdList) String() string {
    return fmt.Sprintf(
    	"PskIdList[PskIds=%v, AdditionalProperties=%v]",
    	p.PskIds, p.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for PskIdList.
// It customizes the JSON marshaling process for PskIdList objects.
func (p PskIdList) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(p.AdditionalProperties,
        "psk_ids"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(p.toMap())
}

// toMap converts the PskIdList object to a map representation for JSON marshaling.
func (p PskIdList) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, p.AdditionalProperties)
    if p.PskIds != nil {
        structMap["psk_ids"] = p.PskIds
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PskIdList.
// It customizes the JSON unmarshaling process for PskIdList objects.
func (p *PskIdList) UnmarshalJSON(input []byte) error {
    var temp tempPskIdList
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "psk_ids")
    if err != nil {
    	return err
    }
    p.AdditionalProperties = additionalProperties
    
    p.PskIds = temp.PskIds
    return nil
}

// tempPskIdList is a temporary struct used for validating the fields of PskIdList.
type tempPskIdList  struct {
    PskIds []uuid.UUID `json:"psk_ids,omitempty"`
}
