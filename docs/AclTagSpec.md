# AclTagSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortRange** | Pointer to **string** | matched dst port, \&quot;0\&quot; means any | [optional] [default to "80"]
**Protocol** | Pointer to **string** | &#x60;tcp&#x60; / &#x60;udp&#x60; / &#x60;icmp&#x60; / &#x60;gre&#x60; / &#x60;any&#x60; / &#x60;:protocol_number&#x60;. &#x60;protocol_number&#x60; is between 1-254 | [optional] [default to "any"]

## Methods

### NewAclTagSpec

`func NewAclTagSpec() *AclTagSpec`

NewAclTagSpec instantiates a new AclTagSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAclTagSpecWithDefaults

`func NewAclTagSpecWithDefaults() *AclTagSpec`

NewAclTagSpecWithDefaults instantiates a new AclTagSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortRange

`func (o *AclTagSpec) GetPortRange() string`

GetPortRange returns the PortRange field if non-nil, zero value otherwise.

### GetPortRangeOk

`func (o *AclTagSpec) GetPortRangeOk() (*string, bool)`

GetPortRangeOk returns a tuple with the PortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRange

`func (o *AclTagSpec) SetPortRange(v string)`

SetPortRange sets PortRange field to given value.

### HasPortRange

`func (o *AclTagSpec) HasPortRange() bool`

HasPortRange returns a boolean if a field has been set.

### GetProtocol

`func (o *AclTagSpec) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *AclTagSpec) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *AclTagSpec) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *AclTagSpec) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


