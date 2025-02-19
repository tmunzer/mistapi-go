
# Response Client Nac Search

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseClientNacSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*int` | Optional | - |
| `Limit` | `*int` | Optional | - |
| `Results` | [`[]models.ClientNac`](../../doc/models/client-nac.md) | Optional | - |
| `Start` | `*int` | Optional | - |
| `Total` | `*int` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "end": 1513362753,
  "limit": 3,
  "start": 1513276353,
  "total": 2,
  "results": [
    {
      "ap": [
        "ap6",
        "ap7",
        "ap8"
      ],
      "auth_type": "peap-tls",
      "cert_cn": [
        "cert_cn9",
        "cert_cn8"
      ],
      "cert_issuer": [
        "cert_issuer2"
      ],
      "cert_serial": [
        "cert_serial8"
      ],
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

