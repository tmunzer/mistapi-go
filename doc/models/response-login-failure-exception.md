
# Response Login Failure Exception

Login failure response returned when authentication cannot continue locally

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseLoginFailureException`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Detail` | `string` | Required | Human-readable login failure message returned by the API |
| `ForwardUrl` | `*string` | Optional | SSO URL where the user should continue login, when provided |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
if err != nil {
    switch typedErr := err.(type) {
    case *errors.ResponseLoginFailureException:
        log.Fatalln(typedErr)
    default:
        log.Fatalln(err)
    }
}
```

