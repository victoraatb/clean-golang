package productservice

import (
	"encoding/json"
	"net/http"

	"github.com/victoraatb/clean-go/core/dto"
)

func (service service) Fetch(response http.ResponseWriter, request *http.Request) {
	paginationRequest, _ := dto.FromValuePaginationRequestParams(request)

	products, err := service.usecase.Fetch(paginationRequest)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(response).Encode(products)
}
