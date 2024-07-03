# ResponseLogout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForwardUrl** | Pointer to **string** | if configured in SSO as custom_logout_url | [optional] 

## Methods

### NewResponseLogout

`func NewResponseLogout() *ResponseLogout`

NewResponseLogout instantiates a new ResponseLogout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseLogoutWithDefaults

`func NewResponseLogoutWithDefaults() *ResponseLogout`

NewResponseLogoutWithDefaults instantiates a new ResponseLogout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForwardUrl

`func (o *ResponseLogout) GetForwardUrl() string`

GetForwardUrl returns the ForwardUrl field if non-nil, zero value otherwise.

### GetForwardUrlOk

`func (o *ResponseLogout) GetForwardUrlOk() (*string, bool)`

GetForwardUrlOk returns a tuple with the ForwardUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardUrl

`func (o *ResponseLogout) SetForwardUrl(v string)`

SetForwardUrl sets ForwardUrl field to given value.

### HasForwardUrl

`func (o *ResponseLogout) HasForwardUrl() bool`

HasForwardUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


