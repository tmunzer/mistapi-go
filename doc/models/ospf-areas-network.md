
# Ospf Areas Network

Property key is the network name. Networks to participate in an OSPF area

## Structure

`OspfAreasNetwork`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AuthKeys` | `map[string]string` | Optional | Required if `auth_type`==`md5`. Property key is the key number |
| `AuthPassword` | `*string` | Optional | Required if `auth_type`==`password`, the password, max length is 8 |
| `AuthType` | [`*models.OspfAreaNetworkAuthTypeEnum`](../../doc/models/ospf-area-network-auth-type-enum.md) | Optional | auth type. enum: `md5`, `none`, `password`<br><br>**Default**: `"none"` |
| `BfdMinimumInterval` | `*int` | Optional | Minimum BFD interval for this OSPF network, in milliseconds<br><br>**Constraints**: `>= 1`, `<= 255000` |
| `DeadInterval` | `*int` | Optional | OSPF dead interval for this network, in seconds<br><br>**Constraints**: `>= 1`, `<= 65535` |
| `ExportPolicy` | `*string` | Optional | Routing policy used to export routes from this OSPF network |
| `HelloInterval` | `*int` | Optional | OSPF hello interval for this network, in seconds<br><br>**Constraints**: `>= 1`, `<= 255` |
| `ImportPolicy` | `*string` | Optional | Routing policy used to import routes for this OSPF network |
| `InterfaceType` | [`*models.OspfAreaNetworkInterfaceTypeEnum`](../../doc/models/ospf-area-network-interface-type-enum.md) | Optional | interface type (nbma = non-broadcast multi-access). enum: `broadcast`, `nbma`, `p2mp`, `p2p`<br><br>**Default**: `"broadcast"` |
| `Metric` | `models.Optional[int]` | Optional | OSPF metric assigned to this network<br><br>**Constraints**: `>= 1`, `<= 65535` |
| `NoReadvertiseToOverlay` | `*bool` | Optional | By default, we'll re-advertise all learned OSPF routes toward overlay<br><br>**Default**: `false` |
| `Passive` | `*bool` | Optional | Whether to send OSPF-Hello<br><br>**Default**: `false` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    ospfAreasNetwork := models.OspfAreasNetwork{
        AuthKeys:               map[string]string{
            "1": "auth-key-1",
        },
        AuthPassword:           models.ToPointer("simple"),
        AuthType:               models.ToPointer(models.OspfAreaNetworkAuthTypeEnum_MD5),
        BfdMinimumInterval:     models.ToPointer(500),
        DeadInterval:           models.ToPointer(40),
        ExportPolicy:           models.ToPointer("export_policy"),
        ImportPolicy:           models.ToPointer("import_policy"),
        InterfaceType:          models.ToPointer(models.OspfAreaNetworkInterfaceTypeEnum_BROADCAST),
        Metric:                 models.NewOptional(models.ToPointer(10000)),
        NoReadvertiseToOverlay: models.ToPointer(false),
        Passive:                models.ToPointer(false),
    }

}
```

