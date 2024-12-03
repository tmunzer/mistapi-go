package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ResponseDeviceConfigCli represents a ResponseDeviceConfigCli struct.
type ResponseDeviceConfigCli struct {
    Cli                  []string               `json:"cli"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseDeviceConfigCli.
// It customizes the JSON marshaling process for ResponseDeviceConfigCli objects.
func (r ResponseDeviceConfigCli) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "cli"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseDeviceConfigCli object to a map representation for JSON marshaling.
func (r ResponseDeviceConfigCli) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["cli"] = r.Cli
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseDeviceConfigCli.
// It customizes the JSON unmarshaling process for ResponseDeviceConfigCli objects.
func (r *ResponseDeviceConfigCli) UnmarshalJSON(input []byte) error {
    var temp tempResponseDeviceConfigCli
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "cli")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.Cli = *temp.Cli
    return nil
}

// tempResponseDeviceConfigCli is a temporary struct used for validating the fields of ResponseDeviceConfigCli.
type tempResponseDeviceConfigCli  struct {
    Cli *[]string `json:"cli"`
}

func (r *tempResponseDeviceConfigCli) validate() error {
    var errs []string
    if r.Cli == nil {
        errs = append(errs, "required field `cli` is missing for type `response_device_config_cli`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
