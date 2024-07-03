# WlanAppQosAppsProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dscp** | Pointer to **int32** |  | [optional] 
**DstSubnet** | Pointer to **string** | subnet filter is not required but helps AP to only inspect certain traffic (thus reducing AP load) | [optional] 
**SrcSubnet** | Pointer to **string** | subnet filter is not required but helps AP to only inspect certain traffic (thus reducing AP load) | [optional] 

## Methods

### NewWlanAppQosAppsProperties

`func NewWlanAppQosAppsProperties() *WlanAppQosAppsProperties`

NewWlanAppQosAppsProperties instantiates a new WlanAppQosAppsProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWlanAppQosAppsPropertiesWithDefaults

`func NewWlanAppQosAppsPropertiesWithDefaults() *WlanAppQosAppsProperties`

NewWlanAppQosAppsPropertiesWithDefaults instantiates a new WlanAppQosAppsProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDscp

`func (o *WlanAppQosAppsProperties) GetDscp() int32`

GetDscp returns the Dscp field if non-nil, zero value otherwise.

### GetDscpOk

`func (o *WlanAppQosAppsProperties) GetDscpOk() (*int32, bool)`

GetDscpOk returns a tuple with the Dscp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDscp

`func (o *WlanAppQosAppsProperties) SetDscp(v int32)`

SetDscp sets Dscp field to given value.

### HasDscp

`func (o *WlanAppQosAppsProperties) HasDscp() bool`

HasDscp returns a boolean if a field has been set.

### GetDstSubnet

`func (o *WlanAppQosAppsProperties) GetDstSubnet() string`

GetDstSubnet returns the DstSubnet field if non-nil, zero value otherwise.

### GetDstSubnetOk

`func (o *WlanAppQosAppsProperties) GetDstSubnetOk() (*string, bool)`

GetDstSubnetOk returns a tuple with the DstSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstSubnet

`func (o *WlanAppQosAppsProperties) SetDstSubnet(v string)`

SetDstSubnet sets DstSubnet field to given value.

### HasDstSubnet

`func (o *WlanAppQosAppsProperties) HasDstSubnet() bool`

HasDstSubnet returns a boolean if a field has been set.

### GetSrcSubnet

`func (o *WlanAppQosAppsProperties) GetSrcSubnet() string`

GetSrcSubnet returns the SrcSubnet field if non-nil, zero value otherwise.

### GetSrcSubnetOk

`func (o *WlanAppQosAppsProperties) GetSrcSubnetOk() (*string, bool)`

GetSrcSubnetOk returns a tuple with the SrcSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcSubnet

`func (o *WlanAppQosAppsProperties) SetSrcSubnet(v string)`

SetSrcSubnet sets SrcSubnet field to given value.

### HasSrcSubnet

`func (o *WlanAppQosAppsProperties) HasSrcSubnet() bool`

HasSrcSubnet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


