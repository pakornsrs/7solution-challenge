package models

type BaseResponse[T any] struct {
	Data                 T      `json:"data"`
	IsError              bool   `json:"isError"`
	ErrorCode            string `json:"errorCode"`
	ErrorMessage         string `json:"message"`
	OriginalErrorMessage string `json:"originalMessage"`
}

func CreateResponseModel[T any](data T, errCode string, errMessage string, originalErrorMessage string) BaseResponse[T] {
	isError := errCode != "0000"
	return BaseResponse[T]{
		Data:                 data,
		IsError:              isError,
		ErrorCode:            errCode,
		ErrorMessage:         errMessage,
		OriginalErrorMessage: originalErrorMessage,
	}
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
