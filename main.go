package main

import (
	"flag"
	"fmt"
	"sync"

	"github.com/johnmikee/ctcc/returner"
	"github.com/johnmikee/ctcc/tcc"
	"github.com/johnmikee/ctcc/version"
)

type Cfg struct {
	FileName string `json:"file_name"`
	ToCSV    bool   `json:"to_csv"`
	ToJSON   bool   `json:"to_json"`
	ToSal    bool   `json:"to_sal"`
	Version  bool   `json:"version"`
}

func systemDefaultPermissions() []*tcc.TCCEntry {
	resp, err := tcc.SystemQuery()
	if err != nil {
		fmt.Printf("error querying system: %s", err)
		return nil
	}

	entries := []*tcc.TCCEntry{}
	for _, r := range resp {
		row := tcc.ProcessRow("System", r)
		entries = append(entries, row)
	}

	return entries
}

func userPermissionOverrides() []*tcc.TCCEntry {
	localUsers := tcc.ListUsers()
	entries := []*tcc.TCCEntry{}
	entryChan := make(chan *tcc.TCCEntry)

	var wg sync.WaitGroup

	for _, user := range localUsers {
		wg.Add(1)
		go func(u tcc.Users) {
			defer wg.Done()
			resp, err := tcc.UserQuery(u.DB)
			if err != nil {
				fmt.Printf("error querying user %s: %s", u.Name, err)
				return
			}
			for _, r := range resp {
				row := tcc.ProcessRow(u.Name, r)
				entryChan <- row
			}
		}(user)
	}

	go func() {
		wg.Wait()
		close(entryChan)
	}()

	for entry := range entryChan {
		entries = append(entries, entry)
	}

	return entries
}

func mdmProfileOverrides() []*tcc.MDMEntry {
	overrides, err := tcc.CheckMDMOverrides()
	if err != nil {
		return nil
	}
	entries := []*tcc.MDMEntry{}
	for item, vals := range overrides {
		entry := tcc.ProcessMDMOverrides(item, vals)
		entries = append(entries, entry)
	}

	return entries
}

func main() {
	o := Cfg{
		ToCSV:   true,
		ToJSON:  false,
		ToSal:   false,
		Version: false,
	}

	flag.StringVar(&o.FileName, "file", "ctcc.csv", "File name to output to")
	flag.BoolVar(&o.ToCSV, "csv", false, "Output to CSV")
	flag.BoolVar(&o.ToJSON, "json", false, "Output to JSON")
	flag.BoolVar(&o.ToSal, "sal", false, "Output for sal to pick it up")
	flag.BoolVar(&o.Version, "version", false, "Print version and exit")
	flag.Parse()

	if o.Version {
		version.Print()
		return
	}

	var m []*tcc.MDMEntry
	var s []*tcc.TCCEntry
	var u []*tcc.TCCEntry

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		m = mdmProfileOverrides()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		s = systemDefaultPermissions()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		u = userPermissionOverrides()
	}()

	wg.Wait()

	// create the report after all functions have finished
	report := &returner.TCCReport{
		MDM:    m,
		System: s,
		User:   u,
	}

	if o.ToCSV {
		err := returner.CSV(o.FileName, report)
		if err != nil {
			fmt.Printf("error generating csv: %s", err)
			return
		}
	}
	if o.ToJSON {
		report, err := returner.Json(report)
		if err != nil {
			fmt.Printf("error marshalling report: %s", err)
			return
		}
		fmt.Println(report)
	}
	if o.ToSal {
		report, err := returner.Sal(report)
		if err != nil {
			fmt.Printf("error generating report: %s", err)
			return
		}
		fmt.Println(report)
	}
}
