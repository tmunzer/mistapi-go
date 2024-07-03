# GatewayComplianceVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MajorVersion** | Pointer to [**map[string]GatewayComplianceMajorVersionProperties**](GatewayComplianceMajorVersionProperties.md) |  | [optional] 
**Score** | Pointer to **float32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewGatewayComplianceVersion

`func NewGatewayComplianceVersion() *GatewayComplianceVersion`

NewGatewayComplianceVersion instantiates a new GatewayComplianceVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayComplianceVersionWithDefaults

`func NewGatewayComplianceVersionWithDefaults() *GatewayComplianceVersion`

NewGatewayComplianceVersionWithDefaults instantiates a new GatewayComplianceVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMajorVersion

`func (o *GatewayComplianceVersion) GetMajorVersion() map[string]GatewayComplianceMajorVersionProperties`

GetMajorVersion returns the MajorVersion field if non-nil, zero value otherwise.

### GetMajorVersionOk

`func (o *GatewayComplianceVersion) GetMajorVersionOk() (*map[string]GatewayComplianceMajorVersionProperties, bool)`

GetMajorVersionOk returns a tuple with the MajorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajorVersion

`func (o *GatewayComplianceVersion) SetMajorVersion(v map[string]GatewayComplianceMajorVersionProperties)`

SetMajorVersion sets MajorVersion field to given value.

### HasMajorVersion

`func (o *GatewayComplianceVersion) HasMajorVersion() bool`

HasMajorVersion returns a boolean if a field has been set.

### GetScore

`func (o *GatewayComplianceVersion) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *GatewayComplianceVersion) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *GatewayComplianceVersion) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *GatewayComplianceVersion) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetType

`func (o *GatewayComplianceVersion) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GatewayComplianceVersion) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GatewayComplianceVersion) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GatewayComplianceVersion) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


