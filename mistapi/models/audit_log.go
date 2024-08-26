package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// AuditLog represents a AuditLog struct.
type AuditLog struct {
    AdminId              uuid.UUID      `json:"admin_id"`
    AdminName            string         `json:"admin_name"`
    // field values after the change
    After                *interface{}   `json:"after,omitempty"`
    // field values prior to the change
    Before               *interface{}   `json:"before,omitempty"`
    ForSite              *bool          `json:"for_site,omitempty"`
    Id                   uuid.UUID      `json:"id"`
    Message              string         `json:"message"`
    OrgId                uuid.UUID      `json:"org_id"`
    SiteId               uuid.UUID      `json:"site_id"`
    Timestamp            float64        `json:"timestamp"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for AuditLog.
// It customizes the JSON marshaling process for AuditLog objects.
func (a AuditLog) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the AuditLog object to a map representation for JSON marshaling.
func (a AuditLog) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    structMap["admin_id"] = a.AdminId
    structMap["admin_name"] = a.AdminName
    if a.After != nil {
        structMap["after"] = a.After
    }
    if a.Before != nil {
        structMap["before"] = a.Before
    }
    if a.ForSite != nil {
        structMap["for_site"] = a.ForSite
    }
    structMap["id"] = a.Id
    structMap["message"] = a.Message
    structMap["org_id"] = a.OrgId
    structMap["site_id"] = a.SiteId
    structMap["timestamp"] = a.Timestamp
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AuditLog.
// It customizes the JSON unmarshaling process for AuditLog objects.
func (a *AuditLog) UnmarshalJSON(input []byte) error {
    var temp tempAuditLog
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "admin_id", "admin_name", "after", "before", "for_site", "id", "message", "org_id", "site_id", "timestamp")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.AdminId = *temp.AdminId
    a.AdminName = *temp.AdminName
    a.After = temp.After
    a.Before = temp.Before
    a.ForSite = temp.ForSite
    a.Id = *temp.Id
    a.Message = *temp.Message
    a.OrgId = *temp.OrgId
    a.SiteId = *temp.SiteId
    a.Timestamp = *temp.Timestamp
    return nil
}

// tempAuditLog is a temporary struct used for validating the fields of AuditLog.
type tempAuditLog  struct {
    AdminId   *uuid.UUID   `json:"admin_id"`
    AdminName *string      `json:"admin_name"`
    After     *interface{} `json:"after,omitempty"`
    Before    *interface{} `json:"before,omitempty"`
    ForSite   *bool        `json:"for_site,omitempty"`
    Id        *uuid.UUID   `json:"id"`
    Message   *string      `json:"message"`
    OrgId     *uuid.UUID   `json:"org_id"`
    SiteId    *uuid.UUID   `json:"site_id"`
    Timestamp *float64     `json:"timestamp"`
}

func (a *tempAuditLog) validate() error {
    var errs []string
    if a.AdminId == nil {
        errs = append(errs, "required field `admin_id` is missing for type `audit_log`")
    }
    if a.AdminName == nil {
        errs = append(errs, "required field `admin_name` is missing for type `audit_log`")
    }
    if a.Id == nil {
        errs = append(errs, "required field `id` is missing for type `audit_log`")
    }
    if a.Message == nil {
        errs = append(errs, "required field `message` is missing for type `audit_log`")
    }
    if a.OrgId == nil {
        errs = append(errs, "required field `org_id` is missing for type `audit_log`")
    }
    if a.SiteId == nil {
        errs = append(errs, "required field `site_id` is missing for type `audit_log`")
    }
    if a.Timestamp == nil {
        errs = append(errs, "required field `timestamp` is missing for type `audit_log`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
