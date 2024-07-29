package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// OrgLicenseAction represents a OrgLicenseAction struct.
type OrgLicenseAction struct {
    // if `op`==`unamend`, the ID of the operation to cancel
    AmendmentId          *uuid.UUID                    `json:"amendment_id,omitempty"`
    // if `op`==`amend`, the id of the org where the license is moved
    DstOrgId             *uuid.UUID                    `json:"dst_org_id,omitempty"`
    // if `op`==`annotate`
    Notes                *string                       `json:"notes,omitempty"`
    // to move a license, use the `amend` operation. enum: `amend`, `annotate`, `delete`, `unamend`
    Op                   OrgLicenseActionOperationEnum `json:"op"`
    // if `op`==`amend`, the number of licenses to move
    Quantity             *int                          `json:"quantity,omitempty"`
    // if `op`==`amend` or `op`==`delete`, the ID of the subscription to use
    SubscriptionId       *string                       `json:"subscription_id,omitempty"`
    AdditionalProperties map[string]any                `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for OrgLicenseAction.
// It customizes the JSON marshaling process for OrgLicenseAction objects.
func (o OrgLicenseAction) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(o.toMap())
}

// toMap converts the OrgLicenseAction object to a map representation for JSON marshaling.
func (o OrgLicenseAction) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, o.AdditionalProperties)
    if o.AmendmentId != nil {
        structMap["amendment_id"] = o.AmendmentId
    }
    if o.DstOrgId != nil {
        structMap["dst_org_id"] = o.DstOrgId
    }
    if o.Notes != nil {
        structMap["notes"] = o.Notes
    }
    structMap["op"] = o.Op
    if o.Quantity != nil {
        structMap["quantity"] = o.Quantity
    }
    if o.SubscriptionId != nil {
        structMap["subscription_id"] = o.SubscriptionId
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgLicenseAction.
// It customizes the JSON unmarshaling process for OrgLicenseAction objects.
func (o *OrgLicenseAction) UnmarshalJSON(input []byte) error {
    var temp orgLicenseAction
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
    
    o.AdditionalProperties = additionalProperties
    o.AmendmentId = temp.AmendmentId
    o.DstOrgId = temp.DstOrgId
    o.Notes = temp.Notes
    o.Op = *temp.Op
    o.Quantity = temp.Quantity
    o.SubscriptionId = temp.SubscriptionId
    return nil
}

// orgLicenseAction is a temporary struct used for validating the fields of OrgLicenseAction.
type orgLicenseAction  struct {
    AmendmentId    *uuid.UUID                     `json:"amendment_id,omitempty"`
    DstOrgId       *uuid.UUID                     `json:"dst_org_id,omitempty"`
    Notes          *string                        `json:"notes,omitempty"`
    Op             *OrgLicenseActionOperationEnum `json:"op"`
    Quantity       *int                           `json:"quantity,omitempty"`
    SubscriptionId *string                        `json:"subscription_id,omitempty"`
}

func (o *orgLicenseAction) validate() error {
    var errs []string
    if o.Op == nil {
        errs = append(errs, "required field `op` is missing for type `Org_License_Action`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
