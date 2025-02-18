
# Secintel Profile

*This model accepts additional fields of type interface{}.*

## Structure

`SecintelProfile`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Name` | `*string` | Optional | - |
| `Profiles` | [`[]models.SecintelProfileProfile`](../../doc/models/secintel-profile-profile.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "name": "secintel-custom",
  "profiles": [
    {
      "action": "strict",
      "category": "IH",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "action": "strict",
      "category": "IH",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

