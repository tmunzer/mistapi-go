# UseAutoApValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accept** | Pointer to **bool** | If accept is true, accepts placement for devices in list otherwise. If false, reject for devices in list. | [optional] [default to false]
**For** | Pointer to [**UseAutoApValuesFor**](UseAutoApValuesFor.md) |  | [optional] [default to USEAUTOAPVALUESFOR_PLACEMENT]
**Macs** | Pointer to **[]string** | A list of macs to accept/reject. If a list is not provided the API will accept/reject for the full map. | [optional] 

## Methods

### NewUseAutoApValues

`func NewUseAutoApValues() *UseAutoApValues`

NewUseAutoApValues instantiates a new UseAutoApValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUseAutoApValuesWithDefaults

`func NewUseAutoApValuesWithDefaults() *UseAutoApValues`

NewUseAutoApValuesWithDefaults instantiates a new UseAutoApValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccept

`func (o *UseAutoApValues) GetAccept() bool`

GetAccept returns the Accept field if non-nil, zero value otherwise.

### GetAcceptOk

`func (o *UseAutoApValues) GetAcceptOk() (*bool, bool)`

GetAcceptOk returns a tuple with the Accept field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccept

`func (o *UseAutoApValues) SetAccept(v bool)`

SetAccept sets Accept field to given value.

### HasAccept

`func (o *UseAutoApValues) HasAccept() bool`

HasAccept returns a boolean if a field has been set.

### GetFor

`func (o *UseAutoApValues) GetFor() UseAutoApValuesFor`

GetFor returns the For field if non-nil, zero value otherwise.

### GetForOk

`func (o *UseAutoApValues) GetForOk() (*UseAutoApValuesFor, bool)`

GetForOk returns a tuple with the For field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFor

`func (o *UseAutoApValues) SetFor(v UseAutoApValuesFor)`

SetFor sets For field to given value.

### HasFor

`func (o *UseAutoApValues) HasFor() bool`

HasFor returns a boolean if a field has been set.

### GetMacs

`func (o *UseAutoApValues) GetMacs() []string`

GetMacs returns the Macs field if non-nil, zero value otherwise.

### GetMacsOk

`func (o *UseAutoApValues) GetMacsOk() (*[]string, bool)`

GetMacsOk returns a tuple with the Macs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacs

`func (o *UseAutoApValues) SetMacs(v []string)`

SetMacs sets Macs field to given value.

### HasMacs

`func (o *UseAutoApValues) HasMacs() bool`

HasMacs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


