# CaptureNewAssoc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApMac** | Pointer to **string** |  | [optional] 
**ClientMac** | Pointer to **string** | client mac, required if &#x60;type&#x60;&#x3D;&#x3D;&#x60;client&#x60;; optional otherwise | [optional] 
**Duration** | Pointer to **int32** | duration of the capture, in seconds | [optional] [default to 600]
**IncludesMcast** | Pointer to **bool** |  | [optional] [default to false]
**MaxPktLen** | Pointer to **int32** |  | [optional] [default to 128]
**NumPackets** | Pointer to **int32** | number of packets to capture, 0 for unlimited | [optional] [default to 100]
**Ssid** | Pointer to **string** | optional filter by ssid | [optional] 
**Type** | [**CaptureNewAssocType**](CaptureNewAssocType.md) |  | 

## Methods

### NewCaptureNewAssoc

`func NewCaptureNewAssoc(type_ CaptureNewAssocType, ) *CaptureNewAssoc`

NewCaptureNewAssoc instantiates a new CaptureNewAssoc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptureNewAssocWithDefaults

`func NewCaptureNewAssocWithDefaults() *CaptureNewAssoc`

NewCaptureNewAssocWithDefaults instantiates a new CaptureNewAssoc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApMac

`func (o *CaptureNewAssoc) GetApMac() string`

GetApMac returns the ApMac field if non-nil, zero value otherwise.

### GetApMacOk

`func (o *CaptureNewAssoc) GetApMacOk() (*string, bool)`

GetApMacOk returns a tuple with the ApMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApMac

`func (o *CaptureNewAssoc) SetApMac(v string)`

SetApMac sets ApMac field to given value.

### HasApMac

`func (o *CaptureNewAssoc) HasApMac() bool`

HasApMac returns a boolean if a field has been set.

### GetClientMac

`func (o *CaptureNewAssoc) GetClientMac() string`

GetClientMac returns the ClientMac field if non-nil, zero value otherwise.

### GetClientMacOk

`func (o *CaptureNewAssoc) GetClientMacOk() (*string, bool)`

GetClientMacOk returns a tuple with the ClientMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientMac

`func (o *CaptureNewAssoc) SetClientMac(v string)`

SetClientMac sets ClientMac field to given value.

### HasClientMac

`func (o *CaptureNewAssoc) HasClientMac() bool`

HasClientMac returns a boolean if a field has been set.

### GetDuration

`func (o *CaptureNewAssoc) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *CaptureNewAssoc) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *CaptureNewAssoc) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *CaptureNewAssoc) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetIncludesMcast

`func (o *CaptureNewAssoc) GetIncludesMcast() bool`

GetIncludesMcast returns the IncludesMcast field if non-nil, zero value otherwise.

### GetIncludesMcastOk

`func (o *CaptureNewAssoc) GetIncludesMcastOk() (*bool, bool)`

GetIncludesMcastOk returns a tuple with the IncludesMcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludesMcast

`func (o *CaptureNewAssoc) SetIncludesMcast(v bool)`

SetIncludesMcast sets IncludesMcast field to given value.

### HasIncludesMcast

`func (o *CaptureNewAssoc) HasIncludesMcast() bool`

HasIncludesMcast returns a boolean if a field has been set.

### GetMaxPktLen

`func (o *CaptureNewAssoc) GetMaxPktLen() int32`

GetMaxPktLen returns the MaxPktLen field if non-nil, zero value otherwise.

### GetMaxPktLenOk

`func (o *CaptureNewAssoc) GetMaxPktLenOk() (*int32, bool)`

GetMaxPktLenOk returns a tuple with the MaxPktLen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPktLen

`func (o *CaptureNewAssoc) SetMaxPktLen(v int32)`

SetMaxPktLen sets MaxPktLen field to given value.

### HasMaxPktLen

`func (o *CaptureNewAssoc) HasMaxPktLen() bool`

HasMaxPktLen returns a boolean if a field has been set.

### GetNumPackets

`func (o *CaptureNewAssoc) GetNumPackets() int32`

GetNumPackets returns the NumPackets field if non-nil, zero value otherwise.

### GetNumPacketsOk

`func (o *CaptureNewAssoc) GetNumPacketsOk() (*int32, bool)`

GetNumPacketsOk returns a tuple with the NumPackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPackets

`func (o *CaptureNewAssoc) SetNumPackets(v int32)`

SetNumPackets sets NumPackets field to given value.

### HasNumPackets

`func (o *CaptureNewAssoc) HasNumPackets() bool`

HasNumPackets returns a boolean if a field has been set.

### GetSsid

`func (o *CaptureNewAssoc) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *CaptureNewAssoc) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *CaptureNewAssoc) SetSsid(v string)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *CaptureNewAssoc) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### GetType

`func (o *CaptureNewAssoc) GetType() CaptureNewAssocType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CaptureNewAssoc) GetTypeOk() (*CaptureNewAssocType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CaptureNewAssoc) SetType(v CaptureNewAssocType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


