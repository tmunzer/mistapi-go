# UserMac

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Mac** | Pointer to **string** | only non-local-admin MAC is accepted | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**Vlan** | Pointer to **int32** |  | [optional] 

## Methods

### NewUserMac

`func NewUserMac() *UserMac`

NewUserMac instantiates a new UserMac object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserMacWithDefaults

`func NewUserMacWithDefaults() *UserMac`

NewUserMacWithDefaults instantiates a new UserMac object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserMac) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserMac) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserMac) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserMac) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabels

`func (o *UserMac) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *UserMac) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *UserMac) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *UserMac) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetMac

`func (o *UserMac) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *UserMac) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *UserMac) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *UserMac) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetNotes

`func (o *UserMac) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *UserMac) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *UserMac) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *UserMac) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetVlan

`func (o *UserMac) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *UserMac) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *UserMac) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *UserMac) HasVlan() bool`

HasVlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


