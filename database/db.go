package database

import (
	"log"

	"github.com/SkyVillageMc/skyvillage-api/db"
)

var DB *db.PrismaClient

func Init() {
	DB = db.NewClient()

	if err := DB.Prisma.Connect(); err != nil {
		//This shouldn't happen
		log.Fatalf("Error connecting to the db\n%s\n", err.Error())
	}
}
