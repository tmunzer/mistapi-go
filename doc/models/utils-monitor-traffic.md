
# Utils Monitor Traffic

*This model accepts additional fields of type interface{}.*

## Structure

`UtilsMonitorTraffic`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Port` | `*string` | Optional | Port name, if no port input is provided then all ports will be monitored |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "port": "ge-0/0/1",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

