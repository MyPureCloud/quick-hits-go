// >> START sdk-overview This example demonstrates authorizing using client credentials and getting a user
import (
	"github.com/MyPureCloud/platform-client-sdk-go/platformclientv2"
)

// Get default config to set config options
config := platformclientv2.GetDefaultConfiguration()

// Create API instance using default config
usersAPI := platformclientv2.NewUsersApi()

// Authorize SDK
err := config.AuthorizeClientCredentials(os.Getenv("GENESYS_CLOUD_CLIENT_ID"), os.Getenv("GENESYS_CLOUD_CLIENT_SECRET"))
if err != nil {
    panic(err)
}

// Make requests
user, response, err := usersAPI.GetUser("asdf1234-5678-90ab-cde1-123456789012", make([]string, 0), "")
fmt.Printf("Response:\n  Success: %v\n  Status code: %v\n  Correlation ID: %v\n", response.IsSuccess, response.StatusCode, response.CorrelationID)
if err != nil {
    fmt.Printf("Error calling GetUser: %v\n", err)
} else {
    fmt.Printf("Hello, %v\n", *user.Name)
}
// >> END sdk-overview
