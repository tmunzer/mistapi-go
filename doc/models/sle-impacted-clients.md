
# Sle Impacted Clients

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
          "switch_name": "switch_name0"
        },
        {
          "interfaces": [
            "interfaces9"
          ],
          "switch_mac": "switch_mac6",
          "switch_name": "switch_name0"
        },
        {
          "interfaces": [
            "interfaces9"
          ],
          "switch_mac": "switch_mac6",
          "switch_name": "switch_name0"
        }
      ]
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
          "switch_name": "switch_name0"
        },
        {
          "interfaces": [
            "interfaces9"
          ],
          "switch_mac": "switch_mac6",
          "switch_name": "switch_name0"
        },
        {
          "interfaces": [
            "interfaces9"
          ],
          "switch_mac": "switch_mac6",
          "switch_name": "switch_name0"
        }
      ]
    }
  ],
  "end": 136,
  "failure": "failure4",
  "limit": 222
}
```

