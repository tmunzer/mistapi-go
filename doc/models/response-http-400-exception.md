
# Response Http 400 Exception

Standard HTTP 400 bad request error response

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseHttp400Exception`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Detail` | `*string` | Optional | Human-readable explanation of the bad request error |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
if err != nil {
    switch typedErr := err.(type) {
    case *errors.ResponseHttp400Exception:
        log.Fatalln(typedErr)
    default:
        log.Fatalln(err)
    }
}
```

