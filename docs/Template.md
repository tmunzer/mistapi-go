# Template

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Applies** | Pointer to [**TemplateApplies**](TemplateApplies.md) |  | [optional] 
**CreatedTime** | Pointer to **float32** |  | [optional] [readonly] 
**DeviceprofileIds** | Pointer to **[]string** | list of Device Profile ids | [optional] 
**Exceptions** | Pointer to [**TemplateExceptions**](TemplateExceptions.md) |  | [optional] 
**FilterByDeviceprofile** | Pointer to **bool** | whether to further filter by Device Profile | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**ModifiedTime** | Pointer to **float32** |  | [optional] [readonly] 
**Name** | **string** |  | 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewTemplate

`func NewTemplate(name string, ) *Template`

NewTemplate instantiates a new Template object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateWithDefaults

`func NewTemplateWithDefaults() *Template`

NewTemplateWithDefaults instantiates a new Template object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplies

`func (o *Template) GetApplies() TemplateApplies`

GetApplies returns the Applies field if non-nil, zero value otherwise.

### GetAppliesOk

`func (o *Template) GetAppliesOk() (*TemplateApplies, bool)`

GetAppliesOk returns a tuple with the Applies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplies

`func (o *Template) SetApplies(v TemplateApplies)`

SetApplies sets Applies field to given value.

### HasApplies

`func (o *Template) HasApplies() bool`

HasApplies returns a boolean if a field has been set.

### GetCreatedTime

`func (o *Template) GetCreatedTime() float32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Template) GetCreatedTimeOk() (*float32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Template) SetCreatedTime(v float32)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *Template) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDeviceprofileIds

`func (o *Template) GetDeviceprofileIds() []string`

GetDeviceprofileIds returns the DeviceprofileIds field if non-nil, zero value otherwise.

### GetDeviceprofileIdsOk

`func (o *Template) GetDeviceprofileIdsOk() (*[]string, bool)`

GetDeviceprofileIdsOk returns a tuple with the DeviceprofileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceprofileIds

`func (o *Template) SetDeviceprofileIds(v []string)`

SetDeviceprofileIds sets DeviceprofileIds field to given value.

### HasDeviceprofileIds

`func (o *Template) HasDeviceprofileIds() bool`

HasDeviceprofileIds returns a boolean if a field has been set.

### GetExceptions

`func (o *Template) GetExceptions() TemplateExceptions`

GetExceptions returns the Exceptions field if non-nil, zero value otherwise.

### GetExceptionsOk

`func (o *Template) GetExceptionsOk() (*TemplateExceptions, bool)`

GetExceptionsOk returns a tuple with the Exceptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptions

`func (o *Template) SetExceptions(v TemplateExceptions)`

SetExceptions sets Exceptions field to given value.

### HasExceptions

`func (o *Template) HasExceptions() bool`

HasExceptions returns a boolean if a field has been set.

### GetFilterByDeviceprofile

`func (o *Template) GetFilterByDeviceprofile() bool`

GetFilterByDeviceprofile returns the FilterByDeviceprofile field if non-nil, zero value otherwise.

### GetFilterByDeviceprofileOk

`func (o *Template) GetFilterByDeviceprofileOk() (*bool, bool)`

GetFilterByDeviceprofileOk returns a tuple with the FilterByDeviceprofile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterByDeviceprofile

`func (o *Template) SetFilterByDeviceprofile(v bool)`

SetFilterByDeviceprofile sets FilterByDeviceprofile field to given value.

### HasFilterByDeviceprofile

`func (o *Template) HasFilterByDeviceprofile() bool`

HasFilterByDeviceprofile returns a boolean if a field has been set.

### GetId

`func (o *Template) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Template) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Template) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Template) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifiedTime

`func (o *Template) GetModifiedTime() float32`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *Template) GetModifiedTimeOk() (*float32, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *Template) SetModifiedTime(v float32)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *Template) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### GetName

`func (o *Template) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Template) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Template) SetName(v string)`

SetName sets Name field to given value.


### GetOrgId

`func (o *Template) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *Template) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *Template) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *Template) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


