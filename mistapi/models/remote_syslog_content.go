package models

import (
	"encoding/json"
)

// RemoteSyslogContent represents a RemoteSyslogContent struct.
type RemoteSyslogContent struct {
	// enum: `any`, `authorization`, `change-log`, `config`, `conflict-log`, `daemon`, `dfc`, `external`, `firewall`, `ftp`, `interactive-commands`, `kernel`, `ntp`, `pfe`, `security`, `user`
	Facility *RemoteSyslogFacilityEnum `json:"facility,omitempty"`
	// enum: `alert`, `any`, `critical`, `emergency`, `error`, `info`, `notice`, `warning`
	Severity             *RemoteSyslogSeverityEnum `json:"severity,omitempty"`
	AdditionalProperties map[string]any            `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for RemoteSyslogContent.
// It customizes the JSON marshaling process for RemoteSyslogContent objects.
func (r RemoteSyslogContent) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RemoteSyslogContent object to a map representation for JSON marshaling.
func (r RemoteSyslogContent) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, r.AdditionalProperties)
	if r.Facility != nil {
		structMap["facility"] = r.Facility
	}
	if r.Severity != nil {
		structMap["severity"] = r.Severity
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RemoteSyslogContent.
// It customizes the JSON unmarshaling process for RemoteSyslogContent objects.
func (r *RemoteSyslogContent) UnmarshalJSON(input []byte) error {
	var temp tempRemoteSyslogContent
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "facility", "severity")
	if err != nil {
		return err
	}

	r.AdditionalProperties = additionalProperties
	r.Facility = temp.Facility
	r.Severity = temp.Severity
	return nil
}

// tempRemoteSyslogContent is a temporary struct used for validating the fields of RemoteSyslogContent.
type tempRemoteSyslogContent struct {
	Facility *RemoteSyslogFacilityEnum `json:"facility,omitempty"`
	Severity *RemoteSyslogSeverityEnum `json:"severity,omitempty"`
}
