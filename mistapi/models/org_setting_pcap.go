// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// OrgSettingPcap represents a OrgSettingPcap struct.
type OrgSettingPcap struct {
    Bucket               *string                `json:"bucket,omitempty"`
    // Max_len of non-management packets to capture
    MaxPktLen            *int                   `json:"max_pkt_len,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for OrgSettingPcap,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OrgSettingPcap) String() string {
    return fmt.Sprintf(
    	"OrgSettingPcap[Bucket=%v, MaxPktLen=%v, AdditionalProperties=%v]",
    	o.Bucket, o.MaxPktLen, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OrgSettingPcap.
// It customizes the JSON marshaling process for OrgSettingPcap objects.
func (o OrgSettingPcap) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(o.AdditionalProperties,
        "bucket", "max_pkt_len"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(o.toMap())
}

// toMap converts the OrgSettingPcap object to a map representation for JSON marshaling.
func (o OrgSettingPcap) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, o.AdditionalProperties)
    if o.Bucket != nil {
        structMap["bucket"] = o.Bucket
    }
    if o.MaxPktLen != nil {
        structMap["max_pkt_len"] = o.MaxPktLen
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSettingPcap.
// It customizes the JSON unmarshaling process for OrgSettingPcap objects.
func (o *OrgSettingPcap) UnmarshalJSON(input []byte) error {
    var temp tempOrgSettingPcap
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "bucket", "max_pkt_len")
    if err != nil {
    	return err
    }
    o.AdditionalProperties = additionalProperties
    
    o.Bucket = temp.Bucket
    o.MaxPktLen = temp.MaxPktLen
    return nil
}

// tempOrgSettingPcap is a temporary struct used for validating the fields of OrgSettingPcap.
type tempOrgSettingPcap  struct {
    Bucket    *string `json:"bucket,omitempty"`
    MaxPktLen *int    `json:"max_pkt_len,omitempty"`
}
