# Orgs Sec Intel Profiles

```go
orgsSecIntelProfiles := client.OrgsSecIntelProfiles()
```

## Class Name

`OrgsSecIntelProfiles`

## Methods

* [Create Org Sec Intel Profile](../../doc/controllers/orgs-sec-intel-profiles.md#create-org-sec-intel-profile)
* [Delete Org Sec Intel Profile](../../doc/controllers/orgs-sec-intel-profiles.md#delete-org-sec-intel-profile)
* [Get Org Sec Intel Profile](../../doc/controllers/orgs-sec-intel-profiles.md#get-org-sec-intel-profile)
* [List Org Sec Intel Profiles](../../doc/controllers/orgs-sec-intel-profiles.md#list-org-sec-intel-profiles)
* [Update Org Sec Intel Profile](../../doc/controllers/orgs-sec-intel-profiles.md#update-org-sec-intel-profile)


# Create Org Sec Intel Profile

Create Sec Intel Profiles

```go
CreateOrgSecIntelProfile(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.SecintelProfile) (
    models.ApiResponse[models.SecintelProfile],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.SecintelProfile`](../../doc/models/secintel-profile.md) | Body, Optional | Request Body |

## Response Type

[`models.SecintelProfile`](../../doc/models/secintel-profile.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.SecintelProfile{
    Name:                 models.ToPointer("secintel-custom"),
    Profiles:             []models.SecintelProfileProfile{
        models.SecintelProfileProfile{
            Action:               models.ToPointer(models.SecintelProfileProfileActionEnum("default")),
            Category:             models.ToPointer(models.SecintelProfileProfileCategoryEnum("CC")),
        },
    },
}

apiResponse, err := orgsSecIntelProfiles.CreateOrgSecIntelProfile(ctx, orgId, &body)
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
  "name": "secintel-custom",
  "profiles": [
    {
      "action": "default",
      "category": "CC"
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


# Delete Org Sec Intel Profile

Delete Sec Intel Profile

```go
DeleteOrgSecIntelProfile(
    ctx context.Context,
    orgId uuid.UUID,
    secintelprofileId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `secintelprofileId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

secintelprofileId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsSecIntelProfiles.DeleteOrgSecIntelProfile(ctx, orgId, secintelprofileId)
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


# Get Org Sec Intel Profile

Get Sec Intel Profile

```go
GetOrgSecIntelProfile(
    ctx context.Context,
    orgId uuid.UUID,
    secintelprofileId uuid.UUID) (
    models.ApiResponse[models.SecintelProfile],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `secintelprofileId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.SecintelProfile`](../../doc/models/secintel-profile.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

secintelprofileId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsSecIntelProfiles.GetOrgSecIntelProfile(ctx, orgId, secintelprofileId)
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
  "name": "secintel-custom",
  "profiles": [
    {
      "action": "default",
      "category": "CC"
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


# List Org Sec Intel Profiles

Get List of Sec Intel Profiles

```go
ListOrgSecIntelProfiles(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[[]models.SecintelProfile],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`[]models.SecintelProfile`](../../doc/models/secintel-profile.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsSecIntelProfiles.ListOrgSecIntelProfiles(ctx, orgId)
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
    "name": "secintel-custom",
    "profiles": [
      {
        "action": "default",
        "category": "CC"
      }
    ]
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


# Update Org Sec Intel Profile

Update Sec Intel Profile

```go
UpdateOrgSecIntelProfile(
    ctx context.Context,
    orgId uuid.UUID,
    secintelprofileId uuid.UUID,
    body *models.SecintelProfile) (
    models.ApiResponse[models.SecintelProfile],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `secintelprofileId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.SecintelProfile`](../../doc/models/secintel-profile.md) | Body, Optional | Request Body |

## Response Type

[`models.SecintelProfile`](../../doc/models/secintel-profile.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

secintelprofileId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.SecintelProfile{
    Name:                 models.ToPointer("secintel-custom"),
    Profiles:             []models.SecintelProfileProfile{
        models.SecintelProfileProfile{
            Action:               models.ToPointer(models.SecintelProfileProfileActionEnum("default")),
            Category:             models.ToPointer(models.SecintelProfileProfileCategoryEnum("CC")),
        },
    },
}

apiResponse, err := orgsSecIntelProfiles.UpdateOrgSecIntelProfile(ctx, orgId, secintelprofileId, &body)
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
  "name": "secintel-custom",
  "profiles": [
    {
      "action": "default",
      "category": "CC"
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

