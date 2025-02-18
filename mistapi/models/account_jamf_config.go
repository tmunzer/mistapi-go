package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// AccountJamfConfig represents a AccountJamfConfig struct.
// OAuth linked Jamf apps account details
type AccountJamfConfig struct {
    // Customer account api client id. Required if `app_name`==`crowdstrike`
    ClientId             string                 `json:"client_id"`
    // Customer account api client secret
    ClientSecret         string                 `json:"client_secret"`
    // Customer account Jamf instance URL
    InstanceUrl          string                 `json:"instance_url"`
    // Smart group membership for determining compliance status
    SmartgroupName       string                 `json:"smartgroup_name"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for AccountJamfConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AccountJamfConfig) String() string {
    return fmt.Sprintf(
    	"AccountJamfConfig[ClientId=%v, ClientSecret=%v, InstanceUrl=%v, SmartgroupName=%v, AdditionalProperties=%v]",
    	a.ClientId, a.ClientSecret, a.InstanceUrl, a.SmartgroupName, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for AccountJamfConfig.
// It customizes the JSON marshaling process for AccountJamfConfig objects.
func (a AccountJamfConfig) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "client_id", "client_secret", "instance_url", "smartgroup_name"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the AccountJamfConfig object to a map representation for JSON marshaling.
func (a AccountJamfConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    structMap["client_id"] = a.ClientId
    structMap["client_secret"] = a.ClientSecret
    structMap["instance_url"] = a.InstanceUrl
    structMap["smartgroup_name"] = a.SmartgroupName
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AccountJamfConfig.
// It customizes the JSON unmarshaling process for AccountJamfConfig objects.
func (a *AccountJamfConfig) UnmarshalJSON(input []byte) error {
    var temp tempAccountJamfConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "client_id", "client_secret", "instance_url", "smartgroup_name")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.ClientId = *temp.ClientId
    a.ClientSecret = *temp.ClientSecret
    a.InstanceUrl = *temp.InstanceUrl
    a.SmartgroupName = *temp.SmartgroupName
    return nil
}

// tempAccountJamfConfig is a temporary struct used for validating the fields of AccountJamfConfig.
type tempAccountJamfConfig  struct {
    ClientId       *string `json:"client_id"`
    ClientSecret   *string `json:"client_secret"`
    InstanceUrl    *string `json:"instance_url"`
    SmartgroupName *string `json:"smartgroup_name"`
}

func (a *tempAccountJamfConfig) validate() error {
    var errs []string
    if a.ClientId == nil {
        errs = append(errs, "required field `client_id` is missing for type `account_jamf_config`")
    }
    if a.ClientSecret == nil {
        errs = append(errs, "required field `client_secret` is missing for type `account_jamf_config`")
    }
    if a.InstanceUrl == nil {
        errs = append(errs, "required field `instance_url` is missing for type `account_jamf_config`")
    }
    if a.SmartgroupName == nil {
        errs = append(errs, "required field `smartgroup_name` is missing for type `account_jamf_config`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
