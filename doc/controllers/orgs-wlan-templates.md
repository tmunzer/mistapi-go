# Orgs WLAN Templates

```go
orgsWLANTemplates := client.OrgsWLANTemplates()
```

## Class Name

`OrgsWLANTemplates`

## Methods

* [Clone Org Template](../../doc/controllers/orgs-wlan-templates.md#clone-org-template)
* [Create Org Template](../../doc/controllers/orgs-wlan-templates.md#create-org-template)
* [Delete Org Template](../../doc/controllers/orgs-wlan-templates.md#delete-org-template)
* [Get Org Template](../../doc/controllers/orgs-wlan-templates.md#get-org-template)
* [List Org Templates](../../doc/controllers/orgs-wlan-templates.md#list-org-templates)
* [Update Org Template](../../doc/controllers/orgs-wlan-templates.md#update-org-template)


# Clone Org Template

Clone Org Template

```go
CloneOrgTemplate(
    ctx context.Context,
    orgId uuid.UUID,
    templateId uuid.UUID,
    body *models.NameString) (
    models.ApiResponse[models.Template],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `templateId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.NameString`](../../doc/models/name-string.md) | Body, Optional | Request Body |

## Response Type

[`models.Template`](../../doc/models/template.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

templateId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.NameString{
    Name:                 models.ToPointer("Cloned"),
}

apiResponse, err := orgsWLANTemplates.CloneOrgTemplate(ctx, orgId, templateId, &body)
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
  "applies": {
    "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "site_ids": [
      "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
    ],
    "sitegroup_ids": [
      "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
    ]
  },
  "created_time": 0,
  "deviceprofile_ids": [
    "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
  ],
  "exceptions": {
    "site_ids": [
      "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
    ],
    "sitegroup_ids": [
      "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
    ]
  },
  "filter_by_deviceprofile": true,
  "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "modified_time": 0,
  "name": "string",
  "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
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


# Create Org Template

Create Org Template

```go
CreateOrgTemplate(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.Template) (
    models.ApiResponse[models.Template],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Template`](../../doc/models/template.md) | Body, Optional | Request Body |

## Response Type

[`models.Template`](../../doc/models/template.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Template{
    Name:                  "name6",
}

apiResponse, err := orgsWLANTemplates.CreateOrgTemplate(ctx, orgId, &body)
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
  "applies": {
    "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "site_ids": [
      "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
    ],
    "sitegroup_ids": [
      "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
    ]
  },
  "created_time": 0,
  "deviceprofile_ids": [
    "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
  ],
  "exceptions": {
    "site_ids": [
      "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
    ],
    "sitegroup_ids": [
      "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
    ]
  },
  "filter_by_deviceprofile": true,
  "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "modified_time": 0,
  "name": "string",
  "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
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


# Delete Org Template

Delete Org Template

```go
DeleteOrgTemplate(
    ctx context.Context,
    orgId uuid.UUID,
    templateId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `templateId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

templateId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsWLANTemplates.DeleteOrgTemplate(ctx, orgId, templateId)
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


# Get Org Template

Get Org Template Details

```go
GetOrgTemplate(
    ctx context.Context,
    orgId uuid.UUID,
    templateId uuid.UUID) (
    models.ApiResponse[models.Template],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `templateId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.Template`](../../doc/models/template.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

templateId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsWLANTemplates.GetOrgTemplate(ctx, orgId, templateId)
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
  "applies": {
    "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "site_ids": [
      "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
    ],
    "sitegroup_ids": [
      "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
    ]
  },
  "created_time": 0,
  "deviceprofile_ids": [
    "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
  ],
  "exceptions": {
    "site_ids": [
      "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
    ],
    "sitegroup_ids": [
      "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
    ]
  },
  "filter_by_deviceprofile": true,
  "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "modified_time": 0,
  "name": "string",
  "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
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


# List Org Templates

Get List of Org WLAN Templates

```go
ListOrgTemplates(
    ctx context.Context,
    orgId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.Template],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br>**Constraints**: `>= 1` |

## Response Type

[`[]models.Template`](../../doc/models/template.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

limit := 100

page := 1

apiResponse, err := orgsWLANTemplates.ListOrgTemplates(ctx, orgId, &limit, &page)
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
    "applies": {
      "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
      "site_ids": [
        "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
      ],
      "sitegroup_ids": [
        "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
      ]
    },
    "created_time": 0,
    "deviceprofile_ids": [
      "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
    ],
    "exceptions": {
      "site_ids": [
        "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
      ],
      "sitegroup_ids": [
        "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
      ]
    },
    "filter_by_deviceprofile": true,
    "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "modified_time": 0,
    "name": "string",
    "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
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


# Update Org Template

Update Org Template

```go
UpdateOrgTemplate(
    ctx context.Context,
    orgId uuid.UUID,
    templateId uuid.UUID,
    body *models.Template) (
    models.ApiResponse[models.Template],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `templateId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Template`](../../doc/models/template.md) | Body, Optional | Request Body |

## Response Type

[`models.Template`](../../doc/models/template.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

templateId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Template{
    Name:                  "name6",
}

apiResponse, err := orgsWLANTemplates.UpdateOrgTemplate(ctx, orgId, templateId, &body)
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
  "applies": {
    "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "site_ids": [
      "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
    ],
    "sitegroup_ids": [
      "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
    ]
  },
  "created_time": 0,
  "deviceprofile_ids": [
    "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
  ],
  "exceptions": {
    "site_ids": [
      "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
    ],
    "sitegroup_ids": [
      "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
    ]
  },
  "filter_by_deviceprofile": true,
  "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "modified_time": 0,
  "name": "string",
  "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
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

