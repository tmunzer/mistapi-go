package models

import (
    "encoding/json"
    "fmt"
)

// TunnelProviderOptionsJse represents a TunnelProviderOptionsJse struct.
// For jse-ipsec, this allow provisioning of adequate resource on JSE. Make sure adequate licenses are added
type TunnelProviderOptionsJse struct {
    NumUsers             *int                   `json:"num_users,omitempty"`
    // JSE Organization name. The list of available organizations can be retrieved with the [Get Org JSE Info]($e/Orgs%20JSE/getOrgJseInfo) API Call
    OrgName              *string                `json:"org_name,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for TunnelProviderOptionsJse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (t TunnelProviderOptionsJse) String() string {
    return fmt.Sprintf(
    	"TunnelProviderOptionsJse[NumUsers=%v, OrgName=%v, AdditionalProperties=%v]",
    	t.NumUsers, t.OrgName, t.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for TunnelProviderOptionsJse.
// It customizes the JSON marshaling process for TunnelProviderOptionsJse objects.
func (t TunnelProviderOptionsJse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(t.AdditionalProperties,
        "num_users", "org_name"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(t.toMap())
}

// toMap converts the TunnelProviderOptionsJse object to a map representation for JSON marshaling.
func (t TunnelProviderOptionsJse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, t.AdditionalProperties)
    if t.NumUsers != nil {
        structMap["num_users"] = t.NumUsers
    }
    if t.OrgName != nil {
        structMap["org_name"] = t.OrgName
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TunnelProviderOptionsJse.
// It customizes the JSON unmarshaling process for TunnelProviderOptionsJse objects.
func (t *TunnelProviderOptionsJse) UnmarshalJSON(input []byte) error {
    var temp tempTunnelProviderOptionsJse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "num_users", "org_name")
    if err != nil {
    	return err
    }
    t.AdditionalProperties = additionalProperties
    
    t.NumUsers = temp.NumUsers
    t.OrgName = temp.OrgName
    return nil
}

// tempTunnelProviderOptionsJse is a temporary struct used for validating the fields of TunnelProviderOptionsJse.
type tempTunnelProviderOptionsJse  struct {
    NumUsers *int    `json:"num_users,omitempty"`
    OrgName  *string `json:"org_name,omitempty"`
}
