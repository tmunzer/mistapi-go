
# Config Switch Local Accounts User

## Structure

`ConfigSwitchLocalAccountsUser`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Password` | `*string` | Optional | - |
| `Role` | [`*models.ConfigSwitchLocalAccountsUserRoleEnum`](../../doc/models/config-switch-local-accounts-user-role-enum.md) | Optional | enum: `admin`, `helpdesk`, `none`, `read`<br><br>**Default**: `"none"` |

## Example (as JSON)

```json
{
  "password": "Juniper123",
  "role": "none"
}
```

