# ApUplinkPortConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dot1x** | Pointer to **bool** | Whether to do 802.1x against uplink switch. When enaled, AP cert will be used to do EAP-TLS and the Org&#39;s CA Cert has to be provisioned at the switch | [optional] [default to false]
**KeepWlansUpIfDown** | Pointer to **bool** | by default, WLANs are disabled when uplink is down. In some scenario, like SiteSurvey, one would want the AP to keep sending beacons. | [optional] [default to false]

## Methods

### NewApUplinkPortConfig

`func NewApUplinkPortConfig() *ApUplinkPortConfig`

NewApUplinkPortConfig instantiates a new ApUplinkPortConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApUplinkPortConfigWithDefaults

`func NewApUplinkPortConfigWithDefaults() *ApUplinkPortConfig`

NewApUplinkPortConfigWithDefaults instantiates a new ApUplinkPortConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDot1x

`func (o *ApUplinkPortConfig) GetDot1x() bool`

GetDot1x returns the Dot1x field if non-nil, zero value otherwise.

### GetDot1xOk

`func (o *ApUplinkPortConfig) GetDot1xOk() (*bool, bool)`

GetDot1xOk returns a tuple with the Dot1x field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDot1x

`func (o *ApUplinkPortConfig) SetDot1x(v bool)`

SetDot1x sets Dot1x field to given value.

### HasDot1x

`func (o *ApUplinkPortConfig) HasDot1x() bool`

HasDot1x returns a boolean if a field has been set.

### GetKeepWlansUpIfDown

`func (o *ApUplinkPortConfig) GetKeepWlansUpIfDown() bool`

GetKeepWlansUpIfDown returns the KeepWlansUpIfDown field if non-nil, zero value otherwise.

### GetKeepWlansUpIfDownOk

`func (o *ApUplinkPortConfig) GetKeepWlansUpIfDownOk() (*bool, bool)`

GetKeepWlansUpIfDownOk returns a tuple with the KeepWlansUpIfDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepWlansUpIfDown

`func (o *ApUplinkPortConfig) SetKeepWlansUpIfDown(v bool)`

SetKeepWlansUpIfDown sets KeepWlansUpIfDown field to given value.

### HasKeepWlansUpIfDown

`func (o *ApUplinkPortConfig) HasKeepWlansUpIfDown() bool`

HasKeepWlansUpIfDown returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


