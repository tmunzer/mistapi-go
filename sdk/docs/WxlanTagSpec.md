# WxlanTagSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortRange** | Pointer to **string** | matched dst port, \&quot;0\&quot; means any | [optional] [default to "0"]
**Protocol** | Pointer to **string** | tcp / udp / icmp / gre / any / \&quot;:protocol_number\&quot;, &#x60;protocol_number&#x60; is between 1-254 | [optional] [default to "any"]
**Subnets** | Pointer to **[]string** | matched dst subnet | [optional] [default to ["0.0.0.0/0"]]

## Methods

### NewWxlanTagSpec

`func NewWxlanTagSpec() *WxlanTagSpec`

NewWxlanTagSpec instantiates a new WxlanTagSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWxlanTagSpecWithDefaults

`func NewWxlanTagSpecWithDefaults() *WxlanTagSpec`

NewWxlanTagSpecWithDefaults instantiates a new WxlanTagSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortRange

`func (o *WxlanTagSpec) GetPortRange() string`

GetPortRange returns the PortRange field if non-nil, zero value otherwise.

### GetPortRangeOk

`func (o *WxlanTagSpec) GetPortRangeOk() (*string, bool)`

GetPortRangeOk returns a tuple with the PortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRange

`func (o *WxlanTagSpec) SetPortRange(v string)`

SetPortRange sets PortRange field to given value.

### HasPortRange

`func (o *WxlanTagSpec) HasPortRange() bool`

HasPortRange returns a boolean if a field has been set.

### GetProtocol

`func (o *WxlanTagSpec) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *WxlanTagSpec) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *WxlanTagSpec) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *WxlanTagSpec) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSubnets

`func (o *WxlanTagSpec) GetSubnets() []string`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *WxlanTagSpec) GetSubnetsOk() (*[]string, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *WxlanTagSpec) SetSubnets(v []string)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *WxlanTagSpec) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


