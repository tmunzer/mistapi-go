# AppProbingCustomApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | if &#x60;protocol&#x60;&#x3D;&#x3D;&#x60;icmp&#x60; | [optional] 
**AppType** | Pointer to **string** |  | [optional] 
**Hostname** | Pointer to **[]string** | if &#x60;protocol&#x60;&#x3D;&#x3D;&#x60;http&#x60; | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Network** | Pointer to **string** |  | [optional] 
**Protocol** | Pointer to [**AppProbingCustomAppProtocol**](AppProbingCustomAppProtocol.md) |  | [optional] [default to APPPROBINGCUSTOMAPPPROTOCOL_HTTP]
**Url** | Pointer to **string** | if &#x60;protocol&#x60;&#x3D;&#x3D;&#x60;http&#x60; | [optional] 
**Vrf** | Pointer to **string** |  | [optional] 

## Methods

### NewAppProbingCustomApp

`func NewAppProbingCustomApp() *AppProbingCustomApp`

NewAppProbingCustomApp instantiates a new AppProbingCustomApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppProbingCustomAppWithDefaults

`func NewAppProbingCustomAppWithDefaults() *AppProbingCustomApp`

NewAppProbingCustomAppWithDefaults instantiates a new AppProbingCustomApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *AppProbingCustomApp) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AppProbingCustomApp) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AppProbingCustomApp) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *AppProbingCustomApp) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAppType

`func (o *AppProbingCustomApp) GetAppType() string`

GetAppType returns the AppType field if non-nil, zero value otherwise.

### GetAppTypeOk

`func (o *AppProbingCustomApp) GetAppTypeOk() (*string, bool)`

GetAppTypeOk returns a tuple with the AppType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppType

`func (o *AppProbingCustomApp) SetAppType(v string)`

SetAppType sets AppType field to given value.

### HasAppType

`func (o *AppProbingCustomApp) HasAppType() bool`

HasAppType returns a boolean if a field has been set.

### GetHostname

`func (o *AppProbingCustomApp) GetHostname() []string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *AppProbingCustomApp) GetHostnameOk() (*[]string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *AppProbingCustomApp) SetHostname(v []string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *AppProbingCustomApp) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetName

`func (o *AppProbingCustomApp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppProbingCustomApp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppProbingCustomApp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AppProbingCustomApp) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetwork

`func (o *AppProbingCustomApp) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *AppProbingCustomApp) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *AppProbingCustomApp) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *AppProbingCustomApp) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetProtocol

`func (o *AppProbingCustomApp) GetProtocol() AppProbingCustomAppProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *AppProbingCustomApp) GetProtocolOk() (*AppProbingCustomAppProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *AppProbingCustomApp) SetProtocol(v AppProbingCustomAppProtocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *AppProbingCustomApp) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetUrl

`func (o *AppProbingCustomApp) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AppProbingCustomApp) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AppProbingCustomApp) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AppProbingCustomApp) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetVrf

`func (o *AppProbingCustomApp) GetVrf() string`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *AppProbingCustomApp) GetVrfOk() (*string, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *AppProbingCustomApp) SetVrf(v string)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *AppProbingCustomApp) HasVrf() bool`

HasVrf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


