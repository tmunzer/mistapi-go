# Orgs EVPN Topologies

```go
orgsEVPNTopologies := client.OrgsEVPNTopologies()
```

## Class Name

`OrgsEVPNTopologies`

## Methods

* [Create Org Evpn Topology](../../doc/controllers/orgs-evpn-topologies.md#create-org-evpn-topology)
* [Delete Org Evpn Topology](../../doc/controllers/orgs-evpn-topologies.md#delete-org-evpn-topology)
* [Get Org Evpn Topology](../../doc/controllers/orgs-evpn-topologies.md#get-org-evpn-topology)
* [List Org Evpn Topologies](../../doc/controllers/orgs-evpn-topologies.md#list-org-evpn-topologies)
* [Update Org Evpn Topology](../../doc/controllers/orgs-evpn-topologies.md#update-org-evpn-topology)


# Create Org Evpn Topology

While all the `evpn_id` / `downlink_ips` can be specified by hand, the easiest way is to call the `build_vpn_topology` API, allowing you to examine the diff, and update it yourself. You can also simply call it with `overwrite=true` which will apply the updates for you.

**Notes:**

1. You can use `core` / `distribution` / `access` to create a CLOS topology
2. You can also use `core` / `distribution` to form a 2-tier EVPN topology where ESI-Lag is configured distribution to connect to access switches
3. In a small/medium campus, `collapsed-core` can be used where core switches are the inter-connected to do EVPN
4. The API uses a few pre-defined parameters and best-practices to generate the configs. It can be customized by using `evpn_options` in Site Setting / Network Template. (e.g. a different subnet for the underlay)

#### Collapsed Core

In a small-medium campus, EVPN can also be enabled only at the core switches (up to 4) by assigning all participating switches with `collapsed-core role`. When there are more than 2 switches, a ring-like topology will be formed.

#### ESI-Lag

If the access switches does not have EVPN support, you can take advantage of EVPN by setting up ESI-Lag on distribution switches

#### Leaf / Access / Collapsed-Core

For leaf nodes in a EVPN topology, you’d have to configure the IPs for networks that would participate in EVPN. Optionally, VRFs to isolate traffic from one tenant verus another

```go
CreateOrgEvpnTopology(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.EvpnTopology) (
    models.ApiResponse[models.EvpnTopology],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.EvpnTopology`](../../doc/models/evpn-topology.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.EvpnTopology](../../doc/models/evpn-topology.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.EvpnTopology{
    Name:                 models.ToPointer("CC"),
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

apiResponse, err := orgsEVPNTopologies.CreateOrgEvpnTopology(ctx, orgId, &body)
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


# Delete Org Evpn Topology

Delete the Org EVPN Topology

```go
DeleteOrgEvpnTopology(
    ctx context.Context,
    orgId uuid.UUID,
    evpnTopologyId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `evpnTopologyId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

evpnTopologyId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsEVPNTopologies.DeleteOrgEvpnTopology(ctx, orgId, evpnTopologyId)
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


# Get Org Evpn Topology

Get One EVPN Topology Detail

```go
GetOrgEvpnTopology(
    ctx context.Context,
    orgId uuid.UUID,
    evpnTopologyId uuid.UUID) (
    models.ApiResponse[models.EvpnTopology],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `evpnTopologyId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.EvpnTopology](../../doc/models/evpn-topology.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

evpnTopologyId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsEVPNTopologies.GetOrgEvpnTopology(ctx, orgId, evpnTopologyId)
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


# List Org Evpn Topologies

Get List of the existing Org EVPN topologies

```go
ListOrgEvpnTopologies(
    ctx context.Context,
    orgId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.EvpnTopologyResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br>**Constraints**: `>= 1` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.EvpnTopologyResponse](../../doc/models/evpn-topology-response.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

limit := 100

page := 1

apiResponse, err := orgsEVPNTopologies.ListOrgEvpnTopologies(ctx, orgId, &limit, &page)
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
    "created_time": 1736421230,
    "evpn_options": {
      "auto_loopback_subnet": "172.16.192.0/24",
      "auto_loopback_subnet6": "fd33:ab00:2::/64",
      "auto_router_id_subnet": "172.16.254.0/23",
      "core_as_border": true,
      "overlay": {
        "as": 65000
      },
      "per_vlan_vga_v4_mac": false,
      "routed_at": "core",
      "underlay": {
        "as_base": 65001,
        "subnet": "10.255.240.0/20",
        "use_ipv6": false
      }
    },
    "for_site": false,
    "id": "764fb173-94f9-447c-8454-def62e5a999f",
    "modified_time": 1736421230,
    "name": "tert",
    "org_id": "3a2627d7-bfbc-45af-b85d-8841581c6d63",
    "pod_names": {
      "1": "Pod 1"
    },
    "site_id": "00000000-0000-0000-0000-000000000000",
    "version": 6
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


# Update Org Evpn Topology

Update the EVPN Topology

```go
UpdateOrgEvpnTopology(
    ctx context.Context,
    orgId uuid.UUID,
    evpnTopologyId uuid.UUID,
    body *models.EvpnTopology) (
    models.ApiResponse[models.EvpnTopology],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `evpnTopologyId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.EvpnTopology`](../../doc/models/evpn-topology.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.EvpnTopology](../../doc/models/evpn-topology.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

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

apiResponse, err := orgsEVPNTopologies.UpdateOrgEvpnTopology(ctx, orgId, evpnTopologyId, &body)
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

