# SyntheticTestRadiusServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | **string** | Specify the password associated with the username | 
**Profile** | Pointer to **string** | Specify the access profile associated with the subscriber | [optional] [default to "dot1x"]
**User** | **string** | Specify the subscriber username to test | 

## Methods

### NewSyntheticTestRadiusServer

`func NewSyntheticTestRadiusServer(password string, user string, ) *SyntheticTestRadiusServer`

NewSyntheticTestRadiusServer instantiates a new SyntheticTestRadiusServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticTestRadiusServerWithDefaults

`func NewSyntheticTestRadiusServerWithDefaults() *SyntheticTestRadiusServer`

NewSyntheticTestRadiusServerWithDefaults instantiates a new SyntheticTestRadiusServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *SyntheticTestRadiusServer) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SyntheticTestRadiusServer) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SyntheticTestRadiusServer) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetProfile

`func (o *SyntheticTestRadiusServer) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *SyntheticTestRadiusServer) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *SyntheticTestRadiusServer) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *SyntheticTestRadiusServer) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetUser

`func (o *SyntheticTestRadiusServer) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *SyntheticTestRadiusServer) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *SyntheticTestRadiusServer) SetUser(v string)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


