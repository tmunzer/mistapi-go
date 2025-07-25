
# Response Mxedge Search

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseMxedgeSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*int` | Optional | - |
| `Limit` | `*int` | Optional | - |
| `Results` | [`[]models.StatsMxedge`](../../doc/models/stats-mxedge.md) | Optional | - |
| `Start` | `*int` | Optional | - |
| `Total` | `*int` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

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
          "key0": null
        },
        "idle": 224,
        "interrupt": 80,
        "system": 80,
        "usage": 46,
        "exampleAdditionalProperty": {
          "key1": "val1",
          "key2": "val2"
        }
      },
      "created_time": 73.76,
      "fips_enabled": false,
      "for_site": false,
      "fwupdate": {
        "progress": 100,
        "status": "inprogress",
        "status_id": 70,
        "timestamp": 147.68,
        "will_retry": false,
        "exampleAdditionalProperty": {
          "key1": "val1",
          "key2": "val2"
        }
      },
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "cpu_stat": {
        "cpus": {
          "key0": null
        },
        "idle": 224,
        "interrupt": 80,
        "system": 80,
        "usage": 46,
        "exampleAdditionalProperty": {
          "key1": "val1",
          "key2": "val2"
        }
      },
      "created_time": 73.76,
      "fips_enabled": false,
      "for_site": false,
      "fwupdate": {
        "progress": 100,
        "status": "inprogress",
        "status_id": 70,
        "timestamp": 147.68,
        "will_retry": false,
        "exampleAdditionalProperty": {
          "key1": "val1",
          "key2": "val2"
        }
      },
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "cpu_stat": {
        "cpus": {
          "key0": null
        },
        "idle": 224,
        "interrupt": 80,
        "system": 80,
        "usage": 46,
        "exampleAdditionalProperty": {
          "key1": "val1",
          "key2": "val2"
        }
      },
      "created_time": 73.76,
      "fips_enabled": false,
      "for_site": false,
      "fwupdate": {
        "progress": 100,
        "status": "inprogress",
        "status_id": 70,
        "timestamp": 147.68,
        "will_retry": false,
        "exampleAdditionalProperty": {
          "key1": "val1",
          "key2": "val2"
        }
      },
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

