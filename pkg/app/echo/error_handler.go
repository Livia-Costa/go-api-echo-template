package echo

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/victorsantosbrazil/go-api-echo-template/pkg/app/errors"
)

const (
	_NOT_FOUND_ERROR = "Not found"
)

func HttpErrorHandler(err error, ctx echo.Context) {
	logError(ctx, err)

	switch err.(type) {
	case *echo.HTTPError:
		handleEchoHttpError(err.(*echo.HTTPError), ctx)
		return
	case errors.ApiError:
		handleApiError(err.(errors.ApiError), ctx)
		return
	default:
		handleError(ctx)
	}
}

func handleEchoHttpError(err *echo.HTTPError, ctx echo.Context) {
	switch err.Code {
	case http.StatusNotFound:
		if err := ctx.JSON(err.Code, errors.NotFoundError(_NOT_FOUND_ERROR)); err != nil {
			logError(ctx, err)
		}
		return
	default:
		if err := ctx.JSON(http.StatusInternalServerError, errors.InternalServerError()); err != nil {
			logError(ctx, err)
		}
	}
}

func handleApiError(err errors.ApiError, ctx echo.Context) {
	if jsonErr := ctx.JSON(err.Status, err); jsonErr != nil {
		logError(ctx, jsonErr)
	}
}

func handleError(ctx echo.Context) {
	if err := ctx.JSON(http.StatusInternalServerError, errors.InternalServerError()); err != nil {
		logError(ctx, err)
	}
}

func logError(ctx echo.Context, err error) {
	logger := EchoLogger(ctx)
	logger.Error(err.Error())
}
