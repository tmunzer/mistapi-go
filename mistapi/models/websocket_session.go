package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// WebsocketSession represents a WebsocketSession struct.
type WebsocketSession struct {
    Session              string                 `json:"session"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WebsocketSession,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WebsocketSession) String() string {
    return fmt.Sprintf(
    	"WebsocketSession[Session=%v, AdditionalProperties=%v]",
    	w.Session, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WebsocketSession.
// It customizes the JSON marshaling process for WebsocketSession objects.
func (w WebsocketSession) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "session"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WebsocketSession object to a map representation for JSON marshaling.
func (w WebsocketSession) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    structMap["session"] = w.Session
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebsocketSession.
// It customizes the JSON unmarshaling process for WebsocketSession objects.
func (w *WebsocketSession) UnmarshalJSON(input []byte) error {
    var temp tempWebsocketSession
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "session")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.Session = *temp.Session
    return nil
}

// tempWebsocketSession is a temporary struct used for validating the fields of WebsocketSession.
type tempWebsocketSession  struct {
    Session *string `json:"session"`
}

func (w *tempWebsocketSession) validate() error {
    var errs []string
    if w.Session == nil {
        errs = append(errs, "required field `session` is missing for type `websocket_session`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
