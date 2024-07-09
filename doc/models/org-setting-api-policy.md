
# Org Setting Api Policy

## Structure

`OrgSettingApiPolicy`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `NoReveal` | `*bool` | Optional | by default, API hides password/secrets when the user doesn't have write access<br><br>* `true`: API will hide passwords/secrets for all users<br>* `false`: API will hide passwords/secrests for read-only users<br>**Default**: `false` |

## Example (as JSON)

```json
{
  "no_reveal": false
}
```

