package models

import (
    "encoding/json"
    "fmt"
)

// VpnPathPeerPathsPeer represents a VpnPathPeerPathsPeer struct.
// Preference indicates which outgoing wan should be preferred
type VpnPathPeerPathsPeer struct {
    Preference           *int                   `json:"preference,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for VpnPathPeerPathsPeer,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (v VpnPathPeerPathsPeer) String() string {
    return fmt.Sprintf(
    	"VpnPathPeerPathsPeer[Preference=%v, AdditionalProperties=%v]",
    	v.Preference, v.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for VpnPathPeerPathsPeer.
// It customizes the JSON marshaling process for VpnPathPeerPathsPeer objects.
func (v VpnPathPeerPathsPeer) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(v.AdditionalProperties,
        "preference"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(v.toMap())
}

// toMap converts the VpnPathPeerPathsPeer object to a map representation for JSON marshaling.
func (v VpnPathPeerPathsPeer) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, v.AdditionalProperties)
    if v.Preference != nil {
        structMap["preference"] = v.Preference
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for VpnPathPeerPathsPeer.
// It customizes the JSON unmarshaling process for VpnPathPeerPathsPeer objects.
func (v *VpnPathPeerPathsPeer) UnmarshalJSON(input []byte) error {
    var temp tempVpnPathPeerPathsPeer
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "preference")
    if err != nil {
    	return err
    }
    v.AdditionalProperties = additionalProperties
    
    v.Preference = temp.Preference
    return nil
}

// tempVpnPathPeerPathsPeer is a temporary struct used for validating the fields of VpnPathPeerPathsPeer.
type tempVpnPathPeerPathsPeer  struct {
    Preference *int `json:"preference,omitempty"`
}
