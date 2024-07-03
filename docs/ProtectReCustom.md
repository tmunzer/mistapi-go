# ProtectReCustom

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortRange** | Pointer to **string** | matched dst port, \&quot;0\&quot; means any | [optional] [default to "0"]
**Protocol** | Pointer to [**ProtectReCustomProtocol**](ProtectReCustomProtocol.md) |  | [optional] [default to PROTECTRECUSTOMPROTOCOL_ANY]
**Subnet** | Pointer to **[]string** |  | [optional] 

## Methods

### NewProtectReCustom

`func NewProtectReCustom() *ProtectReCustom`

NewProtectReCustom instantiates a new ProtectReCustom object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectReCustomWithDefaults

`func NewProtectReCustomWithDefaults() *ProtectReCustom`

NewProtectReCustomWithDefaults instantiates a new ProtectReCustom object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortRange

`func (o *ProtectReCustom) GetPortRange() string`

GetPortRange returns the PortRange field if non-nil, zero value otherwise.

### GetPortRangeOk

`func (o *ProtectReCustom) GetPortRangeOk() (*string, bool)`

GetPortRangeOk returns a tuple with the PortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRange

`func (o *ProtectReCustom) SetPortRange(v string)`

SetPortRange sets PortRange field to given value.

### HasPortRange

`func (o *ProtectReCustom) HasPortRange() bool`

HasPortRange returns a boolean if a field has been set.

### GetProtocol

`func (o *ProtectReCustom) GetProtocol() ProtectReCustomProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ProtectReCustom) GetProtocolOk() (*ProtectReCustomProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ProtectReCustom) SetProtocol(v ProtectReCustomProtocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *ProtectReCustom) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSubnet

`func (o *ProtectReCustom) GetSubnet() []string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *ProtectReCustom) GetSubnetOk() (*[]string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *ProtectReCustom) SetSubnet(v []string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *ProtectReCustom) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


