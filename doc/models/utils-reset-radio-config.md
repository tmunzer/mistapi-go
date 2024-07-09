
# Utils Reset Radio Config

## Structure

`UtilsResetRadioConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Bands` | `[]string` | Required | list of bands |
| `Force` | `*bool` | Optional | whether to reset those with radio disabled. default is false (i.e. if user intentionally disables a radio, honor it)<br>**Default**: `false` |

## Example (as JSON)

```json
{
  "bands": [
    "bands4"
  ],
  "force": false
}
```

