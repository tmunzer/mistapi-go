
# Response Troubleshoot

## Structure

`ResponseTroubleshoot`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*int` | Optional | - |
| `Results` | [`[]models.ResponseTroubleshootItem`](../../doc/models/response-troubleshoot-item.md) | Optional | - |
| `Start` | `*int` | Optional | - |

## Example (as JSON)

```json
{
  "end": 1655151856,
  "start": 1655065456,
  "results": [
    {
      "category": "category4",
      "reason": "reason8",
      "recommendation": "recommendation8",
      "text": "text4"
    }
  ]
}
```
