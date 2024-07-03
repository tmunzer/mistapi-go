# SsoSamlMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcsUrl** | **string** |  | [readonly] 
**EntityId** | **string** |  | [readonly] 
**LogoutUrl** | **string** |  | [readonly] 
**MetadataXml** | **string** |  | [readonly] 

## Methods

### NewSsoSamlMetadata

`func NewSsoSamlMetadata(acsUrl string, entityId string, logoutUrl string, metadataXml string, ) *SsoSamlMetadata`

NewSsoSamlMetadata instantiates a new SsoSamlMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsoSamlMetadataWithDefaults

`func NewSsoSamlMetadataWithDefaults() *SsoSamlMetadata`

NewSsoSamlMetadataWithDefaults instantiates a new SsoSamlMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcsUrl

`func (o *SsoSamlMetadata) GetAcsUrl() string`

GetAcsUrl returns the AcsUrl field if non-nil, zero value otherwise.

### GetAcsUrlOk

`func (o *SsoSamlMetadata) GetAcsUrlOk() (*string, bool)`

GetAcsUrlOk returns a tuple with the AcsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcsUrl

`func (o *SsoSamlMetadata) SetAcsUrl(v string)`

SetAcsUrl sets AcsUrl field to given value.


### GetEntityId

`func (o *SsoSamlMetadata) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *SsoSamlMetadata) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *SsoSamlMetadata) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetLogoutUrl

`func (o *SsoSamlMetadata) GetLogoutUrl() string`

GetLogoutUrl returns the LogoutUrl field if non-nil, zero value otherwise.

### GetLogoutUrlOk

`func (o *SsoSamlMetadata) GetLogoutUrlOk() (*string, bool)`

GetLogoutUrlOk returns a tuple with the LogoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutUrl

`func (o *SsoSamlMetadata) SetLogoutUrl(v string)`

SetLogoutUrl sets LogoutUrl field to given value.


### GetMetadataXml

`func (o *SsoSamlMetadata) GetMetadataXml() string`

GetMetadataXml returns the MetadataXml field if non-nil, zero value otherwise.

### GetMetadataXmlOk

`func (o *SsoSamlMetadata) GetMetadataXmlOk() (*string, bool)`

GetMetadataXmlOk returns a tuple with the MetadataXml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataXml

`func (o *SsoSamlMetadata) SetMetadataXml(v string)`

SetMetadataXml sets MetadataXml field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


