# Snmpv3ConfigTargetAddressItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** |  | [optional] 
**AddressMask** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] [default to 161]
**TagList** | Pointer to **string** | &lt;refer to notify tag, can be multiple with blank | [optional] 
**TargetAddressName** | Pointer to **string** |  | [optional] 
**TargetParameters** | Pointer to **string** | refer to notify target parameters name | [optional] 

## Methods

### NewSnmpv3ConfigTargetAddressItem

`func NewSnmpv3ConfigTargetAddressItem() *Snmpv3ConfigTargetAddressItem`

NewSnmpv3ConfigTargetAddressItem instantiates a new Snmpv3ConfigTargetAddressItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnmpv3ConfigTargetAddressItemWithDefaults

`func NewSnmpv3ConfigTargetAddressItemWithDefaults() *Snmpv3ConfigTargetAddressItem`

NewSnmpv3ConfigTargetAddressItemWithDefaults instantiates a new Snmpv3ConfigTargetAddressItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *Snmpv3ConfigTargetAddressItem) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Snmpv3ConfigTargetAddressItem) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Snmpv3ConfigTargetAddressItem) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Snmpv3ConfigTargetAddressItem) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAddressMask

`func (o *Snmpv3ConfigTargetAddressItem) GetAddressMask() string`

GetAddressMask returns the AddressMask field if non-nil, zero value otherwise.

### GetAddressMaskOk

`func (o *Snmpv3ConfigTargetAddressItem) GetAddressMaskOk() (*string, bool)`

GetAddressMaskOk returns a tuple with the AddressMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressMask

`func (o *Snmpv3ConfigTargetAddressItem) SetAddressMask(v string)`

SetAddressMask sets AddressMask field to given value.

### HasAddressMask

`func (o *Snmpv3ConfigTargetAddressItem) HasAddressMask() bool`

HasAddressMask returns a boolean if a field has been set.

### GetPort

`func (o *Snmpv3ConfigTargetAddressItem) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *Snmpv3ConfigTargetAddressItem) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *Snmpv3ConfigTargetAddressItem) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *Snmpv3ConfigTargetAddressItem) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetTagList

`func (o *Snmpv3ConfigTargetAddressItem) GetTagList() string`

GetTagList returns the TagList field if non-nil, zero value otherwise.

### GetTagListOk

`func (o *Snmpv3ConfigTargetAddressItem) GetTagListOk() (*string, bool)`

GetTagListOk returns a tuple with the TagList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagList

`func (o *Snmpv3ConfigTargetAddressItem) SetTagList(v string)`

SetTagList sets TagList field to given value.

### HasTagList

`func (o *Snmpv3ConfigTargetAddressItem) HasTagList() bool`

HasTagList returns a boolean if a field has been set.

### GetTargetAddressName

`func (o *Snmpv3ConfigTargetAddressItem) GetTargetAddressName() string`

GetTargetAddressName returns the TargetAddressName field if non-nil, zero value otherwise.

### GetTargetAddressNameOk

`func (o *Snmpv3ConfigTargetAddressItem) GetTargetAddressNameOk() (*string, bool)`

GetTargetAddressNameOk returns a tuple with the TargetAddressName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAddressName

`func (o *Snmpv3ConfigTargetAddressItem) SetTargetAddressName(v string)`

SetTargetAddressName sets TargetAddressName field to given value.

### HasTargetAddressName

`func (o *Snmpv3ConfigTargetAddressItem) HasTargetAddressName() bool`

HasTargetAddressName returns a boolean if a field has been set.

### GetTargetParameters

`func (o *Snmpv3ConfigTargetAddressItem) GetTargetParameters() string`

GetTargetParameters returns the TargetParameters field if non-nil, zero value otherwise.

### GetTargetParametersOk

`func (o *Snmpv3ConfigTargetAddressItem) GetTargetParametersOk() (*string, bool)`

GetTargetParametersOk returns a tuple with the TargetParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetParameters

`func (o *Snmpv3ConfigTargetAddressItem) SetTargetParameters(v string)`

SetTargetParameters sets TargetParameters field to given value.

### HasTargetParameters

`func (o *Snmpv3ConfigTargetAddressItem) HasTargetParameters() bool`

HasTargetParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


