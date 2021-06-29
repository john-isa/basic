package constant

//==========================================================================================
// DIGITS
//==========================================================================================
const DIGITS string = "0123456789"

//==========================================================================================
// NON-PRINTING CHARACTERS
//==========================================================================================
const TAB string = "\t"
const SPACE string = " "

//==========================================================================================
// Token list understood by the Lexer
//==========================================================================================
const TT_INT string = "TT_INT"
const TT_FLOAT string = "TT_FLOAT"
const TT_PLUS string = "TT_PLUS"
const TT_MINUS string = "TT_MINUS"
const TT_MUL string = "TT_MUL"
const TT_DIV string = "TT_DIV"
const TT_LPAREN string = "TT_LPAREN"
const TT_RPAREN string = "TT_RPAREN"

//==========================================================================================
// The ERROR token when the Lexer doesn't understand
//==========================================================================================
const TT_ERROR string = "TT_ERROR" // Generic error (for catch-all situations)

//==========================================================================================
// ERROR VALUES
//==========================================================================================
const TT_ILLEGAL_CHAR string = "ILLEGAL_CHAR" // Illegal character error
