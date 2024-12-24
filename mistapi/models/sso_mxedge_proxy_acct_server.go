package models

import (
    "encoding/json"
    "fmt"
)

// SsoMxedgeProxyAcctServer represents a SsoMxedgeProxyAcctServer struct.
type SsoMxedgeProxyAcctServer struct {
    Host                 *string                `json:"host,omitempty"`
    Port                 *int                   `json:"port,omitempty"`
    Secret               *string                `json:"secret,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SsoMxedgeProxyAcctServer,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SsoMxedgeProxyAcctServer) String() string {
    return fmt.Sprintf(
    	"SsoMxedgeProxyAcctServer[Host=%v, Port=%v, Secret=%v, AdditionalProperties=%v]",
    	s.Host, s.Port, s.Secret, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SsoMxedgeProxyAcctServer.
// It customizes the JSON marshaling process for SsoMxedgeProxyAcctServer objects.
func (s SsoMxedgeProxyAcctServer) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "host", "port", "secret"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SsoMxedgeProxyAcctServer object to a map representation for JSON marshaling.
func (s SsoMxedgeProxyAcctServer) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Host != nil {
        structMap["host"] = s.Host
    }
    if s.Port != nil {
        structMap["port"] = s.Port
    }
    if s.Secret != nil {
        structMap["secret"] = s.Secret
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SsoMxedgeProxyAcctServer.
// It customizes the JSON unmarshaling process for SsoMxedgeProxyAcctServer objects.
func (s *SsoMxedgeProxyAcctServer) UnmarshalJSON(input []byte) error {
    var temp tempSsoMxedgeProxyAcctServer
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "host", "port", "secret")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Host = temp.Host
    s.Port = temp.Port
    s.Secret = temp.Secret
    return nil
}

// tempSsoMxedgeProxyAcctServer is a temporary struct used for validating the fields of SsoMxedgeProxyAcctServer.
type tempSsoMxedgeProxyAcctServer  struct {
    Host   *string `json:"host,omitempty"`
    Port   *int    `json:"port,omitempty"`
    Secret *string `json:"secret,omitempty"`
}
