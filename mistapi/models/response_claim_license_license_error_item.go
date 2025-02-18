package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// ResponseClaimLicenseLicenseErrorItem represents a ResponseClaimLicenseLicenseErrorItem struct.
type ResponseClaimLicenseLicenseErrorItem struct {
    Order                string                 `json:"order"`
    Reason               string                 `json:"reason"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseClaimLicenseLicenseErrorItem,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseClaimLicenseLicenseErrorItem) String() string {
    return fmt.Sprintf(
    	"ResponseClaimLicenseLicenseErrorItem[Order=%v, Reason=%v, AdditionalProperties=%v]",
    	r.Order, r.Reason, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseClaimLicenseLicenseErrorItem.
// It customizes the JSON marshaling process for ResponseClaimLicenseLicenseErrorItem objects.
func (r ResponseClaimLicenseLicenseErrorItem) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "order", "reason"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseClaimLicenseLicenseErrorItem object to a map representation for JSON marshaling.
func (r ResponseClaimLicenseLicenseErrorItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["order"] = r.Order
    structMap["reason"] = r.Reason
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseClaimLicenseLicenseErrorItem.
// It customizes the JSON unmarshaling process for ResponseClaimLicenseLicenseErrorItem objects.
func (r *ResponseClaimLicenseLicenseErrorItem) UnmarshalJSON(input []byte) error {
    var temp tempResponseClaimLicenseLicenseErrorItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "order", "reason")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.Order = *temp.Order
    r.Reason = *temp.Reason
    return nil
}

// tempResponseClaimLicenseLicenseErrorItem is a temporary struct used for validating the fields of ResponseClaimLicenseLicenseErrorItem.
type tempResponseClaimLicenseLicenseErrorItem  struct {
    Order  *string `json:"order"`
    Reason *string `json:"reason"`
}

func (r *tempResponseClaimLicenseLicenseErrorItem) validate() error {
    var errs []string
    if r.Order == nil {
        errs = append(errs, "required field `order` is missing for type `response_claim_license_license_error_item`")
    }
    if r.Reason == nil {
        errs = append(errs, "required field `reason` is missing for type `response_claim_license_license_error_item`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
