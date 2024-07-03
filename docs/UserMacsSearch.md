# UserMacsSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int32** |  | [optional] 
**Page** | Pointer to **int32** |  | [optional] 
**Results** | Pointer to [**[]UserMac**](UserMac.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewUserMacsSearch

`func NewUserMacsSearch() *UserMacsSearch`

NewUserMacsSearch instantiates a new UserMacsSearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserMacsSearchWithDefaults

`func NewUserMacsSearchWithDefaults() *UserMacsSearch`

NewUserMacsSearchWithDefaults instantiates a new UserMacsSearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *UserMacsSearch) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *UserMacsSearch) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *UserMacsSearch) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *UserMacsSearch) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetPage

`func (o *UserMacsSearch) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *UserMacsSearch) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *UserMacsSearch) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *UserMacsSearch) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetResults

`func (o *UserMacsSearch) GetResults() []UserMac`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *UserMacsSearch) GetResultsOk() (*[]UserMac, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *UserMacsSearch) SetResults(v []UserMac)`

SetResults sets Results field to given value.

### HasResults

`func (o *UserMacsSearch) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetTotal

`func (o *UserMacsSearch) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *UserMacsSearch) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *UserMacsSearch) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *UserMacsSearch) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


