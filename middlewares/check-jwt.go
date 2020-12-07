package middlewares

import (
	"errors"
	"net/http"

	jwt "github.com/cuonggoja/web_golang/auth"
	res "github.com/cuonggoja/web_golang/utils"
	"github.com/julienschmidt/httprouter"
)

func CheckJwt(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		err := jwt.Verify(r)

		if err != nil {
			res.ERROR(w, 401, errors.New("Unauthorized"))
			return
		}
		next(w, r, ps)
	}
}
