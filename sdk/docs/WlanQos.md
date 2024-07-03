# WlanQos

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Class** | Pointer to [**WlanQosClass**](WlanQosClass.md) |  | [optional] [default to WLANQOSCLASS_BEST_EFFORT]
**Overwrite** | Pointer to **bool** | whether to overwrite QoS | [optional] [default to false]

## Methods

### NewWlanQos

`func NewWlanQos() *WlanQos`

NewWlanQos instantiates a new WlanQos object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWlanQosWithDefaults

`func NewWlanQosWithDefaults() *WlanQos`

NewWlanQosWithDefaults instantiates a new WlanQos object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClass

`func (o *WlanQos) GetClass() WlanQosClass`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *WlanQos) GetClassOk() (*WlanQosClass, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *WlanQos) SetClass(v WlanQosClass)`

SetClass sets Class field to given value.

### HasClass

`func (o *WlanQos) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetOverwrite

`func (o *WlanQos) GetOverwrite() bool`

GetOverwrite returns the Overwrite field if non-nil, zero value otherwise.

### GetOverwriteOk

`func (o *WlanQos) GetOverwriteOk() (*bool, bool)`

GetOverwriteOk returns a tuple with the Overwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwrite

`func (o *WlanQos) SetOverwrite(v bool)`

SetOverwrite sets Overwrite field to given value.

### HasOverwrite

`func (o *WlanQos) HasOverwrite() bool`

HasOverwrite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


