# ClientWireless

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ap** | Pointer to **[]string** | List of AP MAC Addresses the client was connected to | [optional] 
**AppVersion** | Pointer to **[]string** | only when client has the Marvis Client app running. List of the versions of the Marvis Client | [optional] 
**Band** | Pointer to **string** | Wi-Fi Radio band | [optional] 
**Device** | Pointer to **[]string** | only when client has the Marvis Client app running. List of the type of device type detected | [optional] 
**Ftc** | Pointer to **bool** |  | [optional] 
**Hardware** | Pointer to **string** | only when client has the Marvis Client app running. Type of Wi-Fi adapter | [optional] 
**Hostname** | Pointer to **[]string** | list of hostname detected for this client | [optional] 
**Ip** | Pointer to **[]string** | list if the ip addresses detected for this client | [optional] 
**LastAp** | Pointer to **string** | latest AP where the client is/was connected to | [optional] 
**LastDevuce** | Pointer to **string** | latest type of device we identified (e.g. iPhone, Mac, …) | [optional] 
**LastFirmware** | Pointer to **string** | only when client has the Marvis Client app running. Same as \&quot;firmware\&quot; | [optional] 
**LastHostname** | Pointer to **string** | lastest hostname we detected for the the client | [optional] 
**LastIp** | Pointer to **string** | lastest ip address we detected for the client | [optional] 
**LastModel** | Pointer to **string** | only when client has the Marvis Client app running. lastest client hardware model we detected for the client | [optional] 
**LastOs** | Pointer to **string** | only when client has the Marvis Client app running. Latest version of OS Type we detected for the client | [optional] 
**LastOsVersion** | Pointer to **string** | only when client has the Marvis Client app running. Latest version of OS Version we detected for the client | [optional] 
**LastPskId** | Pointer to **string** | only for PPSK authentication. Latest PPSK ID used by the client | [optional] 
**LastPskName** | Pointer to **string** | only for PPSK authentication. Latest PPSK Name used by the client | [optional] 
**LastSsid** | Pointer to **string** | only for PPSK authentication. Latest PPSK Name used by the client | [optional] 
**LastUsername** | Pointer to **string** | only for 802.1X authentifcation. Latest username used by the client | [optional] 
**LastVlan** | Pointer to **int32** | latest VLAN ID assigned to the client | [optional] 
**LastWlanId** | Pointer to **string** | ID of the latest SSID (WLAN) the client is/was connected to | [optional] 
**Mac** | Pointer to **string** | Client MAC Address | [optional] 
**Mfg** | Pointer to **string** | manufacturer of the client hardware (MAC OUI based) | [optional] 
**Model** | Pointer to **string** | only when client has the Marvis Client app running. Client hardware model | [optional] 
**OrgId** | Pointer to **string** | Mist Org ID | [optional] 
**Os** | Pointer to **[]string** | only when client is having the Marvis Client app running. List of OS detected for the client | [optional] 
**OsVersion** | Pointer to **[]string** | only when client is having the Marvis Client app running. List of OS version detected for the client | [optional] 
**Protocol** | Pointer to **string** | 802.11 amendmant | [optional] 
**PskId** | Pointer to **[]string** | list of IDs of the PPSK used by the client | [optional] 
**PskName** | Pointer to **[]string** | list of names of the PPSK used by the client | [optional] 
**RandomMac** | Pointer to **bool** | whether the client is using randomized MAC Address or not | [optional] 
**SdkVersion** | Pointer to **[]string** | only when client has the Marvis Client app running. List of Marvis Client SDK version detected for the client | [optional] 
**SiteId** | Pointer to **string** | Mist Site ID where the client is connected | [optional] 
**SiteIds** | Pointer to **[]string** | list of Mist Site IDs where the client was connected | [optional] 
**Ssid** | Pointer to **[]string** | list of the WLAN names the client was connected to | [optional] 
**Timestamp** | Pointer to **float32** | when the data has been updated | [optional] 
**Username** | Pointer to **[]string** | only for 802.1X authentifcation. List of usernames used by the client | [optional] 
**Vlan** | Pointer to **[]int32** | list of vlans that have been assigned to the client | [optional] 
**WlanId** | Pointer to **[]string** | list of IDs of WLANs the client was connected to | [optional] 

## Methods

### NewClientWireless

`func NewClientWireless() *ClientWireless`

NewClientWireless instantiates a new ClientWireless object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientWirelessWithDefaults

`func NewClientWirelessWithDefaults() *ClientWireless`

NewClientWirelessWithDefaults instantiates a new ClientWireless object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAp

`func (o *ClientWireless) GetAp() []string`

GetAp returns the Ap field if non-nil, zero value otherwise.

### GetApOk

`func (o *ClientWireless) GetApOk() (*[]string, bool)`

GetApOk returns a tuple with the Ap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAp

`func (o *ClientWireless) SetAp(v []string)`

SetAp sets Ap field to given value.

### HasAp

`func (o *ClientWireless) HasAp() bool`

HasAp returns a boolean if a field has been set.

### GetAppVersion

`func (o *ClientWireless) GetAppVersion() []string`

GetAppVersion returns the AppVersion field if non-nil, zero value otherwise.

### GetAppVersionOk

`func (o *ClientWireless) GetAppVersionOk() (*[]string, bool)`

GetAppVersionOk returns a tuple with the AppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppVersion

`func (o *ClientWireless) SetAppVersion(v []string)`

SetAppVersion sets AppVersion field to given value.

### HasAppVersion

`func (o *ClientWireless) HasAppVersion() bool`

HasAppVersion returns a boolean if a field has been set.

### GetBand

`func (o *ClientWireless) GetBand() string`

GetBand returns the Band field if non-nil, zero value otherwise.

### GetBandOk

`func (o *ClientWireless) GetBandOk() (*string, bool)`

GetBandOk returns a tuple with the Band field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand

`func (o *ClientWireless) SetBand(v string)`

SetBand sets Band field to given value.

### HasBand

`func (o *ClientWireless) HasBand() bool`

HasBand returns a boolean if a field has been set.

### GetDevice

`func (o *ClientWireless) GetDevice() []string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *ClientWireless) GetDeviceOk() (*[]string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *ClientWireless) SetDevice(v []string)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *ClientWireless) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetFtc

`func (o *ClientWireless) GetFtc() bool`

GetFtc returns the Ftc field if non-nil, zero value otherwise.

### GetFtcOk

`func (o *ClientWireless) GetFtcOk() (*bool, bool)`

GetFtcOk returns a tuple with the Ftc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFtc

`func (o *ClientWireless) SetFtc(v bool)`

SetFtc sets Ftc field to given value.

### HasFtc

`func (o *ClientWireless) HasFtc() bool`

HasFtc returns a boolean if a field has been set.

### GetHardware

`func (o *ClientWireless) GetHardware() string`

GetHardware returns the Hardware field if non-nil, zero value otherwise.

### GetHardwareOk

`func (o *ClientWireless) GetHardwareOk() (*string, bool)`

GetHardwareOk returns a tuple with the Hardware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardware

`func (o *ClientWireless) SetHardware(v string)`

SetHardware sets Hardware field to given value.

### HasHardware

`func (o *ClientWireless) HasHardware() bool`

HasHardware returns a boolean if a field has been set.

### GetHostname

`func (o *ClientWireless) GetHostname() []string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ClientWireless) GetHostnameOk() (*[]string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ClientWireless) SetHostname(v []string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ClientWireless) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetIp

`func (o *ClientWireless) GetIp() []string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ClientWireless) GetIpOk() (*[]string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ClientWireless) SetIp(v []string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *ClientWireless) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetLastAp

`func (o *ClientWireless) GetLastAp() string`

GetLastAp returns the LastAp field if non-nil, zero value otherwise.

### GetLastApOk

`func (o *ClientWireless) GetLastApOk() (*string, bool)`

GetLastApOk returns a tuple with the LastAp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAp

`func (o *ClientWireless) SetLastAp(v string)`

SetLastAp sets LastAp field to given value.

### HasLastAp

`func (o *ClientWireless) HasLastAp() bool`

HasLastAp returns a boolean if a field has been set.

### GetLastDevuce

`func (o *ClientWireless) GetLastDevuce() string`

GetLastDevuce returns the LastDevuce field if non-nil, zero value otherwise.

### GetLastDevuceOk

`func (o *ClientWireless) GetLastDevuceOk() (*string, bool)`

GetLastDevuceOk returns a tuple with the LastDevuce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDevuce

`func (o *ClientWireless) SetLastDevuce(v string)`

SetLastDevuce sets LastDevuce field to given value.

### HasLastDevuce

`func (o *ClientWireless) HasLastDevuce() bool`

HasLastDevuce returns a boolean if a field has been set.

### GetLastFirmware

`func (o *ClientWireless) GetLastFirmware() string`

GetLastFirmware returns the LastFirmware field if non-nil, zero value otherwise.

### GetLastFirmwareOk

`func (o *ClientWireless) GetLastFirmwareOk() (*string, bool)`

GetLastFirmwareOk returns a tuple with the LastFirmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFirmware

`func (o *ClientWireless) SetLastFirmware(v string)`

SetLastFirmware sets LastFirmware field to given value.

### HasLastFirmware

`func (o *ClientWireless) HasLastFirmware() bool`

HasLastFirmware returns a boolean if a field has been set.

### GetLastHostname

`func (o *ClientWireless) GetLastHostname() string`

GetLastHostname returns the LastHostname field if non-nil, zero value otherwise.

### GetLastHostnameOk

`func (o *ClientWireless) GetLastHostnameOk() (*string, bool)`

GetLastHostnameOk returns a tuple with the LastHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastHostname

`func (o *ClientWireless) SetLastHostname(v string)`

SetLastHostname sets LastHostname field to given value.

### HasLastHostname

`func (o *ClientWireless) HasLastHostname() bool`

HasLastHostname returns a boolean if a field has been set.

### GetLastIp

`func (o *ClientWireless) GetLastIp() string`

GetLastIp returns the LastIp field if non-nil, zero value otherwise.

### GetLastIpOk

`func (o *ClientWireless) GetLastIpOk() (*string, bool)`

GetLastIpOk returns a tuple with the LastIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastIp

`func (o *ClientWireless) SetLastIp(v string)`

SetLastIp sets LastIp field to given value.

### HasLastIp

`func (o *ClientWireless) HasLastIp() bool`

HasLastIp returns a boolean if a field has been set.

### GetLastModel

`func (o *ClientWireless) GetLastModel() string`

GetLastModel returns the LastModel field if non-nil, zero value otherwise.

### GetLastModelOk

`func (o *ClientWireless) GetLastModelOk() (*string, bool)`

GetLastModelOk returns a tuple with the LastModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModel

`func (o *ClientWireless) SetLastModel(v string)`

SetLastModel sets LastModel field to given value.

### HasLastModel

`func (o *ClientWireless) HasLastModel() bool`

HasLastModel returns a boolean if a field has been set.

### GetLastOs

`func (o *ClientWireless) GetLastOs() string`

GetLastOs returns the LastOs field if non-nil, zero value otherwise.

### GetLastOsOk

`func (o *ClientWireless) GetLastOsOk() (*string, bool)`

GetLastOsOk returns a tuple with the LastOs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastOs

`func (o *ClientWireless) SetLastOs(v string)`

SetLastOs sets LastOs field to given value.

### HasLastOs

`func (o *ClientWireless) HasLastOs() bool`

HasLastOs returns a boolean if a field has been set.

### GetLastOsVersion

`func (o *ClientWireless) GetLastOsVersion() string`

GetLastOsVersion returns the LastOsVersion field if non-nil, zero value otherwise.

### GetLastOsVersionOk

`func (o *ClientWireless) GetLastOsVersionOk() (*string, bool)`

GetLastOsVersionOk returns a tuple with the LastOsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastOsVersion

`func (o *ClientWireless) SetLastOsVersion(v string)`

SetLastOsVersion sets LastOsVersion field to given value.

### HasLastOsVersion

`func (o *ClientWireless) HasLastOsVersion() bool`

HasLastOsVersion returns a boolean if a field has been set.

### GetLastPskId

`func (o *ClientWireless) GetLastPskId() string`

GetLastPskId returns the LastPskId field if non-nil, zero value otherwise.

### GetLastPskIdOk

`func (o *ClientWireless) GetLastPskIdOk() (*string, bool)`

GetLastPskIdOk returns a tuple with the LastPskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPskId

`func (o *ClientWireless) SetLastPskId(v string)`

SetLastPskId sets LastPskId field to given value.

### HasLastPskId

`func (o *ClientWireless) HasLastPskId() bool`

HasLastPskId returns a boolean if a field has been set.

### GetLastPskName

`func (o *ClientWireless) GetLastPskName() string`

GetLastPskName returns the LastPskName field if non-nil, zero value otherwise.

### GetLastPskNameOk

`func (o *ClientWireless) GetLastPskNameOk() (*string, bool)`

GetLastPskNameOk returns a tuple with the LastPskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPskName

`func (o *ClientWireless) SetLastPskName(v string)`

SetLastPskName sets LastPskName field to given value.

### HasLastPskName

`func (o *ClientWireless) HasLastPskName() bool`

HasLastPskName returns a boolean if a field has been set.

### GetLastSsid

`func (o *ClientWireless) GetLastSsid() string`

GetLastSsid returns the LastSsid field if non-nil, zero value otherwise.

### GetLastSsidOk

`func (o *ClientWireless) GetLastSsidOk() (*string, bool)`

GetLastSsidOk returns a tuple with the LastSsid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSsid

`func (o *ClientWireless) SetLastSsid(v string)`

SetLastSsid sets LastSsid field to given value.

### HasLastSsid

`func (o *ClientWireless) HasLastSsid() bool`

HasLastSsid returns a boolean if a field has been set.

### GetLastUsername

`func (o *ClientWireless) GetLastUsername() string`

GetLastUsername returns the LastUsername field if non-nil, zero value otherwise.

### GetLastUsernameOk

`func (o *ClientWireless) GetLastUsernameOk() (*string, bool)`

GetLastUsernameOk returns a tuple with the LastUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsername

`func (o *ClientWireless) SetLastUsername(v string)`

SetLastUsername sets LastUsername field to given value.

### HasLastUsername

`func (o *ClientWireless) HasLastUsername() bool`

HasLastUsername returns a boolean if a field has been set.

### GetLastVlan

`func (o *ClientWireless) GetLastVlan() int32`

GetLastVlan returns the LastVlan field if non-nil, zero value otherwise.

### GetLastVlanOk

`func (o *ClientWireless) GetLastVlanOk() (*int32, bool)`

GetLastVlanOk returns a tuple with the LastVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastVlan

`func (o *ClientWireless) SetLastVlan(v int32)`

SetLastVlan sets LastVlan field to given value.

### HasLastVlan

`func (o *ClientWireless) HasLastVlan() bool`

HasLastVlan returns a boolean if a field has been set.

### GetLastWlanId

`func (o *ClientWireless) GetLastWlanId() string`

GetLastWlanId returns the LastWlanId field if non-nil, zero value otherwise.

### GetLastWlanIdOk

`func (o *ClientWireless) GetLastWlanIdOk() (*string, bool)`

GetLastWlanIdOk returns a tuple with the LastWlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastWlanId

`func (o *ClientWireless) SetLastWlanId(v string)`

SetLastWlanId sets LastWlanId field to given value.

### HasLastWlanId

`func (o *ClientWireless) HasLastWlanId() bool`

HasLastWlanId returns a boolean if a field has been set.

### GetMac

`func (o *ClientWireless) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *ClientWireless) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *ClientWireless) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *ClientWireless) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMfg

`func (o *ClientWireless) GetMfg() string`

GetMfg returns the Mfg field if non-nil, zero value otherwise.

### GetMfgOk

`func (o *ClientWireless) GetMfgOk() (*string, bool)`

GetMfgOk returns a tuple with the Mfg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfg

`func (o *ClientWireless) SetMfg(v string)`

SetMfg sets Mfg field to given value.

### HasMfg

`func (o *ClientWireless) HasMfg() bool`

HasMfg returns a boolean if a field has been set.

### GetModel

`func (o *ClientWireless) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ClientWireless) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ClientWireless) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ClientWireless) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetOrgId

`func (o *ClientWireless) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ClientWireless) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ClientWireless) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ClientWireless) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetOs

`func (o *ClientWireless) GetOs() []string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *ClientWireless) GetOsOk() (*[]string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *ClientWireless) SetOs(v []string)`

SetOs sets Os field to given value.

### HasOs

`func (o *ClientWireless) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetOsVersion

`func (o *ClientWireless) GetOsVersion() []string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *ClientWireless) GetOsVersionOk() (*[]string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *ClientWireless) SetOsVersion(v []string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *ClientWireless) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetProtocol

`func (o *ClientWireless) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ClientWireless) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ClientWireless) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *ClientWireless) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetPskId

`func (o *ClientWireless) GetPskId() []string`

GetPskId returns the PskId field if non-nil, zero value otherwise.

### GetPskIdOk

`func (o *ClientWireless) GetPskIdOk() (*[]string, bool)`

GetPskIdOk returns a tuple with the PskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPskId

`func (o *ClientWireless) SetPskId(v []string)`

SetPskId sets PskId field to given value.

### HasPskId

`func (o *ClientWireless) HasPskId() bool`

HasPskId returns a boolean if a field has been set.

### GetPskName

`func (o *ClientWireless) GetPskName() []string`

GetPskName returns the PskName field if non-nil, zero value otherwise.

### GetPskNameOk

`func (o *ClientWireless) GetPskNameOk() (*[]string, bool)`

GetPskNameOk returns a tuple with the PskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPskName

`func (o *ClientWireless) SetPskName(v []string)`

SetPskName sets PskName field to given value.

### HasPskName

`func (o *ClientWireless) HasPskName() bool`

HasPskName returns a boolean if a field has been set.

### GetRandomMac

`func (o *ClientWireless) GetRandomMac() bool`

GetRandomMac returns the RandomMac field if non-nil, zero value otherwise.

### GetRandomMacOk

`func (o *ClientWireless) GetRandomMacOk() (*bool, bool)`

GetRandomMacOk returns a tuple with the RandomMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomMac

`func (o *ClientWireless) SetRandomMac(v bool)`

SetRandomMac sets RandomMac field to given value.

### HasRandomMac

`func (o *ClientWireless) HasRandomMac() bool`

HasRandomMac returns a boolean if a field has been set.

### GetSdkVersion

`func (o *ClientWireless) GetSdkVersion() []string`

GetSdkVersion returns the SdkVersion field if non-nil, zero value otherwise.

### GetSdkVersionOk

`func (o *ClientWireless) GetSdkVersionOk() (*[]string, bool)`

GetSdkVersionOk returns a tuple with the SdkVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkVersion

`func (o *ClientWireless) SetSdkVersion(v []string)`

SetSdkVersion sets SdkVersion field to given value.

### HasSdkVersion

`func (o *ClientWireless) HasSdkVersion() bool`

HasSdkVersion returns a boolean if a field has been set.

### GetSiteId

`func (o *ClientWireless) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *ClientWireless) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *ClientWireless) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *ClientWireless) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSiteIds

`func (o *ClientWireless) GetSiteIds() []string`

GetSiteIds returns the SiteIds field if non-nil, zero value otherwise.

### GetSiteIdsOk

`func (o *ClientWireless) GetSiteIdsOk() (*[]string, bool)`

GetSiteIdsOk returns a tuple with the SiteIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteIds

`func (o *ClientWireless) SetSiteIds(v []string)`

SetSiteIds sets SiteIds field to given value.

### HasSiteIds

`func (o *ClientWireless) HasSiteIds() bool`

HasSiteIds returns a boolean if a field has been set.

### GetSsid

`func (o *ClientWireless) GetSsid() []string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *ClientWireless) GetSsidOk() (*[]string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *ClientWireless) SetSsid(v []string)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *ClientWireless) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### GetTimestamp

`func (o *ClientWireless) GetTimestamp() float32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ClientWireless) GetTimestampOk() (*float32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ClientWireless) SetTimestamp(v float32)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ClientWireless) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetUsername

`func (o *ClientWireless) GetUsername() []string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ClientWireless) GetUsernameOk() (*[]string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ClientWireless) SetUsername(v []string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ClientWireless) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetVlan

`func (o *ClientWireless) GetVlan() []int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *ClientWireless) GetVlanOk() (*[]int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *ClientWireless) SetVlan(v []int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *ClientWireless) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetWlanId

`func (o *ClientWireless) GetWlanId() []string`

GetWlanId returns the WlanId field if non-nil, zero value otherwise.

### GetWlanIdOk

`func (o *ClientWireless) GetWlanIdOk() (*[]string, bool)`

GetWlanIdOk returns a tuple with the WlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWlanId

`func (o *ClientWireless) SetWlanId(v []string)`

SetWlanId sets WlanId field to given value.

### HasWlanId

`func (o *ClientWireless) HasWlanId() bool`

HasWlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


