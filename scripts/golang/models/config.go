package models

type Service struct {
	Name        string  `yaml:"name"`
	URL         string  `yaml:"url"`
	PrefixToAdd *string `yaml:"prefixToAdd"`
}
type Info struct {
	Title       string  `yaml:"title"`
	Version     string  `yaml:"version"`
	Description string  `yaml:"description"`
	Contact     Contact `yaml:"contact"`
}

type Contact struct {
	Name  string `yaml:"name"`
	Email string `yaml:"email"`
}

type Server struct {
	URL         string `yaml:"url"`
	Description string `yaml:"description"`
}

type Config struct {
	Services        []Service              `yaml:"services"`
	Info            Info                   `yaml:"info"`
	Servers         []Server               `yaml:"servers"`
	SecuritySchemes map[string]interface{} `yaml:"securitySchemes"`
}
