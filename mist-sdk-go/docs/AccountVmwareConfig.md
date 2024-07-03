# AccountVmwareConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | customer account Client ID | 
**ClientSecret** | **string** | customer account Client Secret | 
**InstanceUrl** | **string** | customer account VMware instance URL | 

## Methods

### NewAccountVmwareConfig

`func NewAccountVmwareConfig(clientId string, clientSecret string, instanceUrl string, ) *AccountVmwareConfig`

NewAccountVmwareConfig instantiates a new AccountVmwareConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountVmwareConfigWithDefaults

`func NewAccountVmwareConfigWithDefaults() *AccountVmwareConfig`

NewAccountVmwareConfigWithDefaults instantiates a new AccountVmwareConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *AccountVmwareConfig) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *AccountVmwareConfig) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *AccountVmwareConfig) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *AccountVmwareConfig) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *AccountVmwareConfig) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *AccountVmwareConfig) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetInstanceUrl

`func (o *AccountVmwareConfig) GetInstanceUrl() string`

GetInstanceUrl returns the InstanceUrl field if non-nil, zero value otherwise.

### GetInstanceUrlOk

`func (o *AccountVmwareConfig) GetInstanceUrlOk() (*string, bool)`

GetInstanceUrlOk returns a tuple with the InstanceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceUrl

`func (o *AccountVmwareConfig) SetInstanceUrl(v string)`

SetInstanceUrl sets InstanceUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


