package gapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// Org represents a Grafana org.
type Org struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

// Orgs fetches and returns the Grafana orgs.
func (c *Client) Orgs() ([]Org, error) {
	orgs := make([]Org, 0)
	err := c.request("GET", "/api/orgs/", nil, nil, &orgs)
	if err != nil {
		return orgs, err
	}

	return orgs, err
}

// OrgByName fetches and returns the org whose name it's passed.
func (c *Client) OrgByName(name string) (Org, error) {
	org := Org{}
	err := c.request("GET", fmt.Sprintf("/api/orgs/name/%s", name), nil, nil, &org)
	if err != nil {
		return org, err
	}

	return org, err
}

// Org fetches and returns the org whose ID it's passed.
func (c *Client) Org(id int64) (Org, error) {
	org := Org{}
	err := c.request("GET", fmt.Sprintf("/api/orgs/%d", id), nil, nil, &org)
	if err != nil {
		return org, err
	}

	return org, err
}

// NewOrg creates a new Grafana org.
func (c *Client) NewOrg(name string) (int64, error) {
	id := int64(0)

	dataMap := map[string]string{
		"name": name,
	}
	data, err := json.Marshal(dataMap)
	if err != nil {
		return id, err
	}
	tmp := struct {
		Id int64 `json:"orgId"`
	}{}

	err = c.request("POST", "/api/orgs", nil, bytes.NewBuffer(data), &tmp)
	if err != nil {
		return id, err
	}

	return tmp.Id, err
}

// UpdateOrg updates a Grafana org.
func (c *Client) UpdateOrg(id int64, name string) error {
	dataMap := map[string]string{
		"name": name,
	}
	data, err := json.Marshal(dataMap)
	if err != nil {
		return err
	}

	return c.request("PUT", fmt.Sprintf("/api/orgs/%d", id), nil, bytes.NewBuffer(data), nil)
}

// DeleteOrg deletes the Grafana org whose ID it's passed.
func (c *Client) DeleteOrg(id int64) error {
	return c.request("DELETE", fmt.Sprintf("/api/orgs/%d", id), nil, nil, nil)
}

// UpdateCurrentOrgPreferences changes the preferences of the currently-selected organization
// https://grafana.com/docs/grafana/latest/http_api/preferences/#update-current-org-prefs
func (c *Client) UpdateCurrentOrgPreferences(prefs map[string]interface{}) error {
	payload, err := json.Marshal(prefs)
	if err != nil {
		return err
	}

	req, err := c.newRequest("PUT", "/api/org/preferences", nil, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	response, err := c.Do(req)
	if err != nil {
		return err
	} else if response.StatusCode != 200 {
		return errors.New(response.Status)
	}
	return nil
}
