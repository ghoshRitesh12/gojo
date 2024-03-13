package home

type Status bool

type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// NOTE: using *interface{} to have "null" in json, because default value of string is ""
type CommonAnimeResult struct {
	IDNumber         int          `json:"id_number"`
	MalID            int          `json:"mal_id,omitempty"`
	Name             string       `json:"name"`
	NameEnglish      string       `json:"name_english"`
	Slug             string       `json:"slug"`
	Overview         string       `json:"overview,omitempty"`
	Type             *interface{} `json:"type"`
	CountryCode      string       `json:"country_code"`
	PosterPath       string       `json:"poster_path"`
	BackdropPath     string       `json:"backdrop_path,omitempty"`
	TotalEpisodes    int          `json:"total_episodes"`
	LatestEpisodeSub *int         `json:"latest_episode_sub"` // using pointer to have "null" in json
	LatestEpisodeDub *int         `json:"latest_episode_dub"`
	AiringStatus     int          `json:"airing_status"`
	AiredFrom        string       `json:"aired_from,omitempty"`
	Runtime          int          `json:"runtime"`
	RatingType       *interface{} `json:"rating_type"`
	Genres           []Genre      `json:"genres"`
	ID               string       `json:"id"`
	TotalViewTop     int          `json:"totalViewTop,omitempty"`
}

type TopAiringAnime struct {
	Status bool                `json:"status"`
	Result []CommonAnimeResult `json:"result"`
}
type MostFavoriteAnime struct {
	Status bool                `json:"status"`
	Result []CommonAnimeResult `json:"result"`
}

type SpotlightAnime struct {
	Status `json:"status"`
	Result []CommonAnimeResult `json:"result"`
}

type RecentlyUpdatedAnime struct {
	Status `json:"status"`
	Result []CommonAnimeResult `json:"result"`
}

type TrendingAnime struct {
	Status `json:"status"`
	Result []CommonAnimeResult `json:"result"`
}

type RecentlyAddedAnime struct {
	Status `json:"status"`
	Result []CommonAnimeResult `json:"result"`
}

type TopViewedAnimeResult struct {
	TopViewedDay   []CommonAnimeResult `json:"topViewedDay"`
	TopViewedWeek  []CommonAnimeResult `json:"topViewedWeek"`
	TopViewedMonth []CommonAnimeResult `json:"topViewedMonth"`
}

type TopViewedAnime struct {
	Status `json:"status"`
	Result TopViewedAnimeResult `json:"result"`
}

type UpcomingAnime struct {
	Status `json:"status"`
	Result []CommonAnimeResult `json:"result"`
}

type AllGenres struct {
	Status `json:"status"`
	Result []Genre `json:"result"`
}

type HomePageResources interface {
	TopAiringAnime | MostFavoriteAnime | SpotlightAnime
}
