# WxlanTunnelSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApAsSessionId** | Pointer to **string** | if &#x60;use_ap_as_session_ids&#x60;&#x3D;&#x3D;&#x60;true&#x60;, only apmac is supported right now. This is the name WLAN should use for wxtunnel_remote_id | [optional] 
**Comment** | Pointer to **string** | optional, user-specified string for display purpose | [optional] 
**EnableCookie** | Pointer to **bool** |  | [optional] 
**Ethertype** | Pointer to [**WxlanTunnelSessionEthertype**](WxlanTunnelSessionEthertype.md) |  | [optional] 
**LocalSessionId** | Pointer to **int32** | 1-4294967295 | [optional] 
**Pseudo8021adEnabled** | Pointer to **bool** | optional. Enables the pseudo 802.1ad QinQ mode where the AP device drops the outer vlan tag (QinQ). This mode is useful when tunneling Mist AP’s to some aggregation routers. | [optional] [default to false]
**RemoteId** | Pointer to **string** | remote-id of the session, has to be unique in the same tunnel | [optional] 
**RemoteSessionId** | Pointer to **int32** | 1-4294967295 | [optional] 
**UseApAsSessionIds** | Pointer to **bool** | whether to use AP (last 4 bytes of MAC currently) as session ids | [optional] [default to false]

## Methods

### NewWxlanTunnelSession

`func NewWxlanTunnelSession() *WxlanTunnelSession`

NewWxlanTunnelSession instantiates a new WxlanTunnelSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWxlanTunnelSessionWithDefaults

`func NewWxlanTunnelSessionWithDefaults() *WxlanTunnelSession`

NewWxlanTunnelSessionWithDefaults instantiates a new WxlanTunnelSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApAsSessionId

`func (o *WxlanTunnelSession) GetApAsSessionId() string`

GetApAsSessionId returns the ApAsSessionId field if non-nil, zero value otherwise.

### GetApAsSessionIdOk

`func (o *WxlanTunnelSession) GetApAsSessionIdOk() (*string, bool)`

GetApAsSessionIdOk returns a tuple with the ApAsSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApAsSessionId

`func (o *WxlanTunnelSession) SetApAsSessionId(v string)`

SetApAsSessionId sets ApAsSessionId field to given value.

### HasApAsSessionId

`func (o *WxlanTunnelSession) HasApAsSessionId() bool`

HasApAsSessionId returns a boolean if a field has been set.

### GetComment

`func (o *WxlanTunnelSession) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *WxlanTunnelSession) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *WxlanTunnelSession) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *WxlanTunnelSession) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetEnableCookie

`func (o *WxlanTunnelSession) GetEnableCookie() bool`

GetEnableCookie returns the EnableCookie field if non-nil, zero value otherwise.

### GetEnableCookieOk

`func (o *WxlanTunnelSession) GetEnableCookieOk() (*bool, bool)`

GetEnableCookieOk returns a tuple with the EnableCookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCookie

`func (o *WxlanTunnelSession) SetEnableCookie(v bool)`

SetEnableCookie sets EnableCookie field to given value.

### HasEnableCookie

`func (o *WxlanTunnelSession) HasEnableCookie() bool`

HasEnableCookie returns a boolean if a field has been set.

### GetEthertype

`func (o *WxlanTunnelSession) GetEthertype() WxlanTunnelSessionEthertype`

GetEthertype returns the Ethertype field if non-nil, zero value otherwise.

### GetEthertypeOk

`func (o *WxlanTunnelSession) GetEthertypeOk() (*WxlanTunnelSessionEthertype, bool)`

GetEthertypeOk returns a tuple with the Ethertype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthertype

`func (o *WxlanTunnelSession) SetEthertype(v WxlanTunnelSessionEthertype)`

SetEthertype sets Ethertype field to given value.

### HasEthertype

`func (o *WxlanTunnelSession) HasEthertype() bool`

HasEthertype returns a boolean if a field has been set.

### GetLocalSessionId

`func (o *WxlanTunnelSession) GetLocalSessionId() int32`

GetLocalSessionId returns the LocalSessionId field if non-nil, zero value otherwise.

### GetLocalSessionIdOk

`func (o *WxlanTunnelSession) GetLocalSessionIdOk() (*int32, bool)`

GetLocalSessionIdOk returns a tuple with the LocalSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSessionId

`func (o *WxlanTunnelSession) SetLocalSessionId(v int32)`

SetLocalSessionId sets LocalSessionId field to given value.

### HasLocalSessionId

`func (o *WxlanTunnelSession) HasLocalSessionId() bool`

HasLocalSessionId returns a boolean if a field has been set.

### GetPseudo8021adEnabled

`func (o *WxlanTunnelSession) GetPseudo8021adEnabled() bool`

GetPseudo8021adEnabled returns the Pseudo8021adEnabled field if non-nil, zero value otherwise.

### GetPseudo8021adEnabledOk

`func (o *WxlanTunnelSession) GetPseudo8021adEnabledOk() (*bool, bool)`

GetPseudo8021adEnabledOk returns a tuple with the Pseudo8021adEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPseudo8021adEnabled

`func (o *WxlanTunnelSession) SetPseudo8021adEnabled(v bool)`

SetPseudo8021adEnabled sets Pseudo8021adEnabled field to given value.

### HasPseudo8021adEnabled

`func (o *WxlanTunnelSession) HasPseudo8021adEnabled() bool`

HasPseudo8021adEnabled returns a boolean if a field has been set.

### GetRemoteId

`func (o *WxlanTunnelSession) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *WxlanTunnelSession) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *WxlanTunnelSession) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *WxlanTunnelSession) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### GetRemoteSessionId

`func (o *WxlanTunnelSession) GetRemoteSessionId() int32`

GetRemoteSessionId returns the RemoteSessionId field if non-nil, zero value otherwise.

### GetRemoteSessionIdOk

`func (o *WxlanTunnelSession) GetRemoteSessionIdOk() (*int32, bool)`

GetRemoteSessionIdOk returns a tuple with the RemoteSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSessionId

`func (o *WxlanTunnelSession) SetRemoteSessionId(v int32)`

SetRemoteSessionId sets RemoteSessionId field to given value.

### HasRemoteSessionId

`func (o *WxlanTunnelSession) HasRemoteSessionId() bool`

HasRemoteSessionId returns a boolean if a field has been set.

### GetUseApAsSessionIds

`func (o *WxlanTunnelSession) GetUseApAsSessionIds() bool`

GetUseApAsSessionIds returns the UseApAsSessionIds field if non-nil, zero value otherwise.

### GetUseApAsSessionIdsOk

`func (o *WxlanTunnelSession) GetUseApAsSessionIdsOk() (*bool, bool)`

GetUseApAsSessionIdsOk returns a tuple with the UseApAsSessionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseApAsSessionIds

`func (o *WxlanTunnelSession) SetUseApAsSessionIds(v bool)`

SetUseApAsSessionIds sets UseApAsSessionIds field to given value.

### HasUseApAsSessionIds

`func (o *WxlanTunnelSession) HasUseApAsSessionIds() bool`

HasUseApAsSessionIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


