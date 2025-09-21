package controller

import (
	"encoding/json"
	"net/http"

	"github.com/audricimanuel/awb-stock-allocation/src/internals/dto"
	"github.com/audricimanuel/awb-stock-allocation/src/internals/service"
	"github.com/audricimanuel/awb-stock-allocation/utils/httputils"
)

type (
	AWBStockController interface {
		GetAWBStock(w http.ResponseWriter, r *http.Request)
	}

	AWBStockControllerImpl struct {
		awbStockService service.AWBStockService
	}
)

func NewAWBStockController(awbStockService service.AWBStockService) AWBStockController {
	return &AWBStockControllerImpl{
		awbStockService: awbStockService,
	}
}

// @Tags			AWB Stock
// @Summary		List of AWB stocks
// @Description	"List of AWB stocks"
// @Accept			json
// @Produce		json
// @Success		200	{object}	httputils.BaseResponse
// @Router			/awb-stocks [get]
func (a *AWBStockControllerImpl) GetAWBStock(w http.ResponseWriter, r *http.Request) {
	// TODO: implement the API to view lift of AWB stocks
	meta := httputils.SetBaseMeta(1, 10, 100)
	var err error
	httputils.MapBaseResponse(w, r, nil, err, &meta)

	awbStocks, err := a.awbStockService.GetAWBStock()
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	var resp []*dto.AWBStockResponse

	for _, awbStocks := range awbStocks {
		r := &dto.AWBStockResponse{
			AWBNumber: awbStocks.AWBNumber,
			Status:    awbStocks.Status,
		}

		resp = append(resp, r)
	}

	json.NewEncoder(w).Encode(resp)
}
