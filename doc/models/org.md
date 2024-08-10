
# Org

## Structure

`Org`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AlarmtemplateId` | `models.Optional[uuid.UUID]` | Optional | - |
| `AllowMist` | `*bool` | Optional | **Default**: `true` |
| `CreatedTime` | `*float64` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | - |
| `ModifiedTime` | `*float64` | Optional | - |
| `MspId` | `*uuid.UUID` | Optional | - |
| `MspLogoUrl` | `*string` | Optional | logo uploaded by the MSP with advanced tier, only present if provided |
| `MspName` | `*string` | Optional | name of the msp the org belongs to |
| `Name` | `string` | Required | - |
| `OrggroupIds` | `[]uuid.UUID` | Optional | - |
| `SessionExpiry` | `*int` | Optional | **Default**: `1440`<br>**Constraints**: `>= 10`, `<= 20160` |

## Example (as JSON)

```json
{
  "allow_mist": true,
  "created_time": 1652905706.0,
  "id": "2818e386-8dec-2562-9ede-5b8a0fbbdc71",
  "modified_time": 1652905706.0,
  "msp_id": "b9d42c2e-88ee-41f8-b798-f009ce7fe909",
  "msp_logo_url": "https://example.com/logo/b9d42c2e-88ee-41f8-b798-f009ce7fe909.jpeg",
  "msp_name": "MSP",
  "name": "Org",
  "session_expiry": 1440,
  "alarmtemplate_id": "00000682-0000-0000-0000-000000000000"
}
```

