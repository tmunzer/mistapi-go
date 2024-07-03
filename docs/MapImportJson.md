# MapImportJson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImportAllFloorplans** | Pointer to **bool** |  | [optional] [default to false]
**ImportHeight** | Pointer to **bool** |  | [optional] [default to true]
**ImportOrientation** | Pointer to **bool** |  | [optional] [default to true]
**VendorName** | [**MapImportJsonVendorName**](MapImportJsonVendorName.md) |  | 

## Methods

### NewMapImportJson

`func NewMapImportJson(vendorName MapImportJsonVendorName, ) *MapImportJson`

NewMapImportJson instantiates a new MapImportJson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMapImportJsonWithDefaults

`func NewMapImportJsonWithDefaults() *MapImportJson`

NewMapImportJsonWithDefaults instantiates a new MapImportJson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImportAllFloorplans

`func (o *MapImportJson) GetImportAllFloorplans() bool`

GetImportAllFloorplans returns the ImportAllFloorplans field if non-nil, zero value otherwise.

### GetImportAllFloorplansOk

`func (o *MapImportJson) GetImportAllFloorplansOk() (*bool, bool)`

GetImportAllFloorplansOk returns a tuple with the ImportAllFloorplans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportAllFloorplans

`func (o *MapImportJson) SetImportAllFloorplans(v bool)`

SetImportAllFloorplans sets ImportAllFloorplans field to given value.

### HasImportAllFloorplans

`func (o *MapImportJson) HasImportAllFloorplans() bool`

HasImportAllFloorplans returns a boolean if a field has been set.

### GetImportHeight

`func (o *MapImportJson) GetImportHeight() bool`

GetImportHeight returns the ImportHeight field if non-nil, zero value otherwise.

### GetImportHeightOk

`func (o *MapImportJson) GetImportHeightOk() (*bool, bool)`

GetImportHeightOk returns a tuple with the ImportHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportHeight

`func (o *MapImportJson) SetImportHeight(v bool)`

SetImportHeight sets ImportHeight field to given value.

### HasImportHeight

`func (o *MapImportJson) HasImportHeight() bool`

HasImportHeight returns a boolean if a field has been set.

### GetImportOrientation

`func (o *MapImportJson) GetImportOrientation() bool`

GetImportOrientation returns the ImportOrientation field if non-nil, zero value otherwise.

### GetImportOrientationOk

`func (o *MapImportJson) GetImportOrientationOk() (*bool, bool)`

GetImportOrientationOk returns a tuple with the ImportOrientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportOrientation

`func (o *MapImportJson) SetImportOrientation(v bool)`

SetImportOrientation sets ImportOrientation field to given value.

### HasImportOrientation

`func (o *MapImportJson) HasImportOrientation() bool`

HasImportOrientation returns a boolean if a field has been set.

### GetVendorName

`func (o *MapImportJson) GetVendorName() MapImportJsonVendorName`

GetVendorName returns the VendorName field if non-nil, zero value otherwise.

### GetVendorNameOk

`func (o *MapImportJson) GetVendorNameOk() (*MapImportJsonVendorName, bool)`

GetVendorNameOk returns a tuple with the VendorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorName

`func (o *MapImportJson) SetVendorName(v MapImportJsonVendorName)`

SetVendorName sets VendorName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


