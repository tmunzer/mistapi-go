
# Rf Diag

RF Diag

## Structure

`RfDiag`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Duration` | `*int` | Optional | recording length in seconds, max is 180. Default value is also 180.<br>**Default**: `180`<br>**Constraints**: `<= 180` |
| `Mac` | `*string` | Optional | if `type`==`client` or `asset`, mac of the device |
| `Name` | `string` | Required | name of the recording, the name of the sdk client would be a good default choice |
| `SdkclientId` | `*uuid.UUID` | Optional | if `type`==`sdkclient`, sdkclient_id of this recording |
| `Type` | [`models.RfClientTypeEnum`](../../doc/models/rf-client-type-enum.md) | Required | enum: `asset`, `client`, `sdkclient` |

## Example (as JSON)

```json
{
  "duration": 180,
  "name": "name6",
  "type": "client",
  "mac": "mac0",
  "sdkclient_id": "00001116-0000-0000-0000-000000000000"
}
```

