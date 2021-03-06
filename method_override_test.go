package buffalo

import (
	"net/url"
	"testing"

	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/httptest"
	"github.com/stretchr/testify/require"
)

func Test_MethodOverride(t *testing.T) {
	r := require.New(t)

	a := New(Options{})
	a.PUT("/", func(c Context) error {
		return c.Render(200, render.String("you put me!"))
	})

	w := httptest.New(a)
	res := w.HTML("/").Post(url.Values{"_method": []string{"PUT"}})
	r.Equal(200, res.Code)
	r.Equal("you put me!", res.Body.String())
}
