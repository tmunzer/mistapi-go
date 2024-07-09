package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// ResponseUpgradeSiteDevices represents a ResponseUpgradeSiteDevices struct.
type ResponseUpgradeSiteDevices struct {
    UpgradeId            uuid.UUID      `json:"upgrade_id"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseUpgradeSiteDevices.
// It customizes the JSON marshaling process for ResponseUpgradeSiteDevices objects.
func (r ResponseUpgradeSiteDevices) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseUpgradeSiteDevices object to a map representation for JSON marshaling.
func (r ResponseUpgradeSiteDevices) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["upgrade_id"] = r.UpgradeId
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseUpgradeSiteDevices.
// It customizes the JSON unmarshaling process for ResponseUpgradeSiteDevices objects.
func (r *ResponseUpgradeSiteDevices) UnmarshalJSON(input []byte) error {
    var temp responseUpgradeSiteDevices
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "upgrade_id")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.UpgradeId = *temp.UpgradeId
    return nil
}

// responseUpgradeSiteDevices is a temporary struct used for validating the fields of ResponseUpgradeSiteDevices.
type responseUpgradeSiteDevices  struct {
    UpgradeId *uuid.UUID `json:"upgrade_id"`
}

func (r *responseUpgradeSiteDevices) validate() error {
    var errs []string
    if r.UpgradeId == nil {
        errs = append(errs, "required field `upgrade_id` is missing for type `Response_Upgrade_Site_Devices`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
