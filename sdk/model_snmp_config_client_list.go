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
)

// checks if the SnmpConfigClientList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SnmpConfigClientList{}

// SnmpConfigClientList struct for SnmpConfigClientList
type SnmpConfigClientList struct {
	ClientListName *string `json:"client_list_name,omitempty"`
	Clients []string `json:"clients,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SnmpConfigClientList SnmpConfigClientList

// NewSnmpConfigClientList instantiates a new SnmpConfigClientList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnmpConfigClientList() *SnmpConfigClientList {
	this := SnmpConfigClientList{}
	return &this
}

// NewSnmpConfigClientListWithDefaults instantiates a new SnmpConfigClientList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnmpConfigClientListWithDefaults() *SnmpConfigClientList {
	this := SnmpConfigClientList{}
	return &this
}

// GetClientListName returns the ClientListName field value if set, zero value otherwise.
func (o *SnmpConfigClientList) GetClientListName() string {
	if o == nil || IsNil(o.ClientListName) {
		var ret string
		return ret
	}
	return *o.ClientListName
}

// GetClientListNameOk returns a tuple with the ClientListName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpConfigClientList) GetClientListNameOk() (*string, bool) {
	if o == nil || IsNil(o.ClientListName) {
		return nil, false
	}
	return o.ClientListName, true
}

// HasClientListName returns a boolean if a field has been set.
func (o *SnmpConfigClientList) HasClientListName() bool {
	if o != nil && !IsNil(o.ClientListName) {
		return true
	}

	return false
}

// SetClientListName gets a reference to the given string and assigns it to the ClientListName field.
func (o *SnmpConfigClientList) SetClientListName(v string) {
	o.ClientListName = &v
}

// GetClients returns the Clients field value if set, zero value otherwise.
func (o *SnmpConfigClientList) GetClients() []string {
	if o == nil || IsNil(o.Clients) {
		var ret []string
		return ret
	}
	return o.Clients
}

// GetClientsOk returns a tuple with the Clients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpConfigClientList) GetClientsOk() ([]string, bool) {
	if o == nil || IsNil(o.Clients) {
		return nil, false
	}
	return o.Clients, true
}

// HasClients returns a boolean if a field has been set.
func (o *SnmpConfigClientList) HasClients() bool {
	if o != nil && !IsNil(o.Clients) {
		return true
	}

	return false
}

// SetClients gets a reference to the given []string and assigns it to the Clients field.
func (o *SnmpConfigClientList) SetClients(v []string) {
	o.Clients = v
}

func (o SnmpConfigClientList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SnmpConfigClientList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClientListName) {
		toSerialize["client_list_name"] = o.ClientListName
	}
	if !IsNil(o.Clients) {
		toSerialize["clients"] = o.Clients
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SnmpConfigClientList) UnmarshalJSON(data []byte) (err error) {
	varSnmpConfigClientList := _SnmpConfigClientList{}

	err = json.Unmarshal(data, &varSnmpConfigClientList)

	if err != nil {
		return err
	}

	*o = SnmpConfigClientList(varSnmpConfigClientList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "client_list_name")
		delete(additionalProperties, "clients")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSnmpConfigClientList struct {
	value *SnmpConfigClientList
	isSet bool
}

func (v NullableSnmpConfigClientList) Get() *SnmpConfigClientList {
	return v.value
}

func (v *NullableSnmpConfigClientList) Set(val *SnmpConfigClientList) {
	v.value = val
	v.isSet = true
}

func (v NullableSnmpConfigClientList) IsSet() bool {
	return v.isSet
}

func (v *NullableSnmpConfigClientList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnmpConfigClientList(val *SnmpConfigClientList) *NullableSnmpConfigClientList {
	return &NullableSnmpConfigClientList{value: val, isSet: true}
}

func (v NullableSnmpConfigClientList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnmpConfigClientList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

