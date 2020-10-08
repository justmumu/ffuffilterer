package main

import "github.com/ffuf/ffuf/pkg/output"

type Filter func(record output.Result) bool

func GetFilters(opts Options) []Filter {
	filters := make([]Filter, 0)

	for _, host := range opts.Host {
		filters = append(filters, NewHostFilter(host))
	}

	for _, wc := range opts.NotCode {
		filters = append(filters, NewNotStatusCodeFilter(wc))
	}

	for _, ws := range opts.NotLengths {
		filters = append(filters, NewNotLengthFilter(ws))
	}

	for _, wl := range opts.NotLines {
		filters = append(filters, NewNotLineFilter(wl))
	}

	for _, ww := range opts.NotWords {
		filters = append(filters, NewNotWordsFilter(ww))
	}

	for _, fc := range opts.ShouldCode {
		filters = append(filters, NewStatusCodeFilter(fc))
	}

	for _, fs := range opts.NotLengths {
		filters = append(filters, NewLengthFilter(fs))
	}

	for _, fl := range opts.NotLines {
		filters = append(filters, NewLineFilter(fl))
	}

	for _, fw := range opts.NotWords {
		filters = append(filters, NewWordsFilter(fw))
	}

	return filters
}

func NewHostFilter(host string) Filter {
	return func(record output.Result) bool {
		if record.Host == host {
			return true
		}
		return false
	}
}

func NewStatusCodeFilter(statusCode int64) Filter {
	return func(record output.Result) bool {
		if record.StatusCode == statusCode {
			return true
		}
		return false
	}
}

func NewWordsFilter(wordCount int64) Filter {
	return func(record output.Result) bool {
		if record.ContentWords == wordCount {
			return true
		}
		return false
	}
}

func NewLengthFilter(lengthCount int64) Filter {
	return func(record output.Result) bool {
		if record.ContentLength == lengthCount {
			return true
		}
		return false
	}
}

func NewLineFilter(lineCount int64) Filter {
	return func(record output.Result) bool {
		if record.ContentLines == lineCount {
			return true
		}
		return false
	}
}

func NewNotStatusCodeFilter(statusCode int64) Filter {
	return func(record output.Result) bool {
		if record.StatusCode != statusCode {
			return true
		}
		return false
	}
}

func NewNotWordsFilter(wordCount int64) Filter {
	return func(record output.Result) bool {
		if record.ContentWords != wordCount {
			return true
		}
		return false
	}
}

func NewNotLengthFilter(lengthCount int64) Filter {
	return func(record output.Result) bool {
		if record.ContentLength != lengthCount {
			return true
		}
		return false
	}
}

func NewNotLineFilter(lineCount int64) Filter {
	return func(record output.Result) bool {
		if record.ContentLines != lineCount {
			return true
		}
		return false
	}
}

func ApplyFilters(slice []output.Result, filters ...Filter) []output.Result {
	if len(filters) == 0 {
		return slice
	}

	filteredRecords := make([]output.Result, 0, len(slice))

	for _, s := range slice {
		keep := true

		for _, f := range filters {
			if !f(s) {
				keep = false
				break
			}
		}

		if keep {
			filteredRecords = append(filteredRecords, s)
		}
	}

	return filteredRecords
}
