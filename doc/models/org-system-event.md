
# Org System Event

## Structure

`OrgSystemEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ChangeCat` | `*string` | Optional | - |
| `Metadata` | `*string` | Optional | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Scope` | `*string` | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Timestamp` | `*float64` | Optional | Epoch (seconds) |
| `Type` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "change_cat": "admin_action",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "scope": "org",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "type": "delete-wlan",
  "metadata": "metadata2"
}
```

