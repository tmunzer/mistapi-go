# SiteSettingGatewayMgmt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminSshkeys** | Pointer to **[]string** | for SSR only, as direct root access is not allowed | [optional] 
**AppProbing** | Pointer to [**AppProbing**](AppProbing.md) |  | [optional] 
**AppUsage** | Pointer to **bool** | consumes uplink bandwidth, requires WA license | [optional] 
**AutoSignatureUpdate** | Pointer to [**SiteSettingGatewayMgmtAutoSignatureUpdate**](SiteSettingGatewayMgmtAutoSignatureUpdate.md) |  | [optional] 
**ConfigRevertTimer** | Pointer to **float32** | he rollback timer for commit confirmed | [optional] [default to 10]
**ProbeHosts** | Pointer to **[]string** |  | [optional] 
**RootPassword** | Pointer to **string** | for SRX only | [optional] 
**SecurityLogSourceAddress** | Pointer to **string** |  | [optional] 
**SecurityLogSourceInterface** | Pointer to **string** |  | [optional] 

## Methods

### NewSiteSettingGatewayMgmt

`func NewSiteSettingGatewayMgmt() *SiteSettingGatewayMgmt`

NewSiteSettingGatewayMgmt instantiates a new SiteSettingGatewayMgmt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteSettingGatewayMgmtWithDefaults

`func NewSiteSettingGatewayMgmtWithDefaults() *SiteSettingGatewayMgmt`

NewSiteSettingGatewayMgmtWithDefaults instantiates a new SiteSettingGatewayMgmt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminSshkeys

`func (o *SiteSettingGatewayMgmt) GetAdminSshkeys() []string`

GetAdminSshkeys returns the AdminSshkeys field if non-nil, zero value otherwise.

### GetAdminSshkeysOk

`func (o *SiteSettingGatewayMgmt) GetAdminSshkeysOk() (*[]string, bool)`

GetAdminSshkeysOk returns a tuple with the AdminSshkeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminSshkeys

`func (o *SiteSettingGatewayMgmt) SetAdminSshkeys(v []string)`

SetAdminSshkeys sets AdminSshkeys field to given value.

### HasAdminSshkeys

`func (o *SiteSettingGatewayMgmt) HasAdminSshkeys() bool`

HasAdminSshkeys returns a boolean if a field has been set.

### GetAppProbing

`func (o *SiteSettingGatewayMgmt) GetAppProbing() AppProbing`

GetAppProbing returns the AppProbing field if non-nil, zero value otherwise.

### GetAppProbingOk

`func (o *SiteSettingGatewayMgmt) GetAppProbingOk() (*AppProbing, bool)`

GetAppProbingOk returns a tuple with the AppProbing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppProbing

`func (o *SiteSettingGatewayMgmt) SetAppProbing(v AppProbing)`

SetAppProbing sets AppProbing field to given value.

### HasAppProbing

`func (o *SiteSettingGatewayMgmt) HasAppProbing() bool`

HasAppProbing returns a boolean if a field has been set.

### GetAppUsage

`func (o *SiteSettingGatewayMgmt) GetAppUsage() bool`

GetAppUsage returns the AppUsage field if non-nil, zero value otherwise.

### GetAppUsageOk

`func (o *SiteSettingGatewayMgmt) GetAppUsageOk() (*bool, bool)`

GetAppUsageOk returns a tuple with the AppUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppUsage

`func (o *SiteSettingGatewayMgmt) SetAppUsage(v bool)`

SetAppUsage sets AppUsage field to given value.

### HasAppUsage

`func (o *SiteSettingGatewayMgmt) HasAppUsage() bool`

HasAppUsage returns a boolean if a field has been set.

### GetAutoSignatureUpdate

`func (o *SiteSettingGatewayMgmt) GetAutoSignatureUpdate() SiteSettingGatewayMgmtAutoSignatureUpdate`

GetAutoSignatureUpdate returns the AutoSignatureUpdate field if non-nil, zero value otherwise.

### GetAutoSignatureUpdateOk

`func (o *SiteSettingGatewayMgmt) GetAutoSignatureUpdateOk() (*SiteSettingGatewayMgmtAutoSignatureUpdate, bool)`

GetAutoSignatureUpdateOk returns a tuple with the AutoSignatureUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSignatureUpdate

`func (o *SiteSettingGatewayMgmt) SetAutoSignatureUpdate(v SiteSettingGatewayMgmtAutoSignatureUpdate)`

SetAutoSignatureUpdate sets AutoSignatureUpdate field to given value.

### HasAutoSignatureUpdate

`func (o *SiteSettingGatewayMgmt) HasAutoSignatureUpdate() bool`

HasAutoSignatureUpdate returns a boolean if a field has been set.

### GetConfigRevertTimer

`func (o *SiteSettingGatewayMgmt) GetConfigRevertTimer() float32`

GetConfigRevertTimer returns the ConfigRevertTimer field if non-nil, zero value otherwise.

### GetConfigRevertTimerOk

`func (o *SiteSettingGatewayMgmt) GetConfigRevertTimerOk() (*float32, bool)`

GetConfigRevertTimerOk returns a tuple with the ConfigRevertTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigRevertTimer

`func (o *SiteSettingGatewayMgmt) SetConfigRevertTimer(v float32)`

SetConfigRevertTimer sets ConfigRevertTimer field to given value.

### HasConfigRevertTimer

`func (o *SiteSettingGatewayMgmt) HasConfigRevertTimer() bool`

HasConfigRevertTimer returns a boolean if a field has been set.

### GetProbeHosts

`func (o *SiteSettingGatewayMgmt) GetProbeHosts() []string`

GetProbeHosts returns the ProbeHosts field if non-nil, zero value otherwise.

### GetProbeHostsOk

`func (o *SiteSettingGatewayMgmt) GetProbeHostsOk() (*[]string, bool)`

GetProbeHostsOk returns a tuple with the ProbeHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbeHosts

`func (o *SiteSettingGatewayMgmt) SetProbeHosts(v []string)`

SetProbeHosts sets ProbeHosts field to given value.

### HasProbeHosts

`func (o *SiteSettingGatewayMgmt) HasProbeHosts() bool`

HasProbeHosts returns a boolean if a field has been set.

### GetRootPassword

`func (o *SiteSettingGatewayMgmt) GetRootPassword() string`

GetRootPassword returns the RootPassword field if non-nil, zero value otherwise.

### GetRootPasswordOk

`func (o *SiteSettingGatewayMgmt) GetRootPasswordOk() (*string, bool)`

GetRootPasswordOk returns a tuple with the RootPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPassword

`func (o *SiteSettingGatewayMgmt) SetRootPassword(v string)`

SetRootPassword sets RootPassword field to given value.

### HasRootPassword

`func (o *SiteSettingGatewayMgmt) HasRootPassword() bool`

HasRootPassword returns a boolean if a field has been set.

### GetSecurityLogSourceAddress

`func (o *SiteSettingGatewayMgmt) GetSecurityLogSourceAddress() string`

GetSecurityLogSourceAddress returns the SecurityLogSourceAddress field if non-nil, zero value otherwise.

### GetSecurityLogSourceAddressOk

`func (o *SiteSettingGatewayMgmt) GetSecurityLogSourceAddressOk() (*string, bool)`

GetSecurityLogSourceAddressOk returns a tuple with the SecurityLogSourceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityLogSourceAddress

`func (o *SiteSettingGatewayMgmt) SetSecurityLogSourceAddress(v string)`

SetSecurityLogSourceAddress sets SecurityLogSourceAddress field to given value.

### HasSecurityLogSourceAddress

`func (o *SiteSettingGatewayMgmt) HasSecurityLogSourceAddress() bool`

HasSecurityLogSourceAddress returns a boolean if a field has been set.

### GetSecurityLogSourceInterface

`func (o *SiteSettingGatewayMgmt) GetSecurityLogSourceInterface() string`

GetSecurityLogSourceInterface returns the SecurityLogSourceInterface field if non-nil, zero value otherwise.

### GetSecurityLogSourceInterfaceOk

`func (o *SiteSettingGatewayMgmt) GetSecurityLogSourceInterfaceOk() (*string, bool)`

GetSecurityLogSourceInterfaceOk returns a tuple with the SecurityLogSourceInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityLogSourceInterface

`func (o *SiteSettingGatewayMgmt) SetSecurityLogSourceInterface(v string)`

SetSecurityLogSourceInterface sets SecurityLogSourceInterface field to given value.

### HasSecurityLogSourceInterface

`func (o *SiteSettingGatewayMgmt) HasSecurityLogSourceInterface() bool`

HasSecurityLogSourceInterface returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


