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

// checks if the ApStatsAutoPlacement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApStatsAutoPlacement{}

// ApStatsAutoPlacement struct for ApStatsAutoPlacement
type ApStatsAutoPlacement struct {
	Id *string `json:"_id,omitempty"`
	Info *ApStatsAutoPlacementInfo `json:"info,omitempty"`
	// Flag to represent if AP is recommended as an anchor by auto placement service
	RecommendedAnchor *bool `json:"recommended_anchor,omitempty"`
	// Basic Placement Status
	Status *string `json:"status,omitempty"`
	// Additional info about placement status
	StatusDetail *string `json:"status_detail,omitempty"`
	// Flag to represent if auto_placement values are currently utilized
	UseAutoPlacement *bool `json:"use_auto_placement,omitempty"`
	// X Autoplaced Position in pixels
	X *float32 `json:"x,omitempty"`
	// X Autoplaced Position in meters
	XM *float32 `json:"x_m,omitempty"`
	// Y Autoplaced Position in pixels
	Y *float32 `json:"y,omitempty"`
	// X Autoplaced Position in meters
	YM *float32 `json:"y_m,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApStatsAutoPlacement ApStatsAutoPlacement

// NewApStatsAutoPlacement instantiates a new ApStatsAutoPlacement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApStatsAutoPlacement() *ApStatsAutoPlacement {
	this := ApStatsAutoPlacement{}
	return &this
}

// NewApStatsAutoPlacementWithDefaults instantiates a new ApStatsAutoPlacement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApStatsAutoPlacementWithDefaults() *ApStatsAutoPlacement {
	this := ApStatsAutoPlacement{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApStatsAutoPlacement) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApStatsAutoPlacement) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApStatsAutoPlacement) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApStatsAutoPlacement) SetId(v string) {
	o.Id = &v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *ApStatsAutoPlacement) GetInfo() ApStatsAutoPlacementInfo {
	if o == nil || IsNil(o.Info) {
		var ret ApStatsAutoPlacementInfo
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApStatsAutoPlacement) GetInfoOk() (*ApStatsAutoPlacementInfo, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *ApStatsAutoPlacement) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given ApStatsAutoPlacementInfo and assigns it to the Info field.
func (o *ApStatsAutoPlacement) SetInfo(v ApStatsAutoPlacementInfo) {
	o.Info = &v
}

// GetRecommendedAnchor returns the RecommendedAnchor field value if set, zero value otherwise.
func (o *ApStatsAutoPlacement) GetRecommendedAnchor() bool {
	if o == nil || IsNil(o.RecommendedAnchor) {
		var ret bool
		return ret
	}
	return *o.RecommendedAnchor
}

// GetRecommendedAnchorOk returns a tuple with the RecommendedAnchor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApStatsAutoPlacement) GetRecommendedAnchorOk() (*bool, bool) {
	if o == nil || IsNil(o.RecommendedAnchor) {
		return nil, false
	}
	return o.RecommendedAnchor, true
}

// HasRecommendedAnchor returns a boolean if a field has been set.
func (o *ApStatsAutoPlacement) HasRecommendedAnchor() bool {
	if o != nil && !IsNil(o.RecommendedAnchor) {
		return true
	}

	return false
}

// SetRecommendedAnchor gets a reference to the given bool and assigns it to the RecommendedAnchor field.
func (o *ApStatsAutoPlacement) SetRecommendedAnchor(v bool) {
	o.RecommendedAnchor = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ApStatsAutoPlacement) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApStatsAutoPlacement) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ApStatsAutoPlacement) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ApStatsAutoPlacement) SetStatus(v string) {
	o.Status = &v
}

// GetStatusDetail returns the StatusDetail field value if set, zero value otherwise.
func (o *ApStatsAutoPlacement) GetStatusDetail() string {
	if o == nil || IsNil(o.StatusDetail) {
		var ret string
		return ret
	}
	return *o.StatusDetail
}

// GetStatusDetailOk returns a tuple with the StatusDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApStatsAutoPlacement) GetStatusDetailOk() (*string, bool) {
	if o == nil || IsNil(o.StatusDetail) {
		return nil, false
	}
	return o.StatusDetail, true
}

// HasStatusDetail returns a boolean if a field has been set.
func (o *ApStatsAutoPlacement) HasStatusDetail() bool {
	if o != nil && !IsNil(o.StatusDetail) {
		return true
	}

	return false
}

// SetStatusDetail gets a reference to the given string and assigns it to the StatusDetail field.
func (o *ApStatsAutoPlacement) SetStatusDetail(v string) {
	o.StatusDetail = &v
}

// GetUseAutoPlacement returns the UseAutoPlacement field value if set, zero value otherwise.
func (o *ApStatsAutoPlacement) GetUseAutoPlacement() bool {
	if o == nil || IsNil(o.UseAutoPlacement) {
		var ret bool
		return ret
	}
	return *o.UseAutoPlacement
}

// GetUseAutoPlacementOk returns a tuple with the UseAutoPlacement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApStatsAutoPlacement) GetUseAutoPlacementOk() (*bool, bool) {
	if o == nil || IsNil(o.UseAutoPlacement) {
		return nil, false
	}
	return o.UseAutoPlacement, true
}

// HasUseAutoPlacement returns a boolean if a field has been set.
func (o *ApStatsAutoPlacement) HasUseAutoPlacement() bool {
	if o != nil && !IsNil(o.UseAutoPlacement) {
		return true
	}

	return false
}

// SetUseAutoPlacement gets a reference to the given bool and assigns it to the UseAutoPlacement field.
func (o *ApStatsAutoPlacement) SetUseAutoPlacement(v bool) {
	o.UseAutoPlacement = &v
}

// GetX returns the X field value if set, zero value otherwise.
func (o *ApStatsAutoPlacement) GetX() float32 {
	if o == nil || IsNil(o.X) {
		var ret float32
		return ret
	}
	return *o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApStatsAutoPlacement) GetXOk() (*float32, bool) {
	if o == nil || IsNil(o.X) {
		return nil, false
	}
	return o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *ApStatsAutoPlacement) HasX() bool {
	if o != nil && !IsNil(o.X) {
		return true
	}

	return false
}

// SetX gets a reference to the given float32 and assigns it to the X field.
func (o *ApStatsAutoPlacement) SetX(v float32) {
	o.X = &v
}

// GetXM returns the XM field value if set, zero value otherwise.
func (o *ApStatsAutoPlacement) GetXM() float32 {
	if o == nil || IsNil(o.XM) {
		var ret float32
		return ret
	}
	return *o.XM
}

// GetXMOk returns a tuple with the XM field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApStatsAutoPlacement) GetXMOk() (*float32, bool) {
	if o == nil || IsNil(o.XM) {
		return nil, false
	}
	return o.XM, true
}

// HasXM returns a boolean if a field has been set.
func (o *ApStatsAutoPlacement) HasXM() bool {
	if o != nil && !IsNil(o.XM) {
		return true
	}

	return false
}

// SetXM gets a reference to the given float32 and assigns it to the XM field.
func (o *ApStatsAutoPlacement) SetXM(v float32) {
	o.XM = &v
}

// GetY returns the Y field value if set, zero value otherwise.
func (o *ApStatsAutoPlacement) GetY() float32 {
	if o == nil || IsNil(o.Y) {
		var ret float32
		return ret
	}
	return *o.Y
}

// GetYOk returns a tuple with the Y field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApStatsAutoPlacement) GetYOk() (*float32, bool) {
	if o == nil || IsNil(o.Y) {
		return nil, false
	}
	return o.Y, true
}

// HasY returns a boolean if a field has been set.
func (o *ApStatsAutoPlacement) HasY() bool {
	if o != nil && !IsNil(o.Y) {
		return true
	}

	return false
}

// SetY gets a reference to the given float32 and assigns it to the Y field.
func (o *ApStatsAutoPlacement) SetY(v float32) {
	o.Y = &v
}

// GetYM returns the YM field value if set, zero value otherwise.
func (o *ApStatsAutoPlacement) GetYM() float32 {
	if o == nil || IsNil(o.YM) {
		var ret float32
		return ret
	}
	return *o.YM
}

// GetYMOk returns a tuple with the YM field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApStatsAutoPlacement) GetYMOk() (*float32, bool) {
	if o == nil || IsNil(o.YM) {
		return nil, false
	}
	return o.YM, true
}

// HasYM returns a boolean if a field has been set.
func (o *ApStatsAutoPlacement) HasYM() bool {
	if o != nil && !IsNil(o.YM) {
		return true
	}

	return false
}

// SetYM gets a reference to the given float32 and assigns it to the YM field.
func (o *ApStatsAutoPlacement) SetYM(v float32) {
	o.YM = &v
}

func (o ApStatsAutoPlacement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApStatsAutoPlacement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["_id"] = o.Id
	}
	if !IsNil(o.Info) {
		toSerialize["info"] = o.Info
	}
	if !IsNil(o.RecommendedAnchor) {
		toSerialize["recommended_anchor"] = o.RecommendedAnchor
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.StatusDetail) {
		toSerialize["status_detail"] = o.StatusDetail
	}
	if !IsNil(o.UseAutoPlacement) {
		toSerialize["use_auto_placement"] = o.UseAutoPlacement
	}
	if !IsNil(o.X) {
		toSerialize["x"] = o.X
	}
	if !IsNil(o.XM) {
		toSerialize["x_m"] = o.XM
	}
	if !IsNil(o.Y) {
		toSerialize["y"] = o.Y
	}
	if !IsNil(o.YM) {
		toSerialize["y_m"] = o.YM
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApStatsAutoPlacement) UnmarshalJSON(data []byte) (err error) {
	varApStatsAutoPlacement := _ApStatsAutoPlacement{}

	err = json.Unmarshal(data, &varApStatsAutoPlacement)

	if err != nil {
		return err
	}

	*o = ApStatsAutoPlacement(varApStatsAutoPlacement)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_id")
		delete(additionalProperties, "info")
		delete(additionalProperties, "recommended_anchor")
		delete(additionalProperties, "status")
		delete(additionalProperties, "status_detail")
		delete(additionalProperties, "use_auto_placement")
		delete(additionalProperties, "x")
		delete(additionalProperties, "x_m")
		delete(additionalProperties, "y")
		delete(additionalProperties, "y_m")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApStatsAutoPlacement struct {
	value *ApStatsAutoPlacement
	isSet bool
}

func (v NullableApStatsAutoPlacement) Get() *ApStatsAutoPlacement {
	return v.value
}

func (v *NullableApStatsAutoPlacement) Set(val *ApStatsAutoPlacement) {
	v.value = val
	v.isSet = true
}

func (v NullableApStatsAutoPlacement) IsSet() bool {
	return v.isSet
}

func (v *NullableApStatsAutoPlacement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApStatsAutoPlacement(val *ApStatsAutoPlacement) *NullableApStatsAutoPlacement {
	return &NullableApStatsAutoPlacement{value: val, isSet: true}
}

func (v NullableApStatsAutoPlacement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApStatsAutoPlacement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


