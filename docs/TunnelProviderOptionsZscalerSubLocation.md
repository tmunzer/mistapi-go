# TunnelProviderOptionsZscalerSubLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AupAcceptanceRequired** | Pointer to **bool** |  | [optional] [default to true]
**AupExpire** | Pointer to **int32** | days before AUP is requested again | [optional] [default to 1]
**AupSslProxy** | Pointer to **bool** | proxy HTTPs traffic, requiring Zscaler cert to be installed in browser | [optional] [default to false]
**DownloadMbps** | Pointer to **int32** | the download bandwidth cap of the link, in Mbps | [optional] 
**EnableAup** | Pointer to **bool** | if &#x60;use_xff&#x60;&#x3D;&#x3D;&#x60;true&#x60;, display Acceptable Use Policy (AUP) | [optional] 
**EnableCaution** | Pointer to **bool** | when &#x60;enforce_authentication&#x60;&#x3D;&#x3D;&#x60;false&#x60;, display caution notification for non-authenticated users | [optional] [default to false]
**EnforceAuthentication** | Pointer to **bool** |  | [optional] [default to false]
**Subnets** | Pointer to **[]string** |  | [optional] 
**UploadMbps** | Pointer to **int32** | the download bandwidth cap of the link, in Mbps | [optional] 

## Methods

### NewTunnelProviderOptionsZscalerSubLocation

`func NewTunnelProviderOptionsZscalerSubLocation() *TunnelProviderOptionsZscalerSubLocation`

NewTunnelProviderOptionsZscalerSubLocation instantiates a new TunnelProviderOptionsZscalerSubLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTunnelProviderOptionsZscalerSubLocationWithDefaults

`func NewTunnelProviderOptionsZscalerSubLocationWithDefaults() *TunnelProviderOptionsZscalerSubLocation`

NewTunnelProviderOptionsZscalerSubLocationWithDefaults instantiates a new TunnelProviderOptionsZscalerSubLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAupAcceptanceRequired

`func (o *TunnelProviderOptionsZscalerSubLocation) GetAupAcceptanceRequired() bool`

GetAupAcceptanceRequired returns the AupAcceptanceRequired field if non-nil, zero value otherwise.

### GetAupAcceptanceRequiredOk

`func (o *TunnelProviderOptionsZscalerSubLocation) GetAupAcceptanceRequiredOk() (*bool, bool)`

GetAupAcceptanceRequiredOk returns a tuple with the AupAcceptanceRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAupAcceptanceRequired

`func (o *TunnelProviderOptionsZscalerSubLocation) SetAupAcceptanceRequired(v bool)`

SetAupAcceptanceRequired sets AupAcceptanceRequired field to given value.

### HasAupAcceptanceRequired

`func (o *TunnelProviderOptionsZscalerSubLocation) HasAupAcceptanceRequired() bool`

HasAupAcceptanceRequired returns a boolean if a field has been set.

### GetAupExpire

`func (o *TunnelProviderOptionsZscalerSubLocation) GetAupExpire() int32`

GetAupExpire returns the AupExpire field if non-nil, zero value otherwise.

### GetAupExpireOk

`func (o *TunnelProviderOptionsZscalerSubLocation) GetAupExpireOk() (*int32, bool)`

GetAupExpireOk returns a tuple with the AupExpire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAupExpire

`func (o *TunnelProviderOptionsZscalerSubLocation) SetAupExpire(v int32)`

SetAupExpire sets AupExpire field to given value.

### HasAupExpire

`func (o *TunnelProviderOptionsZscalerSubLocation) HasAupExpire() bool`

HasAupExpire returns a boolean if a field has been set.

### GetAupSslProxy

`func (o *TunnelProviderOptionsZscalerSubLocation) GetAupSslProxy() bool`

GetAupSslProxy returns the AupSslProxy field if non-nil, zero value otherwise.

### GetAupSslProxyOk

`func (o *TunnelProviderOptionsZscalerSubLocation) GetAupSslProxyOk() (*bool, bool)`

GetAupSslProxyOk returns a tuple with the AupSslProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAupSslProxy

`func (o *TunnelProviderOptionsZscalerSubLocation) SetAupSslProxy(v bool)`

SetAupSslProxy sets AupSslProxy field to given value.

### HasAupSslProxy

`func (o *TunnelProviderOptionsZscalerSubLocation) HasAupSslProxy() bool`

HasAupSslProxy returns a boolean if a field has been set.

### GetDownloadMbps

`func (o *TunnelProviderOptionsZscalerSubLocation) GetDownloadMbps() int32`

GetDownloadMbps returns the DownloadMbps field if non-nil, zero value otherwise.

### GetDownloadMbpsOk

`func (o *TunnelProviderOptionsZscalerSubLocation) GetDownloadMbpsOk() (*int32, bool)`

GetDownloadMbpsOk returns a tuple with the DownloadMbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadMbps

`func (o *TunnelProviderOptionsZscalerSubLocation) SetDownloadMbps(v int32)`

SetDownloadMbps sets DownloadMbps field to given value.

### HasDownloadMbps

`func (o *TunnelProviderOptionsZscalerSubLocation) HasDownloadMbps() bool`

HasDownloadMbps returns a boolean if a field has been set.

### GetEnableAup

`func (o *TunnelProviderOptionsZscalerSubLocation) GetEnableAup() bool`

GetEnableAup returns the EnableAup field if non-nil, zero value otherwise.

### GetEnableAupOk

`func (o *TunnelProviderOptionsZscalerSubLocation) GetEnableAupOk() (*bool, bool)`

GetEnableAupOk returns a tuple with the EnableAup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAup

`func (o *TunnelProviderOptionsZscalerSubLocation) SetEnableAup(v bool)`

SetEnableAup sets EnableAup field to given value.

### HasEnableAup

`func (o *TunnelProviderOptionsZscalerSubLocation) HasEnableAup() bool`

HasEnableAup returns a boolean if a field has been set.

### GetEnableCaution

`func (o *TunnelProviderOptionsZscalerSubLocation) GetEnableCaution() bool`

GetEnableCaution returns the EnableCaution field if non-nil, zero value otherwise.

### GetEnableCautionOk

`func (o *TunnelProviderOptionsZscalerSubLocation) GetEnableCautionOk() (*bool, bool)`

GetEnableCautionOk returns a tuple with the EnableCaution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCaution

`func (o *TunnelProviderOptionsZscalerSubLocation) SetEnableCaution(v bool)`

SetEnableCaution sets EnableCaution field to given value.

### HasEnableCaution

`func (o *TunnelProviderOptionsZscalerSubLocation) HasEnableCaution() bool`

HasEnableCaution returns a boolean if a field has been set.

### GetEnforceAuthentication

`func (o *TunnelProviderOptionsZscalerSubLocation) GetEnforceAuthentication() bool`

GetEnforceAuthentication returns the EnforceAuthentication field if non-nil, zero value otherwise.

### GetEnforceAuthenticationOk

`func (o *TunnelProviderOptionsZscalerSubLocation) GetEnforceAuthenticationOk() (*bool, bool)`

GetEnforceAuthenticationOk returns a tuple with the EnforceAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceAuthentication

`func (o *TunnelProviderOptionsZscalerSubLocation) SetEnforceAuthentication(v bool)`

SetEnforceAuthentication sets EnforceAuthentication field to given value.

### HasEnforceAuthentication

`func (o *TunnelProviderOptionsZscalerSubLocation) HasEnforceAuthentication() bool`

HasEnforceAuthentication returns a boolean if a field has been set.

### GetSubnets

`func (o *TunnelProviderOptionsZscalerSubLocation) GetSubnets() []string`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *TunnelProviderOptionsZscalerSubLocation) GetSubnetsOk() (*[]string, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *TunnelProviderOptionsZscalerSubLocation) SetSubnets(v []string)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *TunnelProviderOptionsZscalerSubLocation) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.

### GetUploadMbps

`func (o *TunnelProviderOptionsZscalerSubLocation) GetUploadMbps() int32`

GetUploadMbps returns the UploadMbps field if non-nil, zero value otherwise.

### GetUploadMbpsOk

`func (o *TunnelProviderOptionsZscalerSubLocation) GetUploadMbpsOk() (*int32, bool)`

GetUploadMbpsOk returns a tuple with the UploadMbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadMbps

`func (o *TunnelProviderOptionsZscalerSubLocation) SetUploadMbps(v int32)`

SetUploadMbps sets UploadMbps field to given value.

### HasUploadMbps

`func (o *TunnelProviderOptionsZscalerSubLocation) HasUploadMbps() bool`

HasUploadMbps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


