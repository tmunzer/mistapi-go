
# Wlan Inject Dhcp Option 82

## Structure

`WlanInjectDhcpOption82`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CircuitId` | `*string` | Optional | information to set in the `circuit_id` field of the DHCP Option 82. It is possible to use static string or the following variables (e.g. `{{SSID}}:{{AP_MAC}}`):<br><br>* {{AP_MAC}}<br>* {{AP_MAC_DASHED}}<br>* {{AP_MODEL}}<br>* {{AP_NAME}}<br>* {{SITE_NAME}}<br>* {{SSID}} |
| `Enabled` | `*bool` | Optional | whether to inject option 82 when forwarding DHCP packets<br>**Default**: `false` |

## Example (as JSON)

```json
{
  "circuit_id": "{{SSID}}:{{AP_MAC}}",
  "enabled": false
}
```

