# SnmpConfigClientList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientListName** | Pointer to **string** |  | [optional] 
**Clients** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSnmpConfigClientList

`func NewSnmpConfigClientList() *SnmpConfigClientList`

NewSnmpConfigClientList instantiates a new SnmpConfigClientList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnmpConfigClientListWithDefaults

`func NewSnmpConfigClientListWithDefaults() *SnmpConfigClientList`

NewSnmpConfigClientListWithDefaults instantiates a new SnmpConfigClientList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientListName

`func (o *SnmpConfigClientList) GetClientListName() string`

GetClientListName returns the ClientListName field if non-nil, zero value otherwise.

### GetClientListNameOk

`func (o *SnmpConfigClientList) GetClientListNameOk() (*string, bool)`

GetClientListNameOk returns a tuple with the ClientListName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientListName

`func (o *SnmpConfigClientList) SetClientListName(v string)`

SetClientListName sets ClientListName field to given value.

### HasClientListName

`func (o *SnmpConfigClientList) HasClientListName() bool`

HasClientListName returns a boolean if a field has been set.

### GetClients

`func (o *SnmpConfigClientList) GetClients() []string`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *SnmpConfigClientList) GetClientsOk() (*[]string, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *SnmpConfigClientList) SetClients(v []string)`

SetClients sets Clients field to given value.

### HasClients

`func (o *SnmpConfigClientList) HasClients() bool`

HasClients returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


