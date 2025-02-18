
# Site App

*This model accepts additional fields of type interface{}.*

## Structure

`SiteApp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Group` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `Key` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `Name` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "group": "group6",
  "key": "key8",
  "name": "name8",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

