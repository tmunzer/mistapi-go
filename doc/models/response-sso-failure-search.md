
# Response Sso Failure Search

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseSsoFailureSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Results` | [`[]models.ResponseSsoFailureSearchItem`](../../doc/models/response-sso-failure-search-item.md) | Required | **Constraints**: *Unique Items Required* |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "results": [
    {
      "detail": "detail2",
      "saml_assertion_xml": "saml_assertion_xml0",
      "timestamp": 2.64,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

