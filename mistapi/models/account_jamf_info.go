package models

import (
	"encoding/json"
)

// AccountJamfInfo represents a AccountJamfInfo struct.
// OAuth linked Jamf apps account details
type AccountJamfInfo struct {
	// customer account client id
	ClientId *string `json:"client_id,omitempty"`
	// This error is provided when the Jamf account fails to fetch token/data
	Error *string `json:"error,omitempty"`
	// customer account Jamf instance URL
	InstanceUrl *string `json:"instance_url,omitempty"`
	// Is the last data pull for Jamf account is successful or not
	LastStatus *string `json:"last_status,omitempty"`
	// Last data pull timestamp, background jobs that pull Jamf account data
	LastSync *int64 `json:"last_sync,omitempty"`
	// First name of the user who linked the Jamf account
	LinkedBy *string `json:"linked_by,omitempty"`
	// Name of the company whose Jamf account mist has subscribed to
	Name *string `json:"name,omitempty"`
	// smart group membership for determining compliance status
	SmartgroupName       *string        `json:"smartgroup_name,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for AccountJamfInfo.
// It customizes the JSON marshaling process for AccountJamfInfo objects.
func (a AccountJamfInfo) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(a.toMap())
}

// toMap converts the AccountJamfInfo object to a map representation for JSON marshaling.
func (a AccountJamfInfo) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, a.AdditionalProperties)
	if a.ClientId != nil {
		structMap["client_id"] = a.ClientId
	}
	if a.Error != nil {
		structMap["error"] = a.Error
	}
	if a.InstanceUrl != nil {
		structMap["instance_url"] = a.InstanceUrl
	}
	if a.LastStatus != nil {
		structMap["last_status"] = a.LastStatus
	}
	if a.LastSync != nil {
		structMap["last_sync"] = a.LastSync
	}
	if a.LinkedBy != nil {
		structMap["linked_by"] = a.LinkedBy
	}
	if a.Name != nil {
		structMap["name"] = a.Name
	}
	if a.SmartgroupName != nil {
		structMap["smartgroup_name"] = a.SmartgroupName
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AccountJamfInfo.
// It customizes the JSON unmarshaling process for AccountJamfInfo objects.
func (a *AccountJamfInfo) UnmarshalJSON(input []byte) error {
	var temp tempAccountJamfInfo
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "client_id", "error", "instance_url", "last_status", "last_sync", "linked_by", "name", "smartgroup_name")
	if err != nil {
		return err
	}

	a.AdditionalProperties = additionalProperties
	a.ClientId = temp.ClientId
	a.Error = temp.Error
	a.InstanceUrl = temp.InstanceUrl
	a.LastStatus = temp.LastStatus
	a.LastSync = temp.LastSync
	a.LinkedBy = temp.LinkedBy
	a.Name = temp.Name
	a.SmartgroupName = temp.SmartgroupName
	return nil
}

// tempAccountJamfInfo is a temporary struct used for validating the fields of AccountJamfInfo.
type tempAccountJamfInfo struct {
	ClientId       *string `json:"client_id,omitempty"`
	Error          *string `json:"error,omitempty"`
	InstanceUrl    *string `json:"instance_url,omitempty"`
	LastStatus     *string `json:"last_status,omitempty"`
	LastSync       *int64  `json:"last_sync,omitempty"`
	LinkedBy       *string `json:"linked_by,omitempty"`
	Name           *string `json:"name,omitempty"`
	SmartgroupName *string `json:"smartgroup_name,omitempty"`
}
