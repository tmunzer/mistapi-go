package models

import (
    "encoding/json"
)

// ResponseAsyncLicenseDetail represents a ResponseAsyncLicenseDetail struct.
// detail claim status per device
type ResponseAsyncLicenseDetail struct {
    // device MAC Address
    Mac                  *string                `json:"mac,omitempty"`
    Status               *string                `json:"status,omitempty"`
    Timestamp            *float64               `json:"timestamp,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseAsyncLicenseDetail.
// It customizes the JSON marshaling process for ResponseAsyncLicenseDetail objects.
func (r ResponseAsyncLicenseDetail) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "mac", "status", "timestamp"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseAsyncLicenseDetail object to a map representation for JSON marshaling.
func (r ResponseAsyncLicenseDetail) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    if r.Mac != nil {
        structMap["mac"] = r.Mac
    }
    if r.Status != nil {
        structMap["status"] = r.Status
    }
    if r.Timestamp != nil {
        structMap["timestamp"] = r.Timestamp
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseAsyncLicenseDetail.
// It customizes the JSON unmarshaling process for ResponseAsyncLicenseDetail objects.
func (r *ResponseAsyncLicenseDetail) UnmarshalJSON(input []byte) error {
    var temp tempResponseAsyncLicenseDetail
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "mac", "status", "timestamp")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.Mac = temp.Mac
    r.Status = temp.Status
    r.Timestamp = temp.Timestamp
    return nil
}

// tempResponseAsyncLicenseDetail is a temporary struct used for validating the fields of ResponseAsyncLicenseDetail.
type tempResponseAsyncLicenseDetail  struct {
    Mac       *string  `json:"mac,omitempty"`
    Status    *string  `json:"status,omitempty"`
    Timestamp *float64 `json:"timestamp,omitempty"`
}