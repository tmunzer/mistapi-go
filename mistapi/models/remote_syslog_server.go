package models

import (
	"encoding/json"
)

// RemoteSyslogServer represents a RemoteSyslogServer struct.
type RemoteSyslogServer struct {
	Contents         []RemoteSyslogContent `json:"contents,omitempty"`
	ExplicitPriority *bool                 `json:"explicit_priority,omitempty"`
	// enum: `any`, `authorization`, `change-log`, `config`, `conflict-log`, `daemon`, `dfc`, `external`, `firewall`, `ftp`, `interactive-commands`, `kernel`, `ntp`, `pfe`, `security`, `user`
	Facility *RemoteSyslogFacilityEnum `json:"facility,omitempty"`
	Host     *string                   `json:"host,omitempty"`
	Match    *string                   `json:"match,omitempty"`
	Port     *int                      `json:"port,omitempty"`
	// enum: `tcp`, `udp`
	Protocol        *RemoteSyslogServerProtocolEnum `json:"protocol,omitempty"`
	RoutingInstance *string                         `json:"routing_instance,omitempty"`
	// enum: `alert`, `any`, `critical`, `emergency`, `error`, `info`, `notice`, `warning`
	Severity *RemoteSyslogSeverityEnum `json:"severity,omitempty"`
	// if source_address is configured, will use the vlan firstly otherwise use source_ip
	SourceAddress        *string        `json:"source_address,omitempty"`
	StructuredData       *bool          `json:"structured_data,omitempty"`
	Tag                  *string        `json:"tag,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for RemoteSyslogServer.
// It customizes the JSON marshaling process for RemoteSyslogServer objects.
func (r RemoteSyslogServer) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RemoteSyslogServer object to a map representation for JSON marshaling.
func (r RemoteSyslogServer) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, r.AdditionalProperties)
	if r.Contents != nil {
		structMap["contents"] = r.Contents
	}
	if r.ExplicitPriority != nil {
		structMap["explicit_priority"] = r.ExplicitPriority
	}
	if r.Facility != nil {
		structMap["facility"] = r.Facility
	}
	if r.Host != nil {
		structMap["host"] = r.Host
	}
	if r.Match != nil {
		structMap["match"] = r.Match
	}
	if r.Port != nil {
		structMap["port"] = r.Port
	}
	if r.Protocol != nil {
		structMap["protocol"] = r.Protocol
	}
	if r.RoutingInstance != nil {
		structMap["routing_instance"] = r.RoutingInstance
	}
	if r.Severity != nil {
		structMap["severity"] = r.Severity
	}
	if r.SourceAddress != nil {
		structMap["source_address"] = r.SourceAddress
	}
	if r.StructuredData != nil {
		structMap["structured_data"] = r.StructuredData
	}
	if r.Tag != nil {
		structMap["tag"] = r.Tag
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RemoteSyslogServer.
// It customizes the JSON unmarshaling process for RemoteSyslogServer objects.
func (r *RemoteSyslogServer) UnmarshalJSON(input []byte) error {
	var temp tempRemoteSyslogServer
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "contents", "explicit_priority", "facility", "host", "match", "port", "protocol", "routing_instance", "severity", "source_address", "structured_data", "tag")
	if err != nil {
		return err
	}

	r.AdditionalProperties = additionalProperties
	r.Contents = temp.Contents
	r.ExplicitPriority = temp.ExplicitPriority
	r.Facility = temp.Facility
	r.Host = temp.Host
	r.Match = temp.Match
	r.Port = temp.Port
	r.Protocol = temp.Protocol
	r.RoutingInstance = temp.RoutingInstance
	r.Severity = temp.Severity
	r.SourceAddress = temp.SourceAddress
	r.StructuredData = temp.StructuredData
	r.Tag = temp.Tag
	return nil
}

// tempRemoteSyslogServer is a temporary struct used for validating the fields of RemoteSyslogServer.
type tempRemoteSyslogServer struct {
	Contents         []RemoteSyslogContent           `json:"contents,omitempty"`
	ExplicitPriority *bool                           `json:"explicit_priority,omitempty"`
	Facility         *RemoteSyslogFacilityEnum       `json:"facility,omitempty"`
	Host             *string                         `json:"host,omitempty"`
	Match            *string                         `json:"match,omitempty"`
	Port             *int                            `json:"port,omitempty"`
	Protocol         *RemoteSyslogServerProtocolEnum `json:"protocol,omitempty"`
	RoutingInstance  *string                         `json:"routing_instance,omitempty"`
	Severity         *RemoteSyslogSeverityEnum       `json:"severity,omitempty"`
	SourceAddress    *string                         `json:"source_address,omitempty"`
	StructuredData   *bool                           `json:"structured_data,omitempty"`
	Tag              *string                         `json:"tag,omitempty"`
}
