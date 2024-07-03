# NacCrlFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | Pointer to **int32** | Time at which the file was first uploaded, in epoch | [optional] 
**Id** | Pointer to **string** | ID for file upload, used to identify file even for deletion | [optional] 
**ModifiedTime** | Pointer to **int32** | Time at which the file was last updated, in epoch | [optional] 
**Name** | Pointer to **string** | Name for the .crl file issuer if set | [optional] 
**Url** | Pointer to **string** | Download URL for the .crl file | [optional] 

## Methods

### NewNacCrlFile

`func NewNacCrlFile() *NacCrlFile`

NewNacCrlFile instantiates a new NacCrlFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNacCrlFileWithDefaults

`func NewNacCrlFileWithDefaults() *NacCrlFile`

NewNacCrlFileWithDefaults instantiates a new NacCrlFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTime

`func (o *NacCrlFile) GetCreatedTime() int32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *NacCrlFile) GetCreatedTimeOk() (*int32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *NacCrlFile) SetCreatedTime(v int32)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *NacCrlFile) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetId

`func (o *NacCrlFile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NacCrlFile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NacCrlFile) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NacCrlFile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifiedTime

`func (o *NacCrlFile) GetModifiedTime() int32`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *NacCrlFile) GetModifiedTimeOk() (*int32, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *NacCrlFile) SetModifiedTime(v int32)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *NacCrlFile) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### GetName

`func (o *NacCrlFile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NacCrlFile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NacCrlFile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NacCrlFile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *NacCrlFile) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NacCrlFile) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NacCrlFile) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *NacCrlFile) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


