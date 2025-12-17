# Orgs SSO

```go
orgsSSO := client.OrgsSSO()
```

## Class Name

`OrgsSSO`

## Methods

* [Create Org Sso](../../doc/controllers/orgs-sso.md#create-org-sso)
* [Delete Org Sso](../../doc/controllers/orgs-sso.md#delete-org-sso)
* [Download Org Saml Metadata](../../doc/controllers/orgs-sso.md#download-org-saml-metadata)
* [Get Org Saml Metadata](../../doc/controllers/orgs-sso.md#get-org-saml-metadata)
* [Get Org Sso](../../doc/controllers/orgs-sso.md#get-org-sso)
* [List Org Sso Latest Failures](../../doc/controllers/orgs-sso.md#list-org-sso-latest-failures)
* [List Org Ssos](../../doc/controllers/orgs-sso.md#list-org-ssos)
* [Update Org Sso](../../doc/controllers/orgs-sso.md#update-org-sso)


# Create Org Sso

Create Org SSO Configuration

```go
CreateOrgSso(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.Sso) (
    models.ApiResponse[models.Sso],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Sso`](../../doc/models/sso.md) | Body, Optional | Request Body |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Sso](../../doc/models/sso.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Sso{
    CustomLogoutUrl:         models.ToPointer("string"),
    IdpCert:                 models.ToPointer("string"),
    IdpSignAlgo:             models.ToPointer(models.SsoIdpSignAlgoEnum_SHA256),
    IdpSsoUrl:               models.ToPointer("string"),
    IgnoreUnmatchedRoles:    models.ToPointer(true),
    Issuer:                  models.ToPointer("string"),
    Name:                    "string",
    NameidFormat:            models.ToPointer(models.SsoNameidFormatEnum_EMAIL),
}

apiResponse, err := orgsSSO.CreateOrgSso(ctx, orgId, &body)
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
  "created_time": 0,
  "custom_logout_url": "string",
  "default_role": "string",
  "domain": "string",
  "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "idp_cert": "string",
  "idp_sign_algo": "sha256",
  "idp_sso_url": "string",
  "ignore_unmatched_roles": true,
  "issuer": "string",
  "modified_time": 0,
  "msp_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "name": "string",
  "nameid_format": "email",
  "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "site_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "type": "string"
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


# Delete Org Sso

Delete Org SSO Configuration

```go
DeleteOrgSso(
    ctx context.Context,
    orgId uuid.UUID,
    ssoId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `ssoId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

ssoId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsSSO.DeleteOrgSso(ctx, orgId, ssoId)
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


# Download Org Saml Metadata

Download Org SAML Metadata

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
DownloadOrgSamlMetadata(
    ctx context.Context,
    orgId uuid.UUID,
    ssoId uuid.UUID) (
    models.ApiResponse[[]byte],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `ssoId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type []byte.

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

ssoId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsSSO.DownloadOrgSamlMetadata(ctx, orgId, ssoId)
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


# Get Org Saml Metadata

Get Org SAML Metadata

```go
GetOrgSamlMetadata(
    ctx context.Context,
    orgId uuid.UUID,
    ssoId uuid.UUID) (
    models.ApiResponse[models.SamlMetadata],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `ssoId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.SamlMetadata](../../doc/models/saml-metadata.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

ssoId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsSSO.GetOrgSamlMetadata(ctx, orgId, ssoId)
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
  "metadata": "<?xml version=\"1.0\" encoding=\"UTF-8\"?><md:EntityDescriptor xmlns:md=\"urn:oasis:names:tc:SAML:2.0:metadata\" entityID=\"https://api.mist.com/api/v1/saml/llDfa13f/login\" validUntil=\"2027-10-12T21:59:01Z\" xmlns:ds=\"http://www.w3.org/2000/09/xmldsig#\"><md:SPSSODescriptor AuthnRequestsSigned=\"false\" WantAssertionsSigned=\"true\" protocolSupportEnumeration=\"urn:oasis:names:tc:SAML:2.0:protocol\"><md:NameIDFormat>urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified</md:NameIDFormat><md:AssertionConsumerService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST\" Location=\"https://api.mist.com/api/v1/saml/llDfa13f/login\" index=\"0\" isDefault=\"true\"/></md:SPSSODescriptor></md:EntityDescriptor>"
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


# Get Org Sso

Get Org SSO Configuration Details

```go
GetOrgSso(
    ctx context.Context,
    orgId uuid.UUID,
    ssoId uuid.UUID) (
    models.ApiResponse[models.Sso],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `ssoId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Sso](../../doc/models/sso.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

ssoId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsSSO.GetOrgSso(ctx, orgId, ssoId)
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
  "created_time": 0,
  "custom_logout_url": "string",
  "default_role": "string",
  "domain": "string",
  "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "idp_cert": "string",
  "idp_sign_algo": "sha256",
  "idp_sso_url": "string",
  "ignore_unmatched_roles": true,
  "issuer": "string",
  "modified_time": 0,
  "msp_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "name": "string",
  "nameid_format": "email",
  "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "site_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "type": "string"
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


# List Org Sso Latest Failures

Get List of Org SSO Latest Failures

```go
ListOrgSsoLatestFailures(
    ctx context.Context,
    orgId uuid.UUID,
    ssoId uuid.UUID,
    start *string,
    end *string,
    duration *string,
    limit *int,
    page *int) (
    models.ApiResponse[models.ResponseSsoFailureSearch],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `ssoId` | `uuid.UUID` | Template, Required | - |
| `start` | `*string` | Query, Optional | Start time (epoch timestamp in seconds, or relative string like "-1d", "-1w") |
| `end` | `*string` | Query, Optional | End time (epoch timestamp in seconds, or relative string like "-1d", "-2h", "now") |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br><br>**Constraints**: `>= 1` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseSsoFailureSearch](../../doc/models/response-sso-failure-search.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

ssoId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

duration := "10m"

limit := 100

page := 1

apiResponse, err := orgsSSO.ListOrgSsoLatestFailures(ctx, orgId, ssoId, nil, nil, &duration, &limit, &page)
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


# List Org Ssos

Get List of Org SSO Configuration

```go
ListOrgSsos(
    ctx context.Context,
    orgId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.Sso],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br><br>**Constraints**: `>= 1` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.Sso](../../doc/models/sso.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

limit := 100

page := 1

apiResponse, err := orgsSSO.ListOrgSsos(ctx, orgId, &limit, &page)
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
[
  {
    "created_time": 0,
    "custom_logout_url": "string",
    "default_role": "string",
    "domain": "string",
    "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "idp_cert": "string",
    "idp_sign_algo": "sha256",
    "idp_sso_url": "string",
    "ignore_unmatched_roles": true,
    "issuer": "string",
    "modified_time": 0,
    "msp_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "name": "string",
    "nameid_format": "email",
    "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "site_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "type": "string"
  }
]
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Update Org Sso

Update Org SSO Configuration

```go
UpdateOrgSso(
    ctx context.Context,
    orgId uuid.UUID,
    ssoId uuid.UUID,
    body *models.Sso) (
    models.ApiResponse[models.Sso],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `ssoId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Sso`](../../doc/models/sso.md) | Body, Optional | Request Body |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Sso](../../doc/models/sso.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

ssoId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Sso{
    IdpType:                 models.ToPointer(models.SsoIdpTypeEnum_SAML),
    LdapBaseDn:              models.ToPointer("DC=abc,DC=com"),
    LdapBindDn:              models.ToPointer("CN=nas,CN=users,DC=abc,DC=com"),
    LdapBindPassword:        models.ToPointer("secret"),
    LdapCacerts:             []string{
        "-----BEGIN CERTIFICATE-----\\nMIIFZjCCA06gAwIBAgIIP61/1qm/uDowDQYJKoZIhvcNAQELBQE\\n-----END CERTIFICATE-----",
        "-----BEGIN CERTIFICATE-----\\nBhMCRVMxFDASBgNVBAoMC1N0YXJ0Q29tIENBMSwwKgYDVn-----END CERTIFICATE-----",
    },
    LdapClientCert:          models.ToPointer("-----BEGIN CERTIFICATE-----\\nMIIFZjCCA06gAwIBAgIIP61/1qm/uDowDQYJKoZIhvcNAQELBQE\\n-----END CERTIFICATE-----"),
    LdapClientKey:           models.ToPointer("-----BEGIN PRI..."),
    LdapGroupAttr:           models.ToPointer("memberOf"),
    LdapGroupDn:             models.ToPointer("base_dn"),
    LdapResolveGroups:       models.ToPointer(false),
    LdapServerHosts:         []string{
        "hostname",
        "63.1.3.5",
    },
    LdapType:                models.ToPointer(models.SsoLdapTypeEnum_AZURE),
    LdapUserFilter:          models.ToPointer("(mail=%s)"),
    MemberFilter:            models.ToPointer("(CN=%s)"),
    Name:                    "name6",
    NameidFormat:            models.ToPointer(models.SsoNameidFormatEnum_EMAIL),
    OauthCcClientId:         models.ToPointer("e60da615-7def-4c5a-8196-43675f45e174"),
    OauthCcClientSecret:     models.ToPointer("akL8Q~5kWFMVFYl4TFZ3fi~7cMdyDONi6cj01cpH"),
    OauthPingIdentityRegion: models.ToPointer(models.OauthPingIdentityRegionEnum_US),
    OauthRopcClientId:       models.ToPointer("9ce04c97-b5b1-4ec8-af17-f5ed42d2daf7"),
    OauthRopcClientSecret:   models.ToPointer("blM9R~6kWFMVFYl4TFZ3fi~8cMdyDONi6cj01dqI"),
    OauthTenantId:           models.ToPointer("dev-88336535"),
    OauthType:               models.ToPointer(models.SsoOauthTypeEnum_AZURE),
    RoleAttrFrom:            models.ToPointer("Role"),
    ScimEnabled:             models.ToPointer(false),
    ScimSecretToken:         models.ToPointer("FBitbKPE1aecSloPGBuqqPxDUrFeZyZk"),
}

apiResponse, err := orgsSSO.UpdateOrgSso(ctx, orgId, ssoId, &body)
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
  "created_time": 0,
  "custom_logout_url": "string",
  "default_role": "string",
  "domain": "string",
  "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "idp_cert": "string",
  "idp_sign_algo": "sha256",
  "idp_sso_url": "string",
  "ignore_unmatched_roles": true,
  "issuer": "string",
  "modified_time": 0,
  "msp_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "name": "string",
  "nameid_format": "email",
  "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "site_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "type": "string"
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

