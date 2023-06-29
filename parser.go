package main

import (
	"fmt"
	"github.com/tslight/naeq/assets/books"
	"github.com/tslight/naeq/pkg/alw"
	"github.com/tslight/naeq/pkg/json"
	"log"
	"strconv"
)

func Parse(words string, count int, efsBook string) (string, error) {

	i, err := alw.GetSum(words)
	if err != nil {
		log.Fatalln(err)
	}

	var book map[string]interface{}
	book, err = json.FromEFSPath(books.EFS, fmt.Sprint(efsBook, ".json"))
	if err != nil {
		log.Fatalln(err)
	}

	matches := alw.GetMatches(i, book)

	output := "Sum: " + strconv.Itoa(i) + "\nFirst " + strconv.Itoa(count) + " matches in " + efsBook + ": "

	for k, v := range matches {
		if count > 0 && k >= count {
			break
		}
		// append to output with a space, unless it's the last one
		if k == len(matches)-1 || k == count-1 {
			output = fmt.Sprint(output, v)
		} else {
			output = fmt.Sprint(output, v, ", ")
		}
	}

	fmt.Println(output)

	return output, nil
}
