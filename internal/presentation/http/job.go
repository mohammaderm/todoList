package http

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mohammaderm/todoList/internal/dto"
	"github.com/mohammaderm/todoList/internal/models"
	"github.com/mohammaderm/todoList/internal/service/job"
	"github.com/mohammaderm/todoList/log"
)

type (
	JobHandler struct {
		logger      log.Logger
		jobServices job.JobServices
		HandlerHelper
	}
	JobHandlerContract interface {
		Create(w http.ResponseWriter, r *http.Request)
		GetAll(w http.ResponseWriter, r *http.Request)
		Delete(w http.ResponseWriter, r *http.Request)
		Update(w http.ResponseWriter, r *http.Request)
	}
)

func NewJobHandller(logger log.Logger, jobservice job.JobServices) JobHandlerContract {
	return &JobHandler{
		logger:        logger,
		jobServices:   jobservice,
		HandlerHelper: HandlerHelper{logger: logger},
	}
}

func (h *JobHandler) Create(w http.ResponseWriter, r *http.Request) {
	var job models.CreateJob
	err := h.readJSON(w, r, &job)
	if err != nil {
		h.errorJSON(w, errors.New("can not parse values"), http.StatusNotFound)
		return
	}
	jobreq := dto.CreateJobReq{
		Job: models.CreateJob{
			Name:        job.Name,
			Description: job.Description,
			AccountId:   r.Context().Value("UserId").(uint),
			Status:      false,
		},
	}
	err = h.jobServices.Create(r.Context(), jobreq)
	if err != nil {
		h.errorJSON(w, errors.New("can not save job"), http.StatusInternalServerError)
	}
	defer r.Body.Close()
	payload := jsonResponse{
		Error:   false,
		Message: "your job saved succesfully",
		Data:    jobreq,
	}
	h.writeJSON(w, http.StatusOK, payload)
}

func (h *JobHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	offset := r.URL.Query().Get("offset")
	offsetint, err := strconv.Atoi(offset)
	if err != nil {
		h.errorJSON(w, errors.New("failed to handle request"), http.StatusBadRequest)
		return
	}
	result, err := h.jobServices.GetAll(r.Context(), dto.GetAllJobReq{
		AccountId: r.Context().Value("UserId").(uint),
		Offset:    offsetint,
	})
	if err != nil {
		print(err.Error())
		h.errorJSON(w, errors.New("can not found any job"), http.StatusNotFound)
		return
	}
	payload := jsonResponse{
		Error:   false,
		Message: "succesfully",
		Data:    result,
	}
	h.writeJSON(w, http.StatusOK, payload)
}

func (h *JobHandler) Delete(w http.ResponseWriter, r *http.Request) {
	JobId := mux.Vars(r)["jobid"]
	JobIdUint, _ := strconv.Atoi(JobId)
	err := h.jobServices.Delete(r.Context(), dto.DeleteJobReq{
		Id:        uint(JobIdUint),
		AccountId: r.Context().Value("UserId").(uint),
	})
	if err != nil {
		h.errorJSON(w, errors.New("can not delete job"), http.StatusNotFound)
		return
	}
	payload := jsonResponse{
		Error:   false,
		Message: "succesfully",
		Data:    JobId,
	}
	h.writeJSON(w, http.StatusOK, payload)
}

func (h *JobHandler) Update(w http.ResponseWriter, r *http.Request) {
	JobId := mux.Vars(r)["jobid"]
	JobIdUint, _ := strconv.Atoi(JobId)
	err := h.jobServices.Update(r.Context(), dto.UpdateJob{
		Id:        uint(JobIdUint),
		AccountId: r.Context().Value("UserId").(uint),
	})
	if err != nil {
		h.errorJSON(w, errors.New("can not update job"), http.StatusNotFound)
		return
	}
	payload := jsonResponse{
		Error:   false,
		Message: "succesfully",
		Data:    JobId,
	}
	h.writeJSON(w, http.StatusOK, payload)
}
