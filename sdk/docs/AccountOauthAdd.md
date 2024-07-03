# AccountOauthAdd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | customer account Client ID | 
**ClientSecret** | **string** | customer account Client Secret | 
**InstanceUrl** | **string** | customer account VMware instance URL | 
**SmartgroupName** | **string** | smart group membership for determining compliance status | 

## Methods

### NewAccountOauthAdd

`func NewAccountOauthAdd(clientId string, clientSecret string, instanceUrl string, smartgroupName string, ) *AccountOauthAdd`

NewAccountOauthAdd instantiates a new AccountOauthAdd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountOauthAddWithDefaults

`func NewAccountOauthAddWithDefaults() *AccountOauthAdd`

NewAccountOauthAddWithDefaults instantiates a new AccountOauthAdd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *AccountOauthAdd) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *AccountOauthAdd) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *AccountOauthAdd) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *AccountOauthAdd) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *AccountOauthAdd) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *AccountOauthAdd) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetInstanceUrl

`func (o *AccountOauthAdd) GetInstanceUrl() string`

GetInstanceUrl returns the InstanceUrl field if non-nil, zero value otherwise.

### GetInstanceUrlOk

`func (o *AccountOauthAdd) GetInstanceUrlOk() (*string, bool)`

GetInstanceUrlOk returns a tuple with the InstanceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceUrl

`func (o *AccountOauthAdd) SetInstanceUrl(v string)`

SetInstanceUrl sets InstanceUrl field to given value.


### GetSmartgroupName

`func (o *AccountOauthAdd) GetSmartgroupName() string`

GetSmartgroupName returns the SmartgroupName field if non-nil, zero value otherwise.

### GetSmartgroupNameOk

`func (o *AccountOauthAdd) GetSmartgroupNameOk() (*string, bool)`

GetSmartgroupNameOk returns a tuple with the SmartgroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartgroupName

`func (o *AccountOauthAdd) SetSmartgroupName(v string)`

SetSmartgroupName sets SmartgroupName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


