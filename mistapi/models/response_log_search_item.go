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

// ResponseLogSearchItem represents a ResponseLogSearchItem struct.
type ResponseLogSearchItem struct {
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

// String implements the fmt.Stringer interface for ResponseLogSearchItem,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseLogSearchItem) String() string {
	return fmt.Sprintf(
		"ResponseLogSearchItem[AdminId=%v, AdminName=%v, After=%v, Before=%v, ForSite=%v, Id=%v, Message=%v, OrgId=%v, SiteId=%v, SrcIp=%v, Timestamp=%v, AdditionalProperties=%v]",
		r.AdminId, r.AdminName, r.After, r.Before, r.ForSite, r.Id, r.Message, r.OrgId, r.SiteId, r.SrcIp, r.Timestamp, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseLogSearchItem.
// It customizes the JSON marshaling process for ResponseLogSearchItem objects.
func (r ResponseLogSearchItem) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"admin_id", "admin_name", "after", "before", "for_site", "id", "message", "org_id", "site_id", "src_ip", "timestamp"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseLogSearchItem object to a map representation for JSON marshaling.
func (r ResponseLogSearchItem) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, r.AdditionalProperties)
	if r.AdminId.IsValueSet() {
		if r.AdminId.Value() != nil {
			structMap["admin_id"] = r.AdminId.Value()
		} else {
			structMap["admin_id"] = nil
		}
	}
	if r.AdminName.IsValueSet() {
		if r.AdminName.Value() != nil {
			structMap["admin_name"] = r.AdminName.Value()
		} else {
			structMap["admin_name"] = nil
		}
	}
	if r.After != nil {
		structMap["after"] = r.After
	}
	if r.Before != nil {
		structMap["before"] = r.Before
	}
	if r.ForSite != nil {
		structMap["for_site"] = r.ForSite
	}
	if r.Id != nil {
		structMap["id"] = r.Id
	}
	structMap["message"] = r.Message
	structMap["org_id"] = r.OrgId
	if r.SiteId.IsValueSet() {
		if r.SiteId.Value() != nil {
			structMap["site_id"] = r.SiteId.Value()
		} else {
			structMap["site_id"] = nil
		}
	}
	if r.SrcIp != nil {
		structMap["src_ip"] = r.SrcIp
	}
	structMap["timestamp"] = r.Timestamp
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseLogSearchItem.
// It customizes the JSON unmarshaling process for ResponseLogSearchItem objects.
func (r *ResponseLogSearchItem) UnmarshalJSON(input []byte) error {
	var temp tempResponseLogSearchItem
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
	r.AdditionalProperties = additionalProperties

	r.AdminId = temp.AdminId
	r.AdminName = temp.AdminName
	r.After = temp.After
	r.Before = temp.Before
	r.ForSite = temp.ForSite
	r.Id = temp.Id
	r.Message = *temp.Message
	r.OrgId = *temp.OrgId
	r.SiteId = temp.SiteId
	r.SrcIp = temp.SrcIp
	r.Timestamp = *temp.Timestamp
	return nil
}

// tempResponseLogSearchItem is a temporary struct used for validating the fields of ResponseLogSearchItem.
type tempResponseLogSearchItem struct {
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

func (r *tempResponseLogSearchItem) validate() error {
	var errs []string
	if r.Message == nil {
		errs = append(errs, "required field `message` is missing for type `response_log_search_item`")
	}
	if r.OrgId == nil {
		errs = append(errs, "required field `org_id` is missing for type `response_log_search_item`")
	}
	if r.Timestamp == nil {
		errs = append(errs, "required field `timestamp` is missing for type `response_log_search_item`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
