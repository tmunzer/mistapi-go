
# App Probing Custom App

*This model accepts additional fields of type interface{}.*

## Structure

`AppProbingCustomApp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Address` | `*string` | Optional | Required if `protocol`==`icmp` |
| `AppType` | `*string` | Optional | - |
| `Hostnames` | `[]string` | Optional | If `protocol`==`http` |
| `Key` | `*string` | Optional | - |
| `Name` | `*string` | Optional | - |
| `Network` | `*string` | Optional | - |
| `PacketSize` | `*int` | Optional | If `protocol`==`icmp`<br><br>**Constraints**: `>= 0`, `<= 65400` |
| `Protocol` | [`*models.AppProbingCustomAppProtocolEnum`](../../doc/models/app-probing-custom-app-protocol-enum.md) | Optional | enum: `http`, `icmp`<br><br>**Default**: `"http"` |
| `Url` | `*string` | Optional | If `protocol`==`http` |
| `Vrf` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "address": "192.168.1.1",
  "hostnames": [
    "https://www.abc.com"
  ],
  "name": "pos_app",
  "network": "lan",
  "protocol": "http",
  "url": "www.abc.com",
  "vrf": "lan",
  "app_type": "app_type8",
  "key": "key4",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

