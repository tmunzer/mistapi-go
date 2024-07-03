# ProtectRe

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedServices** | Pointer to **[]string** | optionally, services we&#39;ll allow | [optional] 
**Custom** | Pointer to [**[]ProtectReCustom**](ProtectReCustom.md) |  | [optional] 
**Enabled** | Pointer to **bool** | when enabled, all traffic that is not essential to our operation will be dropped e.g. ntp / dns / traffic to mist will be allowed by default      if dhcpd is enabled, we&#39;ll make sure it works | [optional] [default to false]
**TrustedHosts** | Pointer to **[]string** | host/subnets we&#39;ll allow traffic to/from | [optional] 

## Methods

### NewProtectRe

`func NewProtectRe() *ProtectRe`

NewProtectRe instantiates a new ProtectRe object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectReWithDefaults

`func NewProtectReWithDefaults() *ProtectRe`

NewProtectReWithDefaults instantiates a new ProtectRe object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedServices

`func (o *ProtectRe) GetAllowedServices() []string`

GetAllowedServices returns the AllowedServices field if non-nil, zero value otherwise.

### GetAllowedServicesOk

`func (o *ProtectRe) GetAllowedServicesOk() (*[]string, bool)`

GetAllowedServicesOk returns a tuple with the AllowedServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedServices

`func (o *ProtectRe) SetAllowedServices(v []string)`

SetAllowedServices sets AllowedServices field to given value.

### HasAllowedServices

`func (o *ProtectRe) HasAllowedServices() bool`

HasAllowedServices returns a boolean if a field has been set.

### GetCustom

`func (o *ProtectRe) GetCustom() []ProtectReCustom`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *ProtectRe) GetCustomOk() (*[]ProtectReCustom, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *ProtectRe) SetCustom(v []ProtectReCustom)`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *ProtectRe) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetEnabled

`func (o *ProtectRe) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ProtectRe) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ProtectRe) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ProtectRe) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTrustedHosts

`func (o *ProtectRe) GetTrustedHosts() []string`

GetTrustedHosts returns the TrustedHosts field if non-nil, zero value otherwise.

### GetTrustedHostsOk

`func (o *ProtectRe) GetTrustedHostsOk() (*[]string, bool)`

GetTrustedHostsOk returns a tuple with the TrustedHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedHosts

`func (o *ProtectRe) SetTrustedHosts(v []string)`

SetTrustedHosts sets TrustedHosts field to given value.

### HasTrustedHosts

`func (o *ProtectRe) HasTrustedHosts() bool`

HasTrustedHosts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


