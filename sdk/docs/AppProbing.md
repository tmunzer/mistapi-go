# AppProbing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apps** | Pointer to **[]string** |  | [optional] 
**CustomApps** | Pointer to [**[]AppProbingCustomApp**](AppProbingCustomApp.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewAppProbing

`func NewAppProbing() *AppProbing`

NewAppProbing instantiates a new AppProbing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppProbingWithDefaults

`func NewAppProbingWithDefaults() *AppProbing`

NewAppProbingWithDefaults instantiates a new AppProbing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApps

`func (o *AppProbing) GetApps() []string`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *AppProbing) GetAppsOk() (*[]string, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *AppProbing) SetApps(v []string)`

SetApps sets Apps field to given value.

### HasApps

`func (o *AppProbing) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetCustomApps

`func (o *AppProbing) GetCustomApps() []AppProbingCustomApp`

GetCustomApps returns the CustomApps field if non-nil, zero value otherwise.

### GetCustomAppsOk

`func (o *AppProbing) GetCustomAppsOk() (*[]AppProbingCustomApp, bool)`

GetCustomAppsOk returns a tuple with the CustomApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomApps

`func (o *AppProbing) SetCustomApps(v []AppProbingCustomApp)`

SetCustomApps sets CustomApps field to given value.

### HasCustomApps

`func (o *AppProbing) HasCustomApps() bool`

HasCustomApps returns a boolean if a field has been set.

### GetEnabled

`func (o *AppProbing) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AppProbing) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AppProbing) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AppProbing) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


