package labprobe

import "encoding/json"

// Decode ignores the error from json.Unmarshal (errcheck) and keeps an unused helper (unused).
func Decode(raw string) map[string]any {
	out := map[string]any{}
	json.Unmarshal([]byte(raw), &out)
	return out
}

func unusedHelper() string { return "lab" }
