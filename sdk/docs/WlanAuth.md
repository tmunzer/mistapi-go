# WlanAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnticlogThreshold** | Pointer to **int32** | SAE anti-clogging token threshold | [optional] [default to 16]
**EapReauth** | Pointer to **bool** | whether to trigger EAP reauth when the session ends | [optional] [default to false]
**EnableMacAuth** | Pointer to **bool** | whether to enable MAC Auth, uses the same auth_servers | [optional] [default to false]
**KeyIdx** | Pointer to **int32** | when type&#x3D;wep | [optional] [default to 1]
**Keys** | Pointer to **[]string** | when type&#x3D;wep, four 10-character or 26-character hex string, null can be used. All keys, if provided, have to be in the same length | [optional] [default to []]
**MultiPskOnly** | Pointer to **bool** | whether to only use multi_psk | [optional] [default to false]
**Owe** | Pointer to [**WlanAuthOwe**](WlanAuthOwe.md) |  | [optional] [default to WLANAUTHOWE_DISABLED]
**Pairwise** | Pointer to [**[]WlanAuthPairwiseItem**](WlanAuthPairwiseItem.md) | when type&#x3D;psk / eap, one or more of wpa2-ccmp / wpa1-tkip / wpa1-ccmp / wpa2-tkip | [optional] [default to ["wpa2-ccmp"]]
**PrivateWlan** | Pointer to **bool** | whether private wlan is enabled. only applicable to multi_psk mode | [optional] [default to false]
**Psk** | Pointer to **NullableString** | when type&#x3D;psk, 8-64 characters, or 64 hex characters | [optional] [default to ""]
**Type** | [**WlanAuthType**](WlanAuthType.md) |  | [default to WLANAUTHTYPE_OPEN]
**WepAsSecondaryAuth** | Pointer to **bool** | enable WEP as secondary auth | [optional] [default to false]

## Methods

### NewWlanAuth

`func NewWlanAuth(type_ WlanAuthType, ) *WlanAuth`

NewWlanAuth instantiates a new WlanAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWlanAuthWithDefaults

`func NewWlanAuthWithDefaults() *WlanAuth`

NewWlanAuthWithDefaults instantiates a new WlanAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnticlogThreshold

`func (o *WlanAuth) GetAnticlogThreshold() int32`

GetAnticlogThreshold returns the AnticlogThreshold field if non-nil, zero value otherwise.

### GetAnticlogThresholdOk

`func (o *WlanAuth) GetAnticlogThresholdOk() (*int32, bool)`

GetAnticlogThresholdOk returns a tuple with the AnticlogThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnticlogThreshold

`func (o *WlanAuth) SetAnticlogThreshold(v int32)`

SetAnticlogThreshold sets AnticlogThreshold field to given value.

### HasAnticlogThreshold

`func (o *WlanAuth) HasAnticlogThreshold() bool`

HasAnticlogThreshold returns a boolean if a field has been set.

### GetEapReauth

`func (o *WlanAuth) GetEapReauth() bool`

GetEapReauth returns the EapReauth field if non-nil, zero value otherwise.

### GetEapReauthOk

`func (o *WlanAuth) GetEapReauthOk() (*bool, bool)`

GetEapReauthOk returns a tuple with the EapReauth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEapReauth

`func (o *WlanAuth) SetEapReauth(v bool)`

SetEapReauth sets EapReauth field to given value.

### HasEapReauth

`func (o *WlanAuth) HasEapReauth() bool`

HasEapReauth returns a boolean if a field has been set.

### GetEnableMacAuth

`func (o *WlanAuth) GetEnableMacAuth() bool`

GetEnableMacAuth returns the EnableMacAuth field if non-nil, zero value otherwise.

### GetEnableMacAuthOk

`func (o *WlanAuth) GetEnableMacAuthOk() (*bool, bool)`

GetEnableMacAuthOk returns a tuple with the EnableMacAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMacAuth

`func (o *WlanAuth) SetEnableMacAuth(v bool)`

SetEnableMacAuth sets EnableMacAuth field to given value.

### HasEnableMacAuth

`func (o *WlanAuth) HasEnableMacAuth() bool`

HasEnableMacAuth returns a boolean if a field has been set.

### GetKeyIdx

`func (o *WlanAuth) GetKeyIdx() int32`

GetKeyIdx returns the KeyIdx field if non-nil, zero value otherwise.

### GetKeyIdxOk

`func (o *WlanAuth) GetKeyIdxOk() (*int32, bool)`

GetKeyIdxOk returns a tuple with the KeyIdx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyIdx

`func (o *WlanAuth) SetKeyIdx(v int32)`

SetKeyIdx sets KeyIdx field to given value.

### HasKeyIdx

`func (o *WlanAuth) HasKeyIdx() bool`

HasKeyIdx returns a boolean if a field has been set.

### GetKeys

`func (o *WlanAuth) GetKeys() []*string`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *WlanAuth) GetKeysOk() (*[]*string, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *WlanAuth) SetKeys(v []*string)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *WlanAuth) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### GetMultiPskOnly

`func (o *WlanAuth) GetMultiPskOnly() bool`

GetMultiPskOnly returns the MultiPskOnly field if non-nil, zero value otherwise.

### GetMultiPskOnlyOk

`func (o *WlanAuth) GetMultiPskOnlyOk() (*bool, bool)`

GetMultiPskOnlyOk returns a tuple with the MultiPskOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiPskOnly

`func (o *WlanAuth) SetMultiPskOnly(v bool)`

SetMultiPskOnly sets MultiPskOnly field to given value.

### HasMultiPskOnly

`func (o *WlanAuth) HasMultiPskOnly() bool`

HasMultiPskOnly returns a boolean if a field has been set.

### GetOwe

`func (o *WlanAuth) GetOwe() WlanAuthOwe`

GetOwe returns the Owe field if non-nil, zero value otherwise.

### GetOweOk

`func (o *WlanAuth) GetOweOk() (*WlanAuthOwe, bool)`

GetOweOk returns a tuple with the Owe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwe

`func (o *WlanAuth) SetOwe(v WlanAuthOwe)`

SetOwe sets Owe field to given value.

### HasOwe

`func (o *WlanAuth) HasOwe() bool`

HasOwe returns a boolean if a field has been set.

### GetPairwise

`func (o *WlanAuth) GetPairwise() []WlanAuthPairwiseItem`

GetPairwise returns the Pairwise field if non-nil, zero value otherwise.

### GetPairwiseOk

`func (o *WlanAuth) GetPairwiseOk() (*[]WlanAuthPairwiseItem, bool)`

GetPairwiseOk returns a tuple with the Pairwise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairwise

`func (o *WlanAuth) SetPairwise(v []WlanAuthPairwiseItem)`

SetPairwise sets Pairwise field to given value.

### HasPairwise

`func (o *WlanAuth) HasPairwise() bool`

HasPairwise returns a boolean if a field has been set.

### GetPrivateWlan

`func (o *WlanAuth) GetPrivateWlan() bool`

GetPrivateWlan returns the PrivateWlan field if non-nil, zero value otherwise.

### GetPrivateWlanOk

`func (o *WlanAuth) GetPrivateWlanOk() (*bool, bool)`

GetPrivateWlanOk returns a tuple with the PrivateWlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateWlan

`func (o *WlanAuth) SetPrivateWlan(v bool)`

SetPrivateWlan sets PrivateWlan field to given value.

### HasPrivateWlan

`func (o *WlanAuth) HasPrivateWlan() bool`

HasPrivateWlan returns a boolean if a field has been set.

### GetPsk

`func (o *WlanAuth) GetPsk() string`

GetPsk returns the Psk field if non-nil, zero value otherwise.

### GetPskOk

`func (o *WlanAuth) GetPskOk() (*string, bool)`

GetPskOk returns a tuple with the Psk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsk

`func (o *WlanAuth) SetPsk(v string)`

SetPsk sets Psk field to given value.

### HasPsk

`func (o *WlanAuth) HasPsk() bool`

HasPsk returns a boolean if a field has been set.

### SetPskNil

`func (o *WlanAuth) SetPskNil(b bool)`

 SetPskNil sets the value for Psk to be an explicit nil

### UnsetPsk
`func (o *WlanAuth) UnsetPsk()`

UnsetPsk ensures that no value is present for Psk, not even an explicit nil
### GetType

`func (o *WlanAuth) GetType() WlanAuthType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WlanAuth) GetTypeOk() (*WlanAuthType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WlanAuth) SetType(v WlanAuthType)`

SetType sets Type field to given value.


### GetWepAsSecondaryAuth

`func (o *WlanAuth) GetWepAsSecondaryAuth() bool`

GetWepAsSecondaryAuth returns the WepAsSecondaryAuth field if non-nil, zero value otherwise.

### GetWepAsSecondaryAuthOk

`func (o *WlanAuth) GetWepAsSecondaryAuthOk() (*bool, bool)`

GetWepAsSecondaryAuthOk returns a tuple with the WepAsSecondaryAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWepAsSecondaryAuth

`func (o *WlanAuth) SetWepAsSecondaryAuth(v bool)`

SetWepAsSecondaryAuth sets WepAsSecondaryAuth field to given value.

### HasWepAsSecondaryAuth

`func (o *WlanAuth) HasWepAsSecondaryAuth() bool`

HasWepAsSecondaryAuth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


