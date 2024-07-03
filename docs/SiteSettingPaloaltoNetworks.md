# SiteSettingPaloaltoNetworks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gateways** | Pointer to [**[]SiteSettingPaloaltoNetworkGateway**](SiteSettingPaloaltoNetworkGateway.md) |  | [optional] 
**SendMistNacUserInfo** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewSiteSettingPaloaltoNetworks

`func NewSiteSettingPaloaltoNetworks() *SiteSettingPaloaltoNetworks`

NewSiteSettingPaloaltoNetworks instantiates a new SiteSettingPaloaltoNetworks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteSettingPaloaltoNetworksWithDefaults

`func NewSiteSettingPaloaltoNetworksWithDefaults() *SiteSettingPaloaltoNetworks`

NewSiteSettingPaloaltoNetworksWithDefaults instantiates a new SiteSettingPaloaltoNetworks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGateways

`func (o *SiteSettingPaloaltoNetworks) GetGateways() []SiteSettingPaloaltoNetworkGateway`

GetGateways returns the Gateways field if non-nil, zero value otherwise.

### GetGatewaysOk

`func (o *SiteSettingPaloaltoNetworks) GetGatewaysOk() (*[]SiteSettingPaloaltoNetworkGateway, bool)`

GetGatewaysOk returns a tuple with the Gateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateways

`func (o *SiteSettingPaloaltoNetworks) SetGateways(v []SiteSettingPaloaltoNetworkGateway)`

SetGateways sets Gateways field to given value.

### HasGateways

`func (o *SiteSettingPaloaltoNetworks) HasGateways() bool`

HasGateways returns a boolean if a field has been set.

### GetSendMistNacUserInfo

`func (o *SiteSettingPaloaltoNetworks) GetSendMistNacUserInfo() bool`

GetSendMistNacUserInfo returns the SendMistNacUserInfo field if non-nil, zero value otherwise.

### GetSendMistNacUserInfoOk

`func (o *SiteSettingPaloaltoNetworks) GetSendMistNacUserInfoOk() (*bool, bool)`

GetSendMistNacUserInfoOk returns a tuple with the SendMistNacUserInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendMistNacUserInfo

`func (o *SiteSettingPaloaltoNetworks) SetSendMistNacUserInfo(v bool)`

SetSendMistNacUserInfo sets SendMistNacUserInfo field to given value.

### HasSendMistNacUserInfo

`func (o *SiteSettingPaloaltoNetworks) HasSendMistNacUserInfo() bool`

HasSendMistNacUserInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


