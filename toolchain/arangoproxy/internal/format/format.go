package format

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/arangodb/docs/migration-tools/arangoproxy/internal/utils"
	"github.com/dlclark/regexp2"
)

/*
	Functions to reformat inputs/outputs
*/

func AdjustCodeForArangosh(code string) string {
	out := ""

	if !(strings.Contains(code, "EOFD")) {
		code = fmt.Sprintf("%s\nprint('EOFD');\n\n\n\n", code)
	}

	code = strings.ReplaceAll(code, "\r\n", "\n")

	lines := strings.Split(code, "\n")

	for _, line := range lines {
		re := regexp.MustCompile(`(?m)let |const `)
		line = re.ReplaceAllString(line, "var ")

		tildeRE := regexp.MustCompile(`(?m)^\s*~`)

		if tildeRE.MatchString(line) {
			line = tildeRE.ReplaceAllString(line, "")
			line = fmt.Sprintf("print('HIDED-START')\n%s\nprint('HIDED-END');\n", line)
		}

		re = regexp.MustCompile(`(?m)}\n *catch`)
		line = re.ReplaceAllString(line, "} catch")

		out = fmt.Sprintf("%s\n%s", out, line)
	}

	return out
}

/*
	Custom formatters
*/

/*
	JS Formatter
*/

/*
	Curl Formatter
*/

type CurlFormatter struct {
}

func (formatter CurlFormatter) FormatCommand(code string) string {
	multiLineRE := regexp.MustCompile(`(?m)[+]\s*\n*`)
	allMultiLines := multiLineRE.FindAllString(code, -1)
	for _, multiLine := range allMultiLines {
		noMoreMultiline := strings.ReplaceAll(multiLine, "\n", "")
		code = strings.Replace(code, multiLine, noMoreMultiline, -1)
	}

	return code
}

func (formatter CurlFormatter) FormatCurlOutput(arangoOutput, renderOption string) (input, output string, err error) {
	curlRE := regexp2.MustCompile(`(?ms)REQ(.*?)ENDREQ`, 0)
	curlRequests := utils.Regexp2FindAllString(curlRE, arangoOutput)
	for _, curl := range curlRequests {
		input = fmt.Sprintf("%s\n%s", input, curl)
	}

	input = strings.ReplaceAll(input, "REQ", "")
	input = strings.ReplaceAll(input, "ENDREQ", "")
	input = strings.ReplaceAll(input, "END", "\n")

	respRE := regexp2.MustCompile(`(?ms)RESP(.*?)ENDRESP`, 0)
	responses := utils.Regexp2FindAllString(respRE, arangoOutput)
	for _, response := range responses {
		output = fmt.Sprintf("%s\n%s", output, response)
	}

	output = strings.ReplaceAll(output, "RESP", "")
	output = strings.ReplaceAll(output, "ENDRESP", "")
	output = strings.ReplaceAll(output, "END", "\n")

	return input, output, nil
}

/*
	AQL Formatter
*/

type AQLFormatter struct{}

func (formatter AQLFormatter) FormatRequestCode(code string, bindVars map[string]interface{}) string {
	commands := fmt.Sprintf("db._query(`%s`", code)
	if len(bindVars) != 0 {
		bindVarsJson, _ := json.Marshal(bindVars)
		commands = fmt.Sprintf("%s, %s", commands, bindVarsJson)
	}

	commands = commands + ");"
	return commands
}

/*
	Openapi Formatter
*/

type OpenapiFormatter struct{}

func (formatter OpenapiFormatter) EditDescriptions(req []byte) []byte {
	payloadString := string(req)

	payloadString = strings.Replace(payloadString, "{{< warning >}}", "> **WARNING:**", -1)
	payloadString = strings.Replace(payloadString, "{{< info >}}", "> **INFO:**", -1)
	payloadString = strings.Replace(payloadString, "{{< danger >}}", "> **DANGER:**", -1)
	payloadString = strings.Replace(payloadString, "{{< success >}}", "> **SUCCESS:**", -1)
	payloadString = strings.Replace(payloadString, "{{< tip >}}", "> **TIP:**", -1)

	payloadString = strings.Replace(payloadString, "{{< /tip >}}", "", -1)
	payloadString = strings.Replace(payloadString, "{{< /warning >}}", "", -1)
	payloadString = strings.Replace(payloadString, "{{< /info >}}", "", -1)
	payloadString = strings.Replace(payloadString, "{{< /success >}}", "", -1)
	payloadString = strings.Replace(payloadString, "{{< /danger >}}", "", -1)

	return []byte(payloadString)
}
