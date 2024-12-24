package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/google/uuid"
    "strings"
)

// DeviceIdString represents a DeviceIdString struct.
type DeviceIdString struct {
    DeviceId             uuid.UUID              `json:"device_id"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for DeviceIdString,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (d DeviceIdString) String() string {
    return fmt.Sprintf(
    	"DeviceIdString[DeviceId=%v, AdditionalProperties=%v]",
    	d.DeviceId, d.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for DeviceIdString.
// It customizes the JSON marshaling process for DeviceIdString objects.
func (d DeviceIdString) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(d.AdditionalProperties,
        "device_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(d.toMap())
}

// toMap converts the DeviceIdString object to a map representation for JSON marshaling.
func (d DeviceIdString) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, d.AdditionalProperties)
    structMap["device_id"] = d.DeviceId
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DeviceIdString.
// It customizes the JSON unmarshaling process for DeviceIdString objects.
func (d *DeviceIdString) UnmarshalJSON(input []byte) error {
    var temp tempDeviceIdString
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "device_id")
    if err != nil {
    	return err
    }
    d.AdditionalProperties = additionalProperties
    
    d.DeviceId = *temp.DeviceId
    return nil
}

// tempDeviceIdString is a temporary struct used for validating the fields of DeviceIdString.
type tempDeviceIdString  struct {
    DeviceId *uuid.UUID `json:"device_id"`
}

func (d *tempDeviceIdString) validate() error {
    var errs []string
    if d.DeviceId == nil {
        errs = append(errs, "required field `device_id` is missing for type `device_id_string`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
