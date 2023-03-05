package handler

import (
	"context"
	"encoding/json"
	"fmt"
	// "log"
	"net/http"
)

// 統一したJSONフォーマットでエラー情報を返すための型
type ErrResponse struct {
	Message string   `json:"message"`
	Details []string `json:"details,omitempty"`
}

func RespondJSON(ctx context.Context, w http.ResponseWriter, body any, status int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		// WriteHeaderは、
		// 提供されたステータス・コードでHTTPレスポンスヘッダーを送信します。
		w.WriteHeader(http.StatusInternalServerError)
		rsp := ErrResponse{
			// StatusText は、
			// HTTPステータスコードに対応するテキストを返します。
			// コードが不明な場合は、空文字列を返します。
			Message: http.StatusText(http.StatusInternalServerError),
		}
		if err := json.NewEncoder(w).Encode(rsp); err != nil {
			fmt.Printf("write error response error: %v", err)
		}
		return
	}

	w.WriteHeader(status)
	if _, err := fmt.Fprintf(w, "%s", bodyBytes); err != nil {
		fmt.Printf("write response error: %v", err)
	}
}
