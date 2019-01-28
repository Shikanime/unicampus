package education

type School struct {
	UUID        string
	Name        string
	Description string
	Phone       string
	Email       string
	Links       []Link
	Pictures    []Link
	Locations   []Location
	Sectors     []Sector
}

type Sector struct {
	Name string
}

type Student struct {
	UUID      string
	FirstName string
	LastName  string
	Email     string
	Phone     string
}

var (
	schools = []School{
		School{
			UUID:        "xxx",
			Name:        "ETNA",
			Description: "C'est une ecole d'informatique",
			Phone:       "01.44.08.00.23",
			Email:       "admissions@etna.io",
			Links: []Link{
				Link{
					Type:      "Website",
					Reference: "https://etna.io/",
				},
				Link{
					Type:      "Facebook",
					Reference: "https://www.facebook.com/ecole.etna",
				},
				Link{
					Type:      "Twitter",
					Reference: "https://twitter.com/etna_io",
				},
				Link{
					Type:      "Linkedin",
					Reference: "https://www.linkedin.com/school/etna-%C3%A9cole-d'alternance-en-informatique/",
				},
				Link{
					Type:      "Youtube",
					Reference: "https://www.youtube.com/user/liveETNA",
				},
				Link{
					Type:      "RSS",
					Reference: "https://feeds.feedburner.com/blogetna",
				},
			},
		},
	}
)
