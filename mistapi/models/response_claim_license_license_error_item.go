package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// ResponseClaimLicenseLicenseErrorItem represents a ResponseClaimLicenseLicenseErrorItem struct.
type ResponseClaimLicenseLicenseErrorItem struct {
	Order                string         `json:"order"`
	Reason               string         `json:"reason"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseClaimLicenseLicenseErrorItem.
// It customizes the JSON marshaling process for ResponseClaimLicenseLicenseErrorItem objects.
func (r ResponseClaimLicenseLicenseErrorItem) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseClaimLicenseLicenseErrorItem object to a map representation for JSON marshaling.
func (r ResponseClaimLicenseLicenseErrorItem) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, r.AdditionalProperties)
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
	additionalProperties, err := UnmarshalAdditionalProperties(input, "order", "reason")
	if err != nil {
		return err
	}

	r.AdditionalProperties = additionalProperties
	r.Order = *temp.Order
	r.Reason = *temp.Reason
	return nil
}

// tempResponseClaimLicenseLicenseErrorItem is a temporary struct used for validating the fields of ResponseClaimLicenseLicenseErrorItem.
type tempResponseClaimLicenseLicenseErrorItem struct {
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
	return errors.New(strings.Join(errs, "\n"))
}
