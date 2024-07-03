# DeviceVersionItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | **string** | Device model (as seen in the device stats) | 
**Tag** | Pointer to **string** | annotation, stable / beta / alpha. Or it can be empty or nothing which is likely a dev build | [optional] 
**Version** | **string** | firmware version | 

## Methods

### NewDeviceVersionItem

`func NewDeviceVersionItem(model string, version string, ) *DeviceVersionItem`

NewDeviceVersionItem instantiates a new DeviceVersionItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceVersionItemWithDefaults

`func NewDeviceVersionItemWithDefaults() *DeviceVersionItem`

NewDeviceVersionItemWithDefaults instantiates a new DeviceVersionItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *DeviceVersionItem) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *DeviceVersionItem) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *DeviceVersionItem) SetModel(v string)`

SetModel sets Model field to given value.


### GetTag

`func (o *DeviceVersionItem) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *DeviceVersionItem) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *DeviceVersionItem) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *DeviceVersionItem) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetVersion

`func (o *DeviceVersionItem) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DeviceVersionItem) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DeviceVersionItem) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


