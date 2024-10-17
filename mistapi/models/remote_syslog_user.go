package models

import (
	"encoding/json"
)

// RemoteSyslogUser represents a RemoteSyslogUser struct.
type RemoteSyslogUser struct {
	Contents             []RemoteSyslogContent `json:"contents,omitempty"`
	Match                *string               `json:"match,omitempty"`
	User                 *string               `json:"user,omitempty"`
	AdditionalProperties map[string]any        `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for RemoteSyslogUser.
// It customizes the JSON marshaling process for RemoteSyslogUser objects.
func (r RemoteSyslogUser) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RemoteSyslogUser object to a map representation for JSON marshaling.
func (r RemoteSyslogUser) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, r.AdditionalProperties)
	if r.Contents != nil {
		structMap["contents"] = r.Contents
	}
	if r.Match != nil {
		structMap["match"] = r.Match
	}
	if r.User != nil {
		structMap["user"] = r.User
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RemoteSyslogUser.
// It customizes the JSON unmarshaling process for RemoteSyslogUser objects.
func (r *RemoteSyslogUser) UnmarshalJSON(input []byte) error {
	var temp tempRemoteSyslogUser
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "contents", "match", "user")
	if err != nil {
		return err
	}

	r.AdditionalProperties = additionalProperties
	r.Contents = temp.Contents
	r.Match = temp.Match
	r.User = temp.User
	return nil
}

// tempRemoteSyslogUser is a temporary struct used for validating the fields of RemoteSyslogUser.
type tempRemoteSyslogUser struct {
	Contents []RemoteSyslogContent `json:"contents,omitempty"`
	Match    *string               `json:"match,omitempty"`
	User     *string               `json:"user,omitempty"`
}
