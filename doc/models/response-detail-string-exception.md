
# Response Detail String Exception

Response containing a human-readable detail message

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseDetailStringException`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Detail` | `*string` | Optional | Human-readable detail message returned by the API |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
if err != nil {
    switch typedErr := err.(type) {
    case *errors.ResponseDetailStringException:
        log.Fatalln(typedErr)
    default:
        log.Fatalln(err)
    }
}
```

