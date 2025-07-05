package types

type Media struct {
	ID          int64  `db:"id"`
	Type        string `db:"type"`
	Title       string `db:"title"`
	Year        int32  `db:"year"`
	Genres      string `db:"genres"`
	Seasons     int32  `db:"seasons"`
	Caps        int32  `db:"caps"`
	Description string `db:"description"`
	Rating      int32  `db:"rating"`
	Image       string `db:"image"`
	Fav         bool   `db:"fav"`
	Viewed      bool   `db:"viewed"`
}

type MediaResume struct {
	ID          int64  `db:"id"`
	Type        string `db:"type"`
	Title       string `db:"title"`
	Genres      string `db:"genres"`
	Description string `db:"description"`
	Image       string `db:"image"`
	Fav         bool   `db:"fav"`
	Viewed      bool   `db:"viewed"`
}
