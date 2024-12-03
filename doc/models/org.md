
# Org

*This model accepts additional fields of type interface{}.*

## Structure

`Org`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AlarmtemplateId` | `models.Optional[uuid.UUID]` | Optional | - |
| `AllowMist` | `*bool` | Optional | **Default**: `true` |
| `CreatedTime` | `*float64` | Optional | when the object has been created, in epoch |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organnization |
| `ModifiedTime` | `*float64` | Optional | when the object has been modified for the last time, in epoch |
| `MspId` | `*uuid.UUID` | Optional | - |
| `MspLogoUrl` | `*string` | Optional | logo uploaded by the MSP with advanced tier, only present if provided |
| `MspName` | `*string` | Optional | name of the msp the org belongs to |
| `Name` | `string` | Required | - |
| `OrggroupIds` | `[]uuid.UUID` | Optional | - |
| `SessionExpiry` | `*int` | Optional | **Default**: `1440`<br>**Constraints**: `>= 10`, `<= 20160` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "allow_mist": true,
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "msp_id": "b9d42c2e-88ee-41f8-b798-f009ce7fe909",
  "msp_logo_url": "https://example.com/logo/b9d42c2e-88ee-41f8-b798-f009ce7fe909.jpeg",
  "msp_name": "MSP",
  "name": "Org",
  "session_expiry": 1440,
  "alarmtemplate_id": "00000682-0000-0000-0000-000000000000",
  "created_time": 36.26,
  "modified_time": 42.7,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

