# UtilsClearBpdu

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | Pointer to **string** | the port on which to clear the detected BPDU error, or &#x60;all&#x60; for all ports | [optional] 

## Methods

### NewUtilsClearBpdu

`func NewUtilsClearBpdu() *UtilsClearBpdu`

NewUtilsClearBpdu instantiates a new UtilsClearBpdu object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUtilsClearBpduWithDefaults

`func NewUtilsClearBpduWithDefaults() *UtilsClearBpdu`

NewUtilsClearBpduWithDefaults instantiates a new UtilsClearBpdu object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *UtilsClearBpdu) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *UtilsClearBpdu) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *UtilsClearBpdu) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *UtilsClearBpdu) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


