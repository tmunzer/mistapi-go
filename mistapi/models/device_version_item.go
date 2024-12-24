package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// DeviceVersionItem represents a DeviceVersionItem struct.
type DeviceVersionItem struct {
    // Device model (as seen in the device stats)
    Model                string                 `json:"model"`
    // annotation, stable / beta / alpha. Or it can be empty or nothing which is likely a dev build
    Tag                  *string                `json:"tag,omitempty"`
    // firmware version
    Version              string                 `json:"version"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for DeviceVersionItem,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (d DeviceVersionItem) String() string {
    return fmt.Sprintf(
    	"DeviceVersionItem[Model=%v, Tag=%v, Version=%v, AdditionalProperties=%v]",
    	d.Model, d.Tag, d.Version, d.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for DeviceVersionItem.
// It customizes the JSON marshaling process for DeviceVersionItem objects.
func (d DeviceVersionItem) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(d.AdditionalProperties,
        "model", "tag", "version"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(d.toMap())
}

// toMap converts the DeviceVersionItem object to a map representation for JSON marshaling.
func (d DeviceVersionItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, d.AdditionalProperties)
    structMap["model"] = d.Model
    if d.Tag != nil {
        structMap["tag"] = d.Tag
    }
    structMap["version"] = d.Version
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DeviceVersionItem.
// It customizes the JSON unmarshaling process for DeviceVersionItem objects.
func (d *DeviceVersionItem) UnmarshalJSON(input []byte) error {
    var temp tempDeviceVersionItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "model", "tag", "version")
    if err != nil {
    	return err
    }
    d.AdditionalProperties = additionalProperties
    
    d.Model = *temp.Model
    d.Tag = temp.Tag
    d.Version = *temp.Version
    return nil
}

// tempDeviceVersionItem is a temporary struct used for validating the fields of DeviceVersionItem.
type tempDeviceVersionItem  struct {
    Model   *string `json:"model"`
    Tag     *string `json:"tag,omitempty"`
    Version *string `json:"version"`
}

func (d *tempDeviceVersionItem) validate() error {
    var errs []string
    if d.Model == nil {
        errs = append(errs, "required field `model` is missing for type `device_version_item`")
    }
    if d.Version == nil {
        errs = append(errs, "required field `version` is missing for type `device_version_item`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
