package home

import "github.com/ghoshRitesh12/gojo/src/providers/anicrush/utils"

func scrapeHomePage() (*ScrapedHomePage, error) {
	data := &ScrapedHomePage{
		Status: true,
	}

	utils.Parallelize(
		data.GetMostFavoriteAnime,
		data.GetTopAiringAnime,
		data.GetSpotlightAnime,
		data.GetRecentlyUpdatedAnime,
		data.GetTrendingAnime,
		data.GetRecentlyAddedAnime,
		data.GetTopViewedAnime,
		data.GetUpcomingAnime,
		data.GetAllGenres,
	)

	// data.Status = true

	return data, nil
}
