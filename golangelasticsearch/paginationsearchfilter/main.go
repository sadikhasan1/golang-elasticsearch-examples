package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"github.com/elastic/go-elasticsearch/v8"
)

func main() {
	// Create an Elasticsearch client
	esCfg := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	}
	esClient, err := elasticsearch.NewClient(esCfg)
	if err != nil {
		log.Fatalf("Error creating Elasticsearch client: %s", err)
	}

	// Build the query using the query builder
	query := CustomerListQuery(0, 10, "John", "", "New York", "", "", "")

	// Convert the query map to JSON
	queryJSON, err := json.Marshal(query)
	if err != nil {
		log.Fatalf("Error marshaling query to JSON: %s", err)
	}

	// Send the query to Elasticsearch
	res, err := esClient.Search(
		esClient.Search.WithIndex("your_index_name"), // Replace with your actual index name
		esClient.Search.WithBody(strings.NewReader(string(queryJSON))),
	)
	if err != nil {
		log.Fatalf("Error executing Elasticsearch query: %s", err)
	}
	defer res.Body.Close()

	// Parse and print the response
	var resMap map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&resMap); err != nil {
		log.Fatalf("Error decoding response: %s", err)
	}
	fmt.Println(resMap)
}


func CustomerListQuery(from int32, size int32, FirstName string, LastName string, City string, email string, phoneNumber string, search string) map[string]interface{} {
	filters := map[string]string{
		"FirstName": FirstName,
		"LastName":   LastName,
		"City":            City,
		"Email":                   email,
		"PhoneNumber":             phoneNumber,
	}

	query := map[string]interface{}{
		"from": from,
		"size": size,
		"sort": []map[string]interface{}{
			{
				"CreatedAt": map[string]interface{}{
					"order": "desc",
				},
			},
		},
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []map[string]interface{}{
					{
						"match_all": map[string]interface{}{},
					},
				},
			},
		},
	}

	filterConditions := make([]map[string]interface{}, 0)

	for field, value := range filters {
		if value != "" {
			filterConditions = append(filterConditions, BuildFilterCondition(field, value))
		}
	}

	fieldsToSearch := []string{"FirstName", "LastName"} // Add other fields to search here
	if len(fieldsToSearch) > 0 && search != "" {
		filterConditions = append(filterConditions, BuildMultiMatchFilterCondition(fieldsToSearch, search))
	}

	if len(filterConditions) > 0 {
		query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"] = append(
			query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"].([]map[string]interface{}),
			map[string]interface{}{
				"bool": map[string]interface{}{
					"must": filterConditions,
				},
			},
		)
	}
	return query
}

func BuildFilterCondition(field, value string) map[string]interface{} {
	return map[string]interface{}{
		"match": map[string]interface{}{
			field: value,
		},
	}
}

func BuildMultiMatchFilterCondition(fields []string, value string) map[string]interface{} {
	return map[string]interface{}{
		"multi_match": map[string]interface{}{
			"query":  value,
			"fields": fields,
			"type":   "bool_prefix", // Adjust the type as needed
		},
	}
}
