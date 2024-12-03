
# Remote Syslog Archive

*This model accepts additional fields of type interface{}.*

## Structure

`RemoteSyslogArchive`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Files` | `*int` | Optional | - |
| `Size` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "files": 20,
  "size": "5m",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

