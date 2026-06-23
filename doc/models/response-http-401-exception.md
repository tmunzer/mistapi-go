
# Response Http 401 Exception

Standard HTTP 401 authentication error response

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseHttp401Exception`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Detail` | `*string` | Optional | Human-readable explanation of the authentication error |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
if err != nil {
    switch typedErr := err.(type) {
    case *errors.ResponseHttp401Exception:
        log.Fatalln(typedErr)
    default:
        log.Fatalln(err)
    }
}
```

