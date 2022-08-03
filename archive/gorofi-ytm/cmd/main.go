package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	gorofiytm "github.com/Meta502/gorofi-ytm"
	"github.com/joho/godotenv"
)

func main() {
	errorObject, _ := json.Marshal(gorofiytm.Output{
		Message: "An error occurred.",
		Prompt:  "error",
	})

	err := godotenv.Load()
	if err != nil {
		fmt.Println(string(errorObject))
		return
	}

	yt, err := gorofiytm.NewYoutubeClient()

	if err != nil {
		fmt.Println(string(errorObject))
		return
	}

	initialInput := gorofiytm.Output{
		InputAction: "send",
		Prompt:      "Search YouTube",
	}

	jsonData, _ := json.Marshal(initialInput)

	fmt.Println(string(jsonData))

	stage := 0

	for {
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')

		var inputObject gorofiytm.Input

		err := json.Unmarshal([]byte(input), &inputObject)

		if err != nil {
			fmt.Println(string(errorObject))
		}

		if inputObject.Name == "input change" && stage == 0 {
			inputMode, _ := json.Marshal(gorofiytm.Output{
				InputAction: "send",
				Prompt:      "Search YouTube",
				Lines:       yt.GetSuggestions(inputObject.Value),
				Input: inputObject.Value,
			})
			fmt.Println(string(inputMode))
		}

		if inputObject.Name == "select entry" {
			if stage == 0 {
				stage = 1
				inputMode, _ := json.Marshal(gorofiytm.Output{
					InputAction: "filter",
					Prompt: "Searching",
					ActiveEntry: 0,
					Input: "",
					Lines: []gorofiytm.Line{},
				})
				fmt.Println(string(inputMode))

				data, err := yt.GetSearchResults(inputObject.Value)

				if err != nil {
					error, _ := json.Marshal(gorofiytm.Output{
						Prompt: "error",
						Message: err.Error(),
						InputAction: "send",
					})

					fmt.Println(string(error))
					continue
				}

				lines := []gorofiytm.Line{}

				for _, v := range data.Items {
					lines = append(lines, gorofiytm.Line{
						Text: fmt.Sprintf("%s\n%s", v.Snippet.ChannelTitle, v.Snippet.Title),
						Data: v.ID.VideoID,
					})
				}

				result, _ := json.Marshal(gorofiytm.Output{
					InputAction: "filter",
					Prompt: "Search Results",
					ActiveEntry: 0,
					Lines: lines,
				})

				fmt.Println(string(result))
			} else if stage == 1 {
				cmd := exec.Command("mpv", "https://youtube.com/watch?v=" + inputObject.Data)
				cmd.Start()
				return
			}
		}
	}
}
