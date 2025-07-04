
# Sdkinvite

SDK invite

*This model accepts additional fields of type interface{}.*

## Structure

`Sdkinvite`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional | When the object has been created, in epoch |
| `Enabled` | `*bool` | Optional | **Default**: `true` |
| `ExpireTime` | `*int` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organization |
| `ModifiedTime` | `*float64` | Optional | When the object has been modified for the last time, in epoch |
| `Name` | `string` | Required | Name, will show up in mobile |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Quota` | `*int` | Optional | Number of time this invite can be used |
| `QuotaLimited` | `*bool` | Optional | Whether quota limiting is enabled<br><br>**Default**: `false` |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "enabled": true,
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "name": "name4",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "quota_limited": false,
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "created_time": 203.14,
  "expire_time": 100,
  "modified_time": 131.82,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

