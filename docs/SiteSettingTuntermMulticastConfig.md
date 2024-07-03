# SiteSettingTuntermMulticastConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mdns** | Pointer to [**SiteSettingTuntermMulticastConfigMdns**](SiteSettingTuntermMulticastConfigMdns.md) |  | [optional] 
**MulticastAll** | Pointer to **bool** |  | [optional] [default to false]
**Ssdp** | Pointer to [**SiteSettingTuntermMulticastConfigSsdp**](SiteSettingTuntermMulticastConfigSsdp.md) |  | [optional] 

## Methods

### NewSiteSettingTuntermMulticastConfig

`func NewSiteSettingTuntermMulticastConfig() *SiteSettingTuntermMulticastConfig`

NewSiteSettingTuntermMulticastConfig instantiates a new SiteSettingTuntermMulticastConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteSettingTuntermMulticastConfigWithDefaults

`func NewSiteSettingTuntermMulticastConfigWithDefaults() *SiteSettingTuntermMulticastConfig`

NewSiteSettingTuntermMulticastConfigWithDefaults instantiates a new SiteSettingTuntermMulticastConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMdns

`func (o *SiteSettingTuntermMulticastConfig) GetMdns() SiteSettingTuntermMulticastConfigMdns`

GetMdns returns the Mdns field if non-nil, zero value otherwise.

### GetMdnsOk

`func (o *SiteSettingTuntermMulticastConfig) GetMdnsOk() (*SiteSettingTuntermMulticastConfigMdns, bool)`

GetMdnsOk returns a tuple with the Mdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdns

`func (o *SiteSettingTuntermMulticastConfig) SetMdns(v SiteSettingTuntermMulticastConfigMdns)`

SetMdns sets Mdns field to given value.

### HasMdns

`func (o *SiteSettingTuntermMulticastConfig) HasMdns() bool`

HasMdns returns a boolean if a field has been set.

### GetMulticastAll

`func (o *SiteSettingTuntermMulticastConfig) GetMulticastAll() bool`

GetMulticastAll returns the MulticastAll field if non-nil, zero value otherwise.

### GetMulticastAllOk

`func (o *SiteSettingTuntermMulticastConfig) GetMulticastAllOk() (*bool, bool)`

GetMulticastAllOk returns a tuple with the MulticastAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMulticastAll

`func (o *SiteSettingTuntermMulticastConfig) SetMulticastAll(v bool)`

SetMulticastAll sets MulticastAll field to given value.

### HasMulticastAll

`func (o *SiteSettingTuntermMulticastConfig) HasMulticastAll() bool`

HasMulticastAll returns a boolean if a field has been set.

### GetSsdp

`func (o *SiteSettingTuntermMulticastConfig) GetSsdp() SiteSettingTuntermMulticastConfigSsdp`

GetSsdp returns the Ssdp field if non-nil, zero value otherwise.

### GetSsdpOk

`func (o *SiteSettingTuntermMulticastConfig) GetSsdpOk() (*SiteSettingTuntermMulticastConfigSsdp, bool)`

GetSsdpOk returns a tuple with the Ssdp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsdp

`func (o *SiteSettingTuntermMulticastConfig) SetSsdp(v SiteSettingTuntermMulticastConfigSsdp)`

SetSsdp sets Ssdp field to given value.

### HasSsdp

`func (o *SiteSettingTuntermMulticastConfig) HasSsdp() bool`

HasSsdp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


