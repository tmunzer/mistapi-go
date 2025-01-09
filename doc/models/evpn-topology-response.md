
# Evpn Topology Response

*This model accepts additional fields of type interface{}.*

## Structure

`EvpnTopologyResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional | when the object has been created, in epoch |
| `EvpnOptions` | [`*models.EvpnOptions`](../../doc/models/evpn-options.md) | Optional | EVPN Options |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organnization |
| `ModifiedTime` | `*float64` | Optional | when the object has been modified for the last time, in epoch |
| `Name` | `*string` | Optional | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Overwrite` | `*bool` | Optional | - |
| `PodNames` | `map[string]string` | Optional | Property key is the pod number |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "name": "CC",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "created_time": 252.32,
  "evpn_options": {
    "auto_loopback_subnet": "auto_loopback_subnet4",
    "auto_loopback_subnet6": "auto_loopback_subnet60",
    "auto_router_id_subnet": "auto_router_id_subnet8",
    "auto_router_id_subnet6": "auto_router_id_subnet60",
    "core_as_border": false,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "modified_time": 82.64,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

