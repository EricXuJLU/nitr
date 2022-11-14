package lib

// StringsDeleteNil 字符串数组，除去其中的空元素，并保持相对顺序不变
func StringsDeleteNil(src []string) (dst []string) {
	for _, v := range src {
		if len(v) != 0 {
			dst = append(dst, v)
		}
	}
	return dst
}
