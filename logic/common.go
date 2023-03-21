package logic

func contains(s string, list []string) bool {
	for _, v := range list {
		if s == v {
			return true
		}
	}

	return false
}
