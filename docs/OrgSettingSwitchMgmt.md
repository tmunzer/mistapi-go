# OrgSettingSwitchMgmt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApAffinityThreshold** | Pointer to **int32** | If the field is set in both site/setting and org/setting, the value from site/setting will be used. | [optional] [default to 12]

## Methods

### NewOrgSettingSwitchMgmt

`func NewOrgSettingSwitchMgmt() *OrgSettingSwitchMgmt`

NewOrgSettingSwitchMgmt instantiates a new OrgSettingSwitchMgmt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgSettingSwitchMgmtWithDefaults

`func NewOrgSettingSwitchMgmtWithDefaults() *OrgSettingSwitchMgmt`

NewOrgSettingSwitchMgmtWithDefaults instantiates a new OrgSettingSwitchMgmt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApAffinityThreshold

`func (o *OrgSettingSwitchMgmt) GetApAffinityThreshold() int32`

GetApAffinityThreshold returns the ApAffinityThreshold field if non-nil, zero value otherwise.

### GetApAffinityThresholdOk

`func (o *OrgSettingSwitchMgmt) GetApAffinityThresholdOk() (*int32, bool)`

GetApAffinityThresholdOk returns a tuple with the ApAffinityThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApAffinityThreshold

`func (o *OrgSettingSwitchMgmt) SetApAffinityThreshold(v int32)`

SetApAffinityThreshold sets ApAffinityThreshold field to given value.

### HasApAffinityThreshold

`func (o *OrgSettingSwitchMgmt) HasApAffinityThreshold() bool`

HasApAffinityThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


