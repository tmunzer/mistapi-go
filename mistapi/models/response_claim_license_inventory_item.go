package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// ResponseClaimLicenseInventoryItem represents a ResponseClaimLicenseInventoryItem struct.
type ResponseClaimLicenseInventoryItem struct {
    Mac                  string                 `json:"mac"`
    Magic                string                 `json:"magic"`
    Model                string                 `json:"model"`
    Serial               string                 `json:"serial"`
    Type                 string                 `json:"type"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseClaimLicenseInventoryItem,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseClaimLicenseInventoryItem) String() string {
    return fmt.Sprintf(
    	"ResponseClaimLicenseInventoryItem[Mac=%v, Magic=%v, Model=%v, Serial=%v, Type=%v, AdditionalProperties=%v]",
    	r.Mac, r.Magic, r.Model, r.Serial, r.Type, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseClaimLicenseInventoryItem.
// It customizes the JSON marshaling process for ResponseClaimLicenseInventoryItem objects.
func (r ResponseClaimLicenseInventoryItem) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "mac", "magic", "model", "serial", "type"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseClaimLicenseInventoryItem object to a map representation for JSON marshaling.
func (r ResponseClaimLicenseInventoryItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
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
    var temp tempResponseClaimLicenseInventoryItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "mac", "magic", "model", "serial", "type")
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

// tempResponseClaimLicenseInventoryItem is a temporary struct used for validating the fields of ResponseClaimLicenseInventoryItem.
type tempResponseClaimLicenseInventoryItem  struct {
    Mac    *string `json:"mac"`
    Magic  *string `json:"magic"`
    Model  *string `json:"model"`
    Serial *string `json:"serial"`
    Type   *string `json:"type"`
}

func (r *tempResponseClaimLicenseInventoryItem) validate() error {
    var errs []string
    if r.Mac == nil {
        errs = append(errs, "required field `mac` is missing for type `response_claim_license_inventory_item`")
    }
    if r.Magic == nil {
        errs = append(errs, "required field `magic` is missing for type `response_claim_license_inventory_item`")
    }
    if r.Model == nil {
        errs = append(errs, "required field `model` is missing for type `response_claim_license_inventory_item`")
    }
    if r.Serial == nil {
        errs = append(errs, "required field `serial` is missing for type `response_claim_license_inventory_item`")
    }
    if r.Type == nil {
        errs = append(errs, "required field `type` is missing for type `response_claim_license_inventory_item`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
