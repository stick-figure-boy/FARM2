package apierror

const (
	// NOTE: E01xxx ... Request related errors
	ValidationErrCode = "E01001"

	// NOTE: E02xxx ... Authentication related errors
	RequiredLoginErrCode = "E02001"

	// NOTE: E03xxx ... Role related errors
	UnauthorizedAccessErrCode = "E03001"

	// NOTE: E04xxx ... Not found Data related errors
	NotFoundDataErrCode = "E04001"

	// NOTE: E05xxx ... Mismatch Data related errors
	DuplicateDataErrCode = "E05001"

	// NOTE: E09xxx ... Internal server errors
	InternalServerErrCode = "E09001"
	UnknownErrCode        = "E09999"
)
