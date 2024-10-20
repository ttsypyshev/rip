package backend

import (
	"fmt"
	"rip/database"
	"strings"
	"time"

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

// Получение файлов для проекта
func (app *App) GetFilesForProject(projectID int) ([]DbFile, error) {
	var matchedFiles []DbFile
	result := app.db.db.Where("id_project = ?", projectID).Find(&matchedFiles)
	if result.Error != nil {
		return nil, result.Error
	}

	return matchedFiles, nil
}

// Фильтрация языков по запросу
func (app *App) FilterLangsByQuery(query string) ([]DbLang, error) {
	var filteredLangs []DbLang
	lowerQuery := "%" + strings.ToLower(query) + "%"

	result := app.db.db.Where("LOWER(name) LIKE ?", lowerQuery).Find(&filteredLangs)
	if result.Error != nil {
		return nil, result.Error
	}

	return filteredLangs, nil
}

// Поиск последнего черновика для пользователя
func findLastDraft(app *App, userID int) (int, error) {
	var lastProject DbProject

	if err := app.db.db.Where("status = ? AND id_user = ?", 0, userID).First(&lastProject).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return -1, nil
		}
		return -1, err
	}
	return lastProject.ID, nil
}

// Создание нового черновика или возврат существующего
func createDraft(app *App, userID int) (int, error) {
	projectID, err := findLastDraft(app, userID)
	if err != nil {
		return -1, err
	} else if projectID == -1 {
		newProject := DbProject{
			IDUser:       userID,
			CreationTime: time.Now(),
			Status:       0,
		}

		if err := app.db.db.Create(&newProject).Error; err != nil {
			return -1, err
		}

		return newProject.ID, nil
	}
	return projectID, nil
}

// Добавление файла к проекту для пользователя
func (app *App) AddFile(projectID, langID, userID int) error {
	newFile := DbFile{
		IDLang:    langID,
		IDProject: projectID,
	}

	var existingFile DbFile
	if err := app.db.db.Where("id_lang = ? AND id_project = ?", newFile.IDLang, newFile.IDProject).First(&existingFile).Error; err == nil {
		return fmt.Errorf("file with IDLang %d and IDProject %d already exists", newFile.IDLang, newFile.IDProject)
	} else if err != gorm.ErrRecordNotFound {
		return err
	}

	if err := app.db.db.Create(&newFile).Error; err != nil {
		return err
	}

	return nil
}

// Обновление статуса проекта
func (app *App) UpdateProjectStatus(projectID int, newStatus int) error {
	query := "UPDATE projects SET status = ? WHERE id = ?"

	result := app.db.db.Exec(query, newStatus, projectID)
	if result.Error != nil {
		return fmt.Errorf("failed to update project status: %w", result.Error)
	}

	return nil
}

// Обновление кода файлов по предоставленным мапам идентификаторов и кода
func (app *App) UpdateFilesCode(idToCodeMap map[int]string) error {
	for id, newCode := range idToCodeMap {
		var file DbFile
		if err := app.db.db.Where("id = ?", id).First(&file).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return fmt.Errorf("file with id %d not found", id)
			}
			return err
		}

		file.Code = newCode

		if err := app.db.db.Save(&file).Error; err != nil {
			return fmt.Errorf("failed to update file with id %d: %v", id, err)
		}
	}
	return nil
}

// Подсчет количества файлов в черновике пользователя
func (app *App) CountFiles(userID int) (int64, error) {
	projectID, err := findLastDraft(app, userID)
	if err != nil {
		return 0, err
	}

	var count int64
	if err := app.db.db.Model(&DbFile{}).Where("id_project = ?", projectID).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
