# ResponseSsoFailureSearchItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Detail** | **string** |  | 
**SamlAssertionXml** | **string** |  | 
**Timestamp** | **float32** |  | 

## Methods

### NewResponseSsoFailureSearchItem

`func NewResponseSsoFailureSearchItem(detail string, samlAssertionXml string, timestamp float32, ) *ResponseSsoFailureSearchItem`

NewResponseSsoFailureSearchItem instantiates a new ResponseSsoFailureSearchItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseSsoFailureSearchItemWithDefaults

`func NewResponseSsoFailureSearchItemWithDefaults() *ResponseSsoFailureSearchItem`

NewResponseSsoFailureSearchItemWithDefaults instantiates a new ResponseSsoFailureSearchItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetail

`func (o *ResponseSsoFailureSearchItem) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ResponseSsoFailureSearchItem) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ResponseSsoFailureSearchItem) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetSamlAssertionXml

`func (o *ResponseSsoFailureSearchItem) GetSamlAssertionXml() string`

GetSamlAssertionXml returns the SamlAssertionXml field if non-nil, zero value otherwise.

### GetSamlAssertionXmlOk

`func (o *ResponseSsoFailureSearchItem) GetSamlAssertionXmlOk() (*string, bool)`

GetSamlAssertionXmlOk returns a tuple with the SamlAssertionXml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlAssertionXml

`func (o *ResponseSsoFailureSearchItem) SetSamlAssertionXml(v string)`

SetSamlAssertionXml sets SamlAssertionXml field to given value.


### GetTimestamp

`func (o *ResponseSsoFailureSearchItem) GetTimestamp() float32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ResponseSsoFailureSearchItem) GetTimestampOk() (*float32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ResponseSsoFailureSearchItem) SetTimestamp(v float32)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


