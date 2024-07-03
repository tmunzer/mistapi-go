# UtilsResetRadioConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bands** | **[]string** | list of bands | 
**Force** | Pointer to **bool** | whether to reset those with radio disabled. default is false (i.e. if user intentionally disables a radio, honor it) | [optional] [default to false]

## Methods

### NewUtilsResetRadioConfig

`func NewUtilsResetRadioConfig(bands []string, ) *UtilsResetRadioConfig`

NewUtilsResetRadioConfig instantiates a new UtilsResetRadioConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUtilsResetRadioConfigWithDefaults

`func NewUtilsResetRadioConfigWithDefaults() *UtilsResetRadioConfig`

NewUtilsResetRadioConfigWithDefaults instantiates a new UtilsResetRadioConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBands

`func (o *UtilsResetRadioConfig) GetBands() []string`

GetBands returns the Bands field if non-nil, zero value otherwise.

### GetBandsOk

`func (o *UtilsResetRadioConfig) GetBandsOk() (*[]string, bool)`

GetBandsOk returns a tuple with the Bands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBands

`func (o *UtilsResetRadioConfig) SetBands(v []string)`

SetBands sets Bands field to given value.


### GetForce

`func (o *UtilsResetRadioConfig) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *UtilsResetRadioConfig) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *UtilsResetRadioConfig) SetForce(v bool)`

SetForce sets Force field to given value.

### HasForce

`func (o *UtilsResetRadioConfig) HasForce() bool`

HasForce returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


