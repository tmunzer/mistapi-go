# SiteTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoUpgrade** | Pointer to [**SiteTemplateAutoUpgrade**](SiteTemplateAutoUpgrade.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Vars** | Pointer to **map[string]string** | a dictionary of name-&gt;value, the vars can then be used in Wlans. This can overwrite those from Site Vars | [optional] 

## Methods

### NewSiteTemplate

`func NewSiteTemplate() *SiteTemplate`

NewSiteTemplate instantiates a new SiteTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteTemplateWithDefaults

`func NewSiteTemplateWithDefaults() *SiteTemplate`

NewSiteTemplateWithDefaults instantiates a new SiteTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoUpgrade

`func (o *SiteTemplate) GetAutoUpgrade() SiteTemplateAutoUpgrade`

GetAutoUpgrade returns the AutoUpgrade field if non-nil, zero value otherwise.

### GetAutoUpgradeOk

`func (o *SiteTemplate) GetAutoUpgradeOk() (*SiteTemplateAutoUpgrade, bool)`

GetAutoUpgradeOk returns a tuple with the AutoUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoUpgrade

`func (o *SiteTemplate) SetAutoUpgrade(v SiteTemplateAutoUpgrade)`

SetAutoUpgrade sets AutoUpgrade field to given value.

### HasAutoUpgrade

`func (o *SiteTemplate) HasAutoUpgrade() bool`

HasAutoUpgrade returns a boolean if a field has been set.

### GetName

`func (o *SiteTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SiteTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SiteTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SiteTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVars

`func (o *SiteTemplate) GetVars() map[string]string`

GetVars returns the Vars field if non-nil, zero value otherwise.

### GetVarsOk

`func (o *SiteTemplate) GetVarsOk() (*map[string]string, bool)`

GetVarsOk returns a tuple with the Vars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVars

`func (o *SiteTemplate) SetVars(v map[string]string)`

SetVars sets Vars field to given value.

### HasVars

`func (o *SiteTemplate) HasVars() bool`

HasVars returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


