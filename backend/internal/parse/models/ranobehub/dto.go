package ranobehub

type RanobeHubResponse struct {
	Status int             `json:"status"`
	Data   RanobeHubNovela `json:"data"`
}

type RanobeHubNovela struct {
	ID     		int `json:"id"`
	Rating 		int `json:"rating"`
	Names  		RanobeHubNovelaNames `json:"names"`
	Year   		int `json:"year"`
	Synopsis	string `json:"synopsis"`
	Status 		RanobeHubNovelaStatus `json:"status"`
	URL 		string `json:"url"`
	Posters 	RanobeHubNovelaPosters `json:"posters"`
	Description string `json:"description"`
	Authors 	[]RanobeHubNovelaAuthors `json:"authors"`
	Translators []RanobeHubNovelaAuthors `json:"translators"`
	Volumes 	[]RanobeHubNovelaVolume `json:"volumes"`
}

type RanobeHubNovelaNames struct {
	Rus string `json:"rus"`
	Eng string `json:"eng"`	
}

type RanobeHubNovelaStatus struct {
	ID   int    `json:"id"`
	Title string `json:"title"`
}

type RanobeHubNovelaPosters struct {
	Big  string `json:"big"`
	Medium string `json:"medium"`
	Small  string `json:"small"`
	Tiny  string `json:"tiny"`
	Color string `json:"color"`
}

type RanobeHubNovelaAuthors struct {
	Name string `json:"name_eng"`
	Pivot struct {
		RanobeID int `json:"ranobe_id"`
		AuthorID int `json:"author_id"`
	} `json:"pivot"`
}

type RanobeHubNovelaVolume struct {
	ID   int `json:"id"`
	Num  int `json:"num"`
	Name string `json:"name"`
	Status RanobeHubNovelaStatus `json:"status"`
	Chapters []RanobeHubNovelaChapter `json:"chapters"`
}

type RanobeHubNovelaChapter struct {
	ID   int `json:"id"`
	Num  int `json:"num"`
	Name string `json:"name"`
	URL  string `json:"url"`
	HasImages bool `json:"has_images"`
	Text string `json:"text"`
	Images []string `json:"images"`
}