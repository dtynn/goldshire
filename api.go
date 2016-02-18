package goldshire

import (
	"bytes"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

const (
	StatusInvalidSotre = 591
	StatusGetFailed    = 592
	StatusRenderFailed = 593
)

func handler(cfg *Config, store Store) func(*echo.Context) error {
	return func(c *echo.Context) error {
		path := c.Param("_*")
		if path == "" {
			return echo.NewHTTPError(http.StatusNotFound, "empty path")
		}

		meta, err := store.Get(path)
		if err != nil {
			logger.Warnf("fail to get: %s", err)
			return echo.NewHTTPError(StatusGetFailed, "fail to get: "+path)
		}

		if meta == nil {
			return echo.NewHTTPError(http.StatusNotFound, "project not found: "+path)
		}

		data := struct {
			Domain  string
			Project string
			VCS     string
			Url     string
			Path    string
		}{
			cfg.Domain,
			meta.Base,
			meta.VCS,
			meta.Url,
			path,
		}

		buf := new(bytes.Buffer)
		if err := t.Execute(buf, data); err != nil {
			logger.Warnf("fail to execute template: %s", err)
			return echo.NewHTTPError(StatusRenderFailed, "internal error")
		}

		return c.HTML(http.StatusOK, buf.String())
	}
}

func RegisterApis(mux *echo.Group, cfg *Config, store Store) {
	if cfg == nil || store == nil {
		os.Exit(2)
	}

	group := mux.Group("")
	group.Get("/*", handler(cfg, store))
}
