# PskPortalImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**File** | Pointer to ***os.File** | Binary file | [optional] 
**Json** | Pointer to **string** | JSON string describing the upload | [optional] 

## Methods

### NewPskPortalImage

`func NewPskPortalImage() *PskPortalImage`

NewPskPortalImage instantiates a new PskPortalImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPskPortalImageWithDefaults

`func NewPskPortalImageWithDefaults() *PskPortalImage`

NewPskPortalImageWithDefaults instantiates a new PskPortalImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFile

`func (o *PskPortalImage) GetFile() *os.File`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *PskPortalImage) GetFileOk() (**os.File, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *PskPortalImage) SetFile(v *os.File)`

SetFile sets File field to given value.

### HasFile

`func (o *PskPortalImage) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetJson

`func (o *PskPortalImage) GetJson() string`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *PskPortalImage) GetJsonOk() (*string, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *PskPortalImage) SetJson(v string)`

SetJson sets Json field to given value.

### HasJson

`func (o *PskPortalImage) HasJson() bool`

HasJson returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


