# Orgs Mx Edges

```go
orgsMxEdges := client.OrgsMxEdges()
```

## Class Name

`OrgsMxEdges`

## Methods

* [Add Org Mx Edge Image](../../doc/controllers/orgs-mx-edges.md#add-org-mx-edge-image)
* [Assign Org Mx Edge to Site](../../doc/controllers/orgs-mx-edges.md#assign-org-mx-edge-to-site)
* [Bounce Org Mx Edge Data Ports](../../doc/controllers/orgs-mx-edges.md#bounce-org-mx-edge-data-ports)
* [Claim Org Mx Edge](../../doc/controllers/orgs-mx-edges.md#claim-org-mx-edge)
* [Control Org Mx Edge Services](../../doc/controllers/orgs-mx-edges.md#control-org-mx-edge-services)
* [Count Org Mx Edges](../../doc/controllers/orgs-mx-edges.md#count-org-mx-edges)
* [Count Org Site Mx Edge Events](../../doc/controllers/orgs-mx-edges.md#count-org-site-mx-edge-events)
* [Create Org Mx Edge](../../doc/controllers/orgs-mx-edges.md#create-org-mx-edge)
* [Delete Org Mx Edge](../../doc/controllers/orgs-mx-edges.md#delete-org-mx-edge)
* [Delete Org Mx Edge Image](../../doc/controllers/orgs-mx-edges.md#delete-org-mx-edge-image)
* [Get Org Mx Edge](../../doc/controllers/orgs-mx-edges.md#get-org-mx-edge)
* [Get Org Mx Edge Upgrade Info](../../doc/controllers/orgs-mx-edges.md#get-org-mx-edge-upgrade-info)
* [List Org Mx Edges](../../doc/controllers/orgs-mx-edges.md#list-org-mx-edges)
* [Restart Org Mx Edge](../../doc/controllers/orgs-mx-edges.md#restart-org-mx-edge)
* [Search Org Mist Edge Events](../../doc/controllers/orgs-mx-edges.md#search-org-mist-edge-events)
* [Search Org Mx Edges](../../doc/controllers/orgs-mx-edges.md#search-org-mx-edges)
* [Unassign Org Mx Edge From Site](../../doc/controllers/orgs-mx-edges.md#unassign-org-mx-edge-from-site)
* [Unregister Org Mx Edge](../../doc/controllers/orgs-mx-edges.md#unregister-org-mx-edge)
* [Update Org Mx Edge](../../doc/controllers/orgs-mx-edges.md#update-org-mx-edge)
* [Upload Org Mx Edge Support Files](../../doc/controllers/orgs-mx-edges.md#upload-org-mx-edge-support-files)


# Add Org Mx Edge Image

Attach up to 3 images to a mxedge

```go
AddOrgMxEdgeImage(
    ctx context.Context,
    orgId uuid.UUID,
    mxedgeId uuid.UUID,
    imageNumber int,
    body *models.ImageImport) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mxedgeId` | `uuid.UUID` | Template, Required | - |
| `imageNumber` | `int` | Template, Required | **Constraints**: `>= 1`, `<= 3` |
| `body` | [`*models.ImageImport`](../../doc/models/image-import.md) | Body, Optional | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mxedgeId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

imageNumber := 110



resp, err := orgsMxEdges.AddOrgMxEdgeImage(ctx, orgId, mxedgeId, imageNumber, nil)
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


# Assign Org Mx Edge to Site

Assign Org MxEdge to Site

```go
AssignOrgMxEdgeToSite(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.MxedgesAssign) (
    models.ApiResponse[models.ResponseAssignSuccess],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.MxedgesAssign`](../../doc/models/mxedges-assign.md) | Body, Optional | Request Body |

## Response Type

[`models.ResponseAssignSuccess`](../../doc/models/response-assign-success.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.MxedgesAssign{
    MxedgeIds:            []uuid.UUID{
        uuid.MustParse("387804a7-3474-85ce-15a2-f9a9684c9c90"),
    },
    SiteId:               uuid.MustParse("4ac1dcf4-9d8b-7211-65c4-057819f0862b"),
}

apiResponse, err := orgsMxEdges.AssignOrgMxEdgeToSite(ctx, orgId, &body)
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
  "success": [
    "5c5b350e0001"
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


# Bounce Org Mx Edge Data Ports

Bounce TunTerm Data Ports

```go
BounceOrgMxEdgeDataPorts(
    ctx context.Context,
    orgId uuid.UUID,
    mxedgeId uuid.UUID,
    body *models.UtilsTuntermBouncePort) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mxedgeId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.UtilsTuntermBouncePort`](../../doc/models/utils-tunterm-bounce-port.md) | Body, Optional | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mxedgeId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.UtilsTuntermBouncePort{
    Ports:                []string{
        "0",
        "2",
    },
}

resp, err := orgsMxEdges.BounceOrgMxEdgeDataPorts(ctx, orgId, mxedgeId, &body)
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


# Claim Org Mx Edge

For a Mist Edge in default state, it will show a random claim code like `135-546-673` which you can “claim” it into your Org

```go
ClaimOrgMxEdge(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.CodeString) (
    models.ApiResponse[models.ResponseClaimMxEdge],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.CodeString`](../../doc/models/code-string.md) | Body, Optional | Request Body |

## Response Type

[`models.ResponseClaimMxEdge`](../../doc/models/response-claim-mx-edge.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.CodeString{
    Code:                 "135-546-673",
}

apiResponse, err := orgsMxEdges.ClaimOrgMxEdge(ctx, orgId, &body)
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
  "id": "95ddd29a-6a3c-929e-a431-51a5b09f36a6",
  "magic": "ExNpT5gi-ADR8WTFd4EiQPY3cP5WdSoD"
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


# Control Org Mx Edge Services

Control Services on a Mist Edge

```go
ControlOrgMxEdgeServices(
    ctx context.Context,
    orgId uuid.UUID,
    mxedgeId uuid.UUID,
    name models.MxedgeServiceNameEnum,
    action models.MxedgeServiceActionEnum) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mxedgeId` | `uuid.UUID` | Template, Required | - |
| `name` | [`models.MxedgeServiceNameEnum`](../../doc/models/mxedge-service-name-enum.md) | Template, Required | enum: `mxagent`, `mxdas`, `mxnacedge`, `mxocproxy`, `radsecproxy`, `tunterm` |
| `action` | [`models.MxedgeServiceActionEnum`](../../doc/models/mxedge-service-action-enum.md) | Template, Required | restart or start or stop |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mxedgeId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

name := models.MxedgeServiceNameEnum("mxnacedge")

action := models.MxedgeServiceActionEnum("restart")

resp, err := orgsMxEdges.ControlOrgMxEdgeServices(ctx, orgId, mxedgeId, name, action)
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


# Count Org Mx Edges

Count Org Mist Edges

```go
CountOrgMxEdges(
    ctx context.Context,
    orgId uuid.UUID,
    distinct *models.OrgMxedgeCountDistinctEnum,
    mxedgeId *string,
    siteId *string,
    mxclusterId *string,
    model *string,
    distro *string,
    tuntermVersion *string,
    sort *string,
    stats *bool,
    start *int,
    end *int,
    duration *string,
    limit *int,
    page *int) (
    models.ApiResponse[models.RepsonseCount],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.OrgMxedgeCountDistinctEnum`](../../doc/models/org-mxedge-count-distinct-enum.md) | Query, Optional | **Default**: `"model"` |
| `mxedgeId` | `*string` | Query, Optional | mist edge id |
| `siteId` | `*string` | Query, Optional | mist edge site id |
| `mxclusterId` | `*string` | Query, Optional | mist edge cluster id |
| `model` | `*string` | Query, Optional | model name |
| `distro` | `*string` | Query, Optional | debian code name(buster, bullseye) |
| `tuntermVersion` | `*string` | Query, Optional | tunterm version |
| `sort` | `*string` | Query, Optional | sort options, -prefix represents DESC order, default is -last_seen |
| `stats` | `*bool` | Query, Optional | whether to return device stats, default is false |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w<br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br>**Constraints**: `>= 1` |

## Response Type

[`models.RepsonseCount`](../../doc/models/repsonse-count.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

distinct := models.OrgMxedgeCountDistinctEnum("model")





















duration := "10m"

limit := 100

page := 1

apiResponse, err := orgsMxEdges.CountOrgMxEdges(ctx, orgId, &distinct, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &duration, &limit, &page)
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
  "distinct": "string",
  "end": 0,
  "limit": 0,
  "results": [
    {
      "count": 0,
      "property": "string"
    }
  ],
  "start": 0,
  "total": 0
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


# Count Org Site Mx Edge Events

Count Org Mist Edge Events

```go
CountOrgSiteMxEdgeEvents(
    ctx context.Context,
    orgId uuid.UUID,
    distinct *models.OrgMxedgeEventsCountDistinctEnum,
    mxedgeId *string,
    mxclusterId *string,
    mType *string,
    service *string,
    start *int,
    end *int,
    duration *string,
    limit *int) (
    models.ApiResponse[models.RepsonseCount],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.OrgMxedgeEventsCountDistinctEnum`](../../doc/models/org-mxedge-events-count-distinct-enum.md) | Query, Optional | **Default**: `"mxedge_id"` |
| `mxedgeId` | `*string` | Query, Optional | mist edge id |
| `mxclusterId` | `*string` | Query, Optional | mist edge cluster id |
| `mType` | `*string` | Query, Optional | see [listDeviceEventsDefinitions]($e/Constants%20Events/listDeviceEventsDefinitions) |
| `service` | `*string` | Query, Optional | service running on mist edge(mxagent, tunterm etc) |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w<br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |

## Response Type

[`models.RepsonseCount`](../../doc/models/repsonse-count.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

distinct := models.OrgMxedgeEventsCountDistinctEnum("mxedge_id")













duration := "10m"

limit := 100

apiResponse, err := orgsMxEdges.CountOrgSiteMxEdgeEvents(ctx, orgId, &distinct, nil, nil, nil, nil, nil, nil, &duration, &limit)
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
  "distinct": "string",
  "end": 0,
  "limit": 0,
  "results": [
    {
      "count": 0,
      "property": "string"
    }
  ],
  "start": 0,
  "total": 0
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


# Create Org Mx Edge

Create MxEdge

```go
CreateOrgMxEdge(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.Mxedge) (
    models.ApiResponse[models.Mxedge],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Mxedge`](../../doc/models/mxedge.md) | Body, Optional | Request Body |

## Response Type

[`models.Mxedge`](../../doc/models/mxedge.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Mxedge{
    Model:                     "ME-100",
    MxclusterId:               models.ToPointer(uuid.MustParse("572586b7-f97b-a22b-526c-8b97a3f609c4")),
    MxedgeMgmt:                models.ToPointer(models.MxedgeMgmt{
    }),
    Name:                      "Guest",
    NtpServers:                []string{
    },
    OobIpConfig:               models.ToPointer(models.MxedgeOobIpConfig{
    }),
    Services:                  []string{
        "tunterm",
    },
    TuntermIpConfig:           models.ToPointer(models.MxedgeTuntermIpConfig{
        Gateway:              "10.2.1.254",
        Ip:                   "10.2.1.1",
        Netmask:              "255.255.255.0",
    }),
    TuntermPortConfig:         models.ToPointer(models.TuntermPortConfig{
        DownstreamPorts:            []string{
            "0",
            "1",
            "2",
            "3",
        },
        SeparateUpstreamDownstream: models.ToPointer(true),
        UpstreamPortVlanId:         models.ToPointer(1),
        UpstreamPorts:              []string{
            "0",
            "1",
            "2",
            "3",
        },
    }),
    TuntermSwitchConfig:       models.ToPointer(models.MxedgeTuntermSwitchConfigs{
        Enabled:              models.ToPointer(true),
        AdditionalProperties: map[string]models.MxedgeTuntermSwitchConfig{
            "0": models.MxedgeTuntermSwitchConfig{
                PortVlanId:           models.ToPointer(1),
                VlanIds:              []models.MxedgeTuntermSwitchConfigVlanId{
                    models.MxedgeTuntermSwitchConfigVlanIdContainer.FromNumber(5),
                    models.MxedgeTuntermSwitchConfigVlanIdContainer.FromNumber(2),
                    models.MxedgeTuntermSwitchConfigVlanIdContainer.FromNumber(3),
                },
            },
        },
    }),
}

apiResponse, err := orgsMxEdges.CreateOrgMxEdge(ctx, orgId, &body)
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
  "id": "95ddd29a-6a3c-929e-a431-51a5b09f36a6",
  "magic": "L-NpT5gi-ADR8WTFd4EiQPY3cP5WdSoD",
  "model": "ME-100",
  "mxagent_registered": true,
  "mxcluster_id": "572586b7-f97b-a22b-526c-8b97a3f609c4",
  "mxedge_mgmt": {
    "mist_password": "MIST_PASSWORD",
    "root_password": "ROOT_PASSWORD"
  },
  "name": "Guest",
  "ntp_servers": [],
  "oob_ip_config": {
    "dns": [
      "8.8.8.8",
      "4.4.4.4"
    ],
    "gateway": "10.2.1.254",
    "ip": "10.2.1.10",
    "netmask": "255.255.255.0",
    "type": "static"
  },
  "tunterm_dhcpd_config": {
    "2": {
      "enabled": true,
      "servers": [
        "11.2.3.44"
      ]
    },
    "enabled": false,
    "servers": [
      "11.2.3.4"
    ]
  },
  "tunterm_extra_routes": {
    "11.0.0.0/8": {
      "via": "10.3.3.1"
    }
  },
  "tunterm_ip_config": {
    "dns": [
      "8.8.8.8"
    ],
    "dns_suffix": [
      ".mist.local"
    ],
    "gateway": "10.2.1.254",
    "ip": "10.2.1.1",
    "netmask": "255.255.255.0"
  }
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


# Delete Org Mx Edge

Delete Org MxEdge

```go
DeleteOrgMxEdge(
    ctx context.Context,
    orgId uuid.UUID,
    mxedgeId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mxedgeId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mxedgeId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsMxEdges.DeleteOrgMxEdge(ctx, orgId, mxedgeId)
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


# Delete Org Mx Edge Image

Remove MxEdge Image

```go
DeleteOrgMxEdgeImage(
    ctx context.Context,
    orgId uuid.UUID,
    mxedgeId uuid.UUID,
    imageNumber int) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mxedgeId` | `uuid.UUID` | Template, Required | - |
| `imageNumber` | `int` | Template, Required | **Constraints**: `>= 1`, `<= 3` |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mxedgeId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

imageNumber := 110

resp, err := orgsMxEdges.DeleteOrgMxEdgeImage(ctx, orgId, mxedgeId, imageNumber)
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


# Get Org Mx Edge

Get Org MxEdge details

```go
GetOrgMxEdge(
    ctx context.Context,
    orgId uuid.UUID,
    mxedgeId uuid.UUID) (
    models.ApiResponse[models.Mxedge],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mxedgeId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.Mxedge`](../../doc/models/mxedge.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mxedgeId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsMxEdges.GetOrgMxEdge(ctx, orgId, mxedgeId)
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
  "id": "95ddd29a-6a3c-929e-a431-51a5b09f36a6",
  "magic": "L-NpT5gi-ADR8WTFd4EiQPY3cP5WdSoD",
  "model": "ME-100",
  "mxagent_registered": true,
  "mxcluster_id": "572586b7-f97b-a22b-526c-8b97a3f609c4",
  "mxedge_mgmt": {
    "mist_password": "MIST_PASSWORD",
    "root_password": "ROOT_PASSWORD"
  },
  "name": "Guest",
  "ntp_servers": [],
  "oob_ip_config": {
    "dns": [
      "8.8.8.8",
      "4.4.4.4"
    ],
    "gateway": "10.2.1.254",
    "ip": "10.2.1.10",
    "netmask": "255.255.255.0",
    "type": "static"
  },
  "tunterm_dhcpd_config": {
    "2": {
      "enabled": true,
      "servers": [
        "11.2.3.44"
      ]
    },
    "enabled": false,
    "servers": [
      "11.2.3.4"
    ]
  },
  "tunterm_extra_routes": {
    "11.0.0.0/8": {
      "via": "10.3.3.1"
    }
  },
  "tunterm_ip_config": {
    "dns": [
      "8.8.8.8"
    ],
    "dns_suffix": [
      ".mist.local"
    ],
    "gateway": "10.2.1.254",
    "ip": "10.2.1.1",
    "netmask": "255.255.255.0"
  }
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


# Get Org Mx Edge Upgrade Info

Get Mist Edge Upgrade Information

```go
GetOrgMxEdgeUpgradeInfo(
    ctx context.Context,
    orgId uuid.UUID,
    channel *models.GetOrgMxedgeUpgradeInfoChannelEnum) (
    models.ApiResponse[[]models.MxedgeUpgradeInfoItems],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `channel` | [`*models.GetOrgMxedgeUpgradeInfoChannelEnum`](../../doc/models/get-org-mxedge-upgrade-info-channel-enum.md) | Query, Optional | upgrade channel to follow, stable (default) / beta / alpha<br>**Default**: `"stable"` |

## Response Type

[`[]models.MxedgeUpgradeInfoItems`](../../doc/models/mxedge-upgrade-info-items.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

channel := models.GetOrgMxedgeUpgradeInfoChannelEnum("stable")

apiResponse, err := orgsMxEdges.GetOrgMxEdgeUpgradeInfo(ctx, orgId, &channel)
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
    "default": true,
    "distro": "bullseye",
    "package": "mxagent",
    "version": "2.4.100"
  },
  {
    "distro": "bullseye",
    "package": "tunterm",
    "version": "1.0.0"
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


# List Org Mx Edges

Get List of Org MxEdges

```go
ListOrgMxEdges(
    ctx context.Context,
    orgId uuid.UUID,
    forSites *models.MxedgeForSiteEnum,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.Mxedge],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `forSites` | [`*models.MxedgeForSiteEnum`](../../doc/models/mxedge-for-site-enum.md) | Query, Optional | filter for site level mist edges<br>**Default**: `"any"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br>**Constraints**: `>= 1` |

## Response Type

[`[]models.Mxedge`](../../doc/models/mxedge.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

forSites := models.MxedgeForSiteEnum("any")

limit := 100

page := 1

apiResponse, err := orgsMxEdges.ListOrgMxEdges(ctx, orgId, &forSites, &limit, &page)
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
    "cpu_stat": {
      "cpus": {
        "cpu0": {
          "idle": 79,
          "interrupt": 0,
          "system": 4,
          "usage": 20,
          "user": 16
        },
        "cpu1": {
          "idle": 93,
          "interrupt": 0,
          "system": 4,
          "usage": 6,
          "user": 1
        }
      },
      "idle": 87,
      "interrupt": 0,
      "system": 5,
      "usage": 12,
      "user": 7
    },
    "ext_ip": "116.187.144.16",
    "id": "387804a7-3474-85ce-15a2-f9a9684c9c90",
    "ip_stat": {
      "ip": "172.16.5.3",
      "ips": {
        "ens192": "172.16.5.3/24,fe81::20c:29ff:fef8:d18e/64"
      }
    },
    "lag_stat": {
      "lag0": {
        "active_ports": [
          "0",
          "1"
        ]
      }
    },
    "last_seen": 1547437078,
    "magic": "ExNpT5gi-ADR8WTFd4EiQPY3cP5WdSoD",
    "memory_stats": {
      "active": 1061085184,
      "available": 4124860416,
      "buffers": 789495808,
      "cached": 718016512,
      "free": 2818838528,
      "inactive": 458158080,
      "swap_cached": 0,
      "swap_free": 8161062912,
      "swap_total": 8161062912,
      "total": 7947616256,
      "usage": 65
    },
    "model": "ME-S2019",
    "mxagent_registered": false,
    "mxcluster_id": "572586b7-f97b-a22b-526c-8b97a3f609c4",
    "name": "Guest",
    "num_tunnels": 31,
    "port_stat": {
      "eth0": {
        "full_duplex": true,
        "lldp_stats": {
          "mgmt_addr": "122.16.3.11",
          "port_desc": "GigabitEthernet4/0/16",
          "port_id": "\u0005Gi4/0/16",
          "system_desc": "Cisco IOS Software",
          "system_name": "ME-DC2-DIS-SW"
        },
        "rx_bytes": 2056,
        "rx_errors": 0,
        "rx_pkts": 670,
        "speed": 1000,
        "tx_bytes": 2056,
        "tx_pkts": 670,
        "up": true
      },
      "eth1": {
        "up": false
      },
      "module": {
        "up": false
      }
    },
    "status": "connected",
    "tunterm_registered": false,
    "tunterm_stat": {
      "monitoring_failed": false
    },
    "uptime": 884221,
    "version": "0.1.2",
    "virtualization_type": "VirtualizationVMware"
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


# Restart Org Mx Edge

In the case where a Mist Edge is replaced, you would need to unregister it. Which disconnects the currently the connected Mist Edge and allow another to register.

```go
RestartOrgMxEdge(
    ctx context.Context,
    orgId uuid.UUID,
    mxedgeId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mxedgeId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mxedgeId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsMxEdges.RestartOrgMxEdge(ctx, orgId, mxedgeId)
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


# Search Org Mist Edge Events

Search Org Mist Edge Events

```go
SearchOrgMistEdgeEvents(
    ctx context.Context,
    orgId uuid.UUID,
    mxedgeId *string,
    mxclusterId *string,
    mType *string,
    service *string,
    component *string,
    start *int,
    end *int,
    duration *string,
    limit *int) (
    models.ApiResponse[models.ResponseMxedgeEventsSearch],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mxedgeId` | `*string` | Query, Optional | mist edge id |
| `mxclusterId` | `*string` | Query, Optional | mist edge cluster id |
| `mType` | `*string` | Query, Optional | see [listDeviceEventsDefinitions]($e/Constants%20Events/listDeviceEventsDefinitions) |
| `service` | `*string` | Query, Optional | service running on mist edge(mxagent, tunterm etc) |
| `component` | `*string` | Query, Optional | component like PS1, PS2 |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w<br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |

## Response Type

[`models.ResponseMxedgeEventsSearch`](../../doc/models/response-mxedge-events-search.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")















duration := "10m"

limit := 100

apiResponse, err := orgsMxEdges.SearchOrgMistEdgeEvents(ctx, orgId, nil, nil, nil, nil, nil, nil, nil, &duration, &limit)
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
  "end": 1694708579,
  "limit": 10,
  "page": 3,
  "results": [
    {
      "mxcluster_id": "2815c917-58e7-472f-a190-bfd44fb58d05",
      "mxedge_id": "00000000-0000-0000-1000-020000dc585c",
      "org_id": "f2695c32-0e83-4936-b1b2-96fc88051213",
      "service": "tunterm",
      "timestamp": 1694678225.927,
      "type": "ME_SERVICE_STOPPED"
    }
  ],
  "start": 1694622179
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


# Search Org Mx Edges

Search Org Mist Edges

```go
SearchOrgMxEdges(
    ctx context.Context,
    orgId uuid.UUID,
    mxedgeId *string,
    siteId *string,
    mxclusterId *string,
    model *string,
    distro *string,
    tuntermVersion *string,
    sort *string,
    stats *bool,
    start *int,
    end *int,
    duration *string,
    limit *int,
    page *int) (
    models.ApiResponse[models.ResponseMxedgeSearch],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mxedgeId` | `*string` | Query, Optional | mist edge id |
| `siteId` | `*string` | Query, Optional | mist edge site id |
| `mxclusterId` | `*string` | Query, Optional | mist edge cluster id |
| `model` | `*string` | Query, Optional | model name |
| `distro` | `*string` | Query, Optional | debian code name(buster, bullseye) |
| `tuntermVersion` | `*string` | Query, Optional | tunterm version |
| `sort` | `*string` | Query, Optional | sort options, -prefix represents DESC order, default is -last_seen |
| `stats` | `*bool` | Query, Optional | whether to return device stats, default is false |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w<br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br>**Constraints**: `>= 1` |

## Response Type

[`models.ResponseMxedgeSearch`](../../doc/models/response-mxedge-search.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")





















duration := "10m"

limit := 100

page := 1

apiResponse, err := orgsMxEdges.SearchOrgMxEdges(ctx, orgId, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &duration, &limit, &page)
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
  "end": 1694708579,
  "limit": 10,
  "results": [
    {
      "distro": "buster",
      "last_seen": 1695151551.833,
      "model": "ME-X5",
      "mxedge_id": "00000000-0000-0000-1000-d420b0f0025d",
      "org_id": "35d96b1a-1a13-4ba8-90f5-1e78dd2a10c5",
      "tunterm_version": "0.1.2813",
      "uptime": 5662632
    }
  ],
  "start": 1694622179,
  "total": 2
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


# Unassign Org Mx Edge From Site

Unassign Org MxEdge from Site

```go
UnassignOrgMxEdgeFromSite(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.MxedgesUnassign) (
    models.ApiResponse[models.ResponseAssignSuccess],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.MxedgesUnassign`](../../doc/models/mxedges-unassign.md) | Body, Optional | Request Body |

## Response Type

[`models.ResponseAssignSuccess`](../../doc/models/response-assign-success.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.MxedgesUnassign{
    MxedgeIds:            []uuid.UUID{
        uuid.MustParse("387804a7-3474-85ce-15a2-f9a9684c9c90"),
    },
}

apiResponse, err := orgsMxEdges.UnassignOrgMxEdgeFromSite(ctx, orgId, &body)
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
  "success": [
    "5c5b350e0001"
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


# Unregister Org Mx Edge

In the case where a Mist Edge is replaced, you would need to unregister it. Which disconnects the currently the connected Mist Edge and allow another to register.

```go
UnregisterOrgMxEdge(
    ctx context.Context,
    orgId uuid.UUID,
    mxedgeId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mxedgeId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mxedgeId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsMxEdges.UnregisterOrgMxEdge(ctx, orgId, mxedgeId)
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


# Update Org Mx Edge

Update Org MxEdge

```go
UpdateOrgMxEdge(
    ctx context.Context,
    orgId uuid.UUID,
    mxedgeId uuid.UUID,
    body *models.Mxedge) (
    models.ApiResponse[models.Mxedge],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mxedgeId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Mxedge`](../../doc/models/mxedge.md) | Body, Optional | Request Body |

## Response Type

[`models.Mxedge`](../../doc/models/mxedge.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mxedgeId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Mxedge{
    Model:                     "ME-X1",
    Name:                      "string",
    NtpServers:                []string{
        "string",
    },
    Services:                  []string{
        "tunterm",
    },
    TuntermIpConfig:           models.ToPointer(models.MxedgeTuntermIpConfig{
        Gateway:              "string",
        Ip:                   "string",
        Netmask:              "string",
        AdditionalProperties: map[string]interface{}{
            "dns": interface{}("string"),
            "dns_suffix": interface{}("string"),
        },
    }),
    TuntermPortConfig:         models.ToPointer(models.TuntermPortConfig{
        DownstreamPorts:            []string{
            "string",
        },
        SeparateUpstreamDownstream: models.ToPointer(true),
        UpstreamPortVlanId:         models.ToPointer(1),
        UpstreamPorts:              []string{
            "string",
        },
    }),
}

apiResponse, err := orgsMxEdges.UpdateOrgMxEdge(ctx, orgId, mxedgeId, &body)
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
  "id": "95ddd29a-6a3c-929e-a431-51a5b09f36a6",
  "magic": "L-NpT5gi-ADR8WTFd4EiQPY3cP5WdSoD",
  "model": "ME-100",
  "mxagent_registered": true,
  "mxcluster_id": "572586b7-f97b-a22b-526c-8b97a3f609c4",
  "mxedge_mgmt": {
    "mist_password": "MIST_PASSWORD",
    "root_password": "ROOT_PASSWORD"
  },
  "name": "Guest",
  "ntp_servers": [],
  "oob_ip_config": {
    "dns": [
      "8.8.8.8",
      "4.4.4.4"
    ],
    "gateway": "10.2.1.254",
    "ip": "10.2.1.10",
    "netmask": "255.255.255.0",
    "type": "static"
  },
  "tunterm_dhcpd_config": {
    "2": {
      "enabled": true,
      "servers": [
        "11.2.3.44"
      ]
    },
    "enabled": false,
    "servers": [
      "11.2.3.4"
    ]
  },
  "tunterm_extra_routes": {
    "11.0.0.0/8": {
      "via": "10.3.3.1"
    }
  },
  "tunterm_ip_config": {
    "dns": [
      "8.8.8.8"
    ],
    "dns_suffix": [
      ".mist.local"
    ],
    "gateway": "10.2.1.254",
    "ip": "10.2.1.1",
    "netmask": "255.255.255.0"
  }
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


# Upload Org Mx Edge Support Files

Support / Upload Mist Edge support files

```go
UploadOrgMxEdgeSupportFiles(
    ctx context.Context,
    orgId uuid.UUID,
    mxedgeId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mxedgeId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mxedgeId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsMxEdges.UploadOrgMxEdgeSupportFiles(ctx, orgId, mxedgeId)
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

