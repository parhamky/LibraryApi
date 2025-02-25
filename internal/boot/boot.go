package boot

import (
	"LibraryApi/internal/cache"
	"LibraryApi/internal/config"
	"LibraryApi/internal/db"
	"log"
)

func BootServer() error {
	dbConfig := config.LoadDBConfig()

	err := db.ConnectDB(dbConfig)
	if err != nil {
		return err
	}

	log.Println("Database connected!")

	return nil
}

func BootCache() error {
	conf := config.LoadCacheConfig()

	err := cache.InitializeRedis(conf)

	if err != nil {
		return err
	}

	log.Println("Cache connected!")
	return nil
}
