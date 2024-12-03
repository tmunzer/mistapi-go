package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// ResponseUpgradeSiteDevices represents a ResponseUpgradeSiteDevices struct.
type ResponseUpgradeSiteDevices struct {
    UpgradeId            uuid.UUID              `json:"upgrade_id"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseUpgradeSiteDevices.
// It customizes the JSON marshaling process for ResponseUpgradeSiteDevices objects.
func (r ResponseUpgradeSiteDevices) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "upgrade_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseUpgradeSiteDevices object to a map representation for JSON marshaling.
func (r ResponseUpgradeSiteDevices) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["upgrade_id"] = r.UpgradeId
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseUpgradeSiteDevices.
// It customizes the JSON unmarshaling process for ResponseUpgradeSiteDevices objects.
func (r *ResponseUpgradeSiteDevices) UnmarshalJSON(input []byte) error {
    var temp tempResponseUpgradeSiteDevices
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "upgrade_id")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.UpgradeId = *temp.UpgradeId
    return nil
}

// tempResponseUpgradeSiteDevices is a temporary struct used for validating the fields of ResponseUpgradeSiteDevices.
type tempResponseUpgradeSiteDevices  struct {
    UpgradeId *uuid.UUID `json:"upgrade_id"`
}

func (r *tempResponseUpgradeSiteDevices) validate() error {
    var errs []string
    if r.UpgradeId == nil {
        errs = append(errs, "required field `upgrade_id` is missing for type `response_upgrade_site_devices`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
