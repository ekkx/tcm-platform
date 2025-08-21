package util

import "github.com/ekkx/tcm-platform/internal/platform/ulid"

func ToULIDStrings(ids []ulid.ULID) []string {
	strs := make([]string, len(ids))
	for i, id := range ids {
		strs[i] = id.String()
	}
	return strs
}
