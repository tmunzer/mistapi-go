package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// AccountZscalerConfig represents a AccountZscalerConfig struct.
// OAuth linked Zscaler apps account details
type AccountZscalerConfig struct {
    CloudName            string         `json:"cloud_name"`
    PartnerKey           string         `json:"partner_key"`
    // customer account password
    Password             string         `json:"password"`
    // customer account user name
    Username             string         `json:"username"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for AccountZscalerConfig.
// It customizes the JSON marshaling process for AccountZscalerConfig objects.
func (a AccountZscalerConfig) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the AccountZscalerConfig object to a map representation for JSON marshaling.
func (a AccountZscalerConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    structMap["cloud_name"] = a.CloudName
    structMap["partner_key"] = a.PartnerKey
    structMap["password"] = a.Password
    structMap["username"] = a.Username
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AccountZscalerConfig.
// It customizes the JSON unmarshaling process for AccountZscalerConfig objects.
func (a *AccountZscalerConfig) UnmarshalJSON(input []byte) error {
    var temp tempAccountZscalerConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "cloud_name", "partner_key", "password", "username")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.CloudName = *temp.CloudName
    a.PartnerKey = *temp.PartnerKey
    a.Password = *temp.Password
    a.Username = *temp.Username
    return nil
}

// tempAccountZscalerConfig is a temporary struct used for validating the fields of AccountZscalerConfig.
type tempAccountZscalerConfig  struct {
    CloudName  *string `json:"cloud_name"`
    PartnerKey *string `json:"partner_key"`
    Password   *string `json:"password"`
    Username   *string `json:"username"`
}

func (a *tempAccountZscalerConfig) validate() error {
    var errs []string
    if a.CloudName == nil {
        errs = append(errs, "required field `cloud_name` is missing for type `account_zscaler_config`")
    }
    if a.PartnerKey == nil {
        errs = append(errs, "required field `partner_key` is missing for type `account_zscaler_config`")
    }
    if a.Password == nil {
        errs = append(errs, "required field `password` is missing for type `account_zscaler_config`")
    }
    if a.Username == nil {
        errs = append(errs, "required field `username` is missing for type `account_zscaler_config`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
