package numbersservice

import (
	"testing"

	"github.com/Tiny-Paws/numbers-service/pkg/numbers"
)

func TestNumbersService(t *testing.T) {
	svc := numbers.NewNumbersService()
	expected := 10
	got, err := svc.Add(5, 5)
	if err != nil || got != expected {
		t.Error("svc.Add(5, 5) =", got, "; Want", expected)
	}
}
