package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// AccountJseConfig represents a AccountJseConfig struct.
type AccountJseConfig struct {
    CloudName            *string                `json:"cloud_name,omitempty"`
    Password             string                 `json:"password"`
    Username             string                 `json:"username"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for AccountJseConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AccountJseConfig) String() string {
    return fmt.Sprintf(
    	"AccountJseConfig[CloudName=%v, Password=%v, Username=%v, AdditionalProperties=%v]",
    	a.CloudName, a.Password, a.Username, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for AccountJseConfig.
// It customizes the JSON marshaling process for AccountJseConfig objects.
func (a AccountJseConfig) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "cloud_name", "password", "username"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the AccountJseConfig object to a map representation for JSON marshaling.
func (a AccountJseConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    if a.CloudName != nil {
        structMap["cloud_name"] = a.CloudName
    }
    structMap["password"] = a.Password
    structMap["username"] = a.Username
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AccountJseConfig.
// It customizes the JSON unmarshaling process for AccountJseConfig objects.
func (a *AccountJseConfig) UnmarshalJSON(input []byte) error {
    var temp tempAccountJseConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "cloud_name", "password", "username")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.CloudName = temp.CloudName
    a.Password = *temp.Password
    a.Username = *temp.Username
    return nil
}

// tempAccountJseConfig is a temporary struct used for validating the fields of AccountJseConfig.
type tempAccountJseConfig  struct {
    CloudName *string `json:"cloud_name,omitempty"`
    Password  *string `json:"password"`
    Username  *string `json:"username"`
}

func (a *tempAccountJseConfig) validate() error {
    var errs []string
    if a.Password == nil {
        errs = append(errs, "required field `password` is missing for type `account_jse_config`")
    }
    if a.Username == nil {
        errs = append(errs, "required field `username` is missing for type `account_jse_config`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
