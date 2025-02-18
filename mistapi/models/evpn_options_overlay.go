package models

import (
    "encoding/json"
    "fmt"
)

// EvpnOptionsOverlay represents a EvpnOptionsOverlay struct.
type EvpnOptionsOverlay struct {
    // Overlay BGP Local AS Number
    As                   *int                   `json:"as,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for EvpnOptionsOverlay,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (e EvpnOptionsOverlay) String() string {
    return fmt.Sprintf(
    	"EvpnOptionsOverlay[As=%v, AdditionalProperties=%v]",
    	e.As, e.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for EvpnOptionsOverlay.
// It customizes the JSON marshaling process for EvpnOptionsOverlay objects.
func (e EvpnOptionsOverlay) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(e.AdditionalProperties,
        "as"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(e.toMap())
}

// toMap converts the EvpnOptionsOverlay object to a map representation for JSON marshaling.
func (e EvpnOptionsOverlay) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, e.AdditionalProperties)
    if e.As != nil {
        structMap["as"] = e.As
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for EvpnOptionsOverlay.
// It customizes the JSON unmarshaling process for EvpnOptionsOverlay objects.
func (e *EvpnOptionsOverlay) UnmarshalJSON(input []byte) error {
    var temp tempEvpnOptionsOverlay
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "as")
    if err != nil {
    	return err
    }
    e.AdditionalProperties = additionalProperties
    
    e.As = temp.As
    return nil
}

// tempEvpnOptionsOverlay is a temporary struct used for validating the fields of EvpnOptionsOverlay.
type tempEvpnOptionsOverlay  struct {
    As *int `json:"as,omitempty"`
}
