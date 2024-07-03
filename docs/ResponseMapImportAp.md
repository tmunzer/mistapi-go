# ResponseMapImportAp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**ResponseMapImportApAction**](ResponseMapImportApAction.md) |  | 
**FloorplanId** | **string** |  | 
**Height** | Pointer to **float32** |  | [optional] 
**Mac** | **string** |  | 
**MapId** | **string** |  | 
**Orientation** | **float32** |  | 
**Reason** | Pointer to **string** |  | [optional] 

## Methods

### NewResponseMapImportAp

`func NewResponseMapImportAp(action ResponseMapImportApAction, floorplanId string, mac string, mapId string, orientation float32, ) *ResponseMapImportAp`

NewResponseMapImportAp instantiates a new ResponseMapImportAp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseMapImportApWithDefaults

`func NewResponseMapImportApWithDefaults() *ResponseMapImportAp`

NewResponseMapImportApWithDefaults instantiates a new ResponseMapImportAp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ResponseMapImportAp) GetAction() ResponseMapImportApAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ResponseMapImportAp) GetActionOk() (*ResponseMapImportApAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ResponseMapImportAp) SetAction(v ResponseMapImportApAction)`

SetAction sets Action field to given value.


### GetFloorplanId

`func (o *ResponseMapImportAp) GetFloorplanId() string`

GetFloorplanId returns the FloorplanId field if non-nil, zero value otherwise.

### GetFloorplanIdOk

`func (o *ResponseMapImportAp) GetFloorplanIdOk() (*string, bool)`

GetFloorplanIdOk returns a tuple with the FloorplanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloorplanId

`func (o *ResponseMapImportAp) SetFloorplanId(v string)`

SetFloorplanId sets FloorplanId field to given value.


### GetHeight

`func (o *ResponseMapImportAp) GetHeight() float32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ResponseMapImportAp) GetHeightOk() (*float32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ResponseMapImportAp) SetHeight(v float32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ResponseMapImportAp) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetMac

`func (o *ResponseMapImportAp) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *ResponseMapImportAp) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *ResponseMapImportAp) SetMac(v string)`

SetMac sets Mac field to given value.


### GetMapId

`func (o *ResponseMapImportAp) GetMapId() string`

GetMapId returns the MapId field if non-nil, zero value otherwise.

### GetMapIdOk

`func (o *ResponseMapImportAp) GetMapIdOk() (*string, bool)`

GetMapIdOk returns a tuple with the MapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapId

`func (o *ResponseMapImportAp) SetMapId(v string)`

SetMapId sets MapId field to given value.


### GetOrientation

`func (o *ResponseMapImportAp) GetOrientation() float32`

GetOrientation returns the Orientation field if non-nil, zero value otherwise.

### GetOrientationOk

`func (o *ResponseMapImportAp) GetOrientationOk() (*float32, bool)`

GetOrientationOk returns a tuple with the Orientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrientation

`func (o *ResponseMapImportAp) SetOrientation(v float32)`

SetOrientation sets Orientation field to given value.


### GetReason

`func (o *ResponseMapImportAp) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ResponseMapImportAp) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ResponseMapImportAp) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ResponseMapImportAp) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


