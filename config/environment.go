package config

type APIEnvironment struct {
	HTTPPort    string `envconfig:"PORT" default:"8080"`
	DatabaseURL string `envconfig:"DATABASE_URL" required:"true"`

	IsAdmin bool `envconfig:"ADMIN" default:"false"`

	DocsDir      string `envconfig:"DOCS_DIR" default:"/docs"`
	DocsUser     string `envconfig:"DOCS_USER" default:"go_template"`
	DocsPassword string `envconfig:"DOCS_PASSWORD"`
}
