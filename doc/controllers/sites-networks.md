# Sites Networks

```go
sitesNetworks := client.SitesNetworks()
```

## Class Name

`SitesNetworks`


# List Site Networks Derived

Retrieves the list of Networks available for the Site

```go
ListSiteNetworksDerived(
    ctx context.Context,
    siteId uuid.UUID,
    resolve *bool) (
    models.ApiResponse[[]models.Network],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `resolve` | `*bool` | Query, Optional | whether resolve the site variables |

## Response Type

[`[]models.Network`](../../doc/models/network.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resolve := false

apiResponse, err := sitesNetworks.ListSiteNetworksDerived(ctx, siteId, &resolve)
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
    "created_time": 0,
    "disallow_mist_services": false,
    "gateway": "192.168.70.1",
    "hosts": {
      "property1": {
        "external_ips": "172.16.10.32-172.16.10.35",
        "ips": "192.168.70.32-192.168.70.35"
      },
      "property2": {
        "external_ips": "172.16.10.32-172.16.10.35",
        "ips": "192.168.70.32-192.168.70.35"
      }
    },
    "id": "497f6eca-6276-4993-bfeb-53cbbbba6f13",
    "internal_access": {
      "enabled": true
    },
    "internet_access": {
      "create_simple_service_policy": false,
      "destination_nat": {
        "property1": {
          "internal_ip": "192.168.70.30",
          "name": "web server",
          "port": 443
        },
        "property2": {
          "internal_ip": "192.168.70.30",
          "name": "web server",
          "port": 443
        }
      },
      "enabled": true,
      "restricted": false,
      "static_nat": {
        "property1": {
          "internal_ip": "192.168.70.3",
          "name": "printer-1"
        },
        "property2": {
          "internal_ip": "192.168.70.3",
          "name": "printer-1"
        }
      }
    },
    "isolation": true,
    "modified_time": 0,
    "name": "string",
    "org_id": "a40f5d1f-d889-42e9-94ea-b9b33585fc6b",
    "subnet": "192.168.70.0/24",
    "tenants": {
      "property1": {
        "addresses": [
          "10.10.10.10"
        ]
      },
      "property2": {
        "addresses": [
          "10.10.10.45"
        ]
      }
    },
    "vlan_id": 0,
    "vpn_access": {
      "property1": {
        "allow_ping": true,
        "destination_nat": {
          "property1": {
            "name": "web server",
            "port": 443,
            "to": "192.168.70.5/30"
          },
          "property2": {
            "name": "web server",
            "port": 443,
            "to": "192.168.70.5/30"
          }
        },
        "nat_pool": "172.16.0.0/26",
        "routed": true,
        "source_nat": {
          "external_ip": "172.16.0.8/30"
        },
        "static_nat": {
          "property1": {
            "internal_ip": "192.168.70.3",
            "name": "pos_station-1"
          },
          "property2": {
            "internal_ip": "192.168.70.3",
            "name": "pos_station-1"
          }
        },
        "summarized_subnet": "172.16.0.0/16"
      },
      "property2": {
        "allow_ping": true,
        "destination_nat": {
          "property1": {
            "name": "web server",
            "port": 443,
            "to": "192.168.70.5/30"
          },
          "property2": {
            "name": "web server",
            "port": 443,
            "to": "192.168.70.5/30"
          }
        },
        "nat_pool": "172.16.0.0/26",
        "routed": true,
        "source_nat": {
          "external_ip": "172.16.0.8/30"
        },
        "static_nat": {
          "property1": {
            "internal_ip": "192.168.70.3",
            "name": "pos_station-1"
          },
          "property2": {
            "internal_ip": "192.168.70.3",
            "name": "pos_station-1"
          }
        },
        "summarized_subnet": "172.16.0.0/16"
      }
    }
  }
]
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |

