# ClientsWanStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**When** | Pointer to **string** |  | [optional] 
**Hostname** | Pointer to **[]string** |  | [optional] 
**Ip** | Pointer to **[]string** |  | [optional] 
**LastHostname** | Pointer to **string** |  | [optional] 
**LastIp** | Pointer to **string** |  | [optional] 
**Mfg** | Pointer to **string** |  | [optional] 
**Network** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**SiteId** | Pointer to **string** |  | [optional] [readonly] 
**Wcid** | Pointer to **string** |  | [optional] 

## Methods

### NewClientsWanStats

`func NewClientsWanStats() *ClientsWanStats`

NewClientsWanStats instantiates a new ClientsWanStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientsWanStatsWithDefaults

`func NewClientsWanStatsWithDefaults() *ClientsWanStats`

NewClientsWanStatsWithDefaults instantiates a new ClientsWanStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWhen

`func (o *ClientsWanStats) GetWhen() string`

GetWhen returns the When field if non-nil, zero value otherwise.

### GetWhenOk

`func (o *ClientsWanStats) GetWhenOk() (*string, bool)`

GetWhenOk returns a tuple with the When field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhen

`func (o *ClientsWanStats) SetWhen(v string)`

SetWhen sets When field to given value.

### HasWhen

`func (o *ClientsWanStats) HasWhen() bool`

HasWhen returns a boolean if a field has been set.

### GetHostname

`func (o *ClientsWanStats) GetHostname() []string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ClientsWanStats) GetHostnameOk() (*[]string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ClientsWanStats) SetHostname(v []string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ClientsWanStats) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetIp

`func (o *ClientsWanStats) GetIp() []string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ClientsWanStats) GetIpOk() (*[]string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ClientsWanStats) SetIp(v []string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *ClientsWanStats) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetLastHostname

`func (o *ClientsWanStats) GetLastHostname() string`

GetLastHostname returns the LastHostname field if non-nil, zero value otherwise.

### GetLastHostnameOk

`func (o *ClientsWanStats) GetLastHostnameOk() (*string, bool)`

GetLastHostnameOk returns a tuple with the LastHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastHostname

`func (o *ClientsWanStats) SetLastHostname(v string)`

SetLastHostname sets LastHostname field to given value.

### HasLastHostname

`func (o *ClientsWanStats) HasLastHostname() bool`

HasLastHostname returns a boolean if a field has been set.

### GetLastIp

`func (o *ClientsWanStats) GetLastIp() string`

GetLastIp returns the LastIp field if non-nil, zero value otherwise.

### GetLastIpOk

`func (o *ClientsWanStats) GetLastIpOk() (*string, bool)`

GetLastIpOk returns a tuple with the LastIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastIp

`func (o *ClientsWanStats) SetLastIp(v string)`

SetLastIp sets LastIp field to given value.

### HasLastIp

`func (o *ClientsWanStats) HasLastIp() bool`

HasLastIp returns a boolean if a field has been set.

### GetMfg

`func (o *ClientsWanStats) GetMfg() string`

GetMfg returns the Mfg field if non-nil, zero value otherwise.

### GetMfgOk

`func (o *ClientsWanStats) GetMfgOk() (*string, bool)`

GetMfgOk returns a tuple with the Mfg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfg

`func (o *ClientsWanStats) SetMfg(v string)`

SetMfg sets Mfg field to given value.

### HasMfg

`func (o *ClientsWanStats) HasMfg() bool`

HasMfg returns a boolean if a field has been set.

### GetNetwork

`func (o *ClientsWanStats) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *ClientsWanStats) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *ClientsWanStats) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *ClientsWanStats) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetOrgId

`func (o *ClientsWanStats) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ClientsWanStats) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ClientsWanStats) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ClientsWanStats) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetSiteId

`func (o *ClientsWanStats) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *ClientsWanStats) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *ClientsWanStats) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *ClientsWanStats) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetWcid

`func (o *ClientsWanStats) GetWcid() string`

GetWcid returns the Wcid field if non-nil, zero value otherwise.

### GetWcidOk

`func (o *ClientsWanStats) GetWcidOk() (*string, bool)`

GetWcidOk returns a tuple with the Wcid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWcid

`func (o *ClientsWanStats) SetWcid(v string)`

SetWcid sets Wcid field to given value.

### HasWcid

`func (o *ClientsWanStats) HasWcid() bool`

HasWcid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


