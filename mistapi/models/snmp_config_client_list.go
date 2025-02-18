package models

import (
    "encoding/json"
    "fmt"
)

// SnmpConfigClientList represents a SnmpConfigClientList struct.
type SnmpConfigClientList struct {
    ClientListName       *string                `json:"client_list_name,omitempty"`
    Clients              []string               `json:"clients,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SnmpConfigClientList,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SnmpConfigClientList) String() string {
    return fmt.Sprintf(
    	"SnmpConfigClientList[ClientListName=%v, Clients=%v, AdditionalProperties=%v]",
    	s.ClientListName, s.Clients, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SnmpConfigClientList.
// It customizes the JSON marshaling process for SnmpConfigClientList objects.
func (s SnmpConfigClientList) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "client_list_name", "clients"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SnmpConfigClientList object to a map representation for JSON marshaling.
func (s SnmpConfigClientList) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "client_list_name", "clients")
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
