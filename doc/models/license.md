
# License

License

*This model accepts additional fields of type interface{}.*

## Structure

`License`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Amendments` | [`[]models.LicenseAmendment`](../../doc/models/license-amendment.md) | Optional | **Constraints**: *Unique Items Required* |
| `Entitled` | `map[string]int` | Optional | Property key is license type (e.g. SUB-MAN) and Property value is the number of licenses entitled. |
| `Licenses` | [`[]models.LicenseSub`](../../doc/models/license-sub.md) | Optional | - |
| `Summary` | `map[string]int` | Optional | Property key is license type (e.g. SUB-MAN) and Property value is the number of licenses consumed. |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "amendments": [
    {
      "created_time": 132.88,
      "end_time": 210,
      "id": "00000292-0000-0000-0000-000000000000",
      "modified_time": 202.08,
      "quantity": 182,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "created_time": 132.88,
      "end_time": 210,
      "id": "00000292-0000-0000-0000-000000000000",
      "modified_time": 202.08,
      "quantity": 182,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "entitled": {
    "key0": 120,
    "key1": 119
  },
  "licenses": [
    {
      "created_time": 186.08,
      "end_time": 154,
      "id": "0000017a-0000-0000-0000-000000000000",
      "modified_time": 148.88,
      "order_id": "order_id2",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "created_time": 186.08,
      "end_time": 154,
      "id": "0000017a-0000-0000-0000-000000000000",
      "modified_time": 148.88,
      "order_id": "order_id2",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "summary": {
    "key0": 22
  },
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

