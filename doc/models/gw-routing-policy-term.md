
# Gw Routing Policy Term

## Structure

`GwRoutingPolicyTerm`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Actions` | [`*models.GwRoutingPolicyTermAction`](../../doc/models/gw-routing-policy-term-action.md) | Optional | When used as import policy |
| `Matching` | [`*models.GwRoutingPolicyTermMatching`](../../doc/models/gw-routing-policy-term-matching.md) | Optional | zero or more criteria/filter can be specified to match the term, all criteria have to be met |

## Example (as JSON)

```json
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
```

