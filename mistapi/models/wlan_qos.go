package models

import (
    "encoding/json"
    "fmt"
)

// WlanQos represents a WlanQos struct.
type WlanQos struct {
    // enum: `background`, `best_effort`, `video`, `voice`
    Class                *WlanQosClassEnum      `json:"class,omitempty"`
    // whether to overwrite QoS
    Overwrite            *bool                  `json:"overwrite,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WlanQos,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WlanQos) String() string {
    return fmt.Sprintf(
    	"WlanQos[Class=%v, Overwrite=%v, AdditionalProperties=%v]",
    	w.Class, w.Overwrite, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WlanQos.
// It customizes the JSON marshaling process for WlanQos objects.
func (w WlanQos) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "class", "overwrite"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WlanQos object to a map representation for JSON marshaling.
func (w WlanQos) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    if w.Class != nil {
        structMap["class"] = w.Class
    }
    if w.Overwrite != nil {
        structMap["overwrite"] = w.Overwrite
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WlanQos.
// It customizes the JSON unmarshaling process for WlanQos objects.
func (w *WlanQos) UnmarshalJSON(input []byte) error {
    var temp tempWlanQos
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "class", "overwrite")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.Class = temp.Class
    w.Overwrite = temp.Overwrite
    return nil
}

// tempWlanQos is a temporary struct used for validating the fields of WlanQos.
type tempWlanQos  struct {
    Class     *WlanQosClassEnum `json:"class,omitempty"`
    Overwrite *bool             `json:"overwrite,omitempty"`
}
