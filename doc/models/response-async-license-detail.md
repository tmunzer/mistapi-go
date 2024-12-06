
# Response Async License Detail

detail claim status per device

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseAsyncLicenseDetail`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mac` | `*string` | Optional | device MAC Address |
| `Status` | `*string` | Optional | - |
| `Timestamp` | `*float64` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "mac": "mac0",
  "status": "status8",
  "timestamp": 238.64,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

