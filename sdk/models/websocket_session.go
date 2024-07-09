package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// WebsocketSession represents a WebsocketSession struct.
type WebsocketSession struct {
    Session              string         `json:"session"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WebsocketSession.
// It customizes the JSON marshaling process for WebsocketSession objects.
func (w WebsocketSession) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the WebsocketSession object to a map representation for JSON marshaling.
func (w WebsocketSession) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, w.AdditionalProperties)
    structMap["session"] = w.Session
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebsocketSession.
// It customizes the JSON unmarshaling process for WebsocketSession objects.
func (w *WebsocketSession) UnmarshalJSON(input []byte) error {
    var temp websocketSession
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "session")
    if err != nil {
    	return err
    }
    
    w.AdditionalProperties = additionalProperties
    w.Session = *temp.Session
    return nil
}

// websocketSession is a temporary struct used for validating the fields of WebsocketSession.
type websocketSession  struct {
    Session *string `json:"session"`
}

func (w *websocketSession) validate() error {
    var errs []string
    if w.Session == nil {
        errs = append(errs, "required field `session` is missing for type `Websocket_Session`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
