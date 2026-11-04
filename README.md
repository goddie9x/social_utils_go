# social_utils_go

Shared Go utilities for the Social Platform microservices. Published as the Go module **`github.com/goddie9x/social_utils_go`**.

## Install

```bash
go get github.com/goddie9x/social_utils_go@v1.0.0
```

```go
import (
    "github.com/goddie9x/social_utils_go/constants"
    "github.com/goddie9x/social_utils_go/middlewares"
)
```

## Packages

- `configs` — Eureka client config
- `constants` — shared enums (Role, PrivacyType, PostType)
- `dotenv` — env helper
- `eureka` — health status endpoint
- `exceptions` — common exceptions + gin error handler
- `middlewares` — auth middleware
- `requests` — pagination request
- `validates` — post validation
