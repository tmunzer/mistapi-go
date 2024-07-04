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

// checks if the MspOrgChange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MspOrgChange{}

// MspOrgChange struct for MspOrgChange
type MspOrgChange struct {
	Op MspOrgChangeOperation `json:"op"`
	// list of org_id
	OrgIds []string `json:"org_ids"`
	AdditionalProperties map[string]interface{}
}

type _MspOrgChange MspOrgChange

// NewMspOrgChange instantiates a new MspOrgChange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMspOrgChange(op MspOrgChangeOperation, orgIds []string) *MspOrgChange {
	this := MspOrgChange{}
	this.Op = op
	this.OrgIds = orgIds
	return &this
}

// NewMspOrgChangeWithDefaults instantiates a new MspOrgChange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMspOrgChangeWithDefaults() *MspOrgChange {
	this := MspOrgChange{}
	return &this
}

// GetOp returns the Op field value
func (o *MspOrgChange) GetOp() MspOrgChangeOperation {
	if o == nil {
		var ret MspOrgChangeOperation
		return ret
	}

	return o.Op
}

// GetOpOk returns a tuple with the Op field value
// and a boolean to check if the value has been set.
func (o *MspOrgChange) GetOpOk() (*MspOrgChangeOperation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Op, true
}

// SetOp sets field value
func (o *MspOrgChange) SetOp(v MspOrgChangeOperation) {
	o.Op = v
}

// GetOrgIds returns the OrgIds field value
func (o *MspOrgChange) GetOrgIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.OrgIds
}

// GetOrgIdsOk returns a tuple with the OrgIds field value
// and a boolean to check if the value has been set.
func (o *MspOrgChange) GetOrgIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrgIds, true
}

// SetOrgIds sets field value
func (o *MspOrgChange) SetOrgIds(v []string) {
	o.OrgIds = v
}

func (o MspOrgChange) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MspOrgChange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["op"] = o.Op
	toSerialize["org_ids"] = o.OrgIds

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MspOrgChange) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"op",
		"org_ids",
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

	varMspOrgChange := _MspOrgChange{}

	err = json.Unmarshal(data, &varMspOrgChange)

	if err != nil {
		return err
	}

	*o = MspOrgChange(varMspOrgChange)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "op")
		delete(additionalProperties, "org_ids")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMspOrgChange struct {
	value *MspOrgChange
	isSet bool
}

func (v NullableMspOrgChange) Get() *MspOrgChange {
	return v.value
}

func (v *NullableMspOrgChange) Set(val *MspOrgChange) {
	v.value = val
	v.isSet = true
}

func (v NullableMspOrgChange) IsSet() bool {
	return v.isSet
}

func (v *NullableMspOrgChange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMspOrgChange(val *MspOrgChange) *NullableMspOrgChange {
	return &NullableMspOrgChange{value: val, isSet: true}
}

func (v NullableMspOrgChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMspOrgChange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


