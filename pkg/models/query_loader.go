package models

import (
	"fmt"
	"os"
	"strings"
)

func LoadSQLQuery(queryName string, params map[string]string) (string, error) {

	filePath := fmt.Sprintf("queries/%s.sql", queryName)

	content, err := os.ReadFile(filePath)

	if err != nil {
		return "", fmt.Errorf("error al leer la consulta %s: %w", queryName, err)
	}

	query := string(content)

	//Reemplazar los placeholders {{params}} con los valores proporcionados

	for key, value := range params {
		placeholder := fmt.Sprintf("{{%s}}", key)
		query = strings.ReplaceAll(query, placeholder, value)
	}

	return query, nil
}
