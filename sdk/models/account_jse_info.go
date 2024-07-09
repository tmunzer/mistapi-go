package models

import (
    "encoding/json"
)

// AccountJseInfo represents a AccountJseInfo struct.
type AccountJseInfo struct {
    CloudName            *string        `json:"cloud_name,omitempty"`
    Username             *string        `json:"username,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for AccountJseInfo.
// It customizes the JSON marshaling process for AccountJseInfo objects.
func (a AccountJseInfo) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the AccountJseInfo object to a map representation for JSON marshaling.
func (a AccountJseInfo) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    if a.CloudName != nil {
        structMap["cloud_name"] = a.CloudName
    }
    if a.Username != nil {
        structMap["username"] = a.Username
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AccountJseInfo.
// It customizes the JSON unmarshaling process for AccountJseInfo objects.
func (a *AccountJseInfo) UnmarshalJSON(input []byte) error {
    var temp accountJseInfo
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "cloud_name", "username")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.CloudName = temp.CloudName
    a.Username = temp.Username
    return nil
}

// accountJseInfo is a temporary struct used for validating the fields of AccountJseInfo.
type accountJseInfo  struct {
    CloudName *string `json:"cloud_name,omitempty"`
    Username  *string `json:"username,omitempty"`
}
