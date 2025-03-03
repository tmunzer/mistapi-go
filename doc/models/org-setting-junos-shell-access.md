
# Org Setting Junos Shell Access

by default, webshell access is only enabled for Admin user

*This model accepts additional fields of type interface{}.*

## Structure

`OrgSettingJunosShellAccess`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Admin` | [`*models.OrgSettingJunosShellAccessAdminEnum`](../../doc/models/org-setting-junos-shell-access-admin-enum.md) | Optional | enum: `admin`, `viewer`, `none`<br>**Default**: `"admin"` |
| `Helpdesk` | [`*models.OrgSettingJunosShellAccessHelpdeskEnum`](../../doc/models/org-setting-junos-shell-access-helpdesk-enum.md) | Optional | enum: `admin`, `viewer`, `none`<br>**Default**: `"none"` |
| `Read` | [`*models.OrgSettingJunosShellAccessReadEnum`](../../doc/models/org-setting-junos-shell-access-read-enum.md) | Optional | enum: `admin`, `viewer`, `none`<br>**Default**: `"none"` |
| `Write` | [`*models.OrgSettingJunosShellAccessWriteEnum`](../../doc/models/org-setting-junos-shell-access-write-enum.md) | Optional | enum: `admin`, `viewer`, `none`<br>**Default**: `"admin"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "admin": "admin",
  "helpdesk": "none",
  "read": "none",
  "write": "admin",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

