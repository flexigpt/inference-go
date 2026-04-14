package debugclient

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

// GenerateCurlCommand builds a (mostly) copy-pasteable curl command from
// APIRequestDetails. It uses the already-redacted Data and Headers.
func GenerateCurlCommand(config *APIRequestDetails) string {
	if config == nil || config.URL == nil || config.Method == nil {
		return ""
	}

	var b strings.Builder

	method := strings.ToUpper(*config.Method)
	b.WriteString("curl")
	if method != "" {
		b.WriteString(" -X ")
		b.WriteString(method)
	}

	if config.URL != nil {
		escapedURL := shellQuote(*config.URL)
		b.WriteString(" ")
		b.WriteString(escapedURL)
	}

	// Headers (sorted for stability).
	if len(config.Headers) > 0 {
		keys := make([]string, 0, len(config.Headers))
		for k := range config.Headers {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			v := config.Headers[k]
			headerStr := fmt.Sprintf("%s: %v", k, v)
			b.WriteString(" \\\n  -H ")
			b.WriteString(shellQuote(headerStr))
		}
	}

	if config.Data != nil {
		bodyBytes, err := json.MarshalIndent(config.Data, "", "  ")
		if err == nil {
			b.WriteString(" \\\n  --data-raw ")
			b.WriteString(shellQuote(string(bodyBytes)))
		}
	}

	return b.String()
}
