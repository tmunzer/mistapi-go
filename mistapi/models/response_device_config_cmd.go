package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ResponseDeviceConfigCmd represents a ResponseDeviceConfigCmd struct.
type ResponseDeviceConfigCmd struct {
    Cmd                  string                 `json:"cmd"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseDeviceConfigCmd.
// It customizes the JSON marshaling process for ResponseDeviceConfigCmd objects.
func (r ResponseDeviceConfigCmd) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "cmd"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseDeviceConfigCmd object to a map representation for JSON marshaling.
func (r ResponseDeviceConfigCmd) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["cmd"] = r.Cmd
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseDeviceConfigCmd.
// It customizes the JSON unmarshaling process for ResponseDeviceConfigCmd objects.
func (r *ResponseDeviceConfigCmd) UnmarshalJSON(input []byte) error {
    var temp tempResponseDeviceConfigCmd
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "cmd")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.Cmd = *temp.Cmd
    return nil
}

// tempResponseDeviceConfigCmd is a temporary struct used for validating the fields of ResponseDeviceConfigCmd.
type tempResponseDeviceConfigCmd  struct {
    Cmd *string `json:"cmd"`
}

func (r *tempResponseDeviceConfigCmd) validate() error {
    var errs []string
    if r.Cmd == nil {
        errs = append(errs, "required field `cmd` is missing for type `response_device_config_cmd`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
