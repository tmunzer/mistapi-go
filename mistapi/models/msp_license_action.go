package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// MspLicenseAction represents a MspLicenseAction struct.
type MspLicenseAction struct {
    // required if `op`==`unamend`
    AmendmentId          *string                       `json:"amendment_id,omitempty"`
    // required if `op`==`amend`, destination org id
    DstOrgId             *uuid.UUID                    `json:"dst_org_id,omitempty"`
    // required if `op`== `annotate`
    Notes                *string                       `json:"notes,omitempty"`
    // enum: `amend`, `annotate`, `delete`, `unamend`
    Op                   MspLicenseActionOperationEnum `json:"op"`
    // required if `op`==`amend`
    Quantity             *float64                      `json:"quantity,omitempty"`
    // required if `op`== `annotate`
    SubscriptionId       *string                       `json:"subscription_id,omitempty"`
    AdditionalProperties map[string]any                `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MspLicenseAction.
// It customizes the JSON marshaling process for MspLicenseAction objects.
func (m MspLicenseAction) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the MspLicenseAction object to a map representation for JSON marshaling.
func (m MspLicenseAction) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, m.AdditionalProperties)
    if m.AmendmentId != nil {
        structMap["amendment_id"] = m.AmendmentId
    }
    if m.DstOrgId != nil {
        structMap["dst_org_id"] = m.DstOrgId
    }
    if m.Notes != nil {
        structMap["notes"] = m.Notes
    }
    structMap["op"] = m.Op
    if m.Quantity != nil {
        structMap["quantity"] = m.Quantity
    }
    if m.SubscriptionId != nil {
        structMap["subscription_id"] = m.SubscriptionId
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MspLicenseAction.
// It customizes the JSON unmarshaling process for MspLicenseAction objects.
func (m *MspLicenseAction) UnmarshalJSON(input []byte) error {
    var temp tempMspLicenseAction
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "amendment_id", "dst_org_id", "notes", "op", "quantity", "subscription_id")
    if err != nil {
    	return err
    }
    
    m.AdditionalProperties = additionalProperties
    m.AmendmentId = temp.AmendmentId
    m.DstOrgId = temp.DstOrgId
    m.Notes = temp.Notes
    m.Op = *temp.Op
    m.Quantity = temp.Quantity
    m.SubscriptionId = temp.SubscriptionId
    return nil
}

// tempMspLicenseAction is a temporary struct used for validating the fields of MspLicenseAction.
type tempMspLicenseAction  struct {
    AmendmentId    *string                        `json:"amendment_id,omitempty"`
    DstOrgId       *uuid.UUID                     `json:"dst_org_id,omitempty"`
    Notes          *string                        `json:"notes,omitempty"`
    Op             *MspLicenseActionOperationEnum `json:"op"`
    Quantity       *float64                       `json:"quantity,omitempty"`
    SubscriptionId *string                        `json:"subscription_id,omitempty"`
}

func (m *tempMspLicenseAction) validate() error {
    var errs []string
    if m.Op == nil {
        errs = append(errs, "required field `op` is missing for type `msp_license_action`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
