package models

import (
    "encoding/json"
)

// Proxy represents a Proxy struct.
// Proxy Configuration to talk to Mist
type Proxy struct {
    Url                  *string        `json:"url,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for Proxy.
// It customizes the JSON marshaling process for Proxy objects.
func (p Proxy) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the Proxy object to a map representation for JSON marshaling.
func (p Proxy) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, p.AdditionalProperties)
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
    additionalProperties, err := UnmarshalAdditionalProperties(input, "url")
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
