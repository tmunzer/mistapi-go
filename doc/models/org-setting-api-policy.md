
# Org Setting Api Policy

*This model accepts additional fields of type interface{}.*

## Structure

`OrgSettingApiPolicy`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `NoReveal` | `*bool` | Optional | By default, API hides password/secrets when the user doesn't have write access<br><br>* `true`: API will hide passwords/secrets for all users<br>* `false`: API will hide passwords/secrets for read-only users<br>**Default**: `false` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "no_reveal": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

