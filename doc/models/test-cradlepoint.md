
# Test Cradlepoint

*This model accepts additional fields of type interface{}.*

## Structure

`TestCradlepoint`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Error` | `*string` | Optional | if status is `inactive` this field returns the reason for it being inactive. |
| `LastStatus` | [`*models.TestCradlepointLastStatusEnum`](../../doc/models/test-cradlepoint-last-status-enum.md) | Optional | status of integration detected during last sync. enum: `active`, `inactive` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "error": "Cradlepoint API keys are no longer valid, please verify and update the keys under organization settings.",
  "last_status": "inactive",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

