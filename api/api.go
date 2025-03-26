package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/subrotokumar/go-lgtm/observability"
)

type API struct {
	telem   observability.TelemetryProvider
	httpSrv *http.Server
}

func NewAPI(telem observability.TelemetryProvider, httpSrv *http.Server) *API {
	return &API{
		telem:   telem,
		httpSrv: httpSrv,
	}
}
func (a *API) Start() {
	a.httpSrv.ListenAndServe()
}

func (a *API) GetSomething(c *gin.Context) {
	_, span := a.telem.TraceStart(c.Request.Context(), "get_something")
	defer span.End()

	something := []string{"foo", "bar", "baz"}

	c.JSON(http.StatusOK, something)
}
