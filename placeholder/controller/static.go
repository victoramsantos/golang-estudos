package controller

import (
	"encoding/json"
	"placeholder/service"
)

// GET /poem/static
func GetStatic() string {
	var content = service.GetStaticContent()

	response := Response{
		Status: 200,
		Body:   content.Content,
	}

	jsonResponse, _ := json.MarshalIndent(response, "", " ")

	return string(jsonResponse)
}

// GET /poem/dynamic
func GetDynamic() string {
	var content = service.GetDynamicContent()

	response := Response{
		Status: 200,
		Body:   content.Content,
	}

	jsonResponse, _ := json.MarshalIndent(response, "", " ")

	return string(jsonResponse)
}
