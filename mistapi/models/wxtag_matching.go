package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// WxtagMatching represents a WxtagMatching struct.
type WxtagMatching struct {
    Mac                  string         `json:"mac"`
    Since                int            `json:"since"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WxtagMatching.
// It customizes the JSON marshaling process for WxtagMatching objects.
func (w WxtagMatching) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the WxtagMatching object to a map representation for JSON marshaling.
func (w WxtagMatching) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, w.AdditionalProperties)
    structMap["mac"] = w.Mac
    structMap["since"] = w.Since
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WxtagMatching.
// It customizes the JSON unmarshaling process for WxtagMatching objects.
func (w *WxtagMatching) UnmarshalJSON(input []byte) error {
    var temp wxtagMatching
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "mac", "since")
    if err != nil {
    	return err
    }
    
    w.AdditionalProperties = additionalProperties
    w.Mac = *temp.Mac
    w.Since = *temp.Since
    return nil
}

// wxtagMatching is a temporary struct used for validating the fields of WxtagMatching.
type wxtagMatching  struct {
    Mac   *string `json:"mac"`
    Since *int    `json:"since"`
}

func (w *wxtagMatching) validate() error {
    var errs []string
    if w.Mac == nil {
        errs = append(errs, "required field `mac` is missing for type `Wxtag_Matching`")
    }
    if w.Since == nil {
        errs = append(errs, "required field `since` is missing for type `Wxtag_Matching`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
