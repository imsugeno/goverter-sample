package handler

import (
	"encoding/json"
	"net/http"

	gen "github.com/imsugeno/query-generator-sample/api/gen"
	"github.com/imsugeno/query-generator-sample/converter"
	"github.com/imsugeno/query-generator-sample/query"
)

type Handler struct {
	query     query.TaskQuery
	converter converter.Converter
}

func New(q query.TaskQuery, c converter.Converter) *Handler {
	return &Handler{query: q, converter: c}
}

func (h *Handler) ListTasks(w http.ResponseWriter, r *http.Request, params gen.ListTasksParams) {
	input := query.ListTasksInput{}
	if params.Status != nil {
		s := string(*params.Status)
		input.Status = &s
	}

	results, err := h.query.ListTasks(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := gen.ListTasksResponse{
		Tasks: h.converter.ConvertTaskList(results),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) GetTask(w http.ResponseWriter, r *http.Request, taskId string) {
	result, err := h.query.GetTask(query.GetTaskInput{TaskID: taskId})
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	task := h.converter.ConvertTask(*result)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}
