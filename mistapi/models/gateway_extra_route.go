package models

import (
    "encoding/json"
    "fmt"
)

// GatewayExtraRoute represents a GatewayExtraRoute struct.
type GatewayExtraRoute struct {
    Via                  *string                `json:"via,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for GatewayExtraRoute,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (g GatewayExtraRoute) String() string {
    return fmt.Sprintf(
    	"GatewayExtraRoute[Via=%v, AdditionalProperties=%v]",
    	g.Via, g.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for GatewayExtraRoute.
// It customizes the JSON marshaling process for GatewayExtraRoute objects.
func (g GatewayExtraRoute) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(g.AdditionalProperties,
        "via"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(g.toMap())
}

// toMap converts the GatewayExtraRoute object to a map representation for JSON marshaling.
func (g GatewayExtraRoute) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, g.AdditionalProperties)
    if g.Via != nil {
        structMap["via"] = g.Via
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GatewayExtraRoute.
// It customizes the JSON unmarshaling process for GatewayExtraRoute objects.
func (g *GatewayExtraRoute) UnmarshalJSON(input []byte) error {
    var temp tempGatewayExtraRoute
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "via")
    if err != nil {
    	return err
    }
    g.AdditionalProperties = additionalProperties
    
    g.Via = temp.Via
    return nil
}

// tempGatewayExtraRoute is a temporary struct used for validating the fields of GatewayExtraRoute.
type tempGatewayExtraRoute  struct {
    Via *string `json:"via,omitempty"`
}
