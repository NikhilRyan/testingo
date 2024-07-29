package main

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jarcoal/httpmock"
	"github.com/nikhilryan/testingo/pkg/assertions"
	"github.com/nikhilryan/testingo/pkg/benchmarks"
	"github.com/nikhilryan/testingo/pkg/config"
	"github.com/nikhilryan/testingo/pkg/mocks"
	"github.com/nikhilryan/testingo/pkg/runner"
	"github.com/redis/go-redis/v9"
	"net/http"
	"testing"
	"time"
)

// GetRedisValue Example function using Redis
func GetRedisValue(client *redis.Client, key string) (string, error) {
	return client.Get(context.Background(), key).Result()
}

// TestGetRedisValue Test case for Redis mock
func TestGetRedisValue(t *testing.T) {
	redisMock := mocks.NewRedisMock()
	defer redisMock.Mock.ClearExpect()

	redisMock.Mock.ExpectGet("key").SetVal("value")

	result, err := GetRedisValue(redisMock.Client, "key")
	assertions.AssertNil(t, err)
	assertions.AssertEqual(t, "value", result)
}

// GetUserByID Example function using SQL DB
func GetUserByID(db *sql.DB, id int) (string, error) {
	var name string
	err := db.QueryRow("SELECT name FROM users WHERE id = ?", id).Scan(&name)
	return name, err
}

// TestGetUserByID Test case for SQL DB mock
func TestGetUserByID(t *testing.T) {
	dbMock, err := mocks.NewDBMock()
	assertions.AssertNil(t, err)
	defer dbMock.DB.Close()

	dbMock.Mock.ExpectQuery("SELECT name FROM users WHERE id = ?").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("John Doe"))

	result, err := GetUserByID(dbMock.DB, 1)
	assertions.AssertNil(t, err)
	assertions.AssertEqual(t, "John Doe", result)
}

// GetAPIResponse Example function using API
func GetAPIResponse(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	return resp.StatusCode, nil
}

// TestGetAPIResponse Test case for API mock
func TestGetAPIResponse(t *testing.T) {
	apiMock := mocks.NewAPIMock()
	defer apiMock.DeactivateAndReset()

	apiMock.RegisterResponder("GET", "http://example.com",
		httpmock.NewStringResponder(200, "OK"))

	status, err := GetAPIResponse("http://example.com")
	assertions.AssertNil(t, err)
	assertions.AssertEqual(t, 200, status)
}

// BenchmarkAdd Benchmark example with memory allocation
func BenchmarkAdd(b *testing.B) {
	benchmarks.BenchmarkMemoryAlloc(b, func() {
		// Function to benchmark
	})
}

// Main function to run all tests and benchmarks
func main() {
	t := &testing.T{}
	b := &testing.B{}

	// Configure settings
	config.SetConfig(config.Config{
		Timeout: 60 * time.Second,
		Verbose: true,
	})

	// Run tests in parallel
	runner.RunTestsParallel(t, []func(t *testing.T){
		TestGetRedisValue,
		TestGetUserByID,
		TestGetAPIResponse,
	})

	// Run benchmarks
	runner.RunBenchmarks(b, []func(b *testing.B){
		BenchmarkAdd,
	})
}
