# Mxcluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | Pointer to **float32** |  | [optional] [readonly] 
**ForSite** | Pointer to **bool** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**MistDas** | Pointer to [**MxedgeDas**](MxedgeDas.md) |  | [optional] 
**MistNac** | Pointer to [**MxclusterNac**](MxclusterNac.md) |  | [optional] 
**ModifiedTime** | Pointer to **float32** |  | [optional] [readonly] 
**MxedgeMgmt** | Pointer to [**MxedgeMgmt**](MxedgeMgmt.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**Proxy** | Pointer to [**Proxy**](Proxy.md) |  | [optional] 
**Radsec** | Pointer to [**MxclusterRadsec**](MxclusterRadsec.md) |  | [optional] 
**RadsecTls** | Pointer to [**MxclusterRadsecTls**](MxclusterRadsecTls.md) |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] [readonly] 
**TuntermApSubnets** | Pointer to **[]string** | list of subnets where we allow AP to establish Mist Tunnels from | [optional] 
**TuntermDhcpdConfig** | Pointer to [**TuntermDhcpdConfig**](TuntermDhcpdConfig.md) |  | [optional] 
**TuntermExtraRoutes** | Pointer to [**map[string]MxclusterTuntermExtraRoute**](MxclusterTuntermExtraRoute.md) | extra routes for Mist Tunneled VLANs. Property key is a CIDR | [optional] 
**TuntermHosts** | Pointer to **[]string** | hostnames or IPs where a Mist Tunnel will use as the Peer (i.e. they are reachable from AP) | [optional] 
**TuntermHostsOrder** | Pointer to **[]int32** | list of index of tunterm_hosts | [optional] 
**TuntermHostsSelection** | Pointer to [**MxclusterTuntermHostsSelection**](MxclusterTuntermHostsSelection.md) |  | [optional] [default to MXCLUSTERTUNTERMHOSTSSELECTION_SHUFFLE]
**TuntermMonitoring** | Pointer to [**[][]TuntermMonitoringItem**]([]TuntermMonitoringItem.md) |  | [optional] 
**TuntermMonitoringDisabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewMxcluster

`func NewMxcluster() *Mxcluster`

NewMxcluster instantiates a new Mxcluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMxclusterWithDefaults

`func NewMxclusterWithDefaults() *Mxcluster`

NewMxclusterWithDefaults instantiates a new Mxcluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTime

`func (o *Mxcluster) GetCreatedTime() float32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Mxcluster) GetCreatedTimeOk() (*float32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Mxcluster) SetCreatedTime(v float32)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *Mxcluster) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetForSite

`func (o *Mxcluster) GetForSite() bool`

GetForSite returns the ForSite field if non-nil, zero value otherwise.

### GetForSiteOk

`func (o *Mxcluster) GetForSiteOk() (*bool, bool)`

GetForSiteOk returns a tuple with the ForSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForSite

`func (o *Mxcluster) SetForSite(v bool)`

SetForSite sets ForSite field to given value.

### HasForSite

`func (o *Mxcluster) HasForSite() bool`

HasForSite returns a boolean if a field has been set.

### GetId

`func (o *Mxcluster) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Mxcluster) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Mxcluster) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Mxcluster) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMistDas

`func (o *Mxcluster) GetMistDas() MxedgeDas`

GetMistDas returns the MistDas field if non-nil, zero value otherwise.

### GetMistDasOk

`func (o *Mxcluster) GetMistDasOk() (*MxedgeDas, bool)`

GetMistDasOk returns a tuple with the MistDas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMistDas

`func (o *Mxcluster) SetMistDas(v MxedgeDas)`

SetMistDas sets MistDas field to given value.

### HasMistDas

`func (o *Mxcluster) HasMistDas() bool`

HasMistDas returns a boolean if a field has been set.

### GetMistNac

`func (o *Mxcluster) GetMistNac() MxclusterNac`

GetMistNac returns the MistNac field if non-nil, zero value otherwise.

### GetMistNacOk

`func (o *Mxcluster) GetMistNacOk() (*MxclusterNac, bool)`

GetMistNacOk returns a tuple with the MistNac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMistNac

`func (o *Mxcluster) SetMistNac(v MxclusterNac)`

SetMistNac sets MistNac field to given value.

### HasMistNac

`func (o *Mxcluster) HasMistNac() bool`

HasMistNac returns a boolean if a field has been set.

### GetModifiedTime

`func (o *Mxcluster) GetModifiedTime() float32`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *Mxcluster) GetModifiedTimeOk() (*float32, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *Mxcluster) SetModifiedTime(v float32)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *Mxcluster) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### GetMxedgeMgmt

`func (o *Mxcluster) GetMxedgeMgmt() MxedgeMgmt`

GetMxedgeMgmt returns the MxedgeMgmt field if non-nil, zero value otherwise.

### GetMxedgeMgmtOk

`func (o *Mxcluster) GetMxedgeMgmtOk() (*MxedgeMgmt, bool)`

GetMxedgeMgmtOk returns a tuple with the MxedgeMgmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMxedgeMgmt

`func (o *Mxcluster) SetMxedgeMgmt(v MxedgeMgmt)`

SetMxedgeMgmt sets MxedgeMgmt field to given value.

### HasMxedgeMgmt

`func (o *Mxcluster) HasMxedgeMgmt() bool`

HasMxedgeMgmt returns a boolean if a field has been set.

### GetName

`func (o *Mxcluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Mxcluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Mxcluster) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Mxcluster) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrgId

`func (o *Mxcluster) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *Mxcluster) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *Mxcluster) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *Mxcluster) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetProxy

`func (o *Mxcluster) GetProxy() Proxy`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *Mxcluster) GetProxyOk() (*Proxy, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *Mxcluster) SetProxy(v Proxy)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *Mxcluster) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### GetRadsec

`func (o *Mxcluster) GetRadsec() MxclusterRadsec`

GetRadsec returns the Radsec field if non-nil, zero value otherwise.

### GetRadsecOk

`func (o *Mxcluster) GetRadsecOk() (*MxclusterRadsec, bool)`

GetRadsecOk returns a tuple with the Radsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadsec

`func (o *Mxcluster) SetRadsec(v MxclusterRadsec)`

SetRadsec sets Radsec field to given value.

### HasRadsec

`func (o *Mxcluster) HasRadsec() bool`

HasRadsec returns a boolean if a field has been set.

### GetRadsecTls

`func (o *Mxcluster) GetRadsecTls() MxclusterRadsecTls`

GetRadsecTls returns the RadsecTls field if non-nil, zero value otherwise.

### GetRadsecTlsOk

`func (o *Mxcluster) GetRadsecTlsOk() (*MxclusterRadsecTls, bool)`

GetRadsecTlsOk returns a tuple with the RadsecTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadsecTls

`func (o *Mxcluster) SetRadsecTls(v MxclusterRadsecTls)`

SetRadsecTls sets RadsecTls field to given value.

### HasRadsecTls

`func (o *Mxcluster) HasRadsecTls() bool`

HasRadsecTls returns a boolean if a field has been set.

### GetSiteId

`func (o *Mxcluster) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *Mxcluster) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *Mxcluster) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *Mxcluster) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetTuntermApSubnets

`func (o *Mxcluster) GetTuntermApSubnets() []string`

GetTuntermApSubnets returns the TuntermApSubnets field if non-nil, zero value otherwise.

### GetTuntermApSubnetsOk

`func (o *Mxcluster) GetTuntermApSubnetsOk() (*[]string, bool)`

GetTuntermApSubnetsOk returns a tuple with the TuntermApSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuntermApSubnets

`func (o *Mxcluster) SetTuntermApSubnets(v []string)`

SetTuntermApSubnets sets TuntermApSubnets field to given value.

### HasTuntermApSubnets

`func (o *Mxcluster) HasTuntermApSubnets() bool`

HasTuntermApSubnets returns a boolean if a field has been set.

### GetTuntermDhcpdConfig

`func (o *Mxcluster) GetTuntermDhcpdConfig() TuntermDhcpdConfig`

GetTuntermDhcpdConfig returns the TuntermDhcpdConfig field if non-nil, zero value otherwise.

### GetTuntermDhcpdConfigOk

`func (o *Mxcluster) GetTuntermDhcpdConfigOk() (*TuntermDhcpdConfig, bool)`

GetTuntermDhcpdConfigOk returns a tuple with the TuntermDhcpdConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuntermDhcpdConfig

`func (o *Mxcluster) SetTuntermDhcpdConfig(v TuntermDhcpdConfig)`

SetTuntermDhcpdConfig sets TuntermDhcpdConfig field to given value.

### HasTuntermDhcpdConfig

`func (o *Mxcluster) HasTuntermDhcpdConfig() bool`

HasTuntermDhcpdConfig returns a boolean if a field has been set.

### GetTuntermExtraRoutes

`func (o *Mxcluster) GetTuntermExtraRoutes() map[string]MxclusterTuntermExtraRoute`

GetTuntermExtraRoutes returns the TuntermExtraRoutes field if non-nil, zero value otherwise.

### GetTuntermExtraRoutesOk

`func (o *Mxcluster) GetTuntermExtraRoutesOk() (*map[string]MxclusterTuntermExtraRoute, bool)`

GetTuntermExtraRoutesOk returns a tuple with the TuntermExtraRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuntermExtraRoutes

`func (o *Mxcluster) SetTuntermExtraRoutes(v map[string]MxclusterTuntermExtraRoute)`

SetTuntermExtraRoutes sets TuntermExtraRoutes field to given value.

### HasTuntermExtraRoutes

`func (o *Mxcluster) HasTuntermExtraRoutes() bool`

HasTuntermExtraRoutes returns a boolean if a field has been set.

### GetTuntermHosts

`func (o *Mxcluster) GetTuntermHosts() []string`

GetTuntermHosts returns the TuntermHosts field if non-nil, zero value otherwise.

### GetTuntermHostsOk

`func (o *Mxcluster) GetTuntermHostsOk() (*[]string, bool)`

GetTuntermHostsOk returns a tuple with the TuntermHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuntermHosts

`func (o *Mxcluster) SetTuntermHosts(v []string)`

SetTuntermHosts sets TuntermHosts field to given value.

### HasTuntermHosts

`func (o *Mxcluster) HasTuntermHosts() bool`

HasTuntermHosts returns a boolean if a field has been set.

### GetTuntermHostsOrder

`func (o *Mxcluster) GetTuntermHostsOrder() []int32`

GetTuntermHostsOrder returns the TuntermHostsOrder field if non-nil, zero value otherwise.

### GetTuntermHostsOrderOk

`func (o *Mxcluster) GetTuntermHostsOrderOk() (*[]int32, bool)`

GetTuntermHostsOrderOk returns a tuple with the TuntermHostsOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuntermHostsOrder

`func (o *Mxcluster) SetTuntermHostsOrder(v []int32)`

SetTuntermHostsOrder sets TuntermHostsOrder field to given value.

### HasTuntermHostsOrder

`func (o *Mxcluster) HasTuntermHostsOrder() bool`

HasTuntermHostsOrder returns a boolean if a field has been set.

### GetTuntermHostsSelection

`func (o *Mxcluster) GetTuntermHostsSelection() MxclusterTuntermHostsSelection`

GetTuntermHostsSelection returns the TuntermHostsSelection field if non-nil, zero value otherwise.

### GetTuntermHostsSelectionOk

`func (o *Mxcluster) GetTuntermHostsSelectionOk() (*MxclusterTuntermHostsSelection, bool)`

GetTuntermHostsSelectionOk returns a tuple with the TuntermHostsSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuntermHostsSelection

`func (o *Mxcluster) SetTuntermHostsSelection(v MxclusterTuntermHostsSelection)`

SetTuntermHostsSelection sets TuntermHostsSelection field to given value.

### HasTuntermHostsSelection

`func (o *Mxcluster) HasTuntermHostsSelection() bool`

HasTuntermHostsSelection returns a boolean if a field has been set.

### GetTuntermMonitoring

`func (o *Mxcluster) GetTuntermMonitoring() [][]TuntermMonitoringItem`

GetTuntermMonitoring returns the TuntermMonitoring field if non-nil, zero value otherwise.

### GetTuntermMonitoringOk

`func (o *Mxcluster) GetTuntermMonitoringOk() (*[][]TuntermMonitoringItem, bool)`

GetTuntermMonitoringOk returns a tuple with the TuntermMonitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuntermMonitoring

`func (o *Mxcluster) SetTuntermMonitoring(v [][]TuntermMonitoringItem)`

SetTuntermMonitoring sets TuntermMonitoring field to given value.

### HasTuntermMonitoring

`func (o *Mxcluster) HasTuntermMonitoring() bool`

HasTuntermMonitoring returns a boolean if a field has been set.

### GetTuntermMonitoringDisabled

`func (o *Mxcluster) GetTuntermMonitoringDisabled() bool`

GetTuntermMonitoringDisabled returns the TuntermMonitoringDisabled field if non-nil, zero value otherwise.

### GetTuntermMonitoringDisabledOk

`func (o *Mxcluster) GetTuntermMonitoringDisabledOk() (*bool, bool)`

GetTuntermMonitoringDisabledOk returns a tuple with the TuntermMonitoringDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuntermMonitoringDisabled

`func (o *Mxcluster) SetTuntermMonitoringDisabled(v bool)`

SetTuntermMonitoringDisabled sets TuntermMonitoringDisabled field to given value.

### HasTuntermMonitoringDisabled

`func (o *Mxcluster) HasTuntermMonitoringDisabled() bool`

HasTuntermMonitoringDisabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


