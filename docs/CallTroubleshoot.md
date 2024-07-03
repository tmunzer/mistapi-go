# CallTroubleshoot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** |  | [optional] 
**MeetingId** | Pointer to **string** |  | [optional] 
**Results** | Pointer to [**[]TroubleshootCallItem**](TroubleshootCallItem.md) |  | [optional] 

## Methods

### NewCallTroubleshoot

`func NewCallTroubleshoot() *CallTroubleshoot`

NewCallTroubleshoot instantiates a new CallTroubleshoot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallTroubleshootWithDefaults

`func NewCallTroubleshootWithDefaults() *CallTroubleshoot`

NewCallTroubleshootWithDefaults instantiates a new CallTroubleshoot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *CallTroubleshoot) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *CallTroubleshoot) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *CallTroubleshoot) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *CallTroubleshoot) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMeetingId

`func (o *CallTroubleshoot) GetMeetingId() string`

GetMeetingId returns the MeetingId field if non-nil, zero value otherwise.

### GetMeetingIdOk

`func (o *CallTroubleshoot) GetMeetingIdOk() (*string, bool)`

GetMeetingIdOk returns a tuple with the MeetingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeetingId

`func (o *CallTroubleshoot) SetMeetingId(v string)`

SetMeetingId sets MeetingId field to given value.

### HasMeetingId

`func (o *CallTroubleshoot) HasMeetingId() bool`

HasMeetingId returns a boolean if a field has been set.

### GetResults

`func (o *CallTroubleshoot) GetResults() []TroubleshootCallItem`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *CallTroubleshoot) GetResultsOk() (*[]TroubleshootCallItem, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *CallTroubleshoot) SetResults(v []TroubleshootCallItem)`

SetResults sets Results field to given value.

### HasResults

`func (o *CallTroubleshoot) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


