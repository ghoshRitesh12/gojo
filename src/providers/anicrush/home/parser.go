package home

import (
	"github.com/ghoshRitesh12/gojo/src/providers/anicrush/utils"
	"github.com/go-resty/resty/v2"
)

type HomePageParser struct {
	Resp   ScrapedHomePage
	client *resty.Client
}

func newHomePageParser() *HomePageParser {
	parser := &HomePageParser{
		Resp: ScrapedHomePage{
			Status: true,
		},
	}

	parser.client = utils.NewAnicrushClient().
		SetHeader("X-Site", "anicrush")

	return parser
}

func scrapeHomePage() (*ScrapedHomePage, error) {
	parser := newHomePageParser()

	utils.Parallelize(
		parser.GetMostFavoriteAnime,
		parser.GetTopAiringAnime,
		parser.GetSpotlightAnime,
		parser.GetRecentlyUpdatedAnime,
		parser.GetTrendingAnime,
		parser.GetRecentlyAddedAnime,
		parser.GetTopViewedAnime,
		parser.GetUpcomingAnime,
		parser.GetAllGenres,
	)

	return &parser.Resp, nil
}
