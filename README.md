
# Testingo

Testingo is a standalone and easily integrable Go testing library designed to make writing test cases simple and advanced. It includes features for assertions, mocking, benchmarking, parallel test execution, and more.

## Installation

```sh
go get github.com/nikhilryan/testingo
```

## Usage

### Assertions
```go
import "github.com/nikhilryan/testingo/pkg/assertions"

func TestSomething(t *testing.T) {
    assertions.AssertEqual(t, expected, actual)
    assertions.AssertNotEqual(t, expected, actual)
    assertions.AssertNil(t, actual)
    assertions.AssertNotNil(t, actual)
    assertions.AssertTrue(t, condition)
    assertions.AssertFalse(t, condition)
    assertions.AssertContains(t, container, item)
}
```

### Mocks

#### Redis Mocking
```go
import (
    "context"
    "github.com/redis/go-redis/v9"
    "github.com/nikhilryan/testingo/pkg/mocks"
)

func TestRedis(t *testing.T) {
    redisMock := mocks.NewRedisMock()
    defer redisMock.Mock.ClearExpect()

    redisMock.Mock.ExpectGet("key").SetVal("value")

    result, err := redisMock.Client.Get(redisMock.Context, "key").Result()
    assertions.AssertNil(t, err)
    assertions.AssertEqual(t, "value", result)
}
```

#### Database Mocking
```go
import (
    "database/sql"
    "github.com/DATA-DOG/go-sqlmock"
    "github.com/nikhilryan/testingo/pkg/mocks"
)

func TestDatabase(t *testing.T) {
    dbMock, err := mocks.NewDBMock()
    assertions.AssertNil(t, err)
    defer dbMock.DB.Close()

    dbMock.Mock.ExpectQuery("SELECT name FROM users WHERE id = ?").
        WithArgs(1).
        WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("John Doe"))

    var name string
    err = dbMock.DB.QueryRow("SELECT name FROM users WHERE id = ?", 1).Scan(&name)
    assertions.AssertNil(t, err)
    assertions.AssertEqual(t, "John Doe", name)
}
```

#### API Mocking
```go
import (
    "net/http"
    "github.com/jarcoal/httpmock"
    "github.com/nikhilryan/testingo/pkg/mocks"
)

func TestAPI(t *testing.T) {
    apiMock := mocks.NewAPIMock()
    defer apiMock.DeactivateAndReset()

    apiMock.RegisterResponder("GET", "http://example.com",
        httpmock.NewStringResponder(200, "OK"))

    resp, err := http.Get("http://example.com")
    assertions.AssertNil(t, err)
    assertions.AssertEqual(t, 200, resp.StatusCode)
}
```

### Running Tests in Parallel
```go
import "github.com/nikhilryan/testingo/pkg/runner"

func main() {
    t := &testing.T{}
    runner.RunTestsParallel(t, []func(t *testing.T){
        TestRedis,
        TestDatabase,
        TestAPI,
    })
}
```

### Configuration and Settings
```go
import "github.com/nikhilryan/testingo/pkg/config"

func main() {
    config.SetConfig(config.Config{
        Timeout: 60 * time.Second,
        Verbose: true,
    })
}
```

### License
[MIT](LICENSE)
