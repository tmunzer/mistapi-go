# ClaimActivation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | activation code | 
**DeviceType** | Pointer to [**DeviceType**](DeviceType.md) |  | [optional] [default to DEVICETYPE_AP]
**Type** | [**ClaimType**](ClaimType.md) |  | [default to CLAIMTYPE_ALL]

## Methods

### NewClaimActivation

`func NewClaimActivation(code string, type_ ClaimType, ) *ClaimActivation`

NewClaimActivation instantiates a new ClaimActivation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClaimActivationWithDefaults

`func NewClaimActivationWithDefaults() *ClaimActivation`

NewClaimActivationWithDefaults instantiates a new ClaimActivation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ClaimActivation) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ClaimActivation) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ClaimActivation) SetCode(v string)`

SetCode sets Code field to given value.


### GetDeviceType

`func (o *ClaimActivation) GetDeviceType() DeviceType`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *ClaimActivation) GetDeviceTypeOk() (*DeviceType, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *ClaimActivation) SetDeviceType(v DeviceType)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *ClaimActivation) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetType

`func (o *ClaimActivation) GetType() ClaimType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClaimActivation) GetTypeOk() (*ClaimType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClaimActivation) SetType(v ClaimType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


