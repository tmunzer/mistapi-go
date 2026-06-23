
# Response Self Oauth Link Failure Exception

OAuth2 account-linking failure response

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseSelfOauthLinkFailureException`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Error` | `string` | Required | OAuth2 error code returned during account linking |
| `ErrorDescription` | `string` | Required | Human-readable description of the OAuth2 account-linking failure |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
if err != nil {
    switch typedErr := err.(type) {
    case *errors.ResponseSelfOauthLinkFailureException:
        log.Fatalln(typedErr)
    default:
        log.Fatalln(err)
    }
}
```

