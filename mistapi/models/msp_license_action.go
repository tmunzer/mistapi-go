package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/google/uuid"
    "strings"
)

// MspLicenseAction represents a MspLicenseAction struct.
type MspLicenseAction struct {
    // Required if `op`==`unamend`
    AmendmentId          *string                       `json:"amendment_id,omitempty"`
    // Required if `op`==`amend`, destination org id
    DstOrgId             *uuid.UUID                    `json:"dst_org_id,omitempty"`
    // Required if `op`==`annotate`
    Notes                *string                       `json:"notes,omitempty"`
    // enum: `amend`, `annotate`, `delete`, `unamend`
    Op                   MspLicenseActionOperationEnum `json:"op"`
    // Required if `op`==`amend`
    Quantity             *float64                      `json:"quantity,omitempty"`
    // Required if `op`==`annotate`
    SubscriptionId       *string                       `json:"subscription_id,omitempty"`
    AdditionalProperties map[string]interface{}        `json:"_"`
}

// String implements the fmt.Stringer interface for MspLicenseAction,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MspLicenseAction) String() string {
    return fmt.Sprintf(
    	"MspLicenseAction[AmendmentId=%v, DstOrgId=%v, Notes=%v, Op=%v, Quantity=%v, SubscriptionId=%v, AdditionalProperties=%v]",
    	m.AmendmentId, m.DstOrgId, m.Notes, m.Op, m.Quantity, m.SubscriptionId, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MspLicenseAction.
// It customizes the JSON marshaling process for MspLicenseAction objects.
func (m MspLicenseAction) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "amendment_id", "dst_org_id", "notes", "op", "quantity", "subscription_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the MspLicenseAction object to a map representation for JSON marshaling.
func (m MspLicenseAction) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, m.AdditionalProperties)
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "amendment_id", "dst_org_id", "notes", "op", "quantity", "subscription_id")
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
