package awssession

import (
	"testing"
)

func Test_New(t *testing.T) {
	t.Run("should return a session", func(t *testing.T) {
		k, err := New()

		if err != nil {
			t.Errorf("Should not return an error: %e", err)
		}

		if k == nil {
			t.Error("Should return a session")
		}
	})
}
