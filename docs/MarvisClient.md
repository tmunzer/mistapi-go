# MarvisClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Diabled** | Pointer to **bool** |  | [optional] [default to false]
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**ProvisionUrl** | Pointer to **string** | in MDM, add &#x60;--provision_url &lt;provision_url&gt;&#x60; to the instlal command | [optional] [readonly] 

## Methods

### NewMarvisClient

`func NewMarvisClient() *MarvisClient`

NewMarvisClient instantiates a new MarvisClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarvisClientWithDefaults

`func NewMarvisClientWithDefaults() *MarvisClient`

NewMarvisClientWithDefaults instantiates a new MarvisClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiabled

`func (o *MarvisClient) GetDiabled() bool`

GetDiabled returns the Diabled field if non-nil, zero value otherwise.

### GetDiabledOk

`func (o *MarvisClient) GetDiabledOk() (*bool, bool)`

GetDiabledOk returns a tuple with the Diabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiabled

`func (o *MarvisClient) SetDiabled(v bool)`

SetDiabled sets Diabled field to given value.

### HasDiabled

`func (o *MarvisClient) HasDiabled() bool`

HasDiabled returns a boolean if a field has been set.

### GetId

`func (o *MarvisClient) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MarvisClient) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MarvisClient) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MarvisClient) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *MarvisClient) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MarvisClient) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MarvisClient) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MarvisClient) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProvisionUrl

`func (o *MarvisClient) GetProvisionUrl() string`

GetProvisionUrl returns the ProvisionUrl field if non-nil, zero value otherwise.

### GetProvisionUrlOk

`func (o *MarvisClient) GetProvisionUrlOk() (*string, bool)`

GetProvisionUrlOk returns a tuple with the ProvisionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionUrl

`func (o *MarvisClient) SetProvisionUrl(v string)`

SetProvisionUrl sets ProvisionUrl field to given value.

### HasProvisionUrl

`func (o *MarvisClient) HasProvisionUrl() bool`

HasProvisionUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


