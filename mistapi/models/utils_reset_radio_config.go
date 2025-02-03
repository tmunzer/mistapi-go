package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// UtilsResetRadioConfig represents a UtilsResetRadioConfig struct.
type UtilsResetRadioConfig struct {
    // List of bands
    Bands                []string               `json:"bands"`
    // Whether to reset those with radio disabled. default is false (i.e. if user intentionally disables a radio, honor it)
    Force                *bool                  `json:"force,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UtilsResetRadioConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UtilsResetRadioConfig) String() string {
    return fmt.Sprintf(
    	"UtilsResetRadioConfig[Bands=%v, Force=%v, AdditionalProperties=%v]",
    	u.Bands, u.Force, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UtilsResetRadioConfig.
// It customizes the JSON marshaling process for UtilsResetRadioConfig objects.
func (u UtilsResetRadioConfig) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "bands", "force"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UtilsResetRadioConfig object to a map representation for JSON marshaling.
func (u UtilsResetRadioConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
    structMap["bands"] = u.Bands
    if u.Force != nil {
        structMap["force"] = u.Force
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UtilsResetRadioConfig.
// It customizes the JSON unmarshaling process for UtilsResetRadioConfig objects.
func (u *UtilsResetRadioConfig) UnmarshalJSON(input []byte) error {
    var temp tempUtilsResetRadioConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "bands", "force")
    if err != nil {
    	return err
    }
    u.AdditionalProperties = additionalProperties
    
    u.Bands = *temp.Bands
    u.Force = temp.Force
    return nil
}

// tempUtilsResetRadioConfig is a temporary struct used for validating the fields of UtilsResetRadioConfig.
type tempUtilsResetRadioConfig  struct {
    Bands *[]string `json:"bands"`
    Force *bool     `json:"force,omitempty"`
}

func (u *tempUtilsResetRadioConfig) validate() error {
    var errs []string
    if u.Bands == nil {
        errs = append(errs, "required field `bands` is missing for type `utils_reset_radio_config`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
