# ServicePolicyEwfRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertOnly** | Pointer to **bool** |  | [optional] 
**BlockMessage** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] [default to false]
**Profile** | Pointer to [**ServicePolicyEwfRuleProfile**](ServicePolicyEwfRuleProfile.md) |  | [optional] [default to SERVICEPOLICYEWFRULEPROFILE_STRICT]

## Methods

### NewServicePolicyEwfRule

`func NewServicePolicyEwfRule() *ServicePolicyEwfRule`

NewServicePolicyEwfRule instantiates a new ServicePolicyEwfRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicePolicyEwfRuleWithDefaults

`func NewServicePolicyEwfRuleWithDefaults() *ServicePolicyEwfRule`

NewServicePolicyEwfRuleWithDefaults instantiates a new ServicePolicyEwfRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertOnly

`func (o *ServicePolicyEwfRule) GetAlertOnly() bool`

GetAlertOnly returns the AlertOnly field if non-nil, zero value otherwise.

### GetAlertOnlyOk

`func (o *ServicePolicyEwfRule) GetAlertOnlyOk() (*bool, bool)`

GetAlertOnlyOk returns a tuple with the AlertOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnly

`func (o *ServicePolicyEwfRule) SetAlertOnly(v bool)`

SetAlertOnly sets AlertOnly field to given value.

### HasAlertOnly

`func (o *ServicePolicyEwfRule) HasAlertOnly() bool`

HasAlertOnly returns a boolean if a field has been set.

### GetBlockMessage

`func (o *ServicePolicyEwfRule) GetBlockMessage() string`

GetBlockMessage returns the BlockMessage field if non-nil, zero value otherwise.

### GetBlockMessageOk

`func (o *ServicePolicyEwfRule) GetBlockMessageOk() (*string, bool)`

GetBlockMessageOk returns a tuple with the BlockMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockMessage

`func (o *ServicePolicyEwfRule) SetBlockMessage(v string)`

SetBlockMessage sets BlockMessage field to given value.

### HasBlockMessage

`func (o *ServicePolicyEwfRule) HasBlockMessage() bool`

HasBlockMessage returns a boolean if a field has been set.

### GetEnabled

`func (o *ServicePolicyEwfRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ServicePolicyEwfRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ServicePolicyEwfRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ServicePolicyEwfRule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetProfile

`func (o *ServicePolicyEwfRule) GetProfile() ServicePolicyEwfRuleProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ServicePolicyEwfRule) GetProfileOk() (*ServicePolicyEwfRuleProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ServicePolicyEwfRule) SetProfile(v ServicePolicyEwfRuleProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *ServicePolicyEwfRule) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


