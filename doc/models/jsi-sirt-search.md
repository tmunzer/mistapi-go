
# Jsi Sirt Search

SIRT search response

## Structure

`JsiSirtSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*int` | Optional | End timestamp |
| `Limit` | `*int` | Optional | Number of results to return |
| `Next` | `*string` | Optional | Next page URL |
| `Results` | [`[]models.JsiSirtItem`](../../doc/models/jsi-sirt-item.md) | Optional | List of SIRT advisories |
| `Start` | `*int` | Optional | Start timestamp |
| `Total` | `*int` | Optional | Total number of results |

## Example (as JSON)

```json
{
  "end": 168,
  "limit": 2,
  "next": "next2",
  "results": [
    {
      "cvss_score": 231.84,
      "id": "id6",
      "models": [
        "models2",
        "models3"
      ],
      "problem": "problem6",
      "published_date": 160
    },
    {
      "cvss_score": 231.84,
      "id": "id6",
      "models": [
        "models2",
        "models3"
      ],
      "problem": "problem6",
      "published_date": 160
    },
    {
      "cvss_score": 231.84,
      "id": "id6",
      "models": [
        "models2",
        "models3"
      ],
      "problem": "problem6",
      "published_date": 160
    }
  ],
  "start": 126
}
```

