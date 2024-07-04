# OrgAutoRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expression** | Pointer to **NullableString** | \&quot;[0:3]\&quot;            // \&quot;abcdef\&quot; -&gt; \&quot;abc\&quot; \&quot;split(.)[1]\&quot;      // \&quot;a.b.c\&quot; -&gt; \&quot;b\&quot; \&quot;split(-)[1][0:3]\&quot; // \&quot;a1234-b5678-c90\&quot; -&gt; \&quot;b56\&quot; | [optional] 
**MatchDeviceType** | Pointer to [**OrgAutoRulesMatchDeviceType**](OrgAutoRulesMatchDeviceType.md) |  | [optional] [default to ORGAUTORULESMATCHDEVICETYPE_AP]
**MatchModel** | Pointer to **string** | optional/additional filter | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Prefix** | Pointer to **NullableString** |  | [optional] 
**Src** | [**OrgAutoRulesSrc**](OrgAutoRulesSrc.md) |  | 
**Subnet** | Pointer to **string** |  | [optional] 
**Suffix** | Pointer to **NullableString** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 

## Methods

### NewOrgAutoRules

`func NewOrgAutoRules(src OrgAutoRulesSrc, ) *OrgAutoRules`

NewOrgAutoRules instantiates a new OrgAutoRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgAutoRulesWithDefaults

`func NewOrgAutoRulesWithDefaults() *OrgAutoRules`

NewOrgAutoRulesWithDefaults instantiates a new OrgAutoRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpression

`func (o *OrgAutoRules) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *OrgAutoRules) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *OrgAutoRules) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *OrgAutoRules) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### SetExpressionNil

`func (o *OrgAutoRules) SetExpressionNil(b bool)`

 SetExpressionNil sets the value for Expression to be an explicit nil

### UnsetExpression
`func (o *OrgAutoRules) UnsetExpression()`

UnsetExpression ensures that no value is present for Expression, not even an explicit nil
### GetMatchDeviceType

`func (o *OrgAutoRules) GetMatchDeviceType() OrgAutoRulesMatchDeviceType`

GetMatchDeviceType returns the MatchDeviceType field if non-nil, zero value otherwise.

### GetMatchDeviceTypeOk

`func (o *OrgAutoRules) GetMatchDeviceTypeOk() (*OrgAutoRulesMatchDeviceType, bool)`

GetMatchDeviceTypeOk returns a tuple with the MatchDeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchDeviceType

`func (o *OrgAutoRules) SetMatchDeviceType(v OrgAutoRulesMatchDeviceType)`

SetMatchDeviceType sets MatchDeviceType field to given value.

### HasMatchDeviceType

`func (o *OrgAutoRules) HasMatchDeviceType() bool`

HasMatchDeviceType returns a boolean if a field has been set.

### GetMatchModel

`func (o *OrgAutoRules) GetMatchModel() string`

GetMatchModel returns the MatchModel field if non-nil, zero value otherwise.

### GetMatchModelOk

`func (o *OrgAutoRules) GetMatchModelOk() (*string, bool)`

GetMatchModelOk returns a tuple with the MatchModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchModel

`func (o *OrgAutoRules) SetMatchModel(v string)`

SetMatchModel sets MatchModel field to given value.

### HasMatchModel

`func (o *OrgAutoRules) HasMatchModel() bool`

HasMatchModel returns a boolean if a field has been set.

### GetModel

`func (o *OrgAutoRules) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *OrgAutoRules) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *OrgAutoRules) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *OrgAutoRules) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetPrefix

`func (o *OrgAutoRules) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *OrgAutoRules) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *OrgAutoRules) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *OrgAutoRules) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### SetPrefixNil

`func (o *OrgAutoRules) SetPrefixNil(b bool)`

 SetPrefixNil sets the value for Prefix to be an explicit nil

### UnsetPrefix
`func (o *OrgAutoRules) UnsetPrefix()`

UnsetPrefix ensures that no value is present for Prefix, not even an explicit nil
### GetSrc

`func (o *OrgAutoRules) GetSrc() OrgAutoRulesSrc`

GetSrc returns the Src field if non-nil, zero value otherwise.

### GetSrcOk

`func (o *OrgAutoRules) GetSrcOk() (*OrgAutoRulesSrc, bool)`

GetSrcOk returns a tuple with the Src field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrc

`func (o *OrgAutoRules) SetSrc(v OrgAutoRulesSrc)`

SetSrc sets Src field to given value.


### GetSubnet

`func (o *OrgAutoRules) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *OrgAutoRules) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *OrgAutoRules) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *OrgAutoRules) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetSuffix

`func (o *OrgAutoRules) GetSuffix() string`

GetSuffix returns the Suffix field if non-nil, zero value otherwise.

### GetSuffixOk

`func (o *OrgAutoRules) GetSuffixOk() (*string, bool)`

GetSuffixOk returns a tuple with the Suffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuffix

`func (o *OrgAutoRules) SetSuffix(v string)`

SetSuffix sets Suffix field to given value.

### HasSuffix

`func (o *OrgAutoRules) HasSuffix() bool`

HasSuffix returns a boolean if a field has been set.

### SetSuffixNil

`func (o *OrgAutoRules) SetSuffixNil(b bool)`

 SetSuffixNil sets the value for Suffix to be an explicit nil

### UnsetSuffix
`func (o *OrgAutoRules) UnsetSuffix()`

UnsetSuffix ensures that no value is present for Suffix, not even an explicit nil
### GetValue

`func (o *OrgAutoRules) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *OrgAutoRules) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *OrgAutoRules) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *OrgAutoRules) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


