package database

type Guild struct {
	GuildID uint64 `gorm:"column:GUILDID;unique;primary_key"`
	Prefix  string `gorm:"column:PREFIX;type:varchar(8)"`
}

func SetPrefix(guild uint64, prefix string) {
	Db.Where(Guild{GuildID: guild}).Assign(Guild{Prefix: prefix}).FirstOrCreate(nil)
}

func GetPrefix(guild uint64, ch chan string) {
	node := Guild{
		Prefix: "l!",
	}
	Db.Where(Guild{GuildID: guild}).FirstOrCreate(&node)
	ch <- node.Prefix
}
