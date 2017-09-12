package common

var ERRORS = map[int]string{
	SUCCESS:                   SUCCESS_MSG,
	E_IO_INVALID_LINK_ID:      "Invalid link ID",
	E_IO_INVALID_BOX_ID:       "Invalid box ID",
	E_IO_DECODE_BOX:           "Error decoding box object",
	E_IO_DECODE_LINK:          "Error decoding link object",
	E_IO_AUTH_PARSE:           "Error parsing authentication object",
	E_IO_TOKEN_CREATION:       "Error creating authentication token",
	E_IO_INVALID_LOGIN:        "Invalid login",
	E_DB_BOX_NOT_FOUND:        "Error: Box not found",
	E_DB_LINK_NOT_FOUND:       "Error: Link not found",
	E_DB_CREATE_BOX:           "Error creating box",
	E_DB_CREATE_LINK:          "Error creating link",
	E_DB_UPDATE_BOX:           "Error updating box",
	E_DB_UPDATE_LINK:          "Error updating link",
	E_BOX_INVALID_NAME:        "Invalid box name",
	E_BOX_INVALID_DESCRIPTION: "Invalid box description",
	E_BOX_INVALID_PASSWORD:    "Invalid box password",
}
