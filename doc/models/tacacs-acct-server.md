
# Tacacs Acct Server

*This model accepts additional fields of type interface{}.*

## Structure

`TacacsAcctServer`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Host` | `*string` | Optional | - |
| `Port` | `*string` | Optional | - |
| `Secret` | `*string` | Optional | - |
| `Timeout` | `*int` | Optional | **Default**: `10` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "timeout": 10,
  "host": "host4",
  "port": "port2",
  "secret": "secret8",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

