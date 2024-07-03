# RoutingPolicyTerm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to [**RoutingPolicyTermAction**](RoutingPolicyTermAction.md) |  | [optional] 
**Matching** | Pointer to [**RoutingPolicyTermMatching**](RoutingPolicyTermMatching.md) |  | [optional] 

## Methods

### NewRoutingPolicyTerm

`func NewRoutingPolicyTerm() *RoutingPolicyTerm`

NewRoutingPolicyTerm instantiates a new RoutingPolicyTerm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingPolicyTermWithDefaults

`func NewRoutingPolicyTermWithDefaults() *RoutingPolicyTerm`

NewRoutingPolicyTermWithDefaults instantiates a new RoutingPolicyTerm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *RoutingPolicyTerm) GetAction() RoutingPolicyTermAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *RoutingPolicyTerm) GetActionOk() (*RoutingPolicyTermAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *RoutingPolicyTerm) SetAction(v RoutingPolicyTermAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *RoutingPolicyTerm) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetMatching

`func (o *RoutingPolicyTerm) GetMatching() RoutingPolicyTermMatching`

GetMatching returns the Matching field if non-nil, zero value otherwise.

### GetMatchingOk

`func (o *RoutingPolicyTerm) GetMatchingOk() (*RoutingPolicyTermMatching, bool)`

GetMatchingOk returns a tuple with the Matching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatching

`func (o *RoutingPolicyTerm) SetMatching(v RoutingPolicyTermMatching)`

SetMatching sets Matching field to given value.

### HasMatching

`func (o *RoutingPolicyTerm) HasMatching() bool`

HasMatching returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


