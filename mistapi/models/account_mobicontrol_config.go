package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// AccountMobicontrolConfig represents a AccountMobicontrolConfig struct.
type AccountMobicontrolConfig struct {
    // Customer account Client ID
    ClientId             string                 `json:"client_id"`
    // Customer account Client Secret
    ClientSecret         string                 `json:"client_secret"`
    // Customer account MobiControl instance URL
    InstanceUrl          string                 `json:"instance_url"`
    // Customer account password instance URL
    Password             string                 `json:"password"`
    // Customer account username
    Username             string                 `json:"username"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for AccountMobicontrolConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AccountMobicontrolConfig) String() string {
    return fmt.Sprintf(
    	"AccountMobicontrolConfig[ClientId=%v, ClientSecret=%v, InstanceUrl=%v, Password=%v, Username=%v, AdditionalProperties=%v]",
    	a.ClientId, a.ClientSecret, a.InstanceUrl, a.Password, a.Username, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for AccountMobicontrolConfig.
// It customizes the JSON marshaling process for AccountMobicontrolConfig objects.
func (a AccountMobicontrolConfig) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "client_id", "client_secret", "instance_url", "password", "username"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the AccountMobicontrolConfig object to a map representation for JSON marshaling.
func (a AccountMobicontrolConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    structMap["client_id"] = a.ClientId
    structMap["client_secret"] = a.ClientSecret
    structMap["instance_url"] = a.InstanceUrl
    structMap["password"] = a.Password
    structMap["username"] = a.Username
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AccountMobicontrolConfig.
// It customizes the JSON unmarshaling process for AccountMobicontrolConfig objects.
func (a *AccountMobicontrolConfig) UnmarshalJSON(input []byte) error {
    var temp tempAccountMobicontrolConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "client_id", "client_secret", "instance_url", "password", "username")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.ClientId = *temp.ClientId
    a.ClientSecret = *temp.ClientSecret
    a.InstanceUrl = *temp.InstanceUrl
    a.Password = *temp.Password
    a.Username = *temp.Username
    return nil
}

// tempAccountMobicontrolConfig is a temporary struct used for validating the fields of AccountMobicontrolConfig.
type tempAccountMobicontrolConfig  struct {
    ClientId     *string `json:"client_id"`
    ClientSecret *string `json:"client_secret"`
    InstanceUrl  *string `json:"instance_url"`
    Password     *string `json:"password"`
    Username     *string `json:"username"`
}

func (a *tempAccountMobicontrolConfig) validate() error {
    var errs []string
    if a.ClientId == nil {
        errs = append(errs, "required field `client_id` is missing for type `account_mobicontrol_config`")
    }
    if a.ClientSecret == nil {
        errs = append(errs, "required field `client_secret` is missing for type `account_mobicontrol_config`")
    }
    if a.InstanceUrl == nil {
        errs = append(errs, "required field `instance_url` is missing for type `account_mobicontrol_config`")
    }
    if a.Password == nil {
        errs = append(errs, "required field `password` is missing for type `account_mobicontrol_config`")
    }
    if a.Username == nil {
        errs = append(errs, "required field `username` is missing for type `account_mobicontrol_config`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
