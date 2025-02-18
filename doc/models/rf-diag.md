
# Rf Diag

RF Diag

*This model accepts additional fields of type interface{}.*

## Structure

`RfDiag`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Duration` | `*int` | Optional | recording length in seconds, max is 180. Default value is also 180.<br>**Default**: `180`<br>**Constraints**: `<= 180` |
| `Mac` | `*string` | Optional | If `type`==`client` or `asset`, mac of the device |
| `Name` | `string` | Required | Name of the recording, the name of the sdk client would be a good default choice |
| `SdkclientId` | `*uuid.UUID` | Optional | If `type`==`sdkclient`, sdkclient_id of this recording |
| `Type` | [`models.RfClientTypeEnum`](../../doc/models/rf-client-type-enum.md) | Required | enum: `asset`, `client`, `sdkclient` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "duration": 180,
  "name": "name6",
  "type": "sdkclient",
  "mac": "mac0",
  "sdkclient_id": "00000a0e-0000-0000-0000-000000000000",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

