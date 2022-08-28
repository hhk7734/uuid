package uuid

func (uuid UUID) GormDataType() string {
	return "binary(16)"
}
