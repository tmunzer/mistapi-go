# OrgSettingPcap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | Pointer to **string** |  | [optional] 
**MaxPktLen** | Pointer to **int32** | max_len of non-management packets to capture | [optional] [default to 128]

## Methods

### NewOrgSettingPcap

`func NewOrgSettingPcap() *OrgSettingPcap`

NewOrgSettingPcap instantiates a new OrgSettingPcap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgSettingPcapWithDefaults

`func NewOrgSettingPcapWithDefaults() *OrgSettingPcap`

NewOrgSettingPcapWithDefaults instantiates a new OrgSettingPcap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *OrgSettingPcap) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *OrgSettingPcap) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *OrgSettingPcap) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *OrgSettingPcap) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetMaxPktLen

`func (o *OrgSettingPcap) GetMaxPktLen() int32`

GetMaxPktLen returns the MaxPktLen field if non-nil, zero value otherwise.

### GetMaxPktLenOk

`func (o *OrgSettingPcap) GetMaxPktLenOk() (*int32, bool)`

GetMaxPktLenOk returns a tuple with the MaxPktLen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPktLen

`func (o *OrgSettingPcap) SetMaxPktLen(v int32)`

SetMaxPktLen sets MaxPktLen field to given value.

### HasMaxPktLen

`func (o *OrgSettingPcap) HasMaxPktLen() bool`

HasMaxPktLen returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


