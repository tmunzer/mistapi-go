# SwitchPortUsageDynamicRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Equals** | Pointer to **string** |  | [optional] 
**EqualsAny** | Pointer to **[]string** | use &#x60;equals_any&#x60; to match any item in a list | [optional] 
**Expression** | Pointer to **string** | \&quot;[0:3]\&quot;:\&quot;abcdef\&quot; -&gt; \&quot;abc\&quot; \&quot;split(.)[1]\&quot;: \&quot;a.b.c\&quot; -&gt; \&quot;b\&quot; \&quot;split(-)[1][0:3]: \&quot;a1234-b5678-c90\&quot; -&gt; \&quot;b56\&quot; | [optional] 
**Src** | [**SwitchPortUsageDynamicRuleSrc**](SwitchPortUsageDynamicRuleSrc.md) |  | 
**Usage** | Pointer to **string** | &#x60;port_usage&#x60; name | [optional] 

## Methods

### NewSwitchPortUsageDynamicRule

`func NewSwitchPortUsageDynamicRule(src SwitchPortUsageDynamicRuleSrc, ) *SwitchPortUsageDynamicRule`

NewSwitchPortUsageDynamicRule instantiates a new SwitchPortUsageDynamicRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwitchPortUsageDynamicRuleWithDefaults

`func NewSwitchPortUsageDynamicRuleWithDefaults() *SwitchPortUsageDynamicRule`

NewSwitchPortUsageDynamicRuleWithDefaults instantiates a new SwitchPortUsageDynamicRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEquals

`func (o *SwitchPortUsageDynamicRule) GetEquals() string`

GetEquals returns the Equals field if non-nil, zero value otherwise.

### GetEqualsOk

`func (o *SwitchPortUsageDynamicRule) GetEqualsOk() (*string, bool)`

GetEqualsOk returns a tuple with the Equals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquals

`func (o *SwitchPortUsageDynamicRule) SetEquals(v string)`

SetEquals sets Equals field to given value.

### HasEquals

`func (o *SwitchPortUsageDynamicRule) HasEquals() bool`

HasEquals returns a boolean if a field has been set.

### GetEqualsAny

`func (o *SwitchPortUsageDynamicRule) GetEqualsAny() []string`

GetEqualsAny returns the EqualsAny field if non-nil, zero value otherwise.

### GetEqualsAnyOk

`func (o *SwitchPortUsageDynamicRule) GetEqualsAnyOk() (*[]string, bool)`

GetEqualsAnyOk returns a tuple with the EqualsAny field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEqualsAny

`func (o *SwitchPortUsageDynamicRule) SetEqualsAny(v []string)`

SetEqualsAny sets EqualsAny field to given value.

### HasEqualsAny

`func (o *SwitchPortUsageDynamicRule) HasEqualsAny() bool`

HasEqualsAny returns a boolean if a field has been set.

### GetExpression

`func (o *SwitchPortUsageDynamicRule) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *SwitchPortUsageDynamicRule) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *SwitchPortUsageDynamicRule) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *SwitchPortUsageDynamicRule) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### GetSrc

`func (o *SwitchPortUsageDynamicRule) GetSrc() SwitchPortUsageDynamicRuleSrc`

GetSrc returns the Src field if non-nil, zero value otherwise.

### GetSrcOk

`func (o *SwitchPortUsageDynamicRule) GetSrcOk() (*SwitchPortUsageDynamicRuleSrc, bool)`

GetSrcOk returns a tuple with the Src field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrc

`func (o *SwitchPortUsageDynamicRule) SetSrc(v SwitchPortUsageDynamicRuleSrc)`

SetSrc sets Src field to given value.


### GetUsage

`func (o *SwitchPortUsageDynamicRule) GetUsage() string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *SwitchPortUsageDynamicRule) GetUsageOk() (*string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *SwitchPortUsageDynamicRule) SetUsage(v string)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *SwitchPortUsageDynamicRule) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


