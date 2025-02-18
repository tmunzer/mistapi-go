package models

import (
    "encoding/json"
    "fmt"
)

// Proxy represents a Proxy struct.
// Proxy Configuration to talk to Mist
type Proxy struct {
    Url                  *string                `json:"url,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for Proxy,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (p Proxy) String() string {
    return fmt.Sprintf(
    	"Proxy[Url=%v, AdditionalProperties=%v]",
    	p.Url, p.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Proxy.
// It customizes the JSON marshaling process for Proxy objects.
func (p Proxy) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(p.AdditionalProperties,
        "url"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(p.toMap())
}

// toMap converts the Proxy object to a map representation for JSON marshaling.
func (p Proxy) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, p.AdditionalProperties)
    if p.Url != nil {
        structMap["url"] = p.Url
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Proxy.
// It customizes the JSON unmarshaling process for Proxy objects.
func (p *Proxy) UnmarshalJSON(input []byte) error {
    var temp tempProxy
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "url")
    if err != nil {
    	return err
    }
    p.AdditionalProperties = additionalProperties
    
    p.Url = temp.Url
    return nil
}

// tempProxy is a temporary struct used for validating the fields of Proxy.
type tempProxy  struct {
    Url *string `json:"url,omitempty"`
}
