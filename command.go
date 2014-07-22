package main

import (
	"bytes"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/wsxiaoys/terminal/color"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"sort"
	"strings"
)

var Commands = []cli.Command{
	commandGet,
	commandPost,
	commandPut,
	commandDelete,
}

var commandGet = cli.Command{
	Name:        "get",
	ShortName:   "g",
	Usage:       "get a resource",
	Description: "",
	Action:      doGet,
}

var commandPost = cli.Command{
	Name:        "post",
	ShortName:   "c",
	Usage:       "create a new resource",
	Description: "",
	Action: func(c *cli.Context) {
		println("post")
	},
}
var commandPut = cli.Command{
	Name:        "put",
	ShortName:   "u",
	Usage:       "update a resource",
	Description: "",
	Action: func(c *cli.Context) {
		println("put")
	},
}

var commandDelete = cli.Command{
	Name:        "delete",
	ShortName:   "d",
	Usage:       "delete a resource",
	Description: "",
	Action: func(c *cli.Context) {
		println("delete")
	},
}

func doGet(c *cli.Context) {

	client := &http.Client{}
	args := c.Args()
	fields := make(map[string]string)
	queries := make(map[string]string)
	if len(args) >= 1 {
		params := []string(args)[1:len(args)]
		field_regexp, _ := regexp.Compile("[^=:]+:[^=:]+")
		query_regexp, _ := regexp.Compile("[^=:]+=[^=:]+")
		for i := range params {
			param := params[i]
			switch {
			case field_regexp.MatchString(param):
				d := strings.Split(param, ":")
				fields[d[0]] = d[1]
			case query_regexp.MatchString(params[i]):
				d := strings.Split(param, "=")
				queries[d[0]] = d[1]
			}
		}
	}

	baseUrl, _ := url.Parse(args.Get(0))
	queryValues := url.Values{}
	for key, value := range queries {
		queryValues.Add(key, value)
	}
	if len(queryValues) != 0 {
		baseUrl.RawQuery = queryValues.Encode()
	}
	req, reqErr := http.NewRequest("GET", baseUrl.String(), nil)
	if reqErr != nil {
		fmt.Println(reqErr)
		os.Exit(1)
	}

	resp, respErr := client.Do(req)
	if respErr != nil {
		fmt.Println(respErr)
		os.Exit(1)
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	color.Printf("@{m}%s @{y}%s\n", resp.Proto, resp.Status)
	sortedHeader := make([]string, len(resp.Header))
	i := 0
	for key, _ := range resp.Header {
		sortedHeader[i] = key
		i++
	}
	sort.Strings(sortedHeader)
	for _, key := range sortedHeader {
		if key != "Set-Cookie" && key != "P3p" {
			color.Printf("%s: @{c}%s\n", key, resp.Header[key][0])
		}
	}
	//fmt.Println(buf.String())
}
