// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// SearchWxtagAppsItem represents a SearchWxtagAppsItem struct.
type SearchWxtagAppsItem struct {
    Group                string                 `json:"group"`
    Key                  string                 `json:"key"`
    Name                 string                 `json:"name"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SearchWxtagAppsItem,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SearchWxtagAppsItem) String() string {
    return fmt.Sprintf(
    	"SearchWxtagAppsItem[Group=%v, Key=%v, Name=%v, AdditionalProperties=%v]",
    	s.Group, s.Key, s.Name, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SearchWxtagAppsItem.
// It customizes the JSON marshaling process for SearchWxtagAppsItem objects.
func (s SearchWxtagAppsItem) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "group", "key", "name"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SearchWxtagAppsItem object to a map representation for JSON marshaling.
func (s SearchWxtagAppsItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["group"] = s.Group
    structMap["key"] = s.Key
    structMap["name"] = s.Name
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SearchWxtagAppsItem.
// It customizes the JSON unmarshaling process for SearchWxtagAppsItem objects.
func (s *SearchWxtagAppsItem) UnmarshalJSON(input []byte) error {
    var temp tempSearchWxtagAppsItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "group", "key", "name")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Group = *temp.Group
    s.Key = *temp.Key
    s.Name = *temp.Name
    return nil
}

// tempSearchWxtagAppsItem is a temporary struct used for validating the fields of SearchWxtagAppsItem.
type tempSearchWxtagAppsItem  struct {
    Group *string `json:"group"`
    Key   *string `json:"key"`
    Name  *string `json:"name"`
}

func (s *tempSearchWxtagAppsItem) validate() error {
    var errs []string
    if s.Group == nil {
        errs = append(errs, "required field `group` is missing for type `search_wxtag_apps_item`")
    }
    if s.Key == nil {
        errs = append(errs, "required field `key` is missing for type `search_wxtag_apps_item`")
    }
    if s.Name == nil {
        errs = append(errs, "required field `name` is missing for type `search_wxtag_apps_item`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
