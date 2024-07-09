
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
  "created_time": 247.58,
  "id": "00000850-0000-0000-0000-000000000000",
  "modified_time": 87.38,
  "msp_id": "0000064c-0000-0000-0000-000000000000",
  "name": "name8",
  "org_ids": [
    "00001cd8-0000-0000-0000-000000000000",
    "00001cd7-0000-0000-0000-000000000000"
  ]
}
```

