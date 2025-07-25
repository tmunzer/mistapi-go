// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// WifiBeaconExtendedInfoItems represents a WifiBeaconExtendedInfoItems struct.
type WifiBeaconExtendedInfoItems struct {
    // Frame control field of 802.11 header
    FrameCtrl            *int                   `json:"frame_ctrl,omitempty"`
    // Extended Info Payload associated with frame
    Payload              *string                `json:"payload,omitempty"`
    // Sequence control field of 802.11 header
    SeqCtrl              *int                   `json:"seq_ctrl,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WifiBeaconExtendedInfoItems,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WifiBeaconExtendedInfoItems) String() string {
    return fmt.Sprintf(
    	"WifiBeaconExtendedInfoItems[FrameCtrl=%v, Payload=%v, SeqCtrl=%v, AdditionalProperties=%v]",
    	w.FrameCtrl, w.Payload, w.SeqCtrl, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WifiBeaconExtendedInfoItems.
// It customizes the JSON marshaling process for WifiBeaconExtendedInfoItems objects.
func (w WifiBeaconExtendedInfoItems) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "frame_ctrl", "payload", "seq_ctrl"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WifiBeaconExtendedInfoItems object to a map representation for JSON marshaling.
func (w WifiBeaconExtendedInfoItems) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    if w.FrameCtrl != nil {
        structMap["frame_ctrl"] = w.FrameCtrl
    }
    if w.Payload != nil {
        structMap["payload"] = w.Payload
    }
    if w.SeqCtrl != nil {
        structMap["seq_ctrl"] = w.SeqCtrl
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WifiBeaconExtendedInfoItems.
// It customizes the JSON unmarshaling process for WifiBeaconExtendedInfoItems objects.
func (w *WifiBeaconExtendedInfoItems) UnmarshalJSON(input []byte) error {
    var temp tempWifiBeaconExtendedInfoItems
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "frame_ctrl", "payload", "seq_ctrl")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.FrameCtrl = temp.FrameCtrl
    w.Payload = temp.Payload
    w.SeqCtrl = temp.SeqCtrl
    return nil
}

// tempWifiBeaconExtendedInfoItems is a temporary struct used for validating the fields of WifiBeaconExtendedInfoItems.
type tempWifiBeaconExtendedInfoItems  struct {
    FrameCtrl *int    `json:"frame_ctrl,omitempty"`
    Payload   *string `json:"payload,omitempty"`
    SeqCtrl   *int    `json:"seq_ctrl,omitempty"`
}
