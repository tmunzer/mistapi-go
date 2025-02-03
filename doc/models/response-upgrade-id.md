
# Response Upgrade Id

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseUpgradeId`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `UpgradeId` | `uuid.UUID` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "upgrade_id": "4316c116-0acb-4c43-8f06-6723154e741e",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

