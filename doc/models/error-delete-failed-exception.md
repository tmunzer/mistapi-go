
# Error Delete Failed Exception

Error response returned when an organization delete request is blocked

*This model accepts additional fields of type interface{}.*

## Structure

`ErrorDeleteFailedException`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Detail` | `string` | Required | Reason the delete request failed |
| `OrgId` | `uuid.UUID` | Required, Read-only | Unique identifier of a Mist organization |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
if err != nil {
    switch typedErr := err.(type) {
    case *errors.ErrorDeleteFailedException:
        log.Fatalln(typedErr)
    default:
        log.Fatalln(err)
    }
}
```

