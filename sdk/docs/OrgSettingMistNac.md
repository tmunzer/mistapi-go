# OrgSettingMistNac

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cacerts** | Pointer to **string** | the CA certs we use to verify client certs | [optional] 
**DefaultIdpId** | Pointer to **string** | use this IDP when no explicit realm present in the incoming username/CN OR when no IDP is explicitly mapped to the incoming realm. | [optional] 
**EapSslSecurityLevel** | Pointer to **int32** | eap ssl security level see https://www.openssl.org/docs/man1.1.1/man3/SSL_CTX_set_security_level.html#DEFAULT-CALLBACK-BEHAVIOUR | [optional] [default to 2]
**EuOnly** | Pointer to **bool** | By default NAC POD failover considers all NAC pods available around the globe, i.e. EU, US, or APAC based, failover happens based on geo IP of the originating site. For strict GDPR compliancy NAC POD failover would only happen between the PODs located within the EU environment, and no authentication would take place outside of EU. This is an org setting that is applicable to WLANs, switch templates, mxedge clusters that have mist_nac enabled | [optional] [default to false]
**Idps** | Pointer to [**[]OrgSettingMistNacIdp**](OrgSettingMistNacIdp.md) |  | [optional] 
**ServerCert** | Pointer to [**OrgSettingMistNacServerCert**](OrgSettingMistNacServerCert.md) |  | [optional] 
**UseIpVersion** | Pointer to [**OrgSettingMistNacIpVersion**](OrgSettingMistNacIpVersion.md) |  | [optional] [default to ORGSETTINGMISTNACIPVERSION_V4]
**UseSslPort** | Pointer to **bool** | By default NAS devices (switches/aps) and proxies(mxedge) are configured to use port TCP2083(radsec) to reach mist-nac.  Set &#x60;use_ssl_port&#x60;&#x3D;&#x3D;&#x60;true&#x60; to override that port with TCP43 (ssl),  This is a org level setting that is applicable to wlans, switch_templates, and mxedge_clusters that have mist-nac enabled | [optional] [default to false]

## Methods

### NewOrgSettingMistNac

`func NewOrgSettingMistNac() *OrgSettingMistNac`

NewOrgSettingMistNac instantiates a new OrgSettingMistNac object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgSettingMistNacWithDefaults

`func NewOrgSettingMistNacWithDefaults() *OrgSettingMistNac`

NewOrgSettingMistNacWithDefaults instantiates a new OrgSettingMistNac object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacerts

`func (o *OrgSettingMistNac) GetCacerts() string`

GetCacerts returns the Cacerts field if non-nil, zero value otherwise.

### GetCacertsOk

`func (o *OrgSettingMistNac) GetCacertsOk() (*string, bool)`

GetCacertsOk returns a tuple with the Cacerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacerts

`func (o *OrgSettingMistNac) SetCacerts(v string)`

SetCacerts sets Cacerts field to given value.

### HasCacerts

`func (o *OrgSettingMistNac) HasCacerts() bool`

HasCacerts returns a boolean if a field has been set.

### GetDefaultIdpId

`func (o *OrgSettingMistNac) GetDefaultIdpId() string`

GetDefaultIdpId returns the DefaultIdpId field if non-nil, zero value otherwise.

### GetDefaultIdpIdOk

`func (o *OrgSettingMistNac) GetDefaultIdpIdOk() (*string, bool)`

GetDefaultIdpIdOk returns a tuple with the DefaultIdpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultIdpId

`func (o *OrgSettingMistNac) SetDefaultIdpId(v string)`

SetDefaultIdpId sets DefaultIdpId field to given value.

### HasDefaultIdpId

`func (o *OrgSettingMistNac) HasDefaultIdpId() bool`

HasDefaultIdpId returns a boolean if a field has been set.

### GetEapSslSecurityLevel

`func (o *OrgSettingMistNac) GetEapSslSecurityLevel() int32`

GetEapSslSecurityLevel returns the EapSslSecurityLevel field if non-nil, zero value otherwise.

### GetEapSslSecurityLevelOk

`func (o *OrgSettingMistNac) GetEapSslSecurityLevelOk() (*int32, bool)`

GetEapSslSecurityLevelOk returns a tuple with the EapSslSecurityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEapSslSecurityLevel

`func (o *OrgSettingMistNac) SetEapSslSecurityLevel(v int32)`

SetEapSslSecurityLevel sets EapSslSecurityLevel field to given value.

### HasEapSslSecurityLevel

`func (o *OrgSettingMistNac) HasEapSslSecurityLevel() bool`

HasEapSslSecurityLevel returns a boolean if a field has been set.

### GetEuOnly

`func (o *OrgSettingMistNac) GetEuOnly() bool`

GetEuOnly returns the EuOnly field if non-nil, zero value otherwise.

### GetEuOnlyOk

`func (o *OrgSettingMistNac) GetEuOnlyOk() (*bool, bool)`

GetEuOnlyOk returns a tuple with the EuOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEuOnly

`func (o *OrgSettingMistNac) SetEuOnly(v bool)`

SetEuOnly sets EuOnly field to given value.

### HasEuOnly

`func (o *OrgSettingMistNac) HasEuOnly() bool`

HasEuOnly returns a boolean if a field has been set.

### GetIdps

`func (o *OrgSettingMistNac) GetIdps() []OrgSettingMistNacIdp`

GetIdps returns the Idps field if non-nil, zero value otherwise.

### GetIdpsOk

`func (o *OrgSettingMistNac) GetIdpsOk() (*[]OrgSettingMistNacIdp, bool)`

GetIdpsOk returns a tuple with the Idps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdps

`func (o *OrgSettingMistNac) SetIdps(v []OrgSettingMistNacIdp)`

SetIdps sets Idps field to given value.

### HasIdps

`func (o *OrgSettingMistNac) HasIdps() bool`

HasIdps returns a boolean if a field has been set.

### GetServerCert

`func (o *OrgSettingMistNac) GetServerCert() OrgSettingMistNacServerCert`

GetServerCert returns the ServerCert field if non-nil, zero value otherwise.

### GetServerCertOk

`func (o *OrgSettingMistNac) GetServerCertOk() (*OrgSettingMistNacServerCert, bool)`

GetServerCertOk returns a tuple with the ServerCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCert

`func (o *OrgSettingMistNac) SetServerCert(v OrgSettingMistNacServerCert)`

SetServerCert sets ServerCert field to given value.

### HasServerCert

`func (o *OrgSettingMistNac) HasServerCert() bool`

HasServerCert returns a boolean if a field has been set.

### GetUseIpVersion

`func (o *OrgSettingMistNac) GetUseIpVersion() OrgSettingMistNacIpVersion`

GetUseIpVersion returns the UseIpVersion field if non-nil, zero value otherwise.

### GetUseIpVersionOk

`func (o *OrgSettingMistNac) GetUseIpVersionOk() (*OrgSettingMistNacIpVersion, bool)`

GetUseIpVersionOk returns a tuple with the UseIpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIpVersion

`func (o *OrgSettingMistNac) SetUseIpVersion(v OrgSettingMistNacIpVersion)`

SetUseIpVersion sets UseIpVersion field to given value.

### HasUseIpVersion

`func (o *OrgSettingMistNac) HasUseIpVersion() bool`

HasUseIpVersion returns a boolean if a field has been set.

### GetUseSslPort

`func (o *OrgSettingMistNac) GetUseSslPort() bool`

GetUseSslPort returns the UseSslPort field if non-nil, zero value otherwise.

### GetUseSslPortOk

`func (o *OrgSettingMistNac) GetUseSslPortOk() (*bool, bool)`

GetUseSslPortOk returns a tuple with the UseSslPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSslPort

`func (o *OrgSettingMistNac) SetUseSslPort(v bool)`

SetUseSslPort sets UseSslPort field to given value.

### HasUseSslPort

`func (o *OrgSettingMistNac) HasUseSslPort() bool`

HasUseSslPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


