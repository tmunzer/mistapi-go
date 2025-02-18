
# Response Sso Failure Search Item

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseSsoFailureSearchItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Detail` | `string` | Required | - |
| `SamlAssertionXml` | `string` | Required | - |
| `Timestamp` | `float64` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "detail": "detail6",
  "saml_assertion_xml": "saml_assertion_xml4",
  "timestamp": 49.98,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

