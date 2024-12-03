package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// Org represents a Org struct.
type Org struct {
    AlarmtemplateId      Optional[uuid.UUID]    `json:"alarmtemplate_id"`
    AllowMist            *bool                  `json:"allow_mist,omitempty"`
    // when the object has been created, in epoch
    CreatedTime          *float64               `json:"created_time,omitempty"`
    // Unique ID of the object instance in the Mist Organnization
    Id                   *uuid.UUID             `json:"id,omitempty"`
    // when the object has been modified for the last time, in epoch
    ModifiedTime         *float64               `json:"modified_time,omitempty"`
    MspId                *uuid.UUID             `json:"msp_id,omitempty"`
    // logo uploaded by the MSP with advanced tier, only present if provided
    MspLogoUrl           *string                `json:"msp_logo_url,omitempty"`
    // name of the msp the org belongs to
    MspName              *string                `json:"msp_name,omitempty"`
    Name                 string                 `json:"name"`
    OrggroupIds          []uuid.UUID            `json:"orggroup_ids,omitempty"`
    SessionExpiry        *int                   `json:"session_expiry,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for Org.
// It customizes the JSON marshaling process for Org objects.
func (o Org) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(o.AdditionalProperties,
        "alarmtemplate_id", "allow_mist", "created_time", "id", "modified_time", "msp_id", "msp_logo_url", "msp_name", "name", "orggroup_ids", "session_expiry"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(o.toMap())
}

// toMap converts the Org object to a map representation for JSON marshaling.
func (o Org) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, o.AdditionalProperties)
    if o.AlarmtemplateId.IsValueSet() {
        if o.AlarmtemplateId.Value() != nil {
            structMap["alarmtemplate_id"] = o.AlarmtemplateId.Value()
        } else {
            structMap["alarmtemplate_id"] = nil
        }
    }
    if o.AllowMist != nil {
        structMap["allow_mist"] = o.AllowMist
    }
    if o.CreatedTime != nil {
        structMap["created_time"] = o.CreatedTime
    }
    if o.Id != nil {
        structMap["id"] = o.Id
    }
    if o.ModifiedTime != nil {
        structMap["modified_time"] = o.ModifiedTime
    }
    if o.MspId != nil {
        structMap["msp_id"] = o.MspId
    }
    if o.MspLogoUrl != nil {
        structMap["msp_logo_url"] = o.MspLogoUrl
    }
    if o.MspName != nil {
        structMap["msp_name"] = o.MspName
    }
    structMap["name"] = o.Name
    if o.OrggroupIds != nil {
        structMap["orggroup_ids"] = o.OrggroupIds
    }
    if o.SessionExpiry != nil {
        structMap["session_expiry"] = o.SessionExpiry
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Org.
// It customizes the JSON unmarshaling process for Org objects.
func (o *Org) UnmarshalJSON(input []byte) error {
    var temp tempOrg
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "alarmtemplate_id", "allow_mist", "created_time", "id", "modified_time", "msp_id", "msp_logo_url", "msp_name", "name", "orggroup_ids", "session_expiry")
    if err != nil {
    	return err
    }
    o.AdditionalProperties = additionalProperties
    
    o.AlarmtemplateId = temp.AlarmtemplateId
    o.AllowMist = temp.AllowMist
    o.CreatedTime = temp.CreatedTime
    o.Id = temp.Id
    o.ModifiedTime = temp.ModifiedTime
    o.MspId = temp.MspId
    o.MspLogoUrl = temp.MspLogoUrl
    o.MspName = temp.MspName
    o.Name = *temp.Name
    o.OrggroupIds = temp.OrggroupIds
    o.SessionExpiry = temp.SessionExpiry
    return nil
}

// tempOrg is a temporary struct used for validating the fields of Org.
type tempOrg  struct {
    AlarmtemplateId Optional[uuid.UUID] `json:"alarmtemplate_id"`
    AllowMist       *bool               `json:"allow_mist,omitempty"`
    CreatedTime     *float64            `json:"created_time,omitempty"`
    Id              *uuid.UUID          `json:"id,omitempty"`
    ModifiedTime    *float64            `json:"modified_time,omitempty"`
    MspId           *uuid.UUID          `json:"msp_id,omitempty"`
    MspLogoUrl      *string             `json:"msp_logo_url,omitempty"`
    MspName         *string             `json:"msp_name,omitempty"`
    Name            *string             `json:"name"`
    OrggroupIds     []uuid.UUID         `json:"orggroup_ids,omitempty"`
    SessionExpiry   *int                `json:"session_expiry,omitempty"`
}

func (o *tempOrg) validate() error {
    var errs []string
    if o.Name == nil {
        errs = append(errs, "required field `name` is missing for type `org`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
