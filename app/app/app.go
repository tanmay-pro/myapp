package app

import (
	"myapp/util/logger"
)

type App struct {
	logger *logger.Logger
}

func New(logger *logger.Logger) *App {
	return &App{logger: logger}
}

// package app

// import (
// 	"myapp/util/logger"

// 	"github.com/jinzhu/gorm"
// )

// type App struct {
// 	logger *logger.Logger
// 	db     *gorm.DB
// }

// func New(
// 	logger *logger.Logger,
// 	db *gorm.DB,
// ) *App {
// 	return &App{
// 		logger: logger,
// 		db:     db,
// 	}
// }

func (app *App) Logger() *logger.Logger {
	return app.logger
}
