package others

import (
	"bloghomnay-project/common"
	entityOthers "bloghomnay-project/services/entity/others"
	entityPosts "bloghomnay-project/services/entity/posts"
	"context"
)

type BzPost interface {
	BusinessGetALLPost(ctx context.Context) ([]entityPosts.Posts, *common.AppError)
}

type SiteMap struct {
	bzPost BzPost
}

func (s *SiteMap) BusinessSitemap(cxt context.Context) *entityOthers.Urlset {
	posts, _ := s.bzPost.BusinessGetALLPost(cxt)
	var urls []entityOthers.Url
	for _, post := range posts {
		url := entityOthers.Url{
			Loc:        "http://bloghomnay.com/post/" + post.Slug, // hoặc id hoặc gì phù hợp
			LastMod:    post.UpdatedAt.Format("2006-01-02"),
			ChangeFreq: "always",
			Priority:   "0.8",
		}
		urls = append(urls, url)
	}

	urlset := entityOthers.Urlset{
		Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9",
		Urls:  urls,
	}
	return &urlset
}

func NewBusinessSitePost(bzPost BzPost) *SiteMap {
	return &SiteMap{
		bzPost: bzPost,
	}
}
