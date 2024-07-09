package models

import (
    "encoding/json"
)

// WlanQos represents a WlanQos struct.
type WlanQos struct {
    Class                *WlanQosClassEnum `json:"class,omitempty"`
    // whether to overwrite QoS
    Overwrite            *bool             `json:"overwrite,omitempty"`
    AdditionalProperties map[string]any    `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WlanQos.
// It customizes the JSON marshaling process for WlanQos objects.
func (w WlanQos) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the WlanQos object to a map representation for JSON marshaling.
func (w WlanQos) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, w.AdditionalProperties)
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
    var temp wlanQos
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "class", "overwrite")
    if err != nil {
    	return err
    }
    
    w.AdditionalProperties = additionalProperties
    w.Class = temp.Class
    w.Overwrite = temp.Overwrite
    return nil
}

// wlanQos is a temporary struct used for validating the fields of WlanQos.
type wlanQos  struct {
    Class     *WlanQosClassEnum `json:"class,omitempty"`
    Overwrite *bool             `json:"overwrite,omitempty"`
}