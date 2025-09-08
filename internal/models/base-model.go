package models

type BaseResponse[T any] struct {
	Data         T      `json:"data"`
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
	Error        bool   `json:"error"`
}

func CreateSuccessResponseModel[T any](data T) BaseResponse[T] {
	res := BaseResponse[T]{
		Data:         data,
		Error:        false,
		ErrorCode:    "0000",
		ErrorMessage: "success",
	}
	return res
}

func CreateResponseModel[T any](data T, errCode string, errMessage string) BaseResponse[T] {
	isError := errCode != "0000"
	res := BaseResponse[T]{
		Data:         data,
		Error:        isError,
		ErrorCode:    errCode,
		ErrorMessage: errMessage,
	}
	return res
}

type PaginationRequest struct {
	ItemPerPage int64 `json:"itemPerPage"`
	CurrentPage int64 `json:"currentPage"`
}

type Pagination struct {
	ItemPerPage int64 `json:"itemPerPage"`
	CurrentPage int64 `json:"currentPage"`
	TotalPage   int64 `json:"totalPage"`
	TotalItem   int64 `json:"totalItem"`
}
