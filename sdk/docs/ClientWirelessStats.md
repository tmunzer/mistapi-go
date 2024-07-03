# ClientWirelessStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ttl** | **float32** | TTL of the validity of the stat | 
**Accuracy** | Pointer to **int32** | estimated client location accuracy, in meter | [optional] 
**AirespaceIfname** | Pointer to **string** |  | [optional] 
**Airwatch** | Pointer to [**ClientWirelessStatsAirwatch**](ClientWirelessStatsAirwatch.md) |  | [optional] 
**ApId** | **string** | AP ID the client is connected to | 
**ApMac** | **string** | AP the client is connected to | 
**Band** | [**Dot11Band**](Dot11Band.md) |  | 
**Channel** | **int32** | current channel | 
**DualBand** | **bool** | whether the client is dual_band capable (determined by whether we’ve seen probe requests from both bands) | 
**Family** | **string** | device family, through fingerprinting. iPod / Nexus Galaxy / Windows Mobile or CE … | 
**Guest** | Pointer to [**ClientWirelessStatsGuest**](ClientWirelessStatsGuest.md) |  | [optional] 
**Hostname** | **string** | hostname that we learned from sniffing DHCP | 
**IdleTime** | **float32** | how long, in seconds, has the client been idle (since the last RX packet) | 
**Ip** | **string** |  | 
**IsGuest** | **bool** | whether this is a guest | [default to false]
**KeyMgmt** | **string** | e.g. WPA2-PSK/CCMP | 
**LastSeen** | **float32** | last seen timestamp | 
**Mac** | **string** | client mac | 
**Manufacture** | **string** | device manufacture, through fingerprinting or OUI | 
**MapId** | Pointer to **string** | estimated client location - map_id | [optional] 
**Model** | **string** | device model, may be available if we can identify them | 
**NumLocatingAps** | Pointer to **int32** | number of APs used to locate this client | [optional] 
**Os** | **string** | device os, through fingerprinting | 
**PowerSaving** | **bool** | if it’s currently in power-save mode | 
**Proto** | [**Dot11Proto**](Dot11Proto.md) |  | 
**PskId** | Pointer to **string** | PSK id (if multi-psk is used) | [optional] 
**Rssi** | **float32** | signal strength | 
**Rssizones** | Pointer to [**[]ClientWirelessStatsRssiZone**](ClientWirelessStatsRssiZone.md) | list of rssizone_id’s where client is in and since when (if known) | [optional] 
**RxBps** | **float32** | rate of receiving traffic from the clients, bits/seconds, last known | 
**RxBytes** | **float32** | amount of traffic received from client since client connects | 
**RxPackets** | **float32** | amount of traffic received from client since client connects | 
**RxRate** | **float32** | RX Rate, Mbps | 
**RxRetries** | **float32** | amount of rx retries | 
**Snr** | **float32** | signal over noise | 
**Ssid** | **string** | SSID the client is connected to | 
**TxBps** | **float32** | rate of transmitting traffic to the clients, bits/seconds, last known | 
**TxBytes** | **float32** | amount of traffic sent to client since client connects | 
**TxPackets** | **float32** | amount of traffic sent to client since client connects | 
**TxRate** | **float32** | TX Rate, Mbps | 
**TxRetries** | **float32** | amount of tx retries | 
**Type** | Pointer to **string** | client’s type, regular / vip / resource / blocked (if client object is created) | [optional] 
**Uptime** | **float32** | how long, in seconds, has the client been connected | 
**Username** | **string** | username that we learned from 802.1X exchange or Per_user PSK or User Portal | 
**Vbeacons** | Pointer to [**[]ClientWirelessStatsVbeacon**](ClientWirelessStatsVbeacon.md) | list of beacon_id’s where the client is in and since when (if known) | [optional] 
**VlanId** | Pointer to **int32** | vlan id, could be empty (from older AP) | [optional] 
**WlanId** | **string** | WLAN ID the client is connected to | 
**WxruleId** | Pointer to **string** | current WxlanRule using for a Client or an authorized Guest (portal user). null if default rule is matched. | [optional] 
**WxruleUsage** | Pointer to [**[]ClientWirelessStatsWxruleUsage**](ClientWirelessStatsWxruleUsage.md) | current WxlanRule usage per tag_id | [optional] 
**X** | Pointer to **float32** | estimated clinet location in pixels | [optional] 
**XM** | Pointer to **float32** | estimated client location in meter | [optional] 
**Y** | Pointer to **float32** | estimated clinet location in pixels | [optional] 
**YM** | Pointer to **float32** | estimated client location in meter | [optional] 
**Zones** | Pointer to [**[]ClientWirelessStatsZone**](ClientWirelessStatsZone.md) | list of zone_id’s where client is in and since when (if known) | [optional] 

## Methods

### NewClientWirelessStats

`func NewClientWirelessStats(ttl float32, apId string, apMac string, band Dot11Band, channel int32, dualBand bool, family string, hostname string, idleTime float32, ip string, isGuest bool, keyMgmt string, lastSeen float32, mac string, manufacture string, model string, os string, powerSaving bool, proto Dot11Proto, rssi float32, rxBps float32, rxBytes float32, rxPackets float32, rxRate float32, rxRetries float32, snr float32, ssid string, txBps float32, txBytes float32, txPackets float32, txRate float32, txRetries float32, uptime float32, username string, wlanId string, ) *ClientWirelessStats`

NewClientWirelessStats instantiates a new ClientWirelessStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientWirelessStatsWithDefaults

`func NewClientWirelessStatsWithDefaults() *ClientWirelessStats`

NewClientWirelessStatsWithDefaults instantiates a new ClientWirelessStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTtl

`func (o *ClientWirelessStats) GetTtl() float32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *ClientWirelessStats) GetTtlOk() (*float32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *ClientWirelessStats) SetTtl(v float32)`

SetTtl sets Ttl field to given value.


### GetAccuracy

`func (o *ClientWirelessStats) GetAccuracy() int32`

GetAccuracy returns the Accuracy field if non-nil, zero value otherwise.

### GetAccuracyOk

`func (o *ClientWirelessStats) GetAccuracyOk() (*int32, bool)`

GetAccuracyOk returns a tuple with the Accuracy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccuracy

`func (o *ClientWirelessStats) SetAccuracy(v int32)`

SetAccuracy sets Accuracy field to given value.

### HasAccuracy

`func (o *ClientWirelessStats) HasAccuracy() bool`

HasAccuracy returns a boolean if a field has been set.

### GetAirespaceIfname

`func (o *ClientWirelessStats) GetAirespaceIfname() string`

GetAirespaceIfname returns the AirespaceIfname field if non-nil, zero value otherwise.

### GetAirespaceIfnameOk

`func (o *ClientWirelessStats) GetAirespaceIfnameOk() (*string, bool)`

GetAirespaceIfnameOk returns a tuple with the AirespaceIfname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirespaceIfname

`func (o *ClientWirelessStats) SetAirespaceIfname(v string)`

SetAirespaceIfname sets AirespaceIfname field to given value.

### HasAirespaceIfname

`func (o *ClientWirelessStats) HasAirespaceIfname() bool`

HasAirespaceIfname returns a boolean if a field has been set.

### GetAirwatch

`func (o *ClientWirelessStats) GetAirwatch() ClientWirelessStatsAirwatch`

GetAirwatch returns the Airwatch field if non-nil, zero value otherwise.

### GetAirwatchOk

`func (o *ClientWirelessStats) GetAirwatchOk() (*ClientWirelessStatsAirwatch, bool)`

GetAirwatchOk returns a tuple with the Airwatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirwatch

`func (o *ClientWirelessStats) SetAirwatch(v ClientWirelessStatsAirwatch)`

SetAirwatch sets Airwatch field to given value.

### HasAirwatch

`func (o *ClientWirelessStats) HasAirwatch() bool`

HasAirwatch returns a boolean if a field has been set.

### GetApId

`func (o *ClientWirelessStats) GetApId() string`

GetApId returns the ApId field if non-nil, zero value otherwise.

### GetApIdOk

`func (o *ClientWirelessStats) GetApIdOk() (*string, bool)`

GetApIdOk returns a tuple with the ApId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApId

`func (o *ClientWirelessStats) SetApId(v string)`

SetApId sets ApId field to given value.


### GetApMac

`func (o *ClientWirelessStats) GetApMac() string`

GetApMac returns the ApMac field if non-nil, zero value otherwise.

### GetApMacOk

`func (o *ClientWirelessStats) GetApMacOk() (*string, bool)`

GetApMacOk returns a tuple with the ApMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApMac

`func (o *ClientWirelessStats) SetApMac(v string)`

SetApMac sets ApMac field to given value.


### GetBand

`func (o *ClientWirelessStats) GetBand() Dot11Band`

GetBand returns the Band field if non-nil, zero value otherwise.

### GetBandOk

`func (o *ClientWirelessStats) GetBandOk() (*Dot11Band, bool)`

GetBandOk returns a tuple with the Band field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand

`func (o *ClientWirelessStats) SetBand(v Dot11Band)`

SetBand sets Band field to given value.


### GetChannel

`func (o *ClientWirelessStats) GetChannel() int32`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ClientWirelessStats) GetChannelOk() (*int32, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ClientWirelessStats) SetChannel(v int32)`

SetChannel sets Channel field to given value.


### GetDualBand

`func (o *ClientWirelessStats) GetDualBand() bool`

GetDualBand returns the DualBand field if non-nil, zero value otherwise.

### GetDualBandOk

`func (o *ClientWirelessStats) GetDualBandOk() (*bool, bool)`

GetDualBandOk returns a tuple with the DualBand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDualBand

`func (o *ClientWirelessStats) SetDualBand(v bool)`

SetDualBand sets DualBand field to given value.


### GetFamily

`func (o *ClientWirelessStats) GetFamily() string`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *ClientWirelessStats) GetFamilyOk() (*string, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *ClientWirelessStats) SetFamily(v string)`

SetFamily sets Family field to given value.


### GetGuest

`func (o *ClientWirelessStats) GetGuest() ClientWirelessStatsGuest`

GetGuest returns the Guest field if non-nil, zero value otherwise.

### GetGuestOk

`func (o *ClientWirelessStats) GetGuestOk() (*ClientWirelessStatsGuest, bool)`

GetGuestOk returns a tuple with the Guest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuest

`func (o *ClientWirelessStats) SetGuest(v ClientWirelessStatsGuest)`

SetGuest sets Guest field to given value.

### HasGuest

`func (o *ClientWirelessStats) HasGuest() bool`

HasGuest returns a boolean if a field has been set.

### GetHostname

`func (o *ClientWirelessStats) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ClientWirelessStats) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ClientWirelessStats) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetIdleTime

`func (o *ClientWirelessStats) GetIdleTime() float32`

GetIdleTime returns the IdleTime field if non-nil, zero value otherwise.

### GetIdleTimeOk

`func (o *ClientWirelessStats) GetIdleTimeOk() (*float32, bool)`

GetIdleTimeOk returns a tuple with the IdleTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleTime

`func (o *ClientWirelessStats) SetIdleTime(v float32)`

SetIdleTime sets IdleTime field to given value.


### GetIp

`func (o *ClientWirelessStats) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ClientWirelessStats) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ClientWirelessStats) SetIp(v string)`

SetIp sets Ip field to given value.


### GetIsGuest

`func (o *ClientWirelessStats) GetIsGuest() bool`

GetIsGuest returns the IsGuest field if non-nil, zero value otherwise.

### GetIsGuestOk

`func (o *ClientWirelessStats) GetIsGuestOk() (*bool, bool)`

GetIsGuestOk returns a tuple with the IsGuest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGuest

`func (o *ClientWirelessStats) SetIsGuest(v bool)`

SetIsGuest sets IsGuest field to given value.


### GetKeyMgmt

`func (o *ClientWirelessStats) GetKeyMgmt() string`

GetKeyMgmt returns the KeyMgmt field if non-nil, zero value otherwise.

### GetKeyMgmtOk

`func (o *ClientWirelessStats) GetKeyMgmtOk() (*string, bool)`

GetKeyMgmtOk returns a tuple with the KeyMgmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyMgmt

`func (o *ClientWirelessStats) SetKeyMgmt(v string)`

SetKeyMgmt sets KeyMgmt field to given value.


### GetLastSeen

`func (o *ClientWirelessStats) GetLastSeen() float32`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *ClientWirelessStats) GetLastSeenOk() (*float32, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *ClientWirelessStats) SetLastSeen(v float32)`

SetLastSeen sets LastSeen field to given value.


### GetMac

`func (o *ClientWirelessStats) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *ClientWirelessStats) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *ClientWirelessStats) SetMac(v string)`

SetMac sets Mac field to given value.


### GetManufacture

`func (o *ClientWirelessStats) GetManufacture() string`

GetManufacture returns the Manufacture field if non-nil, zero value otherwise.

### GetManufactureOk

`func (o *ClientWirelessStats) GetManufactureOk() (*string, bool)`

GetManufactureOk returns a tuple with the Manufacture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacture

`func (o *ClientWirelessStats) SetManufacture(v string)`

SetManufacture sets Manufacture field to given value.


### GetMapId

`func (o *ClientWirelessStats) GetMapId() string`

GetMapId returns the MapId field if non-nil, zero value otherwise.

### GetMapIdOk

`func (o *ClientWirelessStats) GetMapIdOk() (*string, bool)`

GetMapIdOk returns a tuple with the MapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapId

`func (o *ClientWirelessStats) SetMapId(v string)`

SetMapId sets MapId field to given value.

### HasMapId

`func (o *ClientWirelessStats) HasMapId() bool`

HasMapId returns a boolean if a field has been set.

### GetModel

`func (o *ClientWirelessStats) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ClientWirelessStats) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ClientWirelessStats) SetModel(v string)`

SetModel sets Model field to given value.


### GetNumLocatingAps

`func (o *ClientWirelessStats) GetNumLocatingAps() int32`

GetNumLocatingAps returns the NumLocatingAps field if non-nil, zero value otherwise.

### GetNumLocatingApsOk

`func (o *ClientWirelessStats) GetNumLocatingApsOk() (*int32, bool)`

GetNumLocatingApsOk returns a tuple with the NumLocatingAps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumLocatingAps

`func (o *ClientWirelessStats) SetNumLocatingAps(v int32)`

SetNumLocatingAps sets NumLocatingAps field to given value.

### HasNumLocatingAps

`func (o *ClientWirelessStats) HasNumLocatingAps() bool`

HasNumLocatingAps returns a boolean if a field has been set.

### GetOs

`func (o *ClientWirelessStats) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *ClientWirelessStats) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *ClientWirelessStats) SetOs(v string)`

SetOs sets Os field to given value.


### GetPowerSaving

`func (o *ClientWirelessStats) GetPowerSaving() bool`

GetPowerSaving returns the PowerSaving field if non-nil, zero value otherwise.

### GetPowerSavingOk

`func (o *ClientWirelessStats) GetPowerSavingOk() (*bool, bool)`

GetPowerSavingOk returns a tuple with the PowerSaving field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerSaving

`func (o *ClientWirelessStats) SetPowerSaving(v bool)`

SetPowerSaving sets PowerSaving field to given value.


### GetProto

`func (o *ClientWirelessStats) GetProto() Dot11Proto`

GetProto returns the Proto field if non-nil, zero value otherwise.

### GetProtoOk

`func (o *ClientWirelessStats) GetProtoOk() (*Dot11Proto, bool)`

GetProtoOk returns a tuple with the Proto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProto

`func (o *ClientWirelessStats) SetProto(v Dot11Proto)`

SetProto sets Proto field to given value.


### GetPskId

`func (o *ClientWirelessStats) GetPskId() string`

GetPskId returns the PskId field if non-nil, zero value otherwise.

### GetPskIdOk

`func (o *ClientWirelessStats) GetPskIdOk() (*string, bool)`

GetPskIdOk returns a tuple with the PskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPskId

`func (o *ClientWirelessStats) SetPskId(v string)`

SetPskId sets PskId field to given value.

### HasPskId

`func (o *ClientWirelessStats) HasPskId() bool`

HasPskId returns a boolean if a field has been set.

### GetRssi

`func (o *ClientWirelessStats) GetRssi() float32`

GetRssi returns the Rssi field if non-nil, zero value otherwise.

### GetRssiOk

`func (o *ClientWirelessStats) GetRssiOk() (*float32, bool)`

GetRssiOk returns a tuple with the Rssi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssi

`func (o *ClientWirelessStats) SetRssi(v float32)`

SetRssi sets Rssi field to given value.


### GetRssizones

`func (o *ClientWirelessStats) GetRssizones() []ClientWirelessStatsRssiZone`

GetRssizones returns the Rssizones field if non-nil, zero value otherwise.

### GetRssizonesOk

`func (o *ClientWirelessStats) GetRssizonesOk() (*[]ClientWirelessStatsRssiZone, bool)`

GetRssizonesOk returns a tuple with the Rssizones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssizones

`func (o *ClientWirelessStats) SetRssizones(v []ClientWirelessStatsRssiZone)`

SetRssizones sets Rssizones field to given value.

### HasRssizones

`func (o *ClientWirelessStats) HasRssizones() bool`

HasRssizones returns a boolean if a field has been set.

### GetRxBps

`func (o *ClientWirelessStats) GetRxBps() float32`

GetRxBps returns the RxBps field if non-nil, zero value otherwise.

### GetRxBpsOk

`func (o *ClientWirelessStats) GetRxBpsOk() (*float32, bool)`

GetRxBpsOk returns a tuple with the RxBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxBps

`func (o *ClientWirelessStats) SetRxBps(v float32)`

SetRxBps sets RxBps field to given value.


### GetRxBytes

`func (o *ClientWirelessStats) GetRxBytes() float32`

GetRxBytes returns the RxBytes field if non-nil, zero value otherwise.

### GetRxBytesOk

`func (o *ClientWirelessStats) GetRxBytesOk() (*float32, bool)`

GetRxBytesOk returns a tuple with the RxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxBytes

`func (o *ClientWirelessStats) SetRxBytes(v float32)`

SetRxBytes sets RxBytes field to given value.


### GetRxPackets

`func (o *ClientWirelessStats) GetRxPackets() float32`

GetRxPackets returns the RxPackets field if non-nil, zero value otherwise.

### GetRxPacketsOk

`func (o *ClientWirelessStats) GetRxPacketsOk() (*float32, bool)`

GetRxPacketsOk returns a tuple with the RxPackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxPackets

`func (o *ClientWirelessStats) SetRxPackets(v float32)`

SetRxPackets sets RxPackets field to given value.


### GetRxRate

`func (o *ClientWirelessStats) GetRxRate() float32`

GetRxRate returns the RxRate field if non-nil, zero value otherwise.

### GetRxRateOk

`func (o *ClientWirelessStats) GetRxRateOk() (*float32, bool)`

GetRxRateOk returns a tuple with the RxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxRate

`func (o *ClientWirelessStats) SetRxRate(v float32)`

SetRxRate sets RxRate field to given value.


### GetRxRetries

`func (o *ClientWirelessStats) GetRxRetries() float32`

GetRxRetries returns the RxRetries field if non-nil, zero value otherwise.

### GetRxRetriesOk

`func (o *ClientWirelessStats) GetRxRetriesOk() (*float32, bool)`

GetRxRetriesOk returns a tuple with the RxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxRetries

`func (o *ClientWirelessStats) SetRxRetries(v float32)`

SetRxRetries sets RxRetries field to given value.


### GetSnr

`func (o *ClientWirelessStats) GetSnr() float32`

GetSnr returns the Snr field if non-nil, zero value otherwise.

### GetSnrOk

`func (o *ClientWirelessStats) GetSnrOk() (*float32, bool)`

GetSnrOk returns a tuple with the Snr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnr

`func (o *ClientWirelessStats) SetSnr(v float32)`

SetSnr sets Snr field to given value.


### GetSsid

`func (o *ClientWirelessStats) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *ClientWirelessStats) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *ClientWirelessStats) SetSsid(v string)`

SetSsid sets Ssid field to given value.


### GetTxBps

`func (o *ClientWirelessStats) GetTxBps() float32`

GetTxBps returns the TxBps field if non-nil, zero value otherwise.

### GetTxBpsOk

`func (o *ClientWirelessStats) GetTxBpsOk() (*float32, bool)`

GetTxBpsOk returns a tuple with the TxBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxBps

`func (o *ClientWirelessStats) SetTxBps(v float32)`

SetTxBps sets TxBps field to given value.


### GetTxBytes

`func (o *ClientWirelessStats) GetTxBytes() float32`

GetTxBytes returns the TxBytes field if non-nil, zero value otherwise.

### GetTxBytesOk

`func (o *ClientWirelessStats) GetTxBytesOk() (*float32, bool)`

GetTxBytesOk returns a tuple with the TxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxBytes

`func (o *ClientWirelessStats) SetTxBytes(v float32)`

SetTxBytes sets TxBytes field to given value.


### GetTxPackets

`func (o *ClientWirelessStats) GetTxPackets() float32`

GetTxPackets returns the TxPackets field if non-nil, zero value otherwise.

### GetTxPacketsOk

`func (o *ClientWirelessStats) GetTxPacketsOk() (*float32, bool)`

GetTxPacketsOk returns a tuple with the TxPackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxPackets

`func (o *ClientWirelessStats) SetTxPackets(v float32)`

SetTxPackets sets TxPackets field to given value.


### GetTxRate

`func (o *ClientWirelessStats) GetTxRate() float32`

GetTxRate returns the TxRate field if non-nil, zero value otherwise.

### GetTxRateOk

`func (o *ClientWirelessStats) GetTxRateOk() (*float32, bool)`

GetTxRateOk returns a tuple with the TxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxRate

`func (o *ClientWirelessStats) SetTxRate(v float32)`

SetTxRate sets TxRate field to given value.


### GetTxRetries

`func (o *ClientWirelessStats) GetTxRetries() float32`

GetTxRetries returns the TxRetries field if non-nil, zero value otherwise.

### GetTxRetriesOk

`func (o *ClientWirelessStats) GetTxRetriesOk() (*float32, bool)`

GetTxRetriesOk returns a tuple with the TxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxRetries

`func (o *ClientWirelessStats) SetTxRetries(v float32)`

SetTxRetries sets TxRetries field to given value.


### GetType

`func (o *ClientWirelessStats) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClientWirelessStats) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClientWirelessStats) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ClientWirelessStats) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUptime

`func (o *ClientWirelessStats) GetUptime() float32`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *ClientWirelessStats) GetUptimeOk() (*float32, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *ClientWirelessStats) SetUptime(v float32)`

SetUptime sets Uptime field to given value.


### GetUsername

`func (o *ClientWirelessStats) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ClientWirelessStats) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ClientWirelessStats) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetVbeacons

`func (o *ClientWirelessStats) GetVbeacons() []ClientWirelessStatsVbeacon`

GetVbeacons returns the Vbeacons field if non-nil, zero value otherwise.

### GetVbeaconsOk

`func (o *ClientWirelessStats) GetVbeaconsOk() (*[]ClientWirelessStatsVbeacon, bool)`

GetVbeaconsOk returns a tuple with the Vbeacons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVbeacons

`func (o *ClientWirelessStats) SetVbeacons(v []ClientWirelessStatsVbeacon)`

SetVbeacons sets Vbeacons field to given value.

### HasVbeacons

`func (o *ClientWirelessStats) HasVbeacons() bool`

HasVbeacons returns a boolean if a field has been set.

### GetVlanId

`func (o *ClientWirelessStats) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *ClientWirelessStats) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *ClientWirelessStats) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *ClientWirelessStats) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetWlanId

`func (o *ClientWirelessStats) GetWlanId() string`

GetWlanId returns the WlanId field if non-nil, zero value otherwise.

### GetWlanIdOk

`func (o *ClientWirelessStats) GetWlanIdOk() (*string, bool)`

GetWlanIdOk returns a tuple with the WlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWlanId

`func (o *ClientWirelessStats) SetWlanId(v string)`

SetWlanId sets WlanId field to given value.


### GetWxruleId

`func (o *ClientWirelessStats) GetWxruleId() string`

GetWxruleId returns the WxruleId field if non-nil, zero value otherwise.

### GetWxruleIdOk

`func (o *ClientWirelessStats) GetWxruleIdOk() (*string, bool)`

GetWxruleIdOk returns a tuple with the WxruleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWxruleId

`func (o *ClientWirelessStats) SetWxruleId(v string)`

SetWxruleId sets WxruleId field to given value.

### HasWxruleId

`func (o *ClientWirelessStats) HasWxruleId() bool`

HasWxruleId returns a boolean if a field has been set.

### GetWxruleUsage

`func (o *ClientWirelessStats) GetWxruleUsage() []ClientWirelessStatsWxruleUsage`

GetWxruleUsage returns the WxruleUsage field if non-nil, zero value otherwise.

### GetWxruleUsageOk

`func (o *ClientWirelessStats) GetWxruleUsageOk() (*[]ClientWirelessStatsWxruleUsage, bool)`

GetWxruleUsageOk returns a tuple with the WxruleUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWxruleUsage

`func (o *ClientWirelessStats) SetWxruleUsage(v []ClientWirelessStatsWxruleUsage)`

SetWxruleUsage sets WxruleUsage field to given value.

### HasWxruleUsage

`func (o *ClientWirelessStats) HasWxruleUsage() bool`

HasWxruleUsage returns a boolean if a field has been set.

### GetX

`func (o *ClientWirelessStats) GetX() float32`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *ClientWirelessStats) GetXOk() (*float32, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *ClientWirelessStats) SetX(v float32)`

SetX sets X field to given value.

### HasX

`func (o *ClientWirelessStats) HasX() bool`

HasX returns a boolean if a field has been set.

### GetXM

`func (o *ClientWirelessStats) GetXM() float32`

GetXM returns the XM field if non-nil, zero value otherwise.

### GetXMOk

`func (o *ClientWirelessStats) GetXMOk() (*float32, bool)`

GetXMOk returns a tuple with the XM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXM

`func (o *ClientWirelessStats) SetXM(v float32)`

SetXM sets XM field to given value.

### HasXM

`func (o *ClientWirelessStats) HasXM() bool`

HasXM returns a boolean if a field has been set.

### GetY

`func (o *ClientWirelessStats) GetY() float32`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *ClientWirelessStats) GetYOk() (*float32, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *ClientWirelessStats) SetY(v float32)`

SetY sets Y field to given value.

### HasY

`func (o *ClientWirelessStats) HasY() bool`

HasY returns a boolean if a field has been set.

### GetYM

`func (o *ClientWirelessStats) GetYM() float32`

GetYM returns the YM field if non-nil, zero value otherwise.

### GetYMOk

`func (o *ClientWirelessStats) GetYMOk() (*float32, bool)`

GetYMOk returns a tuple with the YM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYM

`func (o *ClientWirelessStats) SetYM(v float32)`

SetYM sets YM field to given value.

### HasYM

`func (o *ClientWirelessStats) HasYM() bool`

HasYM returns a boolean if a field has been set.

### GetZones

`func (o *ClientWirelessStats) GetZones() []ClientWirelessStatsZone`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *ClientWirelessStats) GetZonesOk() (*[]ClientWirelessStatsZone, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *ClientWirelessStats) SetZones(v []ClientWirelessStatsZone)`

SetZones sets Zones field to given value.

### HasZones

`func (o *ClientWirelessStats) HasZones() bool`

HasZones returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


