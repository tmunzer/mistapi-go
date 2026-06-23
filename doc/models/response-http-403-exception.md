
# Response Http 403 Exception

Standard HTTP 403 permission error response

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseHttp403Exception`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Detail` | `*string` | Optional | Human-readable explanation of the permission error |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
if err != nil {
    switch typedErr := err.(type) {
    case *errors.ResponseHttp403Exception:
        log.Fatalln(typedErr)
    default:
        log.Fatalln(err)
    }
}
```

