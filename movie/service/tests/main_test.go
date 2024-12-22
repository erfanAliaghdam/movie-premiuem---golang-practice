package tests

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Set environment variables with mock values for testing
	os.Setenv("SECRET_KEY", "test-secret-key")
	os.Setenv("DB_NAME", "test-db")
	os.Setenv("REDIS_URL", "localhost:6379")
	os.Setenv("BUCKET_ACCESS_KEY", "test-access-key")
	os.Setenv("BUCKET_SECRET_KEY", "test-secret-key")
	os.Setenv("BUCKET_ENDPOINT", "https://test.com")
	os.Setenv("BUCKET_NAME", "test-bucket")

	// Run the tests
	exitCode := m.Run()

	// Clean up environment variables if needed
	os.Unsetenv("SECRET_KEY")
	os.Unsetenv("DB_NAME")
	os.Unsetenv("REDIS_URL")
	os.Unsetenv("BUCKET_ACCESS_KEY")
	os.Unsetenv("BUCKET_SECRET_KEY")
	os.Unsetenv("BUCKET_ENDPOINT")
	os.Unsetenv("BUCKET_NAME")

	// Exit with the test result code
	os.Exit(exitCode)
}
