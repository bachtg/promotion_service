package utils

func CoalesceInt32(val, fallback int32) int32 {
	if val == 0 {
		return fallback
	}
	return val
}
