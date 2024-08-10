
# Response Mxedge Search

## Structure

`ResponseMxedgeSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*int` | Optional | - |
| `Limit` | `*int` | Optional | - |
| `Results` | [`[]models.MxedgeStats`](../../doc/models/mxedge-stats.md) | Optional | - |
| `Start` | `*int` | Optional | - |
| `Total` | `*int` | Optional | - |

## Example (as JSON)

```json
{
  "end": 1694708579,
  "limit": 10,
  "start": 1694622179,
  "total": 2,
  "results": [
    {
      "cpu_stat": {
        "cpus": {
          "key0": {
            "idle": 4.8,
            "interrupt": 118.56,
            "load_avg": [
              8.63
            ],
            "system": 110.88,
            "user": 107.24
          }
        },
        "idle": 224,
        "interrupt": 80,
        "system": 80,
        "usage": 46
      },
      "created_time": 73.76,
      "fips_enabled": false,
      "for_site": false,
      "id": "000023ba-0000-0000-0000-000000000000"
    },
    {
      "cpu_stat": {
        "cpus": {
          "key0": {
            "idle": 4.8,
            "interrupt": 118.56,
            "load_avg": [
              8.63
            ],
            "system": 110.88,
            "user": 107.24
          }
        },
        "idle": 224,
        "interrupt": 80,
        "system": 80,
        "usage": 46
      },
      "created_time": 73.76,
      "fips_enabled": false,
      "for_site": false,
      "id": "000023ba-0000-0000-0000-000000000000"
    },
    {
      "cpu_stat": {
        "cpus": {
          "key0": {
            "idle": 4.8,
            "interrupt": 118.56,
            "load_avg": [
              8.63
            ],
            "system": 110.88,
            "user": 107.24
          }
        },
        "idle": 224,
        "interrupt": 80,
        "system": 80,
        "usage": 46
      },
      "created_time": 73.76,
      "fips_enabled": false,
      "for_site": false,
      "id": "000023ba-0000-0000-0000-000000000000"
    }
  ]
}
```

