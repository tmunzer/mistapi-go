
# Routing Policy

*This model accepts additional fields of type interface{}.*

## Structure

`RoutingPolicy`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Terms` | [`[]models.RoutingPolicyTerm`](../../doc/models/routing-policy-term.md) | Optional | zero or more criteria/filter can be specified to match the term, all criteria have to be met<br><br>**Constraints**: *Unique Items Required* |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

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
        ],
        "exampleAdditionalProperty": {
          "key1": "val1",
          "key2": "val2"
        }
      },
      "matching": {
        "as_path": [
          "as_path2"
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
          "protocol5",
          "protocol6"
        ],
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
        ],
        "exampleAdditionalProperty": {
          "key1": "val1",
          "key2": "val2"
        }
      },
      "matching": {
        "as_path": [
          "as_path2"
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
          "protocol5",
          "protocol6"
        ],
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

