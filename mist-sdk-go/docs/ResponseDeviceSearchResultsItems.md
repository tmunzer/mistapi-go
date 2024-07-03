# ResponseDeviceSearchResultsItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Band24Bandwith** | Pointer to **string** | Bandwith of band_24 | [optional] 
**Band24Channel** | Pointer to **int32** | Channel of band_24 | [optional] 
**Band24Power** | Pointer to **int32** |  | [optional] 
**Band5Bandwith** | Pointer to **string** | Bandwith of band_5 | [optional] 
**Band5Channel** | Pointer to **int32** | Channel of band_5 | [optional] 
**Band5Power** | Pointer to **int32** |  | [optional] 
**Band6Bandwith** | Pointer to **string** |  | [optional] 
**Band6Channel** | Pointer to **int32** | Channel of band_6 | [optional] 
**Band6Power** | Pointer to **int32** |  | [optional] 
**Eth0PortSpeed** | Pointer to **int32** | Port speed of eth0 | [optional] 
**ExtIp** | Pointer to **string** |  | [optional] 
**Hostname** | Pointer to **[]string** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**LldpMgmtAddr** | Pointer to **string** | LLDP management ip address | [optional] 
**LldpPortDesc** | Pointer to **string** |  | [optional] 
**LldpPortId** | Pointer to **string** | LLDP port id | [optional] 
**LldpPowerAllocated** | Pointer to **int32** |  | [optional] 
**LldpPowerDraw** | Pointer to **int32** |  | [optional] 
**LldpSystemDesc** | Pointer to **string** | LLDP system description | [optional] 
**LldpSystemName** | Pointer to **string** | LLDP system name | [optional] 
**Mac** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**MxedgeId** | Pointer to **string** | Mist Edge id, if AP is connecting to a Mist Edge | [optional] 
**MxtunnelStatus** | Pointer to **string** | MxTunnel status | [optional] 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**PowerConstrained** | Pointer to **bool** |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] [readonly] 
**Sku** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **float32** |  | [optional] 
**Uptime** | Pointer to **int32** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**LastHostname** | Pointer to **string** |  | [optional] 
**NumMembers** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewResponseDeviceSearchResultsItems

`func NewResponseDeviceSearchResultsItems() *ResponseDeviceSearchResultsItems`

NewResponseDeviceSearchResultsItems instantiates a new ResponseDeviceSearchResultsItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseDeviceSearchResultsItemsWithDefaults

`func NewResponseDeviceSearchResultsItemsWithDefaults() *ResponseDeviceSearchResultsItems`

NewResponseDeviceSearchResultsItemsWithDefaults instantiates a new ResponseDeviceSearchResultsItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBand24Bandwith

`func (o *ResponseDeviceSearchResultsItems) GetBand24Bandwith() string`

GetBand24Bandwith returns the Band24Bandwith field if non-nil, zero value otherwise.

### GetBand24BandwithOk

`func (o *ResponseDeviceSearchResultsItems) GetBand24BandwithOk() (*string, bool)`

GetBand24BandwithOk returns a tuple with the Band24Bandwith field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand24Bandwith

`func (o *ResponseDeviceSearchResultsItems) SetBand24Bandwith(v string)`

SetBand24Bandwith sets Band24Bandwith field to given value.

### HasBand24Bandwith

`func (o *ResponseDeviceSearchResultsItems) HasBand24Bandwith() bool`

HasBand24Bandwith returns a boolean if a field has been set.

### GetBand24Channel

`func (o *ResponseDeviceSearchResultsItems) GetBand24Channel() int32`

GetBand24Channel returns the Band24Channel field if non-nil, zero value otherwise.

### GetBand24ChannelOk

`func (o *ResponseDeviceSearchResultsItems) GetBand24ChannelOk() (*int32, bool)`

GetBand24ChannelOk returns a tuple with the Band24Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand24Channel

`func (o *ResponseDeviceSearchResultsItems) SetBand24Channel(v int32)`

SetBand24Channel sets Band24Channel field to given value.

### HasBand24Channel

`func (o *ResponseDeviceSearchResultsItems) HasBand24Channel() bool`

HasBand24Channel returns a boolean if a field has been set.

### GetBand24Power

`func (o *ResponseDeviceSearchResultsItems) GetBand24Power() int32`

GetBand24Power returns the Band24Power field if non-nil, zero value otherwise.

### GetBand24PowerOk

`func (o *ResponseDeviceSearchResultsItems) GetBand24PowerOk() (*int32, bool)`

GetBand24PowerOk returns a tuple with the Band24Power field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand24Power

`func (o *ResponseDeviceSearchResultsItems) SetBand24Power(v int32)`

SetBand24Power sets Band24Power field to given value.

### HasBand24Power

`func (o *ResponseDeviceSearchResultsItems) HasBand24Power() bool`

HasBand24Power returns a boolean if a field has been set.

### GetBand5Bandwith

`func (o *ResponseDeviceSearchResultsItems) GetBand5Bandwith() string`

GetBand5Bandwith returns the Band5Bandwith field if non-nil, zero value otherwise.

### GetBand5BandwithOk

`func (o *ResponseDeviceSearchResultsItems) GetBand5BandwithOk() (*string, bool)`

GetBand5BandwithOk returns a tuple with the Band5Bandwith field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand5Bandwith

`func (o *ResponseDeviceSearchResultsItems) SetBand5Bandwith(v string)`

SetBand5Bandwith sets Band5Bandwith field to given value.

### HasBand5Bandwith

`func (o *ResponseDeviceSearchResultsItems) HasBand5Bandwith() bool`

HasBand5Bandwith returns a boolean if a field has been set.

### GetBand5Channel

`func (o *ResponseDeviceSearchResultsItems) GetBand5Channel() int32`

GetBand5Channel returns the Band5Channel field if non-nil, zero value otherwise.

### GetBand5ChannelOk

`func (o *ResponseDeviceSearchResultsItems) GetBand5ChannelOk() (*int32, bool)`

GetBand5ChannelOk returns a tuple with the Band5Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand5Channel

`func (o *ResponseDeviceSearchResultsItems) SetBand5Channel(v int32)`

SetBand5Channel sets Band5Channel field to given value.

### HasBand5Channel

`func (o *ResponseDeviceSearchResultsItems) HasBand5Channel() bool`

HasBand5Channel returns a boolean if a field has been set.

### GetBand5Power

`func (o *ResponseDeviceSearchResultsItems) GetBand5Power() int32`

GetBand5Power returns the Band5Power field if non-nil, zero value otherwise.

### GetBand5PowerOk

`func (o *ResponseDeviceSearchResultsItems) GetBand5PowerOk() (*int32, bool)`

GetBand5PowerOk returns a tuple with the Band5Power field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand5Power

`func (o *ResponseDeviceSearchResultsItems) SetBand5Power(v int32)`

SetBand5Power sets Band5Power field to given value.

### HasBand5Power

`func (o *ResponseDeviceSearchResultsItems) HasBand5Power() bool`

HasBand5Power returns a boolean if a field has been set.

### GetBand6Bandwith

`func (o *ResponseDeviceSearchResultsItems) GetBand6Bandwith() string`

GetBand6Bandwith returns the Band6Bandwith field if non-nil, zero value otherwise.

### GetBand6BandwithOk

`func (o *ResponseDeviceSearchResultsItems) GetBand6BandwithOk() (*string, bool)`

GetBand6BandwithOk returns a tuple with the Band6Bandwith field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand6Bandwith

`func (o *ResponseDeviceSearchResultsItems) SetBand6Bandwith(v string)`

SetBand6Bandwith sets Band6Bandwith field to given value.

### HasBand6Bandwith

`func (o *ResponseDeviceSearchResultsItems) HasBand6Bandwith() bool`

HasBand6Bandwith returns a boolean if a field has been set.

### GetBand6Channel

`func (o *ResponseDeviceSearchResultsItems) GetBand6Channel() int32`

GetBand6Channel returns the Band6Channel field if non-nil, zero value otherwise.

### GetBand6ChannelOk

`func (o *ResponseDeviceSearchResultsItems) GetBand6ChannelOk() (*int32, bool)`

GetBand6ChannelOk returns a tuple with the Band6Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand6Channel

`func (o *ResponseDeviceSearchResultsItems) SetBand6Channel(v int32)`

SetBand6Channel sets Band6Channel field to given value.

### HasBand6Channel

`func (o *ResponseDeviceSearchResultsItems) HasBand6Channel() bool`

HasBand6Channel returns a boolean if a field has been set.

### GetBand6Power

`func (o *ResponseDeviceSearchResultsItems) GetBand6Power() int32`

GetBand6Power returns the Band6Power field if non-nil, zero value otherwise.

### GetBand6PowerOk

`func (o *ResponseDeviceSearchResultsItems) GetBand6PowerOk() (*int32, bool)`

GetBand6PowerOk returns a tuple with the Band6Power field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand6Power

`func (o *ResponseDeviceSearchResultsItems) SetBand6Power(v int32)`

SetBand6Power sets Band6Power field to given value.

### HasBand6Power

`func (o *ResponseDeviceSearchResultsItems) HasBand6Power() bool`

HasBand6Power returns a boolean if a field has been set.

### GetEth0PortSpeed

`func (o *ResponseDeviceSearchResultsItems) GetEth0PortSpeed() int32`

GetEth0PortSpeed returns the Eth0PortSpeed field if non-nil, zero value otherwise.

### GetEth0PortSpeedOk

`func (o *ResponseDeviceSearchResultsItems) GetEth0PortSpeedOk() (*int32, bool)`

GetEth0PortSpeedOk returns a tuple with the Eth0PortSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEth0PortSpeed

`func (o *ResponseDeviceSearchResultsItems) SetEth0PortSpeed(v int32)`

SetEth0PortSpeed sets Eth0PortSpeed field to given value.

### HasEth0PortSpeed

`func (o *ResponseDeviceSearchResultsItems) HasEth0PortSpeed() bool`

HasEth0PortSpeed returns a boolean if a field has been set.

### GetExtIp

`func (o *ResponseDeviceSearchResultsItems) GetExtIp() string`

GetExtIp returns the ExtIp field if non-nil, zero value otherwise.

### GetExtIpOk

`func (o *ResponseDeviceSearchResultsItems) GetExtIpOk() (*string, bool)`

GetExtIpOk returns a tuple with the ExtIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtIp

`func (o *ResponseDeviceSearchResultsItems) SetExtIp(v string)`

SetExtIp sets ExtIp field to given value.

### HasExtIp

`func (o *ResponseDeviceSearchResultsItems) HasExtIp() bool`

HasExtIp returns a boolean if a field has been set.

### GetHostname

`func (o *ResponseDeviceSearchResultsItems) GetHostname() []string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ResponseDeviceSearchResultsItems) GetHostnameOk() (*[]string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ResponseDeviceSearchResultsItems) SetHostname(v []string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ResponseDeviceSearchResultsItems) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetIp

`func (o *ResponseDeviceSearchResultsItems) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ResponseDeviceSearchResultsItems) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ResponseDeviceSearchResultsItems) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *ResponseDeviceSearchResultsItems) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetLldpMgmtAddr

`func (o *ResponseDeviceSearchResultsItems) GetLldpMgmtAddr() string`

GetLldpMgmtAddr returns the LldpMgmtAddr field if non-nil, zero value otherwise.

### GetLldpMgmtAddrOk

`func (o *ResponseDeviceSearchResultsItems) GetLldpMgmtAddrOk() (*string, bool)`

GetLldpMgmtAddrOk returns a tuple with the LldpMgmtAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldpMgmtAddr

`func (o *ResponseDeviceSearchResultsItems) SetLldpMgmtAddr(v string)`

SetLldpMgmtAddr sets LldpMgmtAddr field to given value.

### HasLldpMgmtAddr

`func (o *ResponseDeviceSearchResultsItems) HasLldpMgmtAddr() bool`

HasLldpMgmtAddr returns a boolean if a field has been set.

### GetLldpPortDesc

`func (o *ResponseDeviceSearchResultsItems) GetLldpPortDesc() string`

GetLldpPortDesc returns the LldpPortDesc field if non-nil, zero value otherwise.

### GetLldpPortDescOk

`func (o *ResponseDeviceSearchResultsItems) GetLldpPortDescOk() (*string, bool)`

GetLldpPortDescOk returns a tuple with the LldpPortDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldpPortDesc

`func (o *ResponseDeviceSearchResultsItems) SetLldpPortDesc(v string)`

SetLldpPortDesc sets LldpPortDesc field to given value.

### HasLldpPortDesc

`func (o *ResponseDeviceSearchResultsItems) HasLldpPortDesc() bool`

HasLldpPortDesc returns a boolean if a field has been set.

### GetLldpPortId

`func (o *ResponseDeviceSearchResultsItems) GetLldpPortId() string`

GetLldpPortId returns the LldpPortId field if non-nil, zero value otherwise.

### GetLldpPortIdOk

`func (o *ResponseDeviceSearchResultsItems) GetLldpPortIdOk() (*string, bool)`

GetLldpPortIdOk returns a tuple with the LldpPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldpPortId

`func (o *ResponseDeviceSearchResultsItems) SetLldpPortId(v string)`

SetLldpPortId sets LldpPortId field to given value.

### HasLldpPortId

`func (o *ResponseDeviceSearchResultsItems) HasLldpPortId() bool`

HasLldpPortId returns a boolean if a field has been set.

### GetLldpPowerAllocated

`func (o *ResponseDeviceSearchResultsItems) GetLldpPowerAllocated() int32`

GetLldpPowerAllocated returns the LldpPowerAllocated field if non-nil, zero value otherwise.

### GetLldpPowerAllocatedOk

`func (o *ResponseDeviceSearchResultsItems) GetLldpPowerAllocatedOk() (*int32, bool)`

GetLldpPowerAllocatedOk returns a tuple with the LldpPowerAllocated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldpPowerAllocated

`func (o *ResponseDeviceSearchResultsItems) SetLldpPowerAllocated(v int32)`

SetLldpPowerAllocated sets LldpPowerAllocated field to given value.

### HasLldpPowerAllocated

`func (o *ResponseDeviceSearchResultsItems) HasLldpPowerAllocated() bool`

HasLldpPowerAllocated returns a boolean if a field has been set.

### GetLldpPowerDraw

`func (o *ResponseDeviceSearchResultsItems) GetLldpPowerDraw() int32`

GetLldpPowerDraw returns the LldpPowerDraw field if non-nil, zero value otherwise.

### GetLldpPowerDrawOk

`func (o *ResponseDeviceSearchResultsItems) GetLldpPowerDrawOk() (*int32, bool)`

GetLldpPowerDrawOk returns a tuple with the LldpPowerDraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldpPowerDraw

`func (o *ResponseDeviceSearchResultsItems) SetLldpPowerDraw(v int32)`

SetLldpPowerDraw sets LldpPowerDraw field to given value.

### HasLldpPowerDraw

`func (o *ResponseDeviceSearchResultsItems) HasLldpPowerDraw() bool`

HasLldpPowerDraw returns a boolean if a field has been set.

### GetLldpSystemDesc

`func (o *ResponseDeviceSearchResultsItems) GetLldpSystemDesc() string`

GetLldpSystemDesc returns the LldpSystemDesc field if non-nil, zero value otherwise.

### GetLldpSystemDescOk

`func (o *ResponseDeviceSearchResultsItems) GetLldpSystemDescOk() (*string, bool)`

GetLldpSystemDescOk returns a tuple with the LldpSystemDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldpSystemDesc

`func (o *ResponseDeviceSearchResultsItems) SetLldpSystemDesc(v string)`

SetLldpSystemDesc sets LldpSystemDesc field to given value.

### HasLldpSystemDesc

`func (o *ResponseDeviceSearchResultsItems) HasLldpSystemDesc() bool`

HasLldpSystemDesc returns a boolean if a field has been set.

### GetLldpSystemName

`func (o *ResponseDeviceSearchResultsItems) GetLldpSystemName() string`

GetLldpSystemName returns the LldpSystemName field if non-nil, zero value otherwise.

### GetLldpSystemNameOk

`func (o *ResponseDeviceSearchResultsItems) GetLldpSystemNameOk() (*string, bool)`

GetLldpSystemNameOk returns a tuple with the LldpSystemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldpSystemName

`func (o *ResponseDeviceSearchResultsItems) SetLldpSystemName(v string)`

SetLldpSystemName sets LldpSystemName field to given value.

### HasLldpSystemName

`func (o *ResponseDeviceSearchResultsItems) HasLldpSystemName() bool`

HasLldpSystemName returns a boolean if a field has been set.

### GetMac

`func (o *ResponseDeviceSearchResultsItems) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *ResponseDeviceSearchResultsItems) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *ResponseDeviceSearchResultsItems) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *ResponseDeviceSearchResultsItems) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *ResponseDeviceSearchResultsItems) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ResponseDeviceSearchResultsItems) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ResponseDeviceSearchResultsItems) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ResponseDeviceSearchResultsItems) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetMxedgeId

`func (o *ResponseDeviceSearchResultsItems) GetMxedgeId() string`

GetMxedgeId returns the MxedgeId field if non-nil, zero value otherwise.

### GetMxedgeIdOk

`func (o *ResponseDeviceSearchResultsItems) GetMxedgeIdOk() (*string, bool)`

GetMxedgeIdOk returns a tuple with the MxedgeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMxedgeId

`func (o *ResponseDeviceSearchResultsItems) SetMxedgeId(v string)`

SetMxedgeId sets MxedgeId field to given value.

### HasMxedgeId

`func (o *ResponseDeviceSearchResultsItems) HasMxedgeId() bool`

HasMxedgeId returns a boolean if a field has been set.

### GetMxtunnelStatus

`func (o *ResponseDeviceSearchResultsItems) GetMxtunnelStatus() string`

GetMxtunnelStatus returns the MxtunnelStatus field if non-nil, zero value otherwise.

### GetMxtunnelStatusOk

`func (o *ResponseDeviceSearchResultsItems) GetMxtunnelStatusOk() (*string, bool)`

GetMxtunnelStatusOk returns a tuple with the MxtunnelStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMxtunnelStatus

`func (o *ResponseDeviceSearchResultsItems) SetMxtunnelStatus(v string)`

SetMxtunnelStatus sets MxtunnelStatus field to given value.

### HasMxtunnelStatus

`func (o *ResponseDeviceSearchResultsItems) HasMxtunnelStatus() bool`

HasMxtunnelStatus returns a boolean if a field has been set.

### GetOrgId

`func (o *ResponseDeviceSearchResultsItems) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ResponseDeviceSearchResultsItems) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ResponseDeviceSearchResultsItems) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ResponseDeviceSearchResultsItems) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPowerConstrained

`func (o *ResponseDeviceSearchResultsItems) GetPowerConstrained() bool`

GetPowerConstrained returns the PowerConstrained field if non-nil, zero value otherwise.

### GetPowerConstrainedOk

`func (o *ResponseDeviceSearchResultsItems) GetPowerConstrainedOk() (*bool, bool)`

GetPowerConstrainedOk returns a tuple with the PowerConstrained field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerConstrained

`func (o *ResponseDeviceSearchResultsItems) SetPowerConstrained(v bool)`

SetPowerConstrained sets PowerConstrained field to given value.

### HasPowerConstrained

`func (o *ResponseDeviceSearchResultsItems) HasPowerConstrained() bool`

HasPowerConstrained returns a boolean if a field has been set.

### GetSiteId

`func (o *ResponseDeviceSearchResultsItems) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *ResponseDeviceSearchResultsItems) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *ResponseDeviceSearchResultsItems) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *ResponseDeviceSearchResultsItems) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSku

`func (o *ResponseDeviceSearchResultsItems) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *ResponseDeviceSearchResultsItems) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *ResponseDeviceSearchResultsItems) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *ResponseDeviceSearchResultsItems) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetTimestamp

`func (o *ResponseDeviceSearchResultsItems) GetTimestamp() float32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ResponseDeviceSearchResultsItems) GetTimestampOk() (*float32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ResponseDeviceSearchResultsItems) SetTimestamp(v float32)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ResponseDeviceSearchResultsItems) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetUptime

`func (o *ResponseDeviceSearchResultsItems) GetUptime() int32`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *ResponseDeviceSearchResultsItems) GetUptimeOk() (*int32, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *ResponseDeviceSearchResultsItems) SetUptime(v int32)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *ResponseDeviceSearchResultsItems) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetVersion

`func (o *ResponseDeviceSearchResultsItems) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ResponseDeviceSearchResultsItems) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ResponseDeviceSearchResultsItems) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ResponseDeviceSearchResultsItems) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetLastHostname

`func (o *ResponseDeviceSearchResultsItems) GetLastHostname() string`

GetLastHostname returns the LastHostname field if non-nil, zero value otherwise.

### GetLastHostnameOk

`func (o *ResponseDeviceSearchResultsItems) GetLastHostnameOk() (*string, bool)`

GetLastHostnameOk returns a tuple with the LastHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastHostname

`func (o *ResponseDeviceSearchResultsItems) SetLastHostname(v string)`

SetLastHostname sets LastHostname field to given value.

### HasLastHostname

`func (o *ResponseDeviceSearchResultsItems) HasLastHostname() bool`

HasLastHostname returns a boolean if a field has been set.

### GetNumMembers

`func (o *ResponseDeviceSearchResultsItems) GetNumMembers() int32`

GetNumMembers returns the NumMembers field if non-nil, zero value otherwise.

### GetNumMembersOk

`func (o *ResponseDeviceSearchResultsItems) GetNumMembersOk() (*int32, bool)`

GetNumMembersOk returns a tuple with the NumMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumMembers

`func (o *ResponseDeviceSearchResultsItems) SetNumMembers(v int32)`

SetNumMembers sets NumMembers field to given value.

### HasNumMembers

`func (o *ResponseDeviceSearchResultsItems) HasNumMembers() bool`

HasNumMembers returns a boolean if a field has been set.

### GetType

`func (o *ResponseDeviceSearchResultsItems) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponseDeviceSearchResultsItems) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponseDeviceSearchResultsItems) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ResponseDeviceSearchResultsItems) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


