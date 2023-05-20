package gomono

import (
	"testing"
	"time"
)

func TestMeasurement(t *testing.T) {
	thenNano := Now()
	time.Sleep(1 * time.Millisecond)
	nowNano := Now()

	// check nanosecond resolution measurements are OK
	t.Run("BeforeNano", func(t *testing.T) {
		if !thenNano.Before(nowNano) {
			t.Fatalf("%s: then should be before now", t.Name())
		}
	})

	t.Run("AfterNano", func(t *testing.T) {
		if !nowNano.After(thenNano) {
			t.Fatalf("%s: then should be after now", t.Name())
		}
	})

	thenSec := Now()
	time.Sleep(1 * time.Second)
	nowSec := Now()

	// check second resolution measurements are OK
	t.Run("BeforeSec", func(t *testing.T) {
		if !thenSec.Before(nowSec) {
			t.Fatalf("%s: then should be before now", t.Name())
		}
	})

	t.Run("AfterSec", func(t *testing.T) {
		if !nowSec.After(thenSec) {
			t.Fatalf("%s: then should be after now", t.Name())
		}
	})

	t.Run("Diff", func(t *testing.T) {
		if nowSec.Diff(thenSec) != thenSec.Diff(nowSec) {
			t.Fatalf("%s: not symmetric", t.Name())
		}
	})

}

func BenchmarkNow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Now()
	}
}

func BenchmarkTimeNow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		time.Now().UnixNano()
	}
}
