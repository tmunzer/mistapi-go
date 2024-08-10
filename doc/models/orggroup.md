
# Orggroup

Organizations Group

## Structure

`Orggroup`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | - |
| `ModifiedTime` | `*float64` | Optional | - |
| `MspId` | `*uuid.UUID` | Optional | - |
| `Name` | `string` | Required | - |
| `OrgIds` | `[]uuid.UUID` | Optional | - |

## Example (as JSON)

```json
{
  "created_time": 12.46,
  "id": "00000bc8-0000-0000-0000-000000000000",
  "modified_time": 66.5,
  "msp_id": "000009c4-0000-0000-0000-000000000000",
  "name": "name6",
  "org_ids": [
    "00001d40-0000-0000-0000-000000000000",
    "00001d41-0000-0000-0000-000000000000"
  ]
}
```

