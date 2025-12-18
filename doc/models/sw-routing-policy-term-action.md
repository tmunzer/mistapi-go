
# Sw Routing Policy Term Action

When used as import policy

## Structure

`SwRoutingPolicyTermAction`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Accept` | `*bool` | Optional | - |
| `Community` | `[]string` | Optional | When used as export policy, optional |
| `LocalPreference` | [`*models.RoutingPolicyLocalPreference`](../../doc/models/containers/routing-policy-local-preference.md) | Optional | Optional, for an import policy, local_preference can be changed, value in range 1-4294967294. Can be a Variable (e.g. `{{bgp_as}}`) |
| `PrependAsPath` | `[]string` | Optional | When used as export policy, optional. By default, the local AS will be prepended, to change it. Can be a Variable (e.g. `{{as_path}}`) |

## Example (as JSON)

```json
{
  "accept": false,
  "community": [
    "community0",
    "community1"
  ],
  "local_preference": "String1",
  "prepend_as_path": [
    "prepend_as_path5",
    "prepend_as_path4",
    "prepend_as_path3"
  ]
}
```

