package models

import (
    "encoding/json"
    "fmt"
)

// GatewayMgmtHostOutPolicySyslog represents a GatewayMgmtHostOutPolicySyslog struct.
type GatewayMgmtHostOutPolicySyslog struct {
    PathPreference       *string                                `json:"path_preference,omitempty"`
    Servers              []GatewayMgmtHostOutPolicySyslogServer `json:"servers,omitempty"`
    AdditionalProperties map[string]interface{}                 `json:"_"`
}

// String implements the fmt.Stringer interface for GatewayMgmtHostOutPolicySyslog,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (g GatewayMgmtHostOutPolicySyslog) String() string {
    return fmt.Sprintf(
    	"GatewayMgmtHostOutPolicySyslog[PathPreference=%v, Servers=%v, AdditionalProperties=%v]",
    	g.PathPreference, g.Servers, g.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for GatewayMgmtHostOutPolicySyslog.
// It customizes the JSON marshaling process for GatewayMgmtHostOutPolicySyslog objects.
func (g GatewayMgmtHostOutPolicySyslog) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(g.AdditionalProperties,
        "path_preference", "servers"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(g.toMap())
}

// toMap converts the GatewayMgmtHostOutPolicySyslog object to a map representation for JSON marshaling.
func (g GatewayMgmtHostOutPolicySyslog) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, g.AdditionalProperties)
    if g.PathPreference != nil {
        structMap["path_preference"] = g.PathPreference
    }
    if g.Servers != nil {
        structMap["servers"] = g.Servers
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GatewayMgmtHostOutPolicySyslog.
// It customizes the JSON unmarshaling process for GatewayMgmtHostOutPolicySyslog objects.
func (g *GatewayMgmtHostOutPolicySyslog) UnmarshalJSON(input []byte) error {
    var temp tempGatewayMgmtHostOutPolicySyslog
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "path_preference", "servers")
    if err != nil {
    	return err
    }
    g.AdditionalProperties = additionalProperties
    
    g.PathPreference = temp.PathPreference
    g.Servers = temp.Servers
    return nil
}

// tempGatewayMgmtHostOutPolicySyslog is a temporary struct used for validating the fields of GatewayMgmtHostOutPolicySyslog.
type tempGatewayMgmtHostOutPolicySyslog  struct {
    PathPreference *string                                `json:"path_preference,omitempty"`
    Servers        []GatewayMgmtHostOutPolicySyslogServer `json:"servers,omitempty"`
}
