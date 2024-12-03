
# Ospf Areas Network

Property key is the network name. Networks to participate in an OSPF area

*This model accepts additional fields of type interface{}.*

## Structure

`OspfAreasNetwork`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AuthKeys` | `map[string]string` | Optional | Required if `auth_type`==`md5`. Property key is the key number |
| `AuthPassword` | `*string` | Optional | Required if `auth_type`==`password`, the password, max length is 8 |
| `AuthType` | [`*models.OspfAreaNetworkAuthTypeEnum`](../../doc/models/ospf-area-network-auth-type-enum.md) | Optional | auth type. enum: `md5`, `none`, `password`<br>**Default**: `"none"` |
| `BfdMinimumInterval` | `*int` | Optional | **Constraints**: `>= 1`, `<= 255000` |
| `DeadInterval` | `*int` | Optional | **Constraints**: `>= 1`, `<= 65535` |
| `ExportPolicy` | `*string` | Optional | - |
| `HelloInterval` | `*int` | Optional | **Constraints**: `>= 1`, `<= 255` |
| `ImportPolicy` | `*string` | Optional | - |
| `InterfaceType` | [`*models.OspfAreaNetworkInterfaceTypeEnum`](../../doc/models/ospf-area-network-interface-type-enum.md) | Optional | interface type (nbma = non-broadcast multi-access). enum: `broadcast`, `nbma`, `p2mp`, `p2p`<br>**Default**: `"broadcast"` |
| `Metric` | `models.Optional[int]` | Optional | **Constraints**: `>= 1`, `<= 65535` |
| `NoReadvertiseToOverlay` | `*bool` | Optional | by default, we'll re-advertise all learned OSPF routes toward overlay<br>**Default**: `false` |
| `Passive` | `*bool` | Optional | whether to send OSPF-Hello<br>**Default**: `false` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "auth_keys": {
    "1": "auth-key-1"
  },
  "auth_password": "simple",
  "auth_type": "md5",
  "bfd_minimum_interval": 500,
  "dead_interval": 40,
  "export_policy": "export_policy",
  "import_policy": "import_policy",
  "interface_type": "broadcast",
  "metric": 10000,
  "no_readvertise_to_overlay": false,
  "passive": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

