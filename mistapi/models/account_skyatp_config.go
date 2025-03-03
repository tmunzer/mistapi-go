package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// AccountSkyatpConfig represents a AccountSkyatpConfig struct.
type AccountSkyatpConfig struct {
    Password             string                 `json:"password"`
    Realm                string                 `json:"realm"`
    Username             string                 `json:"username"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for AccountSkyatpConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AccountSkyatpConfig) String() string {
    return fmt.Sprintf(
    	"AccountSkyatpConfig[Password=%v, Realm=%v, Username=%v, AdditionalProperties=%v]",
    	a.Password, a.Realm, a.Username, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for AccountSkyatpConfig.
// It customizes the JSON marshaling process for AccountSkyatpConfig objects.
func (a AccountSkyatpConfig) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "password", "realm", "username"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the AccountSkyatpConfig object to a map representation for JSON marshaling.
func (a AccountSkyatpConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    structMap["password"] = a.Password
    structMap["realm"] = a.Realm
    structMap["username"] = a.Username
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AccountSkyatpConfig.
// It customizes the JSON unmarshaling process for AccountSkyatpConfig objects.
func (a *AccountSkyatpConfig) UnmarshalJSON(input []byte) error {
    var temp tempAccountSkyatpConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "password", "realm", "username")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.Password = *temp.Password
    a.Realm = *temp.Realm
    a.Username = *temp.Username
    return nil
}

// tempAccountSkyatpConfig is a temporary struct used for validating the fields of AccountSkyatpConfig.
type tempAccountSkyatpConfig  struct {
    Password *string `json:"password"`
    Realm    *string `json:"realm"`
    Username *string `json:"username"`
}

func (a *tempAccountSkyatpConfig) validate() error {
    var errs []string
    if a.Password == nil {
        errs = append(errs, "required field `password` is missing for type `account_skyatp_config`")
    }
    if a.Realm == nil {
        errs = append(errs, "required field `realm` is missing for type `account_skyatp_config`")
    }
    if a.Username == nil {
        errs = append(errs, "required field `username` is missing for type `account_skyatp_config`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
