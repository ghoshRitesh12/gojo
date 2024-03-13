package home

import (
	"github.com/ghoshRitesh12/gojo/src/config"
	"github.com/ghoshRitesh12/gojo/src/providers/anicrush/utils"
)

type ScrapedHomePage struct {
	Status                `json:"status"`
	MostFavoriteAnimes    []CommonAnimeResult  `json:"mostFavoriteAnimes"`
	TopAiringAnimes       []CommonAnimeResult  `json:"topAiringAnimes"`
	SpotlightAnimes       []CommonAnimeResult  `json:"spotlightAnimes"`
	RecentlyUpdatedAnimes []CommonAnimeResult  `json:"recentlyUpdatedAnimes"`
	TrendingAnimes        []CommonAnimeResult  `json:"trendingAnimes"`
	RecentlyAddedAnimes   []CommonAnimeResult  `json:"recentlyAddedAnimes"`
	TopViewedAnime        TopViewedAnimeResult `json:"topViewedAnime"`
	UpcomingAnimes        []CommonAnimeResult  `json:"upcomingAnimes"`
	Genres                []Genre              `json:"genres"`
}

func (shp *ScrapedHomePage) GetMostFavoriteAnime() {
	data := &MostFavoriteAnime{}
	fallbackErr := config.NewHttpError(500)
	url := utils.AJAX_BASE_URL + "/shared/v2/movie/mostFavorite?type=home"

	client := utils.NewAnicrushClient().R().
		SetHeader("X-Site", "anicrush").
		SetError(fallbackErr).
		SetResult(data)

	res, err := client.Get(url)
	if utils.CaptureSubParserErrs(res, err, url) {
		return
	}

	shp.MostFavoriteAnimes = data.Result
}

func (shp *ScrapedHomePage) GetTopAiringAnime() {
	data := &TopAiringAnime{}
	fallbackErr := config.NewHttpError(500)
	url := utils.AJAX_BASE_URL + "/shared/v2/movie/topAiring?type=home"

	client := utils.NewAnicrushClient().R().
		SetHeader("X-Site", "anicrush").
		SetError(fallbackErr).
		SetResult(data)

	res, err := client.Get(url)
	if utils.CaptureSubParserErrs(res, err, url) {
		return
	}

	shp.TopAiringAnimes = data.Result
}

func (shp *ScrapedHomePage) GetSpotlightAnime() {
	data := &SpotlightAnime{}
	fallbackErr := config.NewHttpError(500)
	url := utils.AJAX_BASE_URL + "/shared/v2/movie/spotlight"

	client := utils.NewAnicrushClient().R().
		SetHeader("X-Site", "anicrush").
		SetError(fallbackErr).
		SetResult(data)

	res, err := client.Get(url)
	if utils.CaptureSubParserErrs(res, err, url) {
		return
	}

	shp.SpotlightAnimes = data.Result
}

func (shp *ScrapedHomePage) GetRecentlyUpdatedAnime() {
	data := &RecentlyUpdatedAnime{}
	fallbackErr := config.NewHttpError(500)
	url := utils.AJAX_BASE_URL + "/shared/v2/movie/recentlyUpdated/home"

	client := utils.NewAnicrushClient().R().
		SetHeader("X-Site", "anicrush").
		SetError(fallbackErr).
		SetResult(data)

	res, err := client.Get(url)
	if utils.CaptureSubParserErrs(res, err, url) {
		return
	}

	shp.RecentlyUpdatedAnimes = data.Result
}

func (shp *ScrapedHomePage) GetTrendingAnime() {
	data := &TrendingAnime{}
	fallbackErr := config.NewHttpError(500)
	url := utils.AJAX_BASE_URL + "/shared/v2/movie/trending"

	client := utils.NewAnicrushClient().R().
		SetHeader("X-Site", "anicrush").
		SetError(fallbackErr).
		SetResult(data)

	res, err := client.Get(url)
	if utils.CaptureSubParserErrs(res, err, url) {
		return
	}

	shp.TrendingAnimes = data.Result
}

func (shp *ScrapedHomePage) GetRecentlyAddedAnime() {
	data := &RecentlyAddedAnime{}
	fallbackErr := config.NewHttpError(500)
	url := utils.AJAX_BASE_URL + "/shared/v2/movie/recentlyAdded/home"

	client := utils.NewAnicrushClient().R().
		SetHeader("X-Site", "anicrush").
		SetError(fallbackErr).
		SetResult(data)

	res, err := client.Get(url)
	if utils.CaptureSubParserErrs(res, err, url) {
		return
	}

	shp.RecentlyAddedAnimes = data.Result
}

func (shp *ScrapedHomePage) GetTopViewedAnime() {
	data := &TopViewedAnime{}
	fallbackErr := config.NewHttpError(500)
	url := utils.AJAX_BASE_URL + "/v1/movie/topViewed"

	client := utils.NewAnicrushClient().R().
		SetHeader("X-Site", "anicrush").
		SetError(fallbackErr).
		SetResult(data)

	res, err := client.Get(url)
	if utils.CaptureSubParserErrs(res, err, url) {
		return
	}

	shp.TopViewedAnime = data.Result
}

func (shp *ScrapedHomePage) GetUpcomingAnime() {
	data := &UpcomingAnime{}
	fallbackErr := config.NewHttpError(500)
	url := utils.AJAX_BASE_URL + "/shared/v2/movie/upcoming/home"

	client := utils.NewAnicrushClient().R().
		SetHeader("X-Site", "anicrush").
		SetError(fallbackErr).
		SetResult(data)

	res, err := client.Get(url)
	if utils.CaptureSubParserErrs(res, err, url) {
		return
	}

	shp.UpcomingAnimes = data.Result
}

func (shp *ScrapedHomePage) GetAllGenres() {
	data := &AllGenres{}
	url := utils.AJAX_BASE_URL + "/shared/v1/genre/all"

	client := utils.NewAnicrushClient().R().
		SetHeader("X-Site", "anicrush").
		SetResult(data)

	res, err := client.Get(url)
	if utils.CaptureSubParserErrs(res, err, url) {
		return
	}

	shp.Genres = data.Result
}
