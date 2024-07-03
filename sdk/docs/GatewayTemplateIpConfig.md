# GatewayTemplateIpConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** |  | [optional] 
**Netmask** | Pointer to **string** |  | [optional] 
**SecondaryIps** | Pointer to **[]string** | optional list of secondary IPs in CIDR format | [optional] 

## Methods

### NewGatewayTemplateIpConfig

`func NewGatewayTemplateIpConfig() *GatewayTemplateIpConfig`

NewGatewayTemplateIpConfig instantiates a new GatewayTemplateIpConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayTemplateIpConfigWithDefaults

`func NewGatewayTemplateIpConfigWithDefaults() *GatewayTemplateIpConfig`

NewGatewayTemplateIpConfigWithDefaults instantiates a new GatewayTemplateIpConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *GatewayTemplateIpConfig) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *GatewayTemplateIpConfig) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *GatewayTemplateIpConfig) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *GatewayTemplateIpConfig) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetNetmask

`func (o *GatewayTemplateIpConfig) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *GatewayTemplateIpConfig) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *GatewayTemplateIpConfig) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *GatewayTemplateIpConfig) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetSecondaryIps

`func (o *GatewayTemplateIpConfig) GetSecondaryIps() []string`

GetSecondaryIps returns the SecondaryIps field if non-nil, zero value otherwise.

### GetSecondaryIpsOk

`func (o *GatewayTemplateIpConfig) GetSecondaryIpsOk() (*[]string, bool)`

GetSecondaryIpsOk returns a tuple with the SecondaryIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryIps

`func (o *GatewayTemplateIpConfig) SetSecondaryIps(v []string)`

SetSecondaryIps sets SecondaryIps field to given value.

### HasSecondaryIps

`func (o *GatewayTemplateIpConfig) HasSecondaryIps() bool`

HasSecondaryIps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


