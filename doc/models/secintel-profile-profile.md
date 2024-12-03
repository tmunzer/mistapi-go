
# Secintel Profile Profile

*This model accepts additional fields of type interface{}.*

## Structure

`SecintelProfileProfile`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Action` | [`*models.SecintelProfileProfileActionEnum`](../../doc/models/secintel-profile-profile-action-enum.md) | Optional | enum: `default`, `standard`, `strict`<br>**Default**: `"default"` |
| `Category` | [`*models.SecintelProfileProfileCategoryEnum`](../../doc/models/secintel-profile-profile-category-enum.md) | Optional | enum: `CC`, `IH` (Infected Host), `DNS` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "action": "default",
  "category": "IH",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

