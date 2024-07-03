# Snmpv3ConfigTargetParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageProcessingModel** | Pointer to [**Snmpv3ConfigTargetParamMessProcessModel**](Snmpv3ConfigTargetParamMessProcessModel.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NotifyFilter** | Pointer to **string** | refer to profile-name in notify_filter | [optional] 
**SecurityLevel** | Pointer to [**Snmpv3ConfigTargetParamSecurityLevel**](Snmpv3ConfigTargetParamSecurityLevel.md) |  | [optional] 
**SecurityModel** | Pointer to [**Snmpv3ConfigTargetParamSecurityModel**](Snmpv3ConfigTargetParamSecurityModel.md) |  | [optional] 
**SecurityName** | Pointer to **string** | refer to security_name in usm | [optional] 

## Methods

### NewSnmpv3ConfigTargetParam

`func NewSnmpv3ConfigTargetParam() *Snmpv3ConfigTargetParam`

NewSnmpv3ConfigTargetParam instantiates a new Snmpv3ConfigTargetParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnmpv3ConfigTargetParamWithDefaults

`func NewSnmpv3ConfigTargetParamWithDefaults() *Snmpv3ConfigTargetParam`

NewSnmpv3ConfigTargetParamWithDefaults instantiates a new Snmpv3ConfigTargetParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageProcessingModel

`func (o *Snmpv3ConfigTargetParam) GetMessageProcessingModel() Snmpv3ConfigTargetParamMessProcessModel`

GetMessageProcessingModel returns the MessageProcessingModel field if non-nil, zero value otherwise.

### GetMessageProcessingModelOk

`func (o *Snmpv3ConfigTargetParam) GetMessageProcessingModelOk() (*Snmpv3ConfigTargetParamMessProcessModel, bool)`

GetMessageProcessingModelOk returns a tuple with the MessageProcessingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageProcessingModel

`func (o *Snmpv3ConfigTargetParam) SetMessageProcessingModel(v Snmpv3ConfigTargetParamMessProcessModel)`

SetMessageProcessingModel sets MessageProcessingModel field to given value.

### HasMessageProcessingModel

`func (o *Snmpv3ConfigTargetParam) HasMessageProcessingModel() bool`

HasMessageProcessingModel returns a boolean if a field has been set.

### GetName

`func (o *Snmpv3ConfigTargetParam) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Snmpv3ConfigTargetParam) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Snmpv3ConfigTargetParam) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Snmpv3ConfigTargetParam) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotifyFilter

`func (o *Snmpv3ConfigTargetParam) GetNotifyFilter() string`

GetNotifyFilter returns the NotifyFilter field if non-nil, zero value otherwise.

### GetNotifyFilterOk

`func (o *Snmpv3ConfigTargetParam) GetNotifyFilterOk() (*string, bool)`

GetNotifyFilterOk returns a tuple with the NotifyFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyFilter

`func (o *Snmpv3ConfigTargetParam) SetNotifyFilter(v string)`

SetNotifyFilter sets NotifyFilter field to given value.

### HasNotifyFilter

`func (o *Snmpv3ConfigTargetParam) HasNotifyFilter() bool`

HasNotifyFilter returns a boolean if a field has been set.

### GetSecurityLevel

`func (o *Snmpv3ConfigTargetParam) GetSecurityLevel() Snmpv3ConfigTargetParamSecurityLevel`

GetSecurityLevel returns the SecurityLevel field if non-nil, zero value otherwise.

### GetSecurityLevelOk

`func (o *Snmpv3ConfigTargetParam) GetSecurityLevelOk() (*Snmpv3ConfigTargetParamSecurityLevel, bool)`

GetSecurityLevelOk returns a tuple with the SecurityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityLevel

`func (o *Snmpv3ConfigTargetParam) SetSecurityLevel(v Snmpv3ConfigTargetParamSecurityLevel)`

SetSecurityLevel sets SecurityLevel field to given value.

### HasSecurityLevel

`func (o *Snmpv3ConfigTargetParam) HasSecurityLevel() bool`

HasSecurityLevel returns a boolean if a field has been set.

### GetSecurityModel

`func (o *Snmpv3ConfigTargetParam) GetSecurityModel() Snmpv3ConfigTargetParamSecurityModel`

GetSecurityModel returns the SecurityModel field if non-nil, zero value otherwise.

### GetSecurityModelOk

`func (o *Snmpv3ConfigTargetParam) GetSecurityModelOk() (*Snmpv3ConfigTargetParamSecurityModel, bool)`

GetSecurityModelOk returns a tuple with the SecurityModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityModel

`func (o *Snmpv3ConfigTargetParam) SetSecurityModel(v Snmpv3ConfigTargetParamSecurityModel)`

SetSecurityModel sets SecurityModel field to given value.

### HasSecurityModel

`func (o *Snmpv3ConfigTargetParam) HasSecurityModel() bool`

HasSecurityModel returns a boolean if a field has been set.

### GetSecurityName

`func (o *Snmpv3ConfigTargetParam) GetSecurityName() string`

GetSecurityName returns the SecurityName field if non-nil, zero value otherwise.

### GetSecurityNameOk

`func (o *Snmpv3ConfigTargetParam) GetSecurityNameOk() (*string, bool)`

GetSecurityNameOk returns a tuple with the SecurityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityName

`func (o *Snmpv3ConfigTargetParam) SetSecurityName(v string)`

SetSecurityName sets SecurityName field to given value.

### HasSecurityName

`func (o *Snmpv3ConfigTargetParam) HasSecurityName() bool`

HasSecurityName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


