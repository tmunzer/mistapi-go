# NacRuleMatching

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthType** | Pointer to [**NacRuleMatchingAuthType**](NacRuleMatchingAuthType.md) |  | [optional] 
**Nactags** | Pointer to **[]string** |  | [optional] 
**PortTypes** | Pointer to [**[]NacRuleMatchingPortType**](NacRuleMatchingPortType.md) |  | [optional] 
**SiteIds** | Pointer to **[]string** | list of site ids to match | [optional] 
**SitegroupIds** | Pointer to **[]string** | list of sitegroup ids to match | [optional] 
**Vendor** | Pointer to **[]string** | list of vendors to match | [optional] 

## Methods

### NewNacRuleMatching

`func NewNacRuleMatching() *NacRuleMatching`

NewNacRuleMatching instantiates a new NacRuleMatching object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNacRuleMatchingWithDefaults

`func NewNacRuleMatchingWithDefaults() *NacRuleMatching`

NewNacRuleMatchingWithDefaults instantiates a new NacRuleMatching object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthType

`func (o *NacRuleMatching) GetAuthType() NacRuleMatchingAuthType`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *NacRuleMatching) GetAuthTypeOk() (*NacRuleMatchingAuthType, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *NacRuleMatching) SetAuthType(v NacRuleMatchingAuthType)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *NacRuleMatching) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetNactags

`func (o *NacRuleMatching) GetNactags() []string`

GetNactags returns the Nactags field if non-nil, zero value otherwise.

### GetNactagsOk

`func (o *NacRuleMatching) GetNactagsOk() (*[]string, bool)`

GetNactagsOk returns a tuple with the Nactags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNactags

`func (o *NacRuleMatching) SetNactags(v []string)`

SetNactags sets Nactags field to given value.

### HasNactags

`func (o *NacRuleMatching) HasNactags() bool`

HasNactags returns a boolean if a field has been set.

### GetPortTypes

`func (o *NacRuleMatching) GetPortTypes() []NacRuleMatchingPortType`

GetPortTypes returns the PortTypes field if non-nil, zero value otherwise.

### GetPortTypesOk

`func (o *NacRuleMatching) GetPortTypesOk() (*[]NacRuleMatchingPortType, bool)`

GetPortTypesOk returns a tuple with the PortTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortTypes

`func (o *NacRuleMatching) SetPortTypes(v []NacRuleMatchingPortType)`

SetPortTypes sets PortTypes field to given value.

### HasPortTypes

`func (o *NacRuleMatching) HasPortTypes() bool`

HasPortTypes returns a boolean if a field has been set.

### GetSiteIds

`func (o *NacRuleMatching) GetSiteIds() []string`

GetSiteIds returns the SiteIds field if non-nil, zero value otherwise.

### GetSiteIdsOk

`func (o *NacRuleMatching) GetSiteIdsOk() (*[]string, bool)`

GetSiteIdsOk returns a tuple with the SiteIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteIds

`func (o *NacRuleMatching) SetSiteIds(v []string)`

SetSiteIds sets SiteIds field to given value.

### HasSiteIds

`func (o *NacRuleMatching) HasSiteIds() bool`

HasSiteIds returns a boolean if a field has been set.

### GetSitegroupIds

`func (o *NacRuleMatching) GetSitegroupIds() []string`

GetSitegroupIds returns the SitegroupIds field if non-nil, zero value otherwise.

### GetSitegroupIdsOk

`func (o *NacRuleMatching) GetSitegroupIdsOk() (*[]string, bool)`

GetSitegroupIdsOk returns a tuple with the SitegroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitegroupIds

`func (o *NacRuleMatching) SetSitegroupIds(v []string)`

SetSitegroupIds sets SitegroupIds field to given value.

### HasSitegroupIds

`func (o *NacRuleMatching) HasSitegroupIds() bool`

HasSitegroupIds returns a boolean if a field has been set.

### GetVendor

`func (o *NacRuleMatching) GetVendor() []string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *NacRuleMatching) GetVendorOk() (*[]string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *NacRuleMatching) SetVendor(v []string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *NacRuleMatching) HasVendor() bool`

HasVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


