package actions

import (
	"fmt"
	mSlack "go-bot-test/lib/m_slack"
	"strconv"
	"strings"
)

func calcResultCallback(params mSlack.RequestParams) error {
	if err := mSlack.DeleteOriginal(params); err != nil {
		return err
	}

	args := strings.Split(params.ActionParams.Value, " ")
	num1, _ := strconv.Atoi(args[0])
	operator := args[1]
	num2, _ := strconv.Atoi(args[2])

	result := 0
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "x":
		result = num1 * num2
	case "÷":
		result = num1 / num2
	}
	blocks := []mSlack.Block{
		mSlack.Text{
			Body: fmt.Sprintf("%s = %d", params.ActionParams.Value, result),
		},
	}

	if err := mSlack.Post(params, blocks); err != nil {
		return err
	}
	return nil
}
