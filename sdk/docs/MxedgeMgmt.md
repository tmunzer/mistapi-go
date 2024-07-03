# MxedgeMgmt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FipsEnabled** | Pointer to **bool** |  | [optional] [default to false]
**MistPassword** | Pointer to **string** |  | [optional] 
**OobIpType** | Pointer to [**MxedgeMgmtOobIpType**](MxedgeMgmtOobIpType.md) |  | [optional] [default to MXEDGEMGMTOOBIPTYPE_DHCP]
**OobIpType6** | Pointer to [**MxedgeMgmtOobIpType6**](MxedgeMgmtOobIpType6.md) |  | [optional] [default to MXEDGEMGMTOOBIPTYPE6_AUTOCONF]
**RootPassword** | Pointer to **string** |  | [optional] 

## Methods

### NewMxedgeMgmt

`func NewMxedgeMgmt() *MxedgeMgmt`

NewMxedgeMgmt instantiates a new MxedgeMgmt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMxedgeMgmtWithDefaults

`func NewMxedgeMgmtWithDefaults() *MxedgeMgmt`

NewMxedgeMgmtWithDefaults instantiates a new MxedgeMgmt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFipsEnabled

`func (o *MxedgeMgmt) GetFipsEnabled() bool`

GetFipsEnabled returns the FipsEnabled field if non-nil, zero value otherwise.

### GetFipsEnabledOk

`func (o *MxedgeMgmt) GetFipsEnabledOk() (*bool, bool)`

GetFipsEnabledOk returns a tuple with the FipsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFipsEnabled

`func (o *MxedgeMgmt) SetFipsEnabled(v bool)`

SetFipsEnabled sets FipsEnabled field to given value.

### HasFipsEnabled

`func (o *MxedgeMgmt) HasFipsEnabled() bool`

HasFipsEnabled returns a boolean if a field has been set.

### GetMistPassword

`func (o *MxedgeMgmt) GetMistPassword() string`

GetMistPassword returns the MistPassword field if non-nil, zero value otherwise.

### GetMistPasswordOk

`func (o *MxedgeMgmt) GetMistPasswordOk() (*string, bool)`

GetMistPasswordOk returns a tuple with the MistPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMistPassword

`func (o *MxedgeMgmt) SetMistPassword(v string)`

SetMistPassword sets MistPassword field to given value.

### HasMistPassword

`func (o *MxedgeMgmt) HasMistPassword() bool`

HasMistPassword returns a boolean if a field has been set.

### GetOobIpType

`func (o *MxedgeMgmt) GetOobIpType() MxedgeMgmtOobIpType`

GetOobIpType returns the OobIpType field if non-nil, zero value otherwise.

### GetOobIpTypeOk

`func (o *MxedgeMgmt) GetOobIpTypeOk() (*MxedgeMgmtOobIpType, bool)`

GetOobIpTypeOk returns a tuple with the OobIpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOobIpType

`func (o *MxedgeMgmt) SetOobIpType(v MxedgeMgmtOobIpType)`

SetOobIpType sets OobIpType field to given value.

### HasOobIpType

`func (o *MxedgeMgmt) HasOobIpType() bool`

HasOobIpType returns a boolean if a field has been set.

### GetOobIpType6

`func (o *MxedgeMgmt) GetOobIpType6() MxedgeMgmtOobIpType6`

GetOobIpType6 returns the OobIpType6 field if non-nil, zero value otherwise.

### GetOobIpType6Ok

`func (o *MxedgeMgmt) GetOobIpType6Ok() (*MxedgeMgmtOobIpType6, bool)`

GetOobIpType6Ok returns a tuple with the OobIpType6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOobIpType6

`func (o *MxedgeMgmt) SetOobIpType6(v MxedgeMgmtOobIpType6)`

SetOobIpType6 sets OobIpType6 field to given value.

### HasOobIpType6

`func (o *MxedgeMgmt) HasOobIpType6() bool`

HasOobIpType6 returns a boolean if a field has been set.

### GetRootPassword

`func (o *MxedgeMgmt) GetRootPassword() string`

GetRootPassword returns the RootPassword field if non-nil, zero value otherwise.

### GetRootPasswordOk

`func (o *MxedgeMgmt) GetRootPasswordOk() (*string, bool)`

GetRootPasswordOk returns a tuple with the RootPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPassword

`func (o *MxedgeMgmt) SetRootPassword(v string)`

SetRootPassword sets RootPassword field to given value.

### HasRootPassword

`func (o *MxedgeMgmt) HasRootPassword() bool`

HasRootPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


