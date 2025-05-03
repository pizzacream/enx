<h1 align="center">ENX</h1>
<p align="center">A simple library to parse environment variables with default value</p>

###### Installation

```bash
go get github.com/pizzacream/enx
```

###### Getting started

```go
import (
	"github.com/pizzacream/enx"
	_ "github.com/pizzacream/enx/autoload"
)

type MyStruct struct {
	Value string
}

var port = enx.MustGetInt("PORT")
var maxConnections = enx.GetInt("MAX_CONNECTIONS", 10)
var myStruct = enx.MustGetStruct[MyStruct]("MY_STRUCT")
var duration = enx.GetDuration("DURATION", time.Second*10)
var regex = enx.MustGetRegex("Regex")

```

## Related projects

- [godotenv](https://github.com/joho/godotenv) - Go port of Ruby's dotenv library (Loads environment variables from .env).
- [genv](https://github.com/sakirsensoy/genv) - Read environment variables easily with dotenv support.
