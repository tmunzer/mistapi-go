# Orgs IDP Profiles

```go
orgsIDPProfiles := client.OrgsIDPProfiles()
```

## Class Name

`OrgsIDPProfiles`

## Methods

* [Create Org Idp Profile](../../doc/controllers/orgs-idp-profiles.md#create-org-idp-profile)
* [Delete Org Idp Profile](../../doc/controllers/orgs-idp-profiles.md#delete-org-idp-profile)
* [Get Org Idp Profile](../../doc/controllers/orgs-idp-profiles.md#get-org-idp-profile)
* [List Org Idp Profiles](../../doc/controllers/orgs-idp-profiles.md#list-org-idp-profiles)
* [Update Org Idp Profile](../../doc/controllers/orgs-idp-profiles.md#update-org-idp-profile)


# Create Org Idp Profile

Create Org IDP Profile

```go
CreateOrgIdpProfile(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.IdpProfile) (
    models.ApiResponse[models.IdpProfile],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.IdpProfile`](../../doc/models/idp-profile.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.IdpProfile](../../doc/models/idp-profile.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.IdpProfile{
    BaseProfile:          models.ToPointer(models.IdpProfileBaseProfileEnum_STRICT),
    Name:                 models.ToPointer("relaxed"),
    Overwrites:           []models.IdpProfileOverwrite{
        models.IdpProfileOverwrite{
            Action:               models.ToPointer(models.IdpProfileActionEnum_ALERT),
            Matching:             models.ToPointer(models.IdpProfileMatching{
                AttackName:           []string{
                    "HTTP:INVALID:HDR-FIELD",
                },
                DstSubnet:            []string{
                    "63.1.2.0/24",
                },
                Severity:             []models.IdpProfileMatchingSeverityValueEnum{
                    models.IdpProfileMatchingSeverityValueEnum_MAJOR,
                },
            }),
        },
    },
}

apiResponse, err := orgsIDPProfiles.CreateOrgIdpProfile(ctx, orgId, &body)
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
  "base_profile": "strict",
  "created_time": 0,
  "id": "874ca978-d736-4d4b-bc90-a49a29eec133",
  "modified_time": 0,
  "name": "relaxed",
  "overwrites": [
    {
      "action": "alert",
      "matching": {
        "attack_name": [
          "HTTP:INVALID:HDR-FIELD"
        ],
        "dst_subnet": [
          "63.1.2.0/24"
        ],
        "severity": [
          "major"
        ]
      }
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


# Delete Org Idp Profile

Delete Org IDP Profile

```go
DeleteOrgIdpProfile(
    ctx context.Context,
    orgId uuid.UUID,
    idpprofileId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `idpprofileId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

idpprofileId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsIDPProfiles.DeleteOrgIdpProfile(ctx, orgId, idpprofileId)
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


# Get Org Idp Profile

Get Org IDP Profile

```go
GetOrgIdpProfile(
    ctx context.Context,
    orgId uuid.UUID,
    idpprofileId uuid.UUID) (
    models.ApiResponse[models.IdpProfile],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `idpprofileId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.IdpProfile](../../doc/models/idp-profile.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

idpprofileId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsIDPProfiles.GetOrgIdpProfile(ctx, orgId, idpprofileId)
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
  "base_profile": "strict",
  "created_time": 0,
  "id": "874ca978-d736-4d4b-bc90-a49a29eec133",
  "modified_time": 0,
  "name": "relaxed",
  "overwrites": [
    {
      "action": "alert",
      "matching": {
        "attack_name": [
          "HTTP:INVALID:HDR-FIELD"
        ],
        "dst_subnet": [
          "63.1.2.0/24"
        ],
        "severity": [
          "major"
        ]
      }
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


# List Org Idp Profiles

Get the list of Org IDP Profiles

```go
ListOrgIdpProfiles(
    ctx context.Context,
    orgId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.IdpProfile],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br><br>**Constraints**: `>= 1` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.IdpProfile](../../doc/models/idp-profile.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

limit := 100

page := 1

apiResponse, err := orgsIDPProfiles.ListOrgIdpProfiles(ctx, orgId, &limit, &page)
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
    "base_profile": "strict",
    "created_time": 0,
    "id": "874ca978-d736-4d4b-bc90-a49a29eec133",
    "modified_time": 0,
    "name": "relaxed",
    "overwrites": [
      {
        "action": "alert",
        "matching": {
          "attack_name": [
            "HTTP:INVALID:HDR-FIELD"
          ],
          "dst_subnet": [
            "63.1.2.0/24"
          ],
          "severity": [
            "major"
          ]
        }
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


# Update Org Idp Profile

Update Org IDP Profile

```go
UpdateOrgIdpProfile(
    ctx context.Context,
    orgId uuid.UUID,
    idpprofileId uuid.UUID,
    body *models.IdpProfile) (
    models.ApiResponse[models.IdpProfile],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `idpprofileId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.IdpProfile`](../../doc/models/idp-profile.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.IdpProfile](../../doc/models/idp-profile.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

idpprofileId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.IdpProfile{
    BaseProfile:          models.ToPointer(models.IdpProfileBaseProfileEnum_STRICT),
    Name:                 models.ToPointer("relaxed"),
    Overwrites:           []models.IdpProfileOverwrite{
        models.IdpProfileOverwrite{
            Action:               models.ToPointer(models.IdpProfileActionEnum_ALERT),
            Matching:             models.ToPointer(models.IdpProfileMatching{
                AttackName:           []string{
                    "HTTP:INVALID:HDR-FIELD",
                },
                DstSubnet:            []string{
                    "63.1.2.0/24",
                },
                Severity:             []models.IdpProfileMatchingSeverityValueEnum{
                    models.IdpProfileMatchingSeverityValueEnum_MAJOR,
                },
            }),
        },
    },
}

apiResponse, err := orgsIDPProfiles.UpdateOrgIdpProfile(ctx, orgId, idpprofileId, &body)
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
  "base_profile": "strict",
  "created_time": 0,
  "id": "874ca978-d736-4d4b-bc90-a49a29eec133",
  "modified_time": 0,
  "name": "relaxed",
  "overwrites": [
    {
      "action": "alert",
      "matching": {
        "attack_name": [
          "HTTP:INVALID:HDR-FIELD"
        ],
        "dst_subnet": [
          "63.1.2.0/24"
        ],
        "severity": [
          "major"
        ]
      }
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

