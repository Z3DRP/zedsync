package usr

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Z3DRP/zedsync/internal/crane"
	"github.com/Z3DRP/zedsync/internal/domain"
	"github.com/Z3DRP/zedsync/internal/repos"
	"github.com/Z3DRP/zedsync/internal/request"
)

type UserAPI struct {
	s      UserService
	logger crane.Zlogrus
}

func (u UserAPI) RegisterRoutes(m *http.ServeMux, prefix string) {
	m.HandleFunc(fmt.Sprintf("POST /%v", prefix), u.HandleAddUser)
}

func (u UserAPI) HandleAddUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	select {
	case <-r.Context().Done():
		u.logger.MustDebugErr(request.ErrReqTimeout)
		request.HandleTimeout(w)
	default:
		var payload domain.User
		if err := request.ParseJSON(r, &payload); err != nil {
			u.logger.MustDebugErr(errors.Join(request.ErrJSONParse, err))
			request.WriteErr(w, http.StatusBadRequest, err)
			return
		}

		// TODO: make a dto type and call request.ValidateDto
		// to validate diferent dtos generically
		usr, err := u.s.Create(r.Context(), &payload)
		if err != nil {
			u.logger.MustDebugErr(err)
			request.WriteErr(w, http.StatusBadRequest, err)
			return
		}

		res := request.JSON{
			"user": usr,
		}

		if err := request.WriteJSON(w, http.StatusOK, res); err != nil {
			u.logger.MustDebugErr(err)
			request.WriteErr(w, http.StatusInternalServerError, err)
		}
	}
}

func (u UserAPI) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	select {
	case <-r.Context().Done():
		u.logger.MustDebugErr(request.ErrReqTimeout)
		request.HandleTimeout(w)
	default:
		id, err := request.ParseUUID(r)
		if err != nil {
			request.WriteErr(w, http.StatusBadRequest, err)
			return
		}

		usr, err := u.s.Get(r.Context(), id)
		if err != nil {
			if err != repos.ErrNoRecords {
				request.WriteErr(w, http.StatusInternalServerError, err)
				return
			}
		}

		res := request.JSON{
			"user": usr,
		}
		if err = request.WriteJSON(w, http.StatusOK, res); err != nil {
			request.WriteErr(w, http.StatusInternalServerError, err)
			return
		}
	}
}
