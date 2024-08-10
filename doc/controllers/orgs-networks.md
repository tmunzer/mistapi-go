# Orgs Networks

```go
orgsNetworks := client.OrgsNetworks()
```

## Class Name

`OrgsNetworks`

## Methods

* [Create Org Network](../../doc/controllers/orgs-networks.md#create-org-network)
* [Delete Org Network](../../doc/controllers/orgs-networks.md#delete-org-network)
* [Get Org Network](../../doc/controllers/orgs-networks.md#get-org-network)
* [List Org Networks](../../doc/controllers/orgs-networks.md#list-org-networks)
* [Update Org Network](../../doc/controllers/orgs-networks.md#update-org-network)


# Create Org Network

Create Organization Network

```go
CreateOrgNetwork(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.Network) (
    models.ApiResponse[models.Network],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Network`](../../doc/models/network.md) | Body, Optional | - |

## Response Type

[`models.Network`](../../doc/models/network.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Network{
    DisallowMistServices: models.ToPointer(false),
    Gateway:              models.ToPointer("192.168.70.1"),
    InternalAccess:       models.ToPointer(models.NetworkInternalAccess{
        Enabled: models.ToPointer(true),
    }),
    InternetAccess:       models.ToPointer(models.NetworkInternetAccess{
        CreateSimpleServicePolicy: models.ToPointer(false),
        DestinationNat:            map[string]models.NetworkDestinationNatProperty{
            "property1": models.NetworkDestinationNatProperty{
                InternalIp: models.ToPointer("192.168.70.30"),
                Name:       models.ToPointer("web server"),
                Port:       models.ToPointer(443),
            },
            "property2": models.NetworkDestinationNatProperty{
                InternalIp: models.ToPointer("192.168.70.30"),
                Name:       models.ToPointer("web server"),
                Port:       models.ToPointer(443),
            },
        },
        Enabled:                   models.ToPointer(true),
        Restricted:                models.ToPointer(false),
        StaticNat:                 map[string]models.NetworkStaticNatProperty{
            "property1": models.NetworkStaticNatProperty{
                InternalIp: models.ToPointer("192.168.70.3"),
                Name:       models.ToPointer("printer-1"),
            },
            "property2": models.NetworkStaticNatProperty{
                InternalIp: models.ToPointer("192.168.70.3"),
                Name:       models.ToPointer("printer-1"),
            },
        },
    }),
    Isolation:            models.ToPointer(true),
    Name:                 "string",
    Subnet:               models.ToPointer("192.168.70.0/24"),
    Tenants:              map[string]models.NetworkTenant{
        "property1": models.NetworkTenant{
            Addresses: []string{
                "10.10.10.10.",
            },
        },
        "property2": models.NetworkTenant{
            Addresses: []string{
                "10.10.10.52",
            },
        },
    },
    VlanId:               models.ToPointer(models.VlanIdWithVariableContainer.FromNumber(10)),
    VpnAccess:            map[string]models.NetworkVpnAccessConfig{
        "property1": models.NetworkVpnAccessConfig{
            AllowPing:                 models.ToPointer(true),
            DestinationNat:            map[string]models.NetworkDestinationNatProperty{
                "property1": models.NetworkDestinationNatProperty{
                    Name:       models.ToPointer("web server"),
                    Port:       models.ToPointer(443),
                },
                "property2": models.NetworkDestinationNatProperty{
                    Name:       models.ToPointer("web server"),
                    Port:       models.ToPointer(443),
                },
            },
            NatPool:                   models.ToPointer("172.16.0.0/26"),
            Routed:                    models.ToPointer(true),
            SourceNat:                 models.ToPointer(models.NetworkSourceNat{
                ExternalIp: models.ToPointer("172.16.0.8/30"),
            }),
            StaticNat:                 map[string]models.NetworkStaticNatProperty{
                "property1": models.NetworkStaticNatProperty{
                    InternalIp: models.ToPointer("192.168.70.3"),
                    Name:       models.ToPointer("pos_station-1"),
                },
                "property2": models.NetworkStaticNatProperty{
                    InternalIp: models.ToPointer("192.168.70.3"),
                    Name:       models.ToPointer("pos_station-1"),
                },
            },
            SummarizedSubnet:          models.ToPointer("172.16.0.0/16"),
        },
        "property2": models.NetworkVpnAccessConfig{
            AllowPing:                 models.ToPointer(true),
            DestinationNat:            map[string]models.NetworkDestinationNatProperty{
                "property1": models.NetworkDestinationNatProperty{
                    Name:       models.ToPointer("web server"),
                    Port:       models.ToPointer(443),
                },
                "property2": models.NetworkDestinationNatProperty{
                    Name:       models.ToPointer("web server"),
                    Port:       models.ToPointer(443),
                },
            },
            NatPool:                   models.ToPointer("172.16.0.0/26"),
            Routed:                    models.ToPointer(true),
            SourceNat:                 models.ToPointer(models.NetworkSourceNat{
                ExternalIp: models.ToPointer("172.16.0.8/30"),
            }),
            StaticNat:                 map[string]models.NetworkStaticNatProperty{
                "property1": models.NetworkStaticNatProperty{
                    InternalIp: models.ToPointer("192.168.70.3"),
                    Name:       models.ToPointer("pos_station-1"),
                },
                "property2": models.NetworkStaticNatProperty{
                    InternalIp: models.ToPointer("192.168.70.3"),
                    Name:       models.ToPointer("pos_station-1"),
                },
            },
            SummarizedSubnet:          models.ToPointer("172.16.0.0/16"),
        },
    },
}

apiResponse, err := orgsNetworks.CreateOrgNetwork(ctx, orgId, &body)
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
  "id": "497f6eca-6276-4993-bfeb-53cbbbba6f12",
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
        "10.10.10.52"
      ]
    }
  },
  "vlan_id": 10,
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
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Delete Org Network

Delete Organization Network

```go
DeleteOrgNetwork(
    ctx context.Context,
    orgId uuid.UUID,
    networkId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `networkId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

networkId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsNetworks.DeleteOrgNetwork(ctx, orgId, networkId)
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


# Get Org Network

Get Organization Network Details

```go
GetOrgNetwork(
    ctx context.Context,
    orgId uuid.UUID,
    networkId uuid.UUID) (
    models.ApiResponse[models.Network],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `networkId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.Network`](../../doc/models/network.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

networkId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsNetworks.GetOrgNetwork(ctx, orgId, networkId)
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
  "id": "497f6eca-6276-4993-bfeb-53cbbbba6f12",
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
        "10.10.10.52"
      ]
    }
  },
  "vlan_id": 10,
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
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# List Org Networks

Get List of Org Networks

```go
ListOrgNetworks(
    ctx context.Context,
    orgId uuid.UUID,
    page *int,
    limit *int) (
    models.ApiResponse[[]models.Network],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `page` | `*int` | Query, Optional | - |
| `limit` | `*int` | Query, Optional | - |

## Response Type

[`[]models.Network`](../../doc/models/network.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

page := 1

limit := 100

apiResponse, err := orgsNetworks.ListOrgNetworks(ctx, orgId, &page, &limit)
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
    "vlan_id": 10,
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Update Org Network

Update Organization Network

```go
UpdateOrgNetwork(
    ctx context.Context,
    orgId uuid.UUID,
    networkId uuid.UUID,
    body *models.Network) (
    models.ApiResponse[models.Network],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `networkId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Network`](../../doc/models/network.md) | Body, Optional | - |

## Response Type

[`models.Network`](../../doc/models/network.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

networkId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Network{
    DisallowMistServices: models.ToPointer(false),
    Gateway:              models.ToPointer("192.168.70.1"),
    InternalAccess:       models.ToPointer(models.NetworkInternalAccess{
        Enabled: models.ToPointer(true),
    }),
    InternetAccess:       models.ToPointer(models.NetworkInternetAccess{
        CreateSimpleServicePolicy: models.ToPointer(false),
        DestinationNat:            map[string]models.NetworkDestinationNatProperty{
            "property1": models.NetworkDestinationNatProperty{
                InternalIp: models.ToPointer("192.168.70.30"),
                Name:       models.ToPointer("web server"),
                Port:       models.ToPointer(443),
            },
            "property2": models.NetworkDestinationNatProperty{
                InternalIp: models.ToPointer("192.168.70.30"),
                Name:       models.ToPointer("web server"),
                Port:       models.ToPointer(443),
            },
        },
        Enabled:                   models.ToPointer(true),
        Restricted:                models.ToPointer(false),
        StaticNat:                 map[string]models.NetworkStaticNatProperty{
            "property1": models.NetworkStaticNatProperty{
                InternalIp: models.ToPointer("192.168.70.3"),
                Name:       models.ToPointer("printer-1"),
            },
            "property2": models.NetworkStaticNatProperty{
                InternalIp: models.ToPointer("192.168.70.3"),
                Name:       models.ToPointer("printer-1"),
            },
        },
    }),
    Isolation:            models.ToPointer(true),
    Name:                 "string",
    Subnet:               models.ToPointer("192.168.70.0/24"),
    Tenants:              map[string]models.NetworkTenant{
        "property1": models.NetworkTenant{
            Addresses: []string{
                "10.10.10.10",
            },
        },
        "property2": models.NetworkTenant{
            Addresses: []string{
                "10.10.10.52",
            },
        },
    },
    VlanId:               models.ToPointer(models.VlanIdWithVariableContainer.FromNumber(10)),
    VpnAccess:            map[string]models.NetworkVpnAccessConfig{
        "property1": models.NetworkVpnAccessConfig{
            AllowPing:                 models.ToPointer(true),
            DestinationNat:            map[string]models.NetworkDestinationNatProperty{
                "property1": models.NetworkDestinationNatProperty{
                    Name:       models.ToPointer("web server"),
                    Port:       models.ToPointer(443),
                },
                "property2": models.NetworkDestinationNatProperty{
                    Name:       models.ToPointer("web server"),
                    Port:       models.ToPointer(443),
                },
            },
            NatPool:                   models.ToPointer("172.16.0.0/26"),
            Routed:                    models.ToPointer(true),
            SourceNat:                 models.ToPointer(models.NetworkSourceNat{
                ExternalIp: models.ToPointer("172.16.0.8/30"),
            }),
            StaticNat:                 map[string]models.NetworkStaticNatProperty{
                "property1": models.NetworkStaticNatProperty{
                    InternalIp: models.ToPointer("192.168.70.3"),
                    Name:       models.ToPointer("pos_station-1"),
                },
                "property2": models.NetworkStaticNatProperty{
                    InternalIp: models.ToPointer("192.168.70.3"),
                    Name:       models.ToPointer("pos_station-1"),
                },
            },
            SummarizedSubnet:          models.ToPointer("172.16.0.0/16"),
        },
        "property2": models.NetworkVpnAccessConfig{
            AllowPing:                 models.ToPointer(true),
            DestinationNat:            map[string]models.NetworkDestinationNatProperty{
                "property1": models.NetworkDestinationNatProperty{
                    Name:       models.ToPointer("web server"),
                    Port:       models.ToPointer(443),
                },
                "property2": models.NetworkDestinationNatProperty{
                    Name:       models.ToPointer("web server"),
                    Port:       models.ToPointer(443),
                },
            },
            NatPool:                   models.ToPointer("172.16.0.0/26"),
            Routed:                    models.ToPointer(true),
            SourceNat:                 models.ToPointer(models.NetworkSourceNat{
                ExternalIp: models.ToPointer("172.16.0.8/30"),
            }),
            StaticNat:                 map[string]models.NetworkStaticNatProperty{
                "property1": models.NetworkStaticNatProperty{
                    InternalIp: models.ToPointer("192.168.70.3"),
                    Name:       models.ToPointer("pos_station-1"),
                },
                "property2": models.NetworkStaticNatProperty{
                    InternalIp: models.ToPointer("192.168.70.3"),
                    Name:       models.ToPointer("pos_station-1"),
                },
            },
            SummarizedSubnet:          models.ToPointer("172.16.0.0/16"),
        },
    },
}

apiResponse, err := orgsNetworks.UpdateOrgNetwork(ctx, orgId, networkId, &body)
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
  "id": "497f6eca-6276-4993-bfeb-53cbbbba6f12",
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
        "10.10.10.52"
      ]
    }
  },
  "vlan_id": 10,
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
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |

