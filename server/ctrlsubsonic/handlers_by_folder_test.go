package ctrlsubsonic

import (
	"net/url"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func TestGetIndexes(t *testing.T) {
	runQueryCases(t, testController.ServeGetIndexes, []*queryCase{
		{url.Values{}, "no_args", false},
	})
}

func TestGetMusicDirectory(t *testing.T) {
	runQueryCases(t, testController.ServeGetMusicDirectory, []*queryCase{
		{url.Values{"id": []string{"2"}}, "without_tracks", false},
		{url.Values{"id": []string{"3"}}, "with_tracks", false},
	})
}

func TestGetAlbumList(t *testing.T) {
	runQueryCases(t, testController.ServeGetAlbumList, []*queryCase{
		{url.Values{"type": []string{"alphabeticalByArtist"}}, "alpha_artist", false},
		{url.Values{"type": []string{"alphabeticalByName"}}, "alpha_name", false},
		{url.Values{"type": []string{"newest"}}, "newest", false},
		{url.Values{"type": []string{"random"}, "size": []string{"15"}}, "random", true},
	})
}

func TestSearchTwo(t *testing.T) {
	runQueryCases(t, testController.ServeSearchTwo, []*queryCase{
		{url.Values{"query": []string{"13"}}, "q_13", false},
		{url.Values{"query": []string{"ani"}}, "q_ani", false},
		{url.Values{"query": []string{"cert"}}, "q_cert", false},
	})
}
