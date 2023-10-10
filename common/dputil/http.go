package dputil

func HttpHeaderMerge(src map[string]string, updated map[string]string) map[string]string {
	header := make(map[string]string)
	for k, v := range src {
		header[k] = v
	}
	for k, v := range updated {
		header[k] = v
	}
	return header
}
