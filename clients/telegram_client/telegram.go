package telegram_client

import (
	"ReminderAssets/lib/e"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

const (
	getUpdatesMethod  = "getUpdates"
	sendMessageMethod = "sendMessage"
)

func NewClient(host, token string) Client {
	return Client{
		host:     host,
		basePath: newBasePath(token),
	}
}
func (c *Client) doRequest(methodApi string, query url.Values) (data []byte, err error) {
	defer func() { err = e.WrapIfErr("Не выполнен request", err) }()

	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath, methodApi),
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = query.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func newBasePath(token string) string {
	return "bot" + token
}

func (c *Client) GetUpdates(offset, limit int) ([]Update, error) {
	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))
	q.Add("limit", strconv.Itoa(limit))

	data, err := c.doRequest(getUpdatesMethod, q)
	if err != nil {
		return nil, err
	}

	var resp UpdateResponse

	err = json.Unmarshal(data, &resp)
	if err != nil {
		return nil, err
	}

	return resp.Result, nil
}

func (c *Client) SendMessage(chatID int, text string) error {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatID))
	q.Add("text", text)

	_, err := c.doRequest(sendMessageMethod, q)
	if err != nil {
		return e.Wrap("Не возможно отправить сообщение", err)
	}
	return nil
}
