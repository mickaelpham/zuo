package print

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"text/tabwriter"
)

func Table(records []map[string]string) {
	headers := headers(records)

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)
	fmt.Fprintln(w, strings.Join(headers, "\t "))
	for _, record := range records {
		values := make([]string, 0, len(record))

		for _, field := range headers {
			values = append(values, record[field])
		}

		fmt.Fprintln(w, strings.Join(values, "\t "))
	}
	w.Flush()
}

func headers(records []map[string]string) []string {
	headers := fieldNames(records[0])

	// Loop through all the records to check if we are missing a field
	for _, record := range records {
		if len(record) > len(headers) {
			headers = fieldNames(record)
		}
	}

	sort.Strings(headers)
	return headers
}

func fieldNames(record map[string]string) []string {
	fields := make([]string, 0, len(record))

	for field := range record {
		fields = append(fields, field)
	}

	return fields
}
