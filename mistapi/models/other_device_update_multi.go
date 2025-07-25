// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/google/uuid"
    "strings"
)

// OtherDeviceUpdateMulti represents a OtherDeviceUpdateMulti struct.
type OtherDeviceUpdateMulti struct {
    // MAC address of the peer device.
    Macs                 []string                       `json:"macs,omitempty"`
    // The operation being performed. enum: `assign`, `unassign`
    Op                   OtherDeviceUpdateOperationEnum `json:"op"`
    SiteId               *uuid.UUID                     `json:"site_id,omitempty"`
    AdditionalProperties map[string]interface{}         `json:"_"`
}

// String implements the fmt.Stringer interface for OtherDeviceUpdateMulti,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OtherDeviceUpdateMulti) String() string {
    return fmt.Sprintf(
    	"OtherDeviceUpdateMulti[Macs=%v, Op=%v, SiteId=%v, AdditionalProperties=%v]",
    	o.Macs, o.Op, o.SiteId, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OtherDeviceUpdateMulti.
// It customizes the JSON marshaling process for OtherDeviceUpdateMulti objects.
func (o OtherDeviceUpdateMulti) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(o.AdditionalProperties,
        "macs", "op", "site_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(o.toMap())
}

// toMap converts the OtherDeviceUpdateMulti object to a map representation for JSON marshaling.
func (o OtherDeviceUpdateMulti) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, o.AdditionalProperties)
    if o.Macs != nil {
        structMap["macs"] = o.Macs
    }
    structMap["op"] = o.Op
    if o.SiteId != nil {
        structMap["site_id"] = o.SiteId
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OtherDeviceUpdateMulti.
// It customizes the JSON unmarshaling process for OtherDeviceUpdateMulti objects.
func (o *OtherDeviceUpdateMulti) UnmarshalJSON(input []byte) error {
    var temp tempOtherDeviceUpdateMulti
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "macs", "op", "site_id")
    if err != nil {
    	return err
    }
    o.AdditionalProperties = additionalProperties
    
    o.Macs = temp.Macs
    o.Op = *temp.Op
    o.SiteId = temp.SiteId
    return nil
}

// tempOtherDeviceUpdateMulti is a temporary struct used for validating the fields of OtherDeviceUpdateMulti.
type tempOtherDeviceUpdateMulti  struct {
    Macs   []string                        `json:"macs,omitempty"`
    Op     *OtherDeviceUpdateOperationEnum `json:"op"`
    SiteId *uuid.UUID                      `json:"site_id,omitempty"`
}

func (o *tempOtherDeviceUpdateMulti) validate() error {
    var errs []string
    if o.Op == nil {
        errs = append(errs, "required field `op` is missing for type `other_device_update_multi`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
