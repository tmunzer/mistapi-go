package models

import (
	"encoding/json"
)

// RemoteSyslogArchive represents a RemoteSyslogArchive struct.
type RemoteSyslogArchive struct {
	Files                *int           `json:"files,omitempty"`
	Size                 *string        `json:"size,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for RemoteSyslogArchive.
// It customizes the JSON marshaling process for RemoteSyslogArchive objects.
func (r RemoteSyslogArchive) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RemoteSyslogArchive object to a map representation for JSON marshaling.
func (r RemoteSyslogArchive) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, r.AdditionalProperties)
	if r.Files != nil {
		structMap["files"] = r.Files
	}
	if r.Size != nil {
		structMap["size"] = r.Size
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RemoteSyslogArchive.
// It customizes the JSON unmarshaling process for RemoteSyslogArchive objects.
func (r *RemoteSyslogArchive) UnmarshalJSON(input []byte) error {
	var temp tempRemoteSyslogArchive
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "files", "size")
	if err != nil {
		return err
	}

	r.AdditionalProperties = additionalProperties
	r.Files = temp.Files
	r.Size = temp.Size
	return nil
}

// tempRemoteSyslogArchive is a temporary struct used for validating the fields of RemoteSyslogArchive.
type tempRemoteSyslogArchive struct {
	Files *int    `json:"files,omitempty"`
	Size  *string `json:"size,omitempty"`
}
