package models

import (
    "encoding/json"
)

// AccountVmwareInfo represents a AccountVmwareInfo struct.
type AccountVmwareInfo struct {
    AccountId            *string        `json:"account_id,omitempty"`
    // Linked VMware Instance URL
    InstanceUrl          *string        `json:"instance_url,omitempty"`
    // Is the last data pull for VMware account is successful or not
    LastStatus           *string        `json:"last_status,omitempty"`
    // Last data pull timestamp, background jobs that pull VMware account data
    LastSync             *int           `json:"last_sync,omitempty"`
    // First name of the user who linked the VMware account
    LinkedBy             *string        `json:"linked_by,omitempty"`
    // This error is provided when the VMware account fails to fetch token/data
    LinkedTimestamp      *int           `json:"linked_timestamp,omitempty"`
    // Name of the company whose VMware account mist has subscribed to
    Name                 *string        `json:"name,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for AccountVmwareInfo.
// It customizes the JSON marshaling process for AccountVmwareInfo objects.
func (a AccountVmwareInfo) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the AccountVmwareInfo object to a map representation for JSON marshaling.
func (a AccountVmwareInfo) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    if a.AccountId != nil {
        structMap["account_id"] = a.AccountId
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

// UnmarshalJSON implements the json.Unmarshaler interface for AccountVmwareInfo.
// It customizes the JSON unmarshaling process for AccountVmwareInfo objects.
func (a *AccountVmwareInfo) UnmarshalJSON(input []byte) error {
    var temp accountVmwareInfo
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "account_id", "instance_url", "last_status", "last_sync", "linked_by", "linked_timestamp", "name")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.AccountId = temp.AccountId
    a.InstanceUrl = temp.InstanceUrl
    a.LastStatus = temp.LastStatus
    a.LastSync = temp.LastSync
    a.LinkedBy = temp.LinkedBy
    a.LinkedTimestamp = temp.LinkedTimestamp
    a.Name = temp.Name
    return nil
}

// accountVmwareInfo is a temporary struct used for validating the fields of AccountVmwareInfo.
type accountVmwareInfo  struct {
    AccountId       *string `json:"account_id,omitempty"`
    InstanceUrl     *string `json:"instance_url,omitempty"`
    LastStatus      *string `json:"last_status,omitempty"`
    LastSync        *int    `json:"last_sync,omitempty"`
    LinkedBy        *string `json:"linked_by,omitempty"`
    LinkedTimestamp *int    `json:"linked_timestamp,omitempty"`
    Name            *string `json:"name,omitempty"`
}
