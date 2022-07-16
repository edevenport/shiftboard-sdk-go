package shiftboard

import (
	"net/http"
	"time"
)

type Client struct {
	BaseURL    string
	Auth       Auth
	HTTPClient *http.Client
	Cookies    []*http.Cookie
}

type Auth struct {
	AccessToken string
	Email       string `json:"email"`
	Password    string `json:"password"`
	UseUUID     bool   `json:"useUuid"`
}

type Response struct {
	Success bool
	Message string `json:"message,omitempty"`
	Data    Data   `json:"data,omitempty"`
	Error   *Error `json:"error,omitempty"`
}

type Error struct {
	App  string `json:"app,omitempty"`
	Code string `json:"code,omitempty"`
}

type Data struct {
	AccessToken string   `json:"access_token,omitempty"`
	Count       string   `json:"count,omitempty"`
	Page        *Page    `json:"page,omitempty"`
	Shifts      *[]Shift `json:"shifts,omitempty"`
	Sites       *[]Site  `json:"sites,omitempty"`
}

type Page struct {
	Next PageInfo `json:"next,omitempty"`
	Prev PageInfo `json:"prev,omitempty"`
	This PageInfo `json:"this,omitempty"`
}

type PageInfo struct {
	Batch string `json:"batch"`
	Start int32  `json:"start"`
}

type Shift struct {
	ID      string    `json:"id"`
	Name    string    `json:"name"`
	Updated time.Time `json:"updated"`
	Created time.Time `json:"created"`
}

type Site struct {
	ContactID string `json:"contact_id"`
	Name      string `json:"name"`
	OrgID     string `json:"org_id"`
	SiteID    string `json:"site_id"`
}
