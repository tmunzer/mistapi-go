# MS Ps SSO

```go
mSPsSSO := client.MSPsSSO()
```

## Class Name

`MSPsSSO`

## Methods

* [Create Msp Sso](../../doc/controllers/ms-ps-sso.md#create-msp-sso)
* [Delete Msp Sso](../../doc/controllers/ms-ps-sso.md#delete-msp-sso)
* [Download Msp Sso Saml Metadata](../../doc/controllers/ms-ps-sso.md#download-msp-sso-saml-metadata)
* [Get Msp Sso](../../doc/controllers/ms-ps-sso.md#get-msp-sso)
* [Get Msp Sso Saml Metadata](../../doc/controllers/ms-ps-sso.md#get-msp-sso-saml-metadata)
* [List Msp Sso Latest Failures](../../doc/controllers/ms-ps-sso.md#list-msp-sso-latest-failures)
* [List Msp Ssos](../../doc/controllers/ms-ps-sso.md#list-msp-ssos)
* [Update Msp Sso](../../doc/controllers/ms-ps-sso.md#update-msp-sso)


# Create Msp Sso

Create MSP SSO profile

```go
CreateMspSso(
    ctx context.Context,
    mspId uuid.UUID,
    body *models.Sso) (
    models.ApiResponse[models.Sso],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `mspId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Sso`](../../doc/models/sso.md) | Body, Optional | Request Body |

## Response Type

[`models.Sso`](../../doc/models/sso.md)

## Example Usage

```go
ctx := context.Background()

mspId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Sso{
    IdpType:               models.ToPointer(models.SsoIdpTypeEnum("saml")),
    LdapBaseDn:            models.ToPointer("DC=abc,DC=com"),
    LdapBindDn:            models.ToPointer("CN=nas,CN=users,DC=abc,DC=com"),
    LdapBindPassword:      models.ToPointer("secret"),
    LdapCacerts:           []string{
        "-----BEGIN CERTIFICATE-----\\nMIIFZjCCA06gAwIBAgIIP61/1qm/uDowDQYJKoZIhvcNAQELBQE\\n-----END CERTIFICATE-----",
        "-----BEGIN CERTIFICATE-----\\nBhMCRVMxFDASBgNVBAoMC1N0YXJ0Q29tIENBMSwwKgYDVn-----END CERTIFICATE-----",
    },
    LdapClientCert:        models.ToPointer("-----BEGIN CERTIFICATE-----\\nMIIFZjCCA06gAwIBAgIIP61/1qm/uDowDQYJKoZIhvcNAQELBQE\\n-----END CERTIFICATE-----"),
    LdapClientKey:         models.ToPointer("-----BEGIN PRI..."),
    LdapGroupAttr:         models.ToPointer("memberOf"),
    LdapGroupDn:           models.ToPointer("base_dn"),
    LdapMemberFilter:      models.ToPointer("(CN=%s)"),
    LdapResolveGroups:     models.ToPointer(false),
    LdapServerHosts:       []string{
        "hostname",
        "63.1.3.5",
    },
    LdapType:              models.ToPointer(models.SsoLdapTypeEnum("azure")),
    LdapUserFilter:        models.ToPointer("(mail=%s)"),
    Name:                  "name6",
    NameidFormat:          models.ToPointer(models.SsoNameidFormatEnum("email")),
    OauthCcClientId:       models.ToPointer("e60da615-7def-4c5a-8196-43675f45e174"),
    OauthCcClientSecret:   models.ToPointer("akL8Q~5kWFMVFYl4TFZ3fi~7cMdyDONi6cj01cpH"),
    OauthRopcClientId:     models.ToPointer("9ce04c97-b5b1-4ec8-af17-f5ed42d2daf7"),
    OauthRopcClientSecret: models.ToPointer("blM9R~6kWFMVFYl4TFZ3fi~8cMdyDONi6cj01dqI"),
    OauthTenantId:         models.ToPointer("dev-88336535"),
    OauthType:             models.ToPointer(models.SsoOauthTypeEnum("azure")),
    OrgId:                 models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
    RoleAttrFrom:          models.ToPointer("Role"),
    ScimEnabled:           models.ToPointer(false),
    ScimSecretToken:       models.ToPointer("secret token"),
    SiteId:                models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
}

apiResponse, err := mSPsSSO.CreateMspSso(ctx, mspId, &body)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Delete Msp Sso

Delete MSP SSO Config

```go
DeleteMspSso(
    ctx context.Context,
    mspId uuid.UUID,
    ssoId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `mspId` | `uuid.UUID` | Template, Required | - |
| `ssoId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

mspId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

ssoId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := mSPsSSO.DeleteMspSso(ctx, mspId, ssoId)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Download Msp Sso Saml Metadata

Download MSP SSO SAML Metadata

Example of metadata.xml:

```xml
<?xml version="1.0" encoding="UTF-8"?><md:EntityDescriptor xmlns:md="urn:oasis:names:tc:SAML:2.0:metadata" entityID="https://api.mist.com/api/v1/saml/5hdF5g/login" validUntil="2027-10-12T21:59:01Z" xmlns:ds="http://www.w3.org/2000/09/xmldsig#">
  <md:SPSSODescriptor AuthnRequestsSigned="false" WantAssertionsSigned="true" protocolSupportEnumeration="urn:oasis:names:tc:SAML:2.0:protocol">
      <md:SingleLogoutService Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST" Location="https://api.mist.com/api/v1/saml/5hdF5g/logout" />
      <md:NameIDFormat>urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified</md:NameIDFormat>
      <md:AssertionConsumerService Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST" Location="https://api.mist.com/api/v1/saml/5hdF5g/login" index="0" isDefault="true"/>
      <md:AttributeConsumingService index="0">
          <md:ServiceName xml:lang="en-US">Mist</md:ServiceName>
          <md:RequestedAttribute Name="Role" NameFormat="urn:oasis:names:tc:SAML:2.0:attrname-format:basic" isRequired="true"/>
          <md:RequestedAttribute Name="FirstName" NameFormat="urn:oasis:names:tc:SAML:2.0:attrname-format:basic" isRequired="false"/>
          <md:RequestedAttribute Name="LastName" NameFormat="urn:oasis:names:tc:SAML:2.0:attrname-format:basic" isRequired="false"/>
      </md:AttributeConsumingService>
  </md:SPSSODescriptor>
</md:EntityDescriptor>
```

```go
DownloadMspSsoSamlMetadata(
    ctx context.Context,
    mspId uuid.UUID,
    ssoId uuid.UUID) (
    models.ApiResponse[[]byte],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `mspId` | `uuid.UUID` | Template, Required | - |
| `ssoId` | `uuid.UUID` | Template, Required | - |

## Response Type

`[]byte`

## Example Usage

```go
ctx := context.Background()

mspId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

ssoId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := mSPsSSO.DownloadMspSsoSamlMetadata(ctx, mspId, ssoId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Get Msp Sso

Get MSP SSO Config

```go
GetMspSso(
    ctx context.Context,
    mspId uuid.UUID,
    ssoId uuid.UUID) (
    models.ApiResponse[models.Sso],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `mspId` | `uuid.UUID` | Template, Required | - |
| `ssoId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.Sso`](../../doc/models/sso.md)

## Example Usage

```go
ctx := context.Background()

mspId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

ssoId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := mSPsSSO.GetMspSso(ctx, mspId, ssoId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Get Msp Sso Saml Metadata

Get MSP SSO SAML Metadata

```go
GetMspSsoSamlMetadata(
    ctx context.Context,
    mspId uuid.UUID,
    ssoId uuid.UUID) (
    models.ApiResponse[models.SsoSamlMetadata],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `mspId` | `uuid.UUID` | Template, Required | - |
| `ssoId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.SsoSamlMetadata`](../../doc/models/sso-saml-metadata.md)

## Example Usage

```go
ctx := context.Background()

mspId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

ssoId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := mSPsSSO.GetMspSsoSamlMetadata(ctx, mspId, ssoId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response *(as JSON)*

```json
{
  "acs_url": "https://api.mist.com/api/v1/saml/llDfa13f/login",
  "entity_id": "https://api.mist.com/api/v1/saml/llDfa13f/login",
  "logout_url": "https://api.mist.com/api/v1/saml/llDfa13f/logout",
  "metadata_xml": "<?xml version=\"1.0\" encoding=\"UTF-8\"?><md:EntityDescriptor xmlns:md=\"urn:oasis:names:tc:SAML:2.0:metadata\" entityID=\"https://api.mist.com/api/v1/saml/llDfa13f/login\" validUntil=\"2027-10-12T21:59:01Z\" xmlns:ds=\"http://www.w3.org/2000/09/xmldsig#\"><md:SPSSODescriptor AuthnRequestsSigned=\"false\" WantAssertionsSigned=\"true\" protocolSupportEnumeration=\"urn:oasis:names:tc:SAML:2.0:protocol\"><md:NameIDFormat>urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified</md:NameIDFormat><md:AssertionConsumerService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST\" Location=\"https://api.mist.com/api/v1/saml/llDfa13f/login\" index=\"0\" isDefault=\"true\"/></md:SPSSODescriptor></md:EntityDescriptor>"
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# List Msp Sso Latest Failures

Get List of MSP SSO Latest Failures

```go
ListMspSsoLatestFailures(
    ctx context.Context,
    mspId uuid.UUID,
    ssoId uuid.UUID) (
    models.ApiResponse[models.ResponseSsoFailureSearch],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `mspId` | `uuid.UUID` | Template, Required | - |
| `ssoId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.ResponseSsoFailureSearch`](../../doc/models/response-sso-failure-search.md)

## Example Usage

```go
ctx := context.Background()

mspId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

ssoId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := mSPsSSO.ListMspSsoLatestFailures(ctx, mspId, ssoId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response *(as JSON)*

```json
{
  "results": [
    {
      "detail": "string",
      "saml_assertion_xml": "string",
      "timestamp": 0
    }
  ]
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# List Msp Ssos

List MSP SSO Configs

```go
ListMspSsos(
    ctx context.Context,
    mspId uuid.UUID) (
    models.ApiResponse[[]models.Sso],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `mspId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`[]models.Sso`](../../doc/models/sso.md)

## Example Usage

```go
ctx := context.Background()

mspId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := mSPsSSO.ListMspSsos(ctx, mspId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Update Msp Sso

Update MSP SSO config

```go
UpdateMspSso(
    ctx context.Context,
    mspId uuid.UUID,
    ssoId uuid.UUID,
    body *models.Sso) (
    models.ApiResponse[models.Sso],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `mspId` | `uuid.UUID` | Template, Required | - |
| `ssoId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Sso`](../../doc/models/sso.md) | Body, Optional | Request Body |

## Response Type

[`models.Sso`](../../doc/models/sso.md)

## Example Usage

```go
ctx := context.Background()

mspId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

ssoId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Sso{
    IdpType:               models.ToPointer(models.SsoIdpTypeEnum("saml")),
    LdapBaseDn:            models.ToPointer("DC=abc,DC=com"),
    LdapBindDn:            models.ToPointer("CN=nas,CN=users,DC=abc,DC=com"),
    LdapBindPassword:      models.ToPointer("secret"),
    LdapCacerts:           []string{
        "-----BEGIN CERTIFICATE-----\\nMIIFZjCCA06gAwIBAgIIP61/1qm/uDowDQYJKoZIhvcNAQELBQE\\n-----END CERTIFICATE-----",
        "-----BEGIN CERTIFICATE-----\\nBhMCRVMxFDASBgNVBAoMC1N0YXJ0Q29tIENBMSwwKgYDVn-----END CERTIFICATE-----",
    },
    LdapClientCert:        models.ToPointer("-----BEGIN CERTIFICATE-----\\nMIIFZjCCA06gAwIBAgIIP61/1qm/uDowDQYJKoZIhvcNAQELBQE\\n-----END CERTIFICATE-----"),
    LdapClientKey:         models.ToPointer("-----BEGIN PRI..."),
    LdapGroupAttr:         models.ToPointer("memberOf"),
    LdapGroupDn:           models.ToPointer("base_dn"),
    LdapMemberFilter:      models.ToPointer("(CN=%s)"),
    LdapResolveGroups:     models.ToPointer(false),
    LdapServerHosts:       []string{
        "hostname",
        "63.1.3.5",
    },
    LdapType:              models.ToPointer(models.SsoLdapTypeEnum("azure")),
    LdapUserFilter:        models.ToPointer("(mail=%s)"),
    Name:                  "name6",
    NameidFormat:          models.ToPointer(models.SsoNameidFormatEnum("email")),
    OauthCcClientId:       models.ToPointer("e60da615-7def-4c5a-8196-43675f45e174"),
    OauthCcClientSecret:   models.ToPointer("akL8Q~5kWFMVFYl4TFZ3fi~7cMdyDONi6cj01cpH"),
    OauthRopcClientId:     models.ToPointer("9ce04c97-b5b1-4ec8-af17-f5ed42d2daf7"),
    OauthRopcClientSecret: models.ToPointer("blM9R~6kWFMVFYl4TFZ3fi~8cMdyDONi6cj01dqI"),
    OauthTenantId:         models.ToPointer("dev-88336535"),
    OauthType:             models.ToPointer(models.SsoOauthTypeEnum("azure")),
    OrgId:                 models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
    RoleAttrFrom:          models.ToPointer("Role"),
    ScimEnabled:           models.ToPointer(false),
    ScimSecretToken:       models.ToPointer("secret token"),
    SiteId:                models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
}

apiResponse, err := mSPsSSO.UpdateMspSso(ctx, mspId, ssoId, &body)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |

