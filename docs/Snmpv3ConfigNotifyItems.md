# Snmpv3ConfigNotifyItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Tag** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**Snmpv3ConfigNotifyType**](Snmpv3ConfigNotifyType.md) |  | [optional] 

## Methods

### NewSnmpv3ConfigNotifyItems

`func NewSnmpv3ConfigNotifyItems() *Snmpv3ConfigNotifyItems`

NewSnmpv3ConfigNotifyItems instantiates a new Snmpv3ConfigNotifyItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnmpv3ConfigNotifyItemsWithDefaults

`func NewSnmpv3ConfigNotifyItemsWithDefaults() *Snmpv3ConfigNotifyItems`

NewSnmpv3ConfigNotifyItemsWithDefaults instantiates a new Snmpv3ConfigNotifyItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Snmpv3ConfigNotifyItems) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Snmpv3ConfigNotifyItems) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Snmpv3ConfigNotifyItems) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Snmpv3ConfigNotifyItems) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTag

`func (o *Snmpv3ConfigNotifyItems) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *Snmpv3ConfigNotifyItems) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *Snmpv3ConfigNotifyItems) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *Snmpv3ConfigNotifyItems) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetType

`func (o *Snmpv3ConfigNotifyItems) GetType() Snmpv3ConfigNotifyType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Snmpv3ConfigNotifyItems) GetTypeOk() (*Snmpv3ConfigNotifyType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Snmpv3ConfigNotifyItems) SetType(v Snmpv3ConfigNotifyType)`

SetType sets Type field to given value.

### HasType

`func (o *Snmpv3ConfigNotifyItems) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


