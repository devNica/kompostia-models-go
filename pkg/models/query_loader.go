package models

import (
	"embed"
	"fmt"
	"strings"
)

//go:embed queries/*.sql
var queriesFS embed.FS

var SQLQueries = map[string]string{
	"registeredCategories":    "registered_categories.sql",
	"hierarchicalCategoryRel": "hierarchical_category_relationship.sql",
}

func LoadSQLQuery(queryKey string, params map[string]string) (string, error) {

	queryFile, exists := SQLQueries[queryKey]

	if !exists {
		return "", fmt.Errorf("consulta SQL no encontrada: %s", queryKey)
	}

	filePath := fmt.Sprintf("queries/%s", queryFile)

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

func ListEmbeddedQueries() {
	files, err := queriesFS.ReadDir("queries")
	if err != nil {
		fmt.Println("Error al listar archivos embebidos:", err)
		return
	}

	fmt.Println("Archivos embebidos en queries/:")
	for _, file := range files {
		fmt.Println("-", file.Name())
	}
}
