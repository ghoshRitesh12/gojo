package home

import (
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

func (p *HomePageParser) GetMostFavoriteAnime() {
	data := &MostFavoriteAnime{}
	url := utils.AJAX_BASE_URL + "/shared/v2/movie/mostFavorite?type=home"

	res, err := p.client.R().
		SetResult(data).Get(url)
	if utils.CaptureSubParserErrs(res, err, url) {
		return
	}

	p.Resp.MostFavoriteAnimes = data.Result
}

func (p *HomePageParser) GetTopAiringAnime() {
	data := &TopAiringAnime{}
	url := utils.AJAX_BASE_URL + "/shared/v2/movie/topAiring?type=home"

	res, err := p.client.R().
		SetResult(data).Get(url)
	if utils.CaptureSubParserErrs(res, err, url) {
		return
	}

	p.Resp.TopAiringAnimes = data.Result
}

func (p *HomePageParser) GetSpotlightAnime() {
	data := &SpotlightAnime{}
	url := utils.AJAX_BASE_URL + "/shared/v2/movie/spotlight"

	res, err := p.client.R().
		SetResult(data).Get(url)
	if utils.CaptureSubParserErrs(res, err, url) {
		return
	}

	p.Resp.SpotlightAnimes = data.Result
}

func (p *HomePageParser) GetRecentlyUpdatedAnime() {
	data := &RecentlyUpdatedAnime{}
	url := utils.AJAX_BASE_URL + "/shared/v2/movie/recentlyUpdated/home"

	res, err := p.client.R().
		SetResult(data).Get(url)
	if utils.CaptureSubParserErrs(res, err, url) {
		return
	}

	p.Resp.RecentlyUpdatedAnimes = data.Result
}

func (p *HomePageParser) GetTrendingAnime() {
	data := &TrendingAnime{}
	url := utils.AJAX_BASE_URL + "/shared/v2/movie/trending"

	res, err := p.client.R().
		SetResult(data).Get(url)
	if utils.CaptureSubParserErrs(res, err, url) {
		return
	}

	p.Resp.TrendingAnimes = data.Result
}

func (p *HomePageParser) GetRecentlyAddedAnime() {
	data := &RecentlyAddedAnime{}
	url := utils.AJAX_BASE_URL + "/shared/v2/movie/recentlyAdded/home"

	res, err := p.client.R().
		SetResult(data).Get(url)
	if utils.CaptureSubParserErrs(res, err, url) {
		return
	}

	p.Resp.RecentlyAddedAnimes = data.Result
}

func (p *HomePageParser) GetTopViewedAnime() {
	data := &TopViewedAnime{}
	url := utils.AJAX_BASE_URL + "/v1/movie/topViewed"

	res, err := p.client.R().
		SetResult(data).Get(url)
	if utils.CaptureSubParserErrs(res, err, url) {
		return
	}

	p.Resp.TopViewedAnime = data.Result
}

func (p *HomePageParser) GetUpcomingAnime() {
	data := &UpcomingAnime{}
	url := utils.AJAX_BASE_URL + "/shared/v2/movie/upcoming/home"

	res, err := p.client.R().
		SetResult(data).Get(url)
	if utils.CaptureSubParserErrs(res, err, url) {
		return
	}

	p.Resp.UpcomingAnimes = data.Result
}

func (p *HomePageParser) GetAllGenres() {
	data := &AllGenres{}
	url := utils.AJAX_BASE_URL + "/shared/v1/genre/all"

	res, err := p.client.R().
		SetResult(data).Get(url)
	if utils.CaptureSubParserErrs(res, err, url) {
		return
	}

	p.Resp.Genres = data.Result
}
