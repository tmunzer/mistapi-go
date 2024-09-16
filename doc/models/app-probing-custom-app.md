
# App Probing Custom App

## Structure

`AppProbingCustomApp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Address` | `*string` | Optional | required if `protocol`==`icmp` |
| `AppType` | `*string` | Optional | - |
| `Hostnames` | `[]string` | Optional | if `protocol`==`http` |
| `Key` | `*string` | Optional | - |
| `Name` | `*string` | Optional | - |
| `Network` | `*string` | Optional | - |
| `PacketSize` | `*int` | Optional | if `protocol`==`icmp`<br>**Constraints**: `>= 0`, `<= 65400` |
| `Protocol` | [`*models.AppProbingCustomAppProtocolEnum`](../../doc/models/app-probing-custom-app-protocol-enum.md) | Optional | enum: `http`, `icmp`<br>**Default**: `"http"` |
| `Url` | `*string` | Optional | if `protocol`==`http` |
| `Vrf` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "address": "192.168.1.1",
  "hostnames": [
    "http://www.abc.com"
  ],
  "name": "pos_app",
  "network": "lan",
  "protocol": "http",
  "url": "www.abc.com",
  "vrf": "lan",
  "app_type": "app_type8",
  "key": "key4"
}
```

