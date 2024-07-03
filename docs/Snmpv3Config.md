# Snmpv3Config

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Notify** | Pointer to [**[]Snmpv3ConfigNotifyItems**](Snmpv3ConfigNotifyItems.md) |  | [optional] 
**NotifyFilter** | Pointer to [**[]Snmpv3ConfigNotifyFilterItem**](Snmpv3ConfigNotifyFilterItem.md) |  | [optional] 
**TargetAddress** | Pointer to [**[]Snmpv3ConfigTargetAddressItem**](Snmpv3ConfigTargetAddressItem.md) |  | [optional] 
**TargetParameters** | Pointer to [**[]Snmpv3ConfigTargetParam**](Snmpv3ConfigTargetParam.md) |  | [optional] 
**Usm** | Pointer to [**SnmpUsm**](SnmpUsm.md) |  | [optional] 
**Vacm** | Pointer to [**SnmpVacm**](SnmpVacm.md) |  | [optional] 

## Methods

### NewSnmpv3Config

`func NewSnmpv3Config() *Snmpv3Config`

NewSnmpv3Config instantiates a new Snmpv3Config object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnmpv3ConfigWithDefaults

`func NewSnmpv3ConfigWithDefaults() *Snmpv3Config`

NewSnmpv3ConfigWithDefaults instantiates a new Snmpv3Config object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotify

`func (o *Snmpv3Config) GetNotify() []Snmpv3ConfigNotifyItems`

GetNotify returns the Notify field if non-nil, zero value otherwise.

### GetNotifyOk

`func (o *Snmpv3Config) GetNotifyOk() (*[]Snmpv3ConfigNotifyItems, bool)`

GetNotifyOk returns a tuple with the Notify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotify

`func (o *Snmpv3Config) SetNotify(v []Snmpv3ConfigNotifyItems)`

SetNotify sets Notify field to given value.

### HasNotify

`func (o *Snmpv3Config) HasNotify() bool`

HasNotify returns a boolean if a field has been set.

### GetNotifyFilter

`func (o *Snmpv3Config) GetNotifyFilter() []Snmpv3ConfigNotifyFilterItem`

GetNotifyFilter returns the NotifyFilter field if non-nil, zero value otherwise.

### GetNotifyFilterOk

`func (o *Snmpv3Config) GetNotifyFilterOk() (*[]Snmpv3ConfigNotifyFilterItem, bool)`

GetNotifyFilterOk returns a tuple with the NotifyFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyFilter

`func (o *Snmpv3Config) SetNotifyFilter(v []Snmpv3ConfigNotifyFilterItem)`

SetNotifyFilter sets NotifyFilter field to given value.

### HasNotifyFilter

`func (o *Snmpv3Config) HasNotifyFilter() bool`

HasNotifyFilter returns a boolean if a field has been set.

### GetTargetAddress

`func (o *Snmpv3Config) GetTargetAddress() []Snmpv3ConfigTargetAddressItem`

GetTargetAddress returns the TargetAddress field if non-nil, zero value otherwise.

### GetTargetAddressOk

`func (o *Snmpv3Config) GetTargetAddressOk() (*[]Snmpv3ConfigTargetAddressItem, bool)`

GetTargetAddressOk returns a tuple with the TargetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAddress

`func (o *Snmpv3Config) SetTargetAddress(v []Snmpv3ConfigTargetAddressItem)`

SetTargetAddress sets TargetAddress field to given value.

### HasTargetAddress

`func (o *Snmpv3Config) HasTargetAddress() bool`

HasTargetAddress returns a boolean if a field has been set.

### GetTargetParameters

`func (o *Snmpv3Config) GetTargetParameters() []Snmpv3ConfigTargetParam`

GetTargetParameters returns the TargetParameters field if non-nil, zero value otherwise.

### GetTargetParametersOk

`func (o *Snmpv3Config) GetTargetParametersOk() (*[]Snmpv3ConfigTargetParam, bool)`

GetTargetParametersOk returns a tuple with the TargetParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetParameters

`func (o *Snmpv3Config) SetTargetParameters(v []Snmpv3ConfigTargetParam)`

SetTargetParameters sets TargetParameters field to given value.

### HasTargetParameters

`func (o *Snmpv3Config) HasTargetParameters() bool`

HasTargetParameters returns a boolean if a field has been set.

### GetUsm

`func (o *Snmpv3Config) GetUsm() SnmpUsm`

GetUsm returns the Usm field if non-nil, zero value otherwise.

### GetUsmOk

`func (o *Snmpv3Config) GetUsmOk() (*SnmpUsm, bool)`

GetUsmOk returns a tuple with the Usm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsm

`func (o *Snmpv3Config) SetUsm(v SnmpUsm)`

SetUsm sets Usm field to given value.

### HasUsm

`func (o *Snmpv3Config) HasUsm() bool`

HasUsm returns a boolean if a field has been set.

### GetVacm

`func (o *Snmpv3Config) GetVacm() SnmpVacm`

GetVacm returns the Vacm field if non-nil, zero value otherwise.

### GetVacmOk

`func (o *Snmpv3Config) GetVacmOk() (*SnmpVacm, bool)`

GetVacmOk returns a tuple with the Vacm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVacm

`func (o *Snmpv3Config) SetVacm(v SnmpVacm)`

SetVacm sets Vacm field to given value.

### HasVacm

`func (o *Snmpv3Config) HasVacm() bool`

HasVacm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


