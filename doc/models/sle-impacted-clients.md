
# Sle Impacted Clients

*This model accepts additional fields of type interface{}.*

## Structure

`SleImpactedClients`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Classifier` | `*string` | Optional | - |
| `Clients` | [`[]models.SleImpactedClientsClient`](../../doc/models/sle-impacted-clients-client.md) | Optional | - |
| `End` | `*int` | Optional | - |
| `Failure` | `*string` | Optional | - |
| `Limit` | `*int` | Optional | - |
| `Metric` | `*string` | Optional | - |
| `Page` | `*int` | Optional | - |
| `Start` | `*int` | Optional | - |
| `TotalCount` | `*int` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "classifier": "classifier6",
  "clients": [
    {
      "degraded": 18,
      "duration": 124,
      "mac": "mac2",
      "name": "name8",
      "switches": [
        {
          "interfaces": [
            "interfaces9"
          ],
          "switch_mac": "switch_mac6",
          "switch_name": "switch_name0",
          "exampleAdditionalProperty": {
            "key1": "val1",
            "key2": "val2"
          }
        },
        {
          "interfaces": [
            "interfaces9"
          ],
          "switch_mac": "switch_mac6",
          "switch_name": "switch_name0",
          "exampleAdditionalProperty": {
            "key1": "val1",
            "key2": "val2"
          }
        },
        {
          "interfaces": [
            "interfaces9"
          ],
          "switch_mac": "switch_mac6",
          "switch_name": "switch_name0",
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
    },
    {
      "degraded": 18,
      "duration": 124,
      "mac": "mac2",
      "name": "name8",
      "switches": [
        {
          "interfaces": [
            "interfaces9"
          ],
          "switch_mac": "switch_mac6",
          "switch_name": "switch_name0",
          "exampleAdditionalProperty": {
            "key1": "val1",
            "key2": "val2"
          }
        },
        {
          "interfaces": [
            "interfaces9"
          ],
          "switch_mac": "switch_mac6",
          "switch_name": "switch_name0",
          "exampleAdditionalProperty": {
            "key1": "val1",
            "key2": "val2"
          }
        },
        {
          "interfaces": [
            "interfaces9"
          ],
          "switch_mac": "switch_mac6",
          "switch_name": "switch_name0",
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
  ],
  "end": 136,
  "failure": "failure4",
  "limit": 222,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

