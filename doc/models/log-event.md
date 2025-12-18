
# Log Event

## Structure

`LogEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AdminId` | `models.Optional[uuid.UUID]` | Optional | admin id |
| `AdminName` | `models.Optional[string]` | Optional | Name of the admin that performs the action |
| `After` | `*interface{}` | Optional | field values after the change |
| `Before` | `*interface{}` | Optional | field values prior to the change |
| `DeviceId` | `models.Optional[uuid.UUID]` | Optional | Device id |
| `ForSite` | `*bool` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organization |
| `Message` | `string` | Required | log message |
| `OrgId` | `uuid.UUID` | Required | - |
| `SiteId` | `models.Optional[uuid.UUID]` | Optional | - |
| `SrcIp` | `*string` | Optional | sender source ip address |
| `Timestamp` | `float64` | Required | Epoch (seconds) |

## Example (as JSON)

```json
{
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "message": "message8",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "timestamp": 29.66,
  "admin_id": "00002082-0000-0000-0000-000000000000",
  "admin_name": "admin_name6",
  "after": {
    "key1": "val1",
    "key2": "val2"
  },
  "before": {
    "key1": "val1",
    "key2": "val2"
  },
  "device_id": "00000e6e-0000-0000-0000-000000000000"
}
```

