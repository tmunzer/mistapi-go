package models

import (
    "encoding/json"
    "fmt"
)

// ResponseAutoplacementDevice represents a ResponseAutoplacementDevice struct.
type ResponseAutoplacementDevice struct {
    // Provides the reason for the status if the AP is invalid.
    Reason               *string                `json:"reason,omitempty"`
    // Indicates whether the ap is valid.
    Valid                *bool                  `json:"valid,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseAutoplacementDevice,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseAutoplacementDevice) String() string {
    return fmt.Sprintf(
    	"ResponseAutoplacementDevice[Reason=%v, Valid=%v, AdditionalProperties=%v]",
    	r.Reason, r.Valid, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseAutoplacementDevice.
// It customizes the JSON marshaling process for ResponseAutoplacementDevice objects.
func (r ResponseAutoplacementDevice) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "reason", "valid"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseAutoplacementDevice object to a map representation for JSON marshaling.
func (r ResponseAutoplacementDevice) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    if r.Reason != nil {
        structMap["reason"] = r.Reason
    }
    if r.Valid != nil {
        structMap["valid"] = r.Valid
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseAutoplacementDevice.
// It customizes the JSON unmarshaling process for ResponseAutoplacementDevice objects.
func (r *ResponseAutoplacementDevice) UnmarshalJSON(input []byte) error {
    var temp tempResponseAutoplacementDevice
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "reason", "valid")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.Reason = temp.Reason
    r.Valid = temp.Valid
    return nil
}

// tempResponseAutoplacementDevice is a temporary struct used for validating the fields of ResponseAutoplacementDevice.
type tempResponseAutoplacementDevice  struct {
    Reason *string `json:"reason,omitempty"`
    Valid  *bool   `json:"valid,omitempty"`
}
