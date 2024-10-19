package backend

import (
	"rip/database"
	"strings"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type (
	DbLang    = database.Lang
	DbProject = database.Project
	DbFile    = database.File
)

type Db struct {
	db *gorm.DB
}

func Migrate() error {
	_ = godotenv.Load()
	db, err := gorm.Open(postgres.Open(FromEnv()), &gorm.Config{})
	if err != nil {
		return err
	}

	// Migrate the schema
	err = db.AutoMigrate(&DbLang{}, &DbProject{}, &DbFile{})
	if err != nil {
		return err
	}

	return nil
}

func NewDB(dsn string) (*App, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &App{
		db: &Db{db},
	}, nil
}

// Получение всех сущностей
func getAll[T any](app *App) ([]T, error) {
	var items []T

	err := app.db.db.Find(&items).Error
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (app *App) GetLangs() ([]DbLang, error) {
	return getAll[DbLang](app)
}

func (app *App) GetProjects() ([]DbProject, error) {
	return getAll[DbProject](app)
}

func (app *App) GetFiles() ([]DbFile, error) {
	return getAll[DbFile](app)
}

// Получение сущностей по ID
func getByID[T any](app *App, id int) (T, error) {
	var item T

	err := app.db.db.First(&item, "id = ?", id).Error
	if err != nil {
		return item, err
	}

	return item, nil
}

func (app *App) GetLangByID(langID int) (DbLang, error) {
	return getByID[DbLang](app, langID)
}

func (app *App) GetProjectByID(projectID int) (DbProject, error) {
	return getByID[DbProject](app, projectID)
}

func (app *App) GetFileByID(fileID int) (DbFile, error) {
	return getByID[DbFile](app, fileID)
}

// func (a *app) GetProductByID(id int) (*Project, error) {
// 	product := &Project{}
// 	err := a.db.First(product, "id = ?", id).Error // find product with id = 1
// 	if err != nil {
// 		return nil, err
// 	}
// 	return product, nil
// }

// Получение файлов для проекта с использованием GORM
func (app *App) GetFilesForProject(projectID int) ([]DbFile, error) {
	var matchedFiles []DbFile
	result := app.db.db.Where("id_project = ?", projectID).Find(&matchedFiles)
	if result.Error != nil {
		return nil, result.Error
	}

	return matchedFiles, nil
}

// Фильтрация языков по запросу с использованием GORM
func (app *App) FilterLangsByQuery(query string) ([]DbLang, error) {
	var filteredLangs []DbLang
	lowerQuery := "%" + strings.ToLower(query) + "%"

	result := app.db.db.Where("LOWER(name) LIKE ?", lowerQuery).Find(&filteredLangs)
	if result.Error != nil {
		return nil, result.Error
	}

	return filteredLangs, nil
}

// func (a *app) CreateProduct(product Project) error {
// 	return a.db.Create(product).Error
// }

// func FindMaxProjectID() int {
// 	maxID := -1
// 	for _, project := range Projects {
// 		if project.ID > maxID {
// 			maxID = project.ID
// 		}
// 	}
// 	return maxID
// }

// func GetLangsForProject(matchedFiles []File, projectID int) []Lang {
// 	var matchedLangs []Lang
// 	for _, file := range matchedFiles {
// 		log.Println(file)
// 		if file.ID_lang < 0 || file.ID_lang >= len(Langs) {
// 			log.Printf("Invalid ID_lang: %d for file %v", file.ID_lang, file)
// 			continue
// 		}
// 		matchedLangs = append(matchedLangs, Langs[file.ID_lang])
// 	}
// 	return matchedLangs
// }
