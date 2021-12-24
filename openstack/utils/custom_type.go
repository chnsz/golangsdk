package utils

// In the function of fmt/log and json.Marshal, all SensitiveString fields will return "******".
type SensitiveString string

func (s SensitiveString) String() string {
	return "******"
}

func (s SensitiveString) MarshalJSON() ([]byte, error) {
	return []byte(`"******"`), nil
}
