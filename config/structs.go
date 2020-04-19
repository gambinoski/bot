package config

type (
	// Config is the global config object
	Config struct {
		Bot Bot
		Db  Database
	}

	// Bot is the bot object
	Bot struct {
		Token      string
		Prefix     string
		OwnerID    string
		LogChannel string
	}

	// Database is the database object
	Database struct {
		Host     string
		Port     string
		User     string
		Name     string
		Password string
	}
)
