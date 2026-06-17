package main

import "encoding/json"

// encodeJSON marshals v with stable indentation and a trailing newline,
// matching the convention used by cmd/nemo's bundle writers. Keeping a
// single helper avoids minor formatting drift between stages.
func encodeJSON(v any) ([]byte, error) {
	body, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return nil, err
	}
	return append(body, '\n'), nil
}
