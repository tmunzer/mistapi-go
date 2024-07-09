
# Synthetictest Radius Server

## Structure

`SynthetictestRadiusServer`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Password` | `string` | Required | Specify the password associated with the username |
| `Profile` | `*string` | Optional | Specify the access profile associated with the subscriber<br>**Default**: `"dot1x"` |
| `User` | `string` | Required | Specify the subscriber username to test |

## Example (as JSON)

```json
{
  "password": "password6",
  "profile": "dot1x",
  "user": "user2"
}
```

