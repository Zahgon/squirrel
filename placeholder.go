package squirrel

// PlaceholderFormat is the interface that wraps the ReplacePlaceholders method.
//
// ReplacePlaceholders takes a SQL statement and replaces each question mark
// placeholder with a (possibly different) SQL placeholder.
type PlaceholderFormat interface {
	ReplacePlaceholders(sql string) (string, error)
}

type placeholderDebugger interface {
	debugPlaceholder() string
}

var (
	// Question is a PlaceholderFormat instance that leaves placeholders as
	// question marks.
	Question = questionFormat{}

	// Dollar is a PlaceholderFormat instance that replaces placeholders with
	// dollar-prefixed positional placeholders (e.g. $1, $2, $3).
	Dollar = dollarFormat{}

	// Colon is a PlaceholderFormat instance that replaces placeholders with
	// colon-prefixed positional placeholders (e.g. :1, :2, :3).
	Colon = colonFormat{}

	// AtP is a PlaceholderFormat instance that replaces placeholders with
	// "@p"-prefixed positional placeholders (e.g. @p1, @p2, @p3).
	AtP = atpFormat{}
)

type questionFormat struct{}

func (questionFormat) ReplacePlaceholders(sql string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (questionFormat) debugPlaceholder() string { _ = "STUB: not implemented"; return "" }

type dollarFormat struct{}

func (dollarFormat) ReplacePlaceholders(sql string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (dollarFormat) debugPlaceholder() string { _ = "STUB: not implemented"; return "" }

type colonFormat struct{}

func (colonFormat) ReplacePlaceholders(sql string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (colonFormat) debugPlaceholder() string { _ = "STUB: not implemented"; return "" }

type atpFormat struct{}

func (atpFormat) ReplacePlaceholders(sql string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (atpFormat) debugPlaceholder() string {
	_ = "STUB: not implemented"

	// Placeholders returns a string with count ? placeholders joined with commas.
	return ""
}

func Placeholders(count int) string { _ = "STUB: not implemented"; return "" }

func replacePositionalPlaceholders(sql, prefix string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// escape ?? => ?
