package models

import (
    "encoding/json"
)

// RemoteSyslogConsole represents a RemoteSyslogConsole struct.
type RemoteSyslogConsole struct {
    Contents             []RemoteSyslogContent `json:"contents,omitempty"`
    AdditionalProperties map[string]any        `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for RemoteSyslogConsole.
// It customizes the JSON marshaling process for RemoteSyslogConsole objects.
func (r RemoteSyslogConsole) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the RemoteSyslogConsole object to a map representation for JSON marshaling.
func (r RemoteSyslogConsole) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    if r.Contents != nil {
        structMap["contents"] = r.Contents
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RemoteSyslogConsole.
// It customizes the JSON unmarshaling process for RemoteSyslogConsole objects.
func (r *RemoteSyslogConsole) UnmarshalJSON(input []byte) error {
    var temp tempRemoteSyslogConsole
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "contents")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.Contents = temp.Contents
    return nil
}

// tempRemoteSyslogConsole is a temporary struct used for validating the fields of RemoteSyslogConsole.
type tempRemoteSyslogConsole  struct {
    Contents []RemoteSyslogContent `json:"contents,omitempty"`
}
