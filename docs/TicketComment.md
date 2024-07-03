# TicketComment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttachmentIds** | Pointer to **[]string** |  | [optional] [readonly] 
**Attachments** | Pointer to [**[]TicketCommentsAttachment**](TicketCommentsAttachment.md) |  | [optional] [readonly] 
**Author** | **string** |  | [readonly] 
**Comment** | **string** |  | 
**CreatedAt** | **int32** |  | [readonly] 

## Methods

### NewTicketComment

`func NewTicketComment(author string, comment string, createdAt int32, ) *TicketComment`

NewTicketComment instantiates a new TicketComment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTicketCommentWithDefaults

`func NewTicketCommentWithDefaults() *TicketComment`

NewTicketCommentWithDefaults instantiates a new TicketComment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachmentIds

`func (o *TicketComment) GetAttachmentIds() []string`

GetAttachmentIds returns the AttachmentIds field if non-nil, zero value otherwise.

### GetAttachmentIdsOk

`func (o *TicketComment) GetAttachmentIdsOk() (*[]string, bool)`

GetAttachmentIdsOk returns a tuple with the AttachmentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentIds

`func (o *TicketComment) SetAttachmentIds(v []string)`

SetAttachmentIds sets AttachmentIds field to given value.

### HasAttachmentIds

`func (o *TicketComment) HasAttachmentIds() bool`

HasAttachmentIds returns a boolean if a field has been set.

### GetAttachments

`func (o *TicketComment) GetAttachments() []TicketCommentsAttachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *TicketComment) GetAttachmentsOk() (*[]TicketCommentsAttachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *TicketComment) SetAttachments(v []TicketCommentsAttachment)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *TicketComment) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetAuthor

`func (o *TicketComment) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *TicketComment) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *TicketComment) SetAuthor(v string)`

SetAuthor sets Author field to given value.


### GetComment

`func (o *TicketComment) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *TicketComment) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *TicketComment) SetComment(v string)`

SetComment sets Comment field to given value.


### GetCreatedAt

`func (o *TicketComment) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TicketComment) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TicketComment) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


