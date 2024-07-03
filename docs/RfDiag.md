# RfDiag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | Pointer to **int32** | recording length in seconds, max is 180. Default value is also 180. | [optional] [default to 180]
**Mac** | Pointer to **string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;client&#x60; or &#x60;asset&#x60;, mac of the device | [optional] 
**Name** | **string** | name of the recording, the name of the sdk client would be a good default choice | 
**SdkclientId** | Pointer to **string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;sdkclient&#x60;, sdkclient_id of this recording | [optional] 
**Type** | [**RfClientType**](RfClientType.md) |  | 

## Methods

### NewRfDiag

`func NewRfDiag(name string, type_ RfClientType, ) *RfDiag`

NewRfDiag instantiates a new RfDiag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRfDiagWithDefaults

`func NewRfDiagWithDefaults() *RfDiag`

NewRfDiagWithDefaults instantiates a new RfDiag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *RfDiag) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *RfDiag) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *RfDiag) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *RfDiag) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetMac

`func (o *RfDiag) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *RfDiag) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *RfDiag) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *RfDiag) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetName

`func (o *RfDiag) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RfDiag) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RfDiag) SetName(v string)`

SetName sets Name field to given value.


### GetSdkclientId

`func (o *RfDiag) GetSdkclientId() string`

GetSdkclientId returns the SdkclientId field if non-nil, zero value otherwise.

### GetSdkclientIdOk

`func (o *RfDiag) GetSdkclientIdOk() (*string, bool)`

GetSdkclientIdOk returns a tuple with the SdkclientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkclientId

`func (o *RfDiag) SetSdkclientId(v string)`

SetSdkclientId sets SdkclientId field to given value.

### HasSdkclientId

`func (o *RfDiag) HasSdkclientId() bool`

HasSdkclientId returns a boolean if a field has been set.

### GetType

`func (o *RfDiag) GetType() RfClientType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RfDiag) GetTypeOk() (*RfClientType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RfDiag) SetType(v RfClientType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


