# MapSiteImportFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoDeviceprofileAssignment** | Pointer to **bool** | whether to auto assign device to deviceprofile by name | [optional] 
**Csv** | Pointer to ***os.File** | csv file for ap name mapping, optional | [optional] 
**File** | Pointer to ***os.File** | ekahau or ibwave file | [optional] 
**Json** | Pointer to [**MapImportJson**](MapImportJson.md) |  | [optional] 

## Methods

### NewMapSiteImportFile

`func NewMapSiteImportFile() *MapSiteImportFile`

NewMapSiteImportFile instantiates a new MapSiteImportFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMapSiteImportFileWithDefaults

`func NewMapSiteImportFileWithDefaults() *MapSiteImportFile`

NewMapSiteImportFileWithDefaults instantiates a new MapSiteImportFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoDeviceprofileAssignment

`func (o *MapSiteImportFile) GetAutoDeviceprofileAssignment() bool`

GetAutoDeviceprofileAssignment returns the AutoDeviceprofileAssignment field if non-nil, zero value otherwise.

### GetAutoDeviceprofileAssignmentOk

`func (o *MapSiteImportFile) GetAutoDeviceprofileAssignmentOk() (*bool, bool)`

GetAutoDeviceprofileAssignmentOk returns a tuple with the AutoDeviceprofileAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDeviceprofileAssignment

`func (o *MapSiteImportFile) SetAutoDeviceprofileAssignment(v bool)`

SetAutoDeviceprofileAssignment sets AutoDeviceprofileAssignment field to given value.

### HasAutoDeviceprofileAssignment

`func (o *MapSiteImportFile) HasAutoDeviceprofileAssignment() bool`

HasAutoDeviceprofileAssignment returns a boolean if a field has been set.

### GetCsv

`func (o *MapSiteImportFile) GetCsv() *os.File`

GetCsv returns the Csv field if non-nil, zero value otherwise.

### GetCsvOk

`func (o *MapSiteImportFile) GetCsvOk() (**os.File, bool)`

GetCsvOk returns a tuple with the Csv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsv

`func (o *MapSiteImportFile) SetCsv(v *os.File)`

SetCsv sets Csv field to given value.

### HasCsv

`func (o *MapSiteImportFile) HasCsv() bool`

HasCsv returns a boolean if a field has been set.

### GetFile

`func (o *MapSiteImportFile) GetFile() *os.File`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *MapSiteImportFile) GetFileOk() (**os.File, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *MapSiteImportFile) SetFile(v *os.File)`

SetFile sets File field to given value.

### HasFile

`func (o *MapSiteImportFile) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetJson

`func (o *MapSiteImportFile) GetJson() MapImportJson`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *MapSiteImportFile) GetJsonOk() (*MapImportJson, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *MapSiteImportFile) SetJson(v MapImportJson)`

SetJson sets Json field to given value.

### HasJson

`func (o *MapSiteImportFile) HasJson() bool`

HasJson returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


