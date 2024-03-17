![Go Clean AI](docs/img/logo.svg)

>>> Begin Filename=error.go, Path=/Users/mac/code/go/src/github.com/mytoolzone/clean-go/internal/controller/http/v1/error.go
```package v1

import (
	"github.com/gin-gonic/gin"
)

type response struct {
	Error string `json:"error" example:"message"`
}

func errorResponse(c *gin.Context, code int, msg string) {
	c.AbortWithStatusJSON(code, response{msg})
}
```
<<< End Filename=error.go
>>> Begin Filename=router.go, Path=/Users/mac/code/go/src/github.com/mytoolzone/clean-go/internal/controller/http/v1/router.go
```// Package v1 implements routing paths. Each services in own file.
package v1

import (
	"github.com/gin-gonic/gin"
	// Swagger docs.
	_ "go-clean/docs"
	"go-clean/internal/usecase"
	"go-clean/pkg/logger"
)

// NewRouter -.
// Swagger spec:
// @title       Go Clean Template API
// @description Using a translation service as an example
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
func NewRouter(handler *gin.Engine, l logger.Interface, t usecase.Translation) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Routers
	h := handler.Group("/v1")
	{
		newTranslationRoutes(h, t, l)
	}
}
```
<<< End Filename=router.go
>>> Begin Filename=translation.go, Path=/Users/mac/code/go/src/github.com/mytoolzone/clean-go/internal/controller/http/v1/translation.go
```package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-clean/internal/entity"
	"go-clean/internal/usecase"
	"go-clean/pkg/logger"
)

type translationRoutes struct {
	t usecase.Translation
	l logger.Interface
}

func newTranslationRoutes(handler *gin.RouterGroup, t usecase.Translation, l logger.Interface) {
	r := &translationRoutes{t, l}

	h := handler.Group("/translation")
	{
		h.GET("/history", r.history)
	}
}

type historyResponse struct {
	History []entity.Translation `json:"history"`
}

// @Summary     Show history
// @Description Show all translation history
// @ID          history
// @Tags  	    translation
// @Accept      json
// @Produce     json
// @Success     200 {object} historyResponse
// @Failure     500 {object} response
// @Router      /translation/history [get]
func (r *translationRoutes) history(c *gin.Context) {
	translations, err := r.t.History(c.Request.Context())
	if err != nil {
		r.l.Error(err, "http - v1 - history")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	c.JSON(http.StatusOK, historyResponse{translations})
}

type doTranslateRequest struct {
	Source      string `json:"source"       binding:"required"  example:"auto"`
	Destination string `json:"destination"  binding:"required"  example:"en"`
	Original    string `json:"original"     binding:"required"  example:"текст для перевода"`
}
```
<<< End Filename=translation.go
>>> Begin Filename=translation.go, Path=/Users/mac/code/go/src/github.com/mytoolzone/clean-go/internal/entity/translation.go
```// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

// Translation -.
type Translation struct {
	Source      string `json:"source"       example:"auto"`
	Destination string `json:"destination"  example:"en"`
	Original    string `json:"original"     example:"текст для перевода"`
	Translation string `json:"translation"  example:"text for translation"`
}
```
<<< End Filename=translation.go
>>> Begin Filename=interfaces.go, Path=/Users/mac/code/go/src/github.com/mytoolzone/clean-go/internal/usecase/interfaces.go
```// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"go-clean/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (
	// Translation -.
	Translation interface {
		History(context.Context) ([]entity.Translation, error)
	}

	// TranslationRepo -.
	TranslationRepo interface {
		GetHistory(context.Context) ([]entity.Translation, error)
	}
)
```
<<< End Filename=interfaces.go
>>> Begin Filename=translation_postgres.go, Path=/Users/mac/code/go/src/github.com/mytoolzone/clean-go/internal/usecase/repo/translation_postgres.go
```package repo

import (
	"context"
	"fmt"

	"go-clean/internal/entity"
	"go-clean/pkg/postgres"
)

const _defaultEntityCap = 64

// TranslationRepo -.
type TranslationRepo struct {
	*postgres.Postgres
}

// New -.
func New(pg *postgres.Postgres) *TranslationRepo {
	return &TranslationRepo{pg}
}

// GetHistory -.
func (r *TranslationRepo) GetHistory(ctx context.Context) ([]entity.Translation, error) {
	sql, _, err := r.Builder.
		Select("source, destination, original, translation").
		From("history").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("TranslationRepo - GetHistory - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("TranslationRepo - GetHistory - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	entities := make([]entity.Translation, 0, _defaultEntityCap)

	for rows.Next() {
		e := entity.Translation{}

		err = rows.Scan(&e.Source, &e.Destination, &e.Original, &e.Translation)
		if err != nil {
			return nil, fmt.Errorf("TranslationRepo - GetHistory - rows.Scan: %w", err)
		}

		entities = append(entities, e)
	}

	return entities, nil
}
```
<<< End Filename=translation_postgres.go
>>> Begin Filename=translation.go, Path=/Users/mac/code/go/src/github.com/mytoolzone/clean-go/internal/usecase/translation.go
```package usecase

import (
	"context"
	"fmt"

	"go-clean/internal/entity"
)

// TranslationUseCase -.
type TranslationUseCase struct {
	repo   TranslationRepo
	webAPI TranslationWebAPI
}

// New -.
func New(r TranslationRepo, w TranslationWebAPI) *TranslationUseCase {
	return &TranslationUseCase{
		repo:   r,
		webAPI: w,
	}
}

// History - getting translate history from store.
func (uc *TranslationUseCase) History(ctx context.Context) ([]entity.Translation, error) {
	translations, err := uc.repo.GetHistory(ctx)
	if err != nil {
		return nil, fmt.Errorf("TranslationUseCase - History - s.repo.GetHistory: %w", err)
	}

	return translations, nil
}
```
<<< End Filename=translation.go


### 提示模板
```
我要实现一个 任务的保存功能：  task 任务表的定义如下:
“CREATE TABLE public.tasks (
id serial4 NOT NULL,
name varchar(255) NOT NULL,
create_by int4 NULL,
created_at timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
updated_at timestamptz NULL,
finished_at timestamptz NULL,
deleted_at timestamptz NULL,
"describe" varchar(1024) NULL,
require varchar(1024) NULL,
"location" varchar NULL,
status public.task_status NULL,
max_user_count int8 NULL DEFAULT 10,
leader int8 NULL,
recorder int8 NULL,
CONSTRAINT tasks_pkey PRIMARY KEY (id),
CONSTRAINT tasks_create_by_fkey FOREIGN KEY (create_by) REFERENCES public.users(id)
);”

请使用整洁架构的思想生成对应的golang代码文件:
1. 实现 repo 接口定义, 包含任务的添加｜删除｜更新｜按照状态查询任务列表｜按照人查询对应的任务列表
2. repo 接口实现，使用 gorm 框架保存数据
3. usecase 接口定义
4. usecase 接口实现
5. repo 接口实现的测试用例编写
6. usecase 接口实现的测试用例编写
7. controller 暴露接口 添加任务|更新任务信息｜按照状态查询任务列表｜按照人查询对应的任务|列表

要求:
1. 错误信息使用 Error 类型
2. 日志使用 glog 类库
3. repo 层数据库使用 gorm 框架
4. repo 层 redis 使用
5. kafka使用 kafkago 库
```
