# MxclusterNacClientIp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Secret** | Pointer to **string** | if different from above | [optional] 
**SiteId** | Pointer to **string** | present only for 3rd party clients | [optional] 
**Vendor** | Pointer to [**MxclusterNacClientVendor**](MxclusterNacClientVendor.md) |  | [optional] 

## Methods

### NewMxclusterNacClientIp

`func NewMxclusterNacClientIp() *MxclusterNacClientIp`

NewMxclusterNacClientIp instantiates a new MxclusterNacClientIp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMxclusterNacClientIpWithDefaults

`func NewMxclusterNacClientIpWithDefaults() *MxclusterNacClientIp`

NewMxclusterNacClientIpWithDefaults instantiates a new MxclusterNacClientIp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecret

`func (o *MxclusterNacClientIp) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *MxclusterNacClientIp) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *MxclusterNacClientIp) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *MxclusterNacClientIp) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetSiteId

`func (o *MxclusterNacClientIp) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *MxclusterNacClientIp) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *MxclusterNacClientIp) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *MxclusterNacClientIp) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetVendor

`func (o *MxclusterNacClientIp) GetVendor() MxclusterNacClientVendor`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *MxclusterNacClientIp) GetVendorOk() (*MxclusterNacClientVendor, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *MxclusterNacClientIp) SetVendor(v MxclusterNacClientVendor)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *MxclusterNacClientIp) HasVendor() bool`

HasVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


