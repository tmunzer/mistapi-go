package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ResponseClaimLicenseInventoryItem represents a ResponseClaimLicenseInventoryItem struct.
type ResponseClaimLicenseInventoryItem struct {
    Mac                  string         `json:"mac"`
    Magic                string         `json:"magic"`
    Model                string         `json:"model"`
    Serial               string         `json:"serial"`
    Type                 string         `json:"type"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseClaimLicenseInventoryItem.
// It customizes the JSON marshaling process for ResponseClaimLicenseInventoryItem objects.
func (r ResponseClaimLicenseInventoryItem) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseClaimLicenseInventoryItem object to a map representation for JSON marshaling.
func (r ResponseClaimLicenseInventoryItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["mac"] = r.Mac
    structMap["magic"] = r.Magic
    structMap["model"] = r.Model
    structMap["serial"] = r.Serial
    structMap["type"] = r.Type
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseClaimLicenseInventoryItem.
// It customizes the JSON unmarshaling process for ResponseClaimLicenseInventoryItem objects.
func (r *ResponseClaimLicenseInventoryItem) UnmarshalJSON(input []byte) error {
    var temp responseClaimLicenseInventoryItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "mac", "magic", "model", "serial", "type")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.Mac = *temp.Mac
    r.Magic = *temp.Magic
    r.Model = *temp.Model
    r.Serial = *temp.Serial
    r.Type = *temp.Type
    return nil
}

// responseClaimLicenseInventoryItem is a temporary struct used for validating the fields of ResponseClaimLicenseInventoryItem.
type responseClaimLicenseInventoryItem  struct {
    Mac    *string `json:"mac"`
    Magic  *string `json:"magic"`
    Model  *string `json:"model"`
    Serial *string `json:"serial"`
    Type   *string `json:"type"`
}

func (r *responseClaimLicenseInventoryItem) validate() error {
    var errs []string
    if r.Mac == nil {
        errs = append(errs, "required field `mac` is missing for type `Response_Claim_License_Inventory_Item`")
    }
    if r.Magic == nil {
        errs = append(errs, "required field `magic` is missing for type `Response_Claim_License_Inventory_Item`")
    }
    if r.Model == nil {
        errs = append(errs, "required field `model` is missing for type `Response_Claim_License_Inventory_Item`")
    }
    if r.Serial == nil {
        errs = append(errs, "required field `serial` is missing for type `Response_Claim_License_Inventory_Item`")
    }
    if r.Type == nil {
        errs = append(errs, "required field `type` is missing for type `Response_Claim_License_Inventory_Item`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
