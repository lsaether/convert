package main

import (
	"bufio";
	"bytes";
	"fmt";
	"encoding/json";
	"os";
	"net/http";
	"strings";
	"strconv"
)

func main() {
	// Read standard input. Expects a bitcoin value with 8 decimal places.
	reader := bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString('\n')

	// Parse the first arg for desired currency, choices are USD or EUR.
	curr := ""
	if len(os.Args) > 1 {
		curr = os.Args[1]
		curr = strings.ToUpper(curr)

		validCurrs := map[string]bool {
			"USD": true,
			"EUR": true,
		}
	
		if !validCurrs[curr] {
			panic("Must specify either USD or EUR.")
		}	
	} else {
		curr = "USD"
	}

	// Hit the API
	uri := "https://api.coindesk.com/v1/bpi/currentprice.json"

	resp, err := http.Get(uri)
	if err != nil {
		panic(err)
	} 

	defer resp.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	var f interface{}
	err = json.Unmarshal(buf.Bytes(), &f)
	m := f.(map[string]interface{})
	r := m["bpi"].(map[string]interface{})
	c := r[curr].(map[string]interface{})
	i, _ := strconv.ParseFloat(strings.TrimSpace(
		strings.Join(
			strings.Split(c["rate"].(string), ","),
			"",
		),
	), 64)
	v, _ := strconv.ParseFloat(strings.TrimSpace(value), 64)

	// fmt.Println(c["rate"], value)
	fmt.Println(i * v, curr)
}