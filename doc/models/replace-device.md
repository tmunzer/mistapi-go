
# Replace Device

*This model accepts additional fields of type interface{}.*

## Structure

`ReplaceDevice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Discard` | `[]string` | Optional | attributes that you don’t want to copy |
| `InventoryMac` | `*string` | Optional | MAC Address of the inventory that will be replacing the old one. It has to be claimed and unassigned |
| `Mac` | `*string` | Optional | MAC Address of the device to replace |
| `SiteId` | `*string` | Optional | the site_id of the device to be replaced |
| `TuntermPortConfig` | [`*models.TuntermPortConfig`](../../doc/models/tunterm-port-config.md) | Optional | ethernet port configurations |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "inventory_mac": "5c5b35000301",
  "mac": "5c5b35000101",
  "site_id": "4ac1dcf4-9d8b-7211-65c4-057819f0862b",
  "discard": [
    "discard6"
  ],
  "tunterm_port_config": {
    "downstream_ports": [
      "downstream_ports5"
    ],
    "separate_upstream_downstream": false,
    "upstream_port_vlan_id": 16,
    "upstream_ports": [
      "upstream_ports0"
    ],
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

