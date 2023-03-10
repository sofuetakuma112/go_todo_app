package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/sofuetakuma112/go_todo_app/entity"
)

type AddTask struct {
	Service   AddTaskService
	Validator *validator.Validate
}

// ServeHTTPメソッドを実装することで、http.HandlerFunc型を満たす
func (at *AddTask) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var b struct {
		Title string `json:"title" validate:"required"`
	}
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		// リクエストボディを構造体に詰め替える処理でエラー
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	// Structメソッドは、構造体の公開フィールドをバリデーションし、
	// 特に指定がない限り、ネストした構造体も自動的にバリデーションする。
	err := at.Validator.Struct(b)
	if err != nil {
		// バリデーションエラー
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	t, err := at.Service.AddTask(ctx, b.Title)
	if err != nil {
		// データ保存時のエラー
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	rsp := struct {
		ID entity.TaskID `json:"id"`
	}{ID: t.ID}
	RespondJSON(ctx, w, rsp, http.StatusOK)
}
