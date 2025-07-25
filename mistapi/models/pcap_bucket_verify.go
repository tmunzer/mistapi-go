// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// PcapBucketVerify represents a PcapBucketVerify struct.
type PcapBucketVerify struct {
    Bucket               string                 `json:"bucket"`
    VerifyToken          string                 `json:"verify_token"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for PcapBucketVerify,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (p PcapBucketVerify) String() string {
    return fmt.Sprintf(
    	"PcapBucketVerify[Bucket=%v, VerifyToken=%v, AdditionalProperties=%v]",
    	p.Bucket, p.VerifyToken, p.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for PcapBucketVerify.
// It customizes the JSON marshaling process for PcapBucketVerify objects.
func (p PcapBucketVerify) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(p.AdditionalProperties,
        "bucket", "verify_token"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(p.toMap())
}

// toMap converts the PcapBucketVerify object to a map representation for JSON marshaling.
func (p PcapBucketVerify) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, p.AdditionalProperties)
    structMap["bucket"] = p.Bucket
    structMap["verify_token"] = p.VerifyToken
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PcapBucketVerify.
// It customizes the JSON unmarshaling process for PcapBucketVerify objects.
func (p *PcapBucketVerify) UnmarshalJSON(input []byte) error {
    var temp tempPcapBucketVerify
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "bucket", "verify_token")
    if err != nil {
    	return err
    }
    p.AdditionalProperties = additionalProperties
    
    p.Bucket = *temp.Bucket
    p.VerifyToken = *temp.VerifyToken
    return nil
}

// tempPcapBucketVerify is a temporary struct used for validating the fields of PcapBucketVerify.
type tempPcapBucketVerify  struct {
    Bucket      *string `json:"bucket"`
    VerifyToken *string `json:"verify_token"`
}

func (p *tempPcapBucketVerify) validate() error {
    var errs []string
    if p.Bucket == nil {
        errs = append(errs, "required field `bucket` is missing for type `pcap_bucket_verify`")
    }
    if p.VerifyToken == nil {
        errs = append(errs, "required field `verify_token` is missing for type `pcap_bucket_verify`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
