# Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | Pointer to **[]string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;custom&#x60;, ip subnets | [optional] 
**AppCategories** | Pointer to **[]string** | when &#x60;type&#x60;&#x3D;&#x3D;&#x60;app_categories&#x60; list of application categories are available through /api/v1/const/app_categories | [optional] 
**AppSubcategories** | Pointer to **[]string** | when &#x60;type&#x60;&#x3D;&#x3D;&#x60;app_categories&#x60; list of application categories are available through /api/v1/const/app_subcategories | [optional] 
**Apps** | Pointer to **[]string** | when &#x60;type&#x60;&#x3D;&#x3D;&#x60;apps&#x60; list of applications are available through:   - /api/v1/const/applications,   - /api/v1/const/gateway_applications   - /insight/top_app_by-bytes?wired&#x3D;true | [optional] 
**CreatedTime** | Pointer to **int32** |  | [optional] [readonly] 
**Description** | Pointer to **string** |  | [optional] 
**Dscp** | Pointer to **int32** | when &#x60;traffic_type&#x60;&#x3D;&#x3D;&#x60;custom&#x60; | [optional] 
**FailoverPolicy** | Pointer to [**ServiceFailoverPolicy**](ServiceFailoverPolicy.md) |  | [optional] [default to SERVICEFAILOVERPOLICY_REVERTABLE]
**Hostnames** | Pointer to **[]string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;custom&#x60;, web filtering | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**MaxJitter** | Pointer to **int32** | when &#x60;traffic_type&#x60;&#x3D;&#x3D;&#x60;custom&#x60;, for uplink selection | [optional] 
**MaxLatency** | Pointer to **int32** | when &#x60;traffic_type&#x60;&#x3D;&#x3D;&#x60;custom&#x60;, for uplink selection | [optional] 
**MaxLoss** | Pointer to **int32** | when &#x60;traffic_type&#x60;&#x3D;&#x3D;&#x60;custom&#x60;, for uplink selection | [optional] 
**ModifiedTime** | Pointer to **int32** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**SleEnabled** | Pointer to **bool** | whether to enable measure SLE | [optional] [default to false]
**Specs** | Pointer to [**[]ServiceSpec**](ServiceSpec.md) |  | [optional] 
**SsrRelaxedTcpStateEnforcement** | Pointer to **bool** |  | [optional] [default to false]
**TrafficClass** | Pointer to [**ServiceTrafficClass**](ServiceTrafficClass.md) |  | [optional] [default to SERVICETRAFFICCLASS_BEST_EFFORT]
**TrafficType** | Pointer to **string** | values from &#x60;/api/v1/consts/traffic_types&#x60; * when &#x60;type&#x60;&#x3D;&#x3D;&#x60;apps&#x60;, we&#39;ll choose traffic_type automatically * when &#x60;type&#x60;&#x3D;&#x3D;&#x60;addresses&#x60; or &#x60;type&#x60;&#x3D;&#x3D;&#x60;hostnames&#x60;, you can provide your own settings (optional) | [optional] [default to "data_best_effort"]
**Type** | Pointer to [**ServiceType**](ServiceType.md) |  | [optional] [default to SERVICETYPE_CUSTOM]
**Urls** | Pointer to **[]string** | when &#x60;type&#x60;&#x3D;&#x3D;&#x60;urls no need for spec as URL can encode the ports being used&#x60; | [optional] 

## Methods

### NewService

`func NewService() *Service`

NewService instantiates a new Service object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceWithDefaults

`func NewServiceWithDefaults() *Service`

NewServiceWithDefaults instantiates a new Service object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *Service) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *Service) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *Service) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *Service) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetAppCategories

`func (o *Service) GetAppCategories() []string`

GetAppCategories returns the AppCategories field if non-nil, zero value otherwise.

### GetAppCategoriesOk

`func (o *Service) GetAppCategoriesOk() (*[]string, bool)`

GetAppCategoriesOk returns a tuple with the AppCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppCategories

`func (o *Service) SetAppCategories(v []string)`

SetAppCategories sets AppCategories field to given value.

### HasAppCategories

`func (o *Service) HasAppCategories() bool`

HasAppCategories returns a boolean if a field has been set.

### GetAppSubcategories

`func (o *Service) GetAppSubcategories() []string`

GetAppSubcategories returns the AppSubcategories field if non-nil, zero value otherwise.

### GetAppSubcategoriesOk

`func (o *Service) GetAppSubcategoriesOk() (*[]string, bool)`

GetAppSubcategoriesOk returns a tuple with the AppSubcategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppSubcategories

`func (o *Service) SetAppSubcategories(v []string)`

SetAppSubcategories sets AppSubcategories field to given value.

### HasAppSubcategories

`func (o *Service) HasAppSubcategories() bool`

HasAppSubcategories returns a boolean if a field has been set.

### GetApps

`func (o *Service) GetApps() []string`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *Service) GetAppsOk() (*[]string, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *Service) SetApps(v []string)`

SetApps sets Apps field to given value.

### HasApps

`func (o *Service) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetCreatedTime

`func (o *Service) GetCreatedTime() int32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Service) GetCreatedTimeOk() (*int32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Service) SetCreatedTime(v int32)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *Service) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDescription

`func (o *Service) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Service) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Service) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Service) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDscp

`func (o *Service) GetDscp() int32`

GetDscp returns the Dscp field if non-nil, zero value otherwise.

### GetDscpOk

`func (o *Service) GetDscpOk() (*int32, bool)`

GetDscpOk returns a tuple with the Dscp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDscp

`func (o *Service) SetDscp(v int32)`

SetDscp sets Dscp field to given value.

### HasDscp

`func (o *Service) HasDscp() bool`

HasDscp returns a boolean if a field has been set.

### GetFailoverPolicy

`func (o *Service) GetFailoverPolicy() ServiceFailoverPolicy`

GetFailoverPolicy returns the FailoverPolicy field if non-nil, zero value otherwise.

### GetFailoverPolicyOk

`func (o *Service) GetFailoverPolicyOk() (*ServiceFailoverPolicy, bool)`

GetFailoverPolicyOk returns a tuple with the FailoverPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverPolicy

`func (o *Service) SetFailoverPolicy(v ServiceFailoverPolicy)`

SetFailoverPolicy sets FailoverPolicy field to given value.

### HasFailoverPolicy

`func (o *Service) HasFailoverPolicy() bool`

HasFailoverPolicy returns a boolean if a field has been set.

### GetHostnames

`func (o *Service) GetHostnames() []string`

GetHostnames returns the Hostnames field if non-nil, zero value otherwise.

### GetHostnamesOk

`func (o *Service) GetHostnamesOk() (*[]string, bool)`

GetHostnamesOk returns a tuple with the Hostnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnames

`func (o *Service) SetHostnames(v []string)`

SetHostnames sets Hostnames field to given value.

### HasHostnames

`func (o *Service) HasHostnames() bool`

HasHostnames returns a boolean if a field has been set.

### GetId

`func (o *Service) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Service) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Service) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Service) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMaxJitter

`func (o *Service) GetMaxJitter() int32`

GetMaxJitter returns the MaxJitter field if non-nil, zero value otherwise.

### GetMaxJitterOk

`func (o *Service) GetMaxJitterOk() (*int32, bool)`

GetMaxJitterOk returns a tuple with the MaxJitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxJitter

`func (o *Service) SetMaxJitter(v int32)`

SetMaxJitter sets MaxJitter field to given value.

### HasMaxJitter

`func (o *Service) HasMaxJitter() bool`

HasMaxJitter returns a boolean if a field has been set.

### GetMaxLatency

`func (o *Service) GetMaxLatency() int32`

GetMaxLatency returns the MaxLatency field if non-nil, zero value otherwise.

### GetMaxLatencyOk

`func (o *Service) GetMaxLatencyOk() (*int32, bool)`

GetMaxLatencyOk returns a tuple with the MaxLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLatency

`func (o *Service) SetMaxLatency(v int32)`

SetMaxLatency sets MaxLatency field to given value.

### HasMaxLatency

`func (o *Service) HasMaxLatency() bool`

HasMaxLatency returns a boolean if a field has been set.

### GetMaxLoss

`func (o *Service) GetMaxLoss() int32`

GetMaxLoss returns the MaxLoss field if non-nil, zero value otherwise.

### GetMaxLossOk

`func (o *Service) GetMaxLossOk() (*int32, bool)`

GetMaxLossOk returns a tuple with the MaxLoss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLoss

`func (o *Service) SetMaxLoss(v int32)`

SetMaxLoss sets MaxLoss field to given value.

### HasMaxLoss

`func (o *Service) HasMaxLoss() bool`

HasMaxLoss returns a boolean if a field has been set.

### GetModifiedTime

`func (o *Service) GetModifiedTime() int32`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *Service) GetModifiedTimeOk() (*int32, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *Service) SetModifiedTime(v int32)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *Service) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### GetName

`func (o *Service) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Service) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Service) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Service) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrgId

`func (o *Service) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *Service) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *Service) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *Service) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetSleEnabled

`func (o *Service) GetSleEnabled() bool`

GetSleEnabled returns the SleEnabled field if non-nil, zero value otherwise.

### GetSleEnabledOk

`func (o *Service) GetSleEnabledOk() (*bool, bool)`

GetSleEnabledOk returns a tuple with the SleEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleEnabled

`func (o *Service) SetSleEnabled(v bool)`

SetSleEnabled sets SleEnabled field to given value.

### HasSleEnabled

`func (o *Service) HasSleEnabled() bool`

HasSleEnabled returns a boolean if a field has been set.

### GetSpecs

`func (o *Service) GetSpecs() []ServiceSpec`

GetSpecs returns the Specs field if non-nil, zero value otherwise.

### GetSpecsOk

`func (o *Service) GetSpecsOk() (*[]ServiceSpec, bool)`

GetSpecsOk returns a tuple with the Specs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecs

`func (o *Service) SetSpecs(v []ServiceSpec)`

SetSpecs sets Specs field to given value.

### HasSpecs

`func (o *Service) HasSpecs() bool`

HasSpecs returns a boolean if a field has been set.

### GetSsrRelaxedTcpStateEnforcement

`func (o *Service) GetSsrRelaxedTcpStateEnforcement() bool`

GetSsrRelaxedTcpStateEnforcement returns the SsrRelaxedTcpStateEnforcement field if non-nil, zero value otherwise.

### GetSsrRelaxedTcpStateEnforcementOk

`func (o *Service) GetSsrRelaxedTcpStateEnforcementOk() (*bool, bool)`

GetSsrRelaxedTcpStateEnforcementOk returns a tuple with the SsrRelaxedTcpStateEnforcement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsrRelaxedTcpStateEnforcement

`func (o *Service) SetSsrRelaxedTcpStateEnforcement(v bool)`

SetSsrRelaxedTcpStateEnforcement sets SsrRelaxedTcpStateEnforcement field to given value.

### HasSsrRelaxedTcpStateEnforcement

`func (o *Service) HasSsrRelaxedTcpStateEnforcement() bool`

HasSsrRelaxedTcpStateEnforcement returns a boolean if a field has been set.

### GetTrafficClass

`func (o *Service) GetTrafficClass() ServiceTrafficClass`

GetTrafficClass returns the TrafficClass field if non-nil, zero value otherwise.

### GetTrafficClassOk

`func (o *Service) GetTrafficClassOk() (*ServiceTrafficClass, bool)`

GetTrafficClassOk returns a tuple with the TrafficClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficClass

`func (o *Service) SetTrafficClass(v ServiceTrafficClass)`

SetTrafficClass sets TrafficClass field to given value.

### HasTrafficClass

`func (o *Service) HasTrafficClass() bool`

HasTrafficClass returns a boolean if a field has been set.

### GetTrafficType

`func (o *Service) GetTrafficType() string`

GetTrafficType returns the TrafficType field if non-nil, zero value otherwise.

### GetTrafficTypeOk

`func (o *Service) GetTrafficTypeOk() (*string, bool)`

GetTrafficTypeOk returns a tuple with the TrafficType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficType

`func (o *Service) SetTrafficType(v string)`

SetTrafficType sets TrafficType field to given value.

### HasTrafficType

`func (o *Service) HasTrafficType() bool`

HasTrafficType returns a boolean if a field has been set.

### GetType

`func (o *Service) GetType() ServiceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Service) GetTypeOk() (*ServiceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Service) SetType(v ServiceType)`

SetType sets Type field to given value.

### HasType

`func (o *Service) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrls

`func (o *Service) GetUrls() []string`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *Service) GetUrlsOk() (*[]string, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *Service) SetUrls(v []string)`

SetUrls sets Urls field to given value.

### HasUrls

`func (o *Service) HasUrls() bool`

HasUrls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


