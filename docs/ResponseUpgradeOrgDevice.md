# ResponseUpgradeOrgDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SiteId** | Pointer to **string** |  | [optional] [readonly] 
**Upgrade** | Pointer to [**UpgradeOrgDeviceUpgrade**](UpgradeOrgDeviceUpgrade.md) |  | [optional] 

## Methods

### NewResponseUpgradeOrgDevice

`func NewResponseUpgradeOrgDevice() *ResponseUpgradeOrgDevice`

NewResponseUpgradeOrgDevice instantiates a new ResponseUpgradeOrgDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseUpgradeOrgDeviceWithDefaults

`func NewResponseUpgradeOrgDeviceWithDefaults() *ResponseUpgradeOrgDevice`

NewResponseUpgradeOrgDeviceWithDefaults instantiates a new ResponseUpgradeOrgDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiteId

`func (o *ResponseUpgradeOrgDevice) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *ResponseUpgradeOrgDevice) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *ResponseUpgradeOrgDevice) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *ResponseUpgradeOrgDevice) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetUpgrade

`func (o *ResponseUpgradeOrgDevice) GetUpgrade() UpgradeOrgDeviceUpgrade`

GetUpgrade returns the Upgrade field if non-nil, zero value otherwise.

### GetUpgradeOk

`func (o *ResponseUpgradeOrgDevice) GetUpgradeOk() (*UpgradeOrgDeviceUpgrade, bool)`

GetUpgradeOk returns a tuple with the Upgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgrade

`func (o *ResponseUpgradeOrgDevice) SetUpgrade(v UpgradeOrgDeviceUpgrade)`

SetUpgrade sets Upgrade field to given value.

### HasUpgrade

`func (o *ResponseUpgradeOrgDevice) HasUpgrade() bool`

HasUpgrade returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


