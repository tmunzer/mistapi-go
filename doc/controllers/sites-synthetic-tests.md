# Sites Synthetic Tests

```go
sitesSyntheticTests := client.SitesSyntheticTests()
```

## Class Name

`SitesSyntheticTests`

## Methods

* [Get Site Device Synthetic Test](../../doc/controllers/sites-synthetic-tests.md#get-site-device-synthetic-test)
* [Search Site Synthetic Test](../../doc/controllers/sites-synthetic-tests.md#search-site-synthetic-test)
* [Start Site Switch Radius Synthetic Test](../../doc/controllers/sites-synthetic-tests.md#start-site-switch-radius-synthetic-test)
* [Trigger Site Device Synthetic Test](../../doc/controllers/sites-synthetic-tests.md#trigger-site-device-synthetic-test)
* [Trigger Site Synthetic Test](../../doc/controllers/sites-synthetic-tests.md#trigger-site-synthetic-test)


# Get Site Device Synthetic Test

Get Device Synthetic Test

```go
GetSiteDeviceSyntheticTest(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID) (
    models.ApiResponse[models.SynthetictestInfo],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.SynthetictestInfo](../../doc/models/synthetictest-info.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := sitesSyntheticTests.GetSiteDeviceSyntheticTest(ctx, siteId, deviceId)
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
  "device_type": "gateway",
  "mac": "5c5b35584a6f",
  "port_id": "ge-0/0/1.100",
  "start_time": 1675718807,
  "status": "inprogress",
  "type": "speedtest"
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Device not online / Device not supported / Already in progress | `ApiError` |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Search Site Synthetic Test

Search Site Synthetic Testing

```go
SearchSiteSyntheticTest(
    ctx context.Context,
    siteId uuid.UUID,
    mac *string,
    portId *string,
    vlanId *string,
    by *string,
    reason *string,
    mType *models.SynthetictestTypeEnum,
    protocol *models.SynthetictestProtocolEnum,
    tenant *string) (
    models.ApiResponse[models.ResponseSynthetictestSearch],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mac` | `*string` | Query, Optional | Device MAC Address |
| `portId` | `*string` | Query, Optional | Port_id used to run the test (for SSR only) |
| `vlanId` | `*string` | Query, Optional | VLAN ID |
| `by` | `*string` | Query, Optional | Entity who triggers the test |
| `reason` | `*string` | Query, Optional | Test failure reason |
| `mType` | [`*models.SynthetictestTypeEnum`](../../doc/models/synthetictest-type-enum.md) | Query, Optional | Synthetic test type |
| `protocol` | [`*models.SynthetictestProtocolEnum`](../../doc/models/synthetictest-protocol-enum.md) | Query, Optional | Connectivity protocol |
| `tenant` | `*string` | Query, Optional | Tenant network in which lan_connectivity test was run |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseSynthetictestSearch](../../doc/models/response-synthetictest-search.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mac := "5c5b350e0410"

portId := "ge-1/0/1"

vlanId := "100"

by := "user"

reason := "test failed"







apiResponse, err := sitesSyntheticTests.SearchSiteSyntheticTest(ctx, siteId, &mac, &portId, &vlanId, &by, &reason, nil, nil, nil)
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
  "end": 0,
  "limit": 0,
  "next": "string",
  "results": [
    {
      "by": "user",
      "device_type": "gateway",
      "failed": false,
      "latency": 40,
      "mac": "aff827549235",
      "port_id": "ge-0/0/2",
      "rx_mbps": 322,
      "timestamp": 1706824045.059036,
      "tx_mbps": 199,
      "type": "speedtest",
      "vlan_id": 20
    },
    {
      "by": "marvis",
      "device_type": "gateway",
      "failed": true,
      "latency": 0,
      "mac": "8396cd006c8c",
      "port_id": "ge-0/0/2",
      "reason": "interface not ready to perform test",
      "rx_mbps": 0,
      "timestamp": 1706824045.059036,
      "tx_mbps": 0,
      "type": "speedtest",
      "vlan_id": 100
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


# Start Site Switch Radius Synthetic Test

Ping test from the AP to confirm ‘reachability’ of the Radius server. Utilize Juniper EX switch(to which an AP is connected to) radius test capabilities to get details on the Radius Server ‘availability’ .

```go
StartSiteSwitchRadiusSyntheticTest(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.SynthetictestRadiusServer) (
    models.ApiResponse[models.WebsocketSession],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.SynthetictestRadiusServer`](../../doc/models/synthetictest-radius-server.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.WebsocketSession](../../doc/models/websocket-session.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.SynthetictestRadiusServer{
    Password:             "string",
    Profile:              models.ToPointer("dot1x"),
    User:                 "string",
}

apiResponse, err := sitesSyntheticTests.StartSiteSwitchRadiusSyntheticTest(ctx, siteId, deviceId, &body)
if err != nil {
    log.Fatalln(err)
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Trigger Site Device Synthetic Test

Trigger Device Synthetic Test

```go
TriggerSiteDeviceSyntheticTest(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.SynthetictestDevice) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.SynthetictestDevice`](../../doc/models/synthetictest-device.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.SynthetictestDevice{
    Host:                 models.ToPointer("www.example.com"),
    Hostname:             models.ToPointer("google.com\""),
    Ip:                   models.ToPointer("192.168.3.5"),
    Password:             models.ToPointer("test123"),
    PingCount:            models.ToPointer(10),
    PingDetails:          models.ToPointer(false),
    PingSize:             models.ToPointer(56),
    PortId:               models.ToPointer("wan0"),
    Protocol:             models.ToPointer(models.SynthetictestDeviceProtocolEnum_ENUMPINGTRACEROUTE),
    Tenant:               models.ToPointer("lan_network1"),
    Timeout:              models.ToPointer(60),
    TracerouteUdpPort:    models.ToPointer(33434),
    Type:                 models.SynthetictestTypeEnum_RADIUS,
    Url:                  models.ToPointer("https://www.example.com"),
    Username:             models.ToPointer("user"),
}

resp, err := sitesSyntheticTests.TriggerSiteDeviceSyntheticTest(ctx, siteId, deviceId, &body)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Device not online / Device not supported / Already in progress | `ApiError` |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Trigger Site Synthetic Test

Trigger Synthetic Testing

```go
TriggerSiteSyntheticTest(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.Synthetictest) (
    models.ApiResponse[models.ResponseSynthetictest],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Synthetictest`](../../doc/models/synthetictest.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseSynthetictest](../../doc/models/response-synthetictest.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Synthetictest{
    Email:                models.ToPointer("test@mist.com"),
}

apiResponse, err := sitesSyntheticTests.TriggerSiteSyntheticTest(ctx, siteId, &body)
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
  "id": "a42775f6-edc8-69b5-f979-542fa1b43ff9",
  "message": "Successfully queued synthetic test for the site.",
  "status": "string"
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

