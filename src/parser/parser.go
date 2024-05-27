package parser

func ParseProj(root string, parseStruct bool) (map[string]string, error) {
	m := make(map[string]string)
	if parseStruct {
		str, err := getStructTreeString(root)
		if err != nil {
			return nil, err
		}
		m["Project Struct Tree"] = str
	}
	return m, nil
}
