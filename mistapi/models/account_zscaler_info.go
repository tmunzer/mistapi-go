package models

import (
    "encoding/json"
)

// AccountZscalerInfo represents a AccountZscalerInfo struct.
// OAuth linked Zscaler apps account details
type AccountZscalerInfo struct {
    CloudName            *string        `json:"cloud_name,omitempty"`
    PartnerKey           *string        `json:"partner_key,omitempty"`
    // customer account user name
    Username             *string        `json:"username,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for AccountZscalerInfo.
// It customizes the JSON marshaling process for AccountZscalerInfo objects.
func (a AccountZscalerInfo) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the AccountZscalerInfo object to a map representation for JSON marshaling.
func (a AccountZscalerInfo) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
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
    additionalProperties, err := UnmarshalAdditionalProperties(input, "cloud_name", "partner_key", "username")
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
