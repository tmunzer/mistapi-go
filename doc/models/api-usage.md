
# Api Usage

*This model accepts additional fields of type interface{}.*

## Structure

`ApiUsage`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `RequestLimit` | `int` | Required | max number of request permitted<br><br>**Default**: `5000` |
| `Requests` | `int` | Required | num of request made in the current hour |
| `Seconds` | `*float64` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "request_limit": 5000,
  "requests": 188,
  "seconds": 94.24,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

