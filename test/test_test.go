package test_test

import (
	"testing"

	"github.com/gioxtech/httpcache"
	"github.com/gioxtech/httpcache/test"
)

func TestMemoryCache(t *testing.T) {
	test.Cache(t, httpcache.NewMemoryCache())
}
