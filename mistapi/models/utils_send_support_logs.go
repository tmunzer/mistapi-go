package models

import (
    "encoding/json"
)

// UtilsSendSupportLogs represents a UtilsSendSupportLogs struct.
type UtilsSendSupportLogs struct {
    // optional choice: process, outbound-ssh, and full (default)
    Info                 *UtilsSendSupportLogsInfoEnum `json:"info,omitempty"`
    // optional: for SSR, if node is not present, both nodes support files are uploaded
    Node                 *string                       `json:"node,omitempty"`
    // optional: number of most recent messages files to upload.
    NumMessagesFiles     *int                          `json:"num_messages_files,omitempty"`
    AdditionalProperties map[string]any                `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UtilsSendSupportLogs.
// It customizes the JSON marshaling process for UtilsSendSupportLogs objects.
func (u UtilsSendSupportLogs) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UtilsSendSupportLogs object to a map representation for JSON marshaling.
func (u UtilsSendSupportLogs) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, u.AdditionalProperties)
    if u.Info != nil {
        structMap["info"] = u.Info
    }
    if u.Node != nil {
        structMap["node"] = u.Node
    }
    if u.NumMessagesFiles != nil {
        structMap["num_messages_files"] = u.NumMessagesFiles
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UtilsSendSupportLogs.
// It customizes the JSON unmarshaling process for UtilsSendSupportLogs objects.
func (u *UtilsSendSupportLogs) UnmarshalJSON(input []byte) error {
    var temp utilsSendSupportLogs
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "info", "node", "num_messages_files")
    if err != nil {
    	return err
    }
    
    u.AdditionalProperties = additionalProperties
    u.Info = temp.Info
    u.Node = temp.Node
    u.NumMessagesFiles = temp.NumMessagesFiles
    return nil
}

// utilsSendSupportLogs is a temporary struct used for validating the fields of UtilsSendSupportLogs.
type utilsSendSupportLogs  struct {
    Info             *UtilsSendSupportLogsInfoEnum `json:"info,omitempty"`
    Node             *string                       `json:"node,omitempty"`
    NumMessagesFiles *int                          `json:"num_messages_files,omitempty"`
}