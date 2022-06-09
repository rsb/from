# From
Casting from one type to another in golang.

The primary reason for this package is the need to cast string values into
other types. This dynamic casting is required by my [conf package](https://github.com/rsb/conf)
as well as [pflags](https://github.com/rsb/pflags) and [fuelcell](https://github.com/rsb/fuelcell)
which often take string values from `environment variables` or `cli` and need to convert them to 
`go` types usually to be assigned to some configuration struct.
