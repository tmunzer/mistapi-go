package models

import (
    "encoding/json"
    "fmt"
)

// RemoteSyslog represents a RemoteSyslog struct.
type RemoteSyslog struct {
    Archive              *RemoteSyslogArchive        `json:"archive,omitempty"`
    Console              *RemoteSyslogConsole        `json:"console,omitempty"`
    Enabled              *bool                       `json:"enabled,omitempty"`
    Files                []RemoteSyslogFileConfig    `json:"files,omitempty"`
    // If source_address is configured, will use the vlan firstly otherwise use source_ip
    Network              *string                     `json:"network,omitempty"`
    SendToAllServers     *bool                       `json:"send_to_all_servers,omitempty"`
    Servers              []RemoteSyslogServer        `json:"servers,omitempty"`
    // enum: `millisecond`, `year`, `year millisecond`
    TimeFormat           *RemoteSyslogTimeFormatEnum `json:"time_format,omitempty"`
    Users                []RemoteSyslogUser          `json:"users,omitempty"`
    AdditionalProperties map[string]interface{}      `json:"_"`
}

// String implements the fmt.Stringer interface for RemoteSyslog,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r RemoteSyslog) String() string {
    return fmt.Sprintf(
    	"RemoteSyslog[Archive=%v, Console=%v, Enabled=%v, Files=%v, Network=%v, SendToAllServers=%v, Servers=%v, TimeFormat=%v, Users=%v, AdditionalProperties=%v]",
    	r.Archive, r.Console, r.Enabled, r.Files, r.Network, r.SendToAllServers, r.Servers, r.TimeFormat, r.Users, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for RemoteSyslog.
// It customizes the JSON marshaling process for RemoteSyslog objects.
func (r RemoteSyslog) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "archive", "console", "enabled", "files", "network", "send_to_all_servers", "servers", "time_format", "users"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the RemoteSyslog object to a map representation for JSON marshaling.
func (r RemoteSyslog) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    if r.Archive != nil {
        structMap["archive"] = r.Archive.toMap()
    }
    if r.Console != nil {
        structMap["console"] = r.Console.toMap()
    }
    if r.Enabled != nil {
        structMap["enabled"] = r.Enabled
    }
    if r.Files != nil {
        structMap["files"] = r.Files
    }
    if r.Network != nil {
        structMap["network"] = r.Network
    }
    if r.SendToAllServers != nil {
        structMap["send_to_all_servers"] = r.SendToAllServers
    }
    if r.Servers != nil {
        structMap["servers"] = r.Servers
    }
    if r.TimeFormat != nil {
        structMap["time_format"] = r.TimeFormat
    }
    if r.Users != nil {
        structMap["users"] = r.Users
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RemoteSyslog.
// It customizes the JSON unmarshaling process for RemoteSyslog objects.
func (r *RemoteSyslog) UnmarshalJSON(input []byte) error {
    var temp tempRemoteSyslog
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "archive", "console", "enabled", "files", "network", "send_to_all_servers", "servers", "time_format", "users")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.Archive = temp.Archive
    r.Console = temp.Console
    r.Enabled = temp.Enabled
    r.Files = temp.Files
    r.Network = temp.Network
    r.SendToAllServers = temp.SendToAllServers
    r.Servers = temp.Servers
    r.TimeFormat = temp.TimeFormat
    r.Users = temp.Users
    return nil
}

// tempRemoteSyslog is a temporary struct used for validating the fields of RemoteSyslog.
type tempRemoteSyslog  struct {
    Archive          *RemoteSyslogArchive        `json:"archive,omitempty"`
    Console          *RemoteSyslogConsole        `json:"console,omitempty"`
    Enabled          *bool                       `json:"enabled,omitempty"`
    Files            []RemoteSyslogFileConfig    `json:"files,omitempty"`
    Network          *string                     `json:"network,omitempty"`
    SendToAllServers *bool                       `json:"send_to_all_servers,omitempty"`
    Servers          []RemoteSyslogServer        `json:"servers,omitempty"`
    TimeFormat       *RemoteSyslogTimeFormatEnum `json:"time_format,omitempty"`
    Users            []RemoteSyslogUser          `json:"users,omitempty"`
}
