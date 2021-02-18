package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/Ideolys/carbone-sdk-go/carbone"
)

func main() {
	key := "test_eyJhbGciOiJFUzUxMiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiIzNzk5IiwiYXVkIjoiY2FyYm9uZSIsImV4cCI6MjI0NDM2NDMyMSwiZGF0YSI6eyJpZEFjY291bnQiOjM3OTl9fQ.AVbLnXNvsuV32F7nucMvOw59s0ag-Ea0h_yoAoFAiC_-5k-fIaBPAU6MVflPruJ3ZKC495cNyZ3uawlHJNViffmBAEiltx8YytaUKC-ie1hJHzKXpCkYUSlenZQoDIqAYQoOtizZDLQ8ZzaQ0Nk9NU6H0szB4pWomhTx7NGHK1Aci2dD"

	csdk, err := carbone.NewCarboneSDK(key)
	if err != nil {
		log.Fatal(err)
	}
	// Path to your template
	templateID := "./test.odt"

	file, err := ioutil.ReadFile("./report.json")
	// Add your data here

	jsonData := `{"data":` + string(file) + `,"convertTo":"pdf"}`

	fmt.Println(jsonData, err)
	reportBuffer, err := csdk.Render(templateID, jsonData)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("Report.pdf", reportBuffer, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
