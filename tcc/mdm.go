package tcc

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"howett.net/plist"
)

type KTCCServiceAccessibility struct {
	Allowed             bool        `plist:"Allowed,omitempty"`
	CodeRequirement     string      `plist:"CodeRequirement,omitempty"`
	CodeRequirementData interface{} `plist:"CodeRequirementData,omitempty"`
	Comment             string      `plist:"Comment,omitempty"`
	Identifier          string      `plist:"Identifier,omitempty"`
	IdentifierType      string      `plist:"IdentifierType,omitempty"`
	StaticCode          bool        `plist:"StaticCode,omitempty"`
}

type KTCCServiceAddressBook struct {
	Allowed             bool        `plist:"Allowed,omitempty"`
	CodeRequirement     string      `plist:"CodeRequirement,omitempty"`
	CodeRequirementData interface{} `plist:"CodeRequirementData,omitempty"`
	Identifier          string      `plist:"Identifier,omitempty"`
	IdentifierType      string      `plist:"IdentifierType,omitempty"`
	StaticCode          bool        `plist:"StaticCode,omitempty"`
}

type KTCCServiceCalendar struct {
	Allowed             bool        `plist:"Allowed,omitempty"`
	CodeRequirement     string      `plist:"CodeRequirement,omitempty"`
	CodeRequirementData interface{} `plist:"CodeRequirementData,omitempty"`
	Identifier          string      `plist:"Identifier,omitempty"`
	IdentifierType      string      `plist:"IdentifierType,omitempty"`
	StaticCode          bool        `plist:"StaticCode,omitempty"`
}

type KTCCServiceFileProviderPresence struct {
	Allowed             bool        `plist:"Allowed,omitempty"`
	CodeRequirement     string      `plist:"CodeRequirement,omitempty"`
	CodeRequirementData interface{} `plist:"CodeRequirementData,omitempty"`
	Identifier          string      `plist:"Identifier,omitempty"`
	IdentifierType      string      `plist:"IdentifierType,omitempty"`
	StaticCode          bool        `plist:"StaticCode,omitempty"`
}

type KTCCServicePhotos struct {
	Allowed             bool        `plist:"Allowed,omitempty"`
	CodeRequirement     string      `plist:"CodeRequirement,omitempty"`
	CodeRequirementData interface{} `plist:"CodeRequirementData,omitempty"`
	Identifier          string      `plist:"Identifier,omitempty"`
	IdentifierType      string      `plist:"IdentifierType,omitempty"`
	StaticCode          bool        `plist:"StaticCode,omitempty"`
}

type KTCCServicePostEvent struct {
	Allowed             bool        `plist:"Allowed,omitempty"`
	CodeRequirement     string      `plist:"CodeRequirement,omitempty"`
	CodeRequirementData interface{} `plist:"CodeRequirementData,omitempty"`
	Identifier          string      `plist:"Identifier,omitempty"`
	IdentifierType      string      `plist:"IdentifierType,omitempty"`
	StaticCode          bool        `plist:"StaticCode,omitempty"`
}

type KTCCServiceSystemPolicyDesktopFolder struct {
	Allowed             bool        `plist:"Allowed,omitempty"`
	CodeRequirement     string      `plist:"CodeRequirement,omitempty"`
	CodeRequirementData interface{} `plist:"CodeRequirementData,omitempty"`
	Identifier          string      `plist:"Identifier,omitempty"`
	IdentifierType      string      `plist:"IdentifierType,omitempty"`
	StaticCode          bool        `plist:"StaticCode,omitempty"`
}

type KTCCServiceSystemPolicyDocumentsFolder struct {
	Allowed             bool        `plist:"Allowed,omitempty"`
	CodeRequirement     string      `plist:"CodeRequirement,omitempty"`
	CodeRequirementData interface{} `plist:"CodeRequirementData,omitempty"`
	Identifier          string      `plist:"Identifier,omitempty"`
	IdentifierType      string      `plist:"IdentifierType,omitempty"`
	StaticCode          bool        `plist:"StaticCode,omitempty"`
}

type KTCCServiceSystemPolicyDownloadsFolder struct {
	Allowed             bool        `plist:"Allowed,omitempty"`
	CodeRequirement     string      `plist:"CodeRequirement,omitempty"`
	CodeRequirementData interface{} `plist:"CodeRequirementData,omitempty"`
	Comment             string      `plist:"Comment,omitempty"`
	Identifier          string      `plist:"Identifier,omitempty"`
	IdentifierType      string      `plist:"IdentifierType,omitempty"`
	StaticCode          bool        `plist:"StaticCode,omitempty"`
}

type KTCCServiceSystemPolicyAllFiles struct {
	Allowed             bool        `plist:"Allowed,omitempty"`
	CodeRequirement     string      `plist:"CodeRequirement,omitempty"`
	CodeRequirementData interface{} `plist:"CodeRequirementData,omitempty"`
	Comment             string      `plist:"Comment,omitempty"`
	Identifier          string      `plist:"Identifier,omitempty"`
	IdentifierType      string      `plist:"IdentifierType,omitempty"`
	StaticCode          bool        `plist:"StaticCode,omitempty"`
}

type ComAppleSystemevents struct {
	AEReceiverCodeRequirement     string      `plist:"AEReceiverCodeRequirement, omitempty"`
	AEReceiverCodeRequirementData interface{} `plist:"AEReceiverCodeRequirementData, omitempty"`
	AEReceiverIdentifier          string      `plist:"AEReceiverIdentifier, omitempty"`
	AEReceiverIdentifierType      string      `plist:"AEReceiverIdentifierType, omitempty"`
	Allowed                       bool        `plist:"Allowed, omitempty"`
	CodeRequirement               string      `plist:"CodeRequirement, omitempty"`
	CodeRequirementData           interface{} `plist:"CodeRequirementData, omitempty"`
	Identifier                    string      `plist:"Identifier, omitempty"`
	IdentifierType                string      `plist:"IdentifierType, omitempty"`
	StaticCode                    bool        `plist:"StaticCode, omitempty"`
}

type KTCCServiceAppleEvents struct {
	ComAppleSystemevents ComAppleSystemevents `plist:"com.apple.systemevents"`
}

type Access struct {
	KTCCServiceAppleEvents                 KTCCServiceAppleEvents                 `plist:"kTCCServiceAppleEvents, omitempty"`
	KTCCServiceAccessibility               KTCCServiceAccessibility               `plist:"kTCCServiceAccessibility, omitempty"`
	KTCCServiceAddressBook                 KTCCServiceAddressBook                 `plist:"kTCCServiceAddressBook, omitempty"`
	KTCCServiceCalendar                    KTCCServiceCalendar                    `plist:"kTCCServiceCalendar, omitempty"`
	KTCCServiceFileProviderPresence        KTCCServiceFileProviderPresence        `plist:"kTCCServiceFileProviderPresence, omitempty"`
	KTCCServicePhotos                      KTCCServicePhotos                      `plist:"kTCCServicePhotos, omitempty"`
	KTCCServicePostEvent                   KTCCServicePostEvent                   `plist:"kTCCServicePostEvent, omitempty"`
	KTCCServiceSystemPolicyAllFiles        KTCCServiceSystemPolicyAllFiles        `plist:"kTCCServiceSystemPolicyAllFiles, omitempty"`
	KTCCServiceSystemPolicyDesktopFolder   KTCCServiceSystemPolicyDesktopFolder   `plist:"kTCCServiceSystemPolicyDesktopFolder, omitempty"`
	KTCCServiceSystemPolicyDocumentsFolder KTCCServiceSystemPolicyDocumentsFolder `plist:"kTCCServiceSystemPolicyDocumentsFolder, omitempty"`
	KTCCServiceSystemPolicyDownloadsFolder KTCCServiceSystemPolicyDownloadsFolder `plist:"kTCCServiceSystemPolicyDownloadsFolder, omitempty"`
}

// MDMOverrides is a map of MDM overrides
type MDMOverrides map[string]Access

// MDMEntry is a struct for the response from the MDMOverrides.plist
type MDMEntry struct {
	Source              string
	MDMServer           string
	CodeRequirement     string
	Identifier          string
	IdentifierType      string
	ServiceFriendlyName string
	Services            []ServiceDetail
	StaticCode          string
}

// CheckMDMOverrides checks for MDM overrides
func CheckMDMOverrides() (MDMOverrides, error) {
	_, err := os.Stat(mdmOverrides)
	if err != nil {
		return nil, fmt.Errorf("no MDM TCC Profiles found %v", err)
	}

	file, err := os.Open(mdmOverrides)
	if err != nil {
		return nil, fmt.Errorf("error opening plist file: %v", err)
	}
	defer file.Close()

	var mdmOverrides map[string]Access

	decoder := plist.NewDecoder(file)
	err = decoder.Decode(&mdmOverrides)
	if err != nil {
		return nil, fmt.Errorf("error decoding plist: %v", err)
	}

	return mdmOverrides, nil
}

func getMDMServer() string {
	cmd := exec.Command("profiles", "status", "-type", "enrollment")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error running the command: %s\n", err)
		return ""
	}

	lines := strings.Split(string(output), "\n")

	mdm := ""
	if len(lines) == 0 {
		return mdm
	}
	for _, line := range lines {
		if strings.Contains(line, "MDM server") {
			mdm = strings.TrimSpace(strings.Split(line, ": ")[1])
		}
	}

	return mdm
}

// ServiceDetail is a struct for formatting the response from the MDMOverrides.plist
type ServiceDetail struct {
	Service      string
	FriendlyName string
	Allowed      bool
}

func (m Access) services() []ServiceDetail {
	return []ServiceDetail{
		{
			Service:      "kTCCServiceAccessibility",
			FriendlyName: ServiceArray["kTCCServiceAccessibility"],
			Allowed:      m.KTCCServiceAccessibility.Allowed,
		},
		{
			Service:      "kTCCServiceAddressBook",
			FriendlyName: ServiceArray["kTCCServiceAddressBook"],
			Allowed:      m.KTCCServiceAddressBook.Allowed,
		},
		{
			Service:      "kTCCServiceCalendar",
			FriendlyName: ServiceArray["kTCCServiceCalendar"],
			Allowed:      m.KTCCServiceCalendar.Allowed,
		},
		{
			Service:      "kTCCServiceFileProviderPresence",
			FriendlyName: ServiceArray["kTCCServiceFileProviderPresence"],
			Allowed:      m.KTCCServiceFileProviderPresence.Allowed,
		},
		{
			Service:      "kTCCServicePhotos",
			FriendlyName: ServiceArray["kTCCServicePhotos"],
			Allowed:      m.KTCCServicePhotos.Allowed,
		},
		{
			Service:      "kTCCServicePostEvent",
			FriendlyName: ServiceArray["kTCCServicePostEvent"],
			Allowed:      m.KTCCServicePostEvent.Allowed,
		},
		{
			Service:      "kTCCServiceSystemPolicyAllFiles",
			FriendlyName: ServiceArray["kTCCServiceSystemPolicyAllFiles"],
			Allowed:      m.KTCCServiceSystemPolicyAllFiles.Allowed,
		},
		{
			Service:      "kTCCServiceSystemPolicyDesktopFolder",
			FriendlyName: ServiceArray["kTCCServiceSystemPolicyDesktopFolder"],
			Allowed:      m.KTCCServiceSystemPolicyDesktopFolder.Allowed,
		},
		{
			Service:      "kTCCServiceSystemPolicyDocumentsFolder",
			FriendlyName: ServiceArray["kTCCServiceSystemPolicyDocumentsFolder"],
			Allowed:      m.KTCCServiceSystemPolicyDocumentsFolder.Allowed,
		},
		{
			Service:      "kTCCServiceSystemPolicyDownloadsFolder",
			FriendlyName: ServiceArray["kTCCServiceSystemPolicyDownloadsFolder"],
			Allowed:      m.KTCCServiceSystemPolicyDownloadsFolder.Allowed,
		},
	}
}

func (m Access) codeRequirement() string {
	codeRequirement := []string{
		m.KTCCServiceAccessibility.CodeRequirement,
		m.KTCCServiceAddressBook.CodeRequirement,
		m.KTCCServiceCalendar.CodeRequirement,
		m.KTCCServiceFileProviderPresence.CodeRequirement,
		m.KTCCServicePhotos.CodeRequirement,
		m.KTCCServicePostEvent.CodeRequirement,
		m.KTCCServiceSystemPolicyAllFiles.CodeRequirement,
		m.KTCCServiceSystemPolicyDesktopFolder.CodeRequirement,
		m.KTCCServiceSystemPolicyDocumentsFolder.CodeRequirement,
		m.KTCCServiceSystemPolicyDownloadsFolder.CodeRequirement,
	}

	return sort(codeRequirement)
}

func (m Access) identifier() string {
	identifier := []string{
		m.KTCCServiceAccessibility.Identifier,
		m.KTCCServiceAddressBook.Identifier,
		m.KTCCServiceCalendar.Identifier,
		m.KTCCServiceFileProviderPresence.Identifier,
		m.KTCCServicePhotos.Identifier,
		m.KTCCServicePostEvent.Identifier,
		m.KTCCServiceSystemPolicyAllFiles.Identifier,
		m.KTCCServiceSystemPolicyDesktopFolder.Identifier,
		m.KTCCServiceSystemPolicyDocumentsFolder.Identifier,
		m.KTCCServiceSystemPolicyDownloadsFolder.Identifier,
	}

	return sort(identifier)
}

func (m Access) identifierType() string {
	identifierType := []string{
		m.KTCCServiceAccessibility.IdentifierType,
		m.KTCCServiceAddressBook.IdentifierType,
		m.KTCCServiceCalendar.IdentifierType,
		m.KTCCServiceFileProviderPresence.IdentifierType,
		m.KTCCServicePhotos.IdentifierType,
		m.KTCCServicePostEvent.IdentifierType,
		m.KTCCServiceSystemPolicyAllFiles.IdentifierType,
		m.KTCCServiceSystemPolicyDesktopFolder.IdentifierType,
		m.KTCCServiceSystemPolicyDocumentsFolder.IdentifierType,
		m.KTCCServiceSystemPolicyDownloadsFolder.IdentifierType,
	}

	return sort(identifierType)
}

func (m Access) static(service string) string {
	var required []string

	if m.KTCCServiceAccessibility.StaticCode {
		required = append(required, ServiceArray["kTCCServiceAccessibility"]+" Approved")
	} else {
		required = append(required, ServiceArray["kTCCServiceAccessibility"]+" Not Approved")
	}
	if m.KTCCServiceAddressBook.StaticCode {
		required = append(required, ServiceArray["kTCCServiceAddressBook"]+" Approved")
	} else {
		required = append(required, ServiceArray["kTCCServiceAddressBook"]+" Not Approved")
	}
	if m.KTCCServiceCalendar.StaticCode {
		required = append(required, ServiceArray["kTCCServiceCalendar"]+" Approved")
	} else {
		required = append(required, ServiceArray["kTCCServiceCalendar"]+" Not Approved")
	}
	if m.KTCCServiceFileProviderPresence.StaticCode {
		required = append(required, ServiceArray["kTCCServiceFileProviderPresence"]+" Approved")
	} else {
		required = append(required, ServiceArray["kTCCServiceFileProviderPresence"]+" Not Approved")
	}
	if m.KTCCServicePhotos.StaticCode {
		required = append(required, ServiceArray["kTCCServicePhotos"]+" Approved")
	} else {
		required = append(required, ServiceArray["kTCCServicePhotos"]+" Not Approved")
	}
	if m.KTCCServicePostEvent.StaticCode {
		required = append(required, ServiceArray["kTCCServicePostEvent"]+" Approved")
	} else {
		required = append(required, ServiceArray["kTCCServicePostEvent"]+" Not Approved")
	}
	if m.KTCCServiceSystemPolicyAllFiles.StaticCode {
		required = append(required, ServiceArray["kTCCServiceSystemPolicyAllFiles"]+" Approved")
	} else {
		required = append(required, ServiceArray["kTCCServiceSystemPolicyAllFiles"]+" Not Approved")
	}
	if m.KTCCServiceSystemPolicyDesktopFolder.StaticCode {
		required = append(required, ServiceArray["kTCCServiceSystemPolicyDesktopFolder"]+" Approved")
	} else {
		required = append(required, ServiceArray["kTCCServiceSystemPolicyDesktopFolder"]+" Not Approved")
	}
	if m.KTCCServiceSystemPolicyDocumentsFolder.StaticCode {
		required = append(required, ServiceArray["kTCCServiceSystemPolicyDocumentsFolder"]+" Approved")
	} else {
		required = append(required, ServiceArray["kTCCServiceSystemPolicyDocumentsFolder"]+" Not Approved")
	}
	if m.KTCCServiceSystemPolicyDownloadsFolder.StaticCode {
		required = append(required, ServiceArray["kTCCServiceSystemPolicyDownloadsFolder"]+" Approved")
	} else {
		required = append(required, ServiceArray["kTCCServiceSystemPolicyDownloadsFolder"]+" Not Approved")
	}

	return sort(required)
}

// ProcessMDMOverrides processes the MDMOverrides.plist
func ProcessMDMOverrides(service string, m Access) *MDMEntry {
	return &MDMEntry{
		Source:              "MDM",
		MDMServer:           getMDMServer(),
		CodeRequirement:     m.codeRequirement(),
		Identifier:          m.identifier(),
		IdentifierType:      m.identifierType(),
		ServiceFriendlyName: getAppNameFromClient(service),
		Services:            m.services(),
		StaticCode:          m.static(service),
	}
}
