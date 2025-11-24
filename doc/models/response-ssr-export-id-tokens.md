
# Response Ssr Export Id Tokens

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseSsrExportIdTokens`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Results` | [`[]models.ResponseSsrExportIdTokensResultsItem`](../../doc/models/response-ssr-export-id-tokens-results-item.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "results": [
    {
      "idtoken": "idtoken8",
      "mac": "mac0",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "idtoken": "idtoken8",
      "mac": "mac0",
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

