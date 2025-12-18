
# Gw Routing Policy

## Structure

`GwRoutingPolicy`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Terms` | [`[]models.GwRoutingPolicyTerm`](../../doc/models/gw-routing-policy-term.md) | Optional | zero or more criteria/filter can be specified to match the term, all criteria have to be met<br><br>**Constraints**: *Unique Items Required* |

## Example (as JSON)

```json
{
  "terms": [
    {
      "actions": {
        "accept": false,
        "add_community": [
          "add_community5",
          "add_community6",
          "add_community7"
        ],
        "add_target_vrfs": [
          "add_target_vrfs1"
        ],
        "community": [
          "community4",
          "community5"
        ],
        "exclude_as_path": [
          "exclude_as_path0",
          "exclude_as_path1",
          "exclude_as_path2"
        ]
      },
      "matching": {
        "as_path": [
          "String3"
        ],
        "community": [
          "community4"
        ],
        "network": [
          "network7",
          "network8",
          "network9"
        ],
        "prefix": [
          "prefix5",
          "prefix6",
          "prefix7"
        ],
        "protocol": [
          "aggregate",
          "bgp"
        ]
      }
    },
    {
      "actions": {
        "accept": false,
        "add_community": [
          "add_community5",
          "add_community6",
          "add_community7"
        ],
        "add_target_vrfs": [
          "add_target_vrfs1"
        ],
        "community": [
          "community4",
          "community5"
        ],
        "exclude_as_path": [
          "exclude_as_path0",
          "exclude_as_path1",
          "exclude_as_path2"
        ]
      },
      "matching": {
        "as_path": [
          "String3"
        ],
        "community": [
          "community4"
        ],
        "network": [
          "network7",
          "network8",
          "network9"
        ],
        "prefix": [
          "prefix5",
          "prefix6",
          "prefix7"
        ],
        "protocol": [
          "aggregate",
          "bgp"
        ]
      }
    },
    {
      "actions": {
        "accept": false,
        "add_community": [
          "add_community5",
          "add_community6",
          "add_community7"
        ],
        "add_target_vrfs": [
          "add_target_vrfs1"
        ],
        "community": [
          "community4",
          "community5"
        ],
        "exclude_as_path": [
          "exclude_as_path0",
          "exclude_as_path1",
          "exclude_as_path2"
        ]
      },
      "matching": {
        "as_path": [
          "String3"
        ],
        "community": [
          "community4"
        ],
        "network": [
          "network7",
          "network8",
          "network9"
        ],
        "prefix": [
          "prefix5",
          "prefix6",
          "prefix7"
        ],
        "protocol": [
          "aggregate",
          "bgp"
        ]
      }
    }
  ]
}
```

