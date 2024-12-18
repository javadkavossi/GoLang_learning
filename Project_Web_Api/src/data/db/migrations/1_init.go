package migrations

import (
	"github.com/javadkavossi/GoLang_learning/config"
	"github.com/javadkavossi/GoLang_learning/constants"
	"github.com/javadkavossi/GoLang_learning/data/db"
	"github.com/javadkavossi/GoLang_learning/data/models"
	"github.com/javadkavossi/GoLang_learning/pkg/logging"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var logger = logging.NewLogger(config.GetConfig())

func Up_1() {
	database := db.GetDb()
	CreateTables(database)

}

func CreateTables(database *gorm.DB) {

	tables := []interface{}{}
	country := models.Country{}
	city := models.City{}
	user := models.User{}
	role := models.Role{}
	userRole := models.UserRole{}
	// --------------------- Add Table ------------------------------------
	tables = addNewTable(database, tables, country)
	tables = addNewTable(database, tables, city)
	tables = addNewTable(database, tables, user)
	tables = addNewTable(database, tables, role)
	tables = addNewTable(database, tables, userRole)

	// -------------------- Create Table ----------------------------------

	database.Migrator().CreateTable(tables...)

	// -------------------- Create Default Information ----------------------------------
	createDefaultInformation(database)

	logger.Info(logging.Postgres, logging.Migration, "tables created", nil)
}

// $ ========================================= Add Table Connection ======================================================
func addNewTable(database *gorm.DB, tables []interface{}, model interface{}) []interface{} {
	if !database.Migrator().HasTable(model) {
		tables = append(tables, model)
	}
	return tables
}

// $============================================ Create Default Information ==============================================
func createDefaultInformation(database *gorm.DB) {

	adminRole := models.Role{Name: constants.AdminRoleName}
	createRoleIfNotExists(database, &adminRole)

	defaultRole := models.Role{Name: constants.DefaultRoleName}
	createRoleIfNotExists(database, &defaultRole)

	user := models.User{
		UserName:   constants.DefaultUserName,
		FirstName:  "test Name",
		LastName:   "test name",
		Email:      "test@example.com",
		MobileName: "09919999999",
		Password:   "test",
	}

	pass := "12345678"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	createAdminIfNotExists(database, &user, adminRole.Id)

}

func createRoleIfNotExists(database *gorm.DB, r *models.Role) {
	exists := 0
	database.Model(&models.Role{}).Select("1").Where("name = ?", r.Name).First(&exists)
	if exists == 0 {
		database.Create(r)
	}
}

func createAdminIfNotExists(database *gorm.DB, u *models.User, roleId int) {
	exists := 0
	database.Model(&models.User{}).Select("1").Where("name = ?", u.UserName).First(&exists)
	if exists == 0 {
		database.Create(u)
		ur := models.UserRole{UserId: u.Id, RoleId: roleId}
		database.Create(&ur)
	}
}

func Down_1() {}
