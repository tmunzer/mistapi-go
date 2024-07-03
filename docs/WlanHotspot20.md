# WlanHotspot20

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainName** | Pointer to **[]string** |  | [optional] 
**Enabled** | Pointer to **bool** | whether to enable hotspot 2.0 config | [optional] 
**NaiRealms** | Pointer to **[]string** |  | [optional] 
**Operators** | Pointer to [**[]WlanHotspot20OperatorsItem**](WlanHotspot20OperatorsItem.md) | list of operators to support | [optional] 
**Rcoi** | Pointer to **[]string** |  | [optional] 
**VenueName** | Pointer to **string** | venue name, default is site name | [optional] 

## Methods

### NewWlanHotspot20

`func NewWlanHotspot20() *WlanHotspot20`

NewWlanHotspot20 instantiates a new WlanHotspot20 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWlanHotspot20WithDefaults

`func NewWlanHotspot20WithDefaults() *WlanHotspot20`

NewWlanHotspot20WithDefaults instantiates a new WlanHotspot20 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainName

`func (o *WlanHotspot20) GetDomainName() []string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *WlanHotspot20) GetDomainNameOk() (*[]string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *WlanHotspot20) SetDomainName(v []string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *WlanHotspot20) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetEnabled

`func (o *WlanHotspot20) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WlanHotspot20) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WlanHotspot20) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *WlanHotspot20) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetNaiRealms

`func (o *WlanHotspot20) GetNaiRealms() []string`

GetNaiRealms returns the NaiRealms field if non-nil, zero value otherwise.

### GetNaiRealmsOk

`func (o *WlanHotspot20) GetNaiRealmsOk() (*[]string, bool)`

GetNaiRealmsOk returns a tuple with the NaiRealms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaiRealms

`func (o *WlanHotspot20) SetNaiRealms(v []string)`

SetNaiRealms sets NaiRealms field to given value.

### HasNaiRealms

`func (o *WlanHotspot20) HasNaiRealms() bool`

HasNaiRealms returns a boolean if a field has been set.

### GetOperators

`func (o *WlanHotspot20) GetOperators() []WlanHotspot20OperatorsItem`

GetOperators returns the Operators field if non-nil, zero value otherwise.

### GetOperatorsOk

`func (o *WlanHotspot20) GetOperatorsOk() (*[]WlanHotspot20OperatorsItem, bool)`

GetOperatorsOk returns a tuple with the Operators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperators

`func (o *WlanHotspot20) SetOperators(v []WlanHotspot20OperatorsItem)`

SetOperators sets Operators field to given value.

### HasOperators

`func (o *WlanHotspot20) HasOperators() bool`

HasOperators returns a boolean if a field has been set.

### GetRcoi

`func (o *WlanHotspot20) GetRcoi() []string`

GetRcoi returns the Rcoi field if non-nil, zero value otherwise.

### GetRcoiOk

`func (o *WlanHotspot20) GetRcoiOk() (*[]string, bool)`

GetRcoiOk returns a tuple with the Rcoi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRcoi

`func (o *WlanHotspot20) SetRcoi(v []string)`

SetRcoi sets Rcoi field to given value.

### HasRcoi

`func (o *WlanHotspot20) HasRcoi() bool`

HasRcoi returns a boolean if a field has been set.

### GetVenueName

`func (o *WlanHotspot20) GetVenueName() string`

GetVenueName returns the VenueName field if non-nil, zero value otherwise.

### GetVenueNameOk

`func (o *WlanHotspot20) GetVenueNameOk() (*string, bool)`

GetVenueNameOk returns a tuple with the VenueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenueName

`func (o *WlanHotspot20) SetVenueName(v string)`

SetVenueName sets VenueName field to given value.

### HasVenueName

`func (o *WlanHotspot20) HasVenueName() bool`

HasVenueName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


