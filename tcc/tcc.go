package tcc

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const (
	systemTCCDB  = "/Library/Application Support/com.apple.tcc/tcc.db"
	defaultQuery = "select client, client_type, service, auth_value, auth_reason, last_modified, quote(csreq) from access order by client, auth_value;"
	mdmOverrides = "/Library/Application Support/com.apple.TCC/MDMOverrides.plist"
)

// SQR is a struct for the response from the TCC database
type SQResponse struct {
	Client      string `json:"client"`
	ClientType  int    `json:"client_type"`
	Service     string `json:"service"`
	AuthValue   int    `json:"auth_value"`
	AuthReason  int    `json:"auth_reason"`
	LastMod     int64  `json:"last_modified"`
	Csreq       string `json:"csreq"`
	AuthVersion string `json:"auth_version"`
}

// SystemQuery returns a list of information from the system TCC database
func SystemQuery() ([]SQResponse, error) {
	return tccDBQuery(systemTCCDB, defaultQuery)
}

// UserQuery returns a list of information from the user's TCC database
func UserQuery(userDB string) ([]SQResponse, error) {
	return tccDBQuery(userDB, defaultQuery)
}

func tccDBQuery(dbPath string, query string) ([]SQResponse, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var responses []SQResponse

	for rows.Next() {
		var sqr SQResponse
		err = rows.Scan(&sqr.Client, &sqr.ClientType, &sqr.Service, &sqr.AuthValue, &sqr.AuthReason, &sqr.LastMod, &sqr.Csreq)
		if err != nil {
			return nil, err
		}
		responses = append(responses, sqr)

	}

	return responses, nil
}
