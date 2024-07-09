package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ResponseClaimLicense represents a ResponseClaimLicense struct.
type ResponseClaimLicense struct {
    InventoryAdded       []ResponseClaimLicenseInventoryItem    `json:"inventory_added"`
    InventoryDuplicated  []ResponseClaimLicenseInventoryItem    `json:"inventory_duplicated"`
    LicenseAdded         []ResponseClaimLicenseLicenseItem      `json:"license_added"`
    LicenseDuplicated    []ResponseClaimLicenseLicenseItem      `json:"license_duplicated"`
    LicenseError         []ResponseClaimLicenseLicenseErrorItem `json:"license_error"`
    AdditionalProperties map[string]any                         `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseClaimLicense.
// It customizes the JSON marshaling process for ResponseClaimLicense objects.
func (r ResponseClaimLicense) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseClaimLicense object to a map representation for JSON marshaling.
func (r ResponseClaimLicense) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["inventory_added"] = r.InventoryAdded
    structMap["inventory_duplicated"] = r.InventoryDuplicated
    structMap["license_added"] = r.LicenseAdded
    structMap["license_duplicated"] = r.LicenseDuplicated
    structMap["license_error"] = r.LicenseError
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseClaimLicense.
// It customizes the JSON unmarshaling process for ResponseClaimLicense objects.
func (r *ResponseClaimLicense) UnmarshalJSON(input []byte) error {
    var temp responseClaimLicense
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "inventory_added", "inventory_duplicated", "license_added", "license_duplicated", "license_error")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.InventoryAdded = *temp.InventoryAdded
    r.InventoryDuplicated = *temp.InventoryDuplicated
    r.LicenseAdded = *temp.LicenseAdded
    r.LicenseDuplicated = *temp.LicenseDuplicated
    r.LicenseError = *temp.LicenseError
    return nil
}

// responseClaimLicense is a temporary struct used for validating the fields of ResponseClaimLicense.
type responseClaimLicense  struct {
    InventoryAdded      *[]ResponseClaimLicenseInventoryItem    `json:"inventory_added"`
    InventoryDuplicated *[]ResponseClaimLicenseInventoryItem    `json:"inventory_duplicated"`
    LicenseAdded        *[]ResponseClaimLicenseLicenseItem      `json:"license_added"`
    LicenseDuplicated   *[]ResponseClaimLicenseLicenseItem      `json:"license_duplicated"`
    LicenseError        *[]ResponseClaimLicenseLicenseErrorItem `json:"license_error"`
}

func (r *responseClaimLicense) validate() error {
    var errs []string
    if r.InventoryAdded == nil {
        errs = append(errs, "required field `inventory_added` is missing for type `Response_Claim_License`")
    }
    if r.InventoryDuplicated == nil {
        errs = append(errs, "required field `inventory_duplicated` is missing for type `Response_Claim_License`")
    }
    if r.LicenseAdded == nil {
        errs = append(errs, "required field `license_added` is missing for type `Response_Claim_License`")
    }
    if r.LicenseDuplicated == nil {
        errs = append(errs, "required field `license_duplicated` is missing for type `Response_Claim_License`")
    }
    if r.LicenseError == nil {
        errs = append(errs, "required field `license_error` is missing for type `Response_Claim_License`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
