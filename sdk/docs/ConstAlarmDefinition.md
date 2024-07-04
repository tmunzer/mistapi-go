# ConstAlarmDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Display** | **string** | Description of the alarm type | 
**Example** | Pointer to **map[string]interface{}** |  | [optional] 
**Fields** | **[]string** | List of fields available in an alarm details payload (in REST APIs &amp; Webhooks); e.g. &#x60;aps&#x60;, &#x60;switches&#x60;, &#x60;gateways&#x60;, &#x60;hostnames&#x60;, &#x60;ssids&#x60;, &#x60;bssids&#x60; | 
**Group** | **string** | Group to which the alarm belongs | 
**Key** | **string** | Key name of the alarm type | 
**MarvisSuggestionCategory** | Pointer to **string** | Marvis defined category to which the alarm belongs | [optional] 
**Severity** | **string** | Severity of the alarm | 

## Methods

### NewConstAlarmDefinition

`func NewConstAlarmDefinition(display string, fields []string, group string, key string, severity string, ) *ConstAlarmDefinition`

NewConstAlarmDefinition instantiates a new ConstAlarmDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConstAlarmDefinitionWithDefaults

`func NewConstAlarmDefinitionWithDefaults() *ConstAlarmDefinition`

NewConstAlarmDefinitionWithDefaults instantiates a new ConstAlarmDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplay

`func (o *ConstAlarmDefinition) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *ConstAlarmDefinition) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *ConstAlarmDefinition) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetExample

`func (o *ConstAlarmDefinition) GetExample() map[string]interface{}`

GetExample returns the Example field if non-nil, zero value otherwise.

### GetExampleOk

`func (o *ConstAlarmDefinition) GetExampleOk() (*map[string]interface{}, bool)`

GetExampleOk returns a tuple with the Example field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExample

`func (o *ConstAlarmDefinition) SetExample(v map[string]interface{})`

SetExample sets Example field to given value.

### HasExample

`func (o *ConstAlarmDefinition) HasExample() bool`

HasExample returns a boolean if a field has been set.

### GetFields

`func (o *ConstAlarmDefinition) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ConstAlarmDefinition) GetFieldsOk() (*[]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ConstAlarmDefinition) SetFields(v []string)`

SetFields sets Fields field to given value.


### GetGroup

`func (o *ConstAlarmDefinition) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ConstAlarmDefinition) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ConstAlarmDefinition) SetGroup(v string)`

SetGroup sets Group field to given value.


### GetKey

`func (o *ConstAlarmDefinition) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ConstAlarmDefinition) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ConstAlarmDefinition) SetKey(v string)`

SetKey sets Key field to given value.


### GetMarvisSuggestionCategory

`func (o *ConstAlarmDefinition) GetMarvisSuggestionCategory() string`

GetMarvisSuggestionCategory returns the MarvisSuggestionCategory field if non-nil, zero value otherwise.

### GetMarvisSuggestionCategoryOk

`func (o *ConstAlarmDefinition) GetMarvisSuggestionCategoryOk() (*string, bool)`

GetMarvisSuggestionCategoryOk returns a tuple with the MarvisSuggestionCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarvisSuggestionCategory

`func (o *ConstAlarmDefinition) SetMarvisSuggestionCategory(v string)`

SetMarvisSuggestionCategory sets MarvisSuggestionCategory field to given value.

### HasMarvisSuggestionCategory

`func (o *ConstAlarmDefinition) HasMarvisSuggestionCategory() bool`

HasMarvisSuggestionCategory returns a boolean if a field has been set.

### GetSeverity

`func (o *ConstAlarmDefinition) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ConstAlarmDefinition) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ConstAlarmDefinition) SetSeverity(v string)`

SetSeverity sets Severity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


