
# Idp Profile Overwrite

## Structure

`IdpProfileOverwrite`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Action` | [`*models.IdpProfileActionEnum`](../../doc/models/idp-profile-action-enum.md) | Optional | enum:<br><br>* alert (default)<br>* drop: siliently dropping packets<br>* close: notify client/server to close connection<br>**Default**: `"alert"` |
| `Matching` | [`*models.IdpProfileMatching`](../../doc/models/idp-profile-matching.md) | Optional | - |
| `Name` | `*string` | Optional | - |

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
    ]
  },
  "name": "name8"
}
```

