package returner

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/johnmikee/ctcc/tcc"
)

type ManagedItem struct {
	Name        string                 `json:"name"`
	DateManaged string                 `json:"date_managed"`
	Status      string                 `json:"status"`
	Data        map[string]interface{} `json:"data"`
}

type TCCInfo struct {
	ManagedItems map[string]ManagedItem
}

type Facts struct {
	CheckinModuleVersion string `json:"checkin_module_version"`
}

type Report struct {
	Facts        Facts                  `json:"facts"`
	ManagedItems map[string]ManagedItem `json:"managed_items"`
}

type TCCReport struct {
	MDM    []*tcc.MDMEntry
	System []*tcc.TCCEntry
	User   []*tcc.TCCEntry
}

func addServiceDetail(sd []tcc.ServiceDetail, data map[string]interface{}) map[string]interface{} {
	for _, detail := range sd {
		data[fmt.Sprintf("%s (%s)", detail.Service, detail.FriendlyName)] = detail.Allowed
	}
	return data
}

func serviceDetailString(sd []tcc.ServiceDetail) string {
	var s []string
	for _, detail := range sd {
		s = append(s, fmt.Sprintf("%s (%s): %t", detail.Service, detail.FriendlyName, detail.Allowed))
	}
	return strings.Join(s, ", ")
}

func convert(t *TCCReport) map[string]Report {
	mdmReport := make(map[string]ManagedItem)
	for _, entry := range t.MDM {
		data := make(map[string]interface{})
		data["source"] = entry.Source
		data["mdm_server"] = entry.MDMServer
		data["code_requirement"] = entry.CodeRequirement
		data["identifier"] = entry.Identifier
		data["identifier_type"] = entry.IdentifierType
		data["service_friendly_name"] = entry.ServiceFriendlyName
		data["static_code"] = entry.StaticCode
		data = addServiceDetail(entry.Services, data)

		mdmReport[entry.Identifier] = ManagedItem{
			Name:        entry.Identifier,
			DateManaged: time.Now().UTC().Format(time.RFC3339),
			Status:      "PRESENT",
			Data:        data,
		}
	}

	systemReport := make(map[string]ManagedItem)
	for _, entry := range t.System {
		data := make(map[string]interface{})
		data["source"] = entry.Source
		data["username"] = entry.Username
		data["client"] = entry.Client
		data["client_id"] = entry.ClientID
		data["service_name"] = entry.ServiceName
		data["service_friendly_name"] = entry.ServiceFriendlyName
		data["auth_value"] = entry.AuthValue
		data["auth_reason"] = entry.AuthReason
		data["timestamp"] = entry.Timestamp
		data["formatted_time"] = entry.FormattedTime
		data["code_sign_req"] = entry.CodeSignReq

		systemReport[entry.Client] = ManagedItem{
			Name:        entry.Client,
			DateManaged: time.Now().UTC().Format(time.RFC3339),
			Status:      "PRESENT",
			Data:        data,
		}
	}

	userReport := make(map[string]ManagedItem)
	for _, entry := range t.User {
		data := make(map[string]interface{})
		data["source"] = entry.Source
		data["username"] = entry.Username
		data["client"] = entry.Client
		data["client_id"] = entry.ClientID
		data["service_name"] = entry.ServiceName
		data["service_friendly_name"] = entry.ServiceFriendlyName
		data["auth_value"] = entry.AuthValue
		data["auth_reason"] = entry.AuthReason
		data["timestamp"] = entry.Timestamp
		data["formatted_time"] = entry.FormattedTime
		data["code_sign_req"] = entry.CodeSignReq

		userReport[entry.Client] = ManagedItem{
			Name:        entry.Client,
			DateManaged: time.Now().UTC().Format(time.RFC3339),
			Status:      "PRESENT",
			Data:        data,
		}
	}

	system := Report{
		Facts: Facts{
			CheckinModuleVersion: "1.0",
		},
		ManagedItems: systemReport,
	}

	user := Report{
		Facts: Facts{
			CheckinModuleVersion: "1.0",
		},
		ManagedItems: userReport,
	}

	mdm := Report{
		Facts: Facts{
			CheckinModuleVersion: "1.0",
		},
		ManagedItems: mdmReport,
	}

	report := make(map[string]Report)
	report["TCC-MDMOverrides"] = mdm
	report["TCC-System"] = system
	report["TCC-User"] = user

	return report
}

func CSV(fileName string, t *TCCReport) error {
	err := mdmCsv("MDM-"+fileName, t.MDM)
	if err != nil {
		return err
	}

	systemUser := append(t.System, t.User...)
	err = tccCsv(fileName, systemUser)
	if err != nil {
		return err
	}

	return nil
}

func tccCsv(fileName string, entry []*tcc.TCCEntry) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	// Write the header row to the CSV file
	header := []string{"Source", "Username", "Client", "ClientID", "ServiceName", "ServiceFriendlyName", "AuthValue", "AuthReason", "Timestamp", "FormattedTime", "CodeSignReq"}
	writer.Write(header)
	for _, item := range entry {
		record := []string{
			item.Source,
			item.Username,
			item.Client,
			item.ClientID,
			item.ServiceName,
			item.ServiceFriendlyName,
			item.AuthValue,
			item.AuthReason,
			strconv.FormatInt(item.Timestamp, 10),
			item.FormattedTime,
			item.CodeSignReq,
		}
		writer.Write(record)
	}

	return nil
}

func mdmCsv(fileName string, entry []*tcc.MDMEntry) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	header := []string{"Source", "MDMServer", "CodeRequirement", "Identifier", "IdentifierType", "ServiceFriendlyName", "Services", "StaticCode"}
	writer.Write(header)

	for _, item := range entry {
		record := []string{
			item.Source,
			item.MDMServer,
			item.CodeRequirement,
			item.Identifier,
			item.IdentifierType,
			item.ServiceFriendlyName,
			serviceDetailString(item.Services),
			item.StaticCode,
		}
		writer.Write(record)
	}
	return nil
}

func Sal(t *TCCReport) (string, error) {
	report := convert(t)

	fmt.Println(report)
	file, err := os.Create("/usr/local/sal/tcc_results.json")
	if err != nil {
		return "", fmt.Errorf("error creating file: %s", err)
	}
	defer file.Close()

	data, err := toJson(report)
	if err != nil {
		return "", fmt.Errorf("error marshalling report: %s", err)
	}

	_, err = file.Write(data)
	if err != nil {
		return "", fmt.Errorf("error writing to file: %s", err)
	}

	return string(data), nil
}

func Json(t *TCCReport) (string, error) {
	data, err := toJson(t)
	if err != nil {
		return "", fmt.Errorf("error marshalling report: %s", err)
	}

	return string(data), nil
}

func toJson(report interface{}) ([]byte, error) {
	data, err := json.Marshal(report)
	if err != nil {
		return nil, fmt.Errorf("error marshalling report: %s", err)
	}

	return data, nil
}
