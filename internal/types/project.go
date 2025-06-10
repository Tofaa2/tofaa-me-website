package types

type Project struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	ImageURL    string   `json:"image_url"`
	Tags        []string `json:"tags"`
	GitHubURL   string   `json:"github_url"`
	ProjectURL  string   `json:"project_url,omitempty"`
}

// GetProjects returns a list of all projects
func GetProjects() []Project {
	return []Project{
		{
			Name:        "EntityLib",
			Description: "A packetevents utility library that provides an abstraction layer over raw packets for Entities and their subtypes.",
			ImageURL:    "/static/images/entitylib.jpg",
			Tags:        []string{"Java"},
			GitHubURL:   "https://github.com/Tofaa2/EntityLib",
		},
		{
			Name:        "Tachyon",
			Description: "A fast, lightweight, multithreaded Minecraft server implementation.",
			ImageURL:    "/static/images/tachyon.jpg",
			Tags:        []string{"Java", "C#", "Dotnet"},
			GitHubURL:   "https://github.com/Tofaa2/Tachyon",
		},
	}
}
