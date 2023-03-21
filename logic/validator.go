package logic

import "gear-notification/domain"

func ValidateBrand(selection []string) (string, bool) {
	brandNameList := domain.GetBrandNameList()

	for _, s := range selection {
		if !contains(s, brandNameList) {
			return s, false
		}

	}

	return "", true
}
