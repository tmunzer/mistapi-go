# OspfConfigArea

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NoSummary** | Pointer to **bool** | for a stub/nssa area, where to avoid forwarding type-3 LSA to this area | [optional] 

## Methods

### NewOspfConfigArea

`func NewOspfConfigArea() *OspfConfigArea`

NewOspfConfigArea instantiates a new OspfConfigArea object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOspfConfigAreaWithDefaults

`func NewOspfConfigAreaWithDefaults() *OspfConfigArea`

NewOspfConfigAreaWithDefaults instantiates a new OspfConfigArea object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNoSummary

`func (o *OspfConfigArea) GetNoSummary() bool`

GetNoSummary returns the NoSummary field if non-nil, zero value otherwise.

### GetNoSummaryOk

`func (o *OspfConfigArea) GetNoSummaryOk() (*bool, bool)`

GetNoSummaryOk returns a tuple with the NoSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoSummary

`func (o *OspfConfigArea) SetNoSummary(v bool)`

SetNoSummary sets NoSummary field to given value.

### HasNoSummary

`func (o *OspfConfigArea) HasNoSummary() bool`

HasNoSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


