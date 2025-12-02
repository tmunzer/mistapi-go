
# Sso Openroaming

if `idp_type`==`openroaming`

## Structure

`SsoOpenroaming`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ssids` | `[]string` | Required | SSIDs that support OpenRoaming |
| `WbaCert` | `*string` | Optional | Optional WBA-issued certificate. If not provided, the default WBA-issued certificate for Juniper will be used. |

## Example (as JSON)

```json
{
  "ssids": [
    "ssid_name1",
    "ssid_name2"
  ],
  "wba_cert": "-----BEGIN CERTIFICATE-----\\nMIIFZjCCA06gAwIBAgIIP61/1qm/uDowDQYJKoZIhvcNAQELBQE\\n-----END CERTIFICATE-----"
}
```

