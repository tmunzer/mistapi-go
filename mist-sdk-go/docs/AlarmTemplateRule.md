# AlarmTemplateRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Delivery** | Pointer to [**Delivery**](Delivery.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewAlarmTemplateRule

`func NewAlarmTemplateRule() *AlarmTemplateRule`

NewAlarmTemplateRule instantiates a new AlarmTemplateRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlarmTemplateRuleWithDefaults

`func NewAlarmTemplateRuleWithDefaults() *AlarmTemplateRule`

NewAlarmTemplateRuleWithDefaults instantiates a new AlarmTemplateRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelivery

`func (o *AlarmTemplateRule) GetDelivery() Delivery`

GetDelivery returns the Delivery field if non-nil, zero value otherwise.

### GetDeliveryOk

`func (o *AlarmTemplateRule) GetDeliveryOk() (*Delivery, bool)`

GetDeliveryOk returns a tuple with the Delivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivery

`func (o *AlarmTemplateRule) SetDelivery(v Delivery)`

SetDelivery sets Delivery field to given value.

### HasDelivery

`func (o *AlarmTemplateRule) HasDelivery() bool`

HasDelivery returns a boolean if a field has been set.

### GetEnabled

`func (o *AlarmTemplateRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AlarmTemplateRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AlarmTemplateRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AlarmTemplateRule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


