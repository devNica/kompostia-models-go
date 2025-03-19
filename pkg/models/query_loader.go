package models

import (
	"embed"
	"fmt"
	"strings"
)

//go:embed queries/*.sql
var queriesFS embed.FS

func LoadSQLQuery(queryKey string, params map[string]string) (string, error) {

	queryFile, exists := SQLQueries[queryKey]

	if !exists {
		return "", fmt.Errorf("consulta SQL no encontrada: %s", queryKey)
	}

	filePath := fmt.Sprintf("queries/%s.sql", queryFile)

	content, err := queriesFS.ReadFile(filePath)

	if err != nil {
		return "", fmt.Errorf("error al leer la consulta %s: %w", queryKey, err)
	}

	query := string(content)

	//Reemplazar los placeholders {{params}} con los valores proporcionados

	for key, value := range params {
		placeholder := fmt.Sprintf("{{%s}}", key)
		query = strings.ReplaceAll(query, placeholder, value)
	}

	return query, nil
}
