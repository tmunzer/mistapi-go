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

// checks if the PcapBucketVerify type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PcapBucketVerify{}

// PcapBucketVerify struct for PcapBucketVerify
type PcapBucketVerify struct {
	Bucket string `json:"bucket"`
	VerifyToken string `json:"verify_token"`
	AdditionalProperties map[string]interface{}
}

type _PcapBucketVerify PcapBucketVerify

// NewPcapBucketVerify instantiates a new PcapBucketVerify object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPcapBucketVerify(bucket string, verifyToken string) *PcapBucketVerify {
	this := PcapBucketVerify{}
	this.Bucket = bucket
	this.VerifyToken = verifyToken
	return &this
}

// NewPcapBucketVerifyWithDefaults instantiates a new PcapBucketVerify object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPcapBucketVerifyWithDefaults() *PcapBucketVerify {
	this := PcapBucketVerify{}
	return &this
}

// GetBucket returns the Bucket field value
func (o *PcapBucketVerify) GetBucket() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value
// and a boolean to check if the value has been set.
func (o *PcapBucketVerify) GetBucketOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bucket, true
}

// SetBucket sets field value
func (o *PcapBucketVerify) SetBucket(v string) {
	o.Bucket = v
}

// GetVerifyToken returns the VerifyToken field value
func (o *PcapBucketVerify) GetVerifyToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerifyToken
}

// GetVerifyTokenOk returns a tuple with the VerifyToken field value
// and a boolean to check if the value has been set.
func (o *PcapBucketVerify) GetVerifyTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerifyToken, true
}

// SetVerifyToken sets field value
func (o *PcapBucketVerify) SetVerifyToken(v string) {
	o.VerifyToken = v
}

func (o PcapBucketVerify) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PcapBucketVerify) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bucket"] = o.Bucket
	toSerialize["verify_token"] = o.VerifyToken

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PcapBucketVerify) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"bucket",
		"verify_token",
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

	varPcapBucketVerify := _PcapBucketVerify{}

	err = json.Unmarshal(data, &varPcapBucketVerify)

	if err != nil {
		return err
	}

	*o = PcapBucketVerify(varPcapBucketVerify)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "bucket")
		delete(additionalProperties, "verify_token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePcapBucketVerify struct {
	value *PcapBucketVerify
	isSet bool
}

func (v NullablePcapBucketVerify) Get() *PcapBucketVerify {
	return v.value
}

func (v *NullablePcapBucketVerify) Set(val *PcapBucketVerify) {
	v.value = val
	v.isSet = true
}

func (v NullablePcapBucketVerify) IsSet() bool {
	return v.isSet
}

func (v *NullablePcapBucketVerify) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePcapBucketVerify(val *PcapBucketVerify) *NullablePcapBucketVerify {
	return &NullablePcapBucketVerify{value: val, isSet: true}
}

func (v NullablePcapBucketVerify) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePcapBucketVerify) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


