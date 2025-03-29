package util

import "strings"

func IsValidOrderField(field string) bool {
	validFields := map[string]bool{
		"created_at": true,
		"updated_at": true,
		"user_id":    true,
	}
	return validFields[strings.ToLower(field)]
}
