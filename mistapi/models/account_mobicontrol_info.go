package models

import (
	"encoding/json"
)

// AccountMobicontrolInfo represents a AccountMobicontrolInfo struct.
// OAuth linked Mobicontrol apps account details
type AccountMobicontrolInfo struct {
	// customer account client id
	AccountId *string `json:"account_id,omitempty"`
	// Linked MobiControl Client Id
	ClientId *string `json:"client_id,omitempty"`
	// This error is provided when the MobiControl account fails to fetch token/data
	Error *string `json:"error,omitempty"`
	// Linked MobiControl Instance URL
	InstanceUrl *string `json:"instance_url,omitempty"`
	// Is the last data pull for MobiControl account is successful or not
	LastStatus *string `json:"last_status,omitempty"`
	// Last data pull timestamp, background jobs that pull MobiControl account data
	LastSync *int64 `json:"last_sync,omitempty"`
	// First name of the user who linked the MobiControl account
	LinkedBy        *string `json:"linked_by,omitempty"`
	LinkedTimestamp *int64  `json:"linked_timestamp,omitempty"`
	// Name of the company whose MobiControl account mist has subscribed to
	Name                 *string        `json:"name,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for AccountMobicontrolInfo.
// It customizes the JSON marshaling process for AccountMobicontrolInfo objects.
func (a AccountMobicontrolInfo) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(a.toMap())
}

// toMap converts the AccountMobicontrolInfo object to a map representation for JSON marshaling.
func (a AccountMobicontrolInfo) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, a.AdditionalProperties)
	if a.AccountId != nil {
		structMap["account_id"] = a.AccountId
	}
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
	if a.LinkedTimestamp != nil {
		structMap["linked_timestamp"] = a.LinkedTimestamp
	}
	if a.Name != nil {
		structMap["name"] = a.Name
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AccountMobicontrolInfo.
// It customizes the JSON unmarshaling process for AccountMobicontrolInfo objects.
func (a *AccountMobicontrolInfo) UnmarshalJSON(input []byte) error {
	var temp tempAccountMobicontrolInfo
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "account_id", "client_id", "error", "instance_url", "last_status", "last_sync", "linked_by", "linked_timestamp", "name")
	if err != nil {
		return err
	}

	a.AdditionalProperties = additionalProperties
	a.AccountId = temp.AccountId
	a.ClientId = temp.ClientId
	a.Error = temp.Error
	a.InstanceUrl = temp.InstanceUrl
	a.LastStatus = temp.LastStatus
	a.LastSync = temp.LastSync
	a.LinkedBy = temp.LinkedBy
	a.LinkedTimestamp = temp.LinkedTimestamp
	a.Name = temp.Name
	return nil
}

// tempAccountMobicontrolInfo is a temporary struct used for validating the fields of AccountMobicontrolInfo.
type tempAccountMobicontrolInfo struct {
	AccountId       *string `json:"account_id,omitempty"`
	ClientId        *string `json:"client_id,omitempty"`
	Error           *string `json:"error,omitempty"`
	InstanceUrl     *string `json:"instance_url,omitempty"`
	LastStatus      *string `json:"last_status,omitempty"`
	LastSync        *int64  `json:"last_sync,omitempty"`
	LinkedBy        *string `json:"linked_by,omitempty"`
	LinkedTimestamp *int64  `json:"linked_timestamp,omitempty"`
	Name            *string `json:"name,omitempty"`
}
