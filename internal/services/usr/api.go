package usr

import (
	"errors"
	"net/http"

	"github.com/Z3DRP/zedsync/internal/crane"
	"github.com/Z3DRP/zedsync/internal/domain"
	"github.com/Z3DRP/zedsync/internal/repos"
	"github.com/Z3DRP/zedsync/internal/request"
)

type UserAPI struct {
	s      UserService
	logger *crane.Zlogrus
}

func Initialize(s UserService, l *crane.Zlogrus) UserAPI {
	return UserAPI{
		s:      s,
		logger: l,
	}
}

func (u UserAPI) Name() string {
	return "user"
}

func (u UserAPI) RegisterRoutes(m *http.ServeMux) {
	// this is how i could have the main registerRoutes func call pass in prefixes
	//m.HandleFunc(fmt.Sprintf("GET /%v", prefix), u.HandleFetchUsers)
	m.HandleFunc("GET /user", u.HandleGetUser)
	m.HandleFunc("GET /user/{id}", u.HandleAddUser)
	m.HandleFunc("PUT /user/{id}", u.HandleEditUser)
	m.HandleFunc("PATCH /user/{id}", u.HandleEditUser)
	m.HandleFunc("DELTE /user/{id}", u.HandleDeleteUser)
}

func (u UserAPI) HandleAddUser(w http.ResponseWriter, r *http.Request) {
	request.SetJSONHeader(w)
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

func (u UserAPI) HandleEditUser(w http.ResponseWriter, r *http.Request) {
	request.SetJSONHeader(w)
	select {
	case <-r.Context().Done():
		u.logger.MustDebugErr(request.ErrReqTimeout)
		request.HandleTimeout(w)
	default:
		var payload domain.User
		if err := request.ParseJSON(r, &payload); err != nil {
			u.logger.MustDebugErr(err)
			request.WriteErr(w, http.StatusBadRequest, err)
			return
		}

		usr, err := u.s.Update(r.Context(), payload)
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
			return
		}
	}
}

func (u UserAPI) HandleDeleteUser(w http.ResponseWriter, r *http.Request) {
	request.SetJSONHeader(w)
	select {
	case <-r.Context().Done():
		u.logger.MustDebugErr(request.ErrReqTimeout)
		request.HandleTimeout(w)
	default:
		id, err := request.ParseUUID(r)
		if err != nil {
			u.logger.MustDebugErr(err)
			request.WriteErr(w, http.StatusBadRequest, err)
			return
		}

		err = u.s.Delete(r.Context(), id)
		if err != nil {
			u.logger.MustDebugErr(err)
			request.WriteErr(w, http.StatusInternalServerError, err)
			return
		}

		res := request.JSON{
			"success": true,
		}

		if err := request.WriteJSON(w, http.StatusOK, res); err != nil {
			u.logger.MustDebugErr(err)
			request.WriteErr(w, http.StatusInternalServerError, err)
			return
		}
	}
}

func (u UserAPI) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	request.SetJSONHeader(w)
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

func (u UserAPI) HandleFetchUsers(w http.ResponseWriter, r *http.Request) {
	request.SetJSONHeader(w)
	select {
	case <-r.Context().Done():
		u.logger.MustDebugErr(request.ErrReqTimeout)
		request.HandleTimeout(w)
	default:
		usrs, err := u.s.Fetch(r.Context())
		if err != nil {
			request.WriteErr(w, http.StatusInternalServerError, err)
			return
		}

		res := request.JSON{
			"users": usrs,
		}
		if err := request.WriteJSON(w, http.StatusOK, res); err != nil {
			request.WriteErr(w, http.StatusInternalServerError, err)
			return
		}
	}
}
