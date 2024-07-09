
# Wlan Inject Dhcp Option 82

## Structure

`WlanInjectDhcpOption82`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CircuitId` | `*string` | Optional | - |
| `Enabled` | `*bool` | Optional | whether to inject option 82 when forwarding DHCP packets<br>**Default**: `false` |

## Example (as JSON)

```json
{
  "circuit_id": "{{SSID}}:{{AP_MAC}}",
  "enabled": false
}
```

