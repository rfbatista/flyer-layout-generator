package api

import (
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/rfbatista/apitools"
	"go.uber.org/zap"

	"algvisual/internal/database"
	"algvisual/internal/infra"
	"algvisual/internal/shared"
	"algvisual/internal/usecases"
)

func NewUploadPhotoshopAPI(
	db *database.Queries,
	proc *infra.PhotoshopProcessor,
	storage infra.FileStorage,
	log *zap.Logger,
) apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.POST)
	h.SetPath(shared.EndpointUploadPhotoshop.String())
	h.SetHandle(func(c echo.Context) error {
		file, err := c.FormFile("file")
		if err != nil {
			return err
		}
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()
		req := usecases.UploadPhotoshopFileUseCaseRequest{
			Filename: c.FormValue("filename"),
			File:     src,
		}
		out, err := usecases.UploadPhotoshopFileUseCase(
			c.Request().Context(),
			db,
			req,
			storage.Upload,
			proc.ProcessFile,
			log,
		)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	})
	return h
}

func NewListPhotoshopElementsAPI(db *database.Queries) apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.GET)
	h.SetPath(shared.EndpointListPhotoshopElements.String())
	h.SetHandle(func(c echo.Context) error {
		fmt.Println("teste")
		var req usecases.ListPhotoshopElementsUseCaseRequest
		err := c.Bind(&req)
		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}
		result, err := usecases.ListPhotoshopElementsUseCase(c.Request().Context(), req, db)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, result)
	})
	return h
}

func NewListPhotoshopFilesAPI(db *database.Queries, log *zap.Logger) apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.GET)
	h.SetPath(shared.EndpointListPhotoshop.String())
	h.SetHandle(func(c echo.Context) error {
		var req usecases.ListPhotoshopFilesRequest
		err := c.Bind(&req)
		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}
		result, err := usecases.ListPhotoshopFilesUseCase(c.Request().Context(), req, db, log)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, result)
	})
	return h
}

func NewCreateComponentAPI(db *database.Queries, log *zap.Logger, conn *pgxpool.Pool) apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.POST)
	h.SetPath(shared.EndpointCreateComponent.String())
	h.SetHandle(func(c echo.Context) error {
		var req usecases.CreateComponentRequest
		err := c.Bind(&req)
		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}
		result, err := usecases.CreateComponentUseCase(c.Request().Context(), req, db, conn, log)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, result)
	})
	return h
}

func NewGetPhotoshopByIDAPI(
	db *database.Queries,
	log *zap.Logger,
	conn *pgxpool.Pool,
) apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.GET)
	h.SetPath(shared.EndpointCreateComponent.String())
	h.SetHandle(func(c echo.Context) error {
		var req usecases.GetPhotoshopByIdRequest
		err := c.Bind(&req)
		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}
		result, err := usecases.GetPhotoshopByIdUseCase(c.Request().Context(), req, db, log)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, result)
	})
	return h
}
