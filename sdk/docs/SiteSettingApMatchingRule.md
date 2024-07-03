# SiteSettingApMatchingRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MatchModel** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PortConfig** | Pointer to [**map[string]ApPortConfig**](ApPortConfig.md) | Property key is the interface(s) (e.g. \&quot;eth1,eth2\&quot;) | [optional] 

## Methods

### NewSiteSettingApMatchingRule

`func NewSiteSettingApMatchingRule() *SiteSettingApMatchingRule`

NewSiteSettingApMatchingRule instantiates a new SiteSettingApMatchingRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteSettingApMatchingRuleWithDefaults

`func NewSiteSettingApMatchingRuleWithDefaults() *SiteSettingApMatchingRule`

NewSiteSettingApMatchingRuleWithDefaults instantiates a new SiteSettingApMatchingRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatchModel

`func (o *SiteSettingApMatchingRule) GetMatchModel() string`

GetMatchModel returns the MatchModel field if non-nil, zero value otherwise.

### GetMatchModelOk

`func (o *SiteSettingApMatchingRule) GetMatchModelOk() (*string, bool)`

GetMatchModelOk returns a tuple with the MatchModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchModel

`func (o *SiteSettingApMatchingRule) SetMatchModel(v string)`

SetMatchModel sets MatchModel field to given value.

### HasMatchModel

`func (o *SiteSettingApMatchingRule) HasMatchModel() bool`

HasMatchModel returns a boolean if a field has been set.

### GetName

`func (o *SiteSettingApMatchingRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SiteSettingApMatchingRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SiteSettingApMatchingRule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SiteSettingApMatchingRule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPortConfig

`func (o *SiteSettingApMatchingRule) GetPortConfig() map[string]ApPortConfig`

GetPortConfig returns the PortConfig field if non-nil, zero value otherwise.

### GetPortConfigOk

`func (o *SiteSettingApMatchingRule) GetPortConfigOk() (*map[string]ApPortConfig, bool)`

GetPortConfigOk returns a tuple with the PortConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortConfig

`func (o *SiteSettingApMatchingRule) SetPortConfig(v map[string]ApPortConfig)`

SetPortConfig sets PortConfig field to given value.

### HasPortConfig

`func (o *SiteSettingApMatchingRule) HasPortConfig() bool`

HasPortConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


