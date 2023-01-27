package database

import (
	"waysgalerry_be/models"
	// postgresql "waysgalerry_be/pkg/postgresql"
	"fmt"
	mysql "waysgalerry_be/pkg/mysql"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.User{},
		&models.Art{},
		&models.Post{},
		&models.PostImage{},
		&models.Hired{},
		&models.Project{},
		&models.ProjectImage{},
		&models.Follow{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
