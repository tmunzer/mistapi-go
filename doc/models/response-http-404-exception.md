
# Response Http 404 Exception

Standard HTTP 404 not found error response

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseHttp404Exception`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*string` | Optional | Missing resource identifier, when the API includes one |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
if err != nil {
    switch typedErr := err.(type) {
    case *errors.ResponseHttp404Exception:
        log.Fatalln(typedErr)
    default:
        log.Fatalln(err)
    }
}
```

