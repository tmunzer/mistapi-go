# CaptureMxedge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | Pointer to **int32** | duration of the capture, in seconds | [optional] [default to 600]
**Format** | Pointer to [**CaptureMxedgeFormat**](CaptureMxedgeFormat.md) |  | [optional] [default to CAPTUREMXEDGEFORMAT_STREAM]
**MaxPktLen** | Pointer to **int32** | max_len of each packet to capture | [optional] [default to 128]
**Mxedges** | Pointer to [**map[string]CaptureMxedgeMxedges**](CaptureMxedgeMxedges.md) |  | [optional] 
**NumPackets** | Pointer to **int32** | number of packets to capture, 0 for unlimited | [optional] [default to 1024]
**Type** | [**CaptureMxedgeType**](CaptureMxedgeType.md) |  | 

## Methods

### NewCaptureMxedge

`func NewCaptureMxedge(type_ CaptureMxedgeType, ) *CaptureMxedge`

NewCaptureMxedge instantiates a new CaptureMxedge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptureMxedgeWithDefaults

`func NewCaptureMxedgeWithDefaults() *CaptureMxedge`

NewCaptureMxedgeWithDefaults instantiates a new CaptureMxedge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *CaptureMxedge) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *CaptureMxedge) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *CaptureMxedge) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *CaptureMxedge) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetFormat

`func (o *CaptureMxedge) GetFormat() CaptureMxedgeFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *CaptureMxedge) GetFormatOk() (*CaptureMxedgeFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *CaptureMxedge) SetFormat(v CaptureMxedgeFormat)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *CaptureMxedge) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetMaxPktLen

`func (o *CaptureMxedge) GetMaxPktLen() int32`

GetMaxPktLen returns the MaxPktLen field if non-nil, zero value otherwise.

### GetMaxPktLenOk

`func (o *CaptureMxedge) GetMaxPktLenOk() (*int32, bool)`

GetMaxPktLenOk returns a tuple with the MaxPktLen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPktLen

`func (o *CaptureMxedge) SetMaxPktLen(v int32)`

SetMaxPktLen sets MaxPktLen field to given value.

### HasMaxPktLen

`func (o *CaptureMxedge) HasMaxPktLen() bool`

HasMaxPktLen returns a boolean if a field has been set.

### GetMxedges

`func (o *CaptureMxedge) GetMxedges() map[string]CaptureMxedgeMxedges`

GetMxedges returns the Mxedges field if non-nil, zero value otherwise.

### GetMxedgesOk

`func (o *CaptureMxedge) GetMxedgesOk() (*map[string]CaptureMxedgeMxedges, bool)`

GetMxedgesOk returns a tuple with the Mxedges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMxedges

`func (o *CaptureMxedge) SetMxedges(v map[string]CaptureMxedgeMxedges)`

SetMxedges sets Mxedges field to given value.

### HasMxedges

`func (o *CaptureMxedge) HasMxedges() bool`

HasMxedges returns a boolean if a field has been set.

### GetNumPackets

`func (o *CaptureMxedge) GetNumPackets() int32`

GetNumPackets returns the NumPackets field if non-nil, zero value otherwise.

### GetNumPacketsOk

`func (o *CaptureMxedge) GetNumPacketsOk() (*int32, bool)`

GetNumPacketsOk returns a tuple with the NumPackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPackets

`func (o *CaptureMxedge) SetNumPackets(v int32)`

SetNumPackets sets NumPackets field to given value.

### HasNumPackets

`func (o *CaptureMxedge) HasNumPackets() bool`

HasNumPackets returns a boolean if a field has been set.

### GetType

`func (o *CaptureMxedge) GetType() CaptureMxedgeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CaptureMxedge) GetTypeOk() (*CaptureMxedgeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CaptureMxedge) SetType(v CaptureMxedgeType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


