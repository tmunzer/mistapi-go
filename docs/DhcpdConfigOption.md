# DhcpdConfigOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**DhcpdConfigOptionType**](DhcpdConfigOptionType.md) |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 

## Methods

### NewDhcpdConfigOption

`func NewDhcpdConfigOption() *DhcpdConfigOption`

NewDhcpdConfigOption instantiates a new DhcpdConfigOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpdConfigOptionWithDefaults

`func NewDhcpdConfigOptionWithDefaults() *DhcpdConfigOption`

NewDhcpdConfigOptionWithDefaults instantiates a new DhcpdConfigOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DhcpdConfigOption) GetType() DhcpdConfigOptionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DhcpdConfigOption) GetTypeOk() (*DhcpdConfigOptionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DhcpdConfigOption) SetType(v DhcpdConfigOptionType)`

SetType sets Type field to given value.

### HasType

`func (o *DhcpdConfigOption) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *DhcpdConfigOption) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DhcpdConfigOption) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DhcpdConfigOption) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *DhcpdConfigOption) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


