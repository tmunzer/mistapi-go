# MapOrgImportFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoDeviceprofileAssignment** | Pointer to **bool** | whether to auto assign device to deviceprofile by name | [optional] 
**Csv** | Pointer to ***os.File** | csv file for ap name mapping, optional | [optional] 
**File** | Pointer to ***os.File** | ekahau or ibwave file | [optional] 
**Json** | Pointer to [**MapOrgImportFileJson**](MapOrgImportFileJson.md) |  | [optional] 

## Methods

### NewMapOrgImportFile

`func NewMapOrgImportFile() *MapOrgImportFile`

NewMapOrgImportFile instantiates a new MapOrgImportFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMapOrgImportFileWithDefaults

`func NewMapOrgImportFileWithDefaults() *MapOrgImportFile`

NewMapOrgImportFileWithDefaults instantiates a new MapOrgImportFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoDeviceprofileAssignment

`func (o *MapOrgImportFile) GetAutoDeviceprofileAssignment() bool`

GetAutoDeviceprofileAssignment returns the AutoDeviceprofileAssignment field if non-nil, zero value otherwise.

### GetAutoDeviceprofileAssignmentOk

`func (o *MapOrgImportFile) GetAutoDeviceprofileAssignmentOk() (*bool, bool)`

GetAutoDeviceprofileAssignmentOk returns a tuple with the AutoDeviceprofileAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDeviceprofileAssignment

`func (o *MapOrgImportFile) SetAutoDeviceprofileAssignment(v bool)`

SetAutoDeviceprofileAssignment sets AutoDeviceprofileAssignment field to given value.

### HasAutoDeviceprofileAssignment

`func (o *MapOrgImportFile) HasAutoDeviceprofileAssignment() bool`

HasAutoDeviceprofileAssignment returns a boolean if a field has been set.

### GetCsv

`func (o *MapOrgImportFile) GetCsv() *os.File`

GetCsv returns the Csv field if non-nil, zero value otherwise.

### GetCsvOk

`func (o *MapOrgImportFile) GetCsvOk() (**os.File, bool)`

GetCsvOk returns a tuple with the Csv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsv

`func (o *MapOrgImportFile) SetCsv(v *os.File)`

SetCsv sets Csv field to given value.

### HasCsv

`func (o *MapOrgImportFile) HasCsv() bool`

HasCsv returns a boolean if a field has been set.

### GetFile

`func (o *MapOrgImportFile) GetFile() *os.File`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *MapOrgImportFile) GetFileOk() (**os.File, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *MapOrgImportFile) SetFile(v *os.File)`

SetFile sets File field to given value.

### HasFile

`func (o *MapOrgImportFile) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetJson

`func (o *MapOrgImportFile) GetJson() MapOrgImportFileJson`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *MapOrgImportFile) GetJsonOk() (*MapOrgImportFileJson, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *MapOrgImportFile) SetJson(v MapOrgImportFileJson)`

SetJson sets Json field to given value.

### HasJson

`func (o *MapOrgImportFile) HasJson() bool`

HasJson returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


