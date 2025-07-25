// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// WebsocketSessionWithUrl represents a WebsocketSessionWithUrl struct.
type WebsocketSessionWithUrl struct {
    Session              string                 `json:"session"`
    Url                  string                 `json:"url"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WebsocketSessionWithUrl,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WebsocketSessionWithUrl) String() string {
    return fmt.Sprintf(
    	"WebsocketSessionWithUrl[Session=%v, Url=%v, AdditionalProperties=%v]",
    	w.Session, w.Url, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WebsocketSessionWithUrl.
// It customizes the JSON marshaling process for WebsocketSessionWithUrl objects.
func (w WebsocketSessionWithUrl) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "session", "url"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WebsocketSessionWithUrl object to a map representation for JSON marshaling.
func (w WebsocketSessionWithUrl) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    structMap["session"] = w.Session
    structMap["url"] = w.Url
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebsocketSessionWithUrl.
// It customizes the JSON unmarshaling process for WebsocketSessionWithUrl objects.
func (w *WebsocketSessionWithUrl) UnmarshalJSON(input []byte) error {
    var temp tempWebsocketSessionWithUrl
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "session", "url")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.Session = *temp.Session
    w.Url = *temp.Url
    return nil
}

// tempWebsocketSessionWithUrl is a temporary struct used for validating the fields of WebsocketSessionWithUrl.
type tempWebsocketSessionWithUrl  struct {
    Session *string `json:"session"`
    Url     *string `json:"url"`
}

func (w *tempWebsocketSessionWithUrl) validate() error {
    var errs []string
    if w.Session == nil {
        errs = append(errs, "required field `session` is missing for type `websocket_session_with_url`")
    }
    if w.Url == nil {
        errs = append(errs, "required field `url` is missing for type `websocket_session_with_url`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
