
# Site Setting Skyatp

*This model accepts additional fields of type interface{}.*

## Structure

`SiteSettingSkyatp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | - |
| `SendIpMacMapping` | `*bool` | Optional | whether to send IP-MAC mapping to SkyATP<br>**Default**: `false` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "send_ip_mac_mapping": false,
  "enabled": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

