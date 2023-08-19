package puzzlebot

import (
	"bytes"
	"encoding/json"
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

func (pb Puzzlebot) SendMessage(user User, msg string) error {
	data, _ := json.Marshal(map[string]interface{}{
		"chat_id":    user.Id,
		"text":       msg,
		"parse_mode": "html",
	})

	req, err := http.NewRequest("POST", pb.url+"tg.sendMessage", bytes.NewBuffer(data))

	if err != nil {
		return errors.New("puzzlebot: can't form request")
	}

	_, err = pb.client.Do(req)
	return err
}
