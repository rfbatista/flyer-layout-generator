package api

import (
	"algvisual/internal/designprocessor"
	"algvisual/internal/designs"
	"algvisual/internal/layoutgenerator"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/rfbatista/apitools"
	"go.uber.org/zap"

	"algvisual/internal/database"
	"algvisual/internal/infra"
	"algvisual/internal/shared"
)

func NewGenerateDesignAPI(
	queries *database.Queries,
	conn *pgxpool.Pool,
	log *zap.Logger,
	client *infra.ImageGeneratorClient,
) apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.POST)
	h.SetPath(shared.CreateNewDesignEndpoint.String())
	h.SetHandle(func(c echo.Context) error {
		var req layoutgenerator.GenerateDesignRequest
		err := c.Bind(&req)
		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}
		result, err := layoutgenerator.GenerateDesignUseCase(c.Request().Context(), req, client, queries)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, result)
	})
	return h
}

func NewUploadDesignAPI(
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
		req := designprocessor.UploadDesignFileUseCaseRequest{
			Filename: c.FormValue("filename"),
			File:     src,
		}
		out, err := designprocessor.UploadDesignFileUseCase(
			c.Request().Context(),
			db,
			req,
			storage.Upload,
			log,
		)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	})
	return h
}

func NewListDesignElementsAPI(db *database.Queries) apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.GET)
	h.SetPath(shared.EndpointListPhotoshopElements.String())
	h.SetHandle(func(c echo.Context) error {
		fmt.Println("teste")
		var req designs.ListDesignElementsUseCaseRequest
		err := c.Bind(&req)
		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}
		result, err := designs.ListDesignElementsUseCase(c.Request().Context(), req, db)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, result)
	})
	return h
}

func NewListDesignFilesAPI(db *database.Queries, log *zap.Logger) apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.GET)
	h.SetPath(shared.EndpointListPhotoshop.String())
	h.SetHandle(func(c echo.Context) error {
		var req designprocessor.ListDesignFilesRequest
		err := c.Bind(&req)
		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}
		result, err := designprocessor.ListDesignFiles(c.Request().Context(), req, db, log)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, result)
	})
	return h
}

func NewCreateComponentAPI(
	db *database.Queries,
	log *zap.Logger,
	conn *pgxpool.Pool,
) apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.POST)
	h.SetPath(shared.EndpointCreateComponent.String())
	h.SetHandle(func(c echo.Context) error {
		var req designs.CreateComponentRequest
		err := c.Bind(&req)
		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}
		result, err := designs.CreateComponentUseCase(
			c.Request().Context(),
			req,
			db,
			conn,
			log,
		)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, result)
	})
	return h
}

func NewGetDesignByIDAPI(
	db *database.Queries,
	log *zap.Logger,
	conn *pgxpool.Pool,
) apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.GET)
	h.SetPath(shared.EndpointCreateComponent.String())
	h.SetHandle(func(c echo.Context) error {
		var req designs.GetDesignByIdRequest
		err := c.Bind(&req)
		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}
		result, err := designs.GetPhotoshopByIdUseCase(c.Request().Context(), req, db, log)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, result)
	})
	return h
}
