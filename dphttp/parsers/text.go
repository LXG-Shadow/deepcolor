package parsers

import "github.com/aynakeya/deepcolor/dphttp"

func TextParser(resp *dphttp.Response) (string, error) {
	return resp.String(), nil
}
