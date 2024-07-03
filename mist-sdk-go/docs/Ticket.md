# Ticket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaseNumber** | Pointer to **string** |  | [optional] [readonly] 
**Comments** | Pointer to [**[]TicketComment**](TicketComment.md) |  | [optional] 
**CreatedAt** | Pointer to **int32** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Requester** | Pointer to **string** |  | [optional] [readonly] 
**RequesterEmail** | Pointer to **string** | email of the requester | [optional] 
**Status** | Pointer to [**TicketStatus**](TicketStatus.md) |  | [optional] 
**Subject** | **string** |  | 
**Type** | **string** | question (default) / bug / critical | 
**UpdatedAt** | Pointer to **int32** |  | [optional] [readonly] 

## Methods

### NewTicket

`func NewTicket(subject string, type_ string, ) *Ticket`

NewTicket instantiates a new Ticket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTicketWithDefaults

`func NewTicketWithDefaults() *Ticket`

NewTicketWithDefaults instantiates a new Ticket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaseNumber

`func (o *Ticket) GetCaseNumber() string`

GetCaseNumber returns the CaseNumber field if non-nil, zero value otherwise.

### GetCaseNumberOk

`func (o *Ticket) GetCaseNumberOk() (*string, bool)`

GetCaseNumberOk returns a tuple with the CaseNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseNumber

`func (o *Ticket) SetCaseNumber(v string)`

SetCaseNumber sets CaseNumber field to given value.

### HasCaseNumber

`func (o *Ticket) HasCaseNumber() bool`

HasCaseNumber returns a boolean if a field has been set.

### GetComments

`func (o *Ticket) GetComments() []TicketComment`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *Ticket) GetCommentsOk() (*[]TicketComment, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *Ticket) SetComments(v []TicketComment)`

SetComments sets Comments field to given value.

### HasComments

`func (o *Ticket) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Ticket) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Ticket) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Ticket) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Ticket) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *Ticket) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Ticket) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Ticket) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Ticket) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRequester

`func (o *Ticket) GetRequester() string`

GetRequester returns the Requester field if non-nil, zero value otherwise.

### GetRequesterOk

`func (o *Ticket) GetRequesterOk() (*string, bool)`

GetRequesterOk returns a tuple with the Requester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequester

`func (o *Ticket) SetRequester(v string)`

SetRequester sets Requester field to given value.

### HasRequester

`func (o *Ticket) HasRequester() bool`

HasRequester returns a boolean if a field has been set.

### GetRequesterEmail

`func (o *Ticket) GetRequesterEmail() string`

GetRequesterEmail returns the RequesterEmail field if non-nil, zero value otherwise.

### GetRequesterEmailOk

`func (o *Ticket) GetRequesterEmailOk() (*string, bool)`

GetRequesterEmailOk returns a tuple with the RequesterEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterEmail

`func (o *Ticket) SetRequesterEmail(v string)`

SetRequesterEmail sets RequesterEmail field to given value.

### HasRequesterEmail

`func (o *Ticket) HasRequesterEmail() bool`

HasRequesterEmail returns a boolean if a field has been set.

### GetStatus

`func (o *Ticket) GetStatus() TicketStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Ticket) GetStatusOk() (*TicketStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Ticket) SetStatus(v TicketStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Ticket) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubject

`func (o *Ticket) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *Ticket) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *Ticket) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetType

`func (o *Ticket) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Ticket) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Ticket) SetType(v string)`

SetType sets Type field to given value.


### GetUpdatedAt

`func (o *Ticket) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Ticket) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Ticket) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Ticket) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


