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
* [Disconnect Org Mx Edge Tunterm Aps](../../doc/controllers/orgs-mx-edges.md#disconnect-org-mx-edge-tunterm-aps)
* [Get Org Mx Edge](../../doc/controllers/orgs-mx-edges.md#get-org-mx-edge)
* [Get Org Mx Edge Upgrade Info](../../doc/controllers/orgs-mx-edges.md#get-org-mx-edge-upgrade-info)
* [Get Org Mx Edge Vm Params](../../doc/controllers/orgs-mx-edges.md#get-org-mx-edge-vm-params)
* [List Org Mx Edges](../../doc/controllers/orgs-mx-edges.md#list-org-mx-edges)
* [Restart Org Mx Edge](../../doc/controllers/orgs-mx-edges.md#restart-org-mx-edge)
* [Search Org Mist Edge Events](../../doc/controllers/orgs-mx-edges.md#search-org-mist-edge-events)
* [Search Org Mx Edges](../../doc/controllers/orgs-mx-edges.md#search-org-mx-edges)
* [Unassign Org Mx Edge from Site](../../doc/controllers/orgs-mx-edges.md#unassign-org-mx-edge-from-site)
* [Unregister Org Mx Edge](../../doc/controllers/orgs-mx-edges.md#unregister-org-mx-edge)
* [Update Org Mx Edge](../../doc/controllers/orgs-mx-edges.md#update-org-mx-edge)
* [Upload Org Mx Edge Support Files](../../doc/controllers/orgs-mx-edges.md#upload-org-mx-edge-support-files)


# Add Org Mx Edge Image

Upload and attach an image file to a Mist Edge appliance. A Mist Edge can have up to three image attachments.

```go
AddOrgMxEdgeImage(
    ctx context.Context,
    orgId uuid.UUID,
    mxedgeId uuid.UUID,
    imageNumber int,
    file models.FileWrapper,
    json *string) (
    http.Response,
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mxedgeId` | `uuid.UUID` | Template, Required | - |
| `imageNumber` | `int` | Template, Required | **Constraints**: `>= 1`, `<= 3` |
| `file` | `models.FileWrapper` | Form, Required | Image file content uploaded as multipart form data |
| `json` | `*string` | Form, Optional | Optional JSON metadata submitted with the image upload |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mxedgeId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

imageNumber := 3

file := getFile("dummy_file", func(err error) { log.Fatalln(err) })

resp, err := orgsMxEdges.AddOrgMxEdgeImage(ctx, orgId, mxedgeId, imageNumber, file, nil)
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


# Assign Org Mx Edge to Site

Assign one or more Mist Edge appliances from the organization to a site by Mist Edge ID and site ID.

```go
AssignOrgMxEdgeToSite(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.MxedgesAssign) (
    models.ApiResponse[models.ResponseAssignSuccess],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.MxedgesAssign`](../../doc/models/mxedges-assign.md) | Body, Optional | Request Body |

## Response Type

**200**: OK - list only devices that has deviceprofile_id changed

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseAssignSuccess](../../doc/models/response-assign-success.md).

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
  "success": [
    "5c5b350e0001"
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


# Bounce Org Mx Edge Data Ports

Bounce one or more TunTerm data ports on a Mist Edge, optionally setting the hold time between port bounces.

```go
BounceOrgMxEdgeDataPorts(
    ctx context.Context,
    orgId uuid.UUID,
    mxedgeId uuid.UUID,
    body *models.UtilsTuntermBouncePort) (
    http.Response,
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mxedgeId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.UtilsTuntermBouncePort`](../../doc/models/utils-tunterm-bounce-port.md) | Body, Optional | - |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

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


# Claim Org Mx Edge

Claim one or more Mist Edge appliances into the organization using their claim codes.

```go
ClaimOrgMxEdge(
    ctx context.Context,
    orgId uuid.UUID,
    body []string) (
    models.ApiResponse[models.ResponseClaimMxEdge],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | `[]string` | Body, Optional | Request Body<br><br>**Constraints**: *Unique Items Required* |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseClaimMxEdge](../../doc/models/response-claim-mx-edge.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := []string{
    "6JG8E-PTFV2-A9Z2N",
    "DVH4V-SNMSZ-PDXBR",
}

apiResponse, err := orgsMxEdges.ClaimOrgMxEdge(ctx, orgId, body)
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
  "id": "95ddd29a-6a3c-929e-a431-51a5b09f36a6",
  "magic": "ExNpT5gi-ADR8WTFd4EiQPY3cP5WdSoD"
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


# Control Org Mx Edge Services

Start, stop, or restart a named Mist Edge service such as `tunterm`, `mxagent`, or `radsecproxy`.

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

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mxedgeId` | `uuid.UUID` | Template, Required | - |
| `name` | [`models.MxedgeServiceNameEnum`](../../doc/models/mxedge-service-name-enum.md) | Template, Required | enum: `mxagent`, `mxdas`, `mxnacedge`, `mxocproxy`, `radsecproxy`, `tunterm` |
| `action` | [`models.MxedgeServiceActionEnum`](../../doc/models/mxedge-service-action-enum.md) | Template, Required | Restart or start or stop |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mxedgeId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

name := models.MxedgeServiceNameEnum_MXNACEDGE

action := models.MxedgeServiceActionEnum_RESTART

resp, err := orgsMxEdges.ControlOrgMxEdgeServices(ctx, orgId, mxedgeId, name, action)
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


# Count Org Mx Edges

Count organization Mist Edge records, optionally grouped by `distinct` and filtered by Mist Edge, cluster, site, model, distro, tunnel termination version, and time range.

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
    start *string,
    end *string,
    duration *string,
    limit *int) (
    models.ApiResponse[models.ResponseCount],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.OrgMxedgeCountDistinctEnum`](../../doc/models/org-mxedge-count-distinct-enum.md) | Query, Optional | Field used to group this count response. enum: `distro`, `model`, `mxcluster_id`, `site_id`, `tunterm_version`<br><br>**Default**: `"model"` |
| `mxedgeId` | `*string` | Query, Optional | Filter results by Mist Edge identifier |
| `siteId` | `*string` | Query, Optional | Mist edge site id |
| `mxclusterId` | `*string` | Query, Optional | Mist edge cluster id |
| `model` | `*string` | Query, Optional | Filter results by device model |
| `distro` | `*string` | Query, Optional | Debian code name (buster, bullseye) |
| `tuntermVersion` | `*string` | Query, Optional | Filter results by tunnel termination version |
| `sort` | `*string` | Query, Optional | Field used to sort results |
| `stats` | `*bool` | Query, Optional | Whether to return device stats, default is false |
| `start` | `*string` | Query, Optional | Lower bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d` or `-1w` |
| `end` | `*string` | Query, Optional | Upper bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d`, `-2h`, or `now` |
| `duration` | `*string` | Query, Optional | Time range duration for the query, using relative units such as `10m`, `7d`, or `2w`<br><br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |

## Response Type

**200**: Result of Count

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseCount](../../doc/models/response-count.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

distinct := models.OrgMxedgeCountDistinctEnum_MODEL

duration := "10m"

limit := 100

apiResponse, err := orgsMxEdges.CountOrgMxEdges(ctx, orgId, &distinct, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &duration, &limit)
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# Count Org Site Mx Edge Events

Count Mist Edge event records across the organization, optionally grouped by `distinct` and filtered by Mist Edge, cluster, event type, service, and time range.

```go
CountOrgSiteMxEdgeEvents(
    ctx context.Context,
    orgId uuid.UUID,
    distinct *models.OrgMxedgeEventsCountDistinctEnum,
    mxedgeId *string,
    mxclusterId *string,
    mType *string,
    service *string,
    start *string,
    end *string,
    duration *string,
    limit *int) (
    models.ApiResponse[models.ResponseCount],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.OrgMxedgeEventsCountDistinctEnum`](../../doc/models/org-mxedge-events-count-distinct-enum.md) | Query, Optional | Field used to group this count response. enum: `mxcluster_id`, `mxedge_id`, `package`, `type`<br><br>**Default**: `"mxedge_id"` |
| `mxedgeId` | `*string` | Query, Optional | Filter results by Mist Edge identifier |
| `mxclusterId` | `*string` | Query, Optional | Mist edge cluster id |
| `mType` | `*string` | Query, Optional | See [List Device Events Definitions](../../doc/controllers/constants-events.md#list-device-events-definitions) |
| `service` | `*string` | Query, Optional | Filter results by service name |
| `start` | `*string` | Query, Optional | Lower bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d` or `-1w` |
| `end` | `*string` | Query, Optional | Upper bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d`, `-2h`, or `now` |
| `duration` | `*string` | Query, Optional | Time range duration for the query, using relative units such as `10m`, `7d`, or `2w`<br><br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |

## Response Type

**200**: Result of Count

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseCount](../../doc/models/response-count.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

distinct := models.OrgMxedgeEventsCountDistinctEnum_MXEDGEID

duration := "10m"

limit := 100

apiResponse, err := orgsMxEdges.CountOrgSiteMxEdgeEvents(ctx, orgId, &distinct, nil, nil, nil, nil, nil, nil, &duration, &limit)
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# Create Org Mx Edge

Create a Mist Edge appliance configuration in the organization, including cluster assignment, management, services, and tunnel termination settings.

```go
CreateOrgMxEdge(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.Mxedge) (
    models.ApiResponse[models.Mxedge],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Mxedge`](../../doc/models/mxedge.md) | Body, Optional | Request Body |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Mxedge](../../doc/models/mxedge.md).

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
        UpstreamPortVlanId:         models.ToPointer(models.TuntermPortConfigUpstreamPortVlanIdContainer.FromNumber(1)),
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
                VlanIds:              []models.VlanIdWithVariable{
                    models.VlanIdWithVariableContainer.FromNumber(5),
                    models.VlanIdWithVariableContainer.FromNumber(2),
                    models.VlanIdWithVariableContainer.FromNumber(3),
                },
            },
        },
    }),
}

apiResponse, err := orgsMxEdges.CreateOrgMxEdge(ctx, orgId, &body)
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# Delete Org Mx Edge

Delete a Mist Edge appliance record from the organization by Mist Edge ID.

```go
DeleteOrgMxEdge(
    ctx context.Context,
    orgId uuid.UUID,
    mxedgeId uuid.UUID) (
    http.Response,
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mxedgeId` | `uuid.UUID` | Template, Required | - |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mxedgeId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsMxEdges.DeleteOrgMxEdge(ctx, orgId, mxedgeId)
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


# Delete Org Mx Edge Image

Delete a numbered image attachment from a Mist Edge appliance.

```go
DeleteOrgMxEdgeImage(
    ctx context.Context,
    orgId uuid.UUID,
    mxedgeId uuid.UUID,
    imageNumber int) (
    http.Response,
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mxedgeId` | `uuid.UUID` | Template, Required | - |
| `imageNumber` | `int` | Template, Required | **Constraints**: `>= 1`, `<= 3` |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mxedgeId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

imageNumber := 3

resp, err := orgsMxEdges.DeleteOrgMxEdgeImage(ctx, orgId, mxedgeId, imageNumber)
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


# Disconnect Org Mx Edge Tunterm Aps

Disconnect specific APs from the Mist Edge TunTerm service by AP MAC address.

```go
DisconnectOrgMxEdgeTuntermAps(
    ctx context.Context,
    orgId uuid.UUID,
    mxedgeId uuid.UUID,
    body *models.MacAddresses) (
    http.Response,
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mxedgeId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.MacAddresses`](../../doc/models/mac-addresses.md) | Body, Optional | - |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mxedgeId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.MacAddresses{
    Macs:                 []string{
        "5c5b353e4eb1",
        "5c5b353e4eb2",
    },
}

resp, err := orgsMxEdges.DisconnectOrgMxEdgeTuntermAps(ctx, orgId, mxedgeId, &body)
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


# Get Org Mx Edge

Retrieve configuration and registration details for a specific Mist Edge appliance in the organization.

```go
GetOrgMxEdge(
    ctx context.Context,
    orgId uuid.UUID,
    mxedgeId uuid.UUID) (
    models.ApiResponse[models.Mxedge],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mxedgeId` | `uuid.UUID` | Template, Required | - |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Mxedge](../../doc/models/mxedge.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mxedgeId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsMxEdges.GetOrgMxEdge(ctx, orgId, mxedgeId)
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# Get Org Mx Edge Upgrade Info

Retrieve available Mist Edge package versions by upgrade channel and distro.

```go
GetOrgMxEdgeUpgradeInfo(
    ctx context.Context,
    orgId uuid.UUID,
    channel *models.GetOrgMxedgeUpgradeInfoChannelEnum,
    distro *string) (
    models.ApiResponse[[]models.MxedgeUpgradeInfoItems],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `channel` | [`*models.GetOrgMxedgeUpgradeInfoChannelEnum`](../../doc/models/get-org-mxedge-upgrade-info-channel-enum.md) | Query, Optional | Upgrade channel used to filter available versions. enum: `alpha`, `beta`, `stable`<br><br>**Default**: `"stable"` |
| `distro` | `*string` | Query, Optional | Filter results by distro |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.MxedgeUpgradeInfoItems](../../doc/models/mxedge-upgrade-info-items.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

channel := models.GetOrgMxedgeUpgradeInfoChannelEnum_STABLE

apiResponse, err := orgsMxEdges.GetOrgMxEdgeUpgradeInfo(ctx, orgId, &channel, nil)
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# Get Org Mx Edge Vm Params

Retrieve VM deployment parameters for a Mist Edge, including model, optional name, and base64 user data.

```go
GetOrgMxEdgeVmParams(
    ctx context.Context,
    orgId uuid.UUID,
    mxedgeId uuid.UUID) (
    models.ApiResponse[models.MxedgeVmParams],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mxedgeId` | `uuid.UUID` | Template, Required | - |

## Response Type

**200**: Mist Edge VM Parameters

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.MxedgeVmParams](../../doc/models/mxedge-vm-params.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mxedgeId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsMxEdges.GetOrgMxEdgeVmParams(ctx, orgId, mxedgeId)
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


# List Org Mx Edges

List Mist Edge appliances in the organization, optionally filtering for org-level, site-level, or all Mist Edges with `for_site`.

```go
ListOrgMxEdges(
    ctx context.Context,
    orgId uuid.UUID,
    forSite *models.MxedgeForSiteEnum,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.Mxedge],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `forSite` | [`*models.MxedgeForSiteEnum`](../../doc/models/mxedge-for-site-enum.md) | Query, Optional | Filter for org/site level Mist Edges. enum: `any`, `false`, `true`<br><br>**Default**: `"any"` |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | Select the page number to return when using page-based pagination; starts at `1`<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1` |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.Mxedge](../../doc/models/mxedge.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

forSite := models.MxedgeForSiteEnum_ANY

limit := 100

page := 1

apiResponse, err := orgsMxEdges.ListOrgMxEdges(ctx, orgId, &forSite, &limit, &page)
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# Restart Org Mx Edge

Restart the registration workflow for a Mist Edge replacement by disconnecting the currently registered appliance so another Mist Edge can register.

```go
RestartOrgMxEdge(
    ctx context.Context,
    orgId uuid.UUID,
    mxedgeId uuid.UUID) (
    http.Response,
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mxedgeId` | `uuid.UUID` | Template, Required | - |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mxedgeId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsMxEdges.RestartOrgMxEdge(ctx, orgId, mxedgeId)
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


# Search Org Mist Edge Events

Search Mist Edge event records across the organization with filters for Mist Edge, cluster, event type, service, component, and time range.

```go
SearchOrgMistEdgeEvents(
    ctx context.Context,
    orgId uuid.UUID,
    mxedgeId *string,
    mxclusterId *string,
    mType *string,
    service *string,
    component *string,
    limit *int,
    start *string,
    end *string,
    duration *string,
    sort *string,
    searchAfter *string) (
    models.ApiResponse[models.ResponseMxedgeEventsSearch],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mxedgeId` | `*string` | Query, Optional | Filter results by Mist Edge identifier |
| `mxclusterId` | `*string` | Query, Optional | Mist edge cluster id |
| `mType` | `*string` | Query, Optional | See [List Device Events Definitions](../../doc/controllers/constants-events.md#list-device-events-definitions) |
| `service` | `*string` | Query, Optional | Filter results by service name |
| `component` | `*string` | Query, Optional | Filter results by component name |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `start` | `*string` | Query, Optional | Lower bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d` or `-1w` |
| `end` | `*string` | Query, Optional | Upper bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d`, `-2h`, or `now` |
| `duration` | `*string` | Query, Optional | Time range duration for the query, using relative units such as `10m`, `7d`, or `2w`<br><br>**Default**: `"1d"` |
| `sort` | `*string` | Query, Optional | On which field the list should be sorted, -prefix represents DESC order<br><br>**Default**: `"timestamp"` |
| `searchAfter` | `*string` | Query, Optional | Pagination cursor for retrieving subsequent pages of results. This value is automatically populated by Mist in the `next` URL from the previous response and should not be manually constructed. |

## Response Type

**200**: Example response

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseMxedgeEventsSearch](../../doc/models/response-mxedge-events-search.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

limit := 100

duration := "10m"

sort := "-site_id"

apiResponse, err := orgsMxEdges.SearchOrgMistEdgeEvents(ctx, orgId, nil, nil, nil, nil, nil, &limit, nil, nil, &duration, &sort, nil)
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# Search Org Mx Edges

Search organization Mist Edge records with filters for hostname, Mist Edge, cluster, site, model, distro, tunnel termination version, and time range.

```go
SearchOrgMxEdges(
    ctx context.Context,
    orgId uuid.UUID,
    hostname *string,
    mxedgeId *string,
    mxclusterId *string,
    model *string,
    distro *string,
    tuntermVersion *string,
    siteId *string,
    stats *bool,
    limit *int,
    start *string,
    end *string,
    duration *string,
    sort *string,
    searchAfter *string) (
    models.ApiResponse[models.ResponseMxedgeSearch],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `hostname` | `*string` | Query, Optional | Partial / full Device hostname. Use `prefix*` for prefix search or `*substring*` for contains search (e.g. `my-london*` and `*london*` match `my-london-1`). Suffix-only wildcards (e.g. `*london-1`) are not supported. Accepts multiple comma-separated values. |
| `mxedgeId` | `*string` | Query, Optional | Filter results by Mist Edge identifier. Accepts multiple comma-separated values. |
| `mxclusterId` | `*string` | Query, Optional | Mist edge cluster id. Accepts multiple comma-separated values. |
| `model` | `*string` | Query, Optional | Partial / full Device model. Use `prefix*` for prefix search or `*substring*` for contains search (e.g. `AP4*` and `*P4*` match `AP43`). Suffix-only wildcards (e.g. `*43`) are not supported. Accepts multiple comma-separated values. |
| `distro` | `*string` | Query, Optional | Debian code name (buster, bullseye) |
| `tuntermVersion` | `*string` | Query, Optional | Filter results by tunnel termination version |
| `siteId` | `*string` | Query, Optional | Mist edge site id |
| `stats` | `*bool` | Query, Optional | Whether to return device stats, default is false |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `start` | `*string` | Query, Optional | Lower bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d` or `-1w` |
| `end` | `*string` | Query, Optional | Upper bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d`, `-2h`, or `now` |
| `duration` | `*string` | Query, Optional | Time range duration for the query, using relative units such as `10m`, `7d`, or `2w`<br><br>**Default**: `"1d"` |
| `sort` | `*string` | Query, Optional | On which field the list should be sorted, -prefix represents DESC order<br><br>**Default**: `"timestamp"` |
| `searchAfter` | `*string` | Query, Optional | Pagination cursor for retrieving subsequent pages of results. This value is automatically populated by Mist in the `next` URL from the previous response and should not be manually constructed. |

## Response Type

**200**: Example response

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseMxedgeSearch](../../doc/models/response-mxedge-search.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

hostname := "my-london-1,my-london*"

mxedgeId := "00000000-0000-0000-0000-000000000001,00000000-0000-0000-0000-000000000002"

mxclusterId := "00000000-0000-0000-0000-000000000001,00000000-0000-0000-0000-000000000002"

model := "AP43,AP4*"

limit := 100

duration := "10m"

sort := "-site_id"

apiResponse, err := orgsMxEdges.SearchOrgMxEdges(ctx, orgId, &hostname, &mxedgeId, &mxclusterId, &model, nil, nil, nil, nil, &limit, nil, nil, &duration, &sort, nil)
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# Unassign Org Mx Edge from Site

Unassign one or more Mist Edge appliances from their current site while keeping them in the organization.

```go
UnassignOrgMxEdgeFromSite(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.MxedgesUnassign) (
    models.ApiResponse[models.ResponseAssignSuccess],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.MxedgesUnassign`](../../doc/models/mxedges-unassign.md) | Body, Optional | Request Body |

## Response Type

**200**: OK - list only devices that has deviceprofile_id changed

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseAssignSuccess](../../doc/models/response-assign-success.md).

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
  "success": [
    "5c5b350e0001"
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


# Unregister Org Mx Edge

Unregister a Mist Edge during a replacement workflow by disconnecting the currently registered appliance so another Mist Edge can register.

```go
UnregisterOrgMxEdge(
    ctx context.Context,
    orgId uuid.UUID,
    mxedgeId uuid.UUID) (
    http.Response,
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mxedgeId` | `uuid.UUID` | Template, Required | - |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mxedgeId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsMxEdges.UnregisterOrgMxEdge(ctx, orgId, mxedgeId)
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


# Update Org Mx Edge

Update a Mist Edge appliance configuration, including model, name, management IP, services, and tunnel termination settings.

```go
UpdateOrgMxEdge(
    ctx context.Context,
    orgId uuid.UUID,
    mxedgeId uuid.UUID,
    body *models.Mxedge) (
    models.ApiResponse[models.Mxedge],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mxedgeId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Mxedge`](../../doc/models/mxedge.md) | Body, Optional | Request Body |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Mxedge](../../doc/models/mxedge.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mxedgeId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Mxedge{
    Model:                     "ME-X1",
    Name:                      "me-gc1-01",
    OobIpConfig:               models.ToPointer(models.MxedgeOobIpConfig{
        Dns:                  []string{
            "8.8.8.8",
            "1.1.1.1",
        },
        Gateway:              models.ToPointer("10.3.172.9"),
        Ip:                   models.ToPointer("10.3.172.201"),
        Netmask:              models.ToPointer("/24"),
        Type:                 models.ToPointer(models.IpTypeEnum_STATIC),
    }),
    Services:                  []string{
        "tunterm",
    },
    TuntermIpConfig:           models.ToPointer(models.MxedgeTuntermIpConfig{
        Gateway:              "10.10.172.2",
        Ip:                   "10.10.172.201",
        Netmask:              "/24",
    }),
    TuntermPortConfig:         models.ToPointer(models.TuntermPortConfig{
        DownstreamPorts:            []string{
            "0",
        },
        SeparateUpstreamDownstream: models.ToPointer(true),
        UpstreamPortVlanId:         models.ToPointer(models.TuntermPortConfigUpstreamPortVlanIdContainer.FromString("1010")),
        UpstreamPorts:              []string{
            "1",
        },
    }),
}

apiResponse, err := orgsMxEdges.UpdateOrgMxEdge(ctx, orgId, mxedgeId, &body)
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# Upload Org Mx Edge Support Files

Trigger upload of support files from a Mist Edge for troubleshooting.

```go
UploadOrgMxEdgeSupportFiles(
    ctx context.Context,
    orgId uuid.UUID,
    mxedgeId uuid.UUID) (
    http.Response,
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mxedgeId` | `uuid.UUID` | Template, Required | - |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mxedgeId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsMxEdges.UploadOrgMxEdgeSupportFiles(ctx, orgId, mxedgeId)
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

