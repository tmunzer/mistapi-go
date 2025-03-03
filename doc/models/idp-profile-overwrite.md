
# Idp Profile Overwrite

*This model accepts additional fields of type interface{}.*

## Structure

`IdpProfileOverwrite`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Action` | [`*models.IdpProfileActionEnum`](../../doc/models/idp-profile-action-enum.md) | Optional | enum:<br><br>* alert (default)<br>* drop: silently dropping packets<br>* close: notify client/server to close connection<br>**Default**: `"alert"` |
| `Matching` | [`*models.IdpProfileMatching`](../../doc/models/idp-profile-matching.md) | Optional | - |
| `Name` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "action": "alert",
  "matching": {
    "attack_name": [
      "attack_name9"
    ],
    "dst_subnet": [
      "dst_subnet5"
    ],
    "severity": [
      "critical",
      "info"
    ],
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "name": "name4",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

