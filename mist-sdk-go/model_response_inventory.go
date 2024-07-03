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

// checks if the ResponseInventory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseInventory{}

// ResponseInventory struct for ResponseInventory
type ResponseInventory struct {
	Added []string `json:"added,omitempty"`
	Duplicated []string `json:"duplicated,omitempty"`
	Error []string `json:"error,omitempty"`
	InventoryAdded []ResponseInventoryInventoryAddedItems `json:"inventory_added,omitempty"`
	InventoryDuplicated []ResponseInventoryInventoryDuplicatedItems `json:"inventory_duplicated,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResponseInventory ResponseInventory

// NewResponseInventory instantiates a new ResponseInventory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseInventory() *ResponseInventory {
	this := ResponseInventory{}
	return &this
}

// NewResponseInventoryWithDefaults instantiates a new ResponseInventory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseInventoryWithDefaults() *ResponseInventory {
	this := ResponseInventory{}
	return &this
}

// GetAdded returns the Added field value if set, zero value otherwise.
func (o *ResponseInventory) GetAdded() []string {
	if o == nil || IsNil(o.Added) {
		var ret []string
		return ret
	}
	return o.Added
}

// GetAddedOk returns a tuple with the Added field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseInventory) GetAddedOk() ([]string, bool) {
	if o == nil || IsNil(o.Added) {
		return nil, false
	}
	return o.Added, true
}

// HasAdded returns a boolean if a field has been set.
func (o *ResponseInventory) HasAdded() bool {
	if o != nil && !IsNil(o.Added) {
		return true
	}

	return false
}

// SetAdded gets a reference to the given []string and assigns it to the Added field.
func (o *ResponseInventory) SetAdded(v []string) {
	o.Added = v
}

// GetDuplicated returns the Duplicated field value if set, zero value otherwise.
func (o *ResponseInventory) GetDuplicated() []string {
	if o == nil || IsNil(o.Duplicated) {
		var ret []string
		return ret
	}
	return o.Duplicated
}

// GetDuplicatedOk returns a tuple with the Duplicated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseInventory) GetDuplicatedOk() ([]string, bool) {
	if o == nil || IsNil(o.Duplicated) {
		return nil, false
	}
	return o.Duplicated, true
}

// HasDuplicated returns a boolean if a field has been set.
func (o *ResponseInventory) HasDuplicated() bool {
	if o != nil && !IsNil(o.Duplicated) {
		return true
	}

	return false
}

// SetDuplicated gets a reference to the given []string and assigns it to the Duplicated field.
func (o *ResponseInventory) SetDuplicated(v []string) {
	o.Duplicated = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ResponseInventory) GetError() []string {
	if o == nil || IsNil(o.Error) {
		var ret []string
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseInventory) GetErrorOk() ([]string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ResponseInventory) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given []string and assigns it to the Error field.
func (o *ResponseInventory) SetError(v []string) {
	o.Error = v
}

// GetInventoryAdded returns the InventoryAdded field value if set, zero value otherwise.
func (o *ResponseInventory) GetInventoryAdded() []ResponseInventoryInventoryAddedItems {
	if o == nil || IsNil(o.InventoryAdded) {
		var ret []ResponseInventoryInventoryAddedItems
		return ret
	}
	return o.InventoryAdded
}

// GetInventoryAddedOk returns a tuple with the InventoryAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseInventory) GetInventoryAddedOk() ([]ResponseInventoryInventoryAddedItems, bool) {
	if o == nil || IsNil(o.InventoryAdded) {
		return nil, false
	}
	return o.InventoryAdded, true
}

// HasInventoryAdded returns a boolean if a field has been set.
func (o *ResponseInventory) HasInventoryAdded() bool {
	if o != nil && !IsNil(o.InventoryAdded) {
		return true
	}

	return false
}

// SetInventoryAdded gets a reference to the given []ResponseInventoryInventoryAddedItems and assigns it to the InventoryAdded field.
func (o *ResponseInventory) SetInventoryAdded(v []ResponseInventoryInventoryAddedItems) {
	o.InventoryAdded = v
}

// GetInventoryDuplicated returns the InventoryDuplicated field value if set, zero value otherwise.
func (o *ResponseInventory) GetInventoryDuplicated() []ResponseInventoryInventoryDuplicatedItems {
	if o == nil || IsNil(o.InventoryDuplicated) {
		var ret []ResponseInventoryInventoryDuplicatedItems
		return ret
	}
	return o.InventoryDuplicated
}

// GetInventoryDuplicatedOk returns a tuple with the InventoryDuplicated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseInventory) GetInventoryDuplicatedOk() ([]ResponseInventoryInventoryDuplicatedItems, bool) {
	if o == nil || IsNil(o.InventoryDuplicated) {
		return nil, false
	}
	return o.InventoryDuplicated, true
}

// HasInventoryDuplicated returns a boolean if a field has been set.
func (o *ResponseInventory) HasInventoryDuplicated() bool {
	if o != nil && !IsNil(o.InventoryDuplicated) {
		return true
	}

	return false
}

// SetInventoryDuplicated gets a reference to the given []ResponseInventoryInventoryDuplicatedItems and assigns it to the InventoryDuplicated field.
func (o *ResponseInventory) SetInventoryDuplicated(v []ResponseInventoryInventoryDuplicatedItems) {
	o.InventoryDuplicated = v
}

func (o ResponseInventory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseInventory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Added) {
		toSerialize["added"] = o.Added
	}
	if !IsNil(o.Duplicated) {
		toSerialize["duplicated"] = o.Duplicated
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.InventoryAdded) {
		toSerialize["inventory_added"] = o.InventoryAdded
	}
	if !IsNil(o.InventoryDuplicated) {
		toSerialize["inventory_duplicated"] = o.InventoryDuplicated
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResponseInventory) UnmarshalJSON(data []byte) (err error) {
	varResponseInventory := _ResponseInventory{}

	err = json.Unmarshal(data, &varResponseInventory)

	if err != nil {
		return err
	}

	*o = ResponseInventory(varResponseInventory)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "added")
		delete(additionalProperties, "duplicated")
		delete(additionalProperties, "error")
		delete(additionalProperties, "inventory_added")
		delete(additionalProperties, "inventory_duplicated")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResponseInventory struct {
	value *ResponseInventory
	isSet bool
}

func (v NullableResponseInventory) Get() *ResponseInventory {
	return v.value
}

func (v *NullableResponseInventory) Set(val *ResponseInventory) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseInventory) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseInventory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseInventory(val *ResponseInventory) *NullableResponseInventory {
	return &NullableResponseInventory{value: val, isSet: true}
}

func (v NullableResponseInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseInventory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

