package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ModuleStatErrorsItems represents a ModuleStatErrorsItems struct.
type ModuleStatErrorsItems struct {
    Feature              *string        `json:"feature,omitempty"`
    MinimumVersion       *string        `json:"minimum_version,omitempty"`
    Reason               *string        `json:"reason,omitempty"`
    Since                int            `json:"since"`
    Type                 string         `json:"type"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ModuleStatErrorsItems.
// It customizes the JSON marshaling process for ModuleStatErrorsItems objects.
func (m ModuleStatErrorsItems) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the ModuleStatErrorsItems object to a map representation for JSON marshaling.
func (m ModuleStatErrorsItems) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, m.AdditionalProperties)
    if m.Feature != nil {
        structMap["feature"] = m.Feature
    }
    if m.MinimumVersion != nil {
        structMap["minimum_version"] = m.MinimumVersion
    }
    if m.Reason != nil {
        structMap["reason"] = m.Reason
    }
    structMap["since"] = m.Since
    structMap["type"] = m.Type
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ModuleStatErrorsItems.
// It customizes the JSON unmarshaling process for ModuleStatErrorsItems objects.
func (m *ModuleStatErrorsItems) UnmarshalJSON(input []byte) error {
    var temp moduleStatErrorsItems
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "feature", "minimum_version", "reason", "since", "type")
    if err != nil {
    	return err
    }
    
    m.AdditionalProperties = additionalProperties
    m.Feature = temp.Feature
    m.MinimumVersion = temp.MinimumVersion
    m.Reason = temp.Reason
    m.Since = *temp.Since
    m.Type = *temp.Type
    return nil
}

// moduleStatErrorsItems is a temporary struct used for validating the fields of ModuleStatErrorsItems.
type moduleStatErrorsItems  struct {
    Feature        *string `json:"feature,omitempty"`
    MinimumVersion *string `json:"minimum_version,omitempty"`
    Reason         *string `json:"reason,omitempty"`
    Since          *int    `json:"since"`
    Type           *string `json:"type"`
}

func (m *moduleStatErrorsItems) validate() error {
    var errs []string
    if m.Since == nil {
        errs = append(errs, "required field `since` is missing for type `Module_Stat_Errors_Items`")
    }
    if m.Type == nil {
        errs = append(errs, "required field `type` is missing for type `Module_Stat_Errors_Items`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
