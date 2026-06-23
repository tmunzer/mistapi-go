
# Response Http 429 Exception

Standard HTTP 429 rate limit error response

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseHttp429Exception`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Detail` | `*string` | Optional | Human-readable explanation of the rate limit error |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
if err != nil {
    switch typedErr := err.(type) {
    case *errors.ResponseHttp429Exception:
        log.Fatalln(typedErr)
    default:
        log.Fatalln(err)
    }
}
```

