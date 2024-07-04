# ResponsePcapStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApMac** | Pointer to **NullableString** |  | [optional] 
**Aps** | Pointer to **[]string** | List of target APs to capture packets | [optional] 
**ClientMac** | Pointer to **NullableString** |  | [optional] 
**Duration** | Pointer to **int32** |  | [optional] 
**Failed** | Pointer to **[]string** | List of APs where configuration attempt failed | [optional] 
**Gateways** | Pointer to **[]string** | List of target Gateways to capture packets if a gateway capture type is specified | [optional] 
**Id** | **string** | unique id for the capture | 
**IncludesMcast** | Pointer to **bool** |  | [optional] 
**MaxNumPackets** | Pointer to **int32** | max number of packets configured by user | [optional] 
**MaxPktLen** | Pointer to **int32** |  | [optional] 
**NumPackets** | Pointer to **int32** | total number of packets captured by all AP, not applicable for type [client, new_assoc] | [optional] 
**Ok** | Pointer to **[]string** | List of target APs successfully configured to capture packets | [optional] 
**PcapAps** | Pointer to [**map[string]ResponsePcapAp**](ResponsePcapAp.md) |  | [optional] 
**RadiotapTcpdumpExpression** | Pointer to **string** | when &#x60;type&#x60;&#x3D;&#x3D;&#x60;radiotap&#x60;, radiotap_tcpdump_expression expression provided by the user | [optional] 
**ScanTcpdumpExpression** | Pointer to **string** | when &#x60;type&#x60;&#x3D;&#x3D;&#x60;scan&#x60;, scan_tcpdump_expression provided by the user | [optional] 
**Ssid** | Pointer to **NullableString** |  | [optional] 
**StartedTime** | Pointer to **int32** |  | [optional] 
**Switches** | Pointer to **[]string** | List of target Switches to capture packets if a switch capture type is specified | [optional] 
**TcpdumpExpression** | Pointer to **string** | tcpdump expression provided by the user (common) | [optional] 
**Type** | [**PcapType**](PcapType.md) |  | 
**WiredTcpdumpExpression** | Pointer to **string** | when &#x60;type&#x60;&#x3D;&#x3D;&#x60;wired&#x60;, wired_tcpdump_expression provided by the user | [optional] 
**WirelessTcpdumpExpression** | Pointer to **string** | when &#x60;type&#x60;&#x3D;&#x3D;&#x60;‘wireless’&#x60;, wireless_tcpdump_expression provided by the user | [optional] 

## Methods

### NewResponsePcapStatus

`func NewResponsePcapStatus(id string, type_ PcapType, ) *ResponsePcapStatus`

NewResponsePcapStatus instantiates a new ResponsePcapStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsePcapStatusWithDefaults

`func NewResponsePcapStatusWithDefaults() *ResponsePcapStatus`

NewResponsePcapStatusWithDefaults instantiates a new ResponsePcapStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApMac

`func (o *ResponsePcapStatus) GetApMac() string`

GetApMac returns the ApMac field if non-nil, zero value otherwise.

### GetApMacOk

`func (o *ResponsePcapStatus) GetApMacOk() (*string, bool)`

GetApMacOk returns a tuple with the ApMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApMac

`func (o *ResponsePcapStatus) SetApMac(v string)`

SetApMac sets ApMac field to given value.

### HasApMac

`func (o *ResponsePcapStatus) HasApMac() bool`

HasApMac returns a boolean if a field has been set.

### SetApMacNil

`func (o *ResponsePcapStatus) SetApMacNil(b bool)`

 SetApMacNil sets the value for ApMac to be an explicit nil

### UnsetApMac
`func (o *ResponsePcapStatus) UnsetApMac()`

UnsetApMac ensures that no value is present for ApMac, not even an explicit nil
### GetAps

`func (o *ResponsePcapStatus) GetAps() []string`

GetAps returns the Aps field if non-nil, zero value otherwise.

### GetApsOk

`func (o *ResponsePcapStatus) GetApsOk() (*[]string, bool)`

GetApsOk returns a tuple with the Aps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAps

`func (o *ResponsePcapStatus) SetAps(v []string)`

SetAps sets Aps field to given value.

### HasAps

`func (o *ResponsePcapStatus) HasAps() bool`

HasAps returns a boolean if a field has been set.

### GetClientMac

`func (o *ResponsePcapStatus) GetClientMac() string`

GetClientMac returns the ClientMac field if non-nil, zero value otherwise.

### GetClientMacOk

`func (o *ResponsePcapStatus) GetClientMacOk() (*string, bool)`

GetClientMacOk returns a tuple with the ClientMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientMac

`func (o *ResponsePcapStatus) SetClientMac(v string)`

SetClientMac sets ClientMac field to given value.

### HasClientMac

`func (o *ResponsePcapStatus) HasClientMac() bool`

HasClientMac returns a boolean if a field has been set.

### SetClientMacNil

`func (o *ResponsePcapStatus) SetClientMacNil(b bool)`

 SetClientMacNil sets the value for ClientMac to be an explicit nil

### UnsetClientMac
`func (o *ResponsePcapStatus) UnsetClientMac()`

UnsetClientMac ensures that no value is present for ClientMac, not even an explicit nil
### GetDuration

`func (o *ResponsePcapStatus) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ResponsePcapStatus) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ResponsePcapStatus) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *ResponsePcapStatus) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetFailed

`func (o *ResponsePcapStatus) GetFailed() []string`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *ResponsePcapStatus) GetFailedOk() (*[]string, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *ResponsePcapStatus) SetFailed(v []string)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *ResponsePcapStatus) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetGateways

`func (o *ResponsePcapStatus) GetGateways() []string`

GetGateways returns the Gateways field if non-nil, zero value otherwise.

### GetGatewaysOk

`func (o *ResponsePcapStatus) GetGatewaysOk() (*[]string, bool)`

GetGatewaysOk returns a tuple with the Gateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateways

`func (o *ResponsePcapStatus) SetGateways(v []string)`

SetGateways sets Gateways field to given value.

### HasGateways

`func (o *ResponsePcapStatus) HasGateways() bool`

HasGateways returns a boolean if a field has been set.

### GetId

`func (o *ResponsePcapStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponsePcapStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponsePcapStatus) SetId(v string)`

SetId sets Id field to given value.


### GetIncludesMcast

`func (o *ResponsePcapStatus) GetIncludesMcast() bool`

GetIncludesMcast returns the IncludesMcast field if non-nil, zero value otherwise.

### GetIncludesMcastOk

`func (o *ResponsePcapStatus) GetIncludesMcastOk() (*bool, bool)`

GetIncludesMcastOk returns a tuple with the IncludesMcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludesMcast

`func (o *ResponsePcapStatus) SetIncludesMcast(v bool)`

SetIncludesMcast sets IncludesMcast field to given value.

### HasIncludesMcast

`func (o *ResponsePcapStatus) HasIncludesMcast() bool`

HasIncludesMcast returns a boolean if a field has been set.

### GetMaxNumPackets

`func (o *ResponsePcapStatus) GetMaxNumPackets() int32`

GetMaxNumPackets returns the MaxNumPackets field if non-nil, zero value otherwise.

### GetMaxNumPacketsOk

`func (o *ResponsePcapStatus) GetMaxNumPacketsOk() (*int32, bool)`

GetMaxNumPacketsOk returns a tuple with the MaxNumPackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumPackets

`func (o *ResponsePcapStatus) SetMaxNumPackets(v int32)`

SetMaxNumPackets sets MaxNumPackets field to given value.

### HasMaxNumPackets

`func (o *ResponsePcapStatus) HasMaxNumPackets() bool`

HasMaxNumPackets returns a boolean if a field has been set.

### GetMaxPktLen

`func (o *ResponsePcapStatus) GetMaxPktLen() int32`

GetMaxPktLen returns the MaxPktLen field if non-nil, zero value otherwise.

### GetMaxPktLenOk

`func (o *ResponsePcapStatus) GetMaxPktLenOk() (*int32, bool)`

GetMaxPktLenOk returns a tuple with the MaxPktLen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPktLen

`func (o *ResponsePcapStatus) SetMaxPktLen(v int32)`

SetMaxPktLen sets MaxPktLen field to given value.

### HasMaxPktLen

`func (o *ResponsePcapStatus) HasMaxPktLen() bool`

HasMaxPktLen returns a boolean if a field has been set.

### GetNumPackets

`func (o *ResponsePcapStatus) GetNumPackets() int32`

GetNumPackets returns the NumPackets field if non-nil, zero value otherwise.

### GetNumPacketsOk

`func (o *ResponsePcapStatus) GetNumPacketsOk() (*int32, bool)`

GetNumPacketsOk returns a tuple with the NumPackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPackets

`func (o *ResponsePcapStatus) SetNumPackets(v int32)`

SetNumPackets sets NumPackets field to given value.

### HasNumPackets

`func (o *ResponsePcapStatus) HasNumPackets() bool`

HasNumPackets returns a boolean if a field has been set.

### GetOk

`func (o *ResponsePcapStatus) GetOk() []string`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *ResponsePcapStatus) GetOkOk() (*[]string, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *ResponsePcapStatus) SetOk(v []string)`

SetOk sets Ok field to given value.

### HasOk

`func (o *ResponsePcapStatus) HasOk() bool`

HasOk returns a boolean if a field has been set.

### GetPcapAps

`func (o *ResponsePcapStatus) GetPcapAps() map[string]ResponsePcapAp`

GetPcapAps returns the PcapAps field if non-nil, zero value otherwise.

### GetPcapApsOk

`func (o *ResponsePcapStatus) GetPcapApsOk() (*map[string]ResponsePcapAp, bool)`

GetPcapApsOk returns a tuple with the PcapAps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcapAps

`func (o *ResponsePcapStatus) SetPcapAps(v map[string]ResponsePcapAp)`

SetPcapAps sets PcapAps field to given value.

### HasPcapAps

`func (o *ResponsePcapStatus) HasPcapAps() bool`

HasPcapAps returns a boolean if a field has been set.

### GetRadiotapTcpdumpExpression

`func (o *ResponsePcapStatus) GetRadiotapTcpdumpExpression() string`

GetRadiotapTcpdumpExpression returns the RadiotapTcpdumpExpression field if non-nil, zero value otherwise.

### GetRadiotapTcpdumpExpressionOk

`func (o *ResponsePcapStatus) GetRadiotapTcpdumpExpressionOk() (*string, bool)`

GetRadiotapTcpdumpExpressionOk returns a tuple with the RadiotapTcpdumpExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiotapTcpdumpExpression

`func (o *ResponsePcapStatus) SetRadiotapTcpdumpExpression(v string)`

SetRadiotapTcpdumpExpression sets RadiotapTcpdumpExpression field to given value.

### HasRadiotapTcpdumpExpression

`func (o *ResponsePcapStatus) HasRadiotapTcpdumpExpression() bool`

HasRadiotapTcpdumpExpression returns a boolean if a field has been set.

### GetScanTcpdumpExpression

`func (o *ResponsePcapStatus) GetScanTcpdumpExpression() string`

GetScanTcpdumpExpression returns the ScanTcpdumpExpression field if non-nil, zero value otherwise.

### GetScanTcpdumpExpressionOk

`func (o *ResponsePcapStatus) GetScanTcpdumpExpressionOk() (*string, bool)`

GetScanTcpdumpExpressionOk returns a tuple with the ScanTcpdumpExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanTcpdumpExpression

`func (o *ResponsePcapStatus) SetScanTcpdumpExpression(v string)`

SetScanTcpdumpExpression sets ScanTcpdumpExpression field to given value.

### HasScanTcpdumpExpression

`func (o *ResponsePcapStatus) HasScanTcpdumpExpression() bool`

HasScanTcpdumpExpression returns a boolean if a field has been set.

### GetSsid

`func (o *ResponsePcapStatus) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *ResponsePcapStatus) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *ResponsePcapStatus) SetSsid(v string)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *ResponsePcapStatus) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### SetSsidNil

`func (o *ResponsePcapStatus) SetSsidNil(b bool)`

 SetSsidNil sets the value for Ssid to be an explicit nil

### UnsetSsid
`func (o *ResponsePcapStatus) UnsetSsid()`

UnsetSsid ensures that no value is present for Ssid, not even an explicit nil
### GetStartedTime

`func (o *ResponsePcapStatus) GetStartedTime() int32`

GetStartedTime returns the StartedTime field if non-nil, zero value otherwise.

### GetStartedTimeOk

`func (o *ResponsePcapStatus) GetStartedTimeOk() (*int32, bool)`

GetStartedTimeOk returns a tuple with the StartedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedTime

`func (o *ResponsePcapStatus) SetStartedTime(v int32)`

SetStartedTime sets StartedTime field to given value.

### HasStartedTime

`func (o *ResponsePcapStatus) HasStartedTime() bool`

HasStartedTime returns a boolean if a field has been set.

### GetSwitches

`func (o *ResponsePcapStatus) GetSwitches() []string`

GetSwitches returns the Switches field if non-nil, zero value otherwise.

### GetSwitchesOk

`func (o *ResponsePcapStatus) GetSwitchesOk() (*[]string, bool)`

GetSwitchesOk returns a tuple with the Switches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitches

`func (o *ResponsePcapStatus) SetSwitches(v []string)`

SetSwitches sets Switches field to given value.

### HasSwitches

`func (o *ResponsePcapStatus) HasSwitches() bool`

HasSwitches returns a boolean if a field has been set.

### GetTcpdumpExpression

`func (o *ResponsePcapStatus) GetTcpdumpExpression() string`

GetTcpdumpExpression returns the TcpdumpExpression field if non-nil, zero value otherwise.

### GetTcpdumpExpressionOk

`func (o *ResponsePcapStatus) GetTcpdumpExpressionOk() (*string, bool)`

GetTcpdumpExpressionOk returns a tuple with the TcpdumpExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpdumpExpression

`func (o *ResponsePcapStatus) SetTcpdumpExpression(v string)`

SetTcpdumpExpression sets TcpdumpExpression field to given value.

### HasTcpdumpExpression

`func (o *ResponsePcapStatus) HasTcpdumpExpression() bool`

HasTcpdumpExpression returns a boolean if a field has been set.

### GetType

`func (o *ResponsePcapStatus) GetType() PcapType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponsePcapStatus) GetTypeOk() (*PcapType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponsePcapStatus) SetType(v PcapType)`

SetType sets Type field to given value.


### GetWiredTcpdumpExpression

`func (o *ResponsePcapStatus) GetWiredTcpdumpExpression() string`

GetWiredTcpdumpExpression returns the WiredTcpdumpExpression field if non-nil, zero value otherwise.

### GetWiredTcpdumpExpressionOk

`func (o *ResponsePcapStatus) GetWiredTcpdumpExpressionOk() (*string, bool)`

GetWiredTcpdumpExpressionOk returns a tuple with the WiredTcpdumpExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWiredTcpdumpExpression

`func (o *ResponsePcapStatus) SetWiredTcpdumpExpression(v string)`

SetWiredTcpdumpExpression sets WiredTcpdumpExpression field to given value.

### HasWiredTcpdumpExpression

`func (o *ResponsePcapStatus) HasWiredTcpdumpExpression() bool`

HasWiredTcpdumpExpression returns a boolean if a field has been set.

### GetWirelessTcpdumpExpression

`func (o *ResponsePcapStatus) GetWirelessTcpdumpExpression() string`

GetWirelessTcpdumpExpression returns the WirelessTcpdumpExpression field if non-nil, zero value otherwise.

### GetWirelessTcpdumpExpressionOk

`func (o *ResponsePcapStatus) GetWirelessTcpdumpExpressionOk() (*string, bool)`

GetWirelessTcpdumpExpressionOk returns a tuple with the WirelessTcpdumpExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWirelessTcpdumpExpression

`func (o *ResponsePcapStatus) SetWirelessTcpdumpExpression(v string)`

SetWirelessTcpdumpExpression sets WirelessTcpdumpExpression field to given value.

### HasWirelessTcpdumpExpression

`func (o *ResponsePcapStatus) HasWirelessTcpdumpExpression() bool`

HasWirelessTcpdumpExpression returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


