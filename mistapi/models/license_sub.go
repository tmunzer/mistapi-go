package models

import (
	"encoding/json"
	"github.com/google/uuid"
)

// LicenseSub represents a LicenseSub struct.
type LicenseSub struct {
	CreatedTime *float64 `json:"created_time,omitempty"`
	// end date of the license term
	EndTime      *int       `json:"end_time,omitempty"`
	Id           *uuid.UUID `json:"id,omitempty"`
	ModifiedTime *float64   `json:"modified_time,omitempty"`
	OrderId      *string    `json:"order_id,omitempty"`
	OrgId        *uuid.UUID `json:"org_id,omitempty"`
	// number of devices entitled for this license
	Quantity *int `json:"quantity,omitempty"`
	// Number of licenses left in this subscription
	RemainingQuantity *int `json:"remaining_quantity,omitempty"`
	// start date of the license term
	StartTime      *int    `json:"start_time,omitempty"`
	SubscriptionId *string `json:"subscription_id,omitempty"`
	// enum: `SUB-AST`, `SUB-DATA`, `SUB-ENG`, `SUB-EX12`, `SUB-EX24`, `SUB-EX48`, `SUB-MAN`, `SUB-ME`, `SUB-PMA`, `SUB-SRX1`, `SUB-SRX2`, `SUB-SVNA`, `SUB-VNA`, `SUB-WAN1`, `SUB-WAN2`, `SUB-WVNA1`, `SUB-WVNA2`
	Type                 *LicenseTypeEnum `json:"type,omitempty"`
	AdditionalProperties map[string]any   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for LicenseSub.
// It customizes the JSON marshaling process for LicenseSub objects.
func (l LicenseSub) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(l.toMap())
}

// toMap converts the LicenseSub object to a map representation for JSON marshaling.
func (l LicenseSub) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, l.AdditionalProperties)
	if l.CreatedTime != nil {
		structMap["created_time"] = l.CreatedTime
	}
	if l.EndTime != nil {
		structMap["end_time"] = l.EndTime
	}
	if l.Id != nil {
		structMap["id"] = l.Id
	}
	if l.ModifiedTime != nil {
		structMap["modified_time"] = l.ModifiedTime
	}
	if l.OrderId != nil {
		structMap["order_id"] = l.OrderId
	}
	if l.OrgId != nil {
		structMap["org_id"] = l.OrgId
	}
	if l.Quantity != nil {
		structMap["quantity"] = l.Quantity
	}
	if l.RemainingQuantity != nil {
		structMap["remaining_quantity"] = l.RemainingQuantity
	}
	if l.StartTime != nil {
		structMap["start_time"] = l.StartTime
	}
	if l.SubscriptionId != nil {
		structMap["subscription_id"] = l.SubscriptionId
	}
	if l.Type != nil {
		structMap["type"] = l.Type
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for LicenseSub.
// It customizes the JSON unmarshaling process for LicenseSub objects.
func (l *LicenseSub) UnmarshalJSON(input []byte) error {
	var temp tempLicenseSub
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "created_time", "end_time", "id", "modified_time", "order_id", "org_id", "quantity", "remaining_quantity", "start_time", "subscription_id", "type")
	if err != nil {
		return err
	}

	l.AdditionalProperties = additionalProperties
	l.CreatedTime = temp.CreatedTime
	l.EndTime = temp.EndTime
	l.Id = temp.Id
	l.ModifiedTime = temp.ModifiedTime
	l.OrderId = temp.OrderId
	l.OrgId = temp.OrgId
	l.Quantity = temp.Quantity
	l.RemainingQuantity = temp.RemainingQuantity
	l.StartTime = temp.StartTime
	l.SubscriptionId = temp.SubscriptionId
	l.Type = temp.Type
	return nil
}

// tempLicenseSub is a temporary struct used for validating the fields of LicenseSub.
type tempLicenseSub struct {
	CreatedTime       *float64         `json:"created_time,omitempty"`
	EndTime           *int             `json:"end_time,omitempty"`
	Id                *uuid.UUID       `json:"id,omitempty"`
	ModifiedTime      *float64         `json:"modified_time,omitempty"`
	OrderId           *string          `json:"order_id,omitempty"`
	OrgId             *uuid.UUID       `json:"org_id,omitempty"`
	Quantity          *int             `json:"quantity,omitempty"`
	RemainingQuantity *int             `json:"remaining_quantity,omitempty"`
	StartTime         *int             `json:"start_time,omitempty"`
	SubscriptionId    *string          `json:"subscription_id,omitempty"`
	Type              *LicenseTypeEnum `json:"type,omitempty"`
}
