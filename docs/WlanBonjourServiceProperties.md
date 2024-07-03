# WlanBonjourServiceProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisableLocal** | Pointer to **bool** | whether to prevent wireless clients to discover bonjour devices on the same WLAN | [optional] [default to false]
**RadiusGroups** | Pointer to **[]string** | optional, if the service is further restricted for certain RADIUS groups | [optional] 
**Scope** | Pointer to [**WlanBonjourServicePropertiesScope**](WlanBonjourServicePropertiesScope.md) |  | [optional] [default to WLANBONJOURSERVICEPROPERTIESSCOPE_SAME_SITE]

## Methods

### NewWlanBonjourServiceProperties

`func NewWlanBonjourServiceProperties() *WlanBonjourServiceProperties`

NewWlanBonjourServiceProperties instantiates a new WlanBonjourServiceProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWlanBonjourServicePropertiesWithDefaults

`func NewWlanBonjourServicePropertiesWithDefaults() *WlanBonjourServiceProperties`

NewWlanBonjourServicePropertiesWithDefaults instantiates a new WlanBonjourServiceProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisableLocal

`func (o *WlanBonjourServiceProperties) GetDisableLocal() bool`

GetDisableLocal returns the DisableLocal field if non-nil, zero value otherwise.

### GetDisableLocalOk

`func (o *WlanBonjourServiceProperties) GetDisableLocalOk() (*bool, bool)`

GetDisableLocalOk returns a tuple with the DisableLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableLocal

`func (o *WlanBonjourServiceProperties) SetDisableLocal(v bool)`

SetDisableLocal sets DisableLocal field to given value.

### HasDisableLocal

`func (o *WlanBonjourServiceProperties) HasDisableLocal() bool`

HasDisableLocal returns a boolean if a field has been set.

### GetRadiusGroups

`func (o *WlanBonjourServiceProperties) GetRadiusGroups() []string`

GetRadiusGroups returns the RadiusGroups field if non-nil, zero value otherwise.

### GetRadiusGroupsOk

`func (o *WlanBonjourServiceProperties) GetRadiusGroupsOk() (*[]string, bool)`

GetRadiusGroupsOk returns a tuple with the RadiusGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusGroups

`func (o *WlanBonjourServiceProperties) SetRadiusGroups(v []string)`

SetRadiusGroups sets RadiusGroups field to given value.

### HasRadiusGroups

`func (o *WlanBonjourServiceProperties) HasRadiusGroups() bool`

HasRadiusGroups returns a boolean if a field has been set.

### GetScope

`func (o *WlanBonjourServiceProperties) GetScope() WlanBonjourServicePropertiesScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *WlanBonjourServiceProperties) GetScopeOk() (*WlanBonjourServicePropertiesScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *WlanBonjourServiceProperties) SetScope(v WlanBonjourServicePropertiesScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *WlanBonjourServiceProperties) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


