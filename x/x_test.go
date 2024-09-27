package x

import (
	"testing"
)

func TestElem_Render(t *testing.T) {
	tests := []struct {
		name     string
		elem     Elem
		expected string
	}{
		{
			name:     "Div with text",
			elem:     Div(X{}, C("Hello, World!")),
			expected: "<div>Hello, World!</div>",
		},
		{
			name:     "Span with Class",
			elem:     Span(X{Class: "highlight"}, C("Highlighted text")),
			expected: "<span class=\"highlight\">Highlighted text</span>",
		},
		{
			name:     "Anchor with href",
			elem:     A(X{Att: `href="https://example.com"`}, C("Example")),
			expected: "<a href=\"https://example.com\">Example</a>",
		},
		{
			name:     "Image with Attributes (self-closing)",
			elem:     Img(X{Att: `src="image.png" alt="An image"`}),
			expected: "<img src=\"image.png\" alt=\"An image\" />",
		},
		{
			name:     "Nested elements",
			elem:     Div(X{Class: "container"}, Span(X{}, C("Nested span"))),
			expected: "<div class=\"container\"><span>Nested span</span></div>",
		},
		{
			name:     "DOCTYPE",
			elem:     DOCTYPE(),
			expected: "<!DOCTYPE html>",
		},
		{
			name:     "Html with head and body",
			elem:     Html(X{}, Head(X{}, Title(X{}, C("Page Title"))), Body(X{}, P(X{}, C("Hello, World!")))),
			expected: "<html><head><title>Page Title</title></head><body><p>Hello, World!</p></body></html>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := string(tt.elem.Render())
			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestIF(t *testing.T) {
	tests := []struct {
		condition bool
		trueCase  Elem
		expected  string
	}{
		{
			condition: true,
			trueCase:  Div(X{}, C("Condition is true")),
			expected:  "<div>Condition is true</div>",
		},
		{
			condition: false,
			trueCase:  Div(X{}, C("Condition is true")),
			expected:  "",
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := string(IF(tt.condition, tt.trueCase).Render())
			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestTER(t *testing.T) {
	tests := []struct {
		condition bool
		trueCase  Elem
		falseCase Elem
		expected  string
	}{
		{
			condition: true,
			trueCase:  Div(X{}, C("True case")),
			falseCase: Div(X{}, C("False case")),
			expected:  "<div>True case</div>",
		},
		{
			condition: false,
			trueCase:  Div(X{}, C("True case")),
			falseCase: Div(X{}, C("False case")),
			expected:  "<div>False case</div>",
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := string(TER(tt.condition, tt.trueCase, tt.falseCase).Render())
			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestFOR(t *testing.T) {
	tests := []struct {
		elements []Elem
		expected []string
	}{
		{
			elements: []Elem{
				Div(X{}, C("Item 1")),
				Div(X{}, C("Item 2")),
			},
			expected: []string{
				"<div>Item 1</div>",
				"<div>Item 2</div>",
			},
		},
	}

	for i, tt := range tests {
		t.Run("", func(t *testing.T) {
			results := FOR(tt.elements)
			for j, result := range results {
				if string(result.Render()) != tt.expected[j] {
					t.Errorf("test %d, expected %q, got %q", i, tt.expected[j], result.Render())
				}
			}
		})
	}
}

func TestSTERSIF(t *testing.T) {
	tests := []struct {
		condition bool
		trueCase  string
		falseCase string
		expected  string
	}{
		{
			condition: true,
			trueCase:  "True",
			falseCase: "False",
			expected:  "True",
		},
		{
			condition: false,
			trueCase:  "True",
			falseCase: "False",
			expected:  "False",
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := STER(tt.condition, tt.trueCase, tt.falseCase)
			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestSIF(t *testing.T) {
	tests := []struct {
		condition bool
		trueCase  string
		expected  string
	}{
		{
			condition: true,
			trueCase:  "True",
			expected:  "True",
		},
		{
			condition: false,
			trueCase:  "True",
			expected:  "",
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := SIF(tt.condition, tt.trueCase)
			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
		})
	}
}