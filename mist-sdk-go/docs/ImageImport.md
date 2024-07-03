# ImageImport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**File** | ***os.File** | binary file | 
**Json** | Pointer to **string** | JSON string describing your upload | [optional] 

## Methods

### NewImageImport

`func NewImageImport(file *os.File, ) *ImageImport`

NewImageImport instantiates a new ImageImport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageImportWithDefaults

`func NewImageImportWithDefaults() *ImageImport`

NewImageImportWithDefaults instantiates a new ImageImport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFile

`func (o *ImageImport) GetFile() *os.File`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *ImageImport) GetFileOk() (**os.File, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *ImageImport) SetFile(v *os.File)`

SetFile sets File field to given value.


### GetJson

`func (o *ImageImport) GetJson() string`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *ImageImport) GetJsonOk() (*string, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *ImageImport) SetJson(v string)`

SetJson sets Json field to given value.

### HasJson

`func (o *ImageImport) HasJson() bool`

HasJson returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


