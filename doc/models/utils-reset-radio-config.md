
# Utils Reset Radio Config

*This model accepts additional fields of type interface{}.*

## Structure

`UtilsResetRadioConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Bands` | `[]string` | Required | list of bands |
| `Force` | `*bool` | Optional | whether to reset those with radio disabled. default is false (i.e. if user intentionally disables a radio, honor it)<br>**Default**: `false` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "bands": [
    "bands2",
    "bands3"
  ],
  "force": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

