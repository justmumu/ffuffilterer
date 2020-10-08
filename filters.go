package main

import "github.com/ffuf/ffuf/pkg/output"

type Filter func(record output.Result) bool

func GetFilters(opts Options) []Filter {
	filters := make([]Filter, 0)

	filters = append(filters, NewHostFilter(opts.Host),
		NewNotStatusCodeFilter(opts.NotCode),
		NewNotLengthFilter(opts.NotLengths),
		NewNotLineFilter(opts.NotLines),
		NewNotWordsFilter(opts.NotWords),
		NewStatusCodeFilter(opts.ShouldCode),
		NewLengthFilter(opts.NotLengths),
		NewLineFilter(opts.NotLines),
		NewWordsFilter(opts.NotWords))

	return filters
}

func NewHostFilter(hosts []string) Filter {
	return func(record output.Result) bool {
		if len(hosts) == 0 {
			return true
		}

		found := false
		for _, host := range hosts {
			if record.Host == host {
				found = true
			}
		}
		if found {
			return true
		}
		return false
	}
}

func NewStatusCodeFilter(statusCodes []int64) Filter {
	return func(record output.Result) bool {
		if len(statusCodes) == 0 {
			return true
		}

		found := false
		for _, sc := range statusCodes {
			if record.StatusCode == sc {
				found = true
			}
		}
		if found {
			return true
		}
		return false
	}
}

func NewWordsFilter(wordCounts []int64) Filter {
	return func(record output.Result) bool {
		if len(wordCounts) == 0 {
			return true
		}

		found := false
		for _, wc := range wordCounts {
			if record.ContentWords == wc {
				found = true
			}
		}
		if found {
			return true
		}
		return false
	}
}

func NewLengthFilter(lengthCounts []int64) Filter {
	return func(record output.Result) bool {
		if len(lengthCounts) == 0 {
			return true
		}

		found := false
		for _, lc := range lengthCounts {
			if record.ContentLength == lc {
				found = true
			}
		}
		if found {
			return true
		}
		return false
	}
}

func NewLineFilter(lineCounts []int64) Filter {
	return func(record output.Result) bool {
		if len(lineCounts) == 0 {
			return true
		}

		found := false
		for _, lc := range lineCounts {
			if record.ContentLines == lc {
				found = true
			}
		}
		if found {
			return true
		}
		return false
	}
}

func NewNotStatusCodeFilter(statusCodes []int64) Filter {
	return func(record output.Result) bool {
		if len(statusCodes) == 0 {
			return true
		}

		found := false
		for _, sc := range statusCodes {
			if record.StatusCode == sc {
				found = true
			}
		}

		if !found {
			return true
		}
		return false
	}
}

func NewNotWordsFilter(wordCounts []int64) Filter {
	return func(record output.Result) bool {
		if len(wordCounts) == 0 {
			return true
		}

		found := false
		for _, wc := range wordCounts {
			if record.ContentWords == wc {
				found = true
			}
		}
		if !found {
			return true
		}
		return false
	}
}

func NewNotLengthFilter(lengthCounts []int64) Filter {
	return func(record output.Result) bool {
		if len(lengthCounts) == 0 {
			return true
		}

		found := false
		for _, lc := range lengthCounts {
			if record.ContentLength == lc {
				found = true
			}
		}
		if !found {
			return true
		}
		return false
	}
}

func NewNotLineFilter(lineCounts []int64) Filter {
	return func(record output.Result) bool {
		if len(lineCounts) == 0 {
			return true
		}

		found := false
		for _, wc := range lineCounts {
			if record.ContentLines == wc {
				found = true
			}
		}
		if !found {
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
