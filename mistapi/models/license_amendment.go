package models

import (
    "encoding/json"
)

// LicenseAmendment represents a LicenseAmendment struct.
type LicenseAmendment struct {
    CreatedTime          *float64         `json:"created_time,omitempty"`
    EndTime              *int             `json:"end_time,omitempty"`
    Id                   *string          `json:"id,omitempty"`
    ModifiedTime         *float64         `json:"modified_time,omitempty"`
    Quantity             *int             `json:"quantity,omitempty"`
    StartTime            *int             `json:"start_time,omitempty"`
    SubscriptionId       *string          `json:"subscription_id,omitempty"`
    Type                 *LicenseTypeEnum `json:"type,omitempty"`
    AdditionalProperties map[string]any   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for LicenseAmendment.
// It customizes the JSON marshaling process for LicenseAmendment objects.
func (l LicenseAmendment) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(l.toMap())
}

// toMap converts the LicenseAmendment object to a map representation for JSON marshaling.
func (l LicenseAmendment) toMap() map[string]any {
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
    var temp licenseAmendment
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "created_time", "end_time", "id", "modified_time", "quantity", "start_time", "subscription_id", "type")
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

// licenseAmendment is a temporary struct used for validating the fields of LicenseAmendment.
type licenseAmendment  struct {
    CreatedTime    *float64         `json:"created_time,omitempty"`
    EndTime        *int             `json:"end_time,omitempty"`
    Id             *string          `json:"id,omitempty"`
    ModifiedTime   *float64         `json:"modified_time,omitempty"`
    Quantity       *int             `json:"quantity,omitempty"`
    StartTime      *int             `json:"start_time,omitempty"`
    SubscriptionId *string          `json:"subscription_id,omitempty"`
    Type           *LicenseTypeEnum `json:"type,omitempty"`
}
