
# Org Setting Mgmt

management-related properties

## Structure

`OrgSettingMgmt`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MxtunnelIds` | `[]uuid.UUID` | Optional | list of Mist Tunnels |
| `UseMxtunnel` | `*bool` | Optional | whether to use Mist Tunnel for mgmt connectivity, this takes precedence over use_wxtunnel<br>**Default**: `false` |
| `UseWxtunnel` | `*bool` | Optional | whether to use wxtunnel for mgmt connectivity<br>**Default**: `false` |

## Example (as JSON)

```json
{
  "use_mxtunnel": false,
  "use_wxtunnel": false,
  "mxtunnel_ids": [
    "0000014e-0000-0000-0000-000000000000"
  ]
}
```
