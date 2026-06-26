package bexio

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Client struct {
	httpClient    *http.Client
	BaseURL       *url.URL
	UserAgent     string
	Authorization string

	Contacts *ContactService
	Invoices *InvoiceService
}

const defaultApiUrl = "https://api.bexio.com/2.0/"
const defaultUserAgent = "Golang_go-bexio/0.1"

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	c := Client{httpClient: httpClient, UserAgent: defaultUserAgent}
	c.BaseURL, _ = url.Parse(defaultApiUrl)
	c.Contacts = &ContactService{client: &c}
	c.Invoices = &InvoiceService{client: &c}
	return &c
}

type ContactService struct {
	client *Client
}

func (c *ContactService) ListContacts() ([]Contact, error) {
	req, err := c.client.newRequest("GET", "contact", nil)
	if err != nil {
		return nil, err
	}
	var contacts []Contact
	_, err = c.client.do(req, &contacts)
	return contacts, err
}

// CreateContact creates a contact and returns it with the assigned id.
func (c *ContactService) CreateContact(contact Contact) (Contact, error) {
	req, err := c.client.newRequest("POST", "contact", contact)
	if err != nil {
		return contact, err
	}
	_, err = c.client.do(req, &contact)
	return contact, err
}

func (c *Client) ListContactGroups() ([]ContactGroup, error) {
	req, err := c.newRequest("GET", "contact_group", nil)
	if err != nil {
		return nil, err
	}
	var contactgroups []ContactGroup
	_, err = c.do(req, &contactgroups)
	return contactgroups, err
}

func (c *Client) ListArticles() ([]Article, error) {
	req, err := c.newRequest("GET", "article", nil)
	if err != nil {
		return nil, err
	}
	var article []Article
	_, err = c.do(req, &article)
	return article, err
}

func (c *Client) GetArticle(id int64) (Article, error) {
	var article Article
	req, err := c.newRequest("GET", fmt.Sprintf("article/%d", id), nil)
	if err != nil {
		return article, err
	}
	_, err = c.do(req, &article)
	return article, err
}

func (c *Client) newRequest(method, path string, body interface{}) (*http.Request, error) {
	rel := &url.URL{Path: path}
	u := c.BaseURL.ResolveReference(rel)
	fmt.Println(u)
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Authorization", "Bearer "+c.Authorization)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	log.Println(resp.StatusCode)
	switch resp.StatusCode {
	case 422:
		d, _ := ioutil.ReadAll(resp.Body)
		buf := bytes.NewBuffer(d)
		log.Println(string(d))
		var ve ValidationError
		if err = json.NewDecoder(buf).Decode(&ve); err != nil {
			return resp, err
		}
		return resp, ve
	case 200:
	case 201:
	default:
		b, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(b))
		return resp, fmt.Errorf("Invalid response status code %d", resp.StatusCode)

	}
	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}
