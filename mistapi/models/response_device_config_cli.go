package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ResponseDeviceConfigCli represents a ResponseDeviceConfigCli struct.
type ResponseDeviceConfigCli struct {
    Cli                  []string       `json:"cli"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseDeviceConfigCli.
// It customizes the JSON marshaling process for ResponseDeviceConfigCli objects.
func (r ResponseDeviceConfigCli) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseDeviceConfigCli object to a map representation for JSON marshaling.
func (r ResponseDeviceConfigCli) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["cli"] = r.Cli
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseDeviceConfigCli.
// It customizes the JSON unmarshaling process for ResponseDeviceConfigCli objects.
func (r *ResponseDeviceConfigCli) UnmarshalJSON(input []byte) error {
    var temp responseDeviceConfigCli
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "cli")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.Cli = *temp.Cli
    return nil
}

// responseDeviceConfigCli is a temporary struct used for validating the fields of ResponseDeviceConfigCli.
type responseDeviceConfigCli  struct {
    Cli *[]string `json:"cli"`
}

func (r *responseDeviceConfigCli) validate() error {
    var errs []string
    if r.Cli == nil {
        errs = append(errs, "required field `cli` is missing for type `Response_Device_Config_Cli`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
