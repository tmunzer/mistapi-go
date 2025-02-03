package models

import (
    "encoding/json"
    "fmt"
)

// AutoOrient represents a AutoOrient struct.
type AutoOrient struct {
    // If `force_collection`==`false`, the API attempts to start auto orientation with existing BLE data.
    // If `force_collection`==`true`, the API attempts to start BLE orchestration.
    ForceCollection      *bool                  `json:"force_collection,omitempty"`
    // List of device macs
    Macs                 []string               `json:"macs,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for AutoOrient,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AutoOrient) String() string {
    return fmt.Sprintf(
    	"AutoOrient[ForceCollection=%v, Macs=%v, AdditionalProperties=%v]",
    	a.ForceCollection, a.Macs, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for AutoOrient.
// It customizes the JSON marshaling process for AutoOrient objects.
func (a AutoOrient) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "force_collection", "macs"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the AutoOrient object to a map representation for JSON marshaling.
func (a AutoOrient) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    if a.ForceCollection != nil {
        structMap["force_collection"] = a.ForceCollection
    }
    if a.Macs != nil {
        structMap["macs"] = a.Macs
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AutoOrient.
// It customizes the JSON unmarshaling process for AutoOrient objects.
func (a *AutoOrient) UnmarshalJSON(input []byte) error {
    var temp tempAutoOrient
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "force_collection", "macs")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.ForceCollection = temp.ForceCollection
    a.Macs = temp.Macs
    return nil
}

// tempAutoOrient is a temporary struct used for validating the fields of AutoOrient.
type tempAutoOrient  struct {
    ForceCollection *bool    `json:"force_collection,omitempty"`
    Macs            []string `json:"macs,omitempty"`
}
