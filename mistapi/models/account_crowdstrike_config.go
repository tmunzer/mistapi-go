package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// AccountCrowdstrikeConfig represents a AccountCrowdstrikeConfig struct.
// OAuth linked CrowdStrike apps account details
type AccountCrowdstrikeConfig struct {
    // customer account api client ID
    ClientId             string                 `json:"client_id"`
    // customer account api client Secret
    ClientSecret         string                 `json:"client_secret"`
    // customer id of an admin
    CustomerId           string                 `json:"customer_id"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for AccountCrowdstrikeConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AccountCrowdstrikeConfig) String() string {
    return fmt.Sprintf(
    	"AccountCrowdstrikeConfig[ClientId=%v, ClientSecret=%v, CustomerId=%v, AdditionalProperties=%v]",
    	a.ClientId, a.ClientSecret, a.CustomerId, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for AccountCrowdstrikeConfig.
// It customizes the JSON marshaling process for AccountCrowdstrikeConfig objects.
func (a AccountCrowdstrikeConfig) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "client_id", "client_secret", "customer_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the AccountCrowdstrikeConfig object to a map representation for JSON marshaling.
func (a AccountCrowdstrikeConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    structMap["client_id"] = a.ClientId
    structMap["client_secret"] = a.ClientSecret
    structMap["customer_id"] = a.CustomerId
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AccountCrowdstrikeConfig.
// It customizes the JSON unmarshaling process for AccountCrowdstrikeConfig objects.
func (a *AccountCrowdstrikeConfig) UnmarshalJSON(input []byte) error {
    var temp tempAccountCrowdstrikeConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "client_id", "client_secret", "customer_id")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.ClientId = *temp.ClientId
    a.ClientSecret = *temp.ClientSecret
    a.CustomerId = *temp.CustomerId
    return nil
}

// tempAccountCrowdstrikeConfig is a temporary struct used for validating the fields of AccountCrowdstrikeConfig.
type tempAccountCrowdstrikeConfig  struct {
    ClientId     *string `json:"client_id"`
    ClientSecret *string `json:"client_secret"`
    CustomerId   *string `json:"customer_id"`
}

func (a *tempAccountCrowdstrikeConfig) validate() error {
    var errs []string
    if a.ClientId == nil {
        errs = append(errs, "required field `client_id` is missing for type `account_crowdstrike_config`")
    }
    if a.ClientSecret == nil {
        errs = append(errs, "required field `client_secret` is missing for type `account_crowdstrike_config`")
    }
    if a.CustomerId == nil {
        errs = append(errs, "required field `customer_id` is missing for type `account_crowdstrike_config`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
