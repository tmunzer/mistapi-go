package models

import (
    "encoding/json"
    "fmt"
)

// AccountSkyatpInfo represents a AccountSkyatpInfo struct.
type AccountSkyatpInfo struct {
    Realm                *string                `json:"realm,omitempty"`
    Username             *string                `json:"username,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for AccountSkyatpInfo,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AccountSkyatpInfo) String() string {
    return fmt.Sprintf(
    	"AccountSkyatpInfo[Realm=%v, Username=%v, AdditionalProperties=%v]",
    	a.Realm, a.Username, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for AccountSkyatpInfo.
// It customizes the JSON marshaling process for AccountSkyatpInfo objects.
func (a AccountSkyatpInfo) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "realm", "username"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the AccountSkyatpInfo object to a map representation for JSON marshaling.
func (a AccountSkyatpInfo) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    if a.Realm != nil {
        structMap["realm"] = a.Realm
    }
    if a.Username != nil {
        structMap["username"] = a.Username
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AccountSkyatpInfo.
// It customizes the JSON unmarshaling process for AccountSkyatpInfo objects.
func (a *AccountSkyatpInfo) UnmarshalJSON(input []byte) error {
    var temp tempAccountSkyatpInfo
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "realm", "username")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.Realm = temp.Realm
    a.Username = temp.Username
    return nil
}

// tempAccountSkyatpInfo is a temporary struct used for validating the fields of AccountSkyatpInfo.
type tempAccountSkyatpInfo  struct {
    Realm    *string `json:"realm,omitempty"`
    Username *string `json:"username,omitempty"`
}
