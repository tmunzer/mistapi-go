
# Capture Switch Ports Tcpdump Expression

*This model accepts additional fields of type interface{}.*

## Structure

`CaptureSwitchPortsTcpdumpExpression`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `TcpdumpExpression` | `*string` | Optional | tcpdump expression, port specific if specified under ports dict, otherwise applicable across ports if specified at top level of payload. Port specific value overrides top level value when both exist. |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "tcpdump_expression": "port 443",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

