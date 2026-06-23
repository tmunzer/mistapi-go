
# Response Http 400 Webhook Exception

Webhook-specific HTTP 400 bad request error response

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseHttp400WebhookException`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Detail` | `*string` | Optional | Human-readable explanation of the invalid webhook request |
| `Reason` | `*string` | Optional | Additional reason explaining why the webhook request was invalid |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
if err != nil {
    switch typedErr := err.(type) {
    case *errors.ResponseHttp400WebhookException:
        log.Fatalln(typedErr)
    default:
        log.Fatalln(err)
    }
}
```

