package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// ModuleStatItemErrorsItems represents a ModuleStatItemErrorsItems struct.
type ModuleStatItemErrorsItems struct {
    Feature              *string                `json:"feature,omitempty"`
    MinimumVersion       *string                `json:"minimum_version,omitempty"`
    Reason               *string                `json:"reason,omitempty"`
    Since                int                    `json:"since"`
    Type                 string                 `json:"type"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ModuleStatItemErrorsItems,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m ModuleStatItemErrorsItems) String() string {
    return fmt.Sprintf(
    	"ModuleStatItemErrorsItems[Feature=%v, MinimumVersion=%v, Reason=%v, Since=%v, Type=%v, AdditionalProperties=%v]",
    	m.Feature, m.MinimumVersion, m.Reason, m.Since, m.Type, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ModuleStatItemErrorsItems.
// It customizes the JSON marshaling process for ModuleStatItemErrorsItems objects.
func (m ModuleStatItemErrorsItems) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "feature", "minimum_version", "reason", "since", "type"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the ModuleStatItemErrorsItems object to a map representation for JSON marshaling.
func (m ModuleStatItemErrorsItems) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, m.AdditionalProperties)
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

// UnmarshalJSON implements the json.Unmarshaler interface for ModuleStatItemErrorsItems.
// It customizes the JSON unmarshaling process for ModuleStatItemErrorsItems objects.
func (m *ModuleStatItemErrorsItems) UnmarshalJSON(input []byte) error {
    var temp tempModuleStatItemErrorsItems
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "feature", "minimum_version", "reason", "since", "type")
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

// tempModuleStatItemErrorsItems is a temporary struct used for validating the fields of ModuleStatItemErrorsItems.
type tempModuleStatItemErrorsItems  struct {
    Feature        *string `json:"feature,omitempty"`
    MinimumVersion *string `json:"minimum_version,omitempty"`
    Reason         *string `json:"reason,omitempty"`
    Since          *int    `json:"since"`
    Type           *string `json:"type"`
}

func (m *tempModuleStatItemErrorsItems) validate() error {
    var errs []string
    if m.Since == nil {
        errs = append(errs, "required field `since` is missing for type `module_stat_item_errors_items`")
    }
    if m.Type == nil {
        errs = append(errs, "required field `type` is missing for type `module_stat_item_errors_items`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
