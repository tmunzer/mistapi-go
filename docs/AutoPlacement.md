# AutoPlacement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForceCollection** | Pointer to **bool** | * If &#x60;force_collection&#x60;&#x3D;&#x3D;&#x60;false&#x60;: the API Iattempts to start localization with existing data.  * If &#x60;force_collection&#x60;&#x3D;&#x3D;&#x60;true&#x60;: maintenance the API attempts to start orchestration. | [optional] [default to false]
**Macs** | Pointer to **[]string** | list of device macs | [optional] 

## Methods

### NewAutoPlacement

`func NewAutoPlacement() *AutoPlacement`

NewAutoPlacement instantiates a new AutoPlacement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoPlacementWithDefaults

`func NewAutoPlacementWithDefaults() *AutoPlacement`

NewAutoPlacementWithDefaults instantiates a new AutoPlacement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForceCollection

`func (o *AutoPlacement) GetForceCollection() bool`

GetForceCollection returns the ForceCollection field if non-nil, zero value otherwise.

### GetForceCollectionOk

`func (o *AutoPlacement) GetForceCollectionOk() (*bool, bool)`

GetForceCollectionOk returns a tuple with the ForceCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceCollection

`func (o *AutoPlacement) SetForceCollection(v bool)`

SetForceCollection sets ForceCollection field to given value.

### HasForceCollection

`func (o *AutoPlacement) HasForceCollection() bool`

HasForceCollection returns a boolean if a field has been set.

### GetMacs

`func (o *AutoPlacement) GetMacs() []string`

GetMacs returns the Macs field if non-nil, zero value otherwise.

### GetMacsOk

`func (o *AutoPlacement) GetMacsOk() (*[]string, bool)`

GetMacsOk returns a tuple with the Macs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacs

`func (o *AutoPlacement) SetMacs(v []string)`

SetMacs sets Macs field to given value.

### HasMacs

`func (o *AutoPlacement) HasMacs() bool`

HasMacs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


