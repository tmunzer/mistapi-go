
# App Probing Custom App

## Structure

`AppProbingCustomApp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Address` | `*string` | Optional | if `protocol`==`icmp` |
| `AppType` | `*string` | Optional | - |
| `Hostname` | `[]string` | Optional | if `protocol`==`http` |
| `Name` | `*string` | Optional | - |
| `Network` | `*string` | Optional | - |
| `Protocol` | [`*models.AppProbingCustomAppProtocolEnum`](../../doc/models/app-probing-custom-app-protocol-enum.md) | Optional | enum: `http`, `icmp`<br>**Default**: `"http"` |
| `Url` | `*string` | Optional | if `protocol`==`http` |
| `Vrf` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "address": "192.168.1.1",
  "hostname": [
    "www.abc.com"
  ],
  "name": "pos_app",
  "network": "lan",
  "protocol": "http",
  "url": "www.abc.com",
  "vrf": "lan",
  "app_type": "app_type8"
}
```

