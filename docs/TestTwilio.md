# TestTwilio

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | **string** | One of the numbers you have in your Twilio account | 
**To** | **string** | Phone number of the recipient of SMS | 
**TwilioAuthToken** | **string** | Auth Token associated with twilio account | 
**TwilioSid** | **string** | Twilio Account SID | 

## Methods

### NewTestTwilio

`func NewTestTwilio(from string, to string, twilioAuthToken string, twilioSid string, ) *TestTwilio`

NewTestTwilio instantiates a new TestTwilio object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestTwilioWithDefaults

`func NewTestTwilioWithDefaults() *TestTwilio`

NewTestTwilioWithDefaults instantiates a new TestTwilio object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *TestTwilio) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *TestTwilio) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *TestTwilio) SetFrom(v string)`

SetFrom sets From field to given value.


### GetTo

`func (o *TestTwilio) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *TestTwilio) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *TestTwilio) SetTo(v string)`

SetTo sets To field to given value.


### GetTwilioAuthToken

`func (o *TestTwilio) GetTwilioAuthToken() string`

GetTwilioAuthToken returns the TwilioAuthToken field if non-nil, zero value otherwise.

### GetTwilioAuthTokenOk

`func (o *TestTwilio) GetTwilioAuthTokenOk() (*string, bool)`

GetTwilioAuthTokenOk returns a tuple with the TwilioAuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwilioAuthToken

`func (o *TestTwilio) SetTwilioAuthToken(v string)`

SetTwilioAuthToken sets TwilioAuthToken field to given value.


### GetTwilioSid

`func (o *TestTwilio) GetTwilioSid() string`

GetTwilioSid returns the TwilioSid field if non-nil, zero value otherwise.

### GetTwilioSidOk

`func (o *TestTwilio) GetTwilioSidOk() (*string, bool)`

GetTwilioSidOk returns a tuple with the TwilioSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwilioSid

`func (o *TestTwilio) SetTwilioSid(v string)`

SetTwilioSid sets TwilioSid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


