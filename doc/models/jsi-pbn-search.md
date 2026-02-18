
# Jsi Pbn Search

PBN search response

## Structure

`JsiPbnSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*int` | Optional | End timestamp |
| `Limit` | `*int` | Optional | Number of results to return |
| `Next` | `*string` | Optional | Next page URL |
| `Results` | [`[]models.JsiPbnItem`](../../doc/models/jsi-pbn-item.md) | Optional | List of PBN advisories |
| `Start` | `*int` | Optional | Start timestamp |
| `Total` | `*int` | Optional | Total number of results |

## Example (as JSON)

```json
{
  "end": 194,
  "limit": 232,
  "next": "next8",
  "results": [
    {
      "bug_type": "bug_type0",
      "customer_risk": "customer_risk4",
      "fixed_in": "fixed_in8",
      "id": "id6",
      "introduced_in": "introduced_in2"
    }
  ],
  "start": 152
}
```

