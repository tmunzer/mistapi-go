/*
Mist API

> Version: **2406.1.14** > > Date: **July 3, 2024**  ---  ### Additional Documentation * [Mist Automation Guide](https://www.juniper.net/documentation/us/en/software/mist/automation-integration/index.html) * [Mist Location SDK](https://www.juniper.net/documentation/us/en/software/mist/location_services/topics/concept/mist-how-get-mist-sdk.html) * [Mist Product Updates](https://www.mist.com/documentation/category/product-updates/)  ---  ### Helpful Resources * [API Sandbox and Exercises](https://api-class.mist.com/) * [Postman Collection, Runners and Webhook Samples](https://www.postman.com/juniper-mist/workspace/mist-systems-s-public-workspace) * [API Demo Apps](https://apps.mist-lab.fr/) * [Juniper Blog](https://blogs.juniper.net/)  --- 

API version: 2406.1.14
Contact: tmunzer@juniper.net
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mistapigo

import (
	"encoding/json"
	"fmt"
)

// RemoteSyslogSeverity the model 'RemoteSyslogSeverity'
type RemoteSyslogSeverity string

// List of remote_syslog_severity
const (
	REMOTESYSLOGSEVERITY_EMPTY RemoteSyslogSeverity = ""
	REMOTESYSLOGSEVERITY_ANY RemoteSyslogSeverity = "any"
	REMOTESYSLOGSEVERITY_ALERT RemoteSyslogSeverity = "alert"
	REMOTESYSLOGSEVERITY_EMERGENCY RemoteSyslogSeverity = "emergency"
	REMOTESYSLOGSEVERITY_CRITICAL RemoteSyslogSeverity = "critical"
	REMOTESYSLOGSEVERITY_WARNING RemoteSyslogSeverity = "warning"
	REMOTESYSLOGSEVERITY_INFO RemoteSyslogSeverity = "info"
	REMOTESYSLOGSEVERITY_NOTICE RemoteSyslogSeverity = "notice"
	REMOTESYSLOGSEVERITY_ERROR RemoteSyslogSeverity = "error"
)

// All allowed values of RemoteSyslogSeverity enum
var AllowedRemoteSyslogSeverityEnumValues = []RemoteSyslogSeverity{
	"",
	"any",
	"alert",
	"emergency",
	"critical",
	"warning",
	"info",
	"notice",
	"error",
}

func (v *RemoteSyslogSeverity) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RemoteSyslogSeverity(value)
	for _, existing := range AllowedRemoteSyslogSeverityEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RemoteSyslogSeverity", value)
}

// NewRemoteSyslogSeverityFromValue returns a pointer to a valid RemoteSyslogSeverity
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRemoteSyslogSeverityFromValue(v string) (*RemoteSyslogSeverity, error) {
	ev := RemoteSyslogSeverity(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RemoteSyslogSeverity: valid values are %v", v, AllowedRemoteSyslogSeverityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RemoteSyslogSeverity) IsValid() bool {
	for _, existing := range AllowedRemoteSyslogSeverityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to remote_syslog_severity value
func (v RemoteSyslogSeverity) Ptr() *RemoteSyslogSeverity {
	return &v
}

type NullableRemoteSyslogSeverity struct {
	value *RemoteSyslogSeverity
	isSet bool
}

func (v NullableRemoteSyslogSeverity) Get() *RemoteSyslogSeverity {
	return v.value
}

func (v *NullableRemoteSyslogSeverity) Set(val *RemoteSyslogSeverity) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoteSyslogSeverity) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoteSyslogSeverity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoteSyslogSeverity(val *RemoteSyslogSeverity) *NullableRemoteSyslogSeverity {
	return &NullableRemoteSyslogSeverity{value: val, isSet: true}
}

func (v NullableRemoteSyslogSeverity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoteSyslogSeverity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

