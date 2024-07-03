# ResponseOrgSearchItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MspId** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** | org name | [optional] [readonly] 
**NumAps** | Pointer to **int32** |  | [optional] [readonly] 
**NumGateways** | Pointer to **int32** |  | [optional] [readonly] 
**NumSites** | Pointer to **int32** |  | [optional] [readonly] 
**NumSwitches** | Pointer to **int32** |  | [optional] [readonly] 
**NumUnassignedAps** | Pointer to **int32** |  | [optional] [readonly] 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**SubAnaEntitled** | Pointer to **int32** |  | [optional] [readonly] 
**SubAnaRequired** | Pointer to **int32** |  | [optional] [readonly] 
**SubAstEntitled** | Pointer to **int32** |  | [optional] [readonly] 
**SubAstRequired** | Pointer to **int32** |  | [optional] [readonly] 
**SubEngEntitled** | Pointer to **int32** |  | [optional] [readonly] 
**SubEngRequired** | Pointer to **int32** |  | [optional] [readonly] 
**SubEx12Required** | Pointer to **int32** |  | [optional] [readonly] 
**SubInsufficient** | Pointer to **bool** | if this org has sufficient subscription | [optional] [readonly] 
**SubManEntitled** | Pointer to **int32** |  | [optional] [readonly] 
**SubManRequired** | Pointer to **int32** |  | [optional] [readonly] 
**SubMeEntitled** | Pointer to **int32** |  | [optional] [readonly] 
**SubVnaEntitled** | Pointer to **int32** |  | [optional] [readonly] 
**SubVnaRequired** | Pointer to **int32** |  | [optional] [readonly] 
**Timestamp** | Pointer to **float32** |  | [optional] [readonly] 
**TrialEnabled** | Pointer to **bool** | if this org is under trial period | [optional] [readonly] 
**UsageTypes** | Pointer to **[]string** | a list of types that enabled by usage | [optional] [readonly] 

## Methods

### NewResponseOrgSearchItem

`func NewResponseOrgSearchItem() *ResponseOrgSearchItem`

NewResponseOrgSearchItem instantiates a new ResponseOrgSearchItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseOrgSearchItemWithDefaults

`func NewResponseOrgSearchItemWithDefaults() *ResponseOrgSearchItem`

NewResponseOrgSearchItemWithDefaults instantiates a new ResponseOrgSearchItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMspId

`func (o *ResponseOrgSearchItem) GetMspId() string`

GetMspId returns the MspId field if non-nil, zero value otherwise.

### GetMspIdOk

`func (o *ResponseOrgSearchItem) GetMspIdOk() (*string, bool)`

GetMspIdOk returns a tuple with the MspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMspId

`func (o *ResponseOrgSearchItem) SetMspId(v string)`

SetMspId sets MspId field to given value.

### HasMspId

`func (o *ResponseOrgSearchItem) HasMspId() bool`

HasMspId returns a boolean if a field has been set.

### GetName

`func (o *ResponseOrgSearchItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponseOrgSearchItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponseOrgSearchItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResponseOrgSearchItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumAps

`func (o *ResponseOrgSearchItem) GetNumAps() int32`

GetNumAps returns the NumAps field if non-nil, zero value otherwise.

### GetNumApsOk

`func (o *ResponseOrgSearchItem) GetNumApsOk() (*int32, bool)`

GetNumApsOk returns a tuple with the NumAps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAps

`func (o *ResponseOrgSearchItem) SetNumAps(v int32)`

SetNumAps sets NumAps field to given value.

### HasNumAps

`func (o *ResponseOrgSearchItem) HasNumAps() bool`

HasNumAps returns a boolean if a field has been set.

### GetNumGateways

`func (o *ResponseOrgSearchItem) GetNumGateways() int32`

GetNumGateways returns the NumGateways field if non-nil, zero value otherwise.

### GetNumGatewaysOk

`func (o *ResponseOrgSearchItem) GetNumGatewaysOk() (*int32, bool)`

GetNumGatewaysOk returns a tuple with the NumGateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumGateways

`func (o *ResponseOrgSearchItem) SetNumGateways(v int32)`

SetNumGateways sets NumGateways field to given value.

### HasNumGateways

`func (o *ResponseOrgSearchItem) HasNumGateways() bool`

HasNumGateways returns a boolean if a field has been set.

### GetNumSites

`func (o *ResponseOrgSearchItem) GetNumSites() int32`

GetNumSites returns the NumSites field if non-nil, zero value otherwise.

### GetNumSitesOk

`func (o *ResponseOrgSearchItem) GetNumSitesOk() (*int32, bool)`

GetNumSitesOk returns a tuple with the NumSites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSites

`func (o *ResponseOrgSearchItem) SetNumSites(v int32)`

SetNumSites sets NumSites field to given value.

### HasNumSites

`func (o *ResponseOrgSearchItem) HasNumSites() bool`

HasNumSites returns a boolean if a field has been set.

### GetNumSwitches

`func (o *ResponseOrgSearchItem) GetNumSwitches() int32`

GetNumSwitches returns the NumSwitches field if non-nil, zero value otherwise.

### GetNumSwitchesOk

`func (o *ResponseOrgSearchItem) GetNumSwitchesOk() (*int32, bool)`

GetNumSwitchesOk returns a tuple with the NumSwitches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSwitches

`func (o *ResponseOrgSearchItem) SetNumSwitches(v int32)`

SetNumSwitches sets NumSwitches field to given value.

### HasNumSwitches

`func (o *ResponseOrgSearchItem) HasNumSwitches() bool`

HasNumSwitches returns a boolean if a field has been set.

### GetNumUnassignedAps

`func (o *ResponseOrgSearchItem) GetNumUnassignedAps() int32`

GetNumUnassignedAps returns the NumUnassignedAps field if non-nil, zero value otherwise.

### GetNumUnassignedApsOk

`func (o *ResponseOrgSearchItem) GetNumUnassignedApsOk() (*int32, bool)`

GetNumUnassignedApsOk returns a tuple with the NumUnassignedAps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumUnassignedAps

`func (o *ResponseOrgSearchItem) SetNumUnassignedAps(v int32)`

SetNumUnassignedAps sets NumUnassignedAps field to given value.

### HasNumUnassignedAps

`func (o *ResponseOrgSearchItem) HasNumUnassignedAps() bool`

HasNumUnassignedAps returns a boolean if a field has been set.

### GetOrgId

`func (o *ResponseOrgSearchItem) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ResponseOrgSearchItem) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ResponseOrgSearchItem) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ResponseOrgSearchItem) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetSubAnaEntitled

`func (o *ResponseOrgSearchItem) GetSubAnaEntitled() int32`

GetSubAnaEntitled returns the SubAnaEntitled field if non-nil, zero value otherwise.

### GetSubAnaEntitledOk

`func (o *ResponseOrgSearchItem) GetSubAnaEntitledOk() (*int32, bool)`

GetSubAnaEntitledOk returns a tuple with the SubAnaEntitled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAnaEntitled

`func (o *ResponseOrgSearchItem) SetSubAnaEntitled(v int32)`

SetSubAnaEntitled sets SubAnaEntitled field to given value.

### HasSubAnaEntitled

`func (o *ResponseOrgSearchItem) HasSubAnaEntitled() bool`

HasSubAnaEntitled returns a boolean if a field has been set.

### GetSubAnaRequired

`func (o *ResponseOrgSearchItem) GetSubAnaRequired() int32`

GetSubAnaRequired returns the SubAnaRequired field if non-nil, zero value otherwise.

### GetSubAnaRequiredOk

`func (o *ResponseOrgSearchItem) GetSubAnaRequiredOk() (*int32, bool)`

GetSubAnaRequiredOk returns a tuple with the SubAnaRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAnaRequired

`func (o *ResponseOrgSearchItem) SetSubAnaRequired(v int32)`

SetSubAnaRequired sets SubAnaRequired field to given value.

### HasSubAnaRequired

`func (o *ResponseOrgSearchItem) HasSubAnaRequired() bool`

HasSubAnaRequired returns a boolean if a field has been set.

### GetSubAstEntitled

`func (o *ResponseOrgSearchItem) GetSubAstEntitled() int32`

GetSubAstEntitled returns the SubAstEntitled field if non-nil, zero value otherwise.

### GetSubAstEntitledOk

`func (o *ResponseOrgSearchItem) GetSubAstEntitledOk() (*int32, bool)`

GetSubAstEntitledOk returns a tuple with the SubAstEntitled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAstEntitled

`func (o *ResponseOrgSearchItem) SetSubAstEntitled(v int32)`

SetSubAstEntitled sets SubAstEntitled field to given value.

### HasSubAstEntitled

`func (o *ResponseOrgSearchItem) HasSubAstEntitled() bool`

HasSubAstEntitled returns a boolean if a field has been set.

### GetSubAstRequired

`func (o *ResponseOrgSearchItem) GetSubAstRequired() int32`

GetSubAstRequired returns the SubAstRequired field if non-nil, zero value otherwise.

### GetSubAstRequiredOk

`func (o *ResponseOrgSearchItem) GetSubAstRequiredOk() (*int32, bool)`

GetSubAstRequiredOk returns a tuple with the SubAstRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAstRequired

`func (o *ResponseOrgSearchItem) SetSubAstRequired(v int32)`

SetSubAstRequired sets SubAstRequired field to given value.

### HasSubAstRequired

`func (o *ResponseOrgSearchItem) HasSubAstRequired() bool`

HasSubAstRequired returns a boolean if a field has been set.

### GetSubEngEntitled

`func (o *ResponseOrgSearchItem) GetSubEngEntitled() int32`

GetSubEngEntitled returns the SubEngEntitled field if non-nil, zero value otherwise.

### GetSubEngEntitledOk

`func (o *ResponseOrgSearchItem) GetSubEngEntitledOk() (*int32, bool)`

GetSubEngEntitledOk returns a tuple with the SubEngEntitled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubEngEntitled

`func (o *ResponseOrgSearchItem) SetSubEngEntitled(v int32)`

SetSubEngEntitled sets SubEngEntitled field to given value.

### HasSubEngEntitled

`func (o *ResponseOrgSearchItem) HasSubEngEntitled() bool`

HasSubEngEntitled returns a boolean if a field has been set.

### GetSubEngRequired

`func (o *ResponseOrgSearchItem) GetSubEngRequired() int32`

GetSubEngRequired returns the SubEngRequired field if non-nil, zero value otherwise.

### GetSubEngRequiredOk

`func (o *ResponseOrgSearchItem) GetSubEngRequiredOk() (*int32, bool)`

GetSubEngRequiredOk returns a tuple with the SubEngRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubEngRequired

`func (o *ResponseOrgSearchItem) SetSubEngRequired(v int32)`

SetSubEngRequired sets SubEngRequired field to given value.

### HasSubEngRequired

`func (o *ResponseOrgSearchItem) HasSubEngRequired() bool`

HasSubEngRequired returns a boolean if a field has been set.

### GetSubEx12Required

`func (o *ResponseOrgSearchItem) GetSubEx12Required() int32`

GetSubEx12Required returns the SubEx12Required field if non-nil, zero value otherwise.

### GetSubEx12RequiredOk

`func (o *ResponseOrgSearchItem) GetSubEx12RequiredOk() (*int32, bool)`

GetSubEx12RequiredOk returns a tuple with the SubEx12Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubEx12Required

`func (o *ResponseOrgSearchItem) SetSubEx12Required(v int32)`

SetSubEx12Required sets SubEx12Required field to given value.

### HasSubEx12Required

`func (o *ResponseOrgSearchItem) HasSubEx12Required() bool`

HasSubEx12Required returns a boolean if a field has been set.

### GetSubInsufficient

`func (o *ResponseOrgSearchItem) GetSubInsufficient() bool`

GetSubInsufficient returns the SubInsufficient field if non-nil, zero value otherwise.

### GetSubInsufficientOk

`func (o *ResponseOrgSearchItem) GetSubInsufficientOk() (*bool, bool)`

GetSubInsufficientOk returns a tuple with the SubInsufficient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubInsufficient

`func (o *ResponseOrgSearchItem) SetSubInsufficient(v bool)`

SetSubInsufficient sets SubInsufficient field to given value.

### HasSubInsufficient

`func (o *ResponseOrgSearchItem) HasSubInsufficient() bool`

HasSubInsufficient returns a boolean if a field has been set.

### GetSubManEntitled

`func (o *ResponseOrgSearchItem) GetSubManEntitled() int32`

GetSubManEntitled returns the SubManEntitled field if non-nil, zero value otherwise.

### GetSubManEntitledOk

`func (o *ResponseOrgSearchItem) GetSubManEntitledOk() (*int32, bool)`

GetSubManEntitledOk returns a tuple with the SubManEntitled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubManEntitled

`func (o *ResponseOrgSearchItem) SetSubManEntitled(v int32)`

SetSubManEntitled sets SubManEntitled field to given value.

### HasSubManEntitled

`func (o *ResponseOrgSearchItem) HasSubManEntitled() bool`

HasSubManEntitled returns a boolean if a field has been set.

### GetSubManRequired

`func (o *ResponseOrgSearchItem) GetSubManRequired() int32`

GetSubManRequired returns the SubManRequired field if non-nil, zero value otherwise.

### GetSubManRequiredOk

`func (o *ResponseOrgSearchItem) GetSubManRequiredOk() (*int32, bool)`

GetSubManRequiredOk returns a tuple with the SubManRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubManRequired

`func (o *ResponseOrgSearchItem) SetSubManRequired(v int32)`

SetSubManRequired sets SubManRequired field to given value.

### HasSubManRequired

`func (o *ResponseOrgSearchItem) HasSubManRequired() bool`

HasSubManRequired returns a boolean if a field has been set.

### GetSubMeEntitled

`func (o *ResponseOrgSearchItem) GetSubMeEntitled() int32`

GetSubMeEntitled returns the SubMeEntitled field if non-nil, zero value otherwise.

### GetSubMeEntitledOk

`func (o *ResponseOrgSearchItem) GetSubMeEntitledOk() (*int32, bool)`

GetSubMeEntitledOk returns a tuple with the SubMeEntitled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubMeEntitled

`func (o *ResponseOrgSearchItem) SetSubMeEntitled(v int32)`

SetSubMeEntitled sets SubMeEntitled field to given value.

### HasSubMeEntitled

`func (o *ResponseOrgSearchItem) HasSubMeEntitled() bool`

HasSubMeEntitled returns a boolean if a field has been set.

### GetSubVnaEntitled

`func (o *ResponseOrgSearchItem) GetSubVnaEntitled() int32`

GetSubVnaEntitled returns the SubVnaEntitled field if non-nil, zero value otherwise.

### GetSubVnaEntitledOk

`func (o *ResponseOrgSearchItem) GetSubVnaEntitledOk() (*int32, bool)`

GetSubVnaEntitledOk returns a tuple with the SubVnaEntitled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubVnaEntitled

`func (o *ResponseOrgSearchItem) SetSubVnaEntitled(v int32)`

SetSubVnaEntitled sets SubVnaEntitled field to given value.

### HasSubVnaEntitled

`func (o *ResponseOrgSearchItem) HasSubVnaEntitled() bool`

HasSubVnaEntitled returns a boolean if a field has been set.

### GetSubVnaRequired

`func (o *ResponseOrgSearchItem) GetSubVnaRequired() int32`

GetSubVnaRequired returns the SubVnaRequired field if non-nil, zero value otherwise.

### GetSubVnaRequiredOk

`func (o *ResponseOrgSearchItem) GetSubVnaRequiredOk() (*int32, bool)`

GetSubVnaRequiredOk returns a tuple with the SubVnaRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubVnaRequired

`func (o *ResponseOrgSearchItem) SetSubVnaRequired(v int32)`

SetSubVnaRequired sets SubVnaRequired field to given value.

### HasSubVnaRequired

`func (o *ResponseOrgSearchItem) HasSubVnaRequired() bool`

HasSubVnaRequired returns a boolean if a field has been set.

### GetTimestamp

`func (o *ResponseOrgSearchItem) GetTimestamp() float32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ResponseOrgSearchItem) GetTimestampOk() (*float32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ResponseOrgSearchItem) SetTimestamp(v float32)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ResponseOrgSearchItem) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTrialEnabled

`func (o *ResponseOrgSearchItem) GetTrialEnabled() bool`

GetTrialEnabled returns the TrialEnabled field if non-nil, zero value otherwise.

### GetTrialEnabledOk

`func (o *ResponseOrgSearchItem) GetTrialEnabledOk() (*bool, bool)`

GetTrialEnabledOk returns a tuple with the TrialEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialEnabled

`func (o *ResponseOrgSearchItem) SetTrialEnabled(v bool)`

SetTrialEnabled sets TrialEnabled field to given value.

### HasTrialEnabled

`func (o *ResponseOrgSearchItem) HasTrialEnabled() bool`

HasTrialEnabled returns a boolean if a field has been set.

### GetUsageTypes

`func (o *ResponseOrgSearchItem) GetUsageTypes() []string`

GetUsageTypes returns the UsageTypes field if non-nil, zero value otherwise.

### GetUsageTypesOk

`func (o *ResponseOrgSearchItem) GetUsageTypesOk() (*[]string, bool)`

GetUsageTypesOk returns a tuple with the UsageTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageTypes

`func (o *ResponseOrgSearchItem) SetUsageTypes(v []string)`

SetUsageTypes sets UsageTypes field to given value.

### HasUsageTypes

`func (o *ResponseOrgSearchItem) HasUsageTypes() bool`

HasUsageTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


