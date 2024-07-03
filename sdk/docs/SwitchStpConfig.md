# SwitchStpConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**SwitchStpConfigType**](SwitchStpConfigType.md) |  | [optional] [default to SWITCHSTPCONFIGTYPE_RSTP]

## Methods

### NewSwitchStpConfig

`func NewSwitchStpConfig() *SwitchStpConfig`

NewSwitchStpConfig instantiates a new SwitchStpConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwitchStpConfigWithDefaults

`func NewSwitchStpConfigWithDefaults() *SwitchStpConfig`

NewSwitchStpConfigWithDefaults instantiates a new SwitchStpConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SwitchStpConfig) GetType() SwitchStpConfigType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SwitchStpConfig) GetTypeOk() (*SwitchStpConfigType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SwitchStpConfig) SetType(v SwitchStpConfigType)`

SetType sets Type field to given value.

### HasType

`func (o *SwitchStpConfig) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


