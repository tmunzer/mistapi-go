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
	"fmt"
)

// checks if the Template type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Template{}

// Template Template
type Template struct {
	Applies *TemplateApplies `json:"applies,omitempty"`
	CreatedTime *float32 `json:"created_time,omitempty"`
	// list of Device Profile ids
	DeviceprofileIds []string `json:"deviceprofile_ids,omitempty"`
	Exceptions *TemplateExceptions `json:"exceptions,omitempty"`
	// whether to further filter by Device Profile
	FilterByDeviceprofile *bool `json:"filter_by_deviceprofile,omitempty"`
	Id *string `json:"id,omitempty"`
	ModifiedTime *float32 `json:"modified_time,omitempty"`
	Name string `json:"name"`
	OrgId *string `json:"org_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Template Template

// NewTemplate instantiates a new Template object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplate(name string) *Template {
	this := Template{}
	this.Name = name
	return &this
}

// NewTemplateWithDefaults instantiates a new Template object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateWithDefaults() *Template {
	this := Template{}
	return &this
}

// GetApplies returns the Applies field value if set, zero value otherwise.
func (o *Template) GetApplies() TemplateApplies {
	if o == nil || IsNil(o.Applies) {
		var ret TemplateApplies
		return ret
	}
	return *o.Applies
}

// GetAppliesOk returns a tuple with the Applies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Template) GetAppliesOk() (*TemplateApplies, bool) {
	if o == nil || IsNil(o.Applies) {
		return nil, false
	}
	return o.Applies, true
}

// HasApplies returns a boolean if a field has been set.
func (o *Template) HasApplies() bool {
	if o != nil && !IsNil(o.Applies) {
		return true
	}

	return false
}

// SetApplies gets a reference to the given TemplateApplies and assigns it to the Applies field.
func (o *Template) SetApplies(v TemplateApplies) {
	o.Applies = &v
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *Template) GetCreatedTime() float32 {
	if o == nil || IsNil(o.CreatedTime) {
		var ret float32
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Template) GetCreatedTimeOk() (*float32, bool) {
	if o == nil || IsNil(o.CreatedTime) {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *Template) HasCreatedTime() bool {
	if o != nil && !IsNil(o.CreatedTime) {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given float32 and assigns it to the CreatedTime field.
func (o *Template) SetCreatedTime(v float32) {
	o.CreatedTime = &v
}

// GetDeviceprofileIds returns the DeviceprofileIds field value if set, zero value otherwise.
func (o *Template) GetDeviceprofileIds() []string {
	if o == nil || IsNil(o.DeviceprofileIds) {
		var ret []string
		return ret
	}
	return o.DeviceprofileIds
}

// GetDeviceprofileIdsOk returns a tuple with the DeviceprofileIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Template) GetDeviceprofileIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.DeviceprofileIds) {
		return nil, false
	}
	return o.DeviceprofileIds, true
}

// HasDeviceprofileIds returns a boolean if a field has been set.
func (o *Template) HasDeviceprofileIds() bool {
	if o != nil && !IsNil(o.DeviceprofileIds) {
		return true
	}

	return false
}

// SetDeviceprofileIds gets a reference to the given []string and assigns it to the DeviceprofileIds field.
func (o *Template) SetDeviceprofileIds(v []string) {
	o.DeviceprofileIds = v
}

// GetExceptions returns the Exceptions field value if set, zero value otherwise.
func (o *Template) GetExceptions() TemplateExceptions {
	if o == nil || IsNil(o.Exceptions) {
		var ret TemplateExceptions
		return ret
	}
	return *o.Exceptions
}

// GetExceptionsOk returns a tuple with the Exceptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Template) GetExceptionsOk() (*TemplateExceptions, bool) {
	if o == nil || IsNil(o.Exceptions) {
		return nil, false
	}
	return o.Exceptions, true
}

// HasExceptions returns a boolean if a field has been set.
func (o *Template) HasExceptions() bool {
	if o != nil && !IsNil(o.Exceptions) {
		return true
	}

	return false
}

// SetExceptions gets a reference to the given TemplateExceptions and assigns it to the Exceptions field.
func (o *Template) SetExceptions(v TemplateExceptions) {
	o.Exceptions = &v
}

// GetFilterByDeviceprofile returns the FilterByDeviceprofile field value if set, zero value otherwise.
func (o *Template) GetFilterByDeviceprofile() bool {
	if o == nil || IsNil(o.FilterByDeviceprofile) {
		var ret bool
		return ret
	}
	return *o.FilterByDeviceprofile
}

// GetFilterByDeviceprofileOk returns a tuple with the FilterByDeviceprofile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Template) GetFilterByDeviceprofileOk() (*bool, bool) {
	if o == nil || IsNil(o.FilterByDeviceprofile) {
		return nil, false
	}
	return o.FilterByDeviceprofile, true
}

// HasFilterByDeviceprofile returns a boolean if a field has been set.
func (o *Template) HasFilterByDeviceprofile() bool {
	if o != nil && !IsNil(o.FilterByDeviceprofile) {
		return true
	}

	return false
}

// SetFilterByDeviceprofile gets a reference to the given bool and assigns it to the FilterByDeviceprofile field.
func (o *Template) SetFilterByDeviceprofile(v bool) {
	o.FilterByDeviceprofile = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Template) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Template) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Template) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Template) SetId(v string) {
	o.Id = &v
}

// GetModifiedTime returns the ModifiedTime field value if set, zero value otherwise.
func (o *Template) GetModifiedTime() float32 {
	if o == nil || IsNil(o.ModifiedTime) {
		var ret float32
		return ret
	}
	return *o.ModifiedTime
}

// GetModifiedTimeOk returns a tuple with the ModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Template) GetModifiedTimeOk() (*float32, bool) {
	if o == nil || IsNil(o.ModifiedTime) {
		return nil, false
	}
	return o.ModifiedTime, true
}

// HasModifiedTime returns a boolean if a field has been set.
func (o *Template) HasModifiedTime() bool {
	if o != nil && !IsNil(o.ModifiedTime) {
		return true
	}

	return false
}

// SetModifiedTime gets a reference to the given float32 and assigns it to the ModifiedTime field.
func (o *Template) SetModifiedTime(v float32) {
	o.ModifiedTime = &v
}

// GetName returns the Name field value
func (o *Template) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Template) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Template) SetName(v string) {
	o.Name = v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *Template) GetOrgId() string {
	if o == nil || IsNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Template) GetOrgIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *Template) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *Template) SetOrgId(v string) {
	o.OrgId = &v
}

func (o Template) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Template) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Applies) {
		toSerialize["applies"] = o.Applies
	}
	if !IsNil(o.CreatedTime) {
		toSerialize["created_time"] = o.CreatedTime
	}
	if !IsNil(o.DeviceprofileIds) {
		toSerialize["deviceprofile_ids"] = o.DeviceprofileIds
	}
	if !IsNil(o.Exceptions) {
		toSerialize["exceptions"] = o.Exceptions
	}
	if !IsNil(o.FilterByDeviceprofile) {
		toSerialize["filter_by_deviceprofile"] = o.FilterByDeviceprofile
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ModifiedTime) {
		toSerialize["modified_time"] = o.ModifiedTime
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.OrgId) {
		toSerialize["org_id"] = o.OrgId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Template) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varTemplate := _Template{}

	err = json.Unmarshal(data, &varTemplate)

	if err != nil {
		return err
	}

	*o = Template(varTemplate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "applies")
		delete(additionalProperties, "created_time")
		delete(additionalProperties, "deviceprofile_ids")
		delete(additionalProperties, "exceptions")
		delete(additionalProperties, "filter_by_deviceprofile")
		delete(additionalProperties, "id")
		delete(additionalProperties, "modified_time")
		delete(additionalProperties, "name")
		delete(additionalProperties, "org_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTemplate struct {
	value *Template
	isSet bool
}

func (v NullableTemplate) Get() *Template {
	return v.value
}

func (v *NullableTemplate) Set(val *Template) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplate(val *Template) *NullableTemplate {
	return &NullableTemplate{value: val, isSet: true}
}

func (v NullableTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


