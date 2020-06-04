package env

import "os"

// GetEnv получение значения из окружения, если значение отсутствует, то возвращает пустую строку
func GetEnv(key string) string {
	value, _ := os.LookupEnv(key)
	return value
}
