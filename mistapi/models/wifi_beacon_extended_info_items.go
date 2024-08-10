package models

import (
    "encoding/json"
)

// WifiBeaconExtendedInfoItems represents a WifiBeaconExtendedInfoItems struct.
type WifiBeaconExtendedInfoItems struct {
    // frame control field of 802.11 header
    FrameCtrl            *int           `json:"frame_ctrl,omitempty"`
    // Extended Info Payload associated with frame
    Payload              *string        `json:"payload,omitempty"`
    // sequence control field of 802.11 header
    SeqCtrl              *int           `json:"seq_ctrl,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WifiBeaconExtendedInfoItems.
// It customizes the JSON marshaling process for WifiBeaconExtendedInfoItems objects.
func (w WifiBeaconExtendedInfoItems) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the WifiBeaconExtendedInfoItems object to a map representation for JSON marshaling.
func (w WifiBeaconExtendedInfoItems) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, w.AdditionalProperties)
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
    additionalProperties, err := UnmarshalAdditionalProperties(input, "frame_ctrl", "payload", "seq_ctrl")
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
