# SnmpConfigView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Include** | Pointer to **bool** | if the root oid configured is included | [optional] 
**Oid** | Pointer to **string** |  | [optional] 
**ViewName** | Pointer to **string** |  | [optional] 

## Methods

### NewSnmpConfigView

`func NewSnmpConfigView() *SnmpConfigView`

NewSnmpConfigView instantiates a new SnmpConfigView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnmpConfigViewWithDefaults

`func NewSnmpConfigViewWithDefaults() *SnmpConfigView`

NewSnmpConfigViewWithDefaults instantiates a new SnmpConfigView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInclude

`func (o *SnmpConfigView) GetInclude() bool`

GetInclude returns the Include field if non-nil, zero value otherwise.

### GetIncludeOk

`func (o *SnmpConfigView) GetIncludeOk() (*bool, bool)`

GetIncludeOk returns a tuple with the Include field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInclude

`func (o *SnmpConfigView) SetInclude(v bool)`

SetInclude sets Include field to given value.

### HasInclude

`func (o *SnmpConfigView) HasInclude() bool`

HasInclude returns a boolean if a field has been set.

### GetOid

`func (o *SnmpConfigView) GetOid() string`

GetOid returns the Oid field if non-nil, zero value otherwise.

### GetOidOk

`func (o *SnmpConfigView) GetOidOk() (*string, bool)`

GetOidOk returns a tuple with the Oid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOid

`func (o *SnmpConfigView) SetOid(v string)`

SetOid sets Oid field to given value.

### HasOid

`func (o *SnmpConfigView) HasOid() bool`

HasOid returns a boolean if a field has been set.

### GetViewName

`func (o *SnmpConfigView) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *SnmpConfigView) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *SnmpConfigView) SetViewName(v string)`

SetViewName sets ViewName field to given value.

### HasViewName

`func (o *SnmpConfigView) HasViewName() bool`

HasViewName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


