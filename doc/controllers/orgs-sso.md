# Orgs SSO

```go
orgsSSO := client.OrgsSSO()
```

## Class Name

`OrgsSSO`

## Methods

* [Create Org Sso](../../doc/controllers/orgs-sso.md#create-org-sso)
* [Delete Org Sso](../../doc/controllers/orgs-sso.md#delete-org-sso)
* [Delete Org Sso Admins](../../doc/controllers/orgs-sso.md#delete-org-sso-admins)
* [Download Org Saml Metadata](../../doc/controllers/orgs-sso.md#download-org-saml-metadata)
* [Get Org Saml Metadata](../../doc/controllers/orgs-sso.md#get-org-saml-metadata)
* [Get Org Sso](../../doc/controllers/orgs-sso.md#get-org-sso)
* [List Org Sso Latest Failures](../../doc/controllers/orgs-sso.md#list-org-sso-latest-failures)
* [List Org Ssos](../../doc/controllers/orgs-sso.md#list-org-ssos)
* [Update Org Sso](../../doc/controllers/orgs-sso.md#update-org-sso)


# Create Org Sso

Create an organization SSO identity provider configuration, including provider settings and role-handling behavior.

```go
CreateOrgSso(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.Sso) (
    models.ApiResponse[models.Sso],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Sso`](../../doc/models/sso.md) | Body, Optional | Request Body |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Sso](../../doc/models/sso.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Sso{
    CustomLogoutUrl:          models.ToPointer("string"),
    IdpCert:                  models.ToPointer("string"),
    IdpSignAlgo:              models.ToPointer(models.SsoIdpSignAlgoEnum_SHA256),
    IdpSsoUrl:                models.ToPointer("string"),
    IgnoreUnmatchedRoles:     models.ToPointer(true),
    Issuer:                   models.ToPointer("string"),
    Name:                     "string",
    NameidFormat:             models.ToPointer(models.SsoNameidFormatEnum_EMAIL),
}

apiResponse, err := orgsSSO.CreateOrgSso(ctx, orgId, &body)
if err != nil {
    switch typedErr := err.(type) {
        case *errors.ResponseHttp400:
            log.Fatalln("ResponseHttp400Exception: ", typedErr)
        case *errors.ResponseHttp401:
            log.Fatalln("ResponseHttp401Exception: ", typedErr)
        case *errors.ResponseHttp403:
            log.Fatalln("ResponseHttp403Exception: ", typedErr)
        case *errors.ResponseHttp404:
            log.Fatalln("ResponseHttp404Exception: ", typedErr)
        case *errors.ResponseHttp429:
            log.Fatalln("ResponseHttp429Exception: ", typedErr)
        default:
            log.Fatalln(err)
    }
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# Delete Org Sso

Delete an organization SSO identity provider configuration so it can no longer be used for administrator login.

```go
DeleteOrgSso(
    ctx context.Context,
    orgId uuid.UUID,
    ssoId uuid.UUID) (
    http.Response,
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `ssoId` | `uuid.UUID` | Template, Required | - |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

ssoId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsSSO.DeleteOrgSso(ctx, orgId, ssoId)
if err != nil {
    switch typedErr := err.(type) {
        case *errors.ResponseHttp400:
            log.Fatalln("ResponseHttp400Exception: ", typedErr)
        case *errors.ResponseHttp401:
            log.Fatalln("ResponseHttp401Exception: ", typedErr)
        case *errors.ResponseHttp403:
            log.Fatalln("ResponseHttp403Exception: ", typedErr)
        case *errors.ResponseHttp404:
            log.Fatalln("ResponseHttp404Exception: ", typedErr)
        case *errors.ResponseHttp429:
            log.Fatalln("ResponseHttp429Exception: ", typedErr)
        default:
            log.Fatalln(err)
    }
} else {
    fmt.Println(resp.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# Delete Org Sso Admins

Remove SSO-linked organization administrator accounts by email for this SSO profile.

```go
DeleteOrgSsoAdmins(
    ctx context.Context,
    orgId uuid.UUID,
    ssoId uuid.UUID,
    body *models.SsoDeleteAdmins) (
    models.ApiResponse[models.SsoDeleteAdminsResponse],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `ssoId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.SsoDeleteAdmins`](../../doc/models/sso-delete-admins.md) | Body, Optional | Request Body |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.SsoDeleteAdminsResponse](../../doc/models/sso-delete-admins-response.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

ssoId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.SsoDeleteAdmins{
    Emails:               []string{
        "john@abc.com",
        "may@abc.com",
    },
}

apiResponse, err := orgsSSO.DeleteOrgSsoAdmins(ctx, orgId, ssoId, &body)
if err != nil {
    switch typedErr := err.(type) {
        case *errors.ResponseHttp400:
            log.Fatalln("ResponseHttp400Exception: ", typedErr)
        case *errors.ResponseHttp401:
            log.Fatalln("ResponseHttp401Exception: ", typedErr)
        case *errors.ResponseHttp403:
            log.Fatalln("ResponseHttp403Exception: ", typedErr)
        case *errors.ResponseHttp404:
            log.Fatalln("ResponseHttp404Exception: ", typedErr)
        case *errors.ResponseHttp429:
            log.Fatalln("ResponseHttp429Exception: ", typedErr)
        default:
            log.Fatalln(err)
    }
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# Download Org Saml Metadata

Download generated SAML service provider metadata XML for this
organization SSO configuration. Use this XML to configure the identity provider
with Mist service-provider details such as entity ID, ACS URL, logout URL,
NameID format, and requested attributes.

Example of metadata.xml:

```xml
<?xml version=\"1.0\" encoding=\"UTF-8\"?><md:EntityDescriptor xmlns:md=\"urn:oasis:names:tc:SAML:2.0:metadata\" entityID=\"https://api.mist.com/api/v1/saml/5hdF5g/login\"\ validUntil=\"2027-10-12T21:59:01Z\" xmlns:ds=\"http://www.w3.org/2000/09/xmldsig#\">
    <md:SPSSODescriptor AuthnRequestsSigned=\"false\" WantAssertionsSigned=\"true\" protocolSupportEnumeration=\"urn:oasis:names:tc:SAML:2.0:protocol\">
        <md:SingleLogoutService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST\"\ Location=\"https://api.mist.com/api/v1/saml/5hdF5g/logout\" />
        <md:NameIDFormat>urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified</md:NameIDFormat>
        <md:AssertionConsumerService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST\"\ Location=\"https://api.mist.com/api/v1/saml/5hdF5g/login\" index=\"0\" isDefault=\"true\"/>
        <md:AttributeConsumingService index=\"0\">
            <md:ServiceName xml:lang=\"en-US\">Mist</md:ServiceName>
            <md:RequestedAttribute\ Name=\"Role\" NameFormat=\"urn:oasis:names:tc:SAML:2.0:attrname-format:basic\"\ isRequired=\"true\"/>
            <md:RequestedAttribute Name=\"FirstName\"\ NameFormat=\"urn:oasis:names:tc:SAML:2.0:attrname-format:basic\" isRequired=\"false\"/>
            <md:RequestedAttribute Name=\"LastName\" NameFormat=\"urn:oasis:names:tc:SAML:2.0:attrname-format:basic\" isRequired=\"false\"/>
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

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `ssoId` | `uuid.UUID` | Template, Required | - |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type []byte.

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

ssoId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsSSO.DownloadOrgSamlMetadata(ctx, orgId, ssoId)
if err != nil {
    switch typedErr := err.(type) {
        case *errors.ResponseHttp400:
            log.Fatalln("ResponseHttp400Exception: ", typedErr)
        case *errors.ResponseHttp401:
            log.Fatalln("ResponseHttp401Exception: ", typedErr)
        case *errors.ResponseHttp403:
            log.Fatalln("ResponseHttp403Exception: ", typedErr)
        case *errors.ResponseHttp404:
            log.Fatalln("ResponseHttp404Exception: ", typedErr)
        case *errors.ResponseHttp429:
            log.Fatalln("ResponseHttp429Exception: ", typedErr)
        default:
            log.Fatalln(err)
    }
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# Get Org Saml Metadata

Return generated SAML service provider metadata for this organization SSO configuration as JSON.

```go
GetOrgSamlMetadata(
    ctx context.Context,
    orgId uuid.UUID,
    ssoId uuid.UUID) (
    models.ApiResponse[models.SamlMetadata],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `ssoId` | `uuid.UUID` | Template, Required | - |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.SamlMetadata](../../doc/models/saml-metadata.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

ssoId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsSSO.GetOrgSamlMetadata(ctx, orgId, ssoId)
if err != nil {
    switch typedErr := err.(type) {
        case *errors.ResponseHttp400:
            log.Fatalln("ResponseHttp400Exception: ", typedErr)
        case *errors.ResponseHttp401:
            log.Fatalln("ResponseHttp401Exception: ", typedErr)
        case *errors.ResponseHttp403:
            log.Fatalln("ResponseHttp403Exception: ", typedErr)
        case *errors.ResponseHttp404:
            log.Fatalln("ResponseHttp404Exception: ", typedErr)
        case *errors.ResponseHttp429:
            log.Fatalln("ResponseHttp429Exception: ", typedErr)
        default:
            log.Fatalln(err)
    }
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# Get Org Sso

Return one organization SSO identity provider configuration, including provider settings and generated SSO URLs.

```go
GetOrgSso(
    ctx context.Context,
    orgId uuid.UUID,
    ssoId uuid.UUID) (
    models.ApiResponse[models.Sso],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `ssoId` | `uuid.UUID` | Template, Required | - |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Sso](../../doc/models/sso.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

ssoId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsSSO.GetOrgSso(ctx, orgId, ssoId)
if err != nil {
    switch typedErr := err.(type) {
        case *errors.ResponseHttp400:
            log.Fatalln("ResponseHttp400Exception: ", typedErr)
        case *errors.ResponseHttp401:
            log.Fatalln("ResponseHttp401Exception: ", typedErr)
        case *errors.ResponseHttp403:
            log.Fatalln("ResponseHttp403Exception: ", typedErr)
        case *errors.ResponseHttp404:
            log.Fatalln("ResponseHttp404Exception: ", typedErr)
        case *errors.ResponseHttp429:
            log.Fatalln("ResponseHttp429Exception: ", typedErr)
        default:
            log.Fatalln(err)
    }
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# List Org Sso Latest Failures

List recent authentication failures for this organization SSO configuration, including failure details and captured SAML assertion data when available.

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

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `ssoId` | `uuid.UUID` | Template, Required | - |
| `start` | `*string` | Query, Optional | Lower bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d` or `-1w` |
| `end` | `*string` | Query, Optional | Upper bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d`, `-2h`, or `now` |
| `duration` | `*string` | Query, Optional | Time range duration for the query, using relative units such as `10m`, `7d`, or `2w`<br><br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | Select the page number to return when using page-based pagination; starts at `1`<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1` |

## Response Type

**200**: OK

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
    switch typedErr := err.(type) {
        case *errors.ResponseHttp400:
            log.Fatalln("ResponseHttp400Exception: ", typedErr)
        case *errors.ResponseHttp401:
            log.Fatalln("ResponseHttp401Exception: ", typedErr)
        case *errors.ResponseHttp403:
            log.Fatalln("ResponseHttp403Exception: ", typedErr)
        case *errors.ResponseHttp404:
            log.Fatalln("ResponseHttp404Exception: ", typedErr)
        case *errors.ResponseHttp429:
            log.Fatalln("ResponseHttp429Exception: ", typedErr)
        default:
            log.Fatalln(err)
    }
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# List Org Ssos

List SSO identity provider configurations defined for this organization.

```go
ListOrgSsos(
    ctx context.Context,
    orgId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.Sso],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | Select the page number to return when using page-based pagination; starts at `1`<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1` |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.Sso](../../doc/models/sso.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

limit := 100

page := 1

apiResponse, err := orgsSSO.ListOrgSsos(ctx, orgId, &limit, &page)
if err != nil {
    switch typedErr := err.(type) {
        case *errors.ResponseHttp400:
            log.Fatalln("ResponseHttp400Exception: ", typedErr)
        case *errors.ResponseHttp401:
            log.Fatalln("ResponseHttp401Exception: ", typedErr)
        case *errors.ResponseHttp403:
            log.Fatalln("ResponseHttp403Exception: ", typedErr)
        case *errors.ResponseHttp404:
            log.Fatalln("ResponseHttp404Exception: ", typedErr)
        case *errors.ResponseHttp429:
            log.Fatalln("ResponseHttp429Exception: ", typedErr)
        default:
            log.Fatalln(err)
    }
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# Update Org Sso

Update an organization SSO identity provider configuration, such as IdP URLs, certificates, issuer, NameID format, and unmatched-role handling.

```go
UpdateOrgSso(
    ctx context.Context,
    orgId uuid.UUID,
    ssoId uuid.UUID,
    body *models.Sso) (
    models.ApiResponse[models.Sso],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `ssoId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Sso`](../../doc/models/sso.md) | Body, Optional | Request Body |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Sso](../../doc/models/sso.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

ssoId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Sso{
    IdpType:                  models.ToPointer(models.SsoIdpTypeEnum_SAML),
    LdapBaseDn:               models.ToPointer("DC=abc,DC=com"),
    LdapBindDn:               models.ToPointer("CN=nas,CN=users,DC=abc,DC=com"),
    LdapBindPassword:         models.ToPointer("secret"),
    LdapCacerts:              []string{
        "-----BEGIN CERTIFICATE-----\\nMIIFZjCCA06gAwIBAgIIP61/1qm/uDowDQYJKoZIhvcNAQELBQE\\n-----END CERTIFICATE-----",
        "-----BEGIN CERTIFICATE-----\\nBhMCRVMxFDASBgNVBAoMC1N0YXJ0Q29tIENBMSwwKgYDVn-----END CERTIFICATE-----",
    },
    LdapClientCert:           models.ToPointer("-----BEGIN CERTIFICATE-----\\nMIIFZjCCA06gAwIBAgIIP61/1qm/uDowDQYJKoZIhvcNAQELBQE\\n-----END CERTIFICATE-----"),
    LdapClientKey:            models.ToPointer("-----BEGIN PRI..."),
    LdapGroupAttr:            models.ToPointer("memberOf"),
    LdapGroupDn:              models.ToPointer("base_dn"),
    LdapResolveGroups:        models.ToPointer(false),
    LdapServerHosts:          []string{
        "hostname",
        "63.1.3.5",
    },
    LdapType:                 models.ToPointer(models.SsoLdapTypeEnum_AZURE),
    LdapUserFilter:           models.ToPointer("(mail=%s)"),
    MemberFilter:             models.ToPointer("(CN=%s)"),
    Name:                     "name6",
    NameidFormat:             models.ToPointer(models.SsoNameidFormatEnum_EMAIL),
    OauthCcClientId:          models.ToPointer("e60da615-7def-4c5a-8196-43675f45e174"),
    OauthCcClientSecret:      models.ToPointer("akL8Q~5kWFMVFYl4TFZ3fi~7cMdyDONi6cj01cpH"),
    OauthPingIdentityRegion:  models.ToPointer(models.OauthPingIdentityRegionEnum_US),
    OauthProviderDomain:      models.ToPointer(models.OauthProviderDomainEnum_ENUMOKTACOM),
    OauthRopcClientId:        models.ToPointer("9ce04c97-b5b1-4ec8-af17-f5ed42d2daf7"),
    OauthRopcClientSecret:    models.ToPointer("blM9R~6kWFMVFYl4TFZ3fi~8cMdyDONi6cj01dqI"),
    OauthTenantId:            models.ToPointer("dev-88336535"),
    OauthType:                models.ToPointer(models.SsoOauthTypeEnum_AZURE),
    OpenroamingSsids:         []string{
        "ssid_name1",
        "ssid_name2",
    },
    OpenroamingWbaClientCert: models.ToPointer("-----BEGIN CERTIFICATE-----\\nMIIFZjCCA06gAwIBAgIIP61/1qm/uDowDQYJKoZIhvcNAQELBQE\\n-----END CERTIFICATE-----"),
    OpenroamingWbaClientKey:  models.ToPointer("-----BEGIN PRIVATE KEY-----\\nMIIFZjCCA06gAwIBAgIIP61/1qm/uDowDQYJKoZIhvcNAQELBQE\\n-----END PRIVATE KEY-----"),
    RoleAttrFrom:             models.ToPointer("Role"),
    ScimEnabled:              models.ToPointer(false),
    ScimSecretToken:          models.ToPointer("FBitbKPE1aecSloPGBuqqPxDUrFeZyZk"),
}

apiResponse, err := orgsSSO.UpdateOrgSso(ctx, orgId, ssoId, &body)
if err != nil {
    switch typedErr := err.(type) {
        case *errors.ResponseHttp400:
            log.Fatalln("ResponseHttp400Exception: ", typedErr)
        case *errors.ResponseHttp401:
            log.Fatalln("ResponseHttp401Exception: ", typedErr)
        case *errors.ResponseHttp403:
            log.Fatalln("ResponseHttp403Exception: ", typedErr)
        case *errors.ResponseHttp404:
            log.Fatalln("ResponseHttp404Exception: ", typedErr)
        case *errors.ResponseHttp429:
            log.Fatalln("ResponseHttp429Exception: ", typedErr)
        default:
            log.Fatalln(err)
    }
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |

