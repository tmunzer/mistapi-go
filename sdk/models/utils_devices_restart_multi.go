package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// UtilsDevicesRestartMulti represents a UtilsDevicesRestartMulti struct.
type UtilsDevicesRestartMulti struct {
    DeviceIds            []uuid.UUID    `json:"device_ids"`
    // only for SSR: if node is not present, both nodes are restarted
    // for other devices: node should not be present
    Node                 *string        `json:"node,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UtilsDevicesRestartMulti.
// It customizes the JSON marshaling process for UtilsDevicesRestartMulti objects.
func (u UtilsDevicesRestartMulti) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UtilsDevicesRestartMulti object to a map representation for JSON marshaling.
func (u UtilsDevicesRestartMulti) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, u.AdditionalProperties)
    structMap["device_ids"] = u.DeviceIds
    if u.Node != nil {
        structMap["node"] = u.Node
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UtilsDevicesRestartMulti.
// It customizes the JSON unmarshaling process for UtilsDevicesRestartMulti objects.
func (u *UtilsDevicesRestartMulti) UnmarshalJSON(input []byte) error {
    var temp utilsDevicesRestartMulti
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "device_ids", "node")
    if err != nil {
    	return err
    }
    
    u.AdditionalProperties = additionalProperties
    u.DeviceIds = *temp.DeviceIds
    u.Node = temp.Node
    return nil
}

// utilsDevicesRestartMulti is a temporary struct used for validating the fields of UtilsDevicesRestartMulti.
type utilsDevicesRestartMulti  struct {
    DeviceIds *[]uuid.UUID `json:"device_ids"`
    Node      *string      `json:"node,omitempty"`
}

func (u *utilsDevicesRestartMulti) validate() error {
    var errs []string
    if u.DeviceIds == nil {
        errs = append(errs, "required field `device_ids` is missing for type `Utils_Devices_Restart_Multi`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
