package gateway_42

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

func (g *Gateway42) getLocationByUserID(
	ctx context.Context,
	user42ID string,
	startAt, endAt string, // ISO‑8601
	token string,
) (LocationResponse, error) {

	var zero LocationResponse // valor zero para retorno em erro

	// compõe caminho + querystring
	path := fmt.Sprintf("/v2/users/%s/locations_stats", user42ID)
	q := url.Values{
		"begin_at": {startAt},
		"end_at":   {endAt},
	}

	resp, err := doJSON[LocationResponse](
		ctx,
		g.client, // HTTP client com retry/backoff
		http.MethodGet,
		path+"?"+q.Encode(),
		"Bearer "+token, // cabeçalho Authorization completo
		nil,             // GET não tem body
		ok2xx,
	)
	if err != nil {
		return zero, err
	}
	return resp, nil
}
