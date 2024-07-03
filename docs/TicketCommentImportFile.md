# TicketCommentImportFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** |  | [optional] 
**File** | Pointer to ***os.File** |  | [optional] 

## Methods

### NewTicketCommentImportFile

`func NewTicketCommentImportFile() *TicketCommentImportFile`

NewTicketCommentImportFile instantiates a new TicketCommentImportFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTicketCommentImportFileWithDefaults

`func NewTicketCommentImportFileWithDefaults() *TicketCommentImportFile`

NewTicketCommentImportFileWithDefaults instantiates a new TicketCommentImportFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *TicketCommentImportFile) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *TicketCommentImportFile) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *TicketCommentImportFile) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *TicketCommentImportFile) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetFile

`func (o *TicketCommentImportFile) GetFile() *os.File`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *TicketCommentImportFile) GetFileOk() (**os.File, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *TicketCommentImportFile) SetFile(v *os.File)`

SetFile sets File field to given value.

### HasFile

`func (o *TicketCommentImportFile) HasFile() bool`

HasFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


