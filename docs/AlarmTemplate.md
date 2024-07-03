# AlarmTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | Pointer to **float32** |  | [optional] [readonly] 
**Delivery** | [**Delivery**](Delivery.md) |  | 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**ModifiedTime** | Pointer to **float32** |  | [optional] [readonly] 
**Name** | Pointer to **string** | Some string to name the alarm template | [optional] 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**Rules** | [**map[string]AlarmTemplateRule**](AlarmTemplateRule.md) | Alarm Rules object to configure the individual alarm keys/types. Property key is the alarm name. | 

## Methods

### NewAlarmTemplate

`func NewAlarmTemplate(delivery Delivery, rules map[string]AlarmTemplateRule, ) *AlarmTemplate`

NewAlarmTemplate instantiates a new AlarmTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlarmTemplateWithDefaults

`func NewAlarmTemplateWithDefaults() *AlarmTemplate`

NewAlarmTemplateWithDefaults instantiates a new AlarmTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTime

`func (o *AlarmTemplate) GetCreatedTime() float32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *AlarmTemplate) GetCreatedTimeOk() (*float32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *AlarmTemplate) SetCreatedTime(v float32)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *AlarmTemplate) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDelivery

`func (o *AlarmTemplate) GetDelivery() Delivery`

GetDelivery returns the Delivery field if non-nil, zero value otherwise.

### GetDeliveryOk

`func (o *AlarmTemplate) GetDeliveryOk() (*Delivery, bool)`

GetDeliveryOk returns a tuple with the Delivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivery

`func (o *AlarmTemplate) SetDelivery(v Delivery)`

SetDelivery sets Delivery field to given value.


### GetId

`func (o *AlarmTemplate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AlarmTemplate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AlarmTemplate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AlarmTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifiedTime

`func (o *AlarmTemplate) GetModifiedTime() float32`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *AlarmTemplate) GetModifiedTimeOk() (*float32, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *AlarmTemplate) SetModifiedTime(v float32)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *AlarmTemplate) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### GetName

`func (o *AlarmTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlarmTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlarmTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AlarmTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrgId

`func (o *AlarmTemplate) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *AlarmTemplate) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *AlarmTemplate) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *AlarmTemplate) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetRules

`func (o *AlarmTemplate) GetRules() map[string]AlarmTemplateRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *AlarmTemplate) GetRulesOk() (*map[string]AlarmTemplateRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *AlarmTemplate) SetRules(v map[string]AlarmTemplateRule)`

SetRules sets Rules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


