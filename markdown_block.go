package json_markd

type markdownBlock struct {
	TabCount int
	Key      string
	Value    DataType
}

func newMarkdownBlock(tabCount int, key string, value DataType) markdownBlock {
	return markdownBlock{
		TabCount: tabCount,
		Key:      key,
		Value:    value,
	}
}

func getDatatypeFromVal(value string) DataType {
	switch value {
	case "object":
		return Object
	case "string":
		return String
	case "array":
		return Array
	case "integer":
		return Integer
	case "double":
		return Double
	}
	return Invalid
}

func (client markdownBlock) GetSuffixForDatatype() string {
	switch client.Value {
	case Object:
		return "}"
	case String:
		return "\"random string\""
	case Integer:
		return "0"
	case Double:
		return "0.0"
	case Array:
		return "]"
	}
	return "}"
}

func (client markdownBlock) GetPrefixForDatatype() string {
	switch client.Value {
	case Object:
		return client.Key + " : {\n"
	case String:
		return client.Key + " : "
	case Integer:
		return client.Key + " : "
	case Double:
		return client.Key + " : "
	case Array:
		return client.Key + " : [\n"
	}
	return "{\n"
}

func (client markdownBlock) GetPrefixForDatatypeWhenParentIsArray() string {
	switch client.Value {
	case Object:
		return "{\n"
	case String:
		return ""
	case Integer:
		return ""
	case Double:
		return ""
	case Array:
		return "[\n"
	}
	return "{\n"
}
