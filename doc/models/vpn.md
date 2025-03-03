
# Vpn

*This model accepts additional fields of type interface{}.*

## Structure

`Vpn`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional | When the object has been created, in epoch |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organization |
| `ModifiedTime` | `*float64` | Optional | When the object has been modified for the last time, in epoch |
| `Name` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `PathSelection` | [`*models.VpnPathSelection`](../../doc/models/vpn-path-selection.md) | Optional | Only if `type`==`hub_spoke` |
| `Paths` | [`map[string]models.VpnPath`](../../doc/models/vpn-path.md) | Required | For `type`==`hub_spoke`, Property key is the VPN name. For `type`==`mesh`, Property key is the Interface name |
| `Type` | [`*models.VpnTypeEnum`](../../doc/models/vpn-type-enum.md) | Optional | enum: `hub_spoke`, `mesh`<br>**Default**: `"hub_spoke"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "name": "name0",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "paths": {
    "key0": {
      "bfd_profile": "broadband",
      "bfd_use_tunnel_mode": false,
      "ip": "ip8",
      "peer_paths": {
        "key0": {
          "preference": 144,
          "exampleAdditionalProperty": {
            "key1": "val1",
            "key2": "val2"
          }
        }
      },
      "pod": 146,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "key1": {
      "bfd_profile": "broadband",
      "bfd_use_tunnel_mode": false,
      "ip": "ip8",
      "peer_paths": {
        "key0": {
          "preference": 144,
          "exampleAdditionalProperty": {
            "key1": "val1",
            "key2": "val2"
          }
        }
      },
      "pod": 146,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "key2": {
      "bfd_profile": "broadband",
      "bfd_use_tunnel_mode": false,
      "ip": "ip8",
      "peer_paths": {
        "key0": {
          "preference": 144,
          "exampleAdditionalProperty": {
            "key1": "val1",
            "key2": "val2"
          }
        }
      },
      "pod": 146,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  },
  "type": "hub_spoke",
  "created_time": 229.0,
  "modified_time": 105.96,
  "path_selection": {
    "strategy": "simple",
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

