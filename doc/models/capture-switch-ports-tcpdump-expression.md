
# Capture Switch Ports Tcpdump Expression

## Structure

`CaptureSwitchPortsTcpdumpExpression`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `TcpdumpExpression` | `*string` | Optional | tcpdump expression, port specific if specified under ports dict, otherwise applicable across ports if specified at top level of payload. Port specific value overrides top level value when both exist. |

## Example (as JSON)

```json
{
  "tcpdump_expression": "port 443"
}
```

