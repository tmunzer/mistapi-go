
# Org Setting Switch Mgmt

*This model accepts additional fields of type interface{}.*

## Structure

`OrgSettingSwitchMgmt`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApAffinityThreshold` | `*int` | Optional | If the field is set in both site/setting and org/setting, the value from site/setting will be used.<br><br>**Default**: `12` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "ap_affinity_threshold": 10,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

