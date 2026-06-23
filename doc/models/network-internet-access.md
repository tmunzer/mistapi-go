
# Network Internet Access

Direct internet access settings for an organization network

## Structure

`NetworkInternetAccess`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreateSimpleServicePolicy` | `*bool` | Optional | Whether Mist should create simple service policies for restricted internet access<br><br>**Default**: `false` |
| `DestinationNat` | [`map[string]models.NetworkInternetAccessDestinationNatProperty`](../../doc/models/network-internet-access-destination-nat-property.md) | Optional | Property key can be an External IP (i.e. "63.16.0.3"), an External IP:Port (i.e. "63.16.0.3:443"), an External Port (i.e. ":443"), an External CIDR (i.e. "63.16.0.0/30"), an External CIDR:Port (i.e. "63.16.0.0/30:443") or a Variable (i.e. "{{myvar}}"). At least one of the `internal_ip` or `port` must be defined |
| `Enabled` | `*bool` | Optional | Whether direct internet access is enabled for this network |
| `Restricted` | `*bool` | Optional | By default, all access is allowed, to only allow certain traffic, make `restricted`=`true` and define service_policies<br><br>**Default**: `false` |
| `StaticNat` | [`map[string]models.NetworkInternetAccessStaticNatProperty`](../../doc/models/network-internet-access-static-nat-property.md) | Optional | Property key may be an External IP address (i.e. "63.16.0.3"), a CIDR (i.e. "63.16.0.12/20") or a Variable (i.e. "{{myvar}}") |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    networkInternetAccess := models.NetworkInternetAccess{
        CreateSimpleServicePolicy: models.ToPointer(false),
        DestinationNat:            map[string]models.NetworkInternetAccessDestinationNatProperty{
            "key0": models.NetworkInternetAccessDestinationNatProperty{
                InternalIp:           models.ToPointer("internal_ip0"),
                Name:                 models.ToPointer("name4"),
                Port:                 models.ToPointer("port4"),
                WanName:              models.ToPointer("wan_name0"),
            },
            "key1": models.NetworkInternetAccessDestinationNatProperty{
                InternalIp:           models.ToPointer("internal_ip0"),
                Name:                 models.ToPointer("name4"),
                Port:                 models.ToPointer("port4"),
                WanName:              models.ToPointer("wan_name0"),
            },
            "key2": models.NetworkInternetAccessDestinationNatProperty{
                InternalIp:           models.ToPointer("internal_ip0"),
                Name:                 models.ToPointer("name4"),
                Port:                 models.ToPointer("port4"),
                WanName:              models.ToPointer("wan_name0"),
            },
        },
        Enabled:                   models.ToPointer(false),
        Restricted:                models.ToPointer(false),
        StaticNat:                 map[string]models.NetworkInternetAccessStaticNatProperty{
            "key0": models.NetworkInternetAccessStaticNatProperty{
                InternalIp:           models.ToPointer("internal_ip0"),
                Name:                 models.ToPointer("name4"),
                WanName:              models.ToPointer("wan_name0"),
            },
        },
    }

}
```

