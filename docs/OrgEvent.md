# OrgEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**Text** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **float32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewOrgEvent

`func NewOrgEvent() *OrgEvent`

NewOrgEvent instantiates a new OrgEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgEventWithDefaults

`func NewOrgEventWithDefaults() *OrgEvent`

NewOrgEventWithDefaults instantiates a new OrgEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *OrgEvent) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *OrgEvent) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *OrgEvent) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *OrgEvent) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetText

`func (o *OrgEvent) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *OrgEvent) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *OrgEvent) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *OrgEvent) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTimestamp

`func (o *OrgEvent) GetTimestamp() float32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *OrgEvent) GetTimestampOk() (*float32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *OrgEvent) SetTimestamp(v float32)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *OrgEvent) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetType

`func (o *OrgEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrgEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrgEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OrgEvent) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


