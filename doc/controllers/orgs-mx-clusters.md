# Orgs Mx Clusters

```go
orgsMxClusters := client.OrgsMxClusters()
```

## Class Name

`OrgsMxClusters`

## Methods

* [Create Org Mx Edge Cluster](../../doc/controllers/orgs-mx-clusters.md#create-org-mx-edge-cluster)
* [Delete Org Mx Edge Cluster](../../doc/controllers/orgs-mx-clusters.md#delete-org-mx-edge-cluster)
* [Get Org Mx Edge Cluster](../../doc/controllers/orgs-mx-clusters.md#get-org-mx-edge-cluster)
* [List Org Mx Edge Clusters](../../doc/controllers/orgs-mx-clusters.md#list-org-mx-edge-clusters)
* [Update Org Mx Edge Cluster](../../doc/controllers/orgs-mx-clusters.md#update-org-mx-edge-cluster)


# Create Org Mx Edge Cluster

Create a Mist Edge cluster with tunnel termination, RadSec, NAC,
and management settings.

**Note**: It is not recommended to combine multiple roles (tunnel termination, RadSec, NAC) on the same Mist Edge cluster

```go
CreateOrgMxEdgeCluster(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.Mxcluster) (
    models.ApiResponse[models.Mxcluster],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Mxcluster`](../../doc/models/mxcluster.md) | Body, Optional | Request Body |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Mxcluster](../../doc/models/mxcluster.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Mxcluster{
    Name:                      models.ToPointer("string"),
    Radsec:                    models.ToPointer(models.MxclusterRadsec{
        AuthServers:          []models.MxclusterRadsecAuthServer{
            models.MxclusterRadsecAuthServer{
                Host:                 models.ToPointer("string"),
                Port:                 models.ToPointer(0),
            },
        },
        Enabled:              models.ToPointer(true),
    }),
    TuntermApSubnets:          []string{
        "string",
    },
    TuntermHosts:              []string{
        "string",
    },
}

apiResponse, err := orgsMxClusters.CreateOrgMxEdgeCluster(ctx, orgId, &body)
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
  "for_site": true,
  "id": "468f6eca-6276-4993-bfeb-53cbbbba6f68",
  "modified_time": 0,
  "name": "string",
  "org_id": "a40f5d1f-d889-42e9-94ea-b9b33585fc6b",
  "radsec": {
    "acct_servers": [
      {
        "host": "string",
        "port": 1813,
        "secret": "string"
      }
    ],
    "auth_servers": [
      {
        "host": "string",
        "port": 1812,
        "secret": "string"
      }
    ],
    "enabled": true,
    "server_selection": "ordered"
  },
  "radsec_tls": {
    "keypair": "string"
  },
  "site_id": "72771e6a-6f5e-4de4-a5b9-1266c4197811",
  "tunterm_ap_subnets": [
    "string"
  ],
  "tunterm_dhcpd_config": {
    "enabled": false,
    "property1": {
      "enabled": false,
      "servers": [
        "string"
      ],
      "type": "relay"
    },
    "property2": {
      "enabled": false,
      "servers": [
        "string"
      ],
      "type": "relay"
    },
    "servers": [
      "string"
    ],
    "type": "relay"
  },
  "tunterm_extra_routes": {
    "property1": {
      "via": "string"
    },
    "property2": {
      "via": "string"
    }
  },
  "tunterm_hosts": [
    "string"
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


# Delete Org Mx Edge Cluster

Delete a Mist Edge cluster by cluster ID, removing its cluster configuration from the organization.

```go
DeleteOrgMxEdgeCluster(
    ctx context.Context,
    orgId uuid.UUID,
    mxclusterId uuid.UUID) (
    http.Response,
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mxclusterId` | `uuid.UUID` | Template, Required | - |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mxclusterId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsMxClusters.DeleteOrgMxEdgeCluster(ctx, orgId, mxclusterId)
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


# Get Org Mx Edge Cluster

Retrieve configuration details for a specific Mist Edge cluster, including tunneling, RadSec, NAC, and management settings.

```go
GetOrgMxEdgeCluster(
    ctx context.Context,
    orgId uuid.UUID,
    mxclusterId uuid.UUID) (
    models.ApiResponse[models.Mxcluster],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mxclusterId` | `uuid.UUID` | Template, Required | - |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Mxcluster](../../doc/models/mxcluster.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mxclusterId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsMxClusters.GetOrgMxEdgeCluster(ctx, orgId, mxclusterId)
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
  "for_site": true,
  "id": "468f6eca-6276-4993-bfeb-53cbbbba6f68",
  "modified_time": 0,
  "name": "string",
  "org_id": "a40f5d1f-d889-42e9-94ea-b9b33585fc6b",
  "radsec": {
    "acct_servers": [
      {
        "host": "string",
        "port": 1813,
        "secret": "string"
      }
    ],
    "auth_servers": [
      {
        "host": "string",
        "port": 1812,
        "secret": "string"
      }
    ],
    "enabled": true,
    "server_selection": "ordered"
  },
  "radsec_tls": {
    "keypair": "string"
  },
  "site_id": "72771e6a-6f5e-4de4-a5b9-1266c4197811",
  "tunterm_ap_subnets": [
    "string"
  ],
  "tunterm_dhcpd_config": {
    "enabled": false,
    "property1": {
      "enabled": false,
      "servers": [
        "string"
      ],
      "type": "relay"
    },
    "property2": {
      "enabled": false,
      "servers": [
        "string"
      ],
      "type": "relay"
    },
    "servers": [
      "string"
    ],
    "type": "relay"
  },
  "tunterm_extra_routes": {
    "property1": {
      "via": "string"
    },
    "property2": {
      "via": "string"
    }
  },
  "tunterm_hosts": [
    "string"
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


# List Org Mx Edge Clusters

List Mist Edge clusters in the organization, which group one or more Mist Edge devices for tunneling, RadSec, and related edge services.

```go
ListOrgMxEdgeClusters(
    ctx context.Context,
    orgId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.Mxcluster],
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

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.Mxcluster](../../doc/models/mxcluster.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

limit := 100

page := 1

apiResponse, err := orgsMxClusters.ListOrgMxEdgeClusters(ctx, orgId, &limit, &page)
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
    "for_site": true,
    "id": "467f6eca-6276-4993-bfeb-53cbbbba6f78",
    "modified_time": 0,
    "name": "string",
    "org_id": "a40f5d1f-d889-42e9-94ea-b9b33585fc6b",
    "radsec": {
      "acct_servers": [
        {
          "host": "string",
          "port": 1813,
          "secret": "string"
        }
      ],
      "auth_servers": [
        {
          "host": "string",
          "port": 1812,
          "secret": "string"
        }
      ],
      "enabled": true,
      "server_selection": "ordered"
    },
    "radsec_tls": {
      "keypair": "string"
    },
    "site_id": "72771e6a-6f5e-4de4-a5b9-1266c4197811",
    "tunterm_ap_subnets": [
      "string"
    ],
    "tunterm_dhcpd_config": {
      "enabled": false,
      "property1": {
        "enabled": false,
        "servers": [
          "string"
        ],
        "type": "relay"
      },
      "property2": {
        "enabled": false,
        "servers": [
          "string"
        ],
        "type": "relay"
      },
      "servers": [
        "string"
      ],
      "type": "relay"
    },
    "tunterm_extra_routes": {
      "property1": {
        "via": "string"
      },
      "property2": {
        "via": "string"
      }
    },
    "tunterm_hosts": [
      "string"
    ]
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


# Update Org Mx Edge Cluster

Update a Mist Edge cluster's tunneling, RadSec, NAC, and management settings.

```go
UpdateOrgMxEdgeCluster(
    ctx context.Context,
    orgId uuid.UUID,
    mxclusterId uuid.UUID,
    body *models.Mxcluster) (
    models.ApiResponse[models.Mxcluster],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mxclusterId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Mxcluster`](../../doc/models/mxcluster.md) | Body, Optional | Request Body |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Mxcluster](../../doc/models/mxcluster.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mxclusterId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Mxcluster{
    Name:                      models.ToPointer("string"),
    Radsec:                    models.ToPointer(models.MxclusterRadsec{
        AuthServers:          []models.MxclusterRadsecAuthServer{
            models.MxclusterRadsecAuthServer{
                Host:                 models.ToPointer("string"),
                Port:                 models.ToPointer(0),
            },
        },
        Enabled:              models.ToPointer(true),
    }),
    TuntermApSubnets:          []string{
        "string",
    },
    TuntermHosts:              []string{
        "string",
    },
}

apiResponse, err := orgsMxClusters.UpdateOrgMxEdgeCluster(ctx, orgId, mxclusterId, &body)
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
  "for_site": true,
  "id": "468f6eca-6276-4993-bfeb-53cbbbba6f68",
  "modified_time": 0,
  "name": "string",
  "org_id": "a40f5d1f-d889-42e9-94ea-b9b33585fc6b",
  "radsec": {
    "acct_servers": [
      {
        "host": "string",
        "port": 1813,
        "secret": "string"
      }
    ],
    "auth_servers": [
      {
        "host": "string",
        "port": 1812,
        "secret": "string"
      }
    ],
    "enabled": true,
    "server_selection": "ordered"
  },
  "radsec_tls": {
    "keypair": "string"
  },
  "site_id": "72771e6a-6f5e-4de4-a5b9-1266c4197811",
  "tunterm_ap_subnets": [
    "string"
  ],
  "tunterm_dhcpd_config": {
    "enabled": false,
    "property1": {
      "enabled": false,
      "servers": [
        "string"
      ],
      "type": "relay"
    },
    "property2": {
      "enabled": false,
      "servers": [
        "string"
      ],
      "type": "relay"
    },
    "servers": [
      "string"
    ],
    "type": "relay"
  },
  "tunterm_extra_routes": {
    "property1": {
      "via": "string"
    },
    "property2": {
      "via": "string"
    }
  },
  "tunterm_hosts": [
    "string"
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

