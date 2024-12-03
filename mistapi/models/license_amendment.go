package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// LicenseAmendment represents a LicenseAmendment struct.
type LicenseAmendment struct {
    // when the object has been created, in epoch
    CreatedTime          *float64               `json:"created_time,omitempty"`
    EndTime              *int                   `json:"end_time,omitempty"`
    // Unique ID of the object instance in the Mist Organnization
    Id                   *uuid.UUID             `json:"id,omitempty"`
    // when the object has been modified for the last time, in epoch
    ModifiedTime         *float64               `json:"modified_time,omitempty"`
    Quantity             *int                   `json:"quantity,omitempty"`
    StartTime            *int                   `json:"start_time,omitempty"`
    SubscriptionId       *string                `json:"subscription_id,omitempty"`
    // enum: `SUB-AST`, `SUB-DATA`, `SUB-ENG`, `SUB-EX12`, `SUB-EX24`, `SUB-EX48`, `SUB-MAN`, `SUB-ME`, `SUB-PMA`, `SUB-SRX1`, `SUB-SRX2`, `SUB-SVNA`, `SUB-VNA`, `SUB-WAN1`, `SUB-WAN2`, `SUB-WVNA1`, `SUB-WVNA2`
    Type                 *LicenseTypeEnum       `json:"type,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for LicenseAmendment.
// It customizes the JSON marshaling process for LicenseAmendment objects.
func (l LicenseAmendment) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(l.AdditionalProperties,
        "created_time", "end_time", "id", "modified_time", "quantity", "start_time", "subscription_id", "type"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(l.toMap())
}

// toMap converts the LicenseAmendment object to a map representation for JSON marshaling.
func (l LicenseAmendment) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, l.AdditionalProperties)
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
    if l.Quantity != nil {
        structMap["quantity"] = l.Quantity
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

// UnmarshalJSON implements the json.Unmarshaler interface for LicenseAmendment.
// It customizes the JSON unmarshaling process for LicenseAmendment objects.
func (l *LicenseAmendment) UnmarshalJSON(input []byte) error {
    var temp tempLicenseAmendment
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "created_time", "end_time", "id", "modified_time", "quantity", "start_time", "subscription_id", "type")
    if err != nil {
    	return err
    }
    l.AdditionalProperties = additionalProperties
    
    l.CreatedTime = temp.CreatedTime
    l.EndTime = temp.EndTime
    l.Id = temp.Id
    l.ModifiedTime = temp.ModifiedTime
    l.Quantity = temp.Quantity
    l.StartTime = temp.StartTime
    l.SubscriptionId = temp.SubscriptionId
    l.Type = temp.Type
    return nil
}

// tempLicenseAmendment is a temporary struct used for validating the fields of LicenseAmendment.
type tempLicenseAmendment  struct {
    CreatedTime    *float64         `json:"created_time,omitempty"`
    EndTime        *int             `json:"end_time,omitempty"`
    Id             *uuid.UUID       `json:"id,omitempty"`
    ModifiedTime   *float64         `json:"modified_time,omitempty"`
    Quantity       *int             `json:"quantity,omitempty"`
    StartTime      *int             `json:"start_time,omitempty"`
    SubscriptionId *string          `json:"subscription_id,omitempty"`
    Type           *LicenseTypeEnum `json:"type,omitempty"`
}
