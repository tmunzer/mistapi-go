# CaptureClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApMac** | Pointer to **NullableString** |  | [optional] 
**ClientMac** | Pointer to **NullableString** | client mac, required if &#x60;type&#x60;&#x3D;&#x3D;&#x60;client&#x60;; optional otherwise | [optional] 
**Duration** | Pointer to **NullableInt32** | duration of the capture, in seconds | [optional] [default to 600]
**IncludesMcast** | Pointer to **bool** |  | [optional] [default to false]
**MaxPktLen** | Pointer to **NullableInt32** |  | [optional] [default to 128]
**NumPackets** | Pointer to **NullableInt32** | number of packets to capture, 0 for unlimited, default is 1024 for client-capture | [optional] [default to 1024]
**Ssid** | Pointer to **NullableString** | optional filter by ssid | [optional] 
**Type** | [**CaptureClientType**](CaptureClientType.md) |  | 

## Methods

### NewCaptureClient

`func NewCaptureClient(type_ CaptureClientType, ) *CaptureClient`

NewCaptureClient instantiates a new CaptureClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptureClientWithDefaults

`func NewCaptureClientWithDefaults() *CaptureClient`

NewCaptureClientWithDefaults instantiates a new CaptureClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApMac

`func (o *CaptureClient) GetApMac() string`

GetApMac returns the ApMac field if non-nil, zero value otherwise.

### GetApMacOk

`func (o *CaptureClient) GetApMacOk() (*string, bool)`

GetApMacOk returns a tuple with the ApMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApMac

`func (o *CaptureClient) SetApMac(v string)`

SetApMac sets ApMac field to given value.

### HasApMac

`func (o *CaptureClient) HasApMac() bool`

HasApMac returns a boolean if a field has been set.

### SetApMacNil

`func (o *CaptureClient) SetApMacNil(b bool)`

 SetApMacNil sets the value for ApMac to be an explicit nil

### UnsetApMac
`func (o *CaptureClient) UnsetApMac()`

UnsetApMac ensures that no value is present for ApMac, not even an explicit nil
### GetClientMac

`func (o *CaptureClient) GetClientMac() string`

GetClientMac returns the ClientMac field if non-nil, zero value otherwise.

### GetClientMacOk

`func (o *CaptureClient) GetClientMacOk() (*string, bool)`

GetClientMacOk returns a tuple with the ClientMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientMac

`func (o *CaptureClient) SetClientMac(v string)`

SetClientMac sets ClientMac field to given value.

### HasClientMac

`func (o *CaptureClient) HasClientMac() bool`

HasClientMac returns a boolean if a field has been set.

### SetClientMacNil

`func (o *CaptureClient) SetClientMacNil(b bool)`

 SetClientMacNil sets the value for ClientMac to be an explicit nil

### UnsetClientMac
`func (o *CaptureClient) UnsetClientMac()`

UnsetClientMac ensures that no value is present for ClientMac, not even an explicit nil
### GetDuration

`func (o *CaptureClient) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *CaptureClient) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *CaptureClient) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *CaptureClient) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *CaptureClient) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *CaptureClient) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetIncludesMcast

`func (o *CaptureClient) GetIncludesMcast() bool`

GetIncludesMcast returns the IncludesMcast field if non-nil, zero value otherwise.

### GetIncludesMcastOk

`func (o *CaptureClient) GetIncludesMcastOk() (*bool, bool)`

GetIncludesMcastOk returns a tuple with the IncludesMcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludesMcast

`func (o *CaptureClient) SetIncludesMcast(v bool)`

SetIncludesMcast sets IncludesMcast field to given value.

### HasIncludesMcast

`func (o *CaptureClient) HasIncludesMcast() bool`

HasIncludesMcast returns a boolean if a field has been set.

### GetMaxPktLen

`func (o *CaptureClient) GetMaxPktLen() int32`

GetMaxPktLen returns the MaxPktLen field if non-nil, zero value otherwise.

### GetMaxPktLenOk

`func (o *CaptureClient) GetMaxPktLenOk() (*int32, bool)`

GetMaxPktLenOk returns a tuple with the MaxPktLen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPktLen

`func (o *CaptureClient) SetMaxPktLen(v int32)`

SetMaxPktLen sets MaxPktLen field to given value.

### HasMaxPktLen

`func (o *CaptureClient) HasMaxPktLen() bool`

HasMaxPktLen returns a boolean if a field has been set.

### SetMaxPktLenNil

`func (o *CaptureClient) SetMaxPktLenNil(b bool)`

 SetMaxPktLenNil sets the value for MaxPktLen to be an explicit nil

### UnsetMaxPktLen
`func (o *CaptureClient) UnsetMaxPktLen()`

UnsetMaxPktLen ensures that no value is present for MaxPktLen, not even an explicit nil
### GetNumPackets

`func (o *CaptureClient) GetNumPackets() int32`

GetNumPackets returns the NumPackets field if non-nil, zero value otherwise.

### GetNumPacketsOk

`func (o *CaptureClient) GetNumPacketsOk() (*int32, bool)`

GetNumPacketsOk returns a tuple with the NumPackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPackets

`func (o *CaptureClient) SetNumPackets(v int32)`

SetNumPackets sets NumPackets field to given value.

### HasNumPackets

`func (o *CaptureClient) HasNumPackets() bool`

HasNumPackets returns a boolean if a field has been set.

### SetNumPacketsNil

`func (o *CaptureClient) SetNumPacketsNil(b bool)`

 SetNumPacketsNil sets the value for NumPackets to be an explicit nil

### UnsetNumPackets
`func (o *CaptureClient) UnsetNumPackets()`

UnsetNumPackets ensures that no value is present for NumPackets, not even an explicit nil
### GetSsid

`func (o *CaptureClient) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *CaptureClient) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *CaptureClient) SetSsid(v string)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *CaptureClient) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### SetSsidNil

`func (o *CaptureClient) SetSsidNil(b bool)`

 SetSsidNil sets the value for Ssid to be an explicit nil

### UnsetSsid
`func (o *CaptureClient) UnsetSsid()`

UnsetSsid ensures that no value is present for Ssid, not even an explicit nil
### GetType

`func (o *CaptureClient) GetType() CaptureClientType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CaptureClient) GetTypeOk() (*CaptureClientType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CaptureClient) SetType(v CaptureClientType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


