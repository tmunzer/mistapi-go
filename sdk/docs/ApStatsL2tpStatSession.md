# ApStatsL2tpStatSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalSid** | Pointer to **int32** | remote sessions id (dynamically unless Tunnel is said to be static) | [optional] 
**RemoteId** | Pointer to **string** | WxlanTunnel Remote ID (user-configured) | [optional] 
**RemoteSid** | Pointer to **int32** | remote sessions id (dynamically unless Tunnel is said to be static) | [optional] 
**State** | Pointer to [**L2tpState**](L2tpState.md) |  | [optional] 

## Methods

### NewApStatsL2tpStatSession

`func NewApStatsL2tpStatSession() *ApStatsL2tpStatSession`

NewApStatsL2tpStatSession instantiates a new ApStatsL2tpStatSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApStatsL2tpStatSessionWithDefaults

`func NewApStatsL2tpStatSessionWithDefaults() *ApStatsL2tpStatSession`

NewApStatsL2tpStatSessionWithDefaults instantiates a new ApStatsL2tpStatSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocalSid

`func (o *ApStatsL2tpStatSession) GetLocalSid() int32`

GetLocalSid returns the LocalSid field if non-nil, zero value otherwise.

### GetLocalSidOk

`func (o *ApStatsL2tpStatSession) GetLocalSidOk() (*int32, bool)`

GetLocalSidOk returns a tuple with the LocalSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSid

`func (o *ApStatsL2tpStatSession) SetLocalSid(v int32)`

SetLocalSid sets LocalSid field to given value.

### HasLocalSid

`func (o *ApStatsL2tpStatSession) HasLocalSid() bool`

HasLocalSid returns a boolean if a field has been set.

### GetRemoteId

`func (o *ApStatsL2tpStatSession) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *ApStatsL2tpStatSession) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *ApStatsL2tpStatSession) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *ApStatsL2tpStatSession) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### GetRemoteSid

`func (o *ApStatsL2tpStatSession) GetRemoteSid() int32`

GetRemoteSid returns the RemoteSid field if non-nil, zero value otherwise.

### GetRemoteSidOk

`func (o *ApStatsL2tpStatSession) GetRemoteSidOk() (*int32, bool)`

GetRemoteSidOk returns a tuple with the RemoteSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSid

`func (o *ApStatsL2tpStatSession) SetRemoteSid(v int32)`

SetRemoteSid sets RemoteSid field to given value.

### HasRemoteSid

`func (o *ApStatsL2tpStatSession) HasRemoteSid() bool`

HasRemoteSid returns a boolean if a field has been set.

### GetState

`func (o *ApStatsL2tpStatSession) GetState() L2tpState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ApStatsL2tpStatSession) GetStateOk() (*L2tpState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ApStatsL2tpStatSession) SetState(v L2tpState)`

SetState sets State field to given value.

### HasState

`func (o *ApStatsL2tpStatSession) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


