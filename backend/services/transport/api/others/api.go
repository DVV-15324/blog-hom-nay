package comment

import (
	"context"
	"encoding/xml"
	"net/http"

	entityOthers "bloghomnay-project/services/entity/others"

	"github.com/gin-gonic/gin"
)

type BusinessSitemap interface {
	BusinessSitemap(ctx context.Context) *entityOthers.Urlset
}

type ApiSiteMap struct {
	bz BusinessSitemap
}

func (a *ApiSiteMap) ApiSitemap() func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		urlset := a.bz.BusinessSitemap(ctx)

		c.Header("Content-Type", "application/xml")
		if err := xml.NewEncoder(c.Writer).Encode(urlset); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}
}

func NewApiSiteMap(bz BusinessSitemap) *ApiSiteMap {
	return &ApiSiteMap{
		bz: bz,
	}
}
