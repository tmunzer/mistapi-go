package models

import (
    "encoding/json"
)

// RemoteSyslogFileConfig represents a RemoteSyslogFileConfig struct.
type RemoteSyslogFileConfig struct {
    Archive              *RemoteSyslogArchive   `json:"archive,omitempty"`
    Contents             []RemoteSyslogContent  `json:"contents,omitempty"`
    ExplicitPriority     *bool                  `json:"explicit_priority,omitempty"`
    File                 *string                `json:"file,omitempty"`
    Match                *string                `json:"match,omitempty"`
    StructuredData       *bool                  `json:"structured_data,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for RemoteSyslogFileConfig.
// It customizes the JSON marshaling process for RemoteSyslogFileConfig objects.
func (r RemoteSyslogFileConfig) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "archive", "contents", "explicit_priority", "file", "match", "structured_data"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the RemoteSyslogFileConfig object to a map representation for JSON marshaling.
func (r RemoteSyslogFileConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    if r.Archive != nil {
        structMap["archive"] = r.Archive.toMap()
    }
    if r.Contents != nil {
        structMap["contents"] = r.Contents
    }
    if r.ExplicitPriority != nil {
        structMap["explicit_priority"] = r.ExplicitPriority
    }
    if r.File != nil {
        structMap["file"] = r.File
    }
    if r.Match != nil {
        structMap["match"] = r.Match
    }
    if r.StructuredData != nil {
        structMap["structured_data"] = r.StructuredData
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RemoteSyslogFileConfig.
// It customizes the JSON unmarshaling process for RemoteSyslogFileConfig objects.
func (r *RemoteSyslogFileConfig) UnmarshalJSON(input []byte) error {
    var temp tempRemoteSyslogFileConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "archive", "contents", "explicit_priority", "file", "match", "structured_data")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.Archive = temp.Archive
    r.Contents = temp.Contents
    r.ExplicitPriority = temp.ExplicitPriority
    r.File = temp.File
    r.Match = temp.Match
    r.StructuredData = temp.StructuredData
    return nil
}

// tempRemoteSyslogFileConfig is a temporary struct used for validating the fields of RemoteSyslogFileConfig.
type tempRemoteSyslogFileConfig  struct {
    Archive          *RemoteSyslogArchive  `json:"archive,omitempty"`
    Contents         []RemoteSyslogContent `json:"contents,omitempty"`
    ExplicitPriority *bool                 `json:"explicit_priority,omitempty"`
    File             *string               `json:"file,omitempty"`
    Match            *string               `json:"match,omitempty"`
    StructuredData   *bool                 `json:"structured_data,omitempty"`
}
