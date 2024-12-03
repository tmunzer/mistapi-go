
# Org Crl Import File

*This model accepts additional fields of type interface{}.*

## Structure

`OrgCrlImportFile`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `File` | `*[]byte` | Optional | a PEM or DER formatted CRL file |
| `Json` | `*string` | Optional | a JSON string with "name" field for CRL file issuer (optional) |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "file": "data:text/plain;name=dummy_file;base64,",
  "json": "json2",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

