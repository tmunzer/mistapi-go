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

// LogEvent represents a LogEvent struct.
type LogEvent struct {
	// admin id
	AdminId Optional[uuid.UUID] `json:"admin_id"`
	// Name of the admin that performs the action
	AdminName Optional[string] `json:"admin_name"`
	// field values after the change
	After *interface{} `json:"after,omitempty"`
	// field values prior to the change
	Before  *interface{} `json:"before,omitempty"`
	ForSite *bool        `json:"for_site,omitempty"`
	// Unique ID of the object instance in the Mist Organization
	Id *uuid.UUID `json:"id,omitempty"`
	// log message
	Message string              `json:"message"`
	OrgId   uuid.UUID           `json:"org_id"`
	SiteId  Optional[uuid.UUID] `json:"site_id"`
	// sender source ip address
	SrcIp *string `json:"src_ip,omitempty"`
	// Epoch (seconds)
	Timestamp            float64                `json:"timestamp"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for LogEvent,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (l LogEvent) String() string {
	return fmt.Sprintf(
		"LogEvent[AdminId=%v, AdminName=%v, After=%v, Before=%v, ForSite=%v, Id=%v, Message=%v, OrgId=%v, SiteId=%v, SrcIp=%v, Timestamp=%v, AdditionalProperties=%v]",
		l.AdminId, l.AdminName, l.After, l.Before, l.ForSite, l.Id, l.Message, l.OrgId, l.SiteId, l.SrcIp, l.Timestamp, l.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for LogEvent.
// It customizes the JSON marshaling process for LogEvent objects.
func (l LogEvent) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(l.AdditionalProperties,
		"admin_id", "admin_name", "after", "before", "for_site", "id", "message", "org_id", "site_id", "src_ip", "timestamp"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(l.toMap())
}

// toMap converts the LogEvent object to a map representation for JSON marshaling.
func (l LogEvent) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, l.AdditionalProperties)
	if l.AdminId.IsValueSet() {
		if l.AdminId.Value() != nil {
			structMap["admin_id"] = l.AdminId.Value()
		} else {
			structMap["admin_id"] = nil
		}
	}
	if l.AdminName.IsValueSet() {
		if l.AdminName.Value() != nil {
			structMap["admin_name"] = l.AdminName.Value()
		} else {
			structMap["admin_name"] = nil
		}
	}
	if l.After != nil {
		structMap["after"] = l.After
	}
	if l.Before != nil {
		structMap["before"] = l.Before
	}
	if l.ForSite != nil {
		structMap["for_site"] = l.ForSite
	}
	if l.Id != nil {
		structMap["id"] = l.Id
	}
	structMap["message"] = l.Message
	structMap["org_id"] = l.OrgId
	if l.SiteId.IsValueSet() {
		if l.SiteId.Value() != nil {
			structMap["site_id"] = l.SiteId.Value()
		} else {
			structMap["site_id"] = nil
		}
	}
	if l.SrcIp != nil {
		structMap["src_ip"] = l.SrcIp
	}
	structMap["timestamp"] = l.Timestamp
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for LogEvent.
// It customizes the JSON unmarshaling process for LogEvent objects.
func (l *LogEvent) UnmarshalJSON(input []byte) error {
	var temp tempLogEvent
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "admin_id", "admin_name", "after", "before", "for_site", "id", "message", "org_id", "site_id", "src_ip", "timestamp")
	if err != nil {
		return err
	}
	l.AdditionalProperties = additionalProperties

	l.AdminId = temp.AdminId
	l.AdminName = temp.AdminName
	l.After = temp.After
	l.Before = temp.Before
	l.ForSite = temp.ForSite
	l.Id = temp.Id
	l.Message = *temp.Message
	l.OrgId = *temp.OrgId
	l.SiteId = temp.SiteId
	l.SrcIp = temp.SrcIp
	l.Timestamp = *temp.Timestamp
	return nil
}

// tempLogEvent is a temporary struct used for validating the fields of LogEvent.
type tempLogEvent struct {
	AdminId   Optional[uuid.UUID] `json:"admin_id"`
	AdminName Optional[string]    `json:"admin_name"`
	After     *interface{}        `json:"after,omitempty"`
	Before    *interface{}        `json:"before,omitempty"`
	ForSite   *bool               `json:"for_site,omitempty"`
	Id        *uuid.UUID          `json:"id,omitempty"`
	Message   *string             `json:"message"`
	OrgId     *uuid.UUID          `json:"org_id"`
	SiteId    Optional[uuid.UUID] `json:"site_id"`
	SrcIp     *string             `json:"src_ip,omitempty"`
	Timestamp *float64            `json:"timestamp"`
}

func (l *tempLogEvent) validate() error {
	var errs []string
	if l.Message == nil {
		errs = append(errs, "required field `message` is missing for type `log_event`")
	}
	if l.OrgId == nil {
		errs = append(errs, "required field `org_id` is missing for type `log_event`")
	}
	if l.Timestamp == nil {
		errs = append(errs, "required field `timestamp` is missing for type `log_event`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
