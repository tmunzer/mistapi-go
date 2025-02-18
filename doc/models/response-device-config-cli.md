
# Response Device Config Cli

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseDeviceConfigCli`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Cli` | `[]string` | Required | **Constraints**: *Unique Items Required* |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "cli": [
    "cli0",
    "cli9",
    "cli8"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

