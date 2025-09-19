package httputils

import (
	"encoding/json"
	"fmt"
	"github.com/audricimanuel/awb-stock-allocation/utils"
	"github.com/audricimanuel/awb-stock-allocation/utils/constants"
	"github.com/audricimanuel/errorutils"
	"github.com/pkg/errors"
	"math"
	"net/http"
)

type (
	BaseMeta struct {
		Page      int `json:"page"`
		Limit     int `json:"limit"`
		TotalData int `json:"total_data"`
		TotalPage int `json:"total_page"`
	}

	// BaseResponse is the base response
	BaseResponse struct {
		Status int         `json:"status"`
		Data   interface{} `json:"data"`
		Error  *string     `json:"error_message"`
		Meta   *BaseMeta   `json:"meta,omitempty"`
	}
)

// MapBaseResponse map response
func MapBaseResponse(w http.ResponseWriter, r *http.Request, data interface{}, err error, meta *BaseMeta) {
	var (
		errMsg     *string
		httpStatus = http.StatusOK
	)

	// Check Request ID
	reqId := "-"
	if requestID := r.Header.Get("request-id"); requestID != "" {
		reqId = requestID
	}
	dataByte, _ := json.Marshal(data)
	fmt.Printf("[RESPONSE: [%s] %s] REQUEST_ID: %s DATA: %s\n", r.Method, r.URL.String(), reqId, string(dataByte))

	var httpErr *errorutils.HttpErrorImpl
	if err != nil {
		httpStatus = http.StatusInternalServerError
		if errors.As(err, &httpErr) {
			statusCode, message := errorutils.GetStatusCode(err)
			if message != errorutils.SUCCESS {
				errMsg = &message
			}
			httpStatus = statusCode
		}
	}

	// Payload Response
	payload := BaseResponse{
		Status: httpStatus,
		Data:   data,
		Error:  errMsg,
		Meta:   meta,
	}

	// Marshal json response
	jsonResponse, _ := json.MarshalIndent(payload, "", "	")

	// Write Response
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Date", utils.TimeNow().Format(constants.FORMAT_DATETIME_TEXT))
	w.WriteHeader(httpStatus)
	w.Write(jsonResponse)
}

func SetBaseMeta(page int, limit int, totalData int) BaseMeta {
	totalPage := float64(totalData) / float64(limit)
	return BaseMeta{
		Page:      page,
		Limit:     limit,
		TotalData: totalData,
		TotalPage: int(math.Ceil(totalPage)),
	}
}
