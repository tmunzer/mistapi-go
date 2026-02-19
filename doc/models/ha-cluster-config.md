
# Ha Cluster Config

*This model accepts additional fields of type interface{}.*

## Structure

`HaClusterConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DisableAutoConfig` | `*bool` | Optional | This disables the default behavior of a cloud-ready switch/gateway being managed/configured by Mist. Setting this to `true` means you want to disable the default behavior and do not want the device to be Mist-managed. |
| `Managed` | `*bool` | Optional | An adopted switch/gateway will not be managed/configured by Mist by default. Setting this parameter to `true` enables the adopted switch/gateway to be managed/configured by Mist. |
| `MistConfigured` | `*bool` | Optional | whether the device can be configured by Mist or not. This deprecates `managed` (for adopted device) and `disable_auto_config` for claimed device) |
| `Nodes` | [`[]models.HaClusterConfigNode`](../../doc/models/ha-cluster-config-node.md) | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "site_id": "43e9c864-a7e4-4310-8031-d9817d2c5a43",
  "disable_auto_config": false,
  "managed": false,
  "mist_configured": false,
  "nodes": [
    {
      "mac": "mac0"
    },
    {
      "mac": "mac0"
    }
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

