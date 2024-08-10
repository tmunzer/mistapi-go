
# Response Client Nac Search

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
      "auth_type": "auth_type4",
      "cert_cn": [
        "cert_cn9",
        "cert_cn8"
      ],
      "cert_issuer": [
        "cert_issuer2"
      ],
      "idp_id": "idp_id6"
    }
  ]
}
```

