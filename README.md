# Customer Data Indexing to Elasticsearch and Elasticsearch Query Builder

This repository provides a program to index customer data into Elasticsearch using the provided query builder and the `github.com/elastic/go-elasticsearch/v8` package, along with a query builder utility for constructing Elasticsearch queries in Go.

## Prerequisites

Before using the provided code, make sure you have the following prerequisites installed:

- Go programming language (https://golang.org/)
- Elasticsearch server (https://www.elastic.co/guide/en/elasticsearch/reference/current/install-elasticsearch.html)

## Installation

To use the provided code, you need to have the necessary packages installed. Use the following commands to install them:

```sh
go get github.com/elastic/go-elasticsearch/v8
```

## Customer Data Indexing to Elasticsearch

The provided program demonstrates how to index customer data into Elasticsearch using the query builder and the `github.com/elastic/go-elasticsearch/v8` package.

### Usage

1. Import necessary packages and define the `Customer` struct:

```go
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
```

2. Create an Elasticsearch client:

```go
esCfg := elasticsearch.Config{
	Addresses: []string{"http://localhost:9200"},
}
esClient, err := elasticsearch.NewClient(esCfg)
if err != nil {
	log.Fatalf("Error creating Elasticsearch client: %s", err)
}
```

3. Index customer data into Elasticsearch:

```go
indexName := "customer_index" // Replace with your desired index name
for _, customer := range customerData {
	if err := indexCustomer(esClient, indexName, customer); err != nil {
		log.Printf("Error indexing customer %d: %s", customer.ID, err)
	} else {
		log.Printf("Customer %d indexed successfully", customer.ID)
	}
}
```

4. Replace `indexName` and the sample `customerData` with your actual index name and customer data.

## Elasticsearch Query Builder

The Elasticsearch query builder package (`querybuilder`) offers a set of functions to construct custom Elasticsearch queries with specified filters and search conditions.

### Usage

1. Import the necessary packages:

```go
import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"your/querybuilder/package/import/path"
)
```

2. Create an Elasticsearch client:

```go
esCfg := elasticsearch.Config{
	Addresses: []string{"http://localhost:9200"},
}
esClient, err := elasticsearch.NewClient(esCfg)
if err != nil {
	log.Fatalf("Error creating Elasticsearch client: %s", err)
}
```

3. Build the query using the query builder:

```go
query := querybuilder.CustomerListQuery(0, 10, "John", "", "New York", "", "", "")
```

4. Convert the query map to JSON:

```go
queryJSON, err := json.Marshal(query)
if err != nil {
	log.Fatalf("Error marshaling query to JSON: %s", err)
}
```

5. Send the query to Elasticsearch:

```go
res, err := esClient.Search(
	esClient.Search.WithIndex("your_index_name"), // Replace with your actual index name
	esClient.Search.WithBody(strings.NewReader(string(queryJSON))),
)
if err != nil {
	log.Fatalf("Error executing Elasticsearch query: %s", err)
}
defer res.Body.Close()
```

6. Parse and print the response:

```go
var resMap map[string]interface{}
if err := json.NewDecoder(res.Body).Decode(&resMap); err != nil {
	log.Fatalf("Error decoding response: %s", err)
}
fmt.Println(resMap)
```

## Running the Program

1. Ensure your Elasticsearch server is running.

2. Run the program:

```sh
go run main.go
```

## License

This project is licensed under the [MIT License](LICENSE).
```
