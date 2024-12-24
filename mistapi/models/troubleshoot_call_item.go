package models

import (
    "encoding/json"
    "fmt"
)

// TroubleshootCallItem represents a TroubleshootCallItem struct.
type TroubleshootCallItem struct {
    AudioIn              *CallTroubleshootData  `json:"audio_in,omitempty"`
    AudioOut             *CallTroubleshootData  `json:"audio_out,omitempty"`
    Timestamp            *int                   `json:"timestamp,omitempty"`
    VideoIn              *CallTroubleshootData  `json:"video_in,omitempty"`
    VideoOut             *CallTroubleshootData  `json:"video_out,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for TroubleshootCallItem,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (t TroubleshootCallItem) String() string {
    return fmt.Sprintf(
    	"TroubleshootCallItem[AudioIn=%v, AudioOut=%v, Timestamp=%v, VideoIn=%v, VideoOut=%v, AdditionalProperties=%v]",
    	t.AudioIn, t.AudioOut, t.Timestamp, t.VideoIn, t.VideoOut, t.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for TroubleshootCallItem.
// It customizes the JSON marshaling process for TroubleshootCallItem objects.
func (t TroubleshootCallItem) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(t.AdditionalProperties,
        "audio_in", "audio_out", "timestamp", "video_in", "video_out"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(t.toMap())
}

// toMap converts the TroubleshootCallItem object to a map representation for JSON marshaling.
func (t TroubleshootCallItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, t.AdditionalProperties)
    if t.AudioIn != nil {
        structMap["audio_in"] = t.AudioIn.toMap()
    }
    if t.AudioOut != nil {
        structMap["audio_out"] = t.AudioOut.toMap()
    }
    if t.Timestamp != nil {
        structMap["timestamp"] = t.Timestamp
    }
    if t.VideoIn != nil {
        structMap["video_in"] = t.VideoIn.toMap()
    }
    if t.VideoOut != nil {
        structMap["video_out"] = t.VideoOut.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TroubleshootCallItem.
// It customizes the JSON unmarshaling process for TroubleshootCallItem objects.
func (t *TroubleshootCallItem) UnmarshalJSON(input []byte) error {
    var temp tempTroubleshootCallItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "audio_in", "audio_out", "timestamp", "video_in", "video_out")
    if err != nil {
    	return err
    }
    t.AdditionalProperties = additionalProperties
    
    t.AudioIn = temp.AudioIn
    t.AudioOut = temp.AudioOut
    t.Timestamp = temp.Timestamp
    t.VideoIn = temp.VideoIn
    t.VideoOut = temp.VideoOut
    return nil
}

// tempTroubleshootCallItem is a temporary struct used for validating the fields of TroubleshootCallItem.
type tempTroubleshootCallItem  struct {
    AudioIn   *CallTroubleshootData `json:"audio_in,omitempty"`
    AudioOut  *CallTroubleshootData `json:"audio_out,omitempty"`
    Timestamp *int                  `json:"timestamp,omitempty"`
    VideoIn   *CallTroubleshootData `json:"video_in,omitempty"`
    VideoOut  *CallTroubleshootData `json:"video_out,omitempty"`
}
