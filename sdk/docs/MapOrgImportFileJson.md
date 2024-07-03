# MapOrgImportFileJson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImportAllFloorplans** | Pointer to **bool** |  | [optional] [default to false]
**ImportHeight** | Pointer to **bool** |  | [optional] [default to true]
**ImportOrientation** | Pointer to **bool** |  | [optional] [default to true]
**SiteId** | Pointer to **string** |  | [optional] [readonly] 
**VendorName** | [**MapOrgImportFileJsonVendorName**](MapOrgImportFileJsonVendorName.md) |  | 

## Methods

### NewMapOrgImportFileJson

`func NewMapOrgImportFileJson(vendorName MapOrgImportFileJsonVendorName, ) *MapOrgImportFileJson`

NewMapOrgImportFileJson instantiates a new MapOrgImportFileJson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMapOrgImportFileJsonWithDefaults

`func NewMapOrgImportFileJsonWithDefaults() *MapOrgImportFileJson`

NewMapOrgImportFileJsonWithDefaults instantiates a new MapOrgImportFileJson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImportAllFloorplans

`func (o *MapOrgImportFileJson) GetImportAllFloorplans() bool`

GetImportAllFloorplans returns the ImportAllFloorplans field if non-nil, zero value otherwise.

### GetImportAllFloorplansOk

`func (o *MapOrgImportFileJson) GetImportAllFloorplansOk() (*bool, bool)`

GetImportAllFloorplansOk returns a tuple with the ImportAllFloorplans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportAllFloorplans

`func (o *MapOrgImportFileJson) SetImportAllFloorplans(v bool)`

SetImportAllFloorplans sets ImportAllFloorplans field to given value.

### HasImportAllFloorplans

`func (o *MapOrgImportFileJson) HasImportAllFloorplans() bool`

HasImportAllFloorplans returns a boolean if a field has been set.

### GetImportHeight

`func (o *MapOrgImportFileJson) GetImportHeight() bool`

GetImportHeight returns the ImportHeight field if non-nil, zero value otherwise.

### GetImportHeightOk

`func (o *MapOrgImportFileJson) GetImportHeightOk() (*bool, bool)`

GetImportHeightOk returns a tuple with the ImportHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportHeight

`func (o *MapOrgImportFileJson) SetImportHeight(v bool)`

SetImportHeight sets ImportHeight field to given value.

### HasImportHeight

`func (o *MapOrgImportFileJson) HasImportHeight() bool`

HasImportHeight returns a boolean if a field has been set.

### GetImportOrientation

`func (o *MapOrgImportFileJson) GetImportOrientation() bool`

GetImportOrientation returns the ImportOrientation field if non-nil, zero value otherwise.

### GetImportOrientationOk

`func (o *MapOrgImportFileJson) GetImportOrientationOk() (*bool, bool)`

GetImportOrientationOk returns a tuple with the ImportOrientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportOrientation

`func (o *MapOrgImportFileJson) SetImportOrientation(v bool)`

SetImportOrientation sets ImportOrientation field to given value.

### HasImportOrientation

`func (o *MapOrgImportFileJson) HasImportOrientation() bool`

HasImportOrientation returns a boolean if a field has been set.

### GetSiteId

`func (o *MapOrgImportFileJson) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *MapOrgImportFileJson) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *MapOrgImportFileJson) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *MapOrgImportFileJson) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetVendorName

`func (o *MapOrgImportFileJson) GetVendorName() MapOrgImportFileJsonVendorName`

GetVendorName returns the VendorName field if non-nil, zero value otherwise.

### GetVendorNameOk

`func (o *MapOrgImportFileJson) GetVendorNameOk() (*MapOrgImportFileJsonVendorName, bool)`

GetVendorNameOk returns a tuple with the VendorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorName

`func (o *MapOrgImportFileJson) SetVendorName(v MapOrgImportFileJsonVendorName)`

SetVendorName sets VendorName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


