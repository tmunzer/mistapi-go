# OtherDeviceUpdateMulti

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Macs** | Pointer to **[]string** | The mac address of the peer device. | [optional] 
**Op** | [**OtherDeviceUpdateOperation**](OtherDeviceUpdateOperation.md) |  | 
**SiteId** | Pointer to **string** |  | [optional] 

## Methods

### NewOtherDeviceUpdateMulti

`func NewOtherDeviceUpdateMulti(op OtherDeviceUpdateOperation, ) *OtherDeviceUpdateMulti`

NewOtherDeviceUpdateMulti instantiates a new OtherDeviceUpdateMulti object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOtherDeviceUpdateMultiWithDefaults

`func NewOtherDeviceUpdateMultiWithDefaults() *OtherDeviceUpdateMulti`

NewOtherDeviceUpdateMultiWithDefaults instantiates a new OtherDeviceUpdateMulti object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMacs

`func (o *OtherDeviceUpdateMulti) GetMacs() []string`

GetMacs returns the Macs field if non-nil, zero value otherwise.

### GetMacsOk

`func (o *OtherDeviceUpdateMulti) GetMacsOk() (*[]string, bool)`

GetMacsOk returns a tuple with the Macs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacs

`func (o *OtherDeviceUpdateMulti) SetMacs(v []string)`

SetMacs sets Macs field to given value.

### HasMacs

`func (o *OtherDeviceUpdateMulti) HasMacs() bool`

HasMacs returns a boolean if a field has been set.

### GetOp

`func (o *OtherDeviceUpdateMulti) GetOp() OtherDeviceUpdateOperation`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *OtherDeviceUpdateMulti) GetOpOk() (*OtherDeviceUpdateOperation, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *OtherDeviceUpdateMulti) SetOp(v OtherDeviceUpdateOperation)`

SetOp sets Op field to given value.


### GetSiteId

`func (o *OtherDeviceUpdateMulti) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *OtherDeviceUpdateMulti) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *OtherDeviceUpdateMulti) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *OtherDeviceUpdateMulti) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


