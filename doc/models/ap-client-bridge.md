
# Ap Client Bridge

## Structure

`ApClientBridge`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Auth` | [`*models.ApClientBridgeAuth`](../../doc/models/ap-client-bridge-auth.md) | Optional | - |
| `Enabled` | `*bool` | Optional | when acted as client bridge:<br><br>* only 5G radio can be used<br>* will not serve as AP on any radios<br>**Default**: `false` |
| `Ssid` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |

## Example (as JSON)

```json
{
  "enabled": false,
  "ssid": "Uplink-SSID",
  "auth": {
    "psk": "psk4",
    "type": "open"
  }
}
```

