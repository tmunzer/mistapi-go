package models

import (
    "encoding/json"
)

// RemoteSyslogConsole represents a RemoteSyslogConsole struct.
type RemoteSyslogConsole struct {
    Contents             []RemoteSyslogContent  `json:"contents,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for RemoteSyslogConsole.
// It customizes the JSON marshaling process for RemoteSyslogConsole objects.
func (r RemoteSyslogConsole) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "contents"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the RemoteSyslogConsole object to a map representation for JSON marshaling.
func (r RemoteSyslogConsole) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "contents")
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
