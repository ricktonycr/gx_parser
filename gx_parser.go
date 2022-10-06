// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // GXParser

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type GXParser struct {
	*antlr.BaseParser
}

var gxparserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func gxparserParserInit() {
	staticData := &gxparserParserStaticData
	staticData.literalNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "'('", "')'", "','",
		"'='", "'++'", "'--'", "'+'", "'-'", "'<'", "'>'", "'<='", "'>='", "'<>'",
		"'|'", "'.'", "'*'", "'/'", "'@'",
	}
	staticData.symbolicNames = []string{
		"", "WS", "COMMENT", "DocLineTag", "DocLineInfoPre", "DocLineInfo",
		"COMENTARIO", "FIN", "VARIABLE", "StringLiteral", "StringLiteralDGU",
		"StringLiteralSGU", "StringLiteralDGP", "StringLiteralSGP", "StringLiteralDZG",
		"StringLiteralSZG", "DecimalLiteral", "FOR_", "ADD_", "CLEAR_", "NONE_",
		"WHERE_", "ENDFOR_", "CASE_", "EXIT_", "NEW_", "ENDNEW_", "DO_", "ENDDO_",
		"IF_", "ELSE_", "ENDIF_", "WHILE_", "PRINT_", "EACH_", "WHEN_", "DEFINED_",
		"BY_", "SUB_", "ENDSUB_", "AND_", "OR_", "LIKE_", "DELETE_", "EVENT_",
		"ENDEVENT_", "TO_", "STEP_", "IN_", "ENDCASE_", "OTHERWISE_", "FUNCIONES",
		"CALL_", "YMDHMSTOT_", "ADDMTH_", "ADDYR_", "AFTER_", "AGE_", "ASK_",
		"BROWSERID_", "BROWSERVERSION_", "CDOW_", "CMONTH_", "COLOR_", "COLS_",
		"CONCAT_", "CONFIRMED_", "CREATE_", "CTOD_", "CTOT_", "CURSOR_", "DFRCLOSE_",
		"DFRGDATE_", "DFRGNUM_", "DFRGTXT_", "DFRNEXT_", "DFROPEN_", "DFWCLOSE_",
		"DFWNEXT_", "DFWOPEN_", "DFWPDATE_", "DFWPNUM_", "DFWPTXT_", "DAY_",
		"DECRYPT64_", "DELETEFILE_", "DOW_", "DTOC_", "ENCRYPT64_", "EOM_",
		"FILEEXIST_", "FORMAT_", "GXGETMLI_", "GXMLINES_", "GETCOOKIE_", "GETDATASTORE_",
		"GETENCRYPTIONKEY_", "GETLOCATION_", "GETSOAPERR_", "GETSOAPERRMSG_",
		"HOUR_", "INT_", "ISNULL_", "LTRIM_", "LEN_", "LEVEL_", "LINK_", "LOADBITMAP_",
		"LOWER_", "MINUTE_", "MODIFIED_", "MONTH_", "NEWLINE_", "NOW_", "NULL_",
		"NULLVALUE_", "OLD_", "OPENDOCUMENT_", "PADL_", "PADR_", "PREVIOUS_",
		"PRINTDOCUMENT_", "RGB_", "RTRIM_", "RANDOM_", "READREGKEY_", "REMOTEADDR_",
		"ROUND_", "ROWS_", "RSEED_", "SECOND_", "SERVERDATE_", "SERVERNOW_",
		"SERVERTIME_", "SETCOOKIE_", "SHELL_", "SLEEP_", "SPACE_", "STR_", "STRREPLACE_",
		"STRSEARCH_", "STRSEARCHREV_", "SUBSTR_", "SYSDATE_", "SYSTIME_", "TADD_",
		"TDIFF_", "TIME_", "TOFORMATTEDSTRING_", "TODAY_", "TRIM_", "TRUNC_",
		"TTOC_", "UDF_", "UDP_", "UPPER_", "USERCLS_", "USERID_", "VAL_", "WRKST_",
		"WRITEREGKEY_", "YMDTOD_", "YEAR_", "ATRIBUTO", "ATRIBUTOVAR", "COMPARISON",
		"OpenParen", "CloseParen", "Comma", "Assign", "PlusPlus", "MinusMinus",
		"Plus", "Minus", "LessThan", "MoreThan", "LessThanEquals", "GreaterThanEquals",
		"NotEquals", "Bar", "POINT", "STAR", "SLASH", "AT",
	}
	staticData.ruleNames = []string{
		"gxcode", "newBlock", "forBlock", "whereCondition", "ifBlock", "doCase",
		"doWhile", "condition", "subrutine", "statement", "funcion", "singleExpression",
		"printStatement", "docLine",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 183, 426, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 4, 0, 40, 8, 0, 11, 0, 12, 0, 41,
		1, 1, 1, 1, 3, 1, 46, 8, 1, 1, 1, 5, 1, 49, 8, 1, 10, 1, 12, 1, 52, 9,
		1, 1, 1, 1, 1, 1, 2, 3, 2, 57, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 63,
		8, 2, 5, 2, 65, 8, 2, 10, 2, 12, 2, 68, 9, 2, 1, 2, 3, 2, 71, 8, 2, 1,
		2, 5, 2, 74, 8, 2, 10, 2, 12, 2, 77, 9, 2, 1, 2, 1, 2, 1, 2, 3, 2, 82,
		8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 93, 8,
		2, 10, 2, 12, 2, 96, 9, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 5, 2, 109, 8, 2, 10, 2, 12, 2, 112, 9, 2, 3, 2, 114,
		8, 2, 1, 2, 1, 2, 3, 2, 118, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 3, 2, 128, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 5, 2, 139, 8, 2, 10, 2, 12, 2, 142, 9, 2, 1, 2, 1, 2, 3, 2,
		146, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 5, 2, 160, 8, 2, 10, 2, 12, 2, 163, 9, 2, 1, 2, 3, 2, 166,
		8, 2, 1, 3, 1, 3, 1, 3, 3, 3, 171, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		3, 3, 178, 8, 3, 3, 3, 180, 8, 3, 1, 4, 1, 4, 1, 4, 3, 4, 185, 8, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 196, 8, 4, 10,
		4, 12, 4, 199, 9, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 5, 4, 211, 8, 4, 10, 4, 12, 4, 214, 9, 4, 3, 4, 216, 8, 4, 1,
		4, 1, 4, 1, 5, 1, 5, 1, 5, 5, 5, 223, 8, 5, 10, 5, 12, 5, 226, 9, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 238, 8,
		5, 10, 5, 12, 5, 241, 9, 5, 5, 5, 243, 8, 5, 10, 5, 12, 5, 246, 9, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 258, 8,
		5, 10, 5, 12, 5, 261, 9, 5, 3, 5, 263, 8, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1,
		6, 1, 6, 3, 6, 271, 8, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 5, 6, 282, 8, 6, 10, 6, 12, 6, 285, 9, 6, 1, 6, 1, 6, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 298, 8, 7, 1, 7, 1,
		7, 1, 7, 5, 7, 303, 8, 7, 10, 7, 12, 7, 306, 9, 7, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 319, 8, 8, 10, 8, 12,
		8, 322, 9, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3,
		9, 333, 8, 9, 1, 9, 1, 9, 3, 9, 337, 8, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 3, 9, 345, 8, 9, 1, 9, 1, 9, 3, 9, 349, 8, 9, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 3, 9, 357, 8, 9, 1, 9, 1, 9, 3, 9, 361, 8, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 3, 9, 367, 8, 9, 1, 10, 1, 10, 1, 10, 3, 10, 372, 8, 10, 1,
		10, 1, 10, 5, 10, 376, 8, 10, 10, 10, 12, 10, 379, 9, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 3, 10, 385, 8, 10, 1, 10, 3, 10, 388, 8, 10, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 3, 11, 405, 8, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 5, 11, 413, 8, 11, 10, 11, 12, 11, 416, 9, 11, 1, 12, 1, 12, 1,
		12, 1, 13, 1, 13, 1, 13, 3, 13, 424, 8, 13, 1, 13, 0, 2, 14, 22, 14, 0,
		2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 0, 7, 2, 0, 8, 8, 16, 16,
		2, 0, 8, 8, 164, 164, 2, 0, 2, 2, 6, 6, 2, 0, 165, 165, 169, 169, 1, 0,
		40, 41, 1, 0, 181, 182, 1, 0, 172, 173, 565, 0, 39, 1, 0, 0, 0, 2, 43,
		1, 0, 0, 0, 4, 165, 1, 0, 0, 0, 6, 179, 1, 0, 0, 0, 8, 181, 1, 0, 0, 0,
		10, 219, 1, 0, 0, 0, 12, 266, 1, 0, 0, 0, 14, 297, 1, 0, 0, 0, 16, 307,
		1, 0, 0, 0, 18, 366, 1, 0, 0, 0, 20, 387, 1, 0, 0, 0, 22, 404, 1, 0, 0,
		0, 24, 417, 1, 0, 0, 0, 26, 423, 1, 0, 0, 0, 28, 40, 3, 2, 1, 0, 29, 40,
		3, 4, 2, 0, 30, 40, 3, 8, 4, 0, 31, 40, 3, 10, 5, 0, 32, 40, 3, 12, 6,
		0, 33, 40, 5, 7, 0, 0, 34, 40, 3, 16, 8, 0, 35, 40, 3, 18, 9, 0, 36, 40,
		3, 24, 12, 0, 37, 40, 3, 26, 13, 0, 38, 40, 5, 2, 0, 0, 39, 28, 1, 0, 0,
		0, 39, 29, 1, 0, 0, 0, 39, 30, 1, 0, 0, 0, 39, 31, 1, 0, 0, 0, 39, 32,
		1, 0, 0, 0, 39, 33, 1, 0, 0, 0, 39, 34, 1, 0, 0, 0, 39, 35, 1, 0, 0, 0,
		39, 36, 1, 0, 0, 0, 39, 37, 1, 0, 0, 0, 39, 38, 1, 0, 0, 0, 40, 41, 1,
		0, 0, 0, 41, 39, 1, 0, 0, 0, 41, 42, 1, 0, 0, 0, 42, 1, 1, 0, 0, 0, 43,
		45, 5, 25, 0, 0, 44, 46, 5, 6, 0, 0, 45, 44, 1, 0, 0, 0, 45, 46, 1, 0,
		0, 0, 46, 50, 1, 0, 0, 0, 47, 49, 3, 18, 9, 0, 48, 47, 1, 0, 0, 0, 49,
		52, 1, 0, 0, 0, 50, 48, 1, 0, 0, 0, 50, 51, 1, 0, 0, 0, 51, 53, 1, 0, 0,
		0, 52, 50, 1, 0, 0, 0, 53, 54, 5, 26, 0, 0, 54, 3, 1, 0, 0, 0, 55, 57,
		5, 4, 0, 0, 56, 55, 1, 0, 0, 0, 56, 57, 1, 0, 0, 0, 57, 58, 1, 0, 0, 0,
		58, 59, 5, 17, 0, 0, 59, 66, 5, 34, 0, 0, 60, 62, 5, 163, 0, 0, 61, 63,
		5, 168, 0, 0, 62, 61, 1, 0, 0, 0, 62, 63, 1, 0, 0, 0, 63, 65, 1, 0, 0,
		0, 64, 60, 1, 0, 0, 0, 65, 68, 1, 0, 0, 0, 66, 64, 1, 0, 0, 0, 66, 67,
		1, 0, 0, 0, 67, 70, 1, 0, 0, 0, 68, 66, 1, 0, 0, 0, 69, 71, 5, 6, 0, 0,
		70, 69, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 75, 1, 0, 0, 0, 72, 74, 3,
		6, 3, 0, 73, 72, 1, 0, 0, 0, 74, 77, 1, 0, 0, 0, 75, 73, 1, 0, 0, 0, 75,
		76, 1, 0, 0, 0, 76, 81, 1, 0, 0, 0, 77, 75, 1, 0, 0, 0, 78, 79, 5, 36,
		0, 0, 79, 80, 5, 37, 0, 0, 80, 82, 5, 163, 0, 0, 81, 78, 1, 0, 0, 0, 81,
		82, 1, 0, 0, 0, 82, 94, 1, 0, 0, 0, 83, 93, 3, 18, 9, 0, 84, 93, 3, 24,
		12, 0, 85, 93, 3, 2, 1, 0, 86, 93, 3, 4, 2, 0, 87, 93, 3, 8, 4, 0, 88,
		93, 3, 10, 5, 0, 89, 93, 3, 12, 6, 0, 90, 93, 3, 26, 13, 0, 91, 93, 5,
		2, 0, 0, 92, 83, 1, 0, 0, 0, 92, 84, 1, 0, 0, 0, 92, 85, 1, 0, 0, 0, 92,
		86, 1, 0, 0, 0, 92, 87, 1, 0, 0, 0, 92, 88, 1, 0, 0, 0, 92, 89, 1, 0, 0,
		0, 92, 90, 1, 0, 0, 0, 92, 91, 1, 0, 0, 0, 93, 96, 1, 0, 0, 0, 94, 92,
		1, 0, 0, 0, 94, 95, 1, 0, 0, 0, 95, 113, 1, 0, 0, 0, 96, 94, 1, 0, 0, 0,
		97, 98, 5, 35, 0, 0, 98, 110, 5, 20, 0, 0, 99, 109, 3, 18, 9, 0, 100, 109,
		3, 24, 12, 0, 101, 109, 3, 2, 1, 0, 102, 109, 3, 4, 2, 0, 103, 109, 3,
		8, 4, 0, 104, 109, 3, 10, 5, 0, 105, 109, 3, 12, 6, 0, 106, 109, 3, 26,
		13, 0, 107, 109, 5, 2, 0, 0, 108, 99, 1, 0, 0, 0, 108, 100, 1, 0, 0, 0,
		108, 101, 1, 0, 0, 0, 108, 102, 1, 0, 0, 0, 108, 103, 1, 0, 0, 0, 108,
		104, 1, 0, 0, 0, 108, 105, 1, 0, 0, 0, 108, 106, 1, 0, 0, 0, 108, 107,
		1, 0, 0, 0, 109, 112, 1, 0, 0, 0, 110, 108, 1, 0, 0, 0, 110, 111, 1, 0,
		0, 0, 111, 114, 1, 0, 0, 0, 112, 110, 1, 0, 0, 0, 113, 97, 1, 0, 0, 0,
		113, 114, 1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115, 166, 5, 22, 0, 0, 116,
		118, 5, 4, 0, 0, 117, 116, 1, 0, 0, 0, 117, 118, 1, 0, 0, 0, 118, 119,
		1, 0, 0, 0, 119, 120, 5, 17, 0, 0, 120, 121, 5, 8, 0, 0, 121, 122, 5, 169,
		0, 0, 122, 123, 7, 0, 0, 0, 123, 124, 5, 46, 0, 0, 124, 127, 7, 0, 0, 0,
		125, 126, 5, 47, 0, 0, 126, 128, 7, 0, 0, 0, 127, 125, 1, 0, 0, 0, 127,
		128, 1, 0, 0, 0, 128, 140, 1, 0, 0, 0, 129, 139, 3, 18, 9, 0, 130, 139,
		3, 24, 12, 0, 131, 139, 3, 2, 1, 0, 132, 139, 3, 4, 2, 0, 133, 139, 3,
		8, 4, 0, 134, 139, 3, 10, 5, 0, 135, 139, 3, 12, 6, 0, 136, 139, 3, 26,
		13, 0, 137, 139, 5, 2, 0, 0, 138, 129, 1, 0, 0, 0, 138, 130, 1, 0, 0, 0,
		138, 131, 1, 0, 0, 0, 138, 132, 1, 0, 0, 0, 138, 133, 1, 0, 0, 0, 138,
		134, 1, 0, 0, 0, 138, 135, 1, 0, 0, 0, 138, 136, 1, 0, 0, 0, 138, 137,
		1, 0, 0, 0, 139, 142, 1, 0, 0, 0, 140, 138, 1, 0, 0, 0, 140, 141, 1, 0,
		0, 0, 141, 143, 1, 0, 0, 0, 142, 140, 1, 0, 0, 0, 143, 166, 5, 22, 0, 0,
		144, 146, 5, 4, 0, 0, 145, 144, 1, 0, 0, 0, 145, 146, 1, 0, 0, 0, 146,
		147, 1, 0, 0, 0, 147, 148, 5, 17, 0, 0, 148, 149, 5, 8, 0, 0, 149, 150,
		5, 48, 0, 0, 150, 161, 7, 1, 0, 0, 151, 160, 3, 18, 9, 0, 152, 160, 3,
		2, 1, 0, 153, 160, 3, 4, 2, 0, 154, 160, 3, 8, 4, 0, 155, 160, 3, 10, 5,
		0, 156, 160, 3, 12, 6, 0, 157, 160, 3, 26, 13, 0, 158, 160, 5, 2, 0, 0,
		159, 151, 1, 0, 0, 0, 159, 152, 1, 0, 0, 0, 159, 153, 1, 0, 0, 0, 159,
		154, 1, 0, 0, 0, 159, 155, 1, 0, 0, 0, 159, 156, 1, 0, 0, 0, 159, 157,
		1, 0, 0, 0, 159, 158, 1, 0, 0, 0, 160, 163, 1, 0, 0, 0, 161, 159, 1, 0,
		0, 0, 161, 162, 1, 0, 0, 0, 162, 164, 1, 0, 0, 0, 163, 161, 1, 0, 0, 0,
		164, 166, 5, 22, 0, 0, 165, 56, 1, 0, 0, 0, 165, 117, 1, 0, 0, 0, 165,
		145, 1, 0, 0, 0, 166, 5, 1, 0, 0, 0, 167, 168, 5, 21, 0, 0, 168, 170, 3,
		14, 7, 0, 169, 171, 7, 2, 0, 0, 170, 169, 1, 0, 0, 0, 170, 171, 1, 0, 0,
		0, 171, 180, 1, 0, 0, 0, 172, 173, 5, 21, 0, 0, 173, 174, 3, 14, 7, 0,
		174, 175, 5, 35, 0, 0, 175, 177, 3, 14, 7, 0, 176, 178, 7, 2, 0, 0, 177,
		176, 1, 0, 0, 0, 177, 178, 1, 0, 0, 0, 178, 180, 1, 0, 0, 0, 179, 167,
		1, 0, 0, 0, 179, 172, 1, 0, 0, 0, 180, 7, 1, 0, 0, 0, 181, 182, 5, 29,
		0, 0, 182, 184, 3, 14, 7, 0, 183, 185, 5, 4, 0, 0, 184, 183, 1, 0, 0, 0,
		184, 185, 1, 0, 0, 0, 185, 197, 1, 0, 0, 0, 186, 196, 3, 18, 9, 0, 187,
		196, 3, 24, 12, 0, 188, 196, 3, 2, 1, 0, 189, 196, 3, 4, 2, 0, 190, 196,
		3, 8, 4, 0, 191, 196, 3, 10, 5, 0, 192, 196, 3, 12, 6, 0, 193, 196, 3,
		26, 13, 0, 194, 196, 5, 2, 0, 0, 195, 186, 1, 0, 0, 0, 195, 187, 1, 0,
		0, 0, 195, 188, 1, 0, 0, 0, 195, 189, 1, 0, 0, 0, 195, 190, 1, 0, 0, 0,
		195, 191, 1, 0, 0, 0, 195, 192, 1, 0, 0, 0, 195, 193, 1, 0, 0, 0, 195,
		194, 1, 0, 0, 0, 196, 199, 1, 0, 0, 0, 197, 195, 1, 0, 0, 0, 197, 198,
		1, 0, 0, 0, 198, 215, 1, 0, 0, 0, 199, 197, 1, 0, 0, 0, 200, 212, 5, 30,
		0, 0, 201, 211, 3, 18, 9, 0, 202, 211, 3, 24, 12, 0, 203, 211, 3, 2, 1,
		0, 204, 211, 3, 4, 2, 0, 205, 211, 3, 8, 4, 0, 206, 211, 3, 10, 5, 0, 207,
		211, 3, 12, 6, 0, 208, 211, 3, 26, 13, 0, 209, 211, 5, 2, 0, 0, 210, 201,
		1, 0, 0, 0, 210, 202, 1, 0, 0, 0, 210, 203, 1, 0, 0, 0, 210, 204, 1, 0,
		0, 0, 210, 205, 1, 0, 0, 0, 210, 206, 1, 0, 0, 0, 210, 207, 1, 0, 0, 0,
		210, 208, 1, 0, 0, 0, 210, 209, 1, 0, 0, 0, 211, 214, 1, 0, 0, 0, 212,
		210, 1, 0, 0, 0, 212, 213, 1, 0, 0, 0, 213, 216, 1, 0, 0, 0, 214, 212,
		1, 0, 0, 0, 215, 200, 1, 0, 0, 0, 215, 216, 1, 0, 0, 0, 216, 217, 1, 0,
		0, 0, 217, 218, 5, 31, 0, 0, 218, 9, 1, 0, 0, 0, 219, 220, 5, 27, 0, 0,
		220, 224, 5, 23, 0, 0, 221, 223, 5, 6, 0, 0, 222, 221, 1, 0, 0, 0, 223,
		226, 1, 0, 0, 0, 224, 222, 1, 0, 0, 0, 224, 225, 1, 0, 0, 0, 225, 244,
		1, 0, 0, 0, 226, 224, 1, 0, 0, 0, 227, 228, 5, 23, 0, 0, 228, 239, 3, 14,
		7, 0, 229, 238, 3, 18, 9, 0, 230, 238, 3, 2, 1, 0, 231, 238, 3, 4, 2, 0,
		232, 238, 3, 8, 4, 0, 233, 238, 3, 10, 5, 0, 234, 238, 3, 12, 6, 0, 235,
		238, 3, 26, 13, 0, 236, 238, 5, 2, 0, 0, 237, 229, 1, 0, 0, 0, 237, 230,
		1, 0, 0, 0, 237, 231, 1, 0, 0, 0, 237, 232, 1, 0, 0, 0, 237, 233, 1, 0,
		0, 0, 237, 234, 1, 0, 0, 0, 237, 235, 1, 0, 0, 0, 237, 236, 1, 0, 0, 0,
		238, 241, 1, 0, 0, 0, 239, 237, 1, 0, 0, 0, 239, 240, 1, 0, 0, 0, 240,
		243, 1, 0, 0, 0, 241, 239, 1, 0, 0, 0, 242, 227, 1, 0, 0, 0, 243, 246,
		1, 0, 0, 0, 244, 242, 1, 0, 0, 0, 244, 245, 1, 0, 0, 0, 245, 262, 1, 0,
		0, 0, 246, 244, 1, 0, 0, 0, 247, 259, 5, 50, 0, 0, 248, 258, 3, 18, 9,
		0, 249, 258, 3, 24, 12, 0, 250, 258, 3, 2, 1, 0, 251, 258, 3, 4, 2, 0,
		252, 258, 3, 8, 4, 0, 253, 258, 3, 10, 5, 0, 254, 258, 3, 12, 6, 0, 255,
		258, 3, 26, 13, 0, 256, 258, 5, 2, 0, 0, 257, 248, 1, 0, 0, 0, 257, 249,
		1, 0, 0, 0, 257, 250, 1, 0, 0, 0, 257, 251, 1, 0, 0, 0, 257, 252, 1, 0,
		0, 0, 257, 253, 1, 0, 0, 0, 257, 254, 1, 0, 0, 0, 257, 255, 1, 0, 0, 0,
		257, 256, 1, 0, 0, 0, 258, 261, 1, 0, 0, 0, 259, 257, 1, 0, 0, 0, 259,
		260, 1, 0, 0, 0, 260, 263, 1, 0, 0, 0, 261, 259, 1, 0, 0, 0, 262, 247,
		1, 0, 0, 0, 262, 263, 1, 0, 0, 0, 263, 264, 1, 0, 0, 0, 264, 265, 5, 49,
		0, 0, 265, 11, 1, 0, 0, 0, 266, 267, 5, 27, 0, 0, 267, 268, 5, 32, 0, 0,
		268, 270, 3, 14, 7, 0, 269, 271, 5, 4, 0, 0, 270, 269, 1, 0, 0, 0, 270,
		271, 1, 0, 0, 0, 271, 283, 1, 0, 0, 0, 272, 282, 3, 18, 9, 0, 273, 282,
		3, 24, 12, 0, 274, 282, 3, 2, 1, 0, 275, 282, 3, 4, 2, 0, 276, 282, 3,
		8, 4, 0, 277, 282, 3, 10, 5, 0, 278, 282, 3, 12, 6, 0, 279, 282, 3, 26,
		13, 0, 280, 282, 5, 2, 0, 0, 281, 272, 1, 0, 0, 0, 281, 273, 1, 0, 0, 0,
		281, 274, 1, 0, 0, 0, 281, 275, 1, 0, 0, 0, 281, 276, 1, 0, 0, 0, 281,
		277, 1, 0, 0, 0, 281, 278, 1, 0, 0, 0, 281, 279, 1, 0, 0, 0, 281, 280,
		1, 0, 0, 0, 282, 285, 1, 0, 0, 0, 283, 281, 1, 0, 0, 0, 283, 284, 1, 0,
		0, 0, 284, 286, 1, 0, 0, 0, 285, 283, 1, 0, 0, 0, 286, 287, 5, 28, 0, 0,
		287, 13, 1, 0, 0, 0, 288, 289, 6, 7, -1, 0, 289, 290, 3, 22, 11, 0, 290,
		291, 7, 3, 0, 0, 291, 292, 3, 22, 11, 0, 292, 298, 1, 0, 0, 0, 293, 294,
		5, 166, 0, 0, 294, 295, 3, 14, 7, 0, 295, 296, 5, 167, 0, 0, 296, 298,
		1, 0, 0, 0, 297, 288, 1, 0, 0, 0, 297, 293, 1, 0, 0, 0, 298, 304, 1, 0,
		0, 0, 299, 300, 10, 2, 0, 0, 300, 301, 7, 4, 0, 0, 301, 303, 3, 14, 7,
		3, 302, 299, 1, 0, 0, 0, 303, 306, 1, 0, 0, 0, 304, 302, 1, 0, 0, 0, 304,
		305, 1, 0, 0, 0, 305, 15, 1, 0, 0, 0, 306, 304, 1, 0, 0, 0, 307, 308, 5,
		38, 0, 0, 308, 320, 5, 9, 0, 0, 309, 319, 3, 18, 9, 0, 310, 319, 3, 24,
		12, 0, 311, 319, 3, 2, 1, 0, 312, 319, 3, 4, 2, 0, 313, 319, 3, 8, 4, 0,
		314, 319, 3, 10, 5, 0, 315, 319, 3, 12, 6, 0, 316, 319, 3, 26, 13, 0, 317,
		319, 5, 2, 0, 0, 318, 309, 1, 0, 0, 0, 318, 310, 1, 0, 0, 0, 318, 311,
		1, 0, 0, 0, 318, 312, 1, 0, 0, 0, 318, 313, 1, 0, 0, 0, 318, 314, 1, 0,
		0, 0, 318, 315, 1, 0, 0, 0, 318, 316, 1, 0, 0, 0, 318, 317, 1, 0, 0, 0,
		319, 322, 1, 0, 0, 0, 320, 318, 1, 0, 0, 0, 320, 321, 1, 0, 0, 0, 321,
		323, 1, 0, 0, 0, 322, 320, 1, 0, 0, 0, 323, 324, 5, 39, 0, 0, 324, 17,
		1, 0, 0, 0, 325, 367, 3, 20, 10, 0, 326, 332, 5, 8, 0, 0, 327, 333, 5,
		169, 0, 0, 328, 329, 5, 172, 0, 0, 329, 333, 5, 169, 0, 0, 330, 331, 5,
		173, 0, 0, 331, 333, 5, 169, 0, 0, 332, 327, 1, 0, 0, 0, 332, 328, 1, 0,
		0, 0, 332, 330, 1, 0, 0, 0, 333, 336, 1, 0, 0, 0, 334, 337, 3, 20, 10,
		0, 335, 337, 3, 22, 11, 0, 336, 334, 1, 0, 0, 0, 336, 335, 1, 0, 0, 0,
		337, 367, 1, 0, 0, 0, 338, 344, 5, 163, 0, 0, 339, 345, 5, 169, 0, 0, 340,
		341, 5, 172, 0, 0, 341, 345, 5, 169, 0, 0, 342, 343, 5, 173, 0, 0, 343,
		345, 5, 169, 0, 0, 344, 339, 1, 0, 0, 0, 344, 340, 1, 0, 0, 0, 344, 342,
		1, 0, 0, 0, 345, 348, 1, 0, 0, 0, 346, 349, 3, 20, 10, 0, 347, 349, 3,
		22, 11, 0, 348, 346, 1, 0, 0, 0, 348, 347, 1, 0, 0, 0, 349, 367, 1, 0,
		0, 0, 350, 356, 5, 164, 0, 0, 351, 357, 5, 169, 0, 0, 352, 353, 5, 172,
		0, 0, 353, 357, 5, 169, 0, 0, 354, 355, 5, 173, 0, 0, 355, 357, 5, 169,
		0, 0, 356, 351, 1, 0, 0, 0, 356, 352, 1, 0, 0, 0, 356, 354, 1, 0, 0, 0,
		357, 360, 1, 0, 0, 0, 358, 361, 3, 20, 10, 0, 359, 361, 3, 22, 11, 0, 360,
		358, 1, 0, 0, 0, 360, 359, 1, 0, 0, 0, 361, 367, 1, 0, 0, 0, 362, 363,
		5, 27, 0, 0, 363, 367, 5, 9, 0, 0, 364, 367, 5, 24, 0, 0, 365, 367, 5,
		43, 0, 0, 366, 325, 1, 0, 0, 0, 366, 326, 1, 0, 0, 0, 366, 338, 1, 0, 0,
		0, 366, 350, 1, 0, 0, 0, 366, 362, 1, 0, 0, 0, 366, 364, 1, 0, 0, 0, 366,
		365, 1, 0, 0, 0, 367, 19, 1, 0, 0, 0, 368, 369, 5, 51, 0, 0, 369, 371,
		5, 166, 0, 0, 370, 372, 3, 22, 11, 0, 371, 370, 1, 0, 0, 0, 371, 372, 1,
		0, 0, 0, 372, 377, 1, 0, 0, 0, 373, 374, 5, 168, 0, 0, 374, 376, 3, 22,
		11, 0, 375, 373, 1, 0, 0, 0, 376, 379, 1, 0, 0, 0, 377, 375, 1, 0, 0, 0,
		377, 378, 1, 0, 0, 0, 378, 380, 1, 0, 0, 0, 379, 377, 1, 0, 0, 0, 380,
		388, 5, 167, 0, 0, 381, 382, 5, 164, 0, 0, 382, 384, 5, 166, 0, 0, 383,
		385, 3, 22, 11, 0, 384, 383, 1, 0, 0, 0, 384, 385, 1, 0, 0, 0, 385, 386,
		1, 0, 0, 0, 386, 388, 5, 167, 0, 0, 387, 368, 1, 0, 0, 0, 387, 381, 1,
		0, 0, 0, 388, 21, 1, 0, 0, 0, 389, 390, 6, 11, -1, 0, 390, 405, 5, 8, 0,
		0, 391, 405, 5, 9, 0, 0, 392, 405, 5, 16, 0, 0, 393, 405, 5, 163, 0, 0,
		394, 405, 3, 20, 10, 0, 395, 405, 5, 164, 0, 0, 396, 397, 5, 25, 0, 0,
		397, 398, 5, 163, 0, 0, 398, 399, 5, 166, 0, 0, 399, 405, 5, 167, 0, 0,
		400, 401, 5, 166, 0, 0, 401, 402, 3, 22, 11, 0, 402, 403, 5, 167, 0, 0,
		403, 405, 1, 0, 0, 0, 404, 389, 1, 0, 0, 0, 404, 391, 1, 0, 0, 0, 404,
		392, 1, 0, 0, 0, 404, 393, 1, 0, 0, 0, 404, 394, 1, 0, 0, 0, 404, 395,
		1, 0, 0, 0, 404, 396, 1, 0, 0, 0, 404, 400, 1, 0, 0, 0, 405, 414, 1, 0,
		0, 0, 406, 407, 10, 3, 0, 0, 407, 408, 7, 5, 0, 0, 408, 413, 3, 22, 11,
		4, 409, 410, 10, 2, 0, 0, 410, 411, 7, 6, 0, 0, 411, 413, 3, 22, 11, 3,
		412, 406, 1, 0, 0, 0, 412, 409, 1, 0, 0, 0, 413, 416, 1, 0, 0, 0, 414,
		412, 1, 0, 0, 0, 414, 415, 1, 0, 0, 0, 415, 23, 1, 0, 0, 0, 416, 414, 1,
		0, 0, 0, 417, 418, 5, 33, 0, 0, 418, 419, 5, 163, 0, 0, 419, 25, 1, 0,
		0, 0, 420, 424, 5, 3, 0, 0, 421, 424, 5, 5, 0, 0, 422, 424, 5, 6, 0, 0,
		423, 420, 1, 0, 0, 0, 423, 421, 1, 0, 0, 0, 423, 422, 1, 0, 0, 0, 424,
		27, 1, 0, 0, 0, 61, 39, 41, 45, 50, 56, 62, 66, 70, 75, 81, 92, 94, 108,
		110, 113, 117, 127, 138, 140, 145, 159, 161, 165, 170, 177, 179, 184, 195,
		197, 210, 212, 215, 224, 237, 239, 244, 257, 259, 262, 270, 281, 283, 297,
		304, 318, 320, 332, 336, 344, 348, 356, 360, 366, 371, 377, 384, 387, 404,
		412, 414, 423,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// GXParserInit initializes any static state used to implement GXParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewGXParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func GXParserInit() {
	staticData := &gxparserParserStaticData
	staticData.once.Do(gxparserParserInit)
}

// NewGXParser produces a new parser instance for the optional input antlr.TokenStream.
func NewGXParser(input antlr.TokenStream) *GXParser {
	GXParserInit()
	this := new(GXParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &gxparserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "java-escape"

	return this
}

// GXParser tokens.
const (
	GXParserEOF                = antlr.TokenEOF
	GXParserWS                 = 1
	GXParserCOMMENT            = 2
	GXParserDocLineTag         = 3
	GXParserDocLineInfoPre     = 4
	GXParserDocLineInfo        = 5
	GXParserCOMENTARIO         = 6
	GXParserFIN                = 7
	GXParserVARIABLE           = 8
	GXParserStringLiteral      = 9
	GXParserStringLiteralDGU   = 10
	GXParserStringLiteralSGU   = 11
	GXParserStringLiteralDGP   = 12
	GXParserStringLiteralSGP   = 13
	GXParserStringLiteralDZG   = 14
	GXParserStringLiteralSZG   = 15
	GXParserDecimalLiteral     = 16
	GXParserFOR_               = 17
	GXParserADD_               = 18
	GXParserCLEAR_             = 19
	GXParserNONE_              = 20
	GXParserWHERE_             = 21
	GXParserENDFOR_            = 22
	GXParserCASE_              = 23
	GXParserEXIT_              = 24
	GXParserNEW_               = 25
	GXParserENDNEW_            = 26
	GXParserDO_                = 27
	GXParserENDDO_             = 28
	GXParserIF_                = 29
	GXParserELSE_              = 30
	GXParserENDIF_             = 31
	GXParserWHILE_             = 32
	GXParserPRINT_             = 33
	GXParserEACH_              = 34
	GXParserWHEN_              = 35
	GXParserDEFINED_           = 36
	GXParserBY_                = 37
	GXParserSUB_               = 38
	GXParserENDSUB_            = 39
	GXParserAND_               = 40
	GXParserOR_                = 41
	GXParserLIKE_              = 42
	GXParserDELETE_            = 43
	GXParserEVENT_             = 44
	GXParserENDEVENT_          = 45
	GXParserTO_                = 46
	GXParserSTEP_              = 47
	GXParserIN_                = 48
	GXParserENDCASE_           = 49
	GXParserOTHERWISE_         = 50
	GXParserFUNCIONES          = 51
	GXParserCALL_              = 52
	GXParserYMDHMSTOT_         = 53
	GXParserADDMTH_            = 54
	GXParserADDYR_             = 55
	GXParserAFTER_             = 56
	GXParserAGE_               = 57
	GXParserASK_               = 58
	GXParserBROWSERID_         = 59
	GXParserBROWSERVERSION_    = 60
	GXParserCDOW_              = 61
	GXParserCMONTH_            = 62
	GXParserCOLOR_             = 63
	GXParserCOLS_              = 64
	GXParserCONCAT_            = 65
	GXParserCONFIRMED_         = 66
	GXParserCREATE_            = 67
	GXParserCTOD_              = 68
	GXParserCTOT_              = 69
	GXParserCURSOR_            = 70
	GXParserDFRCLOSE_          = 71
	GXParserDFRGDATE_          = 72
	GXParserDFRGNUM_           = 73
	GXParserDFRGTXT_           = 74
	GXParserDFRNEXT_           = 75
	GXParserDFROPEN_           = 76
	GXParserDFWCLOSE_          = 77
	GXParserDFWNEXT_           = 78
	GXParserDFWOPEN_           = 79
	GXParserDFWPDATE_          = 80
	GXParserDFWPNUM_           = 81
	GXParserDFWPTXT_           = 82
	GXParserDAY_               = 83
	GXParserDECRYPT64_         = 84
	GXParserDELETEFILE_        = 85
	GXParserDOW_               = 86
	GXParserDTOC_              = 87
	GXParserENCRYPT64_         = 88
	GXParserEOM_               = 89
	GXParserFILEEXIST_         = 90
	GXParserFORMAT_            = 91
	GXParserGXGETMLI_          = 92
	GXParserGXMLINES_          = 93
	GXParserGETCOOKIE_         = 94
	GXParserGETDATASTORE_      = 95
	GXParserGETENCRYPTIONKEY_  = 96
	GXParserGETLOCATION_       = 97
	GXParserGETSOAPERR_        = 98
	GXParserGETSOAPERRMSG_     = 99
	GXParserHOUR_              = 100
	GXParserINT_               = 101
	GXParserISNULL_            = 102
	GXParserLTRIM_             = 103
	GXParserLEN_               = 104
	GXParserLEVEL_             = 105
	GXParserLINK_              = 106
	GXParserLOADBITMAP_        = 107
	GXParserLOWER_             = 108
	GXParserMINUTE_            = 109
	GXParserMODIFIED_          = 110
	GXParserMONTH_             = 111
	GXParserNEWLINE_           = 112
	GXParserNOW_               = 113
	GXParserNULL_              = 114
	GXParserNULLVALUE_         = 115
	GXParserOLD_               = 116
	GXParserOPENDOCUMENT_      = 117
	GXParserPADL_              = 118
	GXParserPADR_              = 119
	GXParserPREVIOUS_          = 120
	GXParserPRINTDOCUMENT_     = 121
	GXParserRGB_               = 122
	GXParserRTRIM_             = 123
	GXParserRANDOM_            = 124
	GXParserREADREGKEY_        = 125
	GXParserREMOTEADDR_        = 126
	GXParserROUND_             = 127
	GXParserROWS_              = 128
	GXParserRSEED_             = 129
	GXParserSECOND_            = 130
	GXParserSERVERDATE_        = 131
	GXParserSERVERNOW_         = 132
	GXParserSERVERTIME_        = 133
	GXParserSETCOOKIE_         = 134
	GXParserSHELL_             = 135
	GXParserSLEEP_             = 136
	GXParserSPACE_             = 137
	GXParserSTR_               = 138
	GXParserSTRREPLACE_        = 139
	GXParserSTRSEARCH_         = 140
	GXParserSTRSEARCHREV_      = 141
	GXParserSUBSTR_            = 142
	GXParserSYSDATE_           = 143
	GXParserSYSTIME_           = 144
	GXParserTADD_              = 145
	GXParserTDIFF_             = 146
	GXParserTIME_              = 147
	GXParserTOFORMATTEDSTRING_ = 148
	GXParserTODAY_             = 149
	GXParserTRIM_              = 150
	GXParserTRUNC_             = 151
	GXParserTTOC_              = 152
	GXParserUDF_               = 153
	GXParserUDP_               = 154
	GXParserUPPER_             = 155
	GXParserUSERCLS_           = 156
	GXParserUSERID_            = 157
	GXParserVAL_               = 158
	GXParserWRKST_             = 159
	GXParserWRITEREGKEY_       = 160
	GXParserYMDTOD_            = 161
	GXParserYEAR_              = 162
	GXParserATRIBUTO           = 163
	GXParserATRIBUTOVAR        = 164
	GXParserCOMPARISON         = 165
	GXParserOpenParen          = 166
	GXParserCloseParen         = 167
	GXParserComma              = 168
	GXParserAssign             = 169
	GXParserPlusPlus           = 170
	GXParserMinusMinus         = 171
	GXParserPlus               = 172
	GXParserMinus              = 173
	GXParserLessThan           = 174
	GXParserMoreThan           = 175
	GXParserLessThanEquals     = 176
	GXParserGreaterThanEquals  = 177
	GXParserNotEquals          = 178
	GXParserBar                = 179
	GXParserPOINT              = 180
	GXParserSTAR               = 181
	GXParserSLASH              = 182
	GXParserAT                 = 183
)

// GXParser rules.
const (
	GXParserRULE_gxcode           = 0
	GXParserRULE_newBlock         = 1
	GXParserRULE_forBlock         = 2
	GXParserRULE_whereCondition   = 3
	GXParserRULE_ifBlock          = 4
	GXParserRULE_doCase           = 5
	GXParserRULE_doWhile          = 6
	GXParserRULE_condition        = 7
	GXParserRULE_subrutine        = 8
	GXParserRULE_statement        = 9
	GXParserRULE_funcion          = 10
	GXParserRULE_singleExpression = 11
	GXParserRULE_printStatement   = 12
	GXParserRULE_docLine          = 13
)

// IGxcodeContext is an interface to support dynamic dispatch.
type IGxcodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGxcodeContext differentiates from other interfaces.
	IsGxcodeContext()
}

type GxcodeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGxcodeContext() *GxcodeContext {
	var p = new(GxcodeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GXParserRULE_gxcode
	return p
}

func (*GxcodeContext) IsGxcodeContext() {}

func NewGxcodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GxcodeContext {
	var p = new(GxcodeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GXParserRULE_gxcode

	return p
}

func (s *GxcodeContext) GetParser() antlr.Parser { return s.parser }

func (s *GxcodeContext) AllNewBlock() []INewBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INewBlockContext); ok {
			len++
		}
	}

	tst := make([]INewBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INewBlockContext); ok {
			tst[i] = t.(INewBlockContext)
			i++
		}
	}

	return tst
}

func (s *GxcodeContext) NewBlock(i int) INewBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INewBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INewBlockContext)
}

func (s *GxcodeContext) AllForBlock() []IForBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IForBlockContext); ok {
			len++
		}
	}

	tst := make([]IForBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IForBlockContext); ok {
			tst[i] = t.(IForBlockContext)
			i++
		}
	}

	return tst
}

func (s *GxcodeContext) ForBlock(i int) IForBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForBlockContext)
}

func (s *GxcodeContext) AllIfBlock() []IIfBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIfBlockContext); ok {
			len++
		}
	}

	tst := make([]IIfBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIfBlockContext); ok {
			tst[i] = t.(IIfBlockContext)
			i++
		}
	}

	return tst
}

func (s *GxcodeContext) IfBlock(i int) IIfBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfBlockContext)
}

func (s *GxcodeContext) AllDoCase() []IDoCaseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDoCaseContext); ok {
			len++
		}
	}

	tst := make([]IDoCaseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDoCaseContext); ok {
			tst[i] = t.(IDoCaseContext)
			i++
		}
	}

	return tst
}

func (s *GxcodeContext) DoCase(i int) IDoCaseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDoCaseContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDoCaseContext)
}

func (s *GxcodeContext) AllDoWhile() []IDoWhileContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDoWhileContext); ok {
			len++
		}
	}

	tst := make([]IDoWhileContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDoWhileContext); ok {
			tst[i] = t.(IDoWhileContext)
			i++
		}
	}

	return tst
}

func (s *GxcodeContext) DoWhile(i int) IDoWhileContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDoWhileContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDoWhileContext)
}

func (s *GxcodeContext) AllFIN() []antlr.TerminalNode {
	return s.GetTokens(GXParserFIN)
}

func (s *GxcodeContext) FIN(i int) antlr.TerminalNode {
	return s.GetToken(GXParserFIN, i)
}

func (s *GxcodeContext) AllSubrutine() []ISubrutineContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISubrutineContext); ok {
			len++
		}
	}

	tst := make([]ISubrutineContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISubrutineContext); ok {
			tst[i] = t.(ISubrutineContext)
			i++
		}
	}

	return tst
}

func (s *GxcodeContext) Subrutine(i int) ISubrutineContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubrutineContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubrutineContext)
}

func (s *GxcodeContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *GxcodeContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *GxcodeContext) AllPrintStatement() []IPrintStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPrintStatementContext); ok {
			len++
		}
	}

	tst := make([]IPrintStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPrintStatementContext); ok {
			tst[i] = t.(IPrintStatementContext)
			i++
		}
	}

	return tst
}

func (s *GxcodeContext) PrintStatement(i int) IPrintStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintStatementContext)
}

func (s *GxcodeContext) AllDocLine() []IDocLineContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDocLineContext); ok {
			len++
		}
	}

	tst := make([]IDocLineContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDocLineContext); ok {
			tst[i] = t.(IDocLineContext)
			i++
		}
	}

	return tst
}

func (s *GxcodeContext) DocLine(i int) IDocLineContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDocLineContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDocLineContext)
}

func (s *GxcodeContext) AllCOMMENT() []antlr.TerminalNode {
	return s.GetTokens(GXParserCOMMENT)
}

func (s *GxcodeContext) COMMENT(i int) antlr.TerminalNode {
	return s.GetToken(GXParserCOMMENT, i)
}

func (s *GxcodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GxcodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GxcodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GXParserListener); ok {
		listenerT.EnterGxcode(s)
	}
}

func (s *GxcodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GXParserListener); ok {
		listenerT.ExitGxcode(s)
	}
}

func (p *GXParser) Gxcode() (localctx IGxcodeContext) {
	this := p
	_ = this

	localctx = NewGxcodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GXParserRULE_gxcode)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2260880096100860) != 0 || _la == GXParserATRIBUTO || _la == GXParserATRIBUTOVAR {
		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(28)
				p.NewBlock()
			}

		case 2:
			{
				p.SetState(29)
				p.ForBlock()
			}

		case 3:
			{
				p.SetState(30)
				p.IfBlock()
			}

		case 4:
			{
				p.SetState(31)
				p.DoCase()
			}

		case 5:
			{
				p.SetState(32)
				p.DoWhile()
			}

		case 6:
			{
				p.SetState(33)
				p.Match(GXParserFIN)
			}

		case 7:
			{
				p.SetState(34)
				p.Subrutine()
			}

		case 8:
			{
				p.SetState(35)
				p.Statement()
			}

		case 9:
			{
				p.SetState(36)
				p.PrintStatement()
			}

		case 10:
			{
				p.SetState(37)
				p.DocLine()
			}

		case 11:
			{
				p.SetState(38)
				p.Match(GXParserCOMMENT)
			}

		}

		p.SetState(41)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// INewBlockContext is an interface to support dynamic dispatch.
type INewBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetComentario returns the comentario token.
	GetComentario() antlr.Token

	// SetComentario sets the comentario token.
	SetComentario(antlr.Token)

	// IsNewBlockContext differentiates from other interfaces.
	IsNewBlockContext()
}

type NewBlockContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	comentario antlr.Token
}

func NewEmptyNewBlockContext() *NewBlockContext {
	var p = new(NewBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GXParserRULE_newBlock
	return p
}

func (*NewBlockContext) IsNewBlockContext() {}

func NewNewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NewBlockContext {
	var p = new(NewBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GXParserRULE_newBlock

	return p
}

func (s *NewBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *NewBlockContext) GetComentario() antlr.Token { return s.comentario }

func (s *NewBlockContext) SetComentario(v antlr.Token) { s.comentario = v }

func (s *NewBlockContext) NEW_() antlr.TerminalNode {
	return s.GetToken(GXParserNEW_, 0)
}

func (s *NewBlockContext) ENDNEW_() antlr.TerminalNode {
	return s.GetToken(GXParserENDNEW_, 0)
}

func (s *NewBlockContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *NewBlockContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *NewBlockContext) COMENTARIO() antlr.TerminalNode {
	return s.GetToken(GXParserCOMENTARIO, 0)
}

func (s *NewBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NewBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NewBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GXParserListener); ok {
		listenerT.EnterNewBlock(s)
	}
}

func (s *NewBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GXParserListener); ok {
		listenerT.ExitNewBlock(s)
	}
}

func (p *GXParser) NewBlock() (localctx INewBlockContext) {
	this := p
	_ = this

	localctx = NewNewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GXParserRULE_newBlock)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(43)
		p.Match(GXParserNEW_)
	}
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GXParserCOMENTARIO {
		{
			p.SetState(44)

			var _m = p.Match(GXParserCOMENTARIO)

			localctx.(*NewBlockContext).comentario = _m
		}

	}
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2260596057702656) != 0 || _la == GXParserATRIBUTO || _la == GXParserATRIBUTOVAR {
		{
			p.SetState(47)
			p.Statement()
		}

		p.SetState(52)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(53)
		p.Match(GXParserENDNEW_)
	}

	return localctx
}

// IForBlockContext is an interface to support dynamic dispatch.
type IForBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetDoc returns the doc token.
	GetDoc() antlr.Token

	// Get_ATRIBUTO returns the _ATRIBUTO token.
	Get_ATRIBUTO() antlr.Token

	// GetComentario returns the comentario token.
	GetComentario() antlr.Token

	// GetEn returns the en token.
	GetEn() antlr.Token

	// GetDesde returns the desde token.
	GetDesde() antlr.Token

	// GetHasta returns the hasta token.
	GetHasta() antlr.Token

	// GetCada returns the cada token.
	GetCada() antlr.Token

	// GetSdt returns the sdt token.
	GetSdt() antlr.Token

	// SetDoc sets the doc token.
	SetDoc(antlr.Token)

	// Set_ATRIBUTO sets the _ATRIBUTO token.
	Set_ATRIBUTO(antlr.Token)

	// SetComentario sets the comentario token.
	SetComentario(antlr.Token)

	// SetEn sets the en token.
	SetEn(antlr.Token)

	// SetDesde sets the desde token.
	SetDesde(antlr.Token)

	// SetHasta sets the hasta token.
	SetHasta(antlr.Token)

	// SetCada sets the cada token.
	SetCada(antlr.Token)

	// SetSdt sets the sdt token.
	SetSdt(antlr.Token)

	// GetIndices returns the indices token list.
	GetIndices() []antlr.Token

	// SetIndices sets the indices token list.
	SetIndices([]antlr.Token)

	// Get_whereCondition returns the _whereCondition rule contexts.
	Get_whereCondition() IWhereConditionContext

	// Set_whereCondition sets the _whereCondition rule contexts.
	Set_whereCondition(IWhereConditionContext)

	// GetCondiciones returns the condiciones rule context list.
	GetCondiciones() []IWhereConditionContext

	// SetCondiciones sets the condiciones rule context list.
	SetCondiciones([]IWhereConditionContext)

	// IsForBlockContext differentiates from other interfaces.
	IsForBlockContext()
}

type ForBlockContext struct {
	*antlr.BaseParserRuleContext
	parser          antlr.Parser
	doc             antlr.Token
	_ATRIBUTO       antlr.Token
	indices         []antlr.Token
	comentario      antlr.Token
	_whereCondition IWhereConditionContext
	condiciones     []IWhereConditionContext
	en              antlr.Token
	desde           antlr.Token
	hasta           antlr.Token
	cada            antlr.Token
	sdt             antlr.Token
}

func NewEmptyForBlockContext() *ForBlockContext {
	var p = new(ForBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GXParserRULE_forBlock
	return p
}

func (*ForBlockContext) IsForBlockContext() {}

func NewForBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForBlockContext {
	var p = new(ForBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GXParserRULE_forBlock

	return p
}

func (s *ForBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *ForBlockContext) GetDoc() antlr.Token { return s.doc }

func (s *ForBlockContext) Get_ATRIBUTO() antlr.Token { return s._ATRIBUTO }

func (s *ForBlockContext) GetComentario() antlr.Token { return s.comentario }

func (s *ForBlockContext) GetEn() antlr.Token { return s.en }

func (s *ForBlockContext) GetDesde() antlr.Token { return s.desde }

func (s *ForBlockContext) GetHasta() antlr.Token { return s.hasta }

func (s *ForBlockContext) GetCada() antlr.Token { return s.cada }

func (s *ForBlockContext) GetSdt() antlr.Token { return s.sdt }

func (s *ForBlockContext) SetDoc(v antlr.Token) { s.doc = v }

func (s *ForBlockContext) Set_ATRIBUTO(v antlr.Token) { s._ATRIBUTO = v }

func (s *ForBlockContext) SetComentario(v antlr.Token) { s.comentario = v }

func (s *ForBlockContext) SetEn(v antlr.Token) { s.en = v }

func (s *ForBlockContext) SetDesde(v antlr.Token) { s.desde = v }

func (s *ForBlockContext) SetHasta(v antlr.Token) { s.hasta = v }

func (s *ForBlockContext) SetCada(v antlr.Token) { s.cada = v }

func (s *ForBlockContext) SetSdt(v antlr.Token) { s.sdt = v }

func (s *ForBlockContext) GetIndices() []antlr.Token { return s.indices }

func (s *ForBlockContext) SetIndices(v []antlr.Token) { s.indices = v }

func (s *ForBlockContext) Get_whereCondition() IWhereConditionContext { return s._whereCondition }

func (s *ForBlockContext) Set_whereCondition(v IWhereConditionContext) { s._whereCondition = v }

func (s *ForBlockContext) GetCondiciones() []IWhereConditionContext { return s.condiciones }

func (s *ForBlockContext) SetCondiciones(v []IWhereConditionContext) { s.condiciones = v }

func (s *ForBlockContext) FOR_() antlr.TerminalNode {
	return s.GetToken(GXParserFOR_, 0)
}

func (s *ForBlockContext) EACH_() antlr.TerminalNode {
	return s.GetToken(GXParserEACH_, 0)
}

func (s *ForBlockContext) ENDFOR_() antlr.TerminalNode {
	return s.GetToken(GXParserENDFOR_, 0)
}

func (s *ForBlockContext) DEFINED_() antlr.TerminalNode {
	return s.GetToken(GXParserDEFINED_, 0)
}

func (s *ForBlockContext) BY_() antlr.TerminalNode {
	return s.GetToken(GXParserBY_, 0)
}

func (s *ForBlockContext) AllATRIBUTO() []antlr.TerminalNode {
	return s.GetTokens(GXParserATRIBUTO)
}

func (s *ForBlockContext) ATRIBUTO(i int) antlr.TerminalNode {
	return s.GetToken(GXParserATRIBUTO, i)
}

func (s *ForBlockContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *ForBlockContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ForBlockContext) AllPrintStatement() []IPrintStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPrintStatementContext); ok {
			len++
		}
	}

	tst := make([]IPrintStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPrintStatementContext); ok {
			tst[i] = t.(IPrintStatementContext)
			i++
		}
	}

	return tst
}

func (s *ForBlockContext) PrintStatement(i int) IPrintStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintStatementContext)
}

func (s *ForBlockContext) AllNewBlock() []INewBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INewBlockContext); ok {
			len++
		}
	}

	tst := make([]INewBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INewBlockContext); ok {
			tst[i] = t.(INewBlockContext)
			i++
		}
	}

	return tst
}

func (s *ForBlockContext) NewBlock(i int) INewBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INewBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INewBlockContext)
}

func (s *ForBlockContext) AllForBlock() []IForBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IForBlockContext); ok {
			len++
		}
	}

	tst := make([]IForBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IForBlockContext); ok {
			tst[i] = t.(IForBlockContext)
			i++
		}
	}

	return tst
}

func (s *ForBlockContext) ForBlock(i int) IForBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForBlockContext)
}

func (s *ForBlockContext) AllIfBlock() []IIfBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIfBlockContext); ok {
			len++
		}
	}

	tst := make([]IIfBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIfBlockContext); ok {
			tst[i] = t.(IIfBlockContext)
			i++
		}
	}

	return tst
}

func (s *ForBlockContext) IfBlock(i int) IIfBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfBlockContext)
}

func (s *ForBlockContext) AllDoCase() []IDoCaseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDoCaseContext); ok {
			len++
		}
	}

	tst := make([]IDoCaseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDoCaseContext); ok {
			tst[i] = t.(IDoCaseContext)
			i++
		}
	}

	return tst
}

func (s *ForBlockContext) DoCase(i int) IDoCaseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDoCaseContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDoCaseContext)
}

func (s *ForBlockContext) AllDoWhile() []IDoWhileContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDoWhileContext); ok {
			len++
		}
	}

	tst := make([]IDoWhileContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDoWhileContext); ok {
			tst[i] = t.(IDoWhileContext)
			i++
		}
	}

	return tst
}

func (s *ForBlockContext) DoWhile(i int) IDoWhileContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDoWhileContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDoWhileContext)
}

func (s *ForBlockContext) AllDocLine() []IDocLineContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDocLineContext); ok {
			len++
		}
	}

	tst := make([]IDocLineContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDocLineContext); ok {
			tst[i] = t.(IDocLineContext)
			i++
		}
	}

	return tst
}

func (s *ForBlockContext) DocLine(i int) IDocLineContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDocLineContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDocLineContext)
}

func (s *ForBlockContext) AllCOMMENT() []antlr.TerminalNode {
	return s.GetTokens(GXParserCOMMENT)
}

func (s *ForBlockContext) COMMENT(i int) antlr.TerminalNode {
	return s.GetToken(GXParserCOMMENT, i)
}

func (s *ForBlockContext) WHEN_() antlr.TerminalNode {
	return s.GetToken(GXParserWHEN_, 0)
}

func (s *ForBlockContext) NONE_() antlr.TerminalNode {
	return s.GetToken(GXParserNONE_, 0)
}

func (s *ForBlockContext) DocLineInfoPre() antlr.TerminalNode {
	return s.GetToken(GXParserDocLineInfoPre, 0)
}

func (s *ForBlockContext) COMENTARIO() antlr.TerminalNode {
	return s.GetToken(GXParserCOMENTARIO, 0)
}

func (s *ForBlockContext) AllWhereCondition() []IWhereConditionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IWhereConditionContext); ok {
			len++
		}
	}

	tst := make([]IWhereConditionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IWhereConditionContext); ok {
			tst[i] = t.(IWhereConditionContext)
			i++
		}
	}

	return tst
}

func (s *ForBlockContext) WhereCondition(i int) IWhereConditionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhereConditionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhereConditionContext)
}

func (s *ForBlockContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(GXParserComma)
}

func (s *ForBlockContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(GXParserComma, i)
}

func (s *ForBlockContext) Assign() antlr.TerminalNode {
	return s.GetToken(GXParserAssign, 0)
}

func (s *ForBlockContext) TO_() antlr.TerminalNode {
	return s.GetToken(GXParserTO_, 0)
}

func (s *ForBlockContext) AllVARIABLE() []antlr.TerminalNode {
	return s.GetTokens(GXParserVARIABLE)
}

func (s *ForBlockContext) VARIABLE(i int) antlr.TerminalNode {
	return s.GetToken(GXParserVARIABLE, i)
}

func (s *ForBlockContext) AllDecimalLiteral() []antlr.TerminalNode {
	return s.GetTokens(GXParserDecimalLiteral)
}

func (s *ForBlockContext) DecimalLiteral(i int) antlr.TerminalNode {
	return s.GetToken(GXParserDecimalLiteral, i)
}

func (s *ForBlockContext) STEP_() antlr.TerminalNode {
	return s.GetToken(GXParserSTEP_, 0)
}

func (s *ForBlockContext) IN_() antlr.TerminalNode {
	return s.GetToken(GXParserIN_, 0)
}

func (s *ForBlockContext) ATRIBUTOVAR() antlr.TerminalNode {
	return s.GetToken(GXParserATRIBUTOVAR, 0)
}

func (s *ForBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GXParserListener); ok {
		listenerT.EnterForBlock(s)
	}
}

func (s *ForBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GXParserListener); ok {
		listenerT.ExitForBlock(s)
	}
}

func (p *GXParser) ForBlock() (localctx IForBlockContext) {
	this := p
	_ = this

	localctx = NewForBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GXParserRULE_forBlock)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GXParserDocLineInfoPre {
			{
				p.SetState(55)

				var _m = p.Match(GXParserDocLineInfoPre)

				localctx.(*ForBlockContext).doc = _m
			}

		}
		{
			p.SetState(58)
			p.Match(GXParserFOR_)
		}
		{
			p.SetState(59)
			p.Match(GXParserEACH_)
		}
		p.SetState(66)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(60)

					var _m = p.Match(GXParserATRIBUTO)

					localctx.(*ForBlockContext)._ATRIBUTO = _m
				}
				localctx.(*ForBlockContext).indices = append(localctx.(*ForBlockContext).indices, localctx.(*ForBlockContext)._ATRIBUTO)
				p.SetState(62)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == GXParserComma {
					{
						p.SetState(61)
						p.Match(GXParserComma)
					}

				}

			}
			p.SetState(68)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
		}
		p.SetState(70)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(69)

				var _m = p.Match(GXParserCOMENTARIO)

				localctx.(*ForBlockContext).comentario = _m
			}

		}
		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == GXParserWHERE_ {
			{
				p.SetState(72)

				var _x = p.WhereCondition()

				localctx.(*ForBlockContext)._whereCondition = _x
			}
			localctx.(*ForBlockContext).condiciones = append(localctx.(*ForBlockContext).condiciones, localctx.(*ForBlockContext)._whereCondition)

			p.SetState(77)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GXParserDEFINED_ {
			{
				p.SetState(78)
				p.Match(GXParserDEFINED_)
			}
			{
				p.SetState(79)
				p.Match(GXParserBY_)
			}
			{
				p.SetState(80)
				p.Match(GXParserATRIBUTO)
			}

		}
		p.SetState(94)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2260605218193788) != 0 || _la == GXParserATRIBUTO || _la == GXParserATRIBUTOVAR {
			p.SetState(92)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(83)
					p.Statement()
				}

			case 2:
				{
					p.SetState(84)
					p.PrintStatement()
				}

			case 3:
				{
					p.SetState(85)
					p.NewBlock()
				}

			case 4:
				{
					p.SetState(86)
					p.ForBlock()
				}

			case 5:
				{
					p.SetState(87)
					p.IfBlock()
				}

			case 6:
				{
					p.SetState(88)
					p.DoCase()
				}

			case 7:
				{
					p.SetState(89)
					p.DoWhile()
				}

			case 8:
				{
					p.SetState(90)
					p.DocLine()
				}

			case 9:
				{
					p.SetState(91)
					p.Match(GXParserCOMMENT)
				}

			}

			p.SetState(96)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(113)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GXParserWHEN_ {
			{
				p.SetState(97)
				p.Match(GXParserWHEN_)
			}
			{
				p.SetState(98)
				p.Match(GXParserNONE_)
			}
			p.SetState(110)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2260605218193788) != 0 || _la == GXParserATRIBUTO || _la == GXParserATRIBUTOVAR {
				p.SetState(108)
				p.GetErrorHandler().Sync(p)
				switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
				case 1:
					{
						p.SetState(99)
						p.Statement()
					}

				case 2:
					{
						p.SetState(100)
						p.PrintStatement()
					}

				case 3:
					{
						p.SetState(101)
						p.NewBlock()
					}

				case 4:
					{
						p.SetState(102)
						p.ForBlock()
					}

				case 5:
					{
						p.SetState(103)
						p.IfBlock()
					}

				case 6:
					{
						p.SetState(104)
						p.DoCase()
					}

				case 7:
					{
						p.SetState(105)
						p.DoWhile()
					}

				case 8:
					{
						p.SetState(106)
						p.DocLine()
					}

				case 9:
					{
						p.SetState(107)
						p.Match(GXParserCOMMENT)
					}

				}

				p.SetState(112)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(115)
			p.Match(GXParserENDFOR_)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GXParserDocLineInfoPre {
			{
				p.SetState(116)

				var _m = p.Match(GXParserDocLineInfoPre)

				localctx.(*ForBlockContext).doc = _m
			}

		}
		{
			p.SetState(119)
			p.Match(GXParserFOR_)
		}
		{
			p.SetState(120)

			var _m = p.Match(GXParserVARIABLE)

			localctx.(*ForBlockContext).en = _m
		}
		{
			p.SetState(121)
			p.Match(GXParserAssign)
		}
		{
			p.SetState(122)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ForBlockContext).desde = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == GXParserVARIABLE || _la == GXParserDecimalLiteral) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ForBlockContext).desde = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(123)
			p.Match(GXParserTO_)
		}
		{
			p.SetState(124)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ForBlockContext).hasta = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == GXParserVARIABLE || _la == GXParserDecimalLiteral) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ForBlockContext).hasta = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(127)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GXParserSTEP_ {
			{
				p.SetState(125)
				p.Match(GXParserSTEP_)
			}
			{
				p.SetState(126)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*ForBlockContext).cada = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == GXParserVARIABLE || _la == GXParserDecimalLiteral) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*ForBlockContext).cada = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		p.SetState(140)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2260605218193788) != 0 || _la == GXParserATRIBUTO || _la == GXParserATRIBUTOVAR {
			p.SetState(138)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(129)
					p.Statement()
				}

			case 2:
				{
					p.SetState(130)
					p.PrintStatement()
				}

			case 3:
				{
					p.SetState(131)
					p.NewBlock()
				}

			case 4:
				{
					p.SetState(132)
					p.ForBlock()
				}

			case 5:
				{
					p.SetState(133)
					p.IfBlock()
				}

			case 6:
				{
					p.SetState(134)
					p.DoCase()
				}

			case 7:
				{
					p.SetState(135)
					p.DoWhile()
				}

			case 8:
				{
					p.SetState(136)
					p.DocLine()
				}

			case 9:
				{
					p.SetState(137)
					p.Match(GXParserCOMMENT)
				}

			}

			p.SetState(142)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(143)
			p.Match(GXParserENDFOR_)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(145)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GXParserDocLineInfoPre {
			{
				p.SetState(144)

				var _m = p.Match(GXParserDocLineInfoPre)

				localctx.(*ForBlockContext).doc = _m
			}

		}
		{
			p.SetState(147)
			p.Match(GXParserFOR_)
		}
		{
			p.SetState(148)
			p.Match(GXParserVARIABLE)
		}
		{
			p.SetState(149)
			p.Match(GXParserIN_)
		}
		{
			p.SetState(150)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ForBlockContext).sdt = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == GXParserVARIABLE || _la == GXParserATRIBUTOVAR) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ForBlockContext).sdt = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(161)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2260596628259196) != 0 || _la == GXParserATRIBUTO || _la == GXParserATRIBUTOVAR {
			p.SetState(159)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(151)
					p.Statement()
				}

			case 2:
				{
					p.SetState(152)
					p.NewBlock()
				}

			case 3:
				{
					p.SetState(153)
					p.ForBlock()
				}

			case 4:
				{
					p.SetState(154)
					p.IfBlock()
				}

			case 5:
				{
					p.SetState(155)
					p.DoCase()
				}

			case 6:
				{
					p.SetState(156)
					p.DoWhile()
				}

			case 7:
				{
					p.SetState(157)
					p.DocLine()
				}

			case 8:
				{
					p.SetState(158)
					p.Match(GXParserCOMMENT)
				}

			}

			p.SetState(163)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(164)
			p.Match(GXParserENDFOR_)
		}

	}

	return localctx
}

// IWhereConditionContext is an interface to support dynamic dispatch.
type IWhereConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCondicion returns the condicion rule contexts.
	GetCondicion() IConditionContext

	// SetCondicion sets the condicion rule contexts.
	SetCondicion(IConditionContext)

	// IsWhereConditionContext differentiates from other interfaces.
	IsWhereConditionContext()
}

type WhereConditionContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	condicion IConditionContext
}

func NewEmptyWhereConditionContext() *WhereConditionContext {
	var p = new(WhereConditionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GXParserRULE_whereCondition
	return p
}

func (*WhereConditionContext) IsWhereConditionContext() {}

func NewWhereConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhereConditionContext {
	var p = new(WhereConditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GXParserRULE_whereCondition

	return p
}

func (s *WhereConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *WhereConditionContext) GetCondicion() IConditionContext { return s.condicion }

func (s *WhereConditionContext) SetCondicion(v IConditionContext) { s.condicion = v }

func (s *WhereConditionContext) WHERE_() antlr.TerminalNode {
	return s.GetToken(GXParserWHERE_, 0)
}

func (s *WhereConditionContext) AllCondition() []IConditionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConditionContext); ok {
			len++
		}
	}

	tst := make([]IConditionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConditionContext); ok {
			tst[i] = t.(IConditionContext)
			i++
		}
	}

	return tst
}

func (s *WhereConditionContext) Condition(i int) IConditionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *WhereConditionContext) COMENTARIO() antlr.TerminalNode {
	return s.GetToken(GXParserCOMENTARIO, 0)
}

func (s *WhereConditionContext) COMMENT() antlr.TerminalNode {
	return s.GetToken(GXParserCOMMENT, 0)
}

func (s *WhereConditionContext) WHEN_() antlr.TerminalNode {
	return s.GetToken(GXParserWHEN_, 0)
}

func (s *WhereConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhereConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhereConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GXParserListener); ok {
		listenerT.EnterWhereCondition(s)
	}
}

func (s *WhereConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GXParserListener); ok {
		listenerT.ExitWhereCondition(s)
	}
}

func (p *GXParser) WhereCondition() (localctx IWhereConditionContext) {
	this := p
	_ = this

	localctx = NewWhereConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GXParserRULE_whereCondition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(167)
			p.Match(GXParserWHERE_)
		}
		{
			p.SetState(168)

			var _x = p.condition(0)

			localctx.(*WhereConditionContext).condicion = _x
		}
		p.SetState(170)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(169)
				_la = p.GetTokenStream().LA(1)

				if !(_la == GXParserCOMMENT || _la == GXParserCOMENTARIO) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(172)
			p.Match(GXParserWHERE_)
		}
		{
			p.SetState(173)

			var _x = p.condition(0)

			localctx.(*WhereConditionContext).condicion = _x
		}
		{
			p.SetState(174)
			p.Match(GXParserWHEN_)
		}
		{
			p.SetState(175)
			p.condition(0)
		}
		p.SetState(177)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(176)
				_la = p.GetTokenStream().LA(1)

				if !(_la == GXParserCOMMENT || _la == GXParserCOMENTARIO) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}

	}

	return localctx
}

// IIfBlockContext is an interface to support dynamic dispatch.
type IIfBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetComentario returns the comentario token.
	GetComentario() antlr.Token

	// SetComentario sets the comentario token.
	SetComentario(antlr.Token)

	// GetCondicion returns the condicion rule contexts.
	GetCondicion() IConditionContext

	// SetCondicion sets the condicion rule contexts.
	SetCondicion(IConditionContext)

	// IsIfBlockContext differentiates from other interfaces.
	IsIfBlockContext()
}

type IfBlockContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	condicion  IConditionContext
	comentario antlr.Token
}

func NewEmptyIfBlockContext() *IfBlockContext {
	var p = new(IfBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GXParserRULE_ifBlock
	return p
}

func (*IfBlockContext) IsIfBlockContext() {}

func NewIfBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfBlockContext {
	var p = new(IfBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GXParserRULE_ifBlock

	return p
}

func (s *IfBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *IfBlockContext) GetComentario() antlr.Token { return s.comentario }

func (s *IfBlockContext) SetComentario(v antlr.Token) { s.comentario = v }

func (s *IfBlockContext) GetCondicion() IConditionContext { return s.condicion }

func (s *IfBlockContext) SetCondicion(v IConditionContext) { s.condicion = v }

func (s *IfBlockContext) IF_() antlr.TerminalNode {
	return s.GetToken(GXParserIF_, 0)
}

func (s *IfBlockContext) ENDIF_() antlr.TerminalNode {
	return s.GetToken(GXParserENDIF_, 0)
}

func (s *IfBlockContext) Condition() IConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *IfBlockContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *IfBlockContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *IfBlockContext) AllPrintStatement() []IPrintStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPrintStatementContext); ok {
			len++
		}
	}

	tst := make([]IPrintStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPrintStatementContext); ok {
			tst[i] = t.(IPrintStatementContext)
			i++
		}
	}

	return tst
}

func (s *IfBlockContext) PrintStatement(i int) IPrintStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintStatementContext)
}

func (s *IfBlockContext) AllNewBlock() []INewBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INewBlockContext); ok {
			len++
		}
	}

	tst := make([]INewBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INewBlockContext); ok {
			tst[i] = t.(INewBlockContext)
			i++
		}
	}

	return tst
}

func (s *IfBlockContext) NewBlock(i int) INewBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INewBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INewBlockContext)
}

func (s *IfBlockContext) AllForBlock() []IForBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IForBlockContext); ok {
			len++
		}
	}

	tst := make([]IForBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IForBlockContext); ok {
			tst[i] = t.(IForBlockContext)
			i++
		}
	}

	return tst
}

func (s *IfBlockContext) ForBlock(i int) IForBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForBlockContext)
}

func (s *IfBlockContext) AllIfBlock() []IIfBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIfBlockContext); ok {
			len++
		}
	}

	tst := make([]IIfBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIfBlockContext); ok {
			tst[i] = t.(IIfBlockContext)
			i++
		}
	}

	return tst
}

func (s *IfBlockContext) IfBlock(i int) IIfBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfBlockContext)
}

func (s *IfBlockContext) AllDoCase() []IDoCaseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDoCaseContext); ok {
			len++
		}
	}

	tst := make([]IDoCaseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDoCaseContext); ok {
			tst[i] = t.(IDoCaseContext)
			i++
		}
	}

	return tst
}

func (s *IfBlockContext) DoCase(i int) IDoCaseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDoCaseContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDoCaseContext)
}

func (s *IfBlockContext) AllDoWhile() []IDoWhileContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDoWhileContext); ok {
			len++
		}
	}

	tst := make([]IDoWhileContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDoWhileContext); ok {
			tst[i] = t.(IDoWhileContext)
			i++
		}
	}

	return tst
}

func (s *IfBlockContext) DoWhile(i int) IDoWhileContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDoWhileContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDoWhileContext)
}

func (s *IfBlockContext) AllDocLine() []IDocLineContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDocLineContext); ok {
			len++
		}
	}

	tst := make([]IDocLineContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDocLineContext); ok {
			tst[i] = t.(IDocLineContext)
			i++
		}
	}

	return tst
}

func (s *IfBlockContext) DocLine(i int) IDocLineContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDocLineContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDocLineContext)
}

func (s *IfBlockContext) AllCOMMENT() []antlr.TerminalNode {
	return s.GetTokens(GXParserCOMMENT)
}

func (s *IfBlockContext) COMMENT(i int) antlr.TerminalNode {
	return s.GetToken(GXParserCOMMENT, i)
}

func (s *IfBlockContext) ELSE_() antlr.TerminalNode {
	return s.GetToken(GXParserELSE_, 0)
}

func (s *IfBlockContext) DocLineInfoPre() antlr.TerminalNode {
	return s.GetToken(GXParserDocLineInfoPre, 0)
}

func (s *IfBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GXParserListener); ok {
		listenerT.EnterIfBlock(s)
	}
}

func (s *IfBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GXParserListener); ok {
		listenerT.ExitIfBlock(s)
	}
}

func (p *GXParser) IfBlock() (localctx IIfBlockContext) {
	this := p
	_ = this

	localctx = NewIfBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, GXParserRULE_ifBlock)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(181)
		p.Match(GXParserIF_)
	}
	{
		p.SetState(182)

		var _x = p.condition(0)

		localctx.(*IfBlockContext).condicion = _x
	}
	p.SetState(184)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(183)

			var _m = p.Match(GXParserDocLineInfoPre)

			localctx.(*IfBlockContext).comentario = _m
		}

	}
	p.SetState(197)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2260605218193788) != 0 || _la == GXParserATRIBUTO || _la == GXParserATRIBUTOVAR {
		p.SetState(195)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(186)
				p.Statement()
			}

		case 2:
			{
				p.SetState(187)
				p.PrintStatement()
			}

		case 3:
			{
				p.SetState(188)
				p.NewBlock()
			}

		case 4:
			{
				p.SetState(189)
				p.ForBlock()
			}

		case 5:
			{
				p.SetState(190)
				p.IfBlock()
			}

		case 6:
			{
				p.SetState(191)
				p.DoCase()
			}

		case 7:
			{
				p.SetState(192)
				p.DoWhile()
			}

		case 8:
			{
				p.SetState(193)
				p.DocLine()
			}

		case 9:
			{
				p.SetState(194)
				p.Match(GXParserCOMMENT)
			}

		}

		p.SetState(199)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(215)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GXParserELSE_ {
		{
			p.SetState(200)
			p.Match(GXParserELSE_)
		}
		p.SetState(212)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2260605218193788) != 0 || _la == GXParserATRIBUTO || _la == GXParserATRIBUTOVAR {
			p.SetState(210)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(201)
					p.Statement()
				}

			case 2:
				{
					p.SetState(202)
					p.PrintStatement()
				}

			case 3:
				{
					p.SetState(203)
					p.NewBlock()
				}

			case 4:
				{
					p.SetState(204)
					p.ForBlock()
				}

			case 5:
				{
					p.SetState(205)
					p.IfBlock()
				}

			case 6:
				{
					p.SetState(206)
					p.DoCase()
				}

			case 7:
				{
					p.SetState(207)
					p.DoWhile()
				}

			case 8:
				{
					p.SetState(208)
					p.DocLine()
				}

			case 9:
				{
					p.SetState(209)
					p.Match(GXParserCOMMENT)
				}

			}

			p.SetState(214)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(217)
		p.Match(GXParserENDIF_)
	}

	return localctx
}

// IDoCaseContext is an interface to support dynamic dispatch.
type IDoCaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDoCaseContext differentiates from other interfaces.
	IsDoCaseContext()
}

type DoCaseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDoCaseContext() *DoCaseContext {
	var p = new(DoCaseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GXParserRULE_doCase
	return p
}

func (*DoCaseContext) IsDoCaseContext() {}

func NewDoCaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DoCaseContext {
	var p = new(DoCaseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GXParserRULE_doCase

	return p
}

func (s *DoCaseContext) GetParser() antlr.Parser { return s.parser }

func (s *DoCaseContext) DO_() antlr.TerminalNode {
	return s.GetToken(GXParserDO_, 0)
}

func (s *DoCaseContext) AllCASE_() []antlr.TerminalNode {
	return s.GetTokens(GXParserCASE_)
}

func (s *DoCaseContext) CASE_(i int) antlr.TerminalNode {
	return s.GetToken(GXParserCASE_, i)
}

func (s *DoCaseContext) ENDCASE_() antlr.TerminalNode {
	return s.GetToken(GXParserENDCASE_, 0)
}

func (s *DoCaseContext) AllCOMENTARIO() []antlr.TerminalNode {
	return s.GetTokens(GXParserCOMENTARIO)
}

func (s *DoCaseContext) COMENTARIO(i int) antlr.TerminalNode {
	return s.GetToken(GXParserCOMENTARIO, i)
}

func (s *DoCaseContext) AllCondition() []IConditionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConditionContext); ok {
			len++
		}
	}

	tst := make([]IConditionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConditionContext); ok {
			tst[i] = t.(IConditionContext)
			i++
		}
	}

	return tst
}

func (s *DoCaseContext) Condition(i int) IConditionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *DoCaseContext) OTHERWISE_() antlr.TerminalNode {
	return s.GetToken(GXParserOTHERWISE_, 0)
}

func (s *DoCaseContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *DoCaseContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *DoCaseContext) AllNewBlock() []INewBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INewBlockContext); ok {
			len++
		}
	}

	tst := make([]INewBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INewBlockContext); ok {
			tst[i] = t.(INewBlockContext)
			i++
		}
	}

	return tst
}

func (s *DoCaseContext) NewBlock(i int) INewBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INewBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INewBlockContext)
}

func (s *DoCaseContext) AllForBlock() []IForBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IForBlockContext); ok {
			len++
		}
	}

	tst := make([]IForBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IForBlockContext); ok {
			tst[i] = t.(IForBlockContext)
			i++
		}
	}

	return tst
}

func (s *DoCaseContext) ForBlock(i int) IForBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForBlockContext)
}

func (s *DoCaseContext) AllIfBlock() []IIfBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIfBlockContext); ok {
			len++
		}
	}

	tst := make([]IIfBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIfBlockContext); ok {
			tst[i] = t.(IIfBlockContext)
			i++
		}
	}

	return tst
}

func (s *DoCaseContext) IfBlock(i int) IIfBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfBlockContext)
}

func (s *DoCaseContext) AllDoCase() []IDoCaseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDoCaseContext); ok {
			len++
		}
	}

	tst := make([]IDoCaseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDoCaseContext); ok {
			tst[i] = t.(IDoCaseContext)
			i++
		}
	}

	return tst
}

func (s *DoCaseContext) DoCase(i int) IDoCaseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDoCaseContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDoCaseContext)
}

func (s *DoCaseContext) AllDoWhile() []IDoWhileContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDoWhileContext); ok {
			len++
		}
	}

	tst := make([]IDoWhileContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDoWhileContext); ok {
			tst[i] = t.(IDoWhileContext)
			i++
		}
	}

	return tst
}

func (s *DoCaseContext) DoWhile(i int) IDoWhileContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDoWhileContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDoWhileContext)
}

func (s *DoCaseContext) AllDocLine() []IDocLineContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDocLineContext); ok {
			len++
		}
	}

	tst := make([]IDocLineContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDocLineContext); ok {
			tst[i] = t.(IDocLineContext)
			i++
		}
	}

	return tst
}

func (s *DoCaseContext) DocLine(i int) IDocLineContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDocLineContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDocLineContext)
}

func (s *DoCaseContext) AllCOMMENT() []antlr.TerminalNode {
	return s.GetTokens(GXParserCOMMENT)
}

func (s *DoCaseContext) COMMENT(i int) antlr.TerminalNode {
	return s.GetToken(GXParserCOMMENT, i)
}

func (s *DoCaseContext) AllPrintStatement() []IPrintStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPrintStatementContext); ok {
			len++
		}
	}

	tst := make([]IPrintStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPrintStatementContext); ok {
			tst[i] = t.(IPrintStatementContext)
			i++
		}
	}

	return tst
}

func (s *DoCaseContext) PrintStatement(i int) IPrintStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintStatementContext)
}

func (s *DoCaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DoCaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DoCaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GXParserListener); ok {
		listenerT.EnterDoCase(s)
	}
}

func (s *DoCaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GXParserListener); ok {
		listenerT.ExitDoCase(s)
	}
}

func (p *GXParser) DoCase() (localctx IDoCaseContext) {
	this := p
	_ = this

	localctx = NewDoCaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, GXParserRULE_doCase)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(219)
		p.Match(GXParserDO_)
	}
	{
		p.SetState(220)
		p.Match(GXParserCASE_)
	}
	p.SetState(224)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GXParserCOMENTARIO {
		{
			p.SetState(221)
			p.Match(GXParserCOMENTARIO)
		}

		p.SetState(226)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(244)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GXParserCASE_ {
		{
			p.SetState(227)
			p.Match(GXParserCASE_)
		}
		{
			p.SetState(228)
			p.condition(0)
		}
		p.SetState(239)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2260596628259196) != 0 || _la == GXParserATRIBUTO || _la == GXParserATRIBUTOVAR {
			p.SetState(237)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(229)
					p.Statement()
				}

			case 2:
				{
					p.SetState(230)
					p.NewBlock()
				}

			case 3:
				{
					p.SetState(231)
					p.ForBlock()
				}

			case 4:
				{
					p.SetState(232)
					p.IfBlock()
				}

			case 5:
				{
					p.SetState(233)
					p.DoCase()
				}

			case 6:
				{
					p.SetState(234)
					p.DoWhile()
				}

			case 7:
				{
					p.SetState(235)
					p.DocLine()
				}

			case 8:
				{
					p.SetState(236)
					p.Match(GXParserCOMMENT)
				}

			}

			p.SetState(241)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(246)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(262)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GXParserOTHERWISE_ {
		{
			p.SetState(247)
			p.Match(GXParserOTHERWISE_)
		}
		p.SetState(259)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2260605218193788) != 0 || _la == GXParserATRIBUTO || _la == GXParserATRIBUTOVAR {
			p.SetState(257)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(248)
					p.Statement()
				}

			case 2:
				{
					p.SetState(249)
					p.PrintStatement()
				}

			case 3:
				{
					p.SetState(250)
					p.NewBlock()
				}

			case 4:
				{
					p.SetState(251)
					p.ForBlock()
				}

			case 5:
				{
					p.SetState(252)
					p.IfBlock()
				}

			case 6:
				{
					p.SetState(253)
					p.DoCase()
				}

			case 7:
				{
					p.SetState(254)
					p.DoWhile()
				}

			case 8:
				{
					p.SetState(255)
					p.DocLine()
				}

			case 9:
				{
					p.SetState(256)
					p.Match(GXParserCOMMENT)
				}

			}

			p.SetState(261)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(264)
		p.Match(GXParserENDCASE_)
	}

	return localctx
}

// IDoWhileContext is an interface to support dynamic dispatch.
type IDoWhileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetComentario returns the comentario token.
	GetComentario() antlr.Token

	// SetComentario sets the comentario token.
	SetComentario(antlr.Token)

	// GetCondicion returns the condicion rule contexts.
	GetCondicion() IConditionContext

	// SetCondicion sets the condicion rule contexts.
	SetCondicion(IConditionContext)

	// IsDoWhileContext differentiates from other interfaces.
	IsDoWhileContext()
}

type DoWhileContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	condicion  IConditionContext
	comentario antlr.Token
}

func NewEmptyDoWhileContext() *DoWhileContext {
	var p = new(DoWhileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GXParserRULE_doWhile
	return p
}

func (*DoWhileContext) IsDoWhileContext() {}

func NewDoWhileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DoWhileContext {
	var p = new(DoWhileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GXParserRULE_doWhile

	return p
}

func (s *DoWhileContext) GetParser() antlr.Parser { return s.parser }

func (s *DoWhileContext) GetComentario() antlr.Token { return s.comentario }

func (s *DoWhileContext) SetComentario(v antlr.Token) { s.comentario = v }

func (s *DoWhileContext) GetCondicion() IConditionContext { return s.condicion }

func (s *DoWhileContext) SetCondicion(v IConditionContext) { s.condicion = v }

func (s *DoWhileContext) DO_() antlr.TerminalNode {
	return s.GetToken(GXParserDO_, 0)
}

func (s *DoWhileContext) WHILE_() antlr.TerminalNode {
	return s.GetToken(GXParserWHILE_, 0)
}

func (s *DoWhileContext) ENDDO_() antlr.TerminalNode {
	return s.GetToken(GXParserENDDO_, 0)
}

func (s *DoWhileContext) Condition() IConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *DoWhileContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *DoWhileContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *DoWhileContext) AllPrintStatement() []IPrintStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPrintStatementContext); ok {
			len++
		}
	}

	tst := make([]IPrintStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPrintStatementContext); ok {
			tst[i] = t.(IPrintStatementContext)
			i++
		}
	}

	return tst
}

func (s *DoWhileContext) PrintStatement(i int) IPrintStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintStatementContext)
}

func (s *DoWhileContext) AllNewBlock() []INewBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INewBlockContext); ok {
			len++
		}
	}

	tst := make([]INewBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INewBlockContext); ok {
			tst[i] = t.(INewBlockContext)
			i++
		}
	}

	return tst
}

func (s *DoWhileContext) NewBlock(i int) INewBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INewBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INewBlockContext)
}

func (s *DoWhileContext) AllForBlock() []IForBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IForBlockContext); ok {
			len++
		}
	}

	tst := make([]IForBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IForBlockContext); ok {
			tst[i] = t.(IForBlockContext)
			i++
		}
	}

	return tst
}

func (s *DoWhileContext) ForBlock(i int) IForBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForBlockContext)
}

func (s *DoWhileContext) AllIfBlock() []IIfBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIfBlockContext); ok {
			len++
		}
	}

	tst := make([]IIfBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIfBlockContext); ok {
			tst[i] = t.(IIfBlockContext)
			i++
		}
	}

	return tst
}

func (s *DoWhileContext) IfBlock(i int) IIfBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfBlockContext)
}

func (s *DoWhileContext) AllDoCase() []IDoCaseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDoCaseContext); ok {
			len++
		}
	}

	tst := make([]IDoCaseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDoCaseContext); ok {
			tst[i] = t.(IDoCaseContext)
			i++
		}
	}

	return tst
}

func (s *DoWhileContext) DoCase(i int) IDoCaseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDoCaseContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDoCaseContext)
}

func (s *DoWhileContext) AllDoWhile() []IDoWhileContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDoWhileContext); ok {
			len++
		}
	}

	tst := make([]IDoWhileContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDoWhileContext); ok {
			tst[i] = t.(IDoWhileContext)
			i++
		}
	}

	return tst
}

func (s *DoWhileContext) DoWhile(i int) IDoWhileContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDoWhileContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDoWhileContext)
}

func (s *DoWhileContext) AllDocLine() []IDocLineContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDocLineContext); ok {
			len++
		}
	}

	tst := make([]IDocLineContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDocLineContext); ok {
			tst[i] = t.(IDocLineContext)
			i++
		}
	}

	return tst
}

func (s *DoWhileContext) DocLine(i int) IDocLineContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDocLineContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDocLineContext)
}

func (s *DoWhileContext) AllCOMMENT() []antlr.TerminalNode {
	return s.GetTokens(GXParserCOMMENT)
}

func (s *DoWhileContext) COMMENT(i int) antlr.TerminalNode {
	return s.GetToken(GXParserCOMMENT, i)
}

func (s *DoWhileContext) DocLineInfoPre() antlr.TerminalNode {
	return s.GetToken(GXParserDocLineInfoPre, 0)
}

func (s *DoWhileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DoWhileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DoWhileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GXParserListener); ok {
		listenerT.EnterDoWhile(s)
	}
}

func (s *DoWhileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GXParserListener); ok {
		listenerT.ExitDoWhile(s)
	}
}

func (p *GXParser) DoWhile() (localctx IDoWhileContext) {
	this := p
	_ = this

	localctx = NewDoWhileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GXParserRULE_doWhile)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(266)
		p.Match(GXParserDO_)
	}
	{
		p.SetState(267)
		p.Match(GXParserWHILE_)
	}
	{
		p.SetState(268)

		var _x = p.condition(0)

		localctx.(*DoWhileContext).condicion = _x
	}
	p.SetState(270)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(269)

			var _m = p.Match(GXParserDocLineInfoPre)

			localctx.(*DoWhileContext).comentario = _m
		}

	}
	p.SetState(283)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2260605218193788) != 0 || _la == GXParserATRIBUTO || _la == GXParserATRIBUTOVAR {
		p.SetState(281)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(272)
				p.Statement()
			}

		case 2:
			{
				p.SetState(273)
				p.PrintStatement()
			}

		case 3:
			{
				p.SetState(274)
				p.NewBlock()
			}

		case 4:
			{
				p.SetState(275)
				p.ForBlock()
			}

		case 5:
			{
				p.SetState(276)
				p.IfBlock()
			}

		case 6:
			{
				p.SetState(277)
				p.DoCase()
			}

		case 7:
			{
				p.SetState(278)
				p.DoWhile()
			}

		case 8:
			{
				p.SetState(279)
				p.DocLine()
			}

		case 9:
			{
				p.SetState(280)
				p.Match(GXParserCOMMENT)
			}

		}

		p.SetState(285)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(286)
		p.Match(GXParserENDDO_)
	}

	return localctx
}

// IConditionContext is an interface to support dynamic dispatch.
type IConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_singleExpression returns the _singleExpression rule contexts.
	Get_singleExpression() ISingleExpressionContext

	// Set_singleExpression sets the _singleExpression rule contexts.
	Set_singleExpression(ISingleExpressionContext)

	// GetExpresions returns the expresions rule context list.
	GetExpresions() []ISingleExpressionContext

	// SetExpresions sets the expresions rule context list.
	SetExpresions([]ISingleExpressionContext)

	// IsConditionContext differentiates from other interfaces.
	IsConditionContext()
}

type ConditionContext struct {
	*antlr.BaseParserRuleContext
	parser            antlr.Parser
	_singleExpression ISingleExpressionContext
	expresions        []ISingleExpressionContext
}

func NewEmptyConditionContext() *ConditionContext {
	var p = new(ConditionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GXParserRULE_condition
	return p
}

func (*ConditionContext) IsConditionContext() {}

func NewConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionContext {
	var p = new(ConditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GXParserRULE_condition

	return p
}

func (s *ConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionContext) Get_singleExpression() ISingleExpressionContext {
	return s._singleExpression
}

func (s *ConditionContext) Set_singleExpression(v ISingleExpressionContext) { s._singleExpression = v }

func (s *ConditionContext) GetExpresions() []ISingleExpressionContext { return s.expresions }

func (s *ConditionContext) SetExpresions(v []ISingleExpressionContext) { s.expresions = v }

func (s *ConditionContext) AllSingleExpression() []ISingleExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISingleExpressionContext); ok {
			len++
		}
	}

	tst := make([]ISingleExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISingleExpressionContext); ok {
			tst[i] = t.(ISingleExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ConditionContext) SingleExpression(i int) ISingleExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *ConditionContext) COMPARISON() antlr.TerminalNode {
	return s.GetToken(GXParserCOMPARISON, 0)
}

func (s *ConditionContext) Assign() antlr.TerminalNode {
	return s.GetToken(GXParserAssign, 0)
}

func (s *ConditionContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(GXParserOpenParen, 0)
}

func (s *ConditionContext) AllCondition() []IConditionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConditionContext); ok {
			len++
		}
	}

	tst := make([]IConditionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConditionContext); ok {
			tst[i] = t.(IConditionContext)
			i++
		}
	}

	return tst
}

func (s *ConditionContext) Condition(i int) IConditionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *ConditionContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(GXParserCloseParen, 0)
}

func (s *ConditionContext) AND_() antlr.TerminalNode {
	return s.GetToken(GXParserAND_, 0)
}

func (s *ConditionContext) OR_() antlr.TerminalNode {
	return s.GetToken(GXParserOR_, 0)
}

func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GXParserListener); ok {
		listenerT.EnterCondition(s)
	}
}

func (s *ConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GXParserListener); ok {
		listenerT.ExitCondition(s)
	}
}

func (p *GXParser) Condition() (localctx IConditionContext) {
	return p.condition(0)
}

func (p *GXParser) condition(_p int) (localctx IConditionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewConditionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IConditionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 14
	p.EnterRecursionRule(localctx, 14, GXParserRULE_condition, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(297)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 42, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(289)

			var _x = p.singleExpression(0)

			localctx.(*ConditionContext)._singleExpression = _x
		}
		localctx.(*ConditionContext).expresions = append(localctx.(*ConditionContext).expresions, localctx.(*ConditionContext)._singleExpression)
		{
			p.SetState(290)
			_la = p.GetTokenStream().LA(1)

			if !(_la == GXParserCOMPARISON || _la == GXParserAssign) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(291)

			var _x = p.singleExpression(0)

			localctx.(*ConditionContext)._singleExpression = _x
		}
		localctx.(*ConditionContext).expresions = append(localctx.(*ConditionContext).expresions, localctx.(*ConditionContext)._singleExpression)

	case 2:
		{
			p.SetState(293)
			p.Match(GXParserOpenParen)
		}
		{
			p.SetState(294)
			p.condition(0)
		}
		{
			p.SetState(295)
			p.Match(GXParserCloseParen)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(304)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewConditionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, GXParserRULE_condition)
			p.SetState(299)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(300)
				_la = p.GetTokenStream().LA(1)

				if !(_la == GXParserAND_ || _la == GXParserOR_) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(301)
				p.condition(3)
			}

		}
		p.SetState(306)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext())
	}

	return localctx
}

// ISubrutineContext is an interface to support dynamic dispatch.
type ISubrutineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNombre returns the nombre token.
	GetNombre() antlr.Token

	// SetNombre sets the nombre token.
	SetNombre(antlr.Token)

	// IsSubrutineContext differentiates from other interfaces.
	IsSubrutineContext()
}

type SubrutineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	nombre antlr.Token
}

func NewEmptySubrutineContext() *SubrutineContext {
	var p = new(SubrutineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GXParserRULE_subrutine
	return p
}

func (*SubrutineContext) IsSubrutineContext() {}

func NewSubrutineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubrutineContext {
	var p = new(SubrutineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GXParserRULE_subrutine

	return p
}

func (s *SubrutineContext) GetParser() antlr.Parser { return s.parser }

func (s *SubrutineContext) GetNombre() antlr.Token { return s.nombre }

func (s *SubrutineContext) SetNombre(v antlr.Token) { s.nombre = v }

func (s *SubrutineContext) SUB_() antlr.TerminalNode {
	return s.GetToken(GXParserSUB_, 0)
}

func (s *SubrutineContext) ENDSUB_() antlr.TerminalNode {
	return s.GetToken(GXParserENDSUB_, 0)
}

func (s *SubrutineContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(GXParserStringLiteral, 0)
}

func (s *SubrutineContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *SubrutineContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *SubrutineContext) AllPrintStatement() []IPrintStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPrintStatementContext); ok {
			len++
		}
	}

	tst := make([]IPrintStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPrintStatementContext); ok {
			tst[i] = t.(IPrintStatementContext)
			i++
		}
	}

	return tst
}

func (s *SubrutineContext) PrintStatement(i int) IPrintStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintStatementContext)
}

func (s *SubrutineContext) AllNewBlock() []INewBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INewBlockContext); ok {
			len++
		}
	}

	tst := make([]INewBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INewBlockContext); ok {
			tst[i] = t.(INewBlockContext)
			i++
		}
	}

	return tst
}

func (s *SubrutineContext) NewBlock(i int) INewBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INewBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INewBlockContext)
}

func (s *SubrutineContext) AllForBlock() []IForBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IForBlockContext); ok {
			len++
		}
	}

	tst := make([]IForBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IForBlockContext); ok {
			tst[i] = t.(IForBlockContext)
			i++
		}
	}

	return tst
}

func (s *SubrutineContext) ForBlock(i int) IForBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForBlockContext)
}

func (s *SubrutineContext) AllIfBlock() []IIfBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIfBlockContext); ok {
			len++
		}
	}

	tst := make([]IIfBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIfBlockContext); ok {
			tst[i] = t.(IIfBlockContext)
			i++
		}
	}

	return tst
}

func (s *SubrutineContext) IfBlock(i int) IIfBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfBlockContext)
}

func (s *SubrutineContext) AllDoCase() []IDoCaseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDoCaseContext); ok {
			len++
		}
	}

	tst := make([]IDoCaseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDoCaseContext); ok {
			tst[i] = t.(IDoCaseContext)
			i++
		}
	}

	return tst
}

func (s *SubrutineContext) DoCase(i int) IDoCaseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDoCaseContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDoCaseContext)
}

func (s *SubrutineContext) AllDoWhile() []IDoWhileContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDoWhileContext); ok {
			len++
		}
	}

	tst := make([]IDoWhileContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDoWhileContext); ok {
			tst[i] = t.(IDoWhileContext)
			i++
		}
	}

	return tst
}

func (s *SubrutineContext) DoWhile(i int) IDoWhileContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDoWhileContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDoWhileContext)
}

func (s *SubrutineContext) AllDocLine() []IDocLineContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDocLineContext); ok {
			len++
		}
	}

	tst := make([]IDocLineContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDocLineContext); ok {
			tst[i] = t.(IDocLineContext)
			i++
		}
	}

	return tst
}

func (s *SubrutineContext) DocLine(i int) IDocLineContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDocLineContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDocLineContext)
}

func (s *SubrutineContext) AllCOMMENT() []antlr.TerminalNode {
	return s.GetTokens(GXParserCOMMENT)
}

func (s *SubrutineContext) COMMENT(i int) antlr.TerminalNode {
	return s.GetToken(GXParserCOMMENT, i)
}

func (s *SubrutineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubrutineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubrutineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GXParserListener); ok {
		listenerT.EnterSubrutine(s)
	}
}

func (s *SubrutineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GXParserListener); ok {
		listenerT.ExitSubrutine(s)
	}
}

func (p *GXParser) Subrutine() (localctx ISubrutineContext) {
	this := p
	_ = this

	localctx = NewSubrutineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GXParserRULE_subrutine)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(307)
		p.Match(GXParserSUB_)
	}
	{
		p.SetState(308)

		var _m = p.Match(GXParserStringLiteral)

		localctx.(*SubrutineContext).nombre = _m
	}
	p.SetState(320)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2260605218193788) != 0 || _la == GXParserATRIBUTO || _la == GXParserATRIBUTOVAR {
		p.SetState(318)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(309)
				p.Statement()
			}

		case 2:
			{
				p.SetState(310)
				p.PrintStatement()
			}

		case 3:
			{
				p.SetState(311)
				p.NewBlock()
			}

		case 4:
			{
				p.SetState(312)
				p.ForBlock()
			}

		case 5:
			{
				p.SetState(313)
				p.IfBlock()
			}

		case 6:
			{
				p.SetState(314)
				p.DoCase()
			}

		case 7:
			{
				p.SetState(315)
				p.DoWhile()
			}

		case 8:
			{
				p.SetState(316)
				p.DocLine()
			}

		case 9:
			{
				p.SetState(317)
				p.Match(GXParserCOMMENT)
			}

		}

		p.SetState(322)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(323)
		p.Match(GXParserENDSUB_)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetVariable returns the variable token.
	GetVariable() antlr.Token

	// SetVariable sets the variable token.
	SetVariable(antlr.Token)

	// GetExpresion returns the expresion rule contexts.
	GetExpresion() ISingleExpressionContext

	// SetExpresion sets the expresion rule contexts.
	SetExpresion(ISingleExpressionContext)

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	variable  antlr.Token
	expresion ISingleExpressionContext
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GXParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GXParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) GetVariable() antlr.Token { return s.variable }

func (s *StatementContext) SetVariable(v antlr.Token) { s.variable = v }

func (s *StatementContext) GetExpresion() ISingleExpressionContext { return s.expresion }

func (s *StatementContext) SetExpresion(v ISingleExpressionContext) { s.expresion = v }

func (s *StatementContext) Funcion() IFuncionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncionContext)
}

func (s *StatementContext) VARIABLE() antlr.TerminalNode {
	return s.GetToken(GXParserVARIABLE, 0)
}

func (s *StatementContext) Assign() antlr.TerminalNode {
	return s.GetToken(GXParserAssign, 0)
}

func (s *StatementContext) Plus() antlr.TerminalNode {
	return s.GetToken(GXParserPlus, 0)
}

func (s *StatementContext) Minus() antlr.TerminalNode {
	return s.GetToken(GXParserMinus, 0)
}

func (s *StatementContext) SingleExpression() ISingleExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *StatementContext) ATRIBUTO() antlr.TerminalNode {
	return s.GetToken(GXParserATRIBUTO, 0)
}

func (s *StatementContext) ATRIBUTOVAR() antlr.TerminalNode {
	return s.GetToken(GXParserATRIBUTOVAR, 0)
}

func (s *StatementContext) DO_() antlr.TerminalNode {
	return s.GetToken(GXParserDO_, 0)
}

func (s *StatementContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(GXParserStringLiteral, 0)
}

func (s *StatementContext) EXIT_() antlr.TerminalNode {
	return s.GetToken(GXParserEXIT_, 0)
}

func (s *StatementContext) DELETE_() antlr.TerminalNode {
	return s.GetToken(GXParserDELETE_, 0)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GXParserListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GXParserListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *GXParser) Statement() (localctx IStatementContext) {
	this := p
	_ = this

	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, GXParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(366)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 52, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(325)
			p.Funcion()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(326)

			var _m = p.Match(GXParserVARIABLE)

			localctx.(*StatementContext).variable = _m
		}
		p.SetState(332)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case GXParserAssign:
			{
				p.SetState(327)
				p.Match(GXParserAssign)
			}

		case GXParserPlus:
			{
				p.SetState(328)
				p.Match(GXParserPlus)
			}
			{
				p.SetState(329)
				p.Match(GXParserAssign)
			}

		case GXParserMinus:
			{
				p.SetState(330)
				p.Match(GXParserMinus)
			}
			{
				p.SetState(331)
				p.Match(GXParserAssign)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		p.SetState(336)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 47, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(334)
				p.Funcion()
			}

		case 2:
			{
				p.SetState(335)

				var _x = p.singleExpression(0)

				localctx.(*StatementContext).expresion = _x
			}

		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(338)
			p.Match(GXParserATRIBUTO)
		}
		p.SetState(344)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case GXParserAssign:
			{
				p.SetState(339)
				p.Match(GXParserAssign)
			}

		case GXParserPlus:
			{
				p.SetState(340)
				p.Match(GXParserPlus)
			}
			{
				p.SetState(341)
				p.Match(GXParserAssign)
			}

		case GXParserMinus:
			{
				p.SetState(342)
				p.Match(GXParserMinus)
			}
			{
				p.SetState(343)
				p.Match(GXParserAssign)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		p.SetState(348)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 49, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(346)
				p.Funcion()
			}

		case 2:
			{
				p.SetState(347)
				p.singleExpression(0)
			}

		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(350)
			p.Match(GXParserATRIBUTOVAR)
		}
		p.SetState(356)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case GXParserAssign:
			{
				p.SetState(351)
				p.Match(GXParserAssign)
			}

		case GXParserPlus:
			{
				p.SetState(352)
				p.Match(GXParserPlus)
			}
			{
				p.SetState(353)
				p.Match(GXParserAssign)
			}

		case GXParserMinus:
			{
				p.SetState(354)
				p.Match(GXParserMinus)
			}
			{
				p.SetState(355)
				p.Match(GXParserAssign)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		p.SetState(360)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 51, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(358)
				p.Funcion()
			}

		case 2:
			{
				p.SetState(359)
				p.singleExpression(0)
			}

		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(362)
			p.Match(GXParserDO_)
		}
		{
			p.SetState(363)
			p.Match(GXParserStringLiteral)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(364)
			p.Match(GXParserEXIT_)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(365)
			p.Match(GXParserDELETE_)
		}

	}

	return localctx
}

// IFuncionContext is an interface to support dynamic dispatch.
type IFuncionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncionContext differentiates from other interfaces.
	IsFuncionContext()
}

type FuncionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncionContext() *FuncionContext {
	var p = new(FuncionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GXParserRULE_funcion
	return p
}

func (*FuncionContext) IsFuncionContext() {}

func NewFuncionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncionContext {
	var p = new(FuncionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GXParserRULE_funcion

	return p
}

func (s *FuncionContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncionContext) FUNCIONES() antlr.TerminalNode {
	return s.GetToken(GXParserFUNCIONES, 0)
}

func (s *FuncionContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(GXParserOpenParen, 0)
}

func (s *FuncionContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(GXParserCloseParen, 0)
}

func (s *FuncionContext) AllSingleExpression() []ISingleExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISingleExpressionContext); ok {
			len++
		}
	}

	tst := make([]ISingleExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISingleExpressionContext); ok {
			tst[i] = t.(ISingleExpressionContext)
			i++
		}
	}

	return tst
}

func (s *FuncionContext) SingleExpression(i int) ISingleExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *FuncionContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(GXParserComma)
}

func (s *FuncionContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(GXParserComma, i)
}

func (s *FuncionContext) ATRIBUTOVAR() antlr.TerminalNode {
	return s.GetToken(GXParserATRIBUTOVAR, 0)
}

func (s *FuncionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GXParserListener); ok {
		listenerT.EnterFuncion(s)
	}
}

func (s *FuncionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GXParserListener); ok {
		listenerT.ExitFuncion(s)
	}
}

func (p *GXParser) Funcion() (localctx IFuncionContext) {
	this := p
	_ = this

	localctx = NewFuncionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, GXParserRULE_funcion)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(387)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GXParserFUNCIONES:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(368)
			p.Match(GXParserFUNCIONES)
		}
		{
			p.SetState(369)
			p.Match(GXParserOpenParen)
		}

		p.SetState(371)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2251799847305984) != 0 || (int64((_la-163)) & ^0x3f) == 0 && ((int64(1)<<(_la-163))&11) != 0 {
			{
				p.SetState(370)
				p.singleExpression(0)
			}

		}
		p.SetState(377)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == GXParserComma {
			{
				p.SetState(373)
				p.Match(GXParserComma)
			}
			{
				p.SetState(374)
				p.singleExpression(0)
			}

			p.SetState(379)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

		{
			p.SetState(380)
			p.Match(GXParserCloseParen)
		}

	case GXParserATRIBUTOVAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(381)
			p.Match(GXParserATRIBUTOVAR)
		}
		{
			p.SetState(382)
			p.Match(GXParserOpenParen)
		}
		p.SetState(384)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2251799847305984) != 0 || (int64((_la-163)) & ^0x3f) == 0 && ((int64(1)<<(_la-163))&11) != 0 {
			{
				p.SetState(383)
				p.singleExpression(0)
			}

		}
		{
			p.SetState(386)
			p.Match(GXParserCloseParen)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISingleExpressionContext is an interface to support dynamic dispatch.
type ISingleExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetVariable returns the variable token.
	GetVariable() antlr.Token

	// GetCadena returns the cadena token.
	GetCadena() antlr.Token

	// GetDecimal returns the decimal token.
	GetDecimal() antlr.Token

	// GetAtributo returns the atributo token.
	GetAtributo() antlr.Token

	// SetVariable sets the variable token.
	SetVariable(antlr.Token)

	// SetCadena sets the cadena token.
	SetCadena(antlr.Token)

	// SetDecimal sets the decimal token.
	SetDecimal(antlr.Token)

	// SetAtributo sets the atributo token.
	SetAtributo(antlr.Token)

	// IsSingleExpressionContext differentiates from other interfaces.
	IsSingleExpressionContext()
}

type SingleExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	variable antlr.Token
	cadena   antlr.Token
	decimal  antlr.Token
	atributo antlr.Token
}

func NewEmptySingleExpressionContext() *SingleExpressionContext {
	var p = new(SingleExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GXParserRULE_singleExpression
	return p
}

func (*SingleExpressionContext) IsSingleExpressionContext() {}

func NewSingleExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleExpressionContext {
	var p = new(SingleExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GXParserRULE_singleExpression

	return p
}

func (s *SingleExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleExpressionContext) GetVariable() antlr.Token { return s.variable }

func (s *SingleExpressionContext) GetCadena() antlr.Token { return s.cadena }

func (s *SingleExpressionContext) GetDecimal() antlr.Token { return s.decimal }

func (s *SingleExpressionContext) GetAtributo() antlr.Token { return s.atributo }

func (s *SingleExpressionContext) SetVariable(v antlr.Token) { s.variable = v }

func (s *SingleExpressionContext) SetCadena(v antlr.Token) { s.cadena = v }

func (s *SingleExpressionContext) SetDecimal(v antlr.Token) { s.decimal = v }

func (s *SingleExpressionContext) SetAtributo(v antlr.Token) { s.atributo = v }

func (s *SingleExpressionContext) VARIABLE() antlr.TerminalNode {
	return s.GetToken(GXParserVARIABLE, 0)
}

func (s *SingleExpressionContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(GXParserStringLiteral, 0)
}

func (s *SingleExpressionContext) DecimalLiteral() antlr.TerminalNode {
	return s.GetToken(GXParserDecimalLiteral, 0)
}

func (s *SingleExpressionContext) ATRIBUTO() antlr.TerminalNode {
	return s.GetToken(GXParserATRIBUTO, 0)
}

func (s *SingleExpressionContext) Funcion() IFuncionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncionContext)
}

func (s *SingleExpressionContext) ATRIBUTOVAR() antlr.TerminalNode {
	return s.GetToken(GXParserATRIBUTOVAR, 0)
}

func (s *SingleExpressionContext) NEW_() antlr.TerminalNode {
	return s.GetToken(GXParserNEW_, 0)
}

func (s *SingleExpressionContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(GXParserOpenParen, 0)
}

func (s *SingleExpressionContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(GXParserCloseParen, 0)
}

func (s *SingleExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISingleExpressionContext); ok {
			len++
		}
	}

	tst := make([]ISingleExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISingleExpressionContext); ok {
			tst[i] = t.(ISingleExpressionContext)
			i++
		}
	}

	return tst
}

func (s *SingleExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *SingleExpressionContext) STAR() antlr.TerminalNode {
	return s.GetToken(GXParserSTAR, 0)
}

func (s *SingleExpressionContext) SLASH() antlr.TerminalNode {
	return s.GetToken(GXParserSLASH, 0)
}

func (s *SingleExpressionContext) Plus() antlr.TerminalNode {
	return s.GetToken(GXParserPlus, 0)
}

func (s *SingleExpressionContext) Minus() antlr.TerminalNode {
	return s.GetToken(GXParserMinus, 0)
}

func (s *SingleExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SingleExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GXParserListener); ok {
		listenerT.EnterSingleExpression(s)
	}
}

func (s *SingleExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GXParserListener); ok {
		listenerT.ExitSingleExpression(s)
	}
}

func (p *GXParser) SingleExpression() (localctx ISingleExpressionContext) {
	return p.singleExpression(0)
}

func (p *GXParser) singleExpression(_p int) (localctx ISingleExpressionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewSingleExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ISingleExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 22
	p.EnterRecursionRule(localctx, 22, GXParserRULE_singleExpression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(404)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 57, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(390)

			var _m = p.Match(GXParserVARIABLE)

			localctx.(*SingleExpressionContext).variable = _m
		}

	case 2:
		{
			p.SetState(391)

			var _m = p.Match(GXParserStringLiteral)

			localctx.(*SingleExpressionContext).cadena = _m
		}

	case 3:
		{
			p.SetState(392)

			var _m = p.Match(GXParserDecimalLiteral)

			localctx.(*SingleExpressionContext).decimal = _m
		}

	case 4:
		{
			p.SetState(393)

			var _m = p.Match(GXParserATRIBUTO)

			localctx.(*SingleExpressionContext).atributo = _m
		}

	case 5:
		{
			p.SetState(394)
			p.Funcion()
		}

	case 6:
		{
			p.SetState(395)
			p.Match(GXParserATRIBUTOVAR)
		}

	case 7:
		{
			p.SetState(396)
			p.Match(GXParserNEW_)
		}
		{
			p.SetState(397)
			p.Match(GXParserATRIBUTO)
		}
		{
			p.SetState(398)
			p.Match(GXParserOpenParen)
		}
		{
			p.SetState(399)
			p.Match(GXParserCloseParen)
		}

	case 8:
		{
			p.SetState(400)
			p.Match(GXParserOpenParen)
		}
		{
			p.SetState(401)
			p.singleExpression(0)
		}
		{
			p.SetState(402)
			p.Match(GXParserCloseParen)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(414)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 59, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(412)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 58, p.GetParserRuleContext()) {
			case 1:
				localctx = NewSingleExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, GXParserRULE_singleExpression)
				p.SetState(406)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(407)
					_la = p.GetTokenStream().LA(1)

					if !(_la == GXParserSTAR || _la == GXParserSLASH) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(408)
					p.singleExpression(4)
				}

			case 2:
				localctx = NewSingleExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, GXParserRULE_singleExpression)
				p.SetState(409)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(410)
					_la = p.GetTokenStream().LA(1)

					if !(_la == GXParserPlus || _la == GXParserMinus) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(411)
					p.singleExpression(3)
				}

			}

		}
		p.SetState(416)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 59, p.GetParserRuleContext())
	}

	return localctx
}

// IPrintStatementContext is an interface to support dynamic dispatch.
type IPrintStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrintStatementContext differentiates from other interfaces.
	IsPrintStatementContext()
}

type PrintStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintStatementContext() *PrintStatementContext {
	var p = new(PrintStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GXParserRULE_printStatement
	return p
}

func (*PrintStatementContext) IsPrintStatementContext() {}

func NewPrintStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintStatementContext {
	var p = new(PrintStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GXParserRULE_printStatement

	return p
}

func (s *PrintStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintStatementContext) PRINT_() antlr.TerminalNode {
	return s.GetToken(GXParserPRINT_, 0)
}

func (s *PrintStatementContext) ATRIBUTO() antlr.TerminalNode {
	return s.GetToken(GXParserATRIBUTO, 0)
}

func (s *PrintStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GXParserListener); ok {
		listenerT.EnterPrintStatement(s)
	}
}

func (s *PrintStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GXParserListener); ok {
		listenerT.ExitPrintStatement(s)
	}
}

func (p *GXParser) PrintStatement() (localctx IPrintStatementContext) {
	this := p
	_ = this

	localctx = NewPrintStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, GXParserRULE_printStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(417)
		p.Match(GXParserPRINT_)
	}
	{
		p.SetState(418)
		p.Match(GXParserATRIBUTO)
	}

	return localctx
}

// IDocLineContext is an interface to support dynamic dispatch.
type IDocLineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTag returns the tag token.
	GetTag() antlr.Token

	// GetInfo returns the info token.
	GetInfo() antlr.Token

	// GetComment returns the comment token.
	GetComment() antlr.Token

	// SetTag sets the tag token.
	SetTag(antlr.Token)

	// SetInfo sets the info token.
	SetInfo(antlr.Token)

	// SetComment sets the comment token.
	SetComment(antlr.Token)

	// IsDocLineContext differentiates from other interfaces.
	IsDocLineContext()
}

type DocLineContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	tag     antlr.Token
	info    antlr.Token
	comment antlr.Token
}

func NewEmptyDocLineContext() *DocLineContext {
	var p = new(DocLineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GXParserRULE_docLine
	return p
}

func (*DocLineContext) IsDocLineContext() {}

func NewDocLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DocLineContext {
	var p = new(DocLineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GXParserRULE_docLine

	return p
}

func (s *DocLineContext) GetParser() antlr.Parser { return s.parser }

func (s *DocLineContext) GetTag() antlr.Token { return s.tag }

func (s *DocLineContext) GetInfo() antlr.Token { return s.info }

func (s *DocLineContext) GetComment() antlr.Token { return s.comment }

func (s *DocLineContext) SetTag(v antlr.Token) { s.tag = v }

func (s *DocLineContext) SetInfo(v antlr.Token) { s.info = v }

func (s *DocLineContext) SetComment(v antlr.Token) { s.comment = v }

func (s *DocLineContext) DocLineTag() antlr.TerminalNode {
	return s.GetToken(GXParserDocLineTag, 0)
}

func (s *DocLineContext) DocLineInfo() antlr.TerminalNode {
	return s.GetToken(GXParserDocLineInfo, 0)
}

func (s *DocLineContext) COMENTARIO() antlr.TerminalNode {
	return s.GetToken(GXParserCOMENTARIO, 0)
}

func (s *DocLineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DocLineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DocLineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GXParserListener); ok {
		listenerT.EnterDocLine(s)
	}
}

func (s *DocLineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GXParserListener); ok {
		listenerT.ExitDocLine(s)
	}
}

func (p *GXParser) DocLine() (localctx IDocLineContext) {
	this := p
	_ = this

	localctx = NewDocLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, GXParserRULE_docLine)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(423)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GXParserDocLineTag:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(420)

			var _m = p.Match(GXParserDocLineTag)

			localctx.(*DocLineContext).tag = _m
		}

	case GXParserDocLineInfo:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(421)

			var _m = p.Match(GXParserDocLineInfo)

			localctx.(*DocLineContext).info = _m
		}

	case GXParserCOMENTARIO:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(422)

			var _m = p.Match(GXParserCOMENTARIO)

			localctx.(*DocLineContext).comment = _m
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *GXParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 7:
		var t *ConditionContext = nil
		if localctx != nil {
			t = localctx.(*ConditionContext)
		}
		return p.Condition_Sempred(t, predIndex)

	case 11:
		var t *SingleExpressionContext = nil
		if localctx != nil {
			t = localctx.(*SingleExpressionContext)
		}
		return p.SingleExpression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *GXParser) Condition_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *GXParser) SingleExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
