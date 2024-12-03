package models

import (
    "encoding/json"
)

// ResponsePcapBucketConfig represents a ResponsePcapBucketConfig struct.
type ResponsePcapBucketConfig struct {
    Bucket               *string                `json:"bucket,omitempty"`
    Detail               *string                `json:"detail,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponsePcapBucketConfig.
// It customizes the JSON marshaling process for ResponsePcapBucketConfig objects.
func (r ResponsePcapBucketConfig) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "bucket", "detail"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponsePcapBucketConfig object to a map representation for JSON marshaling.
func (r ResponsePcapBucketConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    if r.Bucket != nil {
        structMap["bucket"] = r.Bucket
    }
    if r.Detail != nil {
        structMap["detail"] = r.Detail
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponsePcapBucketConfig.
// It customizes the JSON unmarshaling process for ResponsePcapBucketConfig objects.
func (r *ResponsePcapBucketConfig) UnmarshalJSON(input []byte) error {
    var temp tempResponsePcapBucketConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "bucket", "detail")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.Bucket = temp.Bucket
    r.Detail = temp.Detail
    return nil
}

// tempResponsePcapBucketConfig is a temporary struct used for validating the fields of ResponsePcapBucketConfig.
type tempResponsePcapBucketConfig  struct {
    Bucket *string `json:"bucket,omitempty"`
    Detail *string `json:"detail,omitempty"`
}
