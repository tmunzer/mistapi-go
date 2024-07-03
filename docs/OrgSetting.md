# OrgSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApUpdownThreshold** | Pointer to **NullableInt32** | enable threshold-based device down delivery for AP devices only. When configured it takes effect for AP devices and &#x60;device_updown_threshold&#x60; is ignored. | [optional] 
**ApiPolicy** | Pointer to [**OrgSettingApiPolicy**](OrgSettingApiPolicy.md) |  | [optional] 
**AutoDeviceNaming** | Pointer to [**OrgSettingAutoDeviceNaming**](OrgSettingAutoDeviceNaming.md) |  | [optional] 
**AutoDeviceprofileAssignment** | Pointer to [**OrgSettingAutoDeviceprofileAssignment**](OrgSettingAutoDeviceprofileAssignment.md) |  | [optional] 
**AutoSiteAssignment** | Pointer to [**OrgSettingAutoSiteAssignment**](OrgSettingAutoSiteAssignment.md) |  | [optional] 
**BlacklistUrl** | Pointer to **string** |  | [optional] [readonly] 
**Cacerts** | Pointer to **[]string** | list of PEM-encoded ca certs | [optional] 
**Celona** | Pointer to [**OrgSettingCelona**](OrgSettingCelona.md) |  | [optional] 
**Cloudshark** | Pointer to [**OrgSettingCloudshark**](OrgSettingCloudshark.md) |  | [optional] 
**Cradlepoint** | Pointer to [**OrgSettingCradlepoint**](OrgSettingCradlepoint.md) |  | [optional] 
**CreatedTime** | Pointer to **float32** |  | [optional] [readonly] 
**DeviceCert** | Pointer to [**OrgSettingDeviceCert**](OrgSettingDeviceCert.md) |  | [optional] 
**DeviceUpdownThreshold** | Pointer to **int32** | enable threshold-based device down delivery via  1) device-updowns webhooks topic,  2) Mist Alert Framework; e.g. send AP/SW/GW down event only if AP/SW/GW Up is not seen within the threshold in minutes; 0 - 240, default is 0 (trigger immediate) | [optional] [default to 0]
**DisablePcap** | Pointer to **bool** | whether to disallow Mist to analyze pcap files (this is required for marvis pcap) | [optional] [default to false]
**DisableRemoteShell** | Pointer to **bool** | whether to disable remote shell access for an entire org | [optional] [default to false]
**ForSite** | Pointer to **bool** |  | [optional] [readonly] 
**GatewayMgmt** | Pointer to [**OrgSettingGatewayMgmt**](OrgSettingGatewayMgmt.md) |  | [optional] 
**GatewayUpdownThreshold** | Pointer to **NullableInt32** | enable threshold-based device down delivery for Gateway devices only. When configured it takes effect for GW devices and &#x60;device_updown_threshold&#x60; is ignored. | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Installer** | Pointer to [**OrgSettingInstaller**](OrgSettingInstaller.md) |  | [optional] 
**Jcloud** | Pointer to [**OrgSettingJcloud**](OrgSettingJcloud.md) |  | [optional] 
**Juniper** | Pointer to [**AccountJuniperInfo**](AccountJuniperInfo.md) |  | [optional] 
**Mgmt** | Pointer to [**OrgSettingMgmt**](OrgSettingMgmt.md) |  | [optional] 
**MistNac** | Pointer to [**OrgSettingMistNac**](OrgSettingMistNac.md) |  | [optional] 
**ModifiedTime** | Pointer to **float32** |  | [optional] [readonly] 
**MspId** | Pointer to **string** |  | [optional] [readonly] 
**MxedgeFipsEnabled** | Pointer to **bool** |  | [optional] [default to false]
**MxedgeMgmt** | Pointer to [**MxedgeMgmt**](MxedgeMgmt.md) |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**PasswordPolicy** | Pointer to [**OrgSettingPasswordPolicy**](OrgSettingPasswordPolicy.md) |  | [optional] 
**Pcap** | Pointer to [**OrgSettingPcap**](OrgSettingPcap.md) |  | [optional] 
**PcapBucketVerified** | Pointer to **bool** |  | [optional] 
**Security** | Pointer to [**OrgSettingSecurity**](OrgSettingSecurity.md) |  | [optional] 
**SimpleAlert** | Pointer to [**SimpleAlert**](SimpleAlert.md) |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] [readonly] 
**SwitchMgmt** | Pointer to [**OrgSettingSwitchMgmt**](OrgSettingSwitchMgmt.md) |  | [optional] 
**SwitchUpdownThreshold** | Pointer to **NullableInt32** | enable threshold-based device down delivery for Switch devices only. When configured it takes effect for SW devices and &#x60;device_updown_threshold&#x60; is ignored. | [optional] 
**SyntheticTest** | Pointer to [**SyntheticTestConfig**](SyntheticTestConfig.md) |  | [optional] 
**Tags** | Pointer to **[]string** | list of tags | [optional] 
**UiIdleTimeout** | Pointer to **int32** | automatically logout the user when UI session is inactive. &#x60;0&#x60; means disabled | [optional] [default to 0]
**VpnOptions** | Pointer to [**OrgSettingVpnOptions**](OrgSettingVpnOptions.md) |  | [optional] 
**WanPma** | Pointer to [**OrgSettingWanPma**](OrgSettingWanPma.md) |  | [optional] 
**WiredPma** | Pointer to [**OrgSettingWiredPma**](OrgSettingWiredPma.md) |  | [optional] 
**WirelessPma** | Pointer to [**OrgSettingWirelessPma**](OrgSettingWirelessPma.md) |  | [optional] 

## Methods

### NewOrgSetting

`func NewOrgSetting() *OrgSetting`

NewOrgSetting instantiates a new OrgSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgSettingWithDefaults

`func NewOrgSettingWithDefaults() *OrgSetting`

NewOrgSettingWithDefaults instantiates a new OrgSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApUpdownThreshold

`func (o *OrgSetting) GetApUpdownThreshold() int32`

GetApUpdownThreshold returns the ApUpdownThreshold field if non-nil, zero value otherwise.

### GetApUpdownThresholdOk

`func (o *OrgSetting) GetApUpdownThresholdOk() (*int32, bool)`

GetApUpdownThresholdOk returns a tuple with the ApUpdownThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApUpdownThreshold

`func (o *OrgSetting) SetApUpdownThreshold(v int32)`

SetApUpdownThreshold sets ApUpdownThreshold field to given value.

### HasApUpdownThreshold

`func (o *OrgSetting) HasApUpdownThreshold() bool`

HasApUpdownThreshold returns a boolean if a field has been set.

### SetApUpdownThresholdNil

`func (o *OrgSetting) SetApUpdownThresholdNil(b bool)`

 SetApUpdownThresholdNil sets the value for ApUpdownThreshold to be an explicit nil

### UnsetApUpdownThreshold
`func (o *OrgSetting) UnsetApUpdownThreshold()`

UnsetApUpdownThreshold ensures that no value is present for ApUpdownThreshold, not even an explicit nil
### GetApiPolicy

`func (o *OrgSetting) GetApiPolicy() OrgSettingApiPolicy`

GetApiPolicy returns the ApiPolicy field if non-nil, zero value otherwise.

### GetApiPolicyOk

`func (o *OrgSetting) GetApiPolicyOk() (*OrgSettingApiPolicy, bool)`

GetApiPolicyOk returns a tuple with the ApiPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiPolicy

`func (o *OrgSetting) SetApiPolicy(v OrgSettingApiPolicy)`

SetApiPolicy sets ApiPolicy field to given value.

### HasApiPolicy

`func (o *OrgSetting) HasApiPolicy() bool`

HasApiPolicy returns a boolean if a field has been set.

### GetAutoDeviceNaming

`func (o *OrgSetting) GetAutoDeviceNaming() OrgSettingAutoDeviceNaming`

GetAutoDeviceNaming returns the AutoDeviceNaming field if non-nil, zero value otherwise.

### GetAutoDeviceNamingOk

`func (o *OrgSetting) GetAutoDeviceNamingOk() (*OrgSettingAutoDeviceNaming, bool)`

GetAutoDeviceNamingOk returns a tuple with the AutoDeviceNaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDeviceNaming

`func (o *OrgSetting) SetAutoDeviceNaming(v OrgSettingAutoDeviceNaming)`

SetAutoDeviceNaming sets AutoDeviceNaming field to given value.

### HasAutoDeviceNaming

`func (o *OrgSetting) HasAutoDeviceNaming() bool`

HasAutoDeviceNaming returns a boolean if a field has been set.

### GetAutoDeviceprofileAssignment

`func (o *OrgSetting) GetAutoDeviceprofileAssignment() OrgSettingAutoDeviceprofileAssignment`

GetAutoDeviceprofileAssignment returns the AutoDeviceprofileAssignment field if non-nil, zero value otherwise.

### GetAutoDeviceprofileAssignmentOk

`func (o *OrgSetting) GetAutoDeviceprofileAssignmentOk() (*OrgSettingAutoDeviceprofileAssignment, bool)`

GetAutoDeviceprofileAssignmentOk returns a tuple with the AutoDeviceprofileAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDeviceprofileAssignment

`func (o *OrgSetting) SetAutoDeviceprofileAssignment(v OrgSettingAutoDeviceprofileAssignment)`

SetAutoDeviceprofileAssignment sets AutoDeviceprofileAssignment field to given value.

### HasAutoDeviceprofileAssignment

`func (o *OrgSetting) HasAutoDeviceprofileAssignment() bool`

HasAutoDeviceprofileAssignment returns a boolean if a field has been set.

### GetAutoSiteAssignment

`func (o *OrgSetting) GetAutoSiteAssignment() OrgSettingAutoSiteAssignment`

GetAutoSiteAssignment returns the AutoSiteAssignment field if non-nil, zero value otherwise.

### GetAutoSiteAssignmentOk

`func (o *OrgSetting) GetAutoSiteAssignmentOk() (*OrgSettingAutoSiteAssignment, bool)`

GetAutoSiteAssignmentOk returns a tuple with the AutoSiteAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSiteAssignment

`func (o *OrgSetting) SetAutoSiteAssignment(v OrgSettingAutoSiteAssignment)`

SetAutoSiteAssignment sets AutoSiteAssignment field to given value.

### HasAutoSiteAssignment

`func (o *OrgSetting) HasAutoSiteAssignment() bool`

HasAutoSiteAssignment returns a boolean if a field has been set.

### GetBlacklistUrl

`func (o *OrgSetting) GetBlacklistUrl() string`

GetBlacklistUrl returns the BlacklistUrl field if non-nil, zero value otherwise.

### GetBlacklistUrlOk

`func (o *OrgSetting) GetBlacklistUrlOk() (*string, bool)`

GetBlacklistUrlOk returns a tuple with the BlacklistUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklistUrl

`func (o *OrgSetting) SetBlacklistUrl(v string)`

SetBlacklistUrl sets BlacklistUrl field to given value.

### HasBlacklistUrl

`func (o *OrgSetting) HasBlacklistUrl() bool`

HasBlacklistUrl returns a boolean if a field has been set.

### GetCacerts

`func (o *OrgSetting) GetCacerts() []string`

GetCacerts returns the Cacerts field if non-nil, zero value otherwise.

### GetCacertsOk

`func (o *OrgSetting) GetCacertsOk() (*[]string, bool)`

GetCacertsOk returns a tuple with the Cacerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacerts

`func (o *OrgSetting) SetCacerts(v []string)`

SetCacerts sets Cacerts field to given value.

### HasCacerts

`func (o *OrgSetting) HasCacerts() bool`

HasCacerts returns a boolean if a field has been set.

### GetCelona

`func (o *OrgSetting) GetCelona() OrgSettingCelona`

GetCelona returns the Celona field if non-nil, zero value otherwise.

### GetCelonaOk

`func (o *OrgSetting) GetCelonaOk() (*OrgSettingCelona, bool)`

GetCelonaOk returns a tuple with the Celona field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCelona

`func (o *OrgSetting) SetCelona(v OrgSettingCelona)`

SetCelona sets Celona field to given value.

### HasCelona

`func (o *OrgSetting) HasCelona() bool`

HasCelona returns a boolean if a field has been set.

### GetCloudshark

`func (o *OrgSetting) GetCloudshark() OrgSettingCloudshark`

GetCloudshark returns the Cloudshark field if non-nil, zero value otherwise.

### GetCloudsharkOk

`func (o *OrgSetting) GetCloudsharkOk() (*OrgSettingCloudshark, bool)`

GetCloudsharkOk returns a tuple with the Cloudshark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudshark

`func (o *OrgSetting) SetCloudshark(v OrgSettingCloudshark)`

SetCloudshark sets Cloudshark field to given value.

### HasCloudshark

`func (o *OrgSetting) HasCloudshark() bool`

HasCloudshark returns a boolean if a field has been set.

### GetCradlepoint

`func (o *OrgSetting) GetCradlepoint() OrgSettingCradlepoint`

GetCradlepoint returns the Cradlepoint field if non-nil, zero value otherwise.

### GetCradlepointOk

`func (o *OrgSetting) GetCradlepointOk() (*OrgSettingCradlepoint, bool)`

GetCradlepointOk returns a tuple with the Cradlepoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCradlepoint

`func (o *OrgSetting) SetCradlepoint(v OrgSettingCradlepoint)`

SetCradlepoint sets Cradlepoint field to given value.

### HasCradlepoint

`func (o *OrgSetting) HasCradlepoint() bool`

HasCradlepoint returns a boolean if a field has been set.

### GetCreatedTime

`func (o *OrgSetting) GetCreatedTime() float32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *OrgSetting) GetCreatedTimeOk() (*float32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *OrgSetting) SetCreatedTime(v float32)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *OrgSetting) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDeviceCert

`func (o *OrgSetting) GetDeviceCert() OrgSettingDeviceCert`

GetDeviceCert returns the DeviceCert field if non-nil, zero value otherwise.

### GetDeviceCertOk

`func (o *OrgSetting) GetDeviceCertOk() (*OrgSettingDeviceCert, bool)`

GetDeviceCertOk returns a tuple with the DeviceCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCert

`func (o *OrgSetting) SetDeviceCert(v OrgSettingDeviceCert)`

SetDeviceCert sets DeviceCert field to given value.

### HasDeviceCert

`func (o *OrgSetting) HasDeviceCert() bool`

HasDeviceCert returns a boolean if a field has been set.

### GetDeviceUpdownThreshold

`func (o *OrgSetting) GetDeviceUpdownThreshold() int32`

GetDeviceUpdownThreshold returns the DeviceUpdownThreshold field if non-nil, zero value otherwise.

### GetDeviceUpdownThresholdOk

`func (o *OrgSetting) GetDeviceUpdownThresholdOk() (*int32, bool)`

GetDeviceUpdownThresholdOk returns a tuple with the DeviceUpdownThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceUpdownThreshold

`func (o *OrgSetting) SetDeviceUpdownThreshold(v int32)`

SetDeviceUpdownThreshold sets DeviceUpdownThreshold field to given value.

### HasDeviceUpdownThreshold

`func (o *OrgSetting) HasDeviceUpdownThreshold() bool`

HasDeviceUpdownThreshold returns a boolean if a field has been set.

### GetDisablePcap

`func (o *OrgSetting) GetDisablePcap() bool`

GetDisablePcap returns the DisablePcap field if non-nil, zero value otherwise.

### GetDisablePcapOk

`func (o *OrgSetting) GetDisablePcapOk() (*bool, bool)`

GetDisablePcapOk returns a tuple with the DisablePcap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisablePcap

`func (o *OrgSetting) SetDisablePcap(v bool)`

SetDisablePcap sets DisablePcap field to given value.

### HasDisablePcap

`func (o *OrgSetting) HasDisablePcap() bool`

HasDisablePcap returns a boolean if a field has been set.

### GetDisableRemoteShell

`func (o *OrgSetting) GetDisableRemoteShell() bool`

GetDisableRemoteShell returns the DisableRemoteShell field if non-nil, zero value otherwise.

### GetDisableRemoteShellOk

`func (o *OrgSetting) GetDisableRemoteShellOk() (*bool, bool)`

GetDisableRemoteShellOk returns a tuple with the DisableRemoteShell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableRemoteShell

`func (o *OrgSetting) SetDisableRemoteShell(v bool)`

SetDisableRemoteShell sets DisableRemoteShell field to given value.

### HasDisableRemoteShell

`func (o *OrgSetting) HasDisableRemoteShell() bool`

HasDisableRemoteShell returns a boolean if a field has been set.

### GetForSite

`func (o *OrgSetting) GetForSite() bool`

GetForSite returns the ForSite field if non-nil, zero value otherwise.

### GetForSiteOk

`func (o *OrgSetting) GetForSiteOk() (*bool, bool)`

GetForSiteOk returns a tuple with the ForSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForSite

`func (o *OrgSetting) SetForSite(v bool)`

SetForSite sets ForSite field to given value.

### HasForSite

`func (o *OrgSetting) HasForSite() bool`

HasForSite returns a boolean if a field has been set.

### GetGatewayMgmt

`func (o *OrgSetting) GetGatewayMgmt() OrgSettingGatewayMgmt`

GetGatewayMgmt returns the GatewayMgmt field if non-nil, zero value otherwise.

### GetGatewayMgmtOk

`func (o *OrgSetting) GetGatewayMgmtOk() (*OrgSettingGatewayMgmt, bool)`

GetGatewayMgmtOk returns a tuple with the GatewayMgmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayMgmt

`func (o *OrgSetting) SetGatewayMgmt(v OrgSettingGatewayMgmt)`

SetGatewayMgmt sets GatewayMgmt field to given value.

### HasGatewayMgmt

`func (o *OrgSetting) HasGatewayMgmt() bool`

HasGatewayMgmt returns a boolean if a field has been set.

### GetGatewayUpdownThreshold

`func (o *OrgSetting) GetGatewayUpdownThreshold() int32`

GetGatewayUpdownThreshold returns the GatewayUpdownThreshold field if non-nil, zero value otherwise.

### GetGatewayUpdownThresholdOk

`func (o *OrgSetting) GetGatewayUpdownThresholdOk() (*int32, bool)`

GetGatewayUpdownThresholdOk returns a tuple with the GatewayUpdownThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayUpdownThreshold

`func (o *OrgSetting) SetGatewayUpdownThreshold(v int32)`

SetGatewayUpdownThreshold sets GatewayUpdownThreshold field to given value.

### HasGatewayUpdownThreshold

`func (o *OrgSetting) HasGatewayUpdownThreshold() bool`

HasGatewayUpdownThreshold returns a boolean if a field has been set.

### SetGatewayUpdownThresholdNil

`func (o *OrgSetting) SetGatewayUpdownThresholdNil(b bool)`

 SetGatewayUpdownThresholdNil sets the value for GatewayUpdownThreshold to be an explicit nil

### UnsetGatewayUpdownThreshold
`func (o *OrgSetting) UnsetGatewayUpdownThreshold()`

UnsetGatewayUpdownThreshold ensures that no value is present for GatewayUpdownThreshold, not even an explicit nil
### GetId

`func (o *OrgSetting) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrgSetting) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrgSetting) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrgSetting) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstaller

`func (o *OrgSetting) GetInstaller() OrgSettingInstaller`

GetInstaller returns the Installer field if non-nil, zero value otherwise.

### GetInstallerOk

`func (o *OrgSetting) GetInstallerOk() (*OrgSettingInstaller, bool)`

GetInstallerOk returns a tuple with the Installer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstaller

`func (o *OrgSetting) SetInstaller(v OrgSettingInstaller)`

SetInstaller sets Installer field to given value.

### HasInstaller

`func (o *OrgSetting) HasInstaller() bool`

HasInstaller returns a boolean if a field has been set.

### GetJcloud

`func (o *OrgSetting) GetJcloud() OrgSettingJcloud`

GetJcloud returns the Jcloud field if non-nil, zero value otherwise.

### GetJcloudOk

`func (o *OrgSetting) GetJcloudOk() (*OrgSettingJcloud, bool)`

GetJcloudOk returns a tuple with the Jcloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJcloud

`func (o *OrgSetting) SetJcloud(v OrgSettingJcloud)`

SetJcloud sets Jcloud field to given value.

### HasJcloud

`func (o *OrgSetting) HasJcloud() bool`

HasJcloud returns a boolean if a field has been set.

### GetJuniper

`func (o *OrgSetting) GetJuniper() AccountJuniperInfo`

GetJuniper returns the Juniper field if non-nil, zero value otherwise.

### GetJuniperOk

`func (o *OrgSetting) GetJuniperOk() (*AccountJuniperInfo, bool)`

GetJuniperOk returns a tuple with the Juniper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJuniper

`func (o *OrgSetting) SetJuniper(v AccountJuniperInfo)`

SetJuniper sets Juniper field to given value.

### HasJuniper

`func (o *OrgSetting) HasJuniper() bool`

HasJuniper returns a boolean if a field has been set.

### GetMgmt

`func (o *OrgSetting) GetMgmt() OrgSettingMgmt`

GetMgmt returns the Mgmt field if non-nil, zero value otherwise.

### GetMgmtOk

`func (o *OrgSetting) GetMgmtOk() (*OrgSettingMgmt, bool)`

GetMgmtOk returns a tuple with the Mgmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmt

`func (o *OrgSetting) SetMgmt(v OrgSettingMgmt)`

SetMgmt sets Mgmt field to given value.

### HasMgmt

`func (o *OrgSetting) HasMgmt() bool`

HasMgmt returns a boolean if a field has been set.

### GetMistNac

`func (o *OrgSetting) GetMistNac() OrgSettingMistNac`

GetMistNac returns the MistNac field if non-nil, zero value otherwise.

### GetMistNacOk

`func (o *OrgSetting) GetMistNacOk() (*OrgSettingMistNac, bool)`

GetMistNacOk returns a tuple with the MistNac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMistNac

`func (o *OrgSetting) SetMistNac(v OrgSettingMistNac)`

SetMistNac sets MistNac field to given value.

### HasMistNac

`func (o *OrgSetting) HasMistNac() bool`

HasMistNac returns a boolean if a field has been set.

### GetModifiedTime

`func (o *OrgSetting) GetModifiedTime() float32`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *OrgSetting) GetModifiedTimeOk() (*float32, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *OrgSetting) SetModifiedTime(v float32)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *OrgSetting) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### GetMspId

`func (o *OrgSetting) GetMspId() string`

GetMspId returns the MspId field if non-nil, zero value otherwise.

### GetMspIdOk

`func (o *OrgSetting) GetMspIdOk() (*string, bool)`

GetMspIdOk returns a tuple with the MspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMspId

`func (o *OrgSetting) SetMspId(v string)`

SetMspId sets MspId field to given value.

### HasMspId

`func (o *OrgSetting) HasMspId() bool`

HasMspId returns a boolean if a field has been set.

### GetMxedgeFipsEnabled

`func (o *OrgSetting) GetMxedgeFipsEnabled() bool`

GetMxedgeFipsEnabled returns the MxedgeFipsEnabled field if non-nil, zero value otherwise.

### GetMxedgeFipsEnabledOk

`func (o *OrgSetting) GetMxedgeFipsEnabledOk() (*bool, bool)`

GetMxedgeFipsEnabledOk returns a tuple with the MxedgeFipsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMxedgeFipsEnabled

`func (o *OrgSetting) SetMxedgeFipsEnabled(v bool)`

SetMxedgeFipsEnabled sets MxedgeFipsEnabled field to given value.

### HasMxedgeFipsEnabled

`func (o *OrgSetting) HasMxedgeFipsEnabled() bool`

HasMxedgeFipsEnabled returns a boolean if a field has been set.

### GetMxedgeMgmt

`func (o *OrgSetting) GetMxedgeMgmt() MxedgeMgmt`

GetMxedgeMgmt returns the MxedgeMgmt field if non-nil, zero value otherwise.

### GetMxedgeMgmtOk

`func (o *OrgSetting) GetMxedgeMgmtOk() (*MxedgeMgmt, bool)`

GetMxedgeMgmtOk returns a tuple with the MxedgeMgmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMxedgeMgmt

`func (o *OrgSetting) SetMxedgeMgmt(v MxedgeMgmt)`

SetMxedgeMgmt sets MxedgeMgmt field to given value.

### HasMxedgeMgmt

`func (o *OrgSetting) HasMxedgeMgmt() bool`

HasMxedgeMgmt returns a boolean if a field has been set.

### GetOrgId

`func (o *OrgSetting) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *OrgSetting) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *OrgSetting) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *OrgSetting) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPasswordPolicy

`func (o *OrgSetting) GetPasswordPolicy() OrgSettingPasswordPolicy`

GetPasswordPolicy returns the PasswordPolicy field if non-nil, zero value otherwise.

### GetPasswordPolicyOk

`func (o *OrgSetting) GetPasswordPolicyOk() (*OrgSettingPasswordPolicy, bool)`

GetPasswordPolicyOk returns a tuple with the PasswordPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordPolicy

`func (o *OrgSetting) SetPasswordPolicy(v OrgSettingPasswordPolicy)`

SetPasswordPolicy sets PasswordPolicy field to given value.

### HasPasswordPolicy

`func (o *OrgSetting) HasPasswordPolicy() bool`

HasPasswordPolicy returns a boolean if a field has been set.

### GetPcap

`func (o *OrgSetting) GetPcap() OrgSettingPcap`

GetPcap returns the Pcap field if non-nil, zero value otherwise.

### GetPcapOk

`func (o *OrgSetting) GetPcapOk() (*OrgSettingPcap, bool)`

GetPcapOk returns a tuple with the Pcap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcap

`func (o *OrgSetting) SetPcap(v OrgSettingPcap)`

SetPcap sets Pcap field to given value.

### HasPcap

`func (o *OrgSetting) HasPcap() bool`

HasPcap returns a boolean if a field has been set.

### GetPcapBucketVerified

`func (o *OrgSetting) GetPcapBucketVerified() bool`

GetPcapBucketVerified returns the PcapBucketVerified field if non-nil, zero value otherwise.

### GetPcapBucketVerifiedOk

`func (o *OrgSetting) GetPcapBucketVerifiedOk() (*bool, bool)`

GetPcapBucketVerifiedOk returns a tuple with the PcapBucketVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcapBucketVerified

`func (o *OrgSetting) SetPcapBucketVerified(v bool)`

SetPcapBucketVerified sets PcapBucketVerified field to given value.

### HasPcapBucketVerified

`func (o *OrgSetting) HasPcapBucketVerified() bool`

HasPcapBucketVerified returns a boolean if a field has been set.

### GetSecurity

`func (o *OrgSetting) GetSecurity() OrgSettingSecurity`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *OrgSetting) GetSecurityOk() (*OrgSettingSecurity, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *OrgSetting) SetSecurity(v OrgSettingSecurity)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *OrgSetting) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetSimpleAlert

`func (o *OrgSetting) GetSimpleAlert() SimpleAlert`

GetSimpleAlert returns the SimpleAlert field if non-nil, zero value otherwise.

### GetSimpleAlertOk

`func (o *OrgSetting) GetSimpleAlertOk() (*SimpleAlert, bool)`

GetSimpleAlertOk returns a tuple with the SimpleAlert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimpleAlert

`func (o *OrgSetting) SetSimpleAlert(v SimpleAlert)`

SetSimpleAlert sets SimpleAlert field to given value.

### HasSimpleAlert

`func (o *OrgSetting) HasSimpleAlert() bool`

HasSimpleAlert returns a boolean if a field has been set.

### GetSiteId

`func (o *OrgSetting) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *OrgSetting) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *OrgSetting) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *OrgSetting) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSwitchMgmt

`func (o *OrgSetting) GetSwitchMgmt() OrgSettingSwitchMgmt`

GetSwitchMgmt returns the SwitchMgmt field if non-nil, zero value otherwise.

### GetSwitchMgmtOk

`func (o *OrgSetting) GetSwitchMgmtOk() (*OrgSettingSwitchMgmt, bool)`

GetSwitchMgmtOk returns a tuple with the SwitchMgmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchMgmt

`func (o *OrgSetting) SetSwitchMgmt(v OrgSettingSwitchMgmt)`

SetSwitchMgmt sets SwitchMgmt field to given value.

### HasSwitchMgmt

`func (o *OrgSetting) HasSwitchMgmt() bool`

HasSwitchMgmt returns a boolean if a field has been set.

### GetSwitchUpdownThreshold

`func (o *OrgSetting) GetSwitchUpdownThreshold() int32`

GetSwitchUpdownThreshold returns the SwitchUpdownThreshold field if non-nil, zero value otherwise.

### GetSwitchUpdownThresholdOk

`func (o *OrgSetting) GetSwitchUpdownThresholdOk() (*int32, bool)`

GetSwitchUpdownThresholdOk returns a tuple with the SwitchUpdownThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchUpdownThreshold

`func (o *OrgSetting) SetSwitchUpdownThreshold(v int32)`

SetSwitchUpdownThreshold sets SwitchUpdownThreshold field to given value.

### HasSwitchUpdownThreshold

`func (o *OrgSetting) HasSwitchUpdownThreshold() bool`

HasSwitchUpdownThreshold returns a boolean if a field has been set.

### SetSwitchUpdownThresholdNil

`func (o *OrgSetting) SetSwitchUpdownThresholdNil(b bool)`

 SetSwitchUpdownThresholdNil sets the value for SwitchUpdownThreshold to be an explicit nil

### UnsetSwitchUpdownThreshold
`func (o *OrgSetting) UnsetSwitchUpdownThreshold()`

UnsetSwitchUpdownThreshold ensures that no value is present for SwitchUpdownThreshold, not even an explicit nil
### GetSyntheticTest

`func (o *OrgSetting) GetSyntheticTest() SyntheticTestConfig`

GetSyntheticTest returns the SyntheticTest field if non-nil, zero value otherwise.

### GetSyntheticTestOk

`func (o *OrgSetting) GetSyntheticTestOk() (*SyntheticTestConfig, bool)`

GetSyntheticTestOk returns a tuple with the SyntheticTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyntheticTest

`func (o *OrgSetting) SetSyntheticTest(v SyntheticTestConfig)`

SetSyntheticTest sets SyntheticTest field to given value.

### HasSyntheticTest

`func (o *OrgSetting) HasSyntheticTest() bool`

HasSyntheticTest returns a boolean if a field has been set.

### GetTags

`func (o *OrgSetting) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *OrgSetting) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *OrgSetting) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *OrgSetting) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUiIdleTimeout

`func (o *OrgSetting) GetUiIdleTimeout() int32`

GetUiIdleTimeout returns the UiIdleTimeout field if non-nil, zero value otherwise.

### GetUiIdleTimeoutOk

`func (o *OrgSetting) GetUiIdleTimeoutOk() (*int32, bool)`

GetUiIdleTimeoutOk returns a tuple with the UiIdleTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiIdleTimeout

`func (o *OrgSetting) SetUiIdleTimeout(v int32)`

SetUiIdleTimeout sets UiIdleTimeout field to given value.

### HasUiIdleTimeout

`func (o *OrgSetting) HasUiIdleTimeout() bool`

HasUiIdleTimeout returns a boolean if a field has been set.

### GetVpnOptions

`func (o *OrgSetting) GetVpnOptions() OrgSettingVpnOptions`

GetVpnOptions returns the VpnOptions field if non-nil, zero value otherwise.

### GetVpnOptionsOk

`func (o *OrgSetting) GetVpnOptionsOk() (*OrgSettingVpnOptions, bool)`

GetVpnOptionsOk returns a tuple with the VpnOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnOptions

`func (o *OrgSetting) SetVpnOptions(v OrgSettingVpnOptions)`

SetVpnOptions sets VpnOptions field to given value.

### HasVpnOptions

`func (o *OrgSetting) HasVpnOptions() bool`

HasVpnOptions returns a boolean if a field has been set.

### GetWanPma

`func (o *OrgSetting) GetWanPma() OrgSettingWanPma`

GetWanPma returns the WanPma field if non-nil, zero value otherwise.

### GetWanPmaOk

`func (o *OrgSetting) GetWanPmaOk() (*OrgSettingWanPma, bool)`

GetWanPmaOk returns a tuple with the WanPma field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanPma

`func (o *OrgSetting) SetWanPma(v OrgSettingWanPma)`

SetWanPma sets WanPma field to given value.

### HasWanPma

`func (o *OrgSetting) HasWanPma() bool`

HasWanPma returns a boolean if a field has been set.

### GetWiredPma

`func (o *OrgSetting) GetWiredPma() OrgSettingWiredPma`

GetWiredPma returns the WiredPma field if non-nil, zero value otherwise.

### GetWiredPmaOk

`func (o *OrgSetting) GetWiredPmaOk() (*OrgSettingWiredPma, bool)`

GetWiredPmaOk returns a tuple with the WiredPma field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWiredPma

`func (o *OrgSetting) SetWiredPma(v OrgSettingWiredPma)`

SetWiredPma sets WiredPma field to given value.

### HasWiredPma

`func (o *OrgSetting) HasWiredPma() bool`

HasWiredPma returns a boolean if a field has been set.

### GetWirelessPma

`func (o *OrgSetting) GetWirelessPma() OrgSettingWirelessPma`

GetWirelessPma returns the WirelessPma field if non-nil, zero value otherwise.

### GetWirelessPmaOk

`func (o *OrgSetting) GetWirelessPmaOk() (*OrgSettingWirelessPma, bool)`

GetWirelessPmaOk returns a tuple with the WirelessPma field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWirelessPma

`func (o *OrgSetting) SetWirelessPma(v OrgSettingWirelessPma)`

SetWirelessPma sets WirelessPma field to given value.

### HasWirelessPma

`func (o *OrgSetting) HasWirelessPma() bool`

HasWirelessPma returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


