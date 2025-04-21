package wrapper

type GeneralError struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

var A_NOT_FOUND = GeneralError{
	Code: "A_001",
	Msg:  "User not found",
}

var A_BAD_REQUEST = GeneralError{
	Code: "A_002",
	Msg:  "Bad request",
}

var A_CREDENTIALS = GeneralError{
	Code: "A_003",
	Msg:  "Invalid credentials",
}

var A_IN_ACTIVE = GeneralError{
	Code: "A_004",
	Msg:  "User is not active",
}

var A_UNAUTHORIZED = GeneralError{
	Code: "A_005",
	Msg:  "Unauthorized",
}

var A_FORBIDDEN = GeneralError{
	Code: "A_006",
	Msg:  "Forbidden",
}

var A_DELETE = GeneralError{
	Code: "A_007",
	Msg:  "Failed to delete",
}

var A_UPDATE = GeneralError{
	Code: "A_008",
	Msg:  "Failed to update",
}

var A_CREATE = GeneralError{
	Code: "A_009",
	Msg:  "Failed to create",
}

var A_INTERNAL = GeneralError{
	Code: "A_010",
	Msg:  "Internal server error",
}
