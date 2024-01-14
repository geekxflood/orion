// internal/httpclient/insecure.go

package httpclient

import "net/http"

// SetInsecure sets the HTTP client to allow insecure connections.
// This function modifies the default transport to skip TLS certificate verification.
// Use with caution as it may expose your application to security risks.
func SetInsecure() {
	tr := http.DefaultTransport.(*http.Transport)
	tr.TLSClientConfig.InsecureSkipVerify = true
}
