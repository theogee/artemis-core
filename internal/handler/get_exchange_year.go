package artemis

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/theogee/artemis-core/internal/model"
	"github.com/theogee/artemis-core/pkg/logger"
	utilsHTTP "github.com/theogee/artemis-core/pkg/utils/http"
)

func (h *ArtemisHandler) GetExchangeYear(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		logPrefix = "[artemis.ArtemisHandler.GetExchangeYear]"
		log       = logger.Log

		statusCode = http.StatusInternalServerError

		resp = &model.DefaultResponse{
			Success:   false,
			ServError: []string{},
		}

		d = &model.GetExchangeYearResponse{
			ExchangeYear: []int{},
		}

		cookies = []*http.Cookie{}
	)

	defer func() {
		resp.Data = d
		utilsHTTP.SendJSON(w, cookies, resp, statusCode)
	}()

	exchangeYear, err := h.artemisUsecase.GetExchangeYear()
	if err != nil {
		log.Printf("%v error calling artemisUsecase.GetExchangeYear. err: %v", logPrefix, err)
		resp.ServError = append(resp.ServError, err.Error())
		return
	}

	statusCode = http.StatusOK
	resp.Success = true
	d.ExchangeYear = exchangeYear
}
