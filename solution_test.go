package solution

import (
	"github.com/kyokomi/emoji/v2"
	"testing"
)

func TestGetMessage(t *testing.T) {
	expected := emoji.Sprint("Hello :world_map:!")
	actual := GetMessage()
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
