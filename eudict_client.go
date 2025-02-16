package dict

import (
	"encoding/json"
	"fmt"
)

// GetCategories returns all word categories for the specified language
func (c *Client) GetCategories(language string) ([]Category, error) {
	path := fmt.Sprintf("/studylist/category?language=%s", language)
	resp, err := c.doRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result CategoryResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result.Data, nil
}

// CreateCategory creates a new word category
func (c *Client) CreateCategory(language, name string) (*Category, error) {
	payload := map[string]string{
		"language": language,
		"name":     name,
	}

	resp, err := c.doRequest("POST", "/studylist/category", payload)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Data    Category `json:"data"`
		Message string   `json:"message"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result.Data, nil
}

// RenameCategory renames an existing category
func (c *Client) RenameCategory(id, language, newName string) error {
	payload := map[string]string{
		"id":       id,
		"language": language,
		"name":     newName,
	}

	resp, err := c.doRequest("PATCH", "/studylist/category", payload)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

// DeleteCategory deletes a category
func (c *Client) DeleteCategory(id, language string) error {
	payload := map[string]string{
		"id":       id,
		"language": language,
	}

	resp, err := c.doRequest("DELETE", "/studylist/category", payload)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

// GetWords returns words from a specific category
func (c *Client) GetWords(categoryID, language string, page, pageSize int) ([]Word, error) {
	path := fmt.Sprintf("/studylist/words/%s?language=%s&page=%d&page_size=%d",
		categoryID, language, page, pageSize)

	resp, err := c.doRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result WordResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result.Data, nil
}

// AddWords adds words to a category
func (c *Client) AddWords(categoryID, language string, words []string) error {
	payload := map[string]interface{}{
		"id":       categoryID,
		"language": language,
		"words":    words,
	}

	resp, err := c.doRequest("POST", "/studylist/words", payload)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

// DeleteWords removes words from a category
func (c *Client) DeleteWords(categoryID, language string, words []string) error {
	payload := map[string]interface{}{
		"id":       categoryID,
		"language": language,
		"words":    words,
	}

	resp, err := c.doRequest("DELETE", "/studylist/words", payload)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
