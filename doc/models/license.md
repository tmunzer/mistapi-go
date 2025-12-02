
# License

License

## Structure

`License`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Amendments` | [`[]models.LicenseAmendment`](../../doc/models/license-amendment.md) | Optional | **Constraints**: *Unique Items Required* |
| `Entitled` | `map[string]int` | Optional | Property key is license type (e.g. SUB-MAN) and Property value is the number of licenses entitled. |
| `FullyLoaded` | `map[string]int` | Optional | Maximum number of licenses that may be required if the service is enabled on all the Organization Devices. Property key is the service name (e.g. "SUB-MAN"). |
| `Licenses` | [`[]models.LicenseSub`](../../doc/models/license-sub.md) | Optional | - |
| `Summary` | `map[string]int` | Optional | Number of licenses currently consumed. Property key is license type (e.g. SUB-MAN). |
| `Usages` | `map[string]int` | Optional | Number of available licenes. Property key is the service name (e.g. "SUB-MAN"). name (e.g. "SUB-MAN") |

## Example (as JSON)

```json
{
  "amendments": [
    {
      "created_time": 132.88,
      "end_time": 210,
      "id": "00000292-0000-0000-0000-000000000000",
      "modified_time": 202.08,
      "quantity": 182
    },
    {
      "created_time": 132.88,
      "end_time": 210,
      "id": "00000292-0000-0000-0000-000000000000",
      "modified_time": 202.08,
      "quantity": 182
    }
  ],
  "entitled": {
    "key0": 120,
    "key1": 119
  },
  "fully_loaded": {
    "key0": 214,
    "key1": 215,
    "key2": 216
  },
  "licenses": [
    {
      "created_time": 186.08,
      "end_time": 154,
      "id": "0000017a-0000-0000-0000-000000000000",
      "modified_time": 148.88,
      "order_id": "order_id2"
    },
    {
      "created_time": 186.08,
      "end_time": 154,
      "id": "0000017a-0000-0000-0000-000000000000",
      "modified_time": 148.88,
      "order_id": "order_id2"
    }
  ],
  "summary": {
    "key0": 22
  }
}
```

