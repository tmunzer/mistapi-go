
# Response Site Search Item

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseSiteSearchItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AutoUpgradeEnabled` | `bool` | Required | - |
| `AutoUpgradeVersion` | `string` | Required | - |
| `CountryCode` | `models.Optional[string]` | Optional | - |
| `HoneypotEnabled` | `bool` | Required | - |
| `Id` | `uuid.UUID` | Required | Unique ID of the object instance in the Mist Organization |
| `Name` | `string` | Required | - |
| `OrgId` | `uuid.UUID` | Required | - |
| `SiteId` | `uuid.UUID` | Required | - |
| `Timestamp` | `float64` | Required | Epoch (seconds) |
| `Timezone` | `string` | Required | - |
| `VnaEnabled` | `bool` | Required | - |
| `WifiEnabled` | `bool` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "auto_upgrade_enabled": false,
  "auto_upgrade_version": "auto_upgrade_version4",
  "honeypot_enabled": false,
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "name": "name6",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "timestamp": 198.34,
  "timezone": "timezone6",
  "vna_enabled": false,
  "wifi_enabled": false,
  "country_code": "country_code6",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

