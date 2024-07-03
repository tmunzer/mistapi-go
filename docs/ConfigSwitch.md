# ConfigSwitch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApAffinityThreshold** | Pointer to **int32** | ap_affinity_threshold ap_affinity_threshold can be added as a field under site/setting. By default this value is set to 12. If the field is set in both site/setting and org/setting, the value from site/setting will be used. | [optional] [default to 10]
**CliBanner** | Pointer to **string** | Set Banners for switches. Allows markup formatting | [optional] 
**CliIdleTimeout** | Pointer to **int32** | Sets timeout for switches | [optional] 
**ConfigRevertTimer** | Pointer to **int32** | the rollback timer for commit confirmed | [optional] [default to 10]
**DhcpOptionFqdn** | Pointer to **bool** | Enable to provide the FQDN with DHCP option 81 | [optional] [default to false]
**LocalAccounts** | Pointer to [**map[string]ConfigSwitchLocalAccountsUser**](ConfigSwitchLocalAccountsUser.md) | Property key is the user name. For Local user authentication | [optional] 
**MxedgeProxyHost** | Pointer to **string** |  | [optional] 
**MxedgeProxyPort** | Pointer to **int32** |  | [optional] [default to 2222]
**ProtectRe** | Pointer to [**ProtectRe**](ProtectRe.md) |  | [optional] 
**Radius** | Pointer to [**ConfigSwitchRadius**](ConfigSwitchRadius.md) |  | [optional] 
**RootPassword** | Pointer to **string** |  | [optional] 
**Tacacs** | Pointer to [**Tacacs**](Tacacs.md) |  | [optional] 
**UseMxedgeProxy** | Pointer to **bool** | to use mxedge as proxy | [optional] 

## Methods

### NewConfigSwitch

`func NewConfigSwitch() *ConfigSwitch`

NewConfigSwitch instantiates a new ConfigSwitch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigSwitchWithDefaults

`func NewConfigSwitchWithDefaults() *ConfigSwitch`

NewConfigSwitchWithDefaults instantiates a new ConfigSwitch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApAffinityThreshold

`func (o *ConfigSwitch) GetApAffinityThreshold() int32`

GetApAffinityThreshold returns the ApAffinityThreshold field if non-nil, zero value otherwise.

### GetApAffinityThresholdOk

`func (o *ConfigSwitch) GetApAffinityThresholdOk() (*int32, bool)`

GetApAffinityThresholdOk returns a tuple with the ApAffinityThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApAffinityThreshold

`func (o *ConfigSwitch) SetApAffinityThreshold(v int32)`

SetApAffinityThreshold sets ApAffinityThreshold field to given value.

### HasApAffinityThreshold

`func (o *ConfigSwitch) HasApAffinityThreshold() bool`

HasApAffinityThreshold returns a boolean if a field has been set.

### GetCliBanner

`func (o *ConfigSwitch) GetCliBanner() string`

GetCliBanner returns the CliBanner field if non-nil, zero value otherwise.

### GetCliBannerOk

`func (o *ConfigSwitch) GetCliBannerOk() (*string, bool)`

GetCliBannerOk returns a tuple with the CliBanner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCliBanner

`func (o *ConfigSwitch) SetCliBanner(v string)`

SetCliBanner sets CliBanner field to given value.

### HasCliBanner

`func (o *ConfigSwitch) HasCliBanner() bool`

HasCliBanner returns a boolean if a field has been set.

### GetCliIdleTimeout

`func (o *ConfigSwitch) GetCliIdleTimeout() int32`

GetCliIdleTimeout returns the CliIdleTimeout field if non-nil, zero value otherwise.

### GetCliIdleTimeoutOk

`func (o *ConfigSwitch) GetCliIdleTimeoutOk() (*int32, bool)`

GetCliIdleTimeoutOk returns a tuple with the CliIdleTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCliIdleTimeout

`func (o *ConfigSwitch) SetCliIdleTimeout(v int32)`

SetCliIdleTimeout sets CliIdleTimeout field to given value.

### HasCliIdleTimeout

`func (o *ConfigSwitch) HasCliIdleTimeout() bool`

HasCliIdleTimeout returns a boolean if a field has been set.

### GetConfigRevertTimer

`func (o *ConfigSwitch) GetConfigRevertTimer() int32`

GetConfigRevertTimer returns the ConfigRevertTimer field if non-nil, zero value otherwise.

### GetConfigRevertTimerOk

`func (o *ConfigSwitch) GetConfigRevertTimerOk() (*int32, bool)`

GetConfigRevertTimerOk returns a tuple with the ConfigRevertTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigRevertTimer

`func (o *ConfigSwitch) SetConfigRevertTimer(v int32)`

SetConfigRevertTimer sets ConfigRevertTimer field to given value.

### HasConfigRevertTimer

`func (o *ConfigSwitch) HasConfigRevertTimer() bool`

HasConfigRevertTimer returns a boolean if a field has been set.

### GetDhcpOptionFqdn

`func (o *ConfigSwitch) GetDhcpOptionFqdn() bool`

GetDhcpOptionFqdn returns the DhcpOptionFqdn field if non-nil, zero value otherwise.

### GetDhcpOptionFqdnOk

`func (o *ConfigSwitch) GetDhcpOptionFqdnOk() (*bool, bool)`

GetDhcpOptionFqdnOk returns a tuple with the DhcpOptionFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpOptionFqdn

`func (o *ConfigSwitch) SetDhcpOptionFqdn(v bool)`

SetDhcpOptionFqdn sets DhcpOptionFqdn field to given value.

### HasDhcpOptionFqdn

`func (o *ConfigSwitch) HasDhcpOptionFqdn() bool`

HasDhcpOptionFqdn returns a boolean if a field has been set.

### GetLocalAccounts

`func (o *ConfigSwitch) GetLocalAccounts() map[string]ConfigSwitchLocalAccountsUser`

GetLocalAccounts returns the LocalAccounts field if non-nil, zero value otherwise.

### GetLocalAccountsOk

`func (o *ConfigSwitch) GetLocalAccountsOk() (*map[string]ConfigSwitchLocalAccountsUser, bool)`

GetLocalAccountsOk returns a tuple with the LocalAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAccounts

`func (o *ConfigSwitch) SetLocalAccounts(v map[string]ConfigSwitchLocalAccountsUser)`

SetLocalAccounts sets LocalAccounts field to given value.

### HasLocalAccounts

`func (o *ConfigSwitch) HasLocalAccounts() bool`

HasLocalAccounts returns a boolean if a field has been set.

### GetMxedgeProxyHost

`func (o *ConfigSwitch) GetMxedgeProxyHost() string`

GetMxedgeProxyHost returns the MxedgeProxyHost field if non-nil, zero value otherwise.

### GetMxedgeProxyHostOk

`func (o *ConfigSwitch) GetMxedgeProxyHostOk() (*string, bool)`

GetMxedgeProxyHostOk returns a tuple with the MxedgeProxyHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMxedgeProxyHost

`func (o *ConfigSwitch) SetMxedgeProxyHost(v string)`

SetMxedgeProxyHost sets MxedgeProxyHost field to given value.

### HasMxedgeProxyHost

`func (o *ConfigSwitch) HasMxedgeProxyHost() bool`

HasMxedgeProxyHost returns a boolean if a field has been set.

### GetMxedgeProxyPort

`func (o *ConfigSwitch) GetMxedgeProxyPort() int32`

GetMxedgeProxyPort returns the MxedgeProxyPort field if non-nil, zero value otherwise.

### GetMxedgeProxyPortOk

`func (o *ConfigSwitch) GetMxedgeProxyPortOk() (*int32, bool)`

GetMxedgeProxyPortOk returns a tuple with the MxedgeProxyPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMxedgeProxyPort

`func (o *ConfigSwitch) SetMxedgeProxyPort(v int32)`

SetMxedgeProxyPort sets MxedgeProxyPort field to given value.

### HasMxedgeProxyPort

`func (o *ConfigSwitch) HasMxedgeProxyPort() bool`

HasMxedgeProxyPort returns a boolean if a field has been set.

### GetProtectRe

`func (o *ConfigSwitch) GetProtectRe() ProtectRe`

GetProtectRe returns the ProtectRe field if non-nil, zero value otherwise.

### GetProtectReOk

`func (o *ConfigSwitch) GetProtectReOk() (*ProtectRe, bool)`

GetProtectReOk returns a tuple with the ProtectRe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectRe

`func (o *ConfigSwitch) SetProtectRe(v ProtectRe)`

SetProtectRe sets ProtectRe field to given value.

### HasProtectRe

`func (o *ConfigSwitch) HasProtectRe() bool`

HasProtectRe returns a boolean if a field has been set.

### GetRadius

`func (o *ConfigSwitch) GetRadius() ConfigSwitchRadius`

GetRadius returns the Radius field if non-nil, zero value otherwise.

### GetRadiusOk

`func (o *ConfigSwitch) GetRadiusOk() (*ConfigSwitchRadius, bool)`

GetRadiusOk returns a tuple with the Radius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadius

`func (o *ConfigSwitch) SetRadius(v ConfigSwitchRadius)`

SetRadius sets Radius field to given value.

### HasRadius

`func (o *ConfigSwitch) HasRadius() bool`

HasRadius returns a boolean if a field has been set.

### GetRootPassword

`func (o *ConfigSwitch) GetRootPassword() string`

GetRootPassword returns the RootPassword field if non-nil, zero value otherwise.

### GetRootPasswordOk

`func (o *ConfigSwitch) GetRootPasswordOk() (*string, bool)`

GetRootPasswordOk returns a tuple with the RootPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPassword

`func (o *ConfigSwitch) SetRootPassword(v string)`

SetRootPassword sets RootPassword field to given value.

### HasRootPassword

`func (o *ConfigSwitch) HasRootPassword() bool`

HasRootPassword returns a boolean if a field has been set.

### GetTacacs

`func (o *ConfigSwitch) GetTacacs() Tacacs`

GetTacacs returns the Tacacs field if non-nil, zero value otherwise.

### GetTacacsOk

`func (o *ConfigSwitch) GetTacacsOk() (*Tacacs, bool)`

GetTacacsOk returns a tuple with the Tacacs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTacacs

`func (o *ConfigSwitch) SetTacacs(v Tacacs)`

SetTacacs sets Tacacs field to given value.

### HasTacacs

`func (o *ConfigSwitch) HasTacacs() bool`

HasTacacs returns a boolean if a field has been set.

### GetUseMxedgeProxy

`func (o *ConfigSwitch) GetUseMxedgeProxy() bool`

GetUseMxedgeProxy returns the UseMxedgeProxy field if non-nil, zero value otherwise.

### GetUseMxedgeProxyOk

`func (o *ConfigSwitch) GetUseMxedgeProxyOk() (*bool, bool)`

GetUseMxedgeProxyOk returns a tuple with the UseMxedgeProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMxedgeProxy

`func (o *ConfigSwitch) SetUseMxedgeProxy(v bool)`

SetUseMxedgeProxy sets UseMxedgeProxy field to given value.

### HasUseMxedgeProxy

`func (o *ConfigSwitch) HasUseMxedgeProxy() bool`

HasUseMxedgeProxy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


