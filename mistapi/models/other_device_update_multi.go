package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// OtherDeviceUpdateMulti represents a OtherDeviceUpdateMulti struct.
type OtherDeviceUpdateMulti struct {
    // The mac address of the peer device.
    Macs                 []string                       `json:"macs,omitempty"`
    // The operation being performed. enum: `assign`, `unassign`
    Op                   OtherDeviceUpdateOperationEnum `json:"op"`
    SiteId               *uuid.UUID                     `json:"site_id,omitempty"`
    AdditionalProperties map[string]any                 `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for OtherDeviceUpdateMulti.
// It customizes the JSON marshaling process for OtherDeviceUpdateMulti objects.
func (o OtherDeviceUpdateMulti) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(o.toMap())
}

// toMap converts the OtherDeviceUpdateMulti object to a map representation for JSON marshaling.
func (o OtherDeviceUpdateMulti) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, o.AdditionalProperties)
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
    additionalProperties, err := UnmarshalAdditionalProperties(input, "macs", "op", "site_id")
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
    return errors.New(strings.Join(errs, "\n"))
}
