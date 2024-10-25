package erika

type StringOrFalse string

func (s *StringOrFalse) UnmarshalJSON(data []byte) error {
	asString := string(data)
	if asString == "false" {
		*s = ""
	} else {
		*s = StringOrFalse(asString)
	}
	return nil
}

// Secret Key types
type SecretKey struct {
	Version        string
	AuthKey        string
	SecureAuthKey  string
	LoggedInKey    string
	NonceKey       string
	AuthSalt       string
	SecureAuthSalt string
	LoggedInSalt   string
	NonceSalt      string
	Raw            []byte
}

// Plugin metadata types
type contributor struct {
	Avatar      string `json:"avatar"`
	DisplayName string `json:"display_name"`
	Profile     string `json:"profile"`
}

type screenshot struct {
	Caption string `json:"caption"`
	Src     string `json:"src"`
}

type PluginMetadata struct {
	ActiveInstalls int    `json:"active_installs"`
	Added          string `json:"added"`
	Author         string `json:"author"`
	AuthorProfile  string `json:"author_profile"`
	Banners        struct {
		High string `json:"high"`
		Low  string `json:"low"`
	} `json:"banners"`
	BusinessModel        bool                   `json:"business_model"`
	Closed               bool                   `json:"closed"`
	ClosedDate           string                 `json:"closed_date"`
	CommercialSupportURL string                 `json:"commercial_support_url"`
	Contributors         map[string]contributor `json:"contributors"`
	Description          string                 `json:"description"`
	DonateLink           string                 `json:"donate_link"`
	DownloadLink         string                 `json:"download_link"`
	Error                string                 `json:"error"`
	Homepage             string                 `json:"homepage"`
	LastUpdated          string                 `json:"last_updated"`
	Name                 string                 `json:"name"`
	NumRatings           int                    `json:"num_ratings"`
	PreviewLink          string                 `json:"preview_link"`
	Rating               int                    `json:"rating"`
	Ratings              struct {
		Num1 int `json:"1"`
		Num2 int `json:"2"`
		Num3 int `json:"3"`
		Num4 int `json:"4"`
		Num5 int `json:"5"`
	} `json:"ratings"`
	Reason          string                `json:"reason"`
	ReasonText      string                `json:"reason_text"`
	RepositoryURL   string                `json:"repository_url"`
	Requires        string                `json:"requires"`
	RequiresPhp     StringOrFalse         `json:"requires_php"`
	RequiresPlugins []interface{}         `json:"requires_plugins"`
	Screenshots     map[string]screenshot `json:"screenshots"`
	Sections        struct {
		Changelog   string `json:"changelog"`
		Description string `json:"description"`
		Reviews     string `json:"reviews"`
		Screenshots string `json:"screenshots"`
	} `json:"sections"`
	Slug                   string            `json:"slug"`
	SupportThreads         int               `json:"support_threads"`
	SupportThreadsResolved int               `json:"support_threads_resolved"`
	SupportURL             string            `json:"support_url"`
	Tags                   map[string]string `json:"tags"`
	Tested                 string            `json:"tested"`
	UpgradeNotice          []interface{}     `json:"upgrade_notice"`
	Version                string            `json:"version"`
	Versions               map[string]string `json:"versions"`
}

// Credits
type creditsGroup struct {
	Name    StringOrFalse       `json:"name"`
	Type    string              `json:"type"`
	Shuffle bool                `json:"shuffle"`
	Data    map[string][]string `json:"data"`
}

type Credits struct {
	Groups map[string]creditsGroup `json:"groups"`
	Data   struct {
		Profiles string `json:"profiles"`
		Version  string `json:"version"`
	} `json:"data"`
}

// Search

type SearchInfo struct {
	Page    int `json:"page"`
	Pages   int `json:"pages"`
	Results int `json:"results"`
}

type PluginSearch struct {
	Info    SearchInfo       `json:"info"`
	Plugins []PluginMetadata `json:"plugins"`
}
