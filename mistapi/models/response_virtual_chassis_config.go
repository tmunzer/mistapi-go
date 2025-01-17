package models

import (
    "encoding/json"
    "fmt"
)

// ResponseVirtualChassisConfig represents a ResponseVirtualChassisConfig struct.
type ResponseVirtualChassisConfig struct {
    // Virtual Chassis
    Id                   *VirtualChassisConfig  `json:"id,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseVirtualChassisConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseVirtualChassisConfig) String() string {
    return fmt.Sprintf(
    	"ResponseVirtualChassisConfig[Id=%v, AdditionalProperties=%v]",
    	r.Id, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseVirtualChassisConfig.
// It customizes the JSON marshaling process for ResponseVirtualChassisConfig objects.
func (r ResponseVirtualChassisConfig) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseVirtualChassisConfig object to a map representation for JSON marshaling.
func (r ResponseVirtualChassisConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    if r.Id != nil {
        structMap["id"] = r.Id.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseVirtualChassisConfig.
// It customizes the JSON unmarshaling process for ResponseVirtualChassisConfig objects.
func (r *ResponseVirtualChassisConfig) UnmarshalJSON(input []byte) error {
    var temp tempResponseVirtualChassisConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.Id = temp.Id
    return nil
}

// tempResponseVirtualChassisConfig is a temporary struct used for validating the fields of ResponseVirtualChassisConfig.
type tempResponseVirtualChassisConfig  struct {
    Id *VirtualChassisConfig `json:"id,omitempty"`
}
