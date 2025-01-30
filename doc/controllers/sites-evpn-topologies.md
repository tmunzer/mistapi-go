# Sites EVPN Topologies

```go
sitesEVPNTopologies := client.SitesEVPNTopologies()
```

## Class Name

`SitesEVPNTopologies`

## Methods

* [Create Site Evpn Topology](../../doc/controllers/sites-evpn-topologies.md#create-site-evpn-topology)
* [Delete Site Evpn Topology](../../doc/controllers/sites-evpn-topologies.md#delete-site-evpn-topology)
* [Get Site Evpn Topology](../../doc/controllers/sites-evpn-topologies.md#get-site-evpn-topology)
* [List Site Evpn Topologies](../../doc/controllers/sites-evpn-topologies.md#list-site-evpn-topologies)
* [Update Site Evpn Topology](../../doc/controllers/sites-evpn-topologies.md#update-site-evpn-topology)


# Create Site Evpn Topology

While all the `evpn_id` / `downlink_ips` can be specifidd by hand, the easiest way is to call the `build_vpn_topology` API, allowing you to examine the diff, and update it yourself. You can also simply call it with `overwrite=true` which will apply the updates for you.

**Notes:**

1. You can use `core` / `distribution` / `access` to create a CLOS topology
2. You can also use `core` / `distribution` to form a 2-tier EVPN topology where ESI-Lag is configured distribution to connect to access switches
3. In a small/medium campus, `collapsed-core` can be used where core switches are the inter-connected to do EVPN
4. The API uses a few pre-defined parameters and best-practices to generate the configs. It can be customized by using `evpn_options` in Site Setting / Network Template. (e.g. a different subnet for the underlay)

#### Collapsed Core

In a small-medium campus, EVPN can also be enabled only at the core switches (up to 4) by assigning all participating switches with `collapsed-core role`. When there are more than 2 switches, a ring-like topology will be formed.

#### ESI-Lag

If the access switchess does not have EVPN support, you can take advantage of EVPN by setting up ESI-Lag on distribution switches

#### Leaf / Access / Collapsed-Core

For leaf nodes in a EVPN topology, you’d have to configure the IPs for networks that would participate in EVPN. Optionally, VRFs to isolate traffic from one tenant verus another

```go
CreateSiteEvpnTopology(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.EvpnTopology) (
    models.ApiResponse[models.EvpnTopology],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.EvpnTopology`](../../doc/models/evpn-topology.md) | Body, Optional | - |

## Response Type

[`models.EvpnTopology`](../../doc/models/evpn-topology.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.EvpnTopology{
    Name:                 models.ToPointer("CC"),
    Overwrite:            models.ToPointer(true),
    PodNames:             map[string]string{
        "1": "default",
        "2": "default",
    },
    Switches:             []models.EvpnTopologySwitch{
        models.EvpnTopologySwitch{
            Mac:                  "5c5b35000003",
            Role:                 models.EvpnTopologySwitchRoleEnum_COLLAPSEDCORE,
        },
        models.EvpnTopologySwitch{
            Mac:                  "5c5b35000004",
            Role:                 models.EvpnTopologySwitchRoleEnum_COLLAPSEDCORE,
        },
    },
}

apiResponse, err := sitesEVPNTopologies.CreateSiteEvpnTopology(ctx, siteId, &body)
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
  "id": "9197ec96-4c8d-529f-c595-035895e688b2",
  "name": "CC",
  "overwrite": true,
  "pod_names": {
    "1": "default",
    "2": "default"
  },
  "switches": [
    {
      "deviceprofile_id": "6a1deab1-96df-4fa2-8455-d5253f943d06",
      "downlink_ips": [
        "10.255.240.6",
        "10.255.240.8"
      ],
      "downlinks": [
        "5c5b35000007",
        "5c5b35000008"
      ],
      "esilaglinks": [
        "5c5b3500000f"
      ],
      "evpn_id": 1,
      "mac": "5c5b35000003",
      "model": "QFX10002-36Q",
      "role": "collapsed-core",
      "site_id": "1916d52a-4a90-11e5-8b45-1258369c38a9",
      "uplinks": [
        "5c5b35000005",
        "5c5b35000006"
      ]
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


# Delete Site Evpn Topology

Delete the site EVPN Topology

```go
DeleteSiteEvpnTopology(
    ctx context.Context,
    siteId uuid.UUID,
    evpnTopologyId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `evpnTopologyId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

evpnTopologyId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := sitesEVPNTopologies.DeleteSiteEvpnTopology(ctx, siteId, evpnTopologyId)
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


# Get Site Evpn Topology

Get One EVPN Topology Detail

```go
GetSiteEvpnTopology(
    ctx context.Context,
    siteId uuid.UUID,
    evpnTopologyId uuid.UUID) (
    models.ApiResponse[models.EvpnTopology],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `evpnTopologyId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.EvpnTopology`](../../doc/models/evpn-topology.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

evpnTopologyId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := sitesEVPNTopologies.GetSiteEvpnTopology(ctx, siteId, evpnTopologyId)
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
  "id": "9197ec96-4c8d-529f-c595-035895e688b2",
  "name": "CC",
  "overwrite": true,
  "pod_names": {
    "1": "default",
    "2": "default"
  },
  "switches": [
    {
      "deviceprofile_id": "6a1deab1-96df-4fa2-8455-d5253f943d06",
      "downlink_ips": [
        "10.255.240.6",
        "10.255.240.8"
      ],
      "downlinks": [
        "5c5b35000007",
        "5c5b35000008"
      ],
      "esilaglinks": [
        "5c5b3500000f"
      ],
      "evpn_id": 1,
      "mac": "5c5b35000003",
      "model": "QFX10002-36Q",
      "role": "collapsed-core",
      "site_id": "1916d52a-4a90-11e5-8b45-1258369c38a9",
      "uplinks": [
        "5c5b35000005",
        "5c5b35000006"
      ]
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


# List Site Evpn Topologies

Get the existing EVPN topology

```go
ListSiteEvpnTopologies(
    ctx context.Context,
    siteId uuid.UUID) (
    models.ApiResponse[models.EvpnTopology],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.EvpnTopology`](../../doc/models/evpn-topology.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := sitesEVPNTopologies.ListSiteEvpnTopologies(ctx, siteId)
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
  "id": "9197ec96-4c8d-529f-c595-035895e688b2",
  "name": "CC",
  "overwrite": true,
  "pod_names": {
    "1": "default",
    "2": "default"
  },
  "switches": [
    {
      "deviceprofile_id": "6a1deab1-96df-4fa2-8455-d5253f943d06",
      "downlink_ips": [
        "10.255.240.6",
        "10.255.240.8"
      ],
      "downlinks": [
        "5c5b35000007",
        "5c5b35000008"
      ],
      "esilaglinks": [
        "5c5b3500000f"
      ],
      "evpn_id": 1,
      "mac": "5c5b35000003",
      "model": "QFX10002-36Q",
      "role": "collapsed-core",
      "site_id": "1916d52a-4a90-11e5-8b45-1258369c38a9",
      "uplinks": [
        "5c5b35000005",
        "5c5b35000006"
      ]
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


# Update Site Evpn Topology

Update the EVPN Topolgy

```go
UpdateSiteEvpnTopology(
    ctx context.Context,
    siteId uuid.UUID,
    evpnTopologyId uuid.UUID,
    body *models.EvpnTopology) (
    models.ApiResponse[models.EvpnTopology],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `evpnTopologyId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.EvpnTopology`](../../doc/models/evpn-topology.md) | Body, Optional | - |

## Response Type

[`models.EvpnTopology`](../../doc/models/evpn-topology.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

evpnTopologyId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.EvpnTopology{
    Overwrite:            models.ToPointer(false),
    Switches:             []models.EvpnTopologySwitch{
        models.EvpnTopologySwitch{
            Mac:                  "5c5b35000003",
            Role:                 models.EvpnTopologySwitchRoleEnum_COLLAPSEDCORE,
        },
        models.EvpnTopologySwitch{
            Mac:                  "5c5b35000004",
            Role:                 models.EvpnTopologySwitchRoleEnum_NONE,
        },
    },
}

apiResponse, err := sitesEVPNTopologies.UpdateSiteEvpnTopology(ctx, siteId, evpnTopologyId, &body)
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
  "id": "9197ec96-4c8d-529f-c595-035895e688b2",
  "name": "CC",
  "overwrite": true,
  "pod_names": {
    "1": "default",
    "2": "default"
  },
  "switches": [
    {
      "deviceprofile_id": "6a1deab1-96df-4fa2-8455-d5253f943d06",
      "downlink_ips": [
        "10.255.240.6",
        "10.255.240.8"
      ],
      "downlinks": [
        "5c5b35000007",
        "5c5b35000008"
      ],
      "esilaglinks": [
        "5c5b3500000f"
      ],
      "evpn_id": 1,
      "mac": "5c5b35000003",
      "model": "QFX10002-36Q",
      "role": "collapsed-core",
      "site_id": "1916d52a-4a90-11e5-8b45-1258369c38a9",
      "uplinks": [
        "5c5b35000005",
        "5c5b35000006"
      ]
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

