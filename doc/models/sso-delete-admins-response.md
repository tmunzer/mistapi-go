
# Sso Delete Admins Response

## Structure

`SsoDeleteAdminsResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Deleted` | `[]string` | Optional | List of email addresses that were successfully deleted |
| `Errors` | `[]string` | Optional | List of error messages for emails that could not be deleted |

## Example (as JSON)

```json
{
  "deleted": [
    "deleted0",
    "deleted1",
    "deleted2"
  ],
  "errors": [
    "errors3"
  ]
}
```

