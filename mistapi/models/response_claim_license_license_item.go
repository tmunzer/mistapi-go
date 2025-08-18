// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// ResponseClaimLicenseLicenseItem represents a ResponseClaimLicenseLicenseItem struct.
type ResponseClaimLicenseLicenseItem struct {
	End                  int                    `json:"end"`
	Quantity             int                    `json:"quantity"`
	Start                int                    `json:"start"`
	Type                 string                 `json:"type"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseClaimLicenseLicenseItem,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseClaimLicenseLicenseItem) String() string {
	return fmt.Sprintf(
		"ResponseClaimLicenseLicenseItem[End=%v, Quantity=%v, Start=%v, Type=%v, AdditionalProperties=%v]",
		r.End, r.Quantity, r.Start, r.Type, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseClaimLicenseLicenseItem.
// It customizes the JSON marshaling process for ResponseClaimLicenseLicenseItem objects.
func (r ResponseClaimLicenseLicenseItem) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"end", "quantity", "start", "type"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseClaimLicenseLicenseItem object to a map representation for JSON marshaling.
func (r ResponseClaimLicenseLicenseItem) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, r.AdditionalProperties)
	structMap["end"] = r.End
	structMap["quantity"] = r.Quantity
	structMap["start"] = r.Start
	structMap["type"] = r.Type
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseClaimLicenseLicenseItem.
// It customizes the JSON unmarshaling process for ResponseClaimLicenseLicenseItem objects.
func (r *ResponseClaimLicenseLicenseItem) UnmarshalJSON(input []byte) error {
	var temp tempResponseClaimLicenseLicenseItem
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "end", "quantity", "start", "type")
	if err != nil {
		return err
	}
	r.AdditionalProperties = additionalProperties

	r.End = *temp.End
	r.Quantity = *temp.Quantity
	r.Start = *temp.Start
	r.Type = *temp.Type
	return nil
}

// tempResponseClaimLicenseLicenseItem is a temporary struct used for validating the fields of ResponseClaimLicenseLicenseItem.
type tempResponseClaimLicenseLicenseItem struct {
	End      *int    `json:"end"`
	Quantity *int    `json:"quantity"`
	Start    *int    `json:"start"`
	Type     *string `json:"type"`
}

func (r *tempResponseClaimLicenseLicenseItem) validate() error {
	var errs []string
	if r.End == nil {
		errs = append(errs, "required field `end` is missing for type `response_claim_license_license_item`")
	}
	if r.Quantity == nil {
		errs = append(errs, "required field `quantity` is missing for type `response_claim_license_license_item`")
	}
	if r.Start == nil {
		errs = append(errs, "required field `start` is missing for type `response_claim_license_license_item`")
	}
	if r.Type == nil {
		errs = append(errs, "required field `type` is missing for type `response_claim_license_license_item`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
