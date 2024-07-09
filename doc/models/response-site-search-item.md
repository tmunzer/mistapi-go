
# Response Site Search Item

## Structure

`ResponseSiteSearchItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AutoUpgradeEnabled` | `bool` | Required | - |
| `AutoUpgradeVersion` | `string` | Required | - |
| `CountryCode` | `*string` | Required | - |
| `HoneypotEnabled` | `bool` | Required | - |
| `Id` | `uuid.UUID` | Required | - |
| `Name` | `string` | Required | - |
| `OrgId` | `uuid.UUID` | Required | - |
| `SiteId` | `uuid.UUID` | Required | - |
| `Timestamp` | `float64` | Required | - |
| `Timezone` | `string` | Required | - |
| `VnaEnabled` | `bool` | Required | - |
| `WifiEnabled` | `bool` | Required | - |

## Example (as JSON)

```json
{
  "auto_upgrade_enabled": false,
  "auto_upgrade_version": "auto_upgrade_version0",
  "country_code": "country_code2",
  "honeypot_enabled": false,
  "id": "00001aec-0000-0000-0000-000000000000",
  "name": "name2",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "timestamp": 124.1,
  "timezone": "timezone2",
  "vna_enabled": false,
  "wifi_enabled": false
}
```

