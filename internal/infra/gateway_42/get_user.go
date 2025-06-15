package gateway_42

import (
	"context"
	"fmt"
	"net/http"

	"github.com/orbit-alliance/orbit-backend/internal/domain/user"
)

func (g *Gateway42) getUserByID(
	ctx context.Context,
	user42ID string,
	token string, // access‑token (user ou app com scope public)
) (userResponse, error) {

	// valor zero p/ retorno em caso de erro
	var zero userResponse

	path := fmt.Sprintf("/v2/users/%s", user42ID)

	resp, err := doJSON[userResponse](
		ctx,
		g.client, // *Client com timeout & retry
		http.MethodGet,
		path,
		"Bearer "+token, // cabeçalho Authorization completo
		nil,             // GET sem body
		ok2xx,
	)
	if err != nil {
		return zero, err
	}
	return resp, nil
}

func (g *Gateway42) GetBasicUserInfo(
	ctx context.Context,
	token string,
) (*user.UserBasicInfoDTO, error) {

	// shape exato que a API devolve
	type userMe struct {
		ID    int    `json:"id"`
		Login string `json:"login"`
	}

	u, err := doJSON[userMe](
		ctx,
		g.client,
		http.MethodGet,
		"/v2/me",
		"Bearer "+token,
		nil,
		ok2xx,
	)
	if err != nil {
		return nil, ErrFailToGetUserIn42
	}
	return &user.UserBasicInfoDTO{
		ID42:  fmt.Sprintf("%d", u.ID),
		Login: u.Login,
	}, nil
}
