# BgpConfigNeighbors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disabled** | Pointer to **bool** | If true, the BGP session to this neighbor will be administratively disabled/shutdown | [optional] [default to false]
**ExportPolicy** | Pointer to **string** |  | [optional] 
**HoldTime** | Pointer to **int32** |  | [optional] [default to 90]
**ImportPolicy** | Pointer to **string** |  | [optional] 
**MultihopTtl** | Pointer to **int32** | assuming BGP neighbor is directly connected | [optional] 
**NeighborAs** | Pointer to **int32** |  | [optional] 

## Methods

### NewBgpConfigNeighbors

`func NewBgpConfigNeighbors() *BgpConfigNeighbors`

NewBgpConfigNeighbors instantiates a new BgpConfigNeighbors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBgpConfigNeighborsWithDefaults

`func NewBgpConfigNeighborsWithDefaults() *BgpConfigNeighbors`

NewBgpConfigNeighborsWithDefaults instantiates a new BgpConfigNeighbors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisabled

`func (o *BgpConfigNeighbors) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *BgpConfigNeighbors) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *BgpConfigNeighbors) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *BgpConfigNeighbors) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetExportPolicy

`func (o *BgpConfigNeighbors) GetExportPolicy() string`

GetExportPolicy returns the ExportPolicy field if non-nil, zero value otherwise.

### GetExportPolicyOk

`func (o *BgpConfigNeighbors) GetExportPolicyOk() (*string, bool)`

GetExportPolicyOk returns a tuple with the ExportPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportPolicy

`func (o *BgpConfigNeighbors) SetExportPolicy(v string)`

SetExportPolicy sets ExportPolicy field to given value.

### HasExportPolicy

`func (o *BgpConfigNeighbors) HasExportPolicy() bool`

HasExportPolicy returns a boolean if a field has been set.

### GetHoldTime

`func (o *BgpConfigNeighbors) GetHoldTime() int32`

GetHoldTime returns the HoldTime field if non-nil, zero value otherwise.

### GetHoldTimeOk

`func (o *BgpConfigNeighbors) GetHoldTimeOk() (*int32, bool)`

GetHoldTimeOk returns a tuple with the HoldTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldTime

`func (o *BgpConfigNeighbors) SetHoldTime(v int32)`

SetHoldTime sets HoldTime field to given value.

### HasHoldTime

`func (o *BgpConfigNeighbors) HasHoldTime() bool`

HasHoldTime returns a boolean if a field has been set.

### GetImportPolicy

`func (o *BgpConfigNeighbors) GetImportPolicy() string`

GetImportPolicy returns the ImportPolicy field if non-nil, zero value otherwise.

### GetImportPolicyOk

`func (o *BgpConfigNeighbors) GetImportPolicyOk() (*string, bool)`

GetImportPolicyOk returns a tuple with the ImportPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportPolicy

`func (o *BgpConfigNeighbors) SetImportPolicy(v string)`

SetImportPolicy sets ImportPolicy field to given value.

### HasImportPolicy

`func (o *BgpConfigNeighbors) HasImportPolicy() bool`

HasImportPolicy returns a boolean if a field has been set.

### GetMultihopTtl

`func (o *BgpConfigNeighbors) GetMultihopTtl() int32`

GetMultihopTtl returns the MultihopTtl field if non-nil, zero value otherwise.

### GetMultihopTtlOk

`func (o *BgpConfigNeighbors) GetMultihopTtlOk() (*int32, bool)`

GetMultihopTtlOk returns a tuple with the MultihopTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultihopTtl

`func (o *BgpConfigNeighbors) SetMultihopTtl(v int32)`

SetMultihopTtl sets MultihopTtl field to given value.

### HasMultihopTtl

`func (o *BgpConfigNeighbors) HasMultihopTtl() bool`

HasMultihopTtl returns a boolean if a field has been set.

### GetNeighborAs

`func (o *BgpConfigNeighbors) GetNeighborAs() int32`

GetNeighborAs returns the NeighborAs field if non-nil, zero value otherwise.

### GetNeighborAsOk

`func (o *BgpConfigNeighbors) GetNeighborAsOk() (*int32, bool)`

GetNeighborAsOk returns a tuple with the NeighborAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighborAs

`func (o *BgpConfigNeighbors) SetNeighborAs(v int32)`

SetNeighborAs sets NeighborAs field to given value.

### HasNeighborAs

`func (o *BgpConfigNeighbors) HasNeighborAs() bool`

HasNeighborAs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


