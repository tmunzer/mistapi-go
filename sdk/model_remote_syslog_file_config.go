/*
Mist API

> Version: **2406.1.16** > > Date: **July 4, 2024**  ---  ### Additional Documentation * [Mist Automation Guide](https://www.juniper.net/documentation/us/en/software/mist/automation-integration/index.html) * [Mist Location SDK](https://www.juniper.net/documentation/us/en/software/mist/location_services/topics/concept/mist-how-get-mist-sdk.html) * [Mist Product Updates](https://www.mist.com/documentation/category/product-updates/)  ---  ### Helpful Resources * [API Sandbox and Exercises](https://api-class.mist.com/) * [Postman Collection, Runners and Webhook Samples](https://www.postman.com/juniper-mist/workspace/mist-systems-s-public-workspace) * [API Demo Apps](https://apps.mist-lab.fr/) * [Juniper Blog](https://blogs.juniper.net/)  --- 

API version: 2406.1.16
Contact: tmunzer@juniper.net
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mistapigo

import (
	"encoding/json"
)

// checks if the RemoteSyslogFileConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemoteSyslogFileConfig{}

// RemoteSyslogFileConfig struct for RemoteSyslogFileConfig
type RemoteSyslogFileConfig struct {
	Archive *RemoteSyslogArchive `json:"archive,omitempty"`
	Contents []RemoteSyslogContent `json:"contents,omitempty"`
	ExplicitPriority *bool `json:"explicit_priority,omitempty"`
	File *string `json:"file,omitempty"`
	Match *string `json:"match,omitempty"`
	StructuredData *bool `json:"structured_data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RemoteSyslogFileConfig RemoteSyslogFileConfig

// NewRemoteSyslogFileConfig instantiates a new RemoteSyslogFileConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoteSyslogFileConfig() *RemoteSyslogFileConfig {
	this := RemoteSyslogFileConfig{}
	return &this
}

// NewRemoteSyslogFileConfigWithDefaults instantiates a new RemoteSyslogFileConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoteSyslogFileConfigWithDefaults() *RemoteSyslogFileConfig {
	this := RemoteSyslogFileConfig{}
	return &this
}

// GetArchive returns the Archive field value if set, zero value otherwise.
func (o *RemoteSyslogFileConfig) GetArchive() RemoteSyslogArchive {
	if o == nil || IsNil(o.Archive) {
		var ret RemoteSyslogArchive
		return ret
	}
	return *o.Archive
}

// GetArchiveOk returns a tuple with the Archive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteSyslogFileConfig) GetArchiveOk() (*RemoteSyslogArchive, bool) {
	if o == nil || IsNil(o.Archive) {
		return nil, false
	}
	return o.Archive, true
}

// HasArchive returns a boolean if a field has been set.
func (o *RemoteSyslogFileConfig) HasArchive() bool {
	if o != nil && !IsNil(o.Archive) {
		return true
	}

	return false
}

// SetArchive gets a reference to the given RemoteSyslogArchive and assigns it to the Archive field.
func (o *RemoteSyslogFileConfig) SetArchive(v RemoteSyslogArchive) {
	o.Archive = &v
}

// GetContents returns the Contents field value if set, zero value otherwise.
func (o *RemoteSyslogFileConfig) GetContents() []RemoteSyslogContent {
	if o == nil || IsNil(o.Contents) {
		var ret []RemoteSyslogContent
		return ret
	}
	return o.Contents
}

// GetContentsOk returns a tuple with the Contents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteSyslogFileConfig) GetContentsOk() ([]RemoteSyslogContent, bool) {
	if o == nil || IsNil(o.Contents) {
		return nil, false
	}
	return o.Contents, true
}

// HasContents returns a boolean if a field has been set.
func (o *RemoteSyslogFileConfig) HasContents() bool {
	if o != nil && !IsNil(o.Contents) {
		return true
	}

	return false
}

// SetContents gets a reference to the given []RemoteSyslogContent and assigns it to the Contents field.
func (o *RemoteSyslogFileConfig) SetContents(v []RemoteSyslogContent) {
	o.Contents = v
}

// GetExplicitPriority returns the ExplicitPriority field value if set, zero value otherwise.
func (o *RemoteSyslogFileConfig) GetExplicitPriority() bool {
	if o == nil || IsNil(o.ExplicitPriority) {
		var ret bool
		return ret
	}
	return *o.ExplicitPriority
}

// GetExplicitPriorityOk returns a tuple with the ExplicitPriority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteSyslogFileConfig) GetExplicitPriorityOk() (*bool, bool) {
	if o == nil || IsNil(o.ExplicitPriority) {
		return nil, false
	}
	return o.ExplicitPriority, true
}

// HasExplicitPriority returns a boolean if a field has been set.
func (o *RemoteSyslogFileConfig) HasExplicitPriority() bool {
	if o != nil && !IsNil(o.ExplicitPriority) {
		return true
	}

	return false
}

// SetExplicitPriority gets a reference to the given bool and assigns it to the ExplicitPriority field.
func (o *RemoteSyslogFileConfig) SetExplicitPriority(v bool) {
	o.ExplicitPriority = &v
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *RemoteSyslogFileConfig) GetFile() string {
	if o == nil || IsNil(o.File) {
		var ret string
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteSyslogFileConfig) GetFileOk() (*string, bool) {
	if o == nil || IsNil(o.File) {
		return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *RemoteSyslogFileConfig) HasFile() bool {
	if o != nil && !IsNil(o.File) {
		return true
	}

	return false
}

// SetFile gets a reference to the given string and assigns it to the File field.
func (o *RemoteSyslogFileConfig) SetFile(v string) {
	o.File = &v
}

// GetMatch returns the Match field value if set, zero value otherwise.
func (o *RemoteSyslogFileConfig) GetMatch() string {
	if o == nil || IsNil(o.Match) {
		var ret string
		return ret
	}
	return *o.Match
}

// GetMatchOk returns a tuple with the Match field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteSyslogFileConfig) GetMatchOk() (*string, bool) {
	if o == nil || IsNil(o.Match) {
		return nil, false
	}
	return o.Match, true
}

// HasMatch returns a boolean if a field has been set.
func (o *RemoteSyslogFileConfig) HasMatch() bool {
	if o != nil && !IsNil(o.Match) {
		return true
	}

	return false
}

// SetMatch gets a reference to the given string and assigns it to the Match field.
func (o *RemoteSyslogFileConfig) SetMatch(v string) {
	o.Match = &v
}

// GetStructuredData returns the StructuredData field value if set, zero value otherwise.
func (o *RemoteSyslogFileConfig) GetStructuredData() bool {
	if o == nil || IsNil(o.StructuredData) {
		var ret bool
		return ret
	}
	return *o.StructuredData
}

// GetStructuredDataOk returns a tuple with the StructuredData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteSyslogFileConfig) GetStructuredDataOk() (*bool, bool) {
	if o == nil || IsNil(o.StructuredData) {
		return nil, false
	}
	return o.StructuredData, true
}

// HasStructuredData returns a boolean if a field has been set.
func (o *RemoteSyslogFileConfig) HasStructuredData() bool {
	if o != nil && !IsNil(o.StructuredData) {
		return true
	}

	return false
}

// SetStructuredData gets a reference to the given bool and assigns it to the StructuredData field.
func (o *RemoteSyslogFileConfig) SetStructuredData(v bool) {
	o.StructuredData = &v
}

func (o RemoteSyslogFileConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemoteSyslogFileConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Archive) {
		toSerialize["archive"] = o.Archive
	}
	if !IsNil(o.Contents) {
		toSerialize["contents"] = o.Contents
	}
	if !IsNil(o.ExplicitPriority) {
		toSerialize["explicit_priority"] = o.ExplicitPriority
	}
	if !IsNil(o.File) {
		toSerialize["file"] = o.File
	}
	if !IsNil(o.Match) {
		toSerialize["match"] = o.Match
	}
	if !IsNil(o.StructuredData) {
		toSerialize["structured_data"] = o.StructuredData
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RemoteSyslogFileConfig) UnmarshalJSON(data []byte) (err error) {
	varRemoteSyslogFileConfig := _RemoteSyslogFileConfig{}

	err = json.Unmarshal(data, &varRemoteSyslogFileConfig)

	if err != nil {
		return err
	}

	*o = RemoteSyslogFileConfig(varRemoteSyslogFileConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "archive")
		delete(additionalProperties, "contents")
		delete(additionalProperties, "explicit_priority")
		delete(additionalProperties, "file")
		delete(additionalProperties, "match")
		delete(additionalProperties, "structured_data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRemoteSyslogFileConfig struct {
	value *RemoteSyslogFileConfig
	isSet bool
}

func (v NullableRemoteSyslogFileConfig) Get() *RemoteSyslogFileConfig {
	return v.value
}

func (v *NullableRemoteSyslogFileConfig) Set(val *RemoteSyslogFileConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoteSyslogFileConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoteSyslogFileConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoteSyslogFileConfig(val *RemoteSyslogFileConfig) *NullableRemoteSyslogFileConfig {
	return &NullableRemoteSyslogFileConfig{value: val, isSet: true}
}

func (v NullableRemoteSyslogFileConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoteSyslogFileConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


