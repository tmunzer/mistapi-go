package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// WxtagMatching represents a WxtagMatching struct.
type WxtagMatching struct {
    Mac                  string                 `json:"mac"`
    Since                int                    `json:"since"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WxtagMatching,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WxtagMatching) String() string {
    return fmt.Sprintf(
    	"WxtagMatching[Mac=%v, Since=%v, AdditionalProperties=%v]",
    	w.Mac, w.Since, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WxtagMatching.
// It customizes the JSON marshaling process for WxtagMatching objects.
func (w WxtagMatching) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "mac", "since"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WxtagMatching object to a map representation for JSON marshaling.
func (w WxtagMatching) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    structMap["mac"] = w.Mac
    structMap["since"] = w.Since
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WxtagMatching.
// It customizes the JSON unmarshaling process for WxtagMatching objects.
func (w *WxtagMatching) UnmarshalJSON(input []byte) error {
    var temp tempWxtagMatching
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "mac", "since")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.Mac = *temp.Mac
    w.Since = *temp.Since
    return nil
}

// tempWxtagMatching is a temporary struct used for validating the fields of WxtagMatching.
type tempWxtagMatching  struct {
    Mac   *string `json:"mac"`
    Since *int    `json:"since"`
}

func (w *tempWxtagMatching) validate() error {
    var errs []string
    if w.Mac == nil {
        errs = append(errs, "required field `mac` is missing for type `wxtag_matching`")
    }
    if w.Since == nil {
        errs = append(errs, "required field `since` is missing for type `wxtag_matching`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
