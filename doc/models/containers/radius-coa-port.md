
# Radius Coa Port

RADIUS CoA Port, value from 1 to 65535, default is 3799

## Class Name

`RadiusCoaPort`

## Cases

| Type | Factory Method |
|  --- | --- |
| `int` | models.RadiusCoaPortContainer.FromNumber(int number) |
| `string` | models.RadiusCoaPortContainer.FromString(string mString) |

## int

### Initialization Code

#### Example

```go
value := models.RadiusCoaPortContainer.FromNumber(1)
```

## string

### Initialization Code

#### Example

```go
value := models.RadiusCoaPortContainer.FromString("String0")
```

