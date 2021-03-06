package session_test

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/session"
	"github.com/stretchr/testify/require"
)

func TestChecksumJSON(t *testing.T) {
	ctx := generateObjectCtx("id")

	data, err := ctx.MarshalJSON()
	require.NoError(t, err)

	ctx2 := new(session.ObjectSessionContext)
	require.NoError(t, ctx2.UnmarshalJSON(data))

	require.Equal(t, ctx, ctx2)
}

func TestTokenLifetimeJSON(t *testing.T) {
	l := generateLifetime(1, 2, 3)

	data, err := l.MarshalJSON()
	require.NoError(t, err)

	l2 := new(session.TokenLifetime)
	require.NoError(t, l2.UnmarshalJSON(data))

	require.Equal(t, l, l2)
}

func TestSessionTokenBodyJSON(t *testing.T) {
	b := generateSessionTokenBody("id")

	data, err := b.MarshalJSON()
	require.NoError(t, err)

	b2 := new(session.SessionTokenBody)
	require.NoError(t, b2.UnmarshalJSON(data))

	require.Equal(t, b, b2)
}

func TestSessionTokenJSON(t *testing.T) {
	tok := generateSessionToken("id")

	data, err := tok.MarshalJSON()
	require.NoError(t, err)

	tok2 := new(session.SessionToken)
	require.NoError(t, tok2.UnmarshalJSON(data))

	require.Equal(t, tok, tok2)
}

func TestXHeaderJSON(t *testing.T) {
	x := generateXHeader("key", "value")

	data, err := x.MarshalJSON()
	require.NoError(t, err)

	x2 := new(session.XHeader)
	require.NoError(t, x2.UnmarshalJSON(data))

	require.Equal(t, x, x2)
}

func TestRequestMetaHeaderJSON(t *testing.T) {
	r := generateRequestMetaHeader(1, "bearer", "session")

	data, err := r.MarshalJSON()
	require.NoError(t, err)

	r2 := new(session.RequestMetaHeader)
	require.NoError(t, r2.UnmarshalJSON(data))

	require.Equal(t, r, r2)
}

func TestRequestVerificationHeaderJSON(t *testing.T) {
	r := generateRequestVerificationHeader("key", "value")

	data, err := r.MarshalJSON()
	require.NoError(t, err)

	r2 := new(session.RequestVerificationHeader)
	require.NoError(t, r2.UnmarshalJSON(data))

	require.Equal(t, r, r2)
}

func TestResponseMetaHeaderJSON(t *testing.T) {
	r := generateResponseMetaHeader(1)

	data, err := r.MarshalJSON()
	require.NoError(t, err)

	r2 := new(session.ResponseMetaHeader)
	require.NoError(t, r2.UnmarshalJSON(data))

	require.Equal(t, r, r2)
}

func TestResponseVerificationHeaderJSON(t *testing.T) {
	r := generateResponseVerificationHeader("key", "value")

	data, err := r.MarshalJSON()
	require.NoError(t, err)

	r2 := new(session.ResponseVerificationHeader)
	require.NoError(t, r2.UnmarshalJSON(data))

	require.Equal(t, r, r2)
}
