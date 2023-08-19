package puzzlebot

type (
	WebHook struct {
		User    User    `json:"user"`
		Command Command `json:"command"`
	}

	User struct {
		Id int `json:"id"`
	}

	Command struct {
		Name string `json:"name"`
	}
)
