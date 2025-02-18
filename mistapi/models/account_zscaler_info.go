package models

import (
    "encoding/json"
    "fmt"
)

// AccountZscalerInfo represents a AccountZscalerInfo struct.
// OAuth linked Zscaler apps account details
type AccountZscalerInfo struct {
    CloudName            *string                `json:"cloud_name,omitempty"`
    PartnerKey           *string                `json:"partner_key,omitempty"`
    // Customer account user name
    Username             *string                `json:"username,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for AccountZscalerInfo,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AccountZscalerInfo) String() string {
    return fmt.Sprintf(
    	"AccountZscalerInfo[CloudName=%v, PartnerKey=%v, Username=%v, AdditionalProperties=%v]",
    	a.CloudName, a.PartnerKey, a.Username, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for AccountZscalerInfo.
// It customizes the JSON marshaling process for AccountZscalerInfo objects.
func (a AccountZscalerInfo) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "cloud_name", "partner_key", "username"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the AccountZscalerInfo object to a map representation for JSON marshaling.
func (a AccountZscalerInfo) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    if a.CloudName != nil {
        structMap["cloud_name"] = a.CloudName
    }
    if a.PartnerKey != nil {
        structMap["partner_key"] = a.PartnerKey
    }
    if a.Username != nil {
        structMap["username"] = a.Username
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AccountZscalerInfo.
// It customizes the JSON unmarshaling process for AccountZscalerInfo objects.
func (a *AccountZscalerInfo) UnmarshalJSON(input []byte) error {
    var temp tempAccountZscalerInfo
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "cloud_name", "partner_key", "username")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.CloudName = temp.CloudName
    a.PartnerKey = temp.PartnerKey
    a.Username = temp.Username
    return nil
}

// tempAccountZscalerInfo is a temporary struct used for validating the fields of AccountZscalerInfo.
type tempAccountZscalerInfo  struct {
    CloudName  *string `json:"cloud_name,omitempty"`
    PartnerKey *string `json:"partner_key,omitempty"`
    Username   *string `json:"username,omitempty"`
}
