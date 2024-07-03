# Snmpv3ConfigNotifyFilterItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contents** | Pointer to [**[]Snmpv3ConfigNotifyFilterItemContent**](Snmpv3ConfigNotifyFilterItemContent.md) |  | [optional] 
**ProfileName** | Pointer to **string** |  | [optional] 

## Methods

### NewSnmpv3ConfigNotifyFilterItem

`func NewSnmpv3ConfigNotifyFilterItem() *Snmpv3ConfigNotifyFilterItem`

NewSnmpv3ConfigNotifyFilterItem instantiates a new Snmpv3ConfigNotifyFilterItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnmpv3ConfigNotifyFilterItemWithDefaults

`func NewSnmpv3ConfigNotifyFilterItemWithDefaults() *Snmpv3ConfigNotifyFilterItem`

NewSnmpv3ConfigNotifyFilterItemWithDefaults instantiates a new Snmpv3ConfigNotifyFilterItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContents

`func (o *Snmpv3ConfigNotifyFilterItem) GetContents() []Snmpv3ConfigNotifyFilterItemContent`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *Snmpv3ConfigNotifyFilterItem) GetContentsOk() (*[]Snmpv3ConfigNotifyFilterItemContent, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *Snmpv3ConfigNotifyFilterItem) SetContents(v []Snmpv3ConfigNotifyFilterItemContent)`

SetContents sets Contents field to given value.

### HasContents

`func (o *Snmpv3ConfigNotifyFilterItem) HasContents() bool`

HasContents returns a boolean if a field has been set.

### GetProfileName

`func (o *Snmpv3ConfigNotifyFilterItem) GetProfileName() string`

GetProfileName returns the ProfileName field if non-nil, zero value otherwise.

### GetProfileNameOk

`func (o *Snmpv3ConfigNotifyFilterItem) GetProfileNameOk() (*string, bool)`

GetProfileNameOk returns a tuple with the ProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileName

`func (o *Snmpv3ConfigNotifyFilterItem) SetProfileName(v string)`

SetProfileName sets ProfileName field to given value.

### HasProfileName

`func (o *Snmpv3ConfigNotifyFilterItem) HasProfileName() bool`

HasProfileName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


