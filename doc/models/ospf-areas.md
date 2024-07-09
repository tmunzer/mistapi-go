
# Ospf Areas

Junos OSPF areas

## Structure

`OspfAreas`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `IncludeLoopback` | `*bool` | Optional | **Default**: `false` |
| `Networks` | [`map[string]models.OspfAreasNetwork`](../../doc/models/ospf-areas-network.md) | Optional | Networks to participate in an OSPF area.<br>Property key is the network name |
| `Type` | [`*models.OspfAreasTypeEnum`](../../doc/models/ospf-areas-type-enum.md) | Optional | OSPF type, default (default) / stub / nssa<br>**Default**: `"default"` |

## Example (as JSON)

```json
{
  "include_loopback": false,
  "networks": {
    "corp": {
      "auth_keys": {
        "1": "auth-key-1"
      },
      "auth_type": "md5",
      "bfd_minimum_interval": 500,
      "dead_interval": 40,
      "hello_interval": 10,
      "interface_type": "nbma",
      "metric": 10000,
      "auth_password": "auth_password4"
    },
    "guest": {
      "passive": true,
      "auth_keys": {
        "key0": "auth_keys4"
      },
      "auth_password": "auth_password4",
      "auth_type": "password",
      "bfd_minimum_interval": 94,
      "dead_interval": 118
    }
  },
  "type": "default"
}
```

