
# Alarms

## Structure

`Alarms`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AlarmIds` | `[]uuid.UUID` | Required | - |
| `Note` | `*string` | Optional | Some text note describing the intent |

## Example (as JSON)

```json
{
  "alarm_ids": [
    "00000638-0000-0000-0000-000000000000",
    "00000639-0000-0000-0000-000000000000",
    "0000063a-0000-0000-0000-000000000000"
  ],
  "note": "note0"
}
```

