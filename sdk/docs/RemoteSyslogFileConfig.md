# RemoteSyslogFileConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Archive** | Pointer to [**RemoteSyslogArchive**](RemoteSyslogArchive.md) |  | [optional] 
**Contents** | Pointer to [**[]RemoteSyslogContent**](RemoteSyslogContent.md) |  | [optional] 
**ExplicitPriority** | Pointer to **bool** |  | [optional] 
**File** | Pointer to **string** |  | [optional] 
**Match** | Pointer to **string** |  | [optional] 
**StructuredData** | Pointer to **bool** |  | [optional] 

## Methods

### NewRemoteSyslogFileConfig

`func NewRemoteSyslogFileConfig() *RemoteSyslogFileConfig`

NewRemoteSyslogFileConfig instantiates a new RemoteSyslogFileConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteSyslogFileConfigWithDefaults

`func NewRemoteSyslogFileConfigWithDefaults() *RemoteSyslogFileConfig`

NewRemoteSyslogFileConfigWithDefaults instantiates a new RemoteSyslogFileConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchive

`func (o *RemoteSyslogFileConfig) GetArchive() RemoteSyslogArchive`

GetArchive returns the Archive field if non-nil, zero value otherwise.

### GetArchiveOk

`func (o *RemoteSyslogFileConfig) GetArchiveOk() (*RemoteSyslogArchive, bool)`

GetArchiveOk returns a tuple with the Archive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchive

`func (o *RemoteSyslogFileConfig) SetArchive(v RemoteSyslogArchive)`

SetArchive sets Archive field to given value.

### HasArchive

`func (o *RemoteSyslogFileConfig) HasArchive() bool`

HasArchive returns a boolean if a field has been set.

### GetContents

`func (o *RemoteSyslogFileConfig) GetContents() []RemoteSyslogContent`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *RemoteSyslogFileConfig) GetContentsOk() (*[]RemoteSyslogContent, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *RemoteSyslogFileConfig) SetContents(v []RemoteSyslogContent)`

SetContents sets Contents field to given value.

### HasContents

`func (o *RemoteSyslogFileConfig) HasContents() bool`

HasContents returns a boolean if a field has been set.

### GetExplicitPriority

`func (o *RemoteSyslogFileConfig) GetExplicitPriority() bool`

GetExplicitPriority returns the ExplicitPriority field if non-nil, zero value otherwise.

### GetExplicitPriorityOk

`func (o *RemoteSyslogFileConfig) GetExplicitPriorityOk() (*bool, bool)`

GetExplicitPriorityOk returns a tuple with the ExplicitPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitPriority

`func (o *RemoteSyslogFileConfig) SetExplicitPriority(v bool)`

SetExplicitPriority sets ExplicitPriority field to given value.

### HasExplicitPriority

`func (o *RemoteSyslogFileConfig) HasExplicitPriority() bool`

HasExplicitPriority returns a boolean if a field has been set.

### GetFile

`func (o *RemoteSyslogFileConfig) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *RemoteSyslogFileConfig) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *RemoteSyslogFileConfig) SetFile(v string)`

SetFile sets File field to given value.

### HasFile

`func (o *RemoteSyslogFileConfig) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetMatch

`func (o *RemoteSyslogFileConfig) GetMatch() string`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *RemoteSyslogFileConfig) GetMatchOk() (*string, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *RemoteSyslogFileConfig) SetMatch(v string)`

SetMatch sets Match field to given value.

### HasMatch

`func (o *RemoteSyslogFileConfig) HasMatch() bool`

HasMatch returns a boolean if a field has been set.

### GetStructuredData

`func (o *RemoteSyslogFileConfig) GetStructuredData() bool`

GetStructuredData returns the StructuredData field if non-nil, zero value otherwise.

### GetStructuredDataOk

`func (o *RemoteSyslogFileConfig) GetStructuredDataOk() (*bool, bool)`

GetStructuredDataOk returns a tuple with the StructuredData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructuredData

`func (o *RemoteSyslogFileConfig) SetStructuredData(v bool)`

SetStructuredData sets StructuredData field to given value.

### HasStructuredData

`func (o *RemoteSyslogFileConfig) HasStructuredData() bool`

HasStructuredData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


