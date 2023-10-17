package tcc

// AuthValueArray is a map of auth values to friendly names
var AuthValueArray = map[int]string{
	0: "Denied",
	1: "Unknown",
	2: "Allowed",
	3: "Limited",
}

// AuthReasonArray is a map of auth reasons to friendly names
var AuthReasonArray = map[int]string{
	0:  "Inherited/Unknown",
	1:  "Error",
	2:  "User Consent",
	3:  "User Set",
	4:  "System Set",
	5:  "Service Policy",
	6:  "MDM Policy",
	7:  "Override Policy",
	8:  "Missing usage string",
	9:  "Prompt Timeout",
	10: "Preflight Unknown",
	11: "Entitled",
	12: "App Type Policy",
}

// ServiceArray is a map of service names to friendly names
var ServiceArray = map[string]string{
	"kTCCServiceAddressBook":                  "Contacts",
	"kTCCServiceAppleEvents":                  "Apple Events",
	"kTCCServiceBluetoothAlways":              "Bluetooth",
	"kTCCServiceCalendar":                     "Calendar",
	"kTCCServiceCamera":                       "Camera",
	"kTCCServiceContactsFull":                 "Full contacts information",
	"kTCCServiceContactsLimited":              "Basic contacts information",
	"kTCCServiceFileProviderDomain":           "Files managed by Apple Events",
	"kTCCServiceFileProviderPresence":         "See when files managed by the client are in use",
	"kTCCServiceLocation":                     "Current location",
	"kTCCServiceMediaLibrary":                 "Apple Music, music and video activity, and media library",
	"kTCCServiceMicrophone":                   "Microphone",
	"kTCCServiceMotion":                       "Motion & Fitness Activity",
	"kTCCServicePhotos":                       "Read Photos",
	"kTCCServicePhotosAdd":                    "Add to Photos",
	"kTCCServicePrototype3Rights":             "Authorization Test Service Proto3Right",
	"kTCCServicePrototype4Rights":             "Authorization Test Service Proto4Right",
	"kTCCServiceReminders":                    "Reminders",
	"kTCCServiceScreenCapture":                "Capture screen contents",
	"kTCCServiceSiri":                         "Use Siri",
	"kTCCServiceSpeechRecognition":            "Speech Recognition",
	"kTCCServiceSystemPolicyDesktopFolder":    "Desktop folder",
	"kTCCServiceSystemPolicyDeveloperFiles":   "Files in Software Development",
	"kTCCServiceSystemPolicyDocumentsFolder":  "Files in Documents folder",
	"kTCCServiceSystemPolicyDownloadsFolder":  "Files in Downloads folder",
	"kTCCServiceSystemPolicyNetworkVolumes":   "Files on a network volume",
	"kTCCServiceSystemPolicyRemovableVolumes": "Files on a removable volume",
	"kTCCServiceSystemPolicySysAdminFiles":    "Administer the computer",
	"kTCCServiceWillow":                       "Home data",
	"kTCCServiceSystemPolicyAllFiles":         "Full Disk Access",
	"kTCCServiceAccessibility":                "Control the computer",
	"kTCCServicePostEvent":                    "Send keystrokes",
	"kTCCServiceListenEvent":                  "Monitor input from the keyboard",
	"kTCCServiceDeveloperTool":                "Run insecure software locally",
	"kTCCServiceLiverpool":                    "Location services",
	"kTCCServiceUbiquity":                     "iCloud",
	"kTCCServiceShareKit":                     "Share features",
	"kTCCServiceLinkedIn":                     "Share via LinkedIn",
	"kTCCServiceTwitter":                      "Share via Twitter",
	"kTCCServiceFacebook":                     "Share via Facebook",
	"kTCCServiceSinaWeibo":                    "Share via Sina Weibo",
	"kTCCServiceTencentWeibo":                 "Share via Tencent Weibo",
}

// TCCEntry is a struct that represents a row in the TCC.db database
type TCCEntry struct {
	Source              string
	Username            string
	Client              string
	ClientID            string
	ServiceName         string
	ServiceFriendlyName string
	AuthValue           string
	AuthReason          string
	Timestamp           int64
	FormattedTime       string
	CodeSignReq         string
}

// ProcessRow takes a row from the TCC.db database and returns a TCCEntry struct
// formatting the input as needed.
func ProcessRow(source string, sqr SQResponse) *TCCEntry {
	return &TCCEntry{
		Source:              sourcerer(source), // System or User
		Username:            source,
		Client:              sqr.Client,
		ClientID:            getAppNameFromClient(sqr.Client),
		ServiceName:         sqr.Service,
		ServiceFriendlyName: ServiceArray[sqr.Service],
		AuthValue:           AuthValueArray[sqr.AuthValue],
		AuthReason:          AuthReasonArray[sqr.AuthReason],
		Timestamp:           sqr.LastMod,
		FormattedTime:       epochToTime(sqr.LastMod),
		CodeSignReq:         decodeCodeSignature(sqr.Csreq),
	}
}
