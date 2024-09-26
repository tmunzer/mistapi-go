
# Ospf Area

Property key is the OSPF Area (Area should be a number (0-255) / IP address)

## Structure

`OspfArea`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `IncludeLoopback` | `*bool` | Optional | **Default**: `false` |
| `Networks` | [`map[string]models.OspfAreasNetwork`](../../doc/models/ospf-areas-network.md) | Optional | - |
| `Type` | [`*models.OspfAreaTypeEnum`](../../doc/models/ospf-area-type-enum.md) | Optional | OSPF type. enum: `default`, `nssa`, `stub`<br>**Default**: `"default"` |

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
