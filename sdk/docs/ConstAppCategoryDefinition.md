# ConstAppCategoryDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Display** | **string** | Description of the app category | 
**Filters** | Pointer to [**ConstAppCategoryDefinitionFilters**](ConstAppCategoryDefinitionFilters.md) |  | [optional] 
**Includes** | Pointer to **[]string** | List of other App Categories contained by this one | [optional] 
**Key** | **string** | Key name of the app category | 

## Methods

### NewConstAppCategoryDefinition

`func NewConstAppCategoryDefinition(display string, key string, ) *ConstAppCategoryDefinition`

NewConstAppCategoryDefinition instantiates a new ConstAppCategoryDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConstAppCategoryDefinitionWithDefaults

`func NewConstAppCategoryDefinitionWithDefaults() *ConstAppCategoryDefinition`

NewConstAppCategoryDefinitionWithDefaults instantiates a new ConstAppCategoryDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplay

`func (o *ConstAppCategoryDefinition) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *ConstAppCategoryDefinition) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *ConstAppCategoryDefinition) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetFilters

`func (o *ConstAppCategoryDefinition) GetFilters() ConstAppCategoryDefinitionFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ConstAppCategoryDefinition) GetFiltersOk() (*ConstAppCategoryDefinitionFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ConstAppCategoryDefinition) SetFilters(v ConstAppCategoryDefinitionFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ConstAppCategoryDefinition) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetIncludes

`func (o *ConstAppCategoryDefinition) GetIncludes() []string`

GetIncludes returns the Includes field if non-nil, zero value otherwise.

### GetIncludesOk

`func (o *ConstAppCategoryDefinition) GetIncludesOk() (*[]string, bool)`

GetIncludesOk returns a tuple with the Includes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludes

`func (o *ConstAppCategoryDefinition) SetIncludes(v []string)`

SetIncludes sets Includes field to given value.

### HasIncludes

`func (o *ConstAppCategoryDefinition) HasIncludes() bool`

HasIncludes returns a boolean if a field has been set.

### GetKey

`func (o *ConstAppCategoryDefinition) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ConstAppCategoryDefinition) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ConstAppCategoryDefinition) SetKey(v string)`

SetKey sets Key field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


