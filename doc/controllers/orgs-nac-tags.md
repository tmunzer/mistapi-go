# Orgs NAC Tags

```go
orgsNACTags := client.OrgsNACTags()
```

## Class Name

`OrgsNACTags`

## Methods

* [Create Org Nac Tag](../../doc/controllers/orgs-nac-tags.md#create-org-nac-tag)
* [Delete Org Nac Tag](../../doc/controllers/orgs-nac-tags.md#delete-org-nac-tag)
* [Get Org Nac Tag](../../doc/controllers/orgs-nac-tags.md#get-org-nac-tag)
* [List Org Nac Tags](../../doc/controllers/orgs-nac-tags.md#list-org-nac-tags)
* [Update Org Nac Tag](../../doc/controllers/orgs-nac-tags.md#update-org-nac-tag)


# Create Org Nac Tag

Create a NAC tag used either as rule-matching criteria or as a result attribute returned when NAC allows access.

```go
CreateOrgNacTag(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.NacTag) (
    models.ApiResponse[models.NacTag],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.NacTag`](../../doc/models/nac-tag.md) | Body, Optional | - |

## Response Type

**200**: Example response

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.NacTag](../../doc/models/nac-tag.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.NacTag{
    AllowUsermacOverride: models.ToPointer(false),
    EgressVlanNames:      []string{
        "1vlan-30",
        "1vlan-20",
        "2-vlan10",
    },
    MatchAll:             models.ToPointer(false),
    NacportalId:          models.ToPointer(uuid.MustParse("1e970fec-0a7a-4d73-a472-3ef3b6a456aa")),
    Name:                 "name6",
    RadiusAttrs:          []string{
        "Idle-Timeout=600",
        "Termination-Action=RADIUS-Request",
    },
    RadiusVendorAttrs:    []string{
        "PaloAlto-Admin-Role=superuser",
        "PaloAlto-Panorama-Admin-Role=administrator",
    },
    SessionTimeout:       models.ToPointer(86000),
    Type:                 models.NacTagTypeEnum_RADIUSGROUP,
}

apiResponse, err := orgsNACTags.CreateOrgNacTag(ctx, orgId, &body)
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
  "egress_vlan_names": [
    "1vlan-30",
    "1vlan-20",
    "2vlan10"
  ],
  "name": "trunk_ap",
  "type": "egress_vlan_names"
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


# Delete Org Nac Tag

Delete an organization NAC tag by tag ID so it can no longer be used by NAC rules.

```go
DeleteOrgNacTag(
    ctx context.Context,
    orgId uuid.UUID,
    nactagId uuid.UUID) (
    http.Response,
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `nactagId` | `uuid.UUID` | Template, Required | - |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

nactagId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsNACTags.DeleteOrgNacTag(ctx, orgId, nactagId)
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


# Get Org Nac Tag

Retrieve configuration details for a specific NAC tag, including type, match values, RADIUS attributes, VLAN or session results, and portal redirection settings.

```go
GetOrgNacTag(
    ctx context.Context,
    orgId uuid.UUID,
    nactagId uuid.UUID) (
    models.ApiResponse[models.NacTag],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `nactagId` | `uuid.UUID` | Template, Required | - |

## Response Type

**200**: Example response

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.NacTag](../../doc/models/nac-tag.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

nactagId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsNACTags.GetOrgNacTag(ctx, orgId, nactagId)
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
  "egress_vlan_names": [
    "1vlan-30",
    "1vlan-20",
    "2vlan10"
  ],
  "name": "trunk_ap",
  "type": "egress_vlan_names"
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


# List Org Nac Tags

List organization NAC tags, optionally filtering by tag type, name, or match attribute.

```go
ListOrgNacTags(
    ctx context.Context,
    orgId uuid.UUID,
    mType *string,
    name *string,
    match *string,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.NacTag],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mType` | `*string` | Query, Optional | Filter results by type. enum: `egress_vlan_names`, `gbp_tag`, `match`, `radius_attrs`, `radius_group`, `radius_vendor_attrs`, `redirect_nacportal_id`, `session_timeout`, `username_attr`, `vlan`. Accepts multiple comma-separated values. |
| `name` | `*string` | Query, Optional | Filter results by name. Accepts multiple comma-separated values. |
| `match` | `*string` | Query, Optional | if `type`==`match`, Type of NAC Tag. enum: `cert_cn`, `cert_eku`, `cert_issuer`, `cert_san`, `cert_serial`, `cert_sub`, `cert_template`, `client_mac`, `edr_status`, `gbp_tag`, `hostname`, `idp_role`, `ingress_vlan`, `mdm_status`, `nas_ip`, `radius_group`, `realm`, `ssid`, `user_name`, `usermac_label`. Accepts multiple comma-separated values. |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | Select the page number to return when using page-based pagination; starts at `1`<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1` |

## Response Type

**200**: Example response

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.NacTag](../../doc/models/nac-tag.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mType := "match,vlan"

name := "name-a,name-b"

match := "ssid,idp_role"

limit := 100

page := 1

apiResponse, err := orgsNACTags.ListOrgNacTags(ctx, orgId, &mType, &name, &match, &limit, &page)
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


# Update Org Nac Tag

Update a NAC tag, including matcher values or result attributes such as RADIUS attributes, VLAN, session timeout, username attribute, or portal redirection.

```go
UpdateOrgNacTag(
    ctx context.Context,
    orgId uuid.UUID,
    nactagId uuid.UUID,
    body *models.NacTag) (
    models.ApiResponse[models.NacTag],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `nactagId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.NacTag`](../../doc/models/nac-tag.md) | Body, Optional | - |

## Response Type

**200**: Example response

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.NacTag](../../doc/models/nac-tag.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

nactagId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.NacTag{
    AllowUsermacOverride: models.ToPointer(false),
    EgressVlanNames:      []string{
        "1vlan-30",
        "1vlan-20",
        "2-vlan10",
    },
    MatchAll:             models.ToPointer(false),
    NacportalId:          models.ToPointer(uuid.MustParse("1e970fec-0a7a-4d73-a472-3ef3b6a456aa")),
    Name:                 "name6",
    RadiusAttrs:          []string{
        "Idle-Timeout=600",
        "Termination-Action=RADIUS-Request",
    },
    RadiusVendorAttrs:    []string{
        "PaloAlto-Admin-Role=superuser",
        "PaloAlto-Panorama-Admin-Role=administrator",
    },
    SessionTimeout:       models.ToPointer(86000),
    Type:                 models.NacTagTypeEnum_RADIUSGROUP,
}

apiResponse, err := orgsNACTags.UpdateOrgNacTag(ctx, orgId, nactagId, &body)
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
  "egress_vlan_names": [
    "1vlan-30",
    "1vlan-20",
    "2vlan10"
  ],
  "name": "trunk_ap",
  "type": "egress_vlan_names"
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

