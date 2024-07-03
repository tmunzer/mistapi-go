# IdpProfileMatching

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttackName** | Pointer to **[]string** |  | [optional] 
**DstSubnet** | Pointer to **[]string** |  | [optional] 
**Severity** | Pointer to [**[]IdpProfileMatchingSeverityValue**](IdpProfileMatchingSeverityValue.md) |  | [optional] 

## Methods

### NewIdpProfileMatching

`func NewIdpProfileMatching() *IdpProfileMatching`

NewIdpProfileMatching instantiates a new IdpProfileMatching object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpProfileMatchingWithDefaults

`func NewIdpProfileMatchingWithDefaults() *IdpProfileMatching`

NewIdpProfileMatchingWithDefaults instantiates a new IdpProfileMatching object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttackName

`func (o *IdpProfileMatching) GetAttackName() []string`

GetAttackName returns the AttackName field if non-nil, zero value otherwise.

### GetAttackNameOk

`func (o *IdpProfileMatching) GetAttackNameOk() (*[]string, bool)`

GetAttackNameOk returns a tuple with the AttackName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackName

`func (o *IdpProfileMatching) SetAttackName(v []string)`

SetAttackName sets AttackName field to given value.

### HasAttackName

`func (o *IdpProfileMatching) HasAttackName() bool`

HasAttackName returns a boolean if a field has been set.

### GetDstSubnet

`func (o *IdpProfileMatching) GetDstSubnet() []string`

GetDstSubnet returns the DstSubnet field if non-nil, zero value otherwise.

### GetDstSubnetOk

`func (o *IdpProfileMatching) GetDstSubnetOk() (*[]string, bool)`

GetDstSubnetOk returns a tuple with the DstSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstSubnet

`func (o *IdpProfileMatching) SetDstSubnet(v []string)`

SetDstSubnet sets DstSubnet field to given value.

### HasDstSubnet

`func (o *IdpProfileMatching) HasDstSubnet() bool`

HasDstSubnet returns a boolean if a field has been set.

### GetSeverity

`func (o *IdpProfileMatching) GetSeverity() []IdpProfileMatchingSeverityValue`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *IdpProfileMatching) GetSeverityOk() (*[]IdpProfileMatchingSeverityValue, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *IdpProfileMatching) SetSeverity(v []IdpProfileMatchingSeverityValue)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *IdpProfileMatching) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


