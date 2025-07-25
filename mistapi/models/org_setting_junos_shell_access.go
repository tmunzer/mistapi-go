// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// OrgSettingJunosShellAccess represents a OrgSettingJunosShellAccess struct.
// by default, webshell access is only enabled for Admin user
type OrgSettingJunosShellAccess struct {
    // enum: `admin`, `viewer`, `none`
    Admin                *OrgSettingJunosShellAccessAdminEnum    `json:"admin,omitempty"`
    // enum: `admin`, `viewer`, `none`
    Helpdesk             *OrgSettingJunosShellAccessHelpdeskEnum `json:"helpdesk,omitempty"`
    // enum: `admin`, `viewer`, `none`
    Read                 *OrgSettingJunosShellAccessReadEnum     `json:"read,omitempty"`
    // enum: `admin`, `viewer`, `none`
    Write                *OrgSettingJunosShellAccessWriteEnum    `json:"write,omitempty"`
    AdditionalProperties map[string]interface{}                  `json:"_"`
}

// String implements the fmt.Stringer interface for OrgSettingJunosShellAccess,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OrgSettingJunosShellAccess) String() string {
    return fmt.Sprintf(
    	"OrgSettingJunosShellAccess[Admin=%v, Helpdesk=%v, Read=%v, Write=%v, AdditionalProperties=%v]",
    	o.Admin, o.Helpdesk, o.Read, o.Write, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OrgSettingJunosShellAccess.
// It customizes the JSON marshaling process for OrgSettingJunosShellAccess objects.
func (o OrgSettingJunosShellAccess) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(o.AdditionalProperties,
        "admin", "helpdesk", "read", "write"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(o.toMap())
}

// toMap converts the OrgSettingJunosShellAccess object to a map representation for JSON marshaling.
func (o OrgSettingJunosShellAccess) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, o.AdditionalProperties)
    if o.Admin != nil {
        structMap["admin"] = o.Admin
    }
    if o.Helpdesk != nil {
        structMap["helpdesk"] = o.Helpdesk
    }
    if o.Read != nil {
        structMap["read"] = o.Read
    }
    if o.Write != nil {
        structMap["write"] = o.Write
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSettingJunosShellAccess.
// It customizes the JSON unmarshaling process for OrgSettingJunosShellAccess objects.
func (o *OrgSettingJunosShellAccess) UnmarshalJSON(input []byte) error {
    var temp tempOrgSettingJunosShellAccess
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "admin", "helpdesk", "read", "write")
    if err != nil {
    	return err
    }
    o.AdditionalProperties = additionalProperties
    
    o.Admin = temp.Admin
    o.Helpdesk = temp.Helpdesk
    o.Read = temp.Read
    o.Write = temp.Write
    return nil
}

// tempOrgSettingJunosShellAccess is a temporary struct used for validating the fields of OrgSettingJunosShellAccess.
type tempOrgSettingJunosShellAccess  struct {
    Admin    *OrgSettingJunosShellAccessAdminEnum    `json:"admin,omitempty"`
    Helpdesk *OrgSettingJunosShellAccessHelpdeskEnum `json:"helpdesk,omitempty"`
    Read     *OrgSettingJunosShellAccessReadEnum     `json:"read,omitempty"`
    Write    *OrgSettingJunosShellAccessWriteEnum    `json:"write,omitempty"`
}
