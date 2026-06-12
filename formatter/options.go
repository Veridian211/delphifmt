package formatter

type KeywordCase string

const (
	KeywordLowercase KeywordCase = "lowercase"
	KeywordUppercase KeywordCase = "uppercase"
)

type BeginPosition string

const (
	BeginEndOfLine BeginPosition = "end_of_line"
	BeginNewLine   BeginPosition = "new_line"
)

type UsesStyle string

const (
	UsesOnePerLine UsesStyle = "one_per_line"
	UsesSingleLine UsesStyle = "single_line"
)

type IndentChar string

const (
	IndentSpace IndentChar = "space"
	IndentTab   IndentChar = "tab"
)

type EndOfLine string

const (
	EOLlf   EndOfLine = "lf"
	EOLcrlf EndOfLine = "crlf"
)

type Options struct {
	Indent                   int           `yaml:"indent"`
	IndentChar               IndentChar    `yaml:"indent_char"`
	LineLength               int           `yaml:"line_length"`
	KeywordCase              KeywordCase   `yaml:"keyword_case"`
	BeginPosition            BeginPosition `yaml:"begin_position"`
	BlankLinesBetweenMethods int           `yaml:"blank_lines_between_methods"`
	UsesStyle                UsesStyle     `yaml:"uses_style"`
	EndOfLine                EndOfLine     `yaml:"end_of_line"`
}

func DefaultOptions() Options {
	return Options{
		Indent:                   4,
		IndentChar:               IndentSpace,
		LineLength:               100,
		KeywordCase:              KeywordLowercase,
		BeginPosition:            BeginEndOfLine,
		BlankLinesBetweenMethods: 1,
		UsesStyle:                UsesOnePerLine,
		EndOfLine:                EOLlf,
	}
}
