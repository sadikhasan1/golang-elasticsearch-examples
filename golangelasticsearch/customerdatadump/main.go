package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
)

// Customer represents customer data
type Customer struct {
	ID          int    `json:"id"`
	FirstName   string `json:"FirstName"`
	LastName    string `json:"LastName"`
	City        string `json:"City"`
	Email       string `json:"Email"`
	PhoneNumber string `json:"PhoneNumber"`
}

func main() {
	// Create an Elasticsearch client
	esCfg := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	}
	esClient, err := elasticsearch.NewClient(esCfg)
	if err != nil {
		log.Fatalf("Error creating Elasticsearch client: %s", err)
	}

	// Sample customer data
	customerData := []Customer{
		{ID: 1, FirstName: "John", LastName: "Doe", City: "New York", Email: "john@example.com", PhoneNumber: "123-456-7890"},
		{ID: 2, FirstName: "Jane", LastName: "Smith", City: "Los Angeles", Email: "jane@example.com", PhoneNumber: "987-654-3210"},
		// Add more customer data here
	}

	// Index the customer data into Elasticsearch
	indexName := "customer_index" // Replace with your desired index name
	for _, customer := range customerData {
		if err := indexCustomer(esClient, indexName, customer); err != nil {
			log.Printf("Error indexing customer %d: %s", customer.ID, err)
		} else {
			log.Printf("Customer %d indexed successfully", customer.ID)
		}
	}
}

func indexCustomer(client *elasticsearch.Client, indexName string, customer Customer) error {
	customerJSON, err := json.Marshal(customer)
	if err != nil {
		return fmt.Errorf("error marshaling customer data: %s", err)
	}

	// Index the customer data
	_, err = client.Index(
		indexName,
		strings.NewReader(string(customerJSON)),
		client.Index.WithDocumentID(fmt.Sprintf("%d", customer.ID)),
		client.Index.WithRefresh("true"),
	)
	if err != nil {
		return fmt.Errorf("error indexing customer: %s", err)
	}

	return nil
}

