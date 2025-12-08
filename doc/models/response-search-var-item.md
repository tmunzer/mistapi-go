
# Response Search Var Item

## Structure

`ResponseSearchVarItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional | When the object has been created, in epoch |
| `ModifiedTime` | `*float64` | Optional | When the object has been modified for the last time, in epoch |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Src` | `*string` | Optional | - |
| `Var` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "created_time": 217.6,
  "modified_time": 117.36,
  "src": "src6"
}
```

