
# Config Switch Local Accounts User

*This model accepts additional fields of type interface{}.*

## Structure

`ConfigSwitchLocalAccountsUser`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Password` | `*string` | Optional | - |
| `Role` | [`*models.ConfigSwitchLocalAccountsUserRoleEnum`](../../doc/models/config-switch-local-accounts-user-role-enum.md) | Optional | enum: `admin`, `helpdesk`, `none`, `read`<br><br>**Default**: `"none"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "password": "Juniper123",
  "role": "none",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

