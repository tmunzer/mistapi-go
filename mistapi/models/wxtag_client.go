package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// WxtagClient represents a WxtagClient struct.
type WxtagClient struct {
    Mac                  string                 `json:"mac"`
    Since                float64                `json:"since"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WxtagClient,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WxtagClient) String() string {
    return fmt.Sprintf(
    	"WxtagClient[Mac=%v, Since=%v, AdditionalProperties=%v]",
    	w.Mac, w.Since, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WxtagClient.
// It customizes the JSON marshaling process for WxtagClient objects.
func (w WxtagClient) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "mac", "since"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WxtagClient object to a map representation for JSON marshaling.
func (w WxtagClient) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    structMap["mac"] = w.Mac
    structMap["since"] = w.Since
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WxtagClient.
// It customizes the JSON unmarshaling process for WxtagClient objects.
func (w *WxtagClient) UnmarshalJSON(input []byte) error {
    var temp tempWxtagClient
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

// tempWxtagClient is a temporary struct used for validating the fields of WxtagClient.
type tempWxtagClient  struct {
    Mac   *string  `json:"mac"`
    Since *float64 `json:"since"`
}

func (w *tempWxtagClient) validate() error {
    var errs []string
    if w.Mac == nil {
        errs = append(errs, "required field `mac` is missing for type `wxtag_client`")
    }
    if w.Since == nil {
        errs = append(errs, "required field `since` is missing for type `wxtag_client`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
