# WlanDatarates

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ht** | Pointer to **NullableString** | MCS bitmasks for 4 streams (16-bit for each stream, MCS0 is least significant bit), e.g. 00ff 00f0 001f limits HT rates to MCS 0-7 for 1 stream, MCS 4-7 for 2 stream (i.e. MCS 12-15), MCS 1-5 for 3 stream (i.e. MCS 16-20) | [optional] 
**Legacy** | Pointer to [**[]WlanDataratesLegacyItem**](WlanDataratesLegacyItem.md) | list of supported rates (IE&#x3D;1) and extended supported rates (IE&#x3D;50) for custom template, append ‘b’ at the end to indicate a rate being basic/mandatory. If &#x60;template&#x60;&#x3D;&#x3D;&#x60;custom&#x60; is configured and legacy does not define at least one basic rate, it will use &#x60;no-legacy&#x60; default values | [optional] 
**MinRssi** | Pointer to **int32** | Minimum RSSI for client to connect, 0 means not enforcing | [optional] 
**Template** | Pointer to **NullableString** | * &#x60;no-legacy&#x60;: no 11b * &#x60;compatible&#x60;: all, like before, default setting that Broadcom/Atheros used * &#x60;legacy-only&#x60;: disable 802.11n and 802.11ac  * &#x60;high-density&#x60;: no 11b, no low rates * &#x60;custom&#x60;: user defined | [optional] 
**Vht** | Pointer to **string** | MCS bitmasks for 4 streams (16-bit for each stream, MCS0 is least significant bit), e.g. 03ff 01ff 00ff limits VHT rates to MCS 0-9 for 1 stream, MCS 0-8 for 2 streams, and MCS 0-7 for 3 streams. | [optional] 

## Methods

### NewWlanDatarates

`func NewWlanDatarates() *WlanDatarates`

NewWlanDatarates instantiates a new WlanDatarates object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWlanDataratesWithDefaults

`func NewWlanDataratesWithDefaults() *WlanDatarates`

NewWlanDataratesWithDefaults instantiates a new WlanDatarates object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHt

`func (o *WlanDatarates) GetHt() string`

GetHt returns the Ht field if non-nil, zero value otherwise.

### GetHtOk

`func (o *WlanDatarates) GetHtOk() (*string, bool)`

GetHtOk returns a tuple with the Ht field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHt

`func (o *WlanDatarates) SetHt(v string)`

SetHt sets Ht field to given value.

### HasHt

`func (o *WlanDatarates) HasHt() bool`

HasHt returns a boolean if a field has been set.

### SetHtNil

`func (o *WlanDatarates) SetHtNil(b bool)`

 SetHtNil sets the value for Ht to be an explicit nil

### UnsetHt
`func (o *WlanDatarates) UnsetHt()`

UnsetHt ensures that no value is present for Ht, not even an explicit nil
### GetLegacy

`func (o *WlanDatarates) GetLegacy() []WlanDataratesLegacyItem`

GetLegacy returns the Legacy field if non-nil, zero value otherwise.

### GetLegacyOk

`func (o *WlanDatarates) GetLegacyOk() (*[]WlanDataratesLegacyItem, bool)`

GetLegacyOk returns a tuple with the Legacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacy

`func (o *WlanDatarates) SetLegacy(v []WlanDataratesLegacyItem)`

SetLegacy sets Legacy field to given value.

### HasLegacy

`func (o *WlanDatarates) HasLegacy() bool`

HasLegacy returns a boolean if a field has been set.

### GetMinRssi

`func (o *WlanDatarates) GetMinRssi() int32`

GetMinRssi returns the MinRssi field if non-nil, zero value otherwise.

### GetMinRssiOk

`func (o *WlanDatarates) GetMinRssiOk() (*int32, bool)`

GetMinRssiOk returns a tuple with the MinRssi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRssi

`func (o *WlanDatarates) SetMinRssi(v int32)`

SetMinRssi sets MinRssi field to given value.

### HasMinRssi

`func (o *WlanDatarates) HasMinRssi() bool`

HasMinRssi returns a boolean if a field has been set.

### GetTemplate

`func (o *WlanDatarates) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *WlanDatarates) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *WlanDatarates) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *WlanDatarates) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### SetTemplateNil

`func (o *WlanDatarates) SetTemplateNil(b bool)`

 SetTemplateNil sets the value for Template to be an explicit nil

### UnsetTemplate
`func (o *WlanDatarates) UnsetTemplate()`

UnsetTemplate ensures that no value is present for Template, not even an explicit nil
### GetVht

`func (o *WlanDatarates) GetVht() string`

GetVht returns the Vht field if non-nil, zero value otherwise.

### GetVhtOk

`func (o *WlanDatarates) GetVhtOk() (*string, bool)`

GetVhtOk returns a tuple with the Vht field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVht

`func (o *WlanDatarates) SetVht(v string)`

SetVht sets Vht field to given value.

### HasVht

`func (o *WlanDatarates) HasVht() bool`

HasVht returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


