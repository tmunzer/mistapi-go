# TestTelstra

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TelstraClientId** | **string** | Telstra client id | 
**TelstraClientSecret** | **string** | Telstra client secret | 
**To** | **string** | Phone number of the recipient of SMS with country code | 

## Methods

### NewTestTelstra

`func NewTestTelstra(telstraClientId string, telstraClientSecret string, to string, ) *TestTelstra`

NewTestTelstra instantiates a new TestTelstra object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestTelstraWithDefaults

`func NewTestTelstraWithDefaults() *TestTelstra`

NewTestTelstraWithDefaults instantiates a new TestTelstra object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTelstraClientId

`func (o *TestTelstra) GetTelstraClientId() string`

GetTelstraClientId returns the TelstraClientId field if non-nil, zero value otherwise.

### GetTelstraClientIdOk

`func (o *TestTelstra) GetTelstraClientIdOk() (*string, bool)`

GetTelstraClientIdOk returns a tuple with the TelstraClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelstraClientId

`func (o *TestTelstra) SetTelstraClientId(v string)`

SetTelstraClientId sets TelstraClientId field to given value.


### GetTelstraClientSecret

`func (o *TestTelstra) GetTelstraClientSecret() string`

GetTelstraClientSecret returns the TelstraClientSecret field if non-nil, zero value otherwise.

### GetTelstraClientSecretOk

`func (o *TestTelstra) GetTelstraClientSecretOk() (*string, bool)`

GetTelstraClientSecretOk returns a tuple with the TelstraClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelstraClientSecret

`func (o *TestTelstra) SetTelstraClientSecret(v string)`

SetTelstraClientSecret sets TelstraClientSecret field to given value.


### GetTo

`func (o *TestTelstra) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *TestTelstra) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *TestTelstra) SetTo(v string)`

SetTo sets To field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


