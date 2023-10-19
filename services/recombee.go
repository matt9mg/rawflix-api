package services

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/matt9mg/rawflix-api/types"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type RecombeePropertyType string

var (
	RecombeePropertyTypeString    RecombeePropertyType = "string"
	RecombeePropertyTypeInt       RecombeePropertyType = "int"
	RecombeePropertyTypeSet       RecombeePropertyType = "set"
	RecombeePropertyTypeImage     RecombeePropertyType = "image"
	RecombeePropertyTypeTimestamp RecombeePropertyType = "timestamp"
)

type Recoombe struct {
	ItemProperties      *ItemProperties
	Item                *Item
	UserProperties      *UserProperties
	User                *User
	Recommendation      *Recommendation
	UserItemInteraction *UserItemInteraction
}

type User struct {
	client *http.Client
}

type UserProperties struct {
	client *http.Client
}

type ItemProperties struct {
	client *http.Client
}

type Item struct {
	client *http.Client
}

type Recommendation struct {
	client *http.Client
}

type UserItemInteraction struct {
	client *http.Client
}

func NewRecoombe(client *http.Client) *Recoombe {
	return &Recoombe{
		ItemProperties: &ItemProperties{
			client: client,
		},
		Item: &Item{
			client: client,
		},
		UserProperties: &UserProperties{
			client: client,
		},
		User: &User{
			client: client,
		},
		Recommendation: &Recommendation{
			client: client,
		},
		UserItemInteraction: &UserItemInteraction{
			client: client,
		},
	}
}

func (r *Recommendation) ReccommendItemsToItem(itemId uint, userId uint, totalRecords int, scenario string) (*types.RecombeeRecommendations, error) {
	path := buildPath(fmt.Sprintf("/recomms/items/%d/items/?targetUserId=%d&count=%d&scenario=%s&cascadeCreate=true", itemId, userId, totalRecords, scenario))

	resp, err := doRequest("GET", path, r.client, nil)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var recommendation *types.RecombeeRecommendations

	if err = json.NewDecoder(resp.Body).Decode(&recommendation); err != nil {
		return nil, buildErrorString(resp)
	}

	return recommendation, nil
}

func (r *Recommendation) ReccommendItemsToUserWithFilter(userId uint, totalRecords int, scenario string, filter string) (*types.RecombeeRecommendations, error) {
	params := struct {
		Scenario string `json:"scenario"`
		Cascade  bool   `json:"cascade"`
		Filter   string
	}{
		Scenario: scenario,
		Cascade:  true,
		Filter:   fmt.Sprintf("\"%s\" in 'genres'", filter),
	}

	path := buildPath(fmt.Sprintf("/recomms/users/%d/items/?count=%d", userId, totalRecords))

	resp, err := doRequest("GET", path, r.client, params)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var recommendation *types.RecombeeRecommendations

	if err = json.NewDecoder(resp.Body).Decode(&recommendation); err != nil {
		return nil, buildErrorString(resp)
	}

	return recommendation, nil
}

func (r *Recommendation) ReccommendItemsToUser(userId uint, totalRecords int, scenario string) (*types.RecombeeRecommendations, error) {
	path := buildPath(fmt.Sprintf("/recomms/users/%d/items/?count=%d&scenario=%s&cascadeCreate=true", userId, totalRecords, scenario))

	resp, err := doRequest("GET", path, r.client, nil)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var recommendation *types.RecombeeRecommendations

	if err = json.NewDecoder(resp.Body).Decode(&recommendation); err != nil {
		return nil, buildErrorString(resp)
	}

	return recommendation, nil
}

func (r *Recommendation) RecommendItemSegmentsToUser(userId uint, totalRecords int, scenario string) (*types.RecombeeRecommendations, error) {
	path := buildPath(fmt.Sprintf("/recomms/users/%d/item-segments/?count=%d&scenario=%s&cascadeCreate=true", userId, totalRecords, scenario))

	resp, err := doRequest("GET", path, r.client, nil)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var recommendation *types.RecombeeRecommendations

	if err = json.NewDecoder(resp.Body).Decode(&recommendation); err != nil {
		return nil, buildErrorString(resp)
	}

	return recommendation, nil
}

func (u *User) SetUserValues(userId string, body interface{}) error {
	path := buildPath(fmt.Sprintf("/users/%s", userId))

	resp, err := doRequest("POST", path, u.client, body)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return buildErrorString(resp)
	}

	return nil
}

func (u *User) AddUser(itemId string) error {
	path := buildPath(fmt.Sprintf("/users/%s", itemId))

	resp, err := doRequest("PUT", path, u.client, nil)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return buildErrorString(resp)
	}

	return nil
}

func (ip *ItemProperties) AddItemProperty(propertyName string, propertyType RecombeePropertyType) error {
	path := buildPath(fmt.Sprintf("/items/properties/%s?type=%s", propertyName, propertyType))

	resp, err := doRequest("PUT", path, ip.client, nil)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return buildErrorString(resp)
	}

	return nil
}

func (up *UserProperties) AddUserProperty(propertyName string, propertyType RecombeePropertyType) error {
	path := buildPath(fmt.Sprintf("/users/properties/%s?type=%s", propertyName, propertyType))

	resp, err := doRequest("PUT", path, up.client, nil)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return buildErrorString(resp)
	}

	return nil
}

func (i *Item) AddItem(itemId string) error {
	path := buildPath(fmt.Sprintf("/items/%s", itemId))

	resp, err := doRequest("PUT", path, i.client, nil)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return buildErrorString(resp)
	}

	return nil
}

func (i *Item) SetItemValues(itemId string, body interface{}) error {
	path := buildPath(fmt.Sprintf("/items/%s", itemId))

	resp, err := doRequest("POST", path, i.client, body)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return buildErrorString(resp)
	}

	return nil
}

func (uii *UserItemInteraction) AddDetailView(view *types.RecombeeUserItemInteraction) error {
	path := buildPath("/detailviews/")

	resp, err := doRequest("POST", path, uii.client, view)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return buildErrorString(resp)
	}

	return nil
}

func (uii *UserItemInteraction) AddBookmark(bookmark *types.RecombeeUserItemInteraction) error {
	path := buildPath("/bookmarks/")

	resp, err := doRequest("POST", path, uii.client, bookmark)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return buildErrorString(resp)
	}

	return nil
}

func signRequest(url string) string {
	mac := hmac.New(sha1.New, []byte(os.Getenv("RECOMBEE_KEY")))
	mac.Write([]byte(url))
	return hex.EncodeToString(mac.Sum(nil))
}

func buildPath(path string) string {
	sep := "?"

	if strings.Contains(path, "?") == true {
		sep = "&"
	}

	return fmt.Sprintf("/%s%s%shmac_timestamp=%d", os.Getenv("RECOMBEE_DB"), path, sep, time.Now().Unix())
}

func doRequest(method string, path string, client *http.Client, body interface{}) (*http.Response, error) {
	hash := signRequest(path)

	var (
		err error
		req *http.Request
	)

	if body != nil {
		var data []byte
		data, err = json.Marshal(body)

		if err != nil {
			return nil, err
		}

		req, err = http.NewRequest(method, fmt.Sprintf("https://rapi.recombee.com%s&hmac_sign=%s", path, hash), bytes.NewBuffer(data))
	} else {
		req, err = http.NewRequest(method, fmt.Sprintf("https://rapi.recombee.com%s&hmac_sign=%s", path, hash), nil)
	}

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	return client.Do(req)
}

func buildErrorString(resp *http.Response) error {
	data, _ := io.ReadAll(resp.Body)
	return fmt.Errorf("status code provided %d with body %s", resp.StatusCode, data)
}
