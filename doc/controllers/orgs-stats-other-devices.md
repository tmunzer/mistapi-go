# Orgs Stats-Other Devices

```go
orgsStatsOtherDevices := client.OrgsStatsOtherDevices()
```

## Class Name

`OrgsStatsOtherDevices`


# Get Org Other Device Stats

Get Otherdevice Stats

```go
GetOrgOtherDeviceStats(
    ctx context.Context,
    orgId uuid.UUID,
    deviceMac string) (
    models.ApiResponse[models.StatsDeviceOther],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `deviceMac` | `string` | Template, Required | **Constraints**: *Pattern*: `^[0-9a-fA-F]{12}$` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.StatsDeviceOther](../../doc/models/stats-device-other.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceMac := "0000000000ab"

apiResponse, err := orgsStatsOtherDevices.GetOrgOtherDeviceStats(ctx, orgId, deviceMac)
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
  "cached_stats": true,
  "connected_devices": {
    "0200010edbca": {
      "mac": "020001abcdef",
      "name": "DNT-NTR-GWE",
      "port_id": "ge-0/0/1",
      "type": "gateway"
    }
  },
  "last_seen": 1740996902,
  "lldp_enabled": true,
  "mac": "00304498a1e8",
  "uptime": 622828,
  "vendor": "cradlepoint",
  "vendor_specific": {
    "interfaces": {
      "ethernet-IPPT": {
        "bytes_in": 331068567,
        "bytes_out": 6763536255,
        "display_name": "Primary LAN",
        "ip": "192.168.0.1",
        "link": true,
        "mode": "lan",
        "port_parent": "Primary LAN",
        "service_mode": "Ethernet",
        "type": "ethernet",
        "uptime": 0
      },
      "ethernet-lan": {
        "bytes_in": 13072566048,
        "bytes_out": 5617915438,
        "display_name": "Secondary LAN",
        "ip": "192.168.0.1",
        "link": false,
        "mode": "lan",
        "port_parent": "Secondary LAN",
        "service_mode": "Ethernet",
        "type": "ethernet",
        "uptime": 0
      },
      "mdm-8a1084c9": {
        "bytes_in": 0,
        "bytes_out": 0,
        "carrier": "Unknown Service",
        "imei": "866401234567894",
        "imsi": "",
        "ip": "",
        "link": false,
        "mode": "wan",
        "mtu": 1400,
        "rsrp": 0,
        "rsrq": 0,
        "rssi": 0,
        "service_mode": "Not Available",
        "sinr": 0,
        "state": "NOSIM",
        "type": "mdm",
        "uptime": 0
      },
      "mdm-8a1fc70c": {
        "bytes_in": 5623096929,
        "bytes_out": 12372750366,
        "carrier": "AT&T",
        "imei": "866401234567893",
        "imsi": "208001234567893",
        "ip": "12.68.86.17",
        "link": true,
        "mode": "wan",
        "mtu": 1400,
        "rsrp": -108,
        "rsrq": -14,
        "rssi": -74,
        "service_mode": "5G NSA",
        "sinr": -1.2,
        "state": "READY",
        "type": "mdm",
        "uptime": 2095779
      }
    }
  },
  "version": "7.24.80"
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

