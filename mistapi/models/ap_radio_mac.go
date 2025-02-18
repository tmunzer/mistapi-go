package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// ApRadioMac represents a ApRadioMac struct.
type ApRadioMac struct {
    Mac                  string                 `json:"mac"`
    RadioMacs            []string               `json:"radio_macs"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ApRadioMac,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a ApRadioMac) String() string {
    return fmt.Sprintf(
    	"ApRadioMac[Mac=%v, RadioMacs=%v, AdditionalProperties=%v]",
    	a.Mac, a.RadioMacs, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ApRadioMac.
// It customizes the JSON marshaling process for ApRadioMac objects.
func (a ApRadioMac) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "mac", "radio_macs"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the ApRadioMac object to a map representation for JSON marshaling.
func (a ApRadioMac) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    structMap["mac"] = a.Mac
    structMap["radio_macs"] = a.RadioMacs
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApRadioMac.
// It customizes the JSON unmarshaling process for ApRadioMac objects.
func (a *ApRadioMac) UnmarshalJSON(input []byte) error {
    var temp tempApRadioMac
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "mac", "radio_macs")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.Mac = *temp.Mac
    a.RadioMacs = *temp.RadioMacs
    return nil
}

// tempApRadioMac is a temporary struct used for validating the fields of ApRadioMac.
type tempApRadioMac  struct {
    Mac       *string   `json:"mac"`
    RadioMacs *[]string `json:"radio_macs"`
}

func (a *tempApRadioMac) validate() error {
    var errs []string
    if a.Mac == nil {
        errs = append(errs, "required field `mac` is missing for type `ap_radio_mac`")
    }
    if a.RadioMacs == nil {
        errs = append(errs, "required field `radio_macs` is missing for type `ap_radio_mac`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
