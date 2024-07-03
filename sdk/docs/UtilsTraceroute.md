# UtilsTraceroute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** | host name | [optional] 
**Network** | Pointer to **string** | for SSR, optional, the source to initiate traceroute from | [optional] [default to "internal"]
**Port** | Pointer to **int32** | when &#x60;protocol&#x60;&#x3D;&#x3D;&#x60;udp&#x60;, the udp port to use | [optional] [default to 33434]
**Protocol** | Pointer to [**UtilsTracerouteProtocol**](UtilsTracerouteProtocol.md) |  | [optional] [default to UTILSTRACEROUTEPROTOCOL_UDP]
**Timeout** | Pointer to **int32** | maximum time in seconds to wait for the response | [optional] [default to 60]
**Vrf** | Pointer to **string** | for SRX, optional, the source to initiate traceroute from. by default, master VRF/RI is assumed | [optional] 

## Methods

### NewUtilsTraceroute

`func NewUtilsTraceroute() *UtilsTraceroute`

NewUtilsTraceroute instantiates a new UtilsTraceroute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUtilsTracerouteWithDefaults

`func NewUtilsTracerouteWithDefaults() *UtilsTraceroute`

NewUtilsTracerouteWithDefaults instantiates a new UtilsTraceroute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *UtilsTraceroute) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *UtilsTraceroute) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *UtilsTraceroute) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *UtilsTraceroute) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetNetwork

`func (o *UtilsTraceroute) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *UtilsTraceroute) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *UtilsTraceroute) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *UtilsTraceroute) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetPort

`func (o *UtilsTraceroute) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *UtilsTraceroute) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *UtilsTraceroute) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *UtilsTraceroute) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProtocol

`func (o *UtilsTraceroute) GetProtocol() UtilsTracerouteProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *UtilsTraceroute) GetProtocolOk() (*UtilsTracerouteProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *UtilsTraceroute) SetProtocol(v UtilsTracerouteProtocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *UtilsTraceroute) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetTimeout

`func (o *UtilsTraceroute) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *UtilsTraceroute) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *UtilsTraceroute) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *UtilsTraceroute) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetVrf

`func (o *UtilsTraceroute) GetVrf() string`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *UtilsTraceroute) GetVrfOk() (*string, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *UtilsTraceroute) SetVrf(v string)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *UtilsTraceroute) HasVrf() bool`

HasVrf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


