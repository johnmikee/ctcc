package tcc

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func decodeCodeSignature(codeHex string) string {
	cmdXxd := exec.Command("xxd", "-r", "-p")
	cmdXxd.Stdin = strings.NewReader(codeHex)
	outXxd, errXxd := cmdXxd.Output()
	if errXxd != nil {
		return formatStringWithQuotes(codeHex)
	}

	cmdCsreq := exec.Command("csreq", "-r", "-", "-t")
	cmdCsreq.Stdin = bytes.NewReader(outXxd)

	output, errCsreq := cmdCsreq.CombinedOutput()
	if errCsreq != nil {
		return formatStringWithQuotes(codeHex)
	}

	csreqOutput := string(output)

	return formatStringWithQuotes(csreqOutput)
}

func epochToTime(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	return t.Format("2006-01-02T15:04-0700")
}

func formatStringWithQuotes(str string) string {
	return strings.TrimSpace(fmt.Sprintf("\"%s\"", strings.Replace(str, "\n", "", -1)))
}

func getAppNameFromClient(client string) string {
	cmd := exec.Command("mdfind", fmt.Sprintf("kMDItemCFBundleIdentifier = %s", client))
	output, err := cmd.Output()
	if err != nil {
		return client
	}

	if len(output) <= 1 {
		return client
	}

	return strings.TrimSpace(string(output))
}

func removeDupes(s []string) []string {
	unique := make(map[string]bool)
	result := []string{}

	for _, value := range s {
		if !unique[value] {
			unique[value] = true
			result = append(result, value)
		}
	}

	return result
}

func sort(input []string) string {
	input = removeDupes(input)
	for i, x := range input {
		if len(x) == 0 {
			input = append(input[:i], input[i+1:]...)
		}
	}

	switch len(input) {
	case 0:
		return ""
	case 1:
		return input[0]
	default:
		return strings.Join(input, ", ")
	}
}

func sourcerer(source string) string {
	switch source {
	case "System":
		return "System"
	default:
		return "User"
	}
}
