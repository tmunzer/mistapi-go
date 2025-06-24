
# Org Setting Mgmt

management-related properties

*This model accepts additional fields of type interface{}.*

## Structure

`OrgSettingMgmt`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MxtunnelIds` | `[]uuid.UUID` | Optional | List of Mist Tunnels |
| `UseMxtunnel` | `*bool` | Optional | Whether to use Mist Tunnel for mgmt connectivity, this takes precedence over use_wxtunnel<br><br>**Default**: `false` |
| `UseWxtunnel` | `*bool` | Optional | Whether to use wxtunnel for mgmt connectivity<br><br>**Default**: `false` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "use_mxtunnel": false,
  "use_wxtunnel": false,
  "mxtunnel_ids": [
    "000015a6-0000-0000-0000-000000000000",
    "000015a7-0000-0000-0000-000000000000"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

