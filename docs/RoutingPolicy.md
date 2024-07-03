# RoutingPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Terms** | Pointer to [**[]RoutingPolicyTerm**](RoutingPolicyTerm.md) | zero or more criteria/filter can be specified to match the term, all criteria have to be met | [optional] 

## Methods

### NewRoutingPolicy

`func NewRoutingPolicy() *RoutingPolicy`

NewRoutingPolicy instantiates a new RoutingPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingPolicyWithDefaults

`func NewRoutingPolicyWithDefaults() *RoutingPolicy`

NewRoutingPolicyWithDefaults instantiates a new RoutingPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTerms

`func (o *RoutingPolicy) GetTerms() []RoutingPolicyTerm`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *RoutingPolicy) GetTermsOk() (*[]RoutingPolicyTerm, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *RoutingPolicy) SetTerms(v []RoutingPolicyTerm)`

SetTerms sets Terms field to given value.

### HasTerms

`func (o *RoutingPolicy) HasTerms() bool`

HasTerms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


