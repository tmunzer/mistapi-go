# PmaDashboard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | description of the dashboard | [optional] 
**Label** | Pointer to **string** | group label name | [optional] 
**Name** | Pointer to **string** | name of the dashboard | [optional] 
**Url** | Pointer to **string** | url to access dashboard. Url will redirect the user to the dashboard | [optional] 

## Methods

### NewPmaDashboard

`func NewPmaDashboard() *PmaDashboard`

NewPmaDashboard instantiates a new PmaDashboard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPmaDashboardWithDefaults

`func NewPmaDashboardWithDefaults() *PmaDashboard`

NewPmaDashboardWithDefaults instantiates a new PmaDashboard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *PmaDashboard) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PmaDashboard) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PmaDashboard) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PmaDashboard) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabel

`func (o *PmaDashboard) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PmaDashboard) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PmaDashboard) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PmaDashboard) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *PmaDashboard) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PmaDashboard) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PmaDashboard) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PmaDashboard) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *PmaDashboard) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PmaDashboard) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PmaDashboard) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PmaDashboard) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


