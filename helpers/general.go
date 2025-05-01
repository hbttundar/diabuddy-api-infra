package helpers

// ToPointer returns a pointer to any value, useful for defaulting.
func ToPointer[T any](v T) *T {
	return &v
}

// IfNotEmpty returns fallback if the input string is empty.
func IfNotEmpty(s, fallback string) string {
	if s == "" {
		return fallback
	}
	return s
}

// Coalesce returns the first non-empty string from the arguments.
func Coalesce(values ...string) string {
	for _, v := range values {
		if v != "" {
			return v
		}
	}
	return ""
}
