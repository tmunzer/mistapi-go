package models

import (
    "encoding/json"
)

// SnmpConfigClientList represents a SnmpConfigClientList struct.
type SnmpConfigClientList struct {
    ClientListName       *string        `json:"client_list_name,omitempty"`
    Clients              []string       `json:"clients,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SnmpConfigClientList.
// It customizes the JSON marshaling process for SnmpConfigClientList objects.
func (s SnmpConfigClientList) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SnmpConfigClientList object to a map representation for JSON marshaling.
func (s SnmpConfigClientList) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.ClientListName != nil {
        structMap["client_list_name"] = s.ClientListName
    }
    if s.Clients != nil {
        structMap["clients"] = s.Clients
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SnmpConfigClientList.
// It customizes the JSON unmarshaling process for SnmpConfigClientList objects.
func (s *SnmpConfigClientList) UnmarshalJSON(input []byte) error {
    var temp tempSnmpConfigClientList
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "client_list_name", "clients")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.ClientListName = temp.ClientListName
    s.Clients = temp.Clients
    return nil
}

// tempSnmpConfigClientList is a temporary struct used for validating the fields of SnmpConfigClientList.
type tempSnmpConfigClientList  struct {
    ClientListName *string  `json:"client_list_name,omitempty"`
    Clients        []string `json:"clients,omitempty"`
}
