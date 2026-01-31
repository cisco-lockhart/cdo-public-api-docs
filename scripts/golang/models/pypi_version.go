package models

type info struct {
	Author                 string   `json:"author"`
	AuthorEmail            string   `json:"author_email"`
	BugtrackURL            *string  `json:"bugtrack_url"`
	Classifiers            []string `json:"classifiers"`
	Description            string   `json:"description"`
	DescriptionContentType string   `json:"description_content_type"`
	DocsURL                *string  `json:"docs_url"`
	DownloadURL            *string  `json:"download_url"`
	Downloads              struct {
		LastDay   int `json:"last_day"`
		LastMonth int `json:"last_month"`
		LastWeek  int `json:"last_week"`
	} `json:"downloads"`
	Dynamic           []string `json:"dynamic"`
	HomePage          *string  `json:"home_page"`
	Keywords          string   `json:"keywords"`
	License           *string  `json:"license"`
	LicenseExpression *string  `json:"license_expression"`
	LicenseFiles      *string  `json:"license_files"`
	Maintainer        *string  `json:"maintainer"`
	MaintainerEmail   *string  `json:"maintainer_email"`
	Name              string   `json:"name"`
	PackageURL        string   `json:"package_url"`
	Platform          *string  `json:"platform"`
	ProjectURL        string   `json:"project_url"`
	ProjectURLs       struct {
		Documentation *string `json:"Documentation"`
		Url           *string `json:"url"`
	}
	ProvidesExtra  *string  `json:"provides_extra"`
	ReleaseURL     string   `json:"release_url"`
	RequiresDist   []string `json:"requires_dist"`
	RequiresPython *string  `json:"requires_python"`
	Summary        string   `json:"summary"`
	Version        string   `json:"version"`
	Yanked         bool     `json:"yanked"`
	YankedReason   *string  `json:"yanked_reason"`
}

type PypiVersionInfo struct {
	Info info `json:"info"`
}
