package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/google/uuid"
    "strings"
)

// ResponseConfigHistorySearchItemWlan represents a ResponseConfigHistorySearchItemWlan struct.
type ResponseConfigHistorySearchItemWlan struct {
    Auth                 string                 `json:"auth"`
    Bands                []string               `json:"bands,omitempty"`
    // Unique ID of the object instance in the Mist Organnization
    Id                   uuid.UUID              `json:"id"`
    Ssid                 string                 `json:"ssid"`
    VlanIds              []string               `json:"vlan_ids,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseConfigHistorySearchItemWlan,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseConfigHistorySearchItemWlan) String() string {
    return fmt.Sprintf(
    	"ResponseConfigHistorySearchItemWlan[Auth=%v, Bands=%v, Id=%v, Ssid=%v, VlanIds=%v, AdditionalProperties=%v]",
    	r.Auth, r.Bands, r.Id, r.Ssid, r.VlanIds, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseConfigHistorySearchItemWlan.
// It customizes the JSON marshaling process for ResponseConfigHistorySearchItemWlan objects.
func (r ResponseConfigHistorySearchItemWlan) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "auth", "bands", "id", "ssid", "vlan_ids"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseConfigHistorySearchItemWlan object to a map representation for JSON marshaling.
func (r ResponseConfigHistorySearchItemWlan) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["auth"] = r.Auth
    if r.Bands != nil {
        structMap["bands"] = r.Bands
    }
    structMap["id"] = r.Id
    structMap["ssid"] = r.Ssid
    if r.VlanIds != nil {
        structMap["vlan_ids"] = r.VlanIds
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseConfigHistorySearchItemWlan.
// It customizes the JSON unmarshaling process for ResponseConfigHistorySearchItemWlan objects.
func (r *ResponseConfigHistorySearchItemWlan) UnmarshalJSON(input []byte) error {
    var temp tempResponseConfigHistorySearchItemWlan
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "auth", "bands", "id", "ssid", "vlan_ids")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.Auth = *temp.Auth
    r.Bands = temp.Bands
    r.Id = *temp.Id
    r.Ssid = *temp.Ssid
    r.VlanIds = temp.VlanIds
    return nil
}

// tempResponseConfigHistorySearchItemWlan is a temporary struct used for validating the fields of ResponseConfigHistorySearchItemWlan.
type tempResponseConfigHistorySearchItemWlan  struct {
    Auth    *string    `json:"auth"`
    Bands   []string   `json:"bands,omitempty"`
    Id      *uuid.UUID `json:"id"`
    Ssid    *string    `json:"ssid"`
    VlanIds []string   `json:"vlan_ids,omitempty"`
}

func (r *tempResponseConfigHistorySearchItemWlan) validate() error {
    var errs []string
    if r.Auth == nil {
        errs = append(errs, "required field `auth` is missing for type `response_config_history_search_item_wlan`")
    }
    if r.Id == nil {
        errs = append(errs, "required field `id` is missing for type `response_config_history_search_item_wlan`")
    }
    if r.Ssid == nil {
        errs = append(errs, "required field `ssid` is missing for type `response_config_history_search_item_wlan`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
