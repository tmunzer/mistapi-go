
# Alarm Ack

## Structure

`AlarmAck`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AlarmIds` | `[]uuid.UUID` | Required | - |
| `Note` | `*string` | Optional | Some text note describing the intent |

## Example (as JSON)

```json
{
  "alarm_ids": [
    "ccb8c94d-ca56-4075-932f-1f2ab444ff2c",
    "98ff4a3d-ec9b-4138-a42e-54fc3335179d"
  ],
  "note": "maintenance window"
}
```

