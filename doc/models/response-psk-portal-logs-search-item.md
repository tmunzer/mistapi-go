
# Response Psk Portal Logs Search Item

## Structure

`ResponsePskPortalLogsSearchItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*uuid.UUID` | Optional | - |
| `Message` | `*string` | Optional | - |
| `NameId` | `*string` | Optional | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `PskId` | `*uuid.UUID` | Optional | - |
| `PskName` | `*string` | Optional | - |
| `PskportalId` | `*uuid.UUID` | Optional | - |
| `Timestamp` | `*float64` | Optional | - |

## Example (as JSON)

```json
{
  "id": "8a3dcaa7-80e3-4bb0-a75b-7bc6322cfd09",
  "message": "Rotate PSK test@mist.com",
  "name_id": "test@mist.com",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "psk_id": "608fe603-f9f0-4ce9-9473-04ef6c6ea749",
  "psk_name": "test@mist.com",
  "pskportal_id": "c1742c09-af35-4161-96ef-7dc65c6d5674",
  "timestamp": 1686346104.096
}
```

