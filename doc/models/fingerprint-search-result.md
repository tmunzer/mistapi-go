
# Fingerprint Search Result

*This model accepts additional fields of type interface{}.*

## Structure

`FingerprintSearchResult`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | - |
| `Limit` | `int` | Required | - |
| `Next` | `*string` | Optional | - |
| `Results` | [`[]models.Fingerprint`](../../doc/models/fingerprint.md) | Required | - |
| `Start` | `int` | Required | - |
| `Total` | `int` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "end": 1711035686,
  "limit": 10,
  "results": [
    {
      "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
      "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
      "family": "family8",
      "mac": "mac0",
      "mfg": "mfg6",
      "model": "model4",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "start": 1710949286,
  "total": 232,
  "next": "next0",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

