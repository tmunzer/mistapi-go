# Sites Device Profiles

```go
sitesDeviceProfiles := client.SitesDeviceProfiles()
```

## Class Name

`SitesDeviceProfiles`


# List Site Device Profiles Derived

Get the list of derived Device Profiles for a Site

```go
ListSiteDeviceProfilesDerived(
    ctx context.Context,
    siteId uuid.UUID,
    resolve *bool) (
    models.ApiResponse[[]models.Deviceprofile],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `resolve` | `*bool` | Query, Optional | Whether resolve the site variables<br><br>**Default**: `false` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type []models.Deviceprofile.

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resolve := false

apiResponse, err := sitesDeviceProfiles.ListSiteDeviceProfilesDerived(ctx, siteId, &resolve)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    responseBody := apiResponse.Data
    for _, item := range responseBody {
        if i, ok := item.AsDeviceprofileAp(); ok {
            fmt.Println("Value narrowed down to models.DeviceprofileAp: ", *i)
        } else if i, ok := item.AsDeviceprofileGateway(); ok {
            fmt.Println("Value narrowed down to models.DeviceprofileGateway: ", *i)
        } else if i, ok := item.AsDeviceprofileSwitch(); ok {
            fmt.Println("Value narrowed down to models.DeviceprofileSwitch: ", *i)
        }
    }

    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response

```
[
  {
    "aeroscout": {
      "enabled": true,
      "host": "string"
    },
    "created_time": 0,
    "disable_eth1": true,
    "disable_module": true,
    "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "ip_config": {
      "dns": [
        "string"
      ],
      "dns_suffix": [
        "string"
      ],
      "gateway": "192.168.0.1",
      "gateway6": "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
      "ip": "192.168.0.1",
      "ip6": "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
      "mtu": 1500,
      "netmask": "192.168.0.1",
      "netmask6": "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
      "type": "static",
      "type6": "static",
      "vlan_id": 1
    },
    "mesh": {
      "enabled": true,
      "group": 1,
      "role": "base"
    },
    "modified_time": 0,
    "name": "string",
    "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "poe_passthrough": true,
    "switch_config": {
      "enabled": true,
      "eth0": {
        "port_vlan_id": 1,
        "vlan_ids": [
          1,
          10
        ]
      },
      "eth1": {
        "port_vlan_id": 1,
        "vlan_ids": [
          10
        ]
      },
      "eth2": {
        "port_vlan_id": 1,
        "vlan_ids": [
          10
        ]
      },
      "eth3": {
        "port_vlan_id": 1,
        "vlan_ids": [
          10
        ]
      },
      "module": {
        "port_vlan_id": 1,
        "vlan_ids": [
          10
        ]
      },
      "wds": {
        "port_vlan_id": 1,
        "vlan_ids": [
          10
        ]
      }
    },
    "type": "ap",
    "usb_config": {
      "cacert": "string",
      "enabled": true,
      "host": "string",
      "type": "imagotag",
      "verify_cert": true
    }
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

