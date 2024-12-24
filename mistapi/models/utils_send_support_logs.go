package models

import (
    "encoding/json"
    "fmt"
)

// UtilsSendSupportLogs represents a UtilsSendSupportLogs struct.
type UtilsSendSupportLogs struct {
    // optional, enum:
    // * `code-dumps`: Upload all core dump files, if any found
    // * `full`: Upload 1 file with output of `request support information`, 1 file that concatenates all `/var/log/outbound-ssh.log*` files, all core dump files, the 5 most recent `/var/log/messages*` files, and Mist agent logs
    // * `messages`: Upload 1 to 10 `/var/log/messages*` files
    // * `outbound-ssh`: Upload 1 file that concatenates all `/var/log/outbound-ssh.log*` files
    // * `process`: Upload 1 file with output of show `system processes extensive``
    // * `var-logs`: Upload all non-empty files in the `/var/log/` directory
    Info                 *UtilsSendSupportLogsInfoEnum `json:"info,omitempty"`
    // optional: for SSR, if node is not present, both nodes support files are uploaded
    Node                 *string                       `json:"node,omitempty"`
    // optional: number of most recent messages files to upload.
    NumMessagesFiles     *int                          `json:"num_messages_files,omitempty"`
    AdditionalProperties map[string]interface{}        `json:"_"`
}

// String implements the fmt.Stringer interface for UtilsSendSupportLogs,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UtilsSendSupportLogs) String() string {
    return fmt.Sprintf(
    	"UtilsSendSupportLogs[Info=%v, Node=%v, NumMessagesFiles=%v, AdditionalProperties=%v]",
    	u.Info, u.Node, u.NumMessagesFiles, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UtilsSendSupportLogs.
// It customizes the JSON marshaling process for UtilsSendSupportLogs objects.
func (u UtilsSendSupportLogs) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "info", "node", "num_messages_files"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UtilsSendSupportLogs object to a map representation for JSON marshaling.
func (u UtilsSendSupportLogs) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
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
    var temp tempUtilsSendSupportLogs
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "info", "node", "num_messages_files")
    if err != nil {
    	return err
    }
    u.AdditionalProperties = additionalProperties
    
    u.Info = temp.Info
    u.Node = temp.Node
    u.NumMessagesFiles = temp.NumMessagesFiles
    return nil
}

// tempUtilsSendSupportLogs is a temporary struct used for validating the fields of UtilsSendSupportLogs.
type tempUtilsSendSupportLogs  struct {
    Info             *UtilsSendSupportLogsInfoEnum `json:"info,omitempty"`
    Node             *string                       `json:"node,omitempty"`
    NumMessagesFiles *int                          `json:"num_messages_files,omitempty"`
}
