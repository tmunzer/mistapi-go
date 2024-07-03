# OspfConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Areas** | Pointer to [**map[string]OspfConfigAreasAdditionalProperties**](OspfConfigAreasAdditionalProperties.md) | OSPF areas to run on this device and the corresponding per-area-specific configs. Property key is the area | [optional] 
**Enabled** | Pointer to **bool** | whether to rung OSPF on this device | [optional] 
**ReferenceBandwidth** | Pointer to **string** | Bandwidth for calculating metric defaults (9600..4000000000000) | [optional] [default to "100M"]

## Methods

### NewOspfConfig

`func NewOspfConfig() *OspfConfig`

NewOspfConfig instantiates a new OspfConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOspfConfigWithDefaults

`func NewOspfConfigWithDefaults() *OspfConfig`

NewOspfConfigWithDefaults instantiates a new OspfConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAreas

`func (o *OspfConfig) GetAreas() map[string]OspfConfigAreasAdditionalProperties`

GetAreas returns the Areas field if non-nil, zero value otherwise.

### GetAreasOk

`func (o *OspfConfig) GetAreasOk() (*map[string]OspfConfigAreasAdditionalProperties, bool)`

GetAreasOk returns a tuple with the Areas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreas

`func (o *OspfConfig) SetAreas(v map[string]OspfConfigAreasAdditionalProperties)`

SetAreas sets Areas field to given value.

### HasAreas

`func (o *OspfConfig) HasAreas() bool`

HasAreas returns a boolean if a field has been set.

### GetEnabled

`func (o *OspfConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OspfConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OspfConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *OspfConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetReferenceBandwidth

`func (o *OspfConfig) GetReferenceBandwidth() string`

GetReferenceBandwidth returns the ReferenceBandwidth field if non-nil, zero value otherwise.

### GetReferenceBandwidthOk

`func (o *OspfConfig) GetReferenceBandwidthOk() (*string, bool)`

GetReferenceBandwidthOk returns a tuple with the ReferenceBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceBandwidth

`func (o *OspfConfig) SetReferenceBandwidth(v string)`

SetReferenceBandwidth sets ReferenceBandwidth field to given value.

### HasReferenceBandwidth

`func (o *OspfConfig) HasReferenceBandwidth() bool`

HasReferenceBandwidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


