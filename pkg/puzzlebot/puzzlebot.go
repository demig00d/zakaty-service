package puzzlebot

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
)

type Puzzlebot struct {
	url    string
	client *http.Client
}

func NewPuzzleBot(client *http.Client, token string) Puzzlebot {
	return Puzzlebot{
		url:    fmt.Sprintf("https://api.puzzlebot.top/?token=%s&method=", token),
		client: client,
	}
}

func (pb Puzzlebot) SendMessage(msg string) error {
	data := fmt.Sprintf(`{'chat_id': userId, 'text': %s, 'parse_mode': "html")`, msg)
	req, err := http.NewRequest("POST", pb.url+"sendMessage", bytes.NewBufferString(data))

	if err != nil {
		return errors.New("puzzlebot: can't form request")
	}

	pb.client.Do(req)
	return nil
}
