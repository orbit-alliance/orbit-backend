package gateway_42

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/cenkalti/backoff"
)

func doJSON[Resp any](
	ctx context.Context,
	c *Client,
	method, path, token string,
	body any,
	statusOK func(int) bool,
) (Resp, error) {

	// zero é o "valor zero" do tipo Resp. Se Resp for struct, é {}.
	var zero Resp

	// --- 1. Prepara o corpo -------------------------------------------------
	var bodyReader io.Reader
	if body != nil {
		switch v := body.(type) {
		case url.Values: // formulário x-www-form-urlencoded
			bodyReader = bytes.NewBufferString(v.Encode())
		default: // JSON
			b, err := json.Marshal(v)
			if err != nil {
				return zero, err
			}
			bodyReader = bytes.NewBuffer(b)
		}
	}

	// --- 2. Função a ser retentada pelo backoff -----------------------------
	op := func() error {
		// cria o *http.Request com contexto (suporta timeout do ctx)
		req, err := http.NewRequestWithContext(
			ctx,
			method,
			c.baseURL+path,
			bodyReader,
		)
		if err != nil {
			// erro de construção → não adianta tentar de novo
			return backoff.Permanent(err)
		}

		// cabeçalhos
		req.Header.Set("Accept", "application/json")
		if _, ok := body.(url.Values); ok {
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		} else if body != nil {
			req.Header.Set("Content-Type", "application/json")
		}
		if token != "" {
			req.Header.Set("Authorization", "Bearer "+token)
		}

		// faz a requisição
		resp, err := c.http.Do(req)
		if err != nil {
			// falha de rede / timeout → deixa o backoff tentar de novo
			return err
		}
		defer resp.Body.Close()

		// --- 3. Valida status HTTP ------------------------------------------
		if !statusOK(resp.StatusCode) {
			// 5xx ⇒ provavelmente temporário, sinaliza p/ retry
			if resp.StatusCode >= 500 {
				return fmt.Errorf("server %d", resp.StatusCode)
			}
			// 4xx ⇒ cliente errou → não adianta retry
			body, _ := io.ReadAll(resp.Body)
			return backoff.Permanent(
				fmt.Errorf("status %d: %s", resp.StatusCode, string(body)),
			)
		}

		// --- 4. Decodifica JSON no valor outer "zero" -----------------------
		if err := json.NewDecoder(resp.Body).Decode(&zero); err != nil {
			return backoff.Permanent(err) // JSON mal‑formado → não tentar de novo
		}

		return nil // sucesso → para o retry
	}

	// --- 5. Executa com backoff/retry ---------------------------------------
	err := backoff.Retry(op, backoff.WithContext(c.retryPolicy, ctx))
	return zero, err // zero agora contém a resposta (se err==nil)
}

func ok2xx(code int) bool {
	return code >= 200 && code < 300
}
