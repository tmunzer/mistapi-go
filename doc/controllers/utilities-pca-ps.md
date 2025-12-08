# Utilities PCA Ps

```go
utilitiesPCAPs := client.UtilitiesPCAPs()
```

## Class Name

`UtilitiesPCAPs`

## Methods

* [Get Org Capturing Status](../../doc/controllers/utilities-pca-ps.md#get-org-capturing-status)
* [Get Site Capturing Status](../../doc/controllers/utilities-pca-ps.md#get-site-capturing-status)
* [List Org Packet Captures](../../doc/controllers/utilities-pca-ps.md#list-org-packet-captures)
* [List Site Packet Captures](../../doc/controllers/utilities-pca-ps.md#list-site-packet-captures)
* [Start Org Packet Capture](../../doc/controllers/utilities-pca-ps.md#start-org-packet-capture)
* [Start Site Packet Capture](../../doc/controllers/utilities-pca-ps.md#start-site-packet-capture)
* [Stop Org Packet Capture](../../doc/controllers/utilities-pca-ps.md#stop-org-packet-capture)
* [Stop Site Packet Capture](../../doc/controllers/utilities-pca-ps.md#stop-site-packet-capture)
* [Update Site Packet Capture](../../doc/controllers/utilities-pca-ps.md#update-site-packet-capture)


# Get Org Capturing Status

Get Org Capturing status

```go
GetOrgCapturingStatus(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[models.ResponsePcapStatus],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponsePcapStatus](../../doc/models/response-pcap-status.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := utilitiesPCAPs.GetOrgCapturingStatus(ctx, orgId)
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
  "aps": [
    "5c5b350e001c",
    "5c5b350e001b"
  ],
  "client_mac": "60a10a773412",
  "duration": 300,
  "failed": [],
  "id": "a9a84e13-a714-b1eb-152f-a434416217d5",
  "includes_mcast": false,
  "max_pkt_len": 128,
  "num_packets": 1000,
  "ok": [
    "5c5b350e001c",
    "5c5b350e001b"
  ],
  "started_time": 1435080709,
  "type": "client"
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


# Get Site Capturing Status

Get Capturing status

```go
GetSiteCapturingStatus(
    ctx context.Context,
    siteId uuid.UUID) (
    models.ApiResponse[models.ResponsePcapStatus],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponsePcapStatus](../../doc/models/response-pcap-status.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := utilitiesPCAPs.GetSiteCapturingStatus(ctx, siteId)
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
  "aps": [
    "5c5b350e001c",
    "5c5b350e001b"
  ],
  "client_mac": "60a10a773412",
  "duration": 300,
  "failed": [],
  "id": "a9a84e13-a714-b1eb-152f-a434416217d5",
  "includes_mcast": false,
  "max_pkt_len": 128,
  "num_packets": 1000,
  "ok": [
    "5c5b350e001c",
    "5c5b350e001b"
  ],
  "started_time": 1435080709,
  "type": "client"
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


# List Org Packet Captures

Get List of Org  Packet Captures

```go
ListOrgPacketCaptures(
    ctx context.Context,
    orgId uuid.UUID,
    start *string,
    end *string,
    duration *string,
    limit *int,
    page *int) (
    models.ApiResponse[models.ResponsePcapSearch],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `start` | `*string` | Query, Optional | Start time (epoch timestamp in seconds, or relative string like "-1d", "-1w") |
| `end` | `*string` | Query, Optional | End time (epoch timestamp in seconds, or relative string like "-1d", "-2h", "now") |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br><br>**Constraints**: `>= 1` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponsePcapSearch](../../doc/models/response-pcap-search.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

duration := "10m"

limit := 100

page := 1

apiResponse, err := utilitiesPCAPs.ListOrgPacketCaptures(ctx, orgId, nil, nil, &duration, &limit, &page)
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
  "end": 1461089816,
  "limit": 100,
  "next": "/api/v1/sites/67970e46-4e12-11e6-9188-0242ac110007/pcaps?start=1461099816&search_after=%5B1694537121217%5D&limit=100&end=1461089816",
  "results": [
    {
      "ap_macs": [
        "5c5b35000010"
      ],
      "timestamp": 1461869041,
      "type": "new_assoc",
      "url": "https://..."
    }
  ],
  "start": 1461099816
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


# List Site Packet Captures

Get List of Site Packet Captures

```go
ListSitePacketCaptures(
    ctx context.Context,
    siteId uuid.UUID,
    clientMac *string,
    start *string,
    end *string,
    duration *string,
    limit *int,
    page *int) (
    models.ApiResponse[models.ResponsePcapSearch],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `clientMac` | `*string` | Query, Optional | Optional client mac filter |
| `start` | `*string` | Query, Optional | Start time (epoch timestamp in seconds, or relative string like "-1d", "-1w") |
| `end` | `*string` | Query, Optional | End time (epoch timestamp in seconds, or relative string like "-1d", "-2h", "now") |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br><br>**Constraints**: `>= 1` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponsePcapSearch](../../doc/models/response-pcap-search.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

duration := "10m"

limit := 100

page := 1

apiResponse, err := utilitiesPCAPs.ListSitePacketCaptures(ctx, siteId, nil, nil, nil, &duration, &limit, &page)
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
  "end": 1461089816,
  "limit": 100,
  "next": "/api/v1/sites/67970e46-4e12-11e6-9188-0242ac110007/pcaps?start=1461099816&search_after=%5B1694537121217%5D&limit=100&end=1461089816",
  "results": [
    {
      "ap_macs": [
        "5c5b35000010"
      ],
      "timestamp": 1461869041,
      "type": "new_assoc",
      "url": "https://..."
    }
  ],
  "start": 1461099816
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


# Start Org Packet Capture

Initiate a Packet Capture

**NOTE**: For packet captures of org level Mist Edges only. Use [Start Site Packet Capture](../../doc/controllers/utilities-pca-ps.md#start-site-packet-capture) for site level Mist Edges.

The output will be available through websocket. As there can be multiple command issued against the same AP at the same time and the output all goes through the same websocket stream, session is introduced for demux.

#### Subscribe to Device Command outputs

`WS /api-ws/v1/stream`

```json
{
    "subscribe": "/sites/{site_id}/pcaps"
}
```

#### Response (Wireless/RadioTap)

```json
{
  "event": "data"
  "channel": "/orgs/67970e46-4e12-11e6-9188-0242ac110007/pcaps"
  "data": {
      "capture_id": "f039b1b4-a23e-48b2-906a-0da40524de73", 
      "pcap_dict": {
          "dst_mac": "68:ec:c5:09:2e:87",
          "src_mac": "8c:3b:ad:e0:47:40", 
          "vlan": 1, 
          "src_ip": "34.224.147.117", 
          "dst_ip": "192.168.1.55",
          "dst_port": 51635, 
          "src_port": 443,
          "protocol": "TCP", 
          "mxedge_id": "00000000-0000-0000-1000-001122334455",
          "direction": "tx", 
          "timestamp": 1652247615, 
          "length": 159.0, 
          "interface": "port0",
          "info": "1652247616.007409 IP ec2-34-224-147-117.compute-1.amazonaws.com.https > ip-192-168-1-55.ec2.internal.51635: Flags [P.], seq 
                    2192123968:2192124057, ack 4035166782, win 12, options [nop,nop,TS val 597467050 ecr 740580660], length 89\\n",
          }, 
      "pcap_raw": "1MOyoQIABAAAAAAAAAAAAP//AAABAAAAQEx7YhMzAACfAAAAnwAAAGjsxQkuh4w7reBHQIEAAAEIAEUAAI1bLEAAKAZ/CiLgk3XAqAE3AbvJs4KpKEDwg8I+gBgADFf9AAABAQgKI5yfqiwkXTQXAwMAVKY5JopoKQrVEn0/3ld4YntctGEH/rTZuwtCvzSncFw71QJveJi9uxHs57KC8w9Apph3YvXJrmWg7M37+o+YV0KH/xmr626s5Bkhb3QhKOu+NoNEmA==\"
    }
}
```

```go
StartOrgPacketCapture(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.CaptureOrg) (
    models.ApiResponse[models.ResponsePcapStart],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.CaptureOrg`](../../doc/models/containers/capture-org.md) | Body, Optional | Request Body |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponsePcapStart](../../doc/models/response-pcap-start.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.CaptureOrgContainer.FromCaptureMxedge(models.CaptureMxedge{
    Duration:             models.ToPointer(600),
    Format:               models.ToPointer(models.CaptureMxedgeFormatEnum_STREAM),
    MaxPktLen:            models.ToPointer(1500),
    Mxedges:              map[string]models.CaptureMxedgeMxedges{
        "00000000-0000-0000-1000-001122334455": models.CaptureMxedgeMxedges{
            Interfaces:           map[string]models.CaptureMxedgeMxedgesInterfaces{
                "port1": models.CaptureMxedgeMxedgesInterfaces{
                    TcpdumpExpression:    models.ToPointer("udp port 67 or udp port 68"),
                },
            },
        },
    },
    NumPackets:           models.ToPointer(100),
    Type:                 "mxedge",
})

apiResponse, err := utilitiesPCAPs.StartOrgPacketCapture(ctx, orgId, &body)
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
  "ap_count": 3,
  "aps": [],
  "duration": 600,
  "enabled": true,
  "expiry": 1614886726.5411825,
  "format": "stream",
  "id": "a9a84e13-a714-b1eb-152f-a434416217d5",
  "include_mcast": false,
  "max_pkt_len": 68,
  "num_packets": 100,
  "org_id": "a9346fba-f920-e99a-cc51-2e8dcc57fa3c",
  "raw": true,
  "site_id": "67970e46-4e12-11e6-9188-0242ac110007",
  "ssid": "",
  "timestamp": 1614886126.5411825,
  "type": "radiotap"
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


# Start Site Packet Capture

Initiate a Site Packet Capture

The output will be available through websocket. As there can be multiple command issued against the same AP at the same time and the output all goes through the same websocket stream, session is introduced for demux.

#### Subscribe to Device Command outputs

`WS /api-ws/v1/stream`

```json
{
    "subscribe": "/sites/{site_id}/pcaps"
}
```

#### Response (MxEdge)

```json
{
    "event": "data"
    "channel": "/sites/{site_id}/pcaps"
    "data": {
         "capture_id": "6b1be4fb-b239-44d9-9d3b-cb1ff3af1721",
     "lost_messages": 0
         "pcap_dict": {
             "channel_frequency": 2412,
             "channel": "1",
             "datarate": "1.0 Mbps",
             "rssi": -75, 
             "dst": "78:bd:bc:ca:0b:0a",
             "src": "18:b8:1f:4c:91:c0",
             "bssid": "18:b8:1f:4c:91:c0",
             "frame_type": "Management", 
             "frame_subtype": "Probe Response", 
         "proto": "802.11", 
             "ap_mac": "d4:20:b0:81:99:2e", 
             "direction": "tx", 
             "timestamp": 1652246543, 
             "length": 416.0,
             "interface": "radiotap",
             "info": "1652246544.467733 1683216786us tsft 1.0 Mb/s 2412 MHz 11g -75dBm signal -82dBm noise antenna 0 Probe Response (ATTKmsWiVS) [1.0* 2.0* 5.5* 11.0* 18.0 24.0 36.0 54.0 Mbit] CH: 2, PRIVACY\\n",
         }, 
        "pcap_raw": "1MOyoQIABAAAAAAAAAAAAP//AAABAAAAEEh7Yh5VBwCgAQAAoAEAAAAAKwBvCADAAQAAAIw7reCS2VNkAAAAABACbAmABLWuAAEAEBgAAwACAABQADoBeL28ygsKGLgfTJHAGLgfTJHAcIZ2WDlBJQAAAGQAERUACkFUVEttc1dpVlMBCIKEi5YkMEhsAwECBwZVUyABCx4gAQAjAhkAKgEEMgQMEhhgMBQBAAAPrAQBAAAPrAQBAAAPrAIMAAsFAQAbAABGBTIIAQAALRqtCR////8AAAAAAAAAAAAAAAAAAAAAAAAAAD0WAggVAAAAAAAAAAAAAAAAAAAAAAAAAH8IBAAIAAAAAEDdkwBQ8gQQSgABEBBEAAECEDsAAQMQRwAQn2481frn3KT+uGod2ERx+RAhAAtBcnJpcywgSW5jLhAjAApCR1cyMTAtNzAwECQACkJHVzIxMC03MDAQQgAKQkdXMjEwLTcwMBBUAAgABgBQ8gQAARARAA5BcnJpcyBXaXJlbGVzcxAIAAIgCBA8AAEBEEkABgA3KgABIN0JABAYAgEQHAAA3RgAUPICAQGEAAOkAAAnpAAAQkNeAGIyLwAzjakr"
}
```

#### Response (Wired)

```json
{
    "event": "data"
    "channel": "/sites/67970e46-4e12-11e6-9188-0242ac110007/pcaps"
    "data": {
        "capture_id": "f039b1b4-a23e-48b2-906a-0da40524de73", 
        "pcap_dict": {
             "dst_mac": "68:ec:c5:09:2e:87",
             "src_mac": "8c:3b:ad:e0:47:40", 
             "vlan": 1, 
             "src_ip": "34.224.147.117", 
             "dst_ip": "192.168.1.55",
             "dst_port": 51635, 
             "src_port": 443,
             "proto": "TCP", 
             "ap_mac": "d4:20:b0:81:99:2e",
             "direction": "tx", 
             "timestamp": 1652247615, 
             "length": 159.0, 
             "interface": "wired",
             "info": "1652247616.007409 IP ec2-34-224-147-117.compute-1.amazonaws.com.https > ip-192-168-1-55.ec2.internal.51635: Flags [P.], seq 2192123968:2192124057, ack 4035166782, win 12, options [nop,nop,TS val 597467050 ecr 740580660], length 89\\n",
             }, 
        "pcap_raw": "1MOyoQIABAAAAAAAAAAAAP//AAABAAAAQEx7YhMzAACfAAAAnwAAAGjsxQkuh4w7reBHQIEAAAEIAEUAAI1bLEAAKAZ/CiLgk3XAqAE3AbvJs4KpKEDwg8I+gBgADFf9AAABAQgKI5yfqiwkXTQXAwMAVKY5JopoKQrVEn0/3ld4YntctGEH/rTZuwtCvzSncFw71QJveJi9uxHs57KC8w9Apph3YvXJrmWg7M37+o+YV0KH/xmr626s5Bkhb3QhKOu+NoNEmA=="

    }
}
```

#### Stop Response (Wired/Wireless)

```json
{
    "event": "data"
    "channel": "/sites/67970e46-4e12-11e6-9188-0242ac110007/pcaps"
    "data": {
      "capture_id": "a2f7374d-6a70-41fd-8a3f-71e42573baaf", 
      "lost_messages": 0,
      "pcap_dict": null
    }
}
```

```go
StartSitePacketCapture(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.CaptureSite) (
    models.ApiResponse[models.ResponsePcapStart],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.CaptureSite`](../../doc/models/containers/capture-site.md) | Body, Optional | Request Body |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponsePcapStart](../../doc/models/response-pcap-start.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.CaptureSiteContainer.FromCaptureNewAssoc(models.CaptureNewAssoc{
    ClientMac:            models.ToPointer("60a10a773412"),
    Duration:             models.NewOptional(models.ToPointer(600)),
    IncludesMcast:        models.ToPointer(false),
    MaxPktLen:            models.NewOptional(models.ToPointer(128)),
    NumPackets:           models.NewOptional(models.ToPointer(100)),
    Type:                 "new_assoc",
})

apiResponse, err := utilitiesPCAPs.StartSitePacketCapture(ctx, siteId, &body)
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
  "ap_count": 3,
  "aps": [],
  "duration": 600,
  "enabled": true,
  "expiry": 1614886726.5411825,
  "format": "stream",
  "id": "a9a84e13-a714-b1eb-152f-a434416217d5",
  "include_mcast": false,
  "max_pkt_len": 68,
  "num_packets": 100,
  "org_id": "a9346fba-f920-e99a-cc51-2e8dcc57fa3c",
  "raw": true,
  "site_id": "67970e46-4e12-11e6-9188-0242ac110007",
  "ssid": "",
  "timestamp": 1614886126.5411825,
  "type": "radiotap"
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


# Stop Org Packet Capture

Stop current Org capture

```go
StopOrgPacketCapture(
    ctx context.Context,
    orgId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := utilitiesPCAPs.StopOrgPacketCapture(ctx, orgId)
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


# Stop Site Packet Capture

Stop current capture

```go
StopSitePacketCapture(
    ctx context.Context,
    siteId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := utilitiesPCAPs.StopSitePacketCapture(ctx, siteId)
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


# Update Site Packet Capture

Update or add notes to a completed packet capture

```go
UpdateSitePacketCapture(
    ctx context.Context,
    siteId uuid.UUID,
    pcapId uuid.UUID,
    body *models.NotesString) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `pcapId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.NotesString`](../../doc/models/notes-string.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

pcapId := uuid.MustParse("00002548-0000-0000-0000-000000000000")

body := models.NotesString{
    Notes:                models.ToPointer("wired pcap test"),
}

resp, err := utilitiesPCAPs.UpdateSitePacketCapture(ctx, siteId, pcapId, &body)
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

