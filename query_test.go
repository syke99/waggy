package waggy

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/syke99/waggy/internal/resources"
	"net/http"
	"testing"
)

func TestQuery(t *testing.T) {
	// Arrange
	r, _ := http.NewRequest(http.MethodGet, resources.TestRoute, nil)

	ctx := context.WithValue(r.Context(), resources.QueryParams, resources.TestQueryMap())

	r = r.Clone(ctx)

	// Act
	q := Query(r)

	// Assert
	assert.IsType(t, &QueryParams{}, q)
	assert.Equal(t, 3, len(q.qp))
	assert.Equal(t, 1, len(q.qp[resources.TestMapKey1]))
	assert.Equal(t, resources.TestMapValue1, q.qp[resources.TestMapKey1])
	assert.Equal(t, 2, len(q.qp[resources.TestMapKey2]))
	assert.Equal(t, resources.TestMapValue2, q.qp[resources.TestMapKey2])
	assert.Equal(t, 0, len(q.qp[resources.TestMapKey3]))
	assert.Equal(t, resources.TestMapValue3, q.qp[resources.TestMapKey3])
}

func TestQueryParams_Get(t *testing.T) {
	// Arrange
	r, _ := http.NewRequest(http.MethodGet, resources.TestRoute, nil)

	ctx := context.WithValue(r.Context(), resources.QueryParams, resources.TestQueryMap())

	r = r.Clone(ctx)

	q := Query(r)

	// Act
	v := q.Get(resources.TestMapKey1)

	// Assert
	assert.Equal(t, resources.TestMapValue1[0], v)
}

func TestQueryParams_Get_NoValue(t *testing.T) {
	// Arrange
	r, _ := http.NewRequest(http.MethodGet, resources.TestRoute, nil)

	ctx := context.WithValue(r.Context(), resources.QueryParams, resources.TestQueryMap())

	r = r.Clone(ctx)

	q := Query(r)

	// Act
	v := q.Get("agaga")

	// Assert
	assert.Equal(t, "", v)
}

func TestQueryParams_Set(t *testing.T) {
	// Arrange
	r, _ := http.NewRequest(http.MethodGet, resources.TestRoute, nil)

	ctx := context.WithValue(r.Context(), resources.QueryParams, resources.TestQueryMap())

	r = r.Clone(ctx)

	q := Query(r)

	// Act
	q.Set(resources.TestMapKey2, resources.Hello)

	// Assert
	assert.Equal(t, 1, len(q.qp[resources.TestMapKey2]))
	assert.Equal(t, resources.Hello, q.qp[resources.TestMapKey2][0])
}

func TestQueryParams_Add(t *testing.T) {
	// Arrange
	r, _ := http.NewRequest(http.MethodGet, resources.TestRoute, nil)

	ctx := context.WithValue(r.Context(), resources.QueryParams, resources.TestQueryMap())

	r = r.Clone(ctx)

	q := Query(r)

	// Act
	q.Add(resources.TestMapKey2, resources.WhereAmI)

	// Asset
	assert.Equal(t, 3, len(q.qp[resources.TestMapKey2]))
	assert.Equal(t, resources.WhereAmI, q.qp[resources.TestMapKey2][2])
}

func TestQueryParams_Del(t *testing.T) {
	// Arrange
	r, _ := http.NewRequest(http.MethodGet, resources.TestRoute, nil)

	ctx := context.WithValue(r.Context(), resources.QueryParams, resources.TestQueryMap())

	r = r.Clone(ctx)

	q := Query(r)

	// Act
	q.Del(resources.TestMapKey2)

	// Act
	assert.Equal(t, 0, len(q.qp[resources.TestMapKey2]))
}

func TestQueryParams_Values(t *testing.T) {
	// Arrange
	r, _ := http.NewRequest(http.MethodGet, resources.TestRoute, nil)

	ctx := context.WithValue(r.Context(), resources.QueryParams, resources.TestQueryMap())

	r = r.Clone(ctx)

	q := Query(r)

	// Act
	v := q.Values(resources.TestMapKey2)

	// Assert
	assert.Equal(t, 2, len(v))
	assert.Equal(t, resources.Hello, v[0])
	assert.Equal(t, resources.Goodbye, v[1])
}
