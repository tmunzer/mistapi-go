
# Vpn

## Structure

`Vpn`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional | when the object has been created, in epoch |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organnization |
| `Links` | [`map[string]models.VpnLinksSteering`](../../doc/models/vpn-links-steering.md) | Optional | Gateways participating in mesh overlay by, on its vpn_path under wan port.<br>The steering behavior is defined at the overlay level.<br><br>Property Key is the Gateway Port name (e.g. `wan0`, `wan1`, `mpls`, ..) |
| `ModifiedTime` | `*float64` | Optional | when the object has been modified for the last time, in epoch |
| `Name` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `PathSelection` | [`*models.VpnPathSelection`](../../doc/models/vpn-path-selection.md) | Optional | Only if `type`==`hub_spoke` |
| `Paths` | [`map[string]models.VpnPath`](../../doc/models/vpn-path.md) | Required | Only if `type`==`hub_spoke`. Property key is the VPN name |
| `Type` | [`*models.VpnTypeEnum`](../../doc/models/vpn-type-enum.md) | Optional | enum: `hub_spoke`, `mesh`<br>**Default**: `"hub_spoke"` |

## Example (as JSON)

```json
{
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "name": "name0",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "paths": {
    "key0": {
      "bfd_profile": "broadband",
      "ip": "ip8",
      "pod": 146
    },
    "key1": {
      "bfd_profile": "broadband",
      "ip": "ip8",
      "pod": 146
    },
    "key2": {
      "bfd_profile": "broadband",
      "ip": "ip8",
      "pod": 146
    }
  },
  "type": "hub_spoke",
  "created_time": 229.0,
  "links": {
    "key0": {
      "preference": 8
    }
  },
  "modified_time": 105.96
}
```

