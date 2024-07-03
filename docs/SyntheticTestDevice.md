# SyntheticTestDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | Pointer to **string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;dns&#x60; | [optional] 
**Ip** | Pointer to **string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;arp&#x60; | [optional] 
**Password** | Pointer to **string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;radius&#x60; | [optional] 
**PortId** | Pointer to **string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;ssr&#x60; | [optional] 
**Type** | [**SyntheticTestType**](SyntheticTestType.md) |  | 
**Url** | Pointer to **string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;curl&#x60; | [optional] 
**Username** | Pointer to **string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;radius&#x60; | [optional] 
**VlanId** | Pointer to **int32** | required for AP | [optional] 

## Methods

### NewSyntheticTestDevice

`func NewSyntheticTestDevice(type_ SyntheticTestType, ) *SyntheticTestDevice`

NewSyntheticTestDevice instantiates a new SyntheticTestDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticTestDeviceWithDefaults

`func NewSyntheticTestDeviceWithDefaults() *SyntheticTestDevice`

NewSyntheticTestDeviceWithDefaults instantiates a new SyntheticTestDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *SyntheticTestDevice) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *SyntheticTestDevice) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *SyntheticTestDevice) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *SyntheticTestDevice) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetIp

`func (o *SyntheticTestDevice) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *SyntheticTestDevice) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *SyntheticTestDevice) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *SyntheticTestDevice) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetPassword

`func (o *SyntheticTestDevice) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SyntheticTestDevice) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SyntheticTestDevice) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *SyntheticTestDevice) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPortId

`func (o *SyntheticTestDevice) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *SyntheticTestDevice) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *SyntheticTestDevice) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *SyntheticTestDevice) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetType

`func (o *SyntheticTestDevice) GetType() SyntheticTestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SyntheticTestDevice) GetTypeOk() (*SyntheticTestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SyntheticTestDevice) SetType(v SyntheticTestType)`

SetType sets Type field to given value.


### GetUrl

`func (o *SyntheticTestDevice) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SyntheticTestDevice) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SyntheticTestDevice) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SyntheticTestDevice) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUsername

`func (o *SyntheticTestDevice) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SyntheticTestDevice) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SyntheticTestDevice) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *SyntheticTestDevice) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetVlanId

`func (o *SyntheticTestDevice) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *SyntheticTestDevice) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *SyntheticTestDevice) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *SyntheticTestDevice) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


