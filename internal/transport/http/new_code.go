package http

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rs/zerolog"
	qrcode "github.com/skip2/go-qrcode"
)

type NewCodeRequest struct {
	Value string `json:"value"`
}

func (h *Handler) NewCode(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context())
	var request NewCodeRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Error().Err(err).Msg("error decoding request")
	}

	png, err := qrcode.Encode(request.Value, qrcode.Medium, 256)
	if err != nil {
		log.Error().Err(err).Msg("error generating code")
	}

	w.Write([]byte(fmt.Sprintf(`<img src="data:image/png;base64, %s"/>`, base64.RawStdEncoding.EncodeToString(png))))
}
