
# Response Psk Portal Logs Search Item

*This model accepts additional fields of type interface{}.*

## Structure

`ResponsePskPortalLogsSearchItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organization |
| `Message` | `*string` | Optional | - |
| `NameId` | `*string` | Optional | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `PskId` | `*uuid.UUID` | Optional | - |
| `PskName` | `*string` | Optional | - |
| `PskportalId` | `*uuid.UUID` | Optional | - |
| `Timestamp` | `*float64` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "message": "Rotate PSK test@mist.com",
  "name_id": "test@mist.com",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "psk_id": "608fe603-f9f0-4ce9-9473-04ef6c6ea749",
  "psk_name": "test@mist.com",
  "pskportal_id": "c1742c09-af35-4161-96ef-7dc65c6d5674",
  "timestamp": 1686346104.096,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

