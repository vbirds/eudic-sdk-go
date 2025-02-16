
package dict

import (
	"os"
	"testing"
)

func TestStudyList(t *testing.T) {
	token := os.Getenv("EUDIC_TOKEN")
	if token == "" {
		t.Skip("EUDIC_TOKEN not set")
	}

	client := NewClient(token)

	t.Run("GetCategories", func(t *testing.T) {
		cats, err := client.GetCategories("en")
		if err != nil {
			t.Fatalf("Failed to get categories: %v", err)
		}
		if len(cats) == 0 {
			t.Error("Expected at least one category")
		}
	})

	// Add more test cases for other methods
}