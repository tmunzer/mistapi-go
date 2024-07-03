# SimpleAlert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArpFailure** | Pointer to [**SimpleAlertArpFailure**](SimpleAlertArpFailure.md) |  | [optional] 
**DhcpFailure** | Pointer to [**SimpleAlertDhcpFailure**](SimpleAlertDhcpFailure.md) |  | [optional] 
**DnsFailure** | Pointer to [**SimpleAlertDnsFailure**](SimpleAlertDnsFailure.md) |  | [optional] 

## Methods

### NewSimpleAlert

`func NewSimpleAlert() *SimpleAlert`

NewSimpleAlert instantiates a new SimpleAlert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimpleAlertWithDefaults

`func NewSimpleAlertWithDefaults() *SimpleAlert`

NewSimpleAlertWithDefaults instantiates a new SimpleAlert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArpFailure

`func (o *SimpleAlert) GetArpFailure() SimpleAlertArpFailure`

GetArpFailure returns the ArpFailure field if non-nil, zero value otherwise.

### GetArpFailureOk

`func (o *SimpleAlert) GetArpFailureOk() (*SimpleAlertArpFailure, bool)`

GetArpFailureOk returns a tuple with the ArpFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArpFailure

`func (o *SimpleAlert) SetArpFailure(v SimpleAlertArpFailure)`

SetArpFailure sets ArpFailure field to given value.

### HasArpFailure

`func (o *SimpleAlert) HasArpFailure() bool`

HasArpFailure returns a boolean if a field has been set.

### GetDhcpFailure

`func (o *SimpleAlert) GetDhcpFailure() SimpleAlertDhcpFailure`

GetDhcpFailure returns the DhcpFailure field if non-nil, zero value otherwise.

### GetDhcpFailureOk

`func (o *SimpleAlert) GetDhcpFailureOk() (*SimpleAlertDhcpFailure, bool)`

GetDhcpFailureOk returns a tuple with the DhcpFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpFailure

`func (o *SimpleAlert) SetDhcpFailure(v SimpleAlertDhcpFailure)`

SetDhcpFailure sets DhcpFailure field to given value.

### HasDhcpFailure

`func (o *SimpleAlert) HasDhcpFailure() bool`

HasDhcpFailure returns a boolean if a field has been set.

### GetDnsFailure

`func (o *SimpleAlert) GetDnsFailure() SimpleAlertDnsFailure`

GetDnsFailure returns the DnsFailure field if non-nil, zero value otherwise.

### GetDnsFailureOk

`func (o *SimpleAlert) GetDnsFailureOk() (*SimpleAlertDnsFailure, bool)`

GetDnsFailureOk returns a tuple with the DnsFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsFailure

`func (o *SimpleAlert) SetDnsFailure(v SimpleAlertDnsFailure)`

SetDnsFailure sets DnsFailure field to given value.

### HasDnsFailure

`func (o *SimpleAlert) HasDnsFailure() bool`

HasDnsFailure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


