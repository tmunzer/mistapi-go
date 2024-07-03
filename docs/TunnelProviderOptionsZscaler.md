# TunnelProviderOptionsZscaler

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AupAcceptanceRequired** | Pointer to **bool** |  | [optional] [default to true]
**AupExpire** | Pointer to **int32** | days before AUP is requested again | [optional] [default to 1]
**AupSslProxy** | Pointer to **bool** | proxy HTTPs traffic, requiring Zscaler cert to be installed in browser | [optional] [default to false]
**DownloadMbps** | Pointer to **int32** | the download bandwidth cap of the link, in Mbps | [optional] 
**EnableAup** | Pointer to **bool** | if &#x60;use_xff&#x60;&#x3D;&#x3D;&#x60;true&#x60;, display Acceptable Use Policy (AUP) | [optional] [default to false]
**EnableCaution** | Pointer to **bool** | when &#x60;enforce_authentication&#x60;&#x3D;&#x3D;&#x60;false&#x60;, display caution notification for non-authenticated users | [optional] [default to false]
**EnforceAuthentication** | Pointer to **bool** |  | [optional] [default to false]
**Name** | Pointer to **string** |  | [optional] 
**SubLocations** | Pointer to [**[]TunnelProviderOptionsZscalerSubLocation**](TunnelProviderOptionsZscalerSubLocation.md) | if &#x60;use_xff&#x60;&#x3D;&#x3D;&#x60;true&#x60; | [optional] 
**UploadMbps** | Pointer to **int32** | the download bandwidth cap of the link, in Mbps | [optional] 
**UseXff** | Pointer to **bool** | location uses proxy chaining to forward traffic | [optional] 

## Methods

### NewTunnelProviderOptionsZscaler

`func NewTunnelProviderOptionsZscaler() *TunnelProviderOptionsZscaler`

NewTunnelProviderOptionsZscaler instantiates a new TunnelProviderOptionsZscaler object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTunnelProviderOptionsZscalerWithDefaults

`func NewTunnelProviderOptionsZscalerWithDefaults() *TunnelProviderOptionsZscaler`

NewTunnelProviderOptionsZscalerWithDefaults instantiates a new TunnelProviderOptionsZscaler object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAupAcceptanceRequired

`func (o *TunnelProviderOptionsZscaler) GetAupAcceptanceRequired() bool`

GetAupAcceptanceRequired returns the AupAcceptanceRequired field if non-nil, zero value otherwise.

### GetAupAcceptanceRequiredOk

`func (o *TunnelProviderOptionsZscaler) GetAupAcceptanceRequiredOk() (*bool, bool)`

GetAupAcceptanceRequiredOk returns a tuple with the AupAcceptanceRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAupAcceptanceRequired

`func (o *TunnelProviderOptionsZscaler) SetAupAcceptanceRequired(v bool)`

SetAupAcceptanceRequired sets AupAcceptanceRequired field to given value.

### HasAupAcceptanceRequired

`func (o *TunnelProviderOptionsZscaler) HasAupAcceptanceRequired() bool`

HasAupAcceptanceRequired returns a boolean if a field has been set.

### GetAupExpire

`func (o *TunnelProviderOptionsZscaler) GetAupExpire() int32`

GetAupExpire returns the AupExpire field if non-nil, zero value otherwise.

### GetAupExpireOk

`func (o *TunnelProviderOptionsZscaler) GetAupExpireOk() (*int32, bool)`

GetAupExpireOk returns a tuple with the AupExpire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAupExpire

`func (o *TunnelProviderOptionsZscaler) SetAupExpire(v int32)`

SetAupExpire sets AupExpire field to given value.

### HasAupExpire

`func (o *TunnelProviderOptionsZscaler) HasAupExpire() bool`

HasAupExpire returns a boolean if a field has been set.

### GetAupSslProxy

`func (o *TunnelProviderOptionsZscaler) GetAupSslProxy() bool`

GetAupSslProxy returns the AupSslProxy field if non-nil, zero value otherwise.

### GetAupSslProxyOk

`func (o *TunnelProviderOptionsZscaler) GetAupSslProxyOk() (*bool, bool)`

GetAupSslProxyOk returns a tuple with the AupSslProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAupSslProxy

`func (o *TunnelProviderOptionsZscaler) SetAupSslProxy(v bool)`

SetAupSslProxy sets AupSslProxy field to given value.

### HasAupSslProxy

`func (o *TunnelProviderOptionsZscaler) HasAupSslProxy() bool`

HasAupSslProxy returns a boolean if a field has been set.

### GetDownloadMbps

`func (o *TunnelProviderOptionsZscaler) GetDownloadMbps() int32`

GetDownloadMbps returns the DownloadMbps field if non-nil, zero value otherwise.

### GetDownloadMbpsOk

`func (o *TunnelProviderOptionsZscaler) GetDownloadMbpsOk() (*int32, bool)`

GetDownloadMbpsOk returns a tuple with the DownloadMbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadMbps

`func (o *TunnelProviderOptionsZscaler) SetDownloadMbps(v int32)`

SetDownloadMbps sets DownloadMbps field to given value.

### HasDownloadMbps

`func (o *TunnelProviderOptionsZscaler) HasDownloadMbps() bool`

HasDownloadMbps returns a boolean if a field has been set.

### GetEnableAup

`func (o *TunnelProviderOptionsZscaler) GetEnableAup() bool`

GetEnableAup returns the EnableAup field if non-nil, zero value otherwise.

### GetEnableAupOk

`func (o *TunnelProviderOptionsZscaler) GetEnableAupOk() (*bool, bool)`

GetEnableAupOk returns a tuple with the EnableAup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAup

`func (o *TunnelProviderOptionsZscaler) SetEnableAup(v bool)`

SetEnableAup sets EnableAup field to given value.

### HasEnableAup

`func (o *TunnelProviderOptionsZscaler) HasEnableAup() bool`

HasEnableAup returns a boolean if a field has been set.

### GetEnableCaution

`func (o *TunnelProviderOptionsZscaler) GetEnableCaution() bool`

GetEnableCaution returns the EnableCaution field if non-nil, zero value otherwise.

### GetEnableCautionOk

`func (o *TunnelProviderOptionsZscaler) GetEnableCautionOk() (*bool, bool)`

GetEnableCautionOk returns a tuple with the EnableCaution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCaution

`func (o *TunnelProviderOptionsZscaler) SetEnableCaution(v bool)`

SetEnableCaution sets EnableCaution field to given value.

### HasEnableCaution

`func (o *TunnelProviderOptionsZscaler) HasEnableCaution() bool`

HasEnableCaution returns a boolean if a field has been set.

### GetEnforceAuthentication

`func (o *TunnelProviderOptionsZscaler) GetEnforceAuthentication() bool`

GetEnforceAuthentication returns the EnforceAuthentication field if non-nil, zero value otherwise.

### GetEnforceAuthenticationOk

`func (o *TunnelProviderOptionsZscaler) GetEnforceAuthenticationOk() (*bool, bool)`

GetEnforceAuthenticationOk returns a tuple with the EnforceAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceAuthentication

`func (o *TunnelProviderOptionsZscaler) SetEnforceAuthentication(v bool)`

SetEnforceAuthentication sets EnforceAuthentication field to given value.

### HasEnforceAuthentication

`func (o *TunnelProviderOptionsZscaler) HasEnforceAuthentication() bool`

HasEnforceAuthentication returns a boolean if a field has been set.

### GetName

`func (o *TunnelProviderOptionsZscaler) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TunnelProviderOptionsZscaler) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TunnelProviderOptionsZscaler) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TunnelProviderOptionsZscaler) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubLocations

`func (o *TunnelProviderOptionsZscaler) GetSubLocations() []TunnelProviderOptionsZscalerSubLocation`

GetSubLocations returns the SubLocations field if non-nil, zero value otherwise.

### GetSubLocationsOk

`func (o *TunnelProviderOptionsZscaler) GetSubLocationsOk() (*[]TunnelProviderOptionsZscalerSubLocation, bool)`

GetSubLocationsOk returns a tuple with the SubLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubLocations

`func (o *TunnelProviderOptionsZscaler) SetSubLocations(v []TunnelProviderOptionsZscalerSubLocation)`

SetSubLocations sets SubLocations field to given value.

### HasSubLocations

`func (o *TunnelProviderOptionsZscaler) HasSubLocations() bool`

HasSubLocations returns a boolean if a field has been set.

### GetUploadMbps

`func (o *TunnelProviderOptionsZscaler) GetUploadMbps() int32`

GetUploadMbps returns the UploadMbps field if non-nil, zero value otherwise.

### GetUploadMbpsOk

`func (o *TunnelProviderOptionsZscaler) GetUploadMbpsOk() (*int32, bool)`

GetUploadMbpsOk returns a tuple with the UploadMbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadMbps

`func (o *TunnelProviderOptionsZscaler) SetUploadMbps(v int32)`

SetUploadMbps sets UploadMbps field to given value.

### HasUploadMbps

`func (o *TunnelProviderOptionsZscaler) HasUploadMbps() bool`

HasUploadMbps returns a boolean if a field has been set.

### GetUseXff

`func (o *TunnelProviderOptionsZscaler) GetUseXff() bool`

GetUseXff returns the UseXff field if non-nil, zero value otherwise.

### GetUseXffOk

`func (o *TunnelProviderOptionsZscaler) GetUseXffOk() (*bool, bool)`

GetUseXffOk returns a tuple with the UseXff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseXff

`func (o *TunnelProviderOptionsZscaler) SetUseXff(v bool)`

SetUseXff sets UseXff field to given value.

### HasUseXff

`func (o *TunnelProviderOptionsZscaler) HasUseXff() bool`

HasUseXff returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


