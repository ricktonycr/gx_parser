// Code generated from GXParser.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // GXParser

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 185, 428,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 6, 2, 42, 10, 2, 13, 2, 14, 2, 43, 3, 3,
	3, 3, 5, 3, 48, 10, 3, 3, 3, 7, 3, 51, 10, 3, 12, 3, 14, 3, 54, 11, 3,
	3, 3, 3, 3, 3, 4, 5, 4, 59, 10, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 65, 10,
	4, 7, 4, 67, 10, 4, 12, 4, 14, 4, 70, 11, 4, 3, 4, 5, 4, 73, 10, 4, 3,
	4, 7, 4, 76, 10, 4, 12, 4, 14, 4, 79, 11, 4, 3, 4, 3, 4, 3, 4, 5, 4, 84,
	10, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 95,
	10, 4, 12, 4, 14, 4, 98, 11, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 111, 10, 4, 12, 4, 14, 4, 114, 11, 4,
	5, 4, 116, 10, 4, 3, 4, 3, 4, 5, 4, 120, 10, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 130, 10, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 141, 10, 4, 12, 4, 14, 4, 144, 11, 4, 3,
	4, 3, 4, 5, 4, 148, 10, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 162, 10, 4, 12, 4, 14, 4, 165, 11, 4,
	3, 4, 5, 4, 168, 10, 4, 3, 5, 3, 5, 3, 5, 5, 5, 173, 10, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 5, 5, 180, 10, 5, 5, 5, 182, 10, 5, 3, 6, 3, 6, 3, 6,
	5, 6, 187, 10, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	7, 6, 198, 10, 6, 12, 6, 14, 6, 201, 11, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 213, 10, 6, 12, 6, 14, 6, 216, 11,
	6, 5, 6, 218, 10, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 7, 7, 225, 10, 7, 12,
	7, 14, 7, 228, 11, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 7, 7, 240, 10, 7, 12, 7, 14, 7, 243, 11, 7, 7, 7, 245, 10, 7,
	12, 7, 14, 7, 248, 11, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 7, 7, 260, 10, 7, 12, 7, 14, 7, 263, 11, 7, 5, 7, 265, 10,
	7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 273, 10, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 7, 8, 284, 10, 8, 12, 8, 14, 8,
	287, 11, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 5, 9, 300, 10, 9, 3, 9, 3, 9, 3, 9, 7, 9, 305, 10, 9, 12, 9, 14,
	9, 308, 11, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 7, 10, 321, 10, 10, 12, 10, 14, 10, 324, 11, 10, 3,
	10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 335,
	10, 11, 3, 11, 3, 11, 5, 11, 339, 10, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 5, 11, 347, 10, 11, 3, 11, 3, 11, 5, 11, 351, 10, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 359, 10, 11, 3, 11, 3, 11, 5,
	11, 363, 10, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 369, 10, 11, 3, 12,
	3, 12, 3, 12, 5, 12, 374, 10, 12, 3, 12, 3, 12, 7, 12, 378, 10, 12, 12,
	12, 14, 12, 381, 11, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 387, 10, 12,
	3, 12, 5, 12, 390, 10, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 407,
	10, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 7, 13, 415, 10, 13, 12,
	13, 14, 13, 418, 11, 13, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 5, 15,
	426, 10, 15, 3, 15, 2, 4, 16, 24, 16, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
	22, 24, 26, 28, 2, 9, 4, 2, 10, 10, 18, 18, 4, 2, 10, 10, 166, 166, 4,
	2, 4, 4, 8, 8, 4, 2, 167, 167, 171, 171, 3, 2, 42, 43, 3, 2, 183, 184,
	3, 2, 174, 175, 2, 567, 2, 41, 3, 2, 2, 2, 4, 45, 3, 2, 2, 2, 6, 167, 3,
	2, 2, 2, 8, 181, 3, 2, 2, 2, 10, 183, 3, 2, 2, 2, 12, 221, 3, 2, 2, 2,
	14, 268, 3, 2, 2, 2, 16, 299, 3, 2, 2, 2, 18, 309, 3, 2, 2, 2, 20, 368,
	3, 2, 2, 2, 22, 389, 3, 2, 2, 2, 24, 406, 3, 2, 2, 2, 26, 419, 3, 2, 2,
	2, 28, 425, 3, 2, 2, 2, 30, 42, 5, 4, 3, 2, 31, 42, 5, 6, 4, 2, 32, 42,
	5, 10, 6, 2, 33, 42, 5, 12, 7, 2, 34, 42, 5, 14, 8, 2, 35, 42, 7, 9, 2,
	2, 36, 42, 5, 18, 10, 2, 37, 42, 5, 20, 11, 2, 38, 42, 5, 26, 14, 2, 39,
	42, 5, 28, 15, 2, 40, 42, 7, 4, 2, 2, 41, 30, 3, 2, 2, 2, 41, 31, 3, 2,
	2, 2, 41, 32, 3, 2, 2, 2, 41, 33, 3, 2, 2, 2, 41, 34, 3, 2, 2, 2, 41, 35,
	3, 2, 2, 2, 41, 36, 3, 2, 2, 2, 41, 37, 3, 2, 2, 2, 41, 38, 3, 2, 2, 2,
	41, 39, 3, 2, 2, 2, 41, 40, 3, 2, 2, 2, 42, 43, 3, 2, 2, 2, 43, 41, 3,
	2, 2, 2, 43, 44, 3, 2, 2, 2, 44, 3, 3, 2, 2, 2, 45, 47, 7, 27, 2, 2, 46,
	48, 7, 8, 2, 2, 47, 46, 3, 2, 2, 2, 47, 48, 3, 2, 2, 2, 48, 52, 3, 2, 2,
	2, 49, 51, 5, 20, 11, 2, 50, 49, 3, 2, 2, 2, 51, 54, 3, 2, 2, 2, 52, 50,
	3, 2, 2, 2, 52, 53, 3, 2, 2, 2, 53, 55, 3, 2, 2, 2, 54, 52, 3, 2, 2, 2,
	55, 56, 7, 28, 2, 2, 56, 5, 3, 2, 2, 2, 57, 59, 7, 6, 2, 2, 58, 57, 3,
	2, 2, 2, 58, 59, 3, 2, 2, 2, 59, 60, 3, 2, 2, 2, 60, 61, 7, 19, 2, 2, 61,
	68, 7, 36, 2, 2, 62, 64, 7, 165, 2, 2, 63, 65, 7, 170, 2, 2, 64, 63, 3,
	2, 2, 2, 64, 65, 3, 2, 2, 2, 65, 67, 3, 2, 2, 2, 66, 62, 3, 2, 2, 2, 67,
	70, 3, 2, 2, 2, 68, 66, 3, 2, 2, 2, 68, 69, 3, 2, 2, 2, 69, 72, 3, 2, 2,
	2, 70, 68, 3, 2, 2, 2, 71, 73, 7, 8, 2, 2, 72, 71, 3, 2, 2, 2, 72, 73,
	3, 2, 2, 2, 73, 77, 3, 2, 2, 2, 74, 76, 5, 8, 5, 2, 75, 74, 3, 2, 2, 2,
	76, 79, 3, 2, 2, 2, 77, 75, 3, 2, 2, 2, 77, 78, 3, 2, 2, 2, 78, 83, 3,
	2, 2, 2, 79, 77, 3, 2, 2, 2, 80, 81, 7, 38, 2, 2, 81, 82, 7, 39, 2, 2,
	82, 84, 7, 165, 2, 2, 83, 80, 3, 2, 2, 2, 83, 84, 3, 2, 2, 2, 84, 96, 3,
	2, 2, 2, 85, 95, 5, 20, 11, 2, 86, 95, 5, 26, 14, 2, 87, 95, 5, 4, 3, 2,
	88, 95, 5, 6, 4, 2, 89, 95, 5, 10, 6, 2, 90, 95, 5, 12, 7, 2, 91, 95, 5,
	14, 8, 2, 92, 95, 5, 28, 15, 2, 93, 95, 7, 4, 2, 2, 94, 85, 3, 2, 2, 2,
	94, 86, 3, 2, 2, 2, 94, 87, 3, 2, 2, 2, 94, 88, 3, 2, 2, 2, 94, 89, 3,
	2, 2, 2, 94, 90, 3, 2, 2, 2, 94, 91, 3, 2, 2, 2, 94, 92, 3, 2, 2, 2, 94,
	93, 3, 2, 2, 2, 95, 98, 3, 2, 2, 2, 96, 94, 3, 2, 2, 2, 96, 97, 3, 2, 2,
	2, 97, 115, 3, 2, 2, 2, 98, 96, 3, 2, 2, 2, 99, 100, 7, 37, 2, 2, 100,
	112, 7, 22, 2, 2, 101, 111, 5, 20, 11, 2, 102, 111, 5, 26, 14, 2, 103,
	111, 5, 4, 3, 2, 104, 111, 5, 6, 4, 2, 105, 111, 5, 10, 6, 2, 106, 111,
	5, 12, 7, 2, 107, 111, 5, 14, 8, 2, 108, 111, 5, 28, 15, 2, 109, 111, 7,
	4, 2, 2, 110, 101, 3, 2, 2, 2, 110, 102, 3, 2, 2, 2, 110, 103, 3, 2, 2,
	2, 110, 104, 3, 2, 2, 2, 110, 105, 3, 2, 2, 2, 110, 106, 3, 2, 2, 2, 110,
	107, 3, 2, 2, 2, 110, 108, 3, 2, 2, 2, 110, 109, 3, 2, 2, 2, 111, 114,
	3, 2, 2, 2, 112, 110, 3, 2, 2, 2, 112, 113, 3, 2, 2, 2, 113, 116, 3, 2,
	2, 2, 114, 112, 3, 2, 2, 2, 115, 99, 3, 2, 2, 2, 115, 116, 3, 2, 2, 2,
	116, 117, 3, 2, 2, 2, 117, 168, 7, 24, 2, 2, 118, 120, 7, 6, 2, 2, 119,
	118, 3, 2, 2, 2, 119, 120, 3, 2, 2, 2, 120, 121, 3, 2, 2, 2, 121, 122,
	7, 19, 2, 2, 122, 123, 7, 10, 2, 2, 123, 124, 7, 171, 2, 2, 124, 125, 9,
	2, 2, 2, 125, 126, 7, 48, 2, 2, 126, 129, 9, 2, 2, 2, 127, 128, 7, 49,
	2, 2, 128, 130, 9, 2, 2, 2, 129, 127, 3, 2, 2, 2, 129, 130, 3, 2, 2, 2,
	130, 142, 3, 2, 2, 2, 131, 141, 5, 20, 11, 2, 132, 141, 5, 26, 14, 2, 133,
	141, 5, 4, 3, 2, 134, 141, 5, 6, 4, 2, 135, 141, 5, 10, 6, 2, 136, 141,
	5, 12, 7, 2, 137, 141, 5, 14, 8, 2, 138, 141, 5, 28, 15, 2, 139, 141, 7,
	4, 2, 2, 140, 131, 3, 2, 2, 2, 140, 132, 3, 2, 2, 2, 140, 133, 3, 2, 2,
	2, 140, 134, 3, 2, 2, 2, 140, 135, 3, 2, 2, 2, 140, 136, 3, 2, 2, 2, 140,
	137, 3, 2, 2, 2, 140, 138, 3, 2, 2, 2, 140, 139, 3, 2, 2, 2, 141, 144,
	3, 2, 2, 2, 142, 140, 3, 2, 2, 2, 142, 143, 3, 2, 2, 2, 143, 145, 3, 2,
	2, 2, 144, 142, 3, 2, 2, 2, 145, 168, 7, 24, 2, 2, 146, 148, 7, 6, 2, 2,
	147, 146, 3, 2, 2, 2, 147, 148, 3, 2, 2, 2, 148, 149, 3, 2, 2, 2, 149,
	150, 7, 19, 2, 2, 150, 151, 7, 10, 2, 2, 151, 152, 7, 50, 2, 2, 152, 163,
	9, 3, 2, 2, 153, 162, 5, 20, 11, 2, 154, 162, 5, 4, 3, 2, 155, 162, 5,
	6, 4, 2, 156, 162, 5, 10, 6, 2, 157, 162, 5, 12, 7, 2, 158, 162, 5, 14,
	8, 2, 159, 162, 5, 28, 15, 2, 160, 162, 7, 4, 2, 2, 161, 153, 3, 2, 2,
	2, 161, 154, 3, 2, 2, 2, 161, 155, 3, 2, 2, 2, 161, 156, 3, 2, 2, 2, 161,
	157, 3, 2, 2, 2, 161, 158, 3, 2, 2, 2, 161, 159, 3, 2, 2, 2, 161, 160,
	3, 2, 2, 2, 162, 165, 3, 2, 2, 2, 163, 161, 3, 2, 2, 2, 163, 164, 3, 2,
	2, 2, 164, 166, 3, 2, 2, 2, 165, 163, 3, 2, 2, 2, 166, 168, 7, 24, 2, 2,
	167, 58, 3, 2, 2, 2, 167, 119, 3, 2, 2, 2, 167, 147, 3, 2, 2, 2, 168, 7,
	3, 2, 2, 2, 169, 170, 7, 23, 2, 2, 170, 172, 5, 16, 9, 2, 171, 173, 9,
	4, 2, 2, 172, 171, 3, 2, 2, 2, 172, 173, 3, 2, 2, 2, 173, 182, 3, 2, 2,
	2, 174, 175, 7, 23, 2, 2, 175, 176, 5, 16, 9, 2, 176, 177, 7, 37, 2, 2,
	177, 179, 5, 16, 9, 2, 178, 180, 9, 4, 2, 2, 179, 178, 3, 2, 2, 2, 179,
	180, 3, 2, 2, 2, 180, 182, 3, 2, 2, 2, 181, 169, 3, 2, 2, 2, 181, 174,
	3, 2, 2, 2, 182, 9, 3, 2, 2, 2, 183, 184, 7, 31, 2, 2, 184, 186, 5, 16,
	9, 2, 185, 187, 7, 6, 2, 2, 186, 185, 3, 2, 2, 2, 186, 187, 3, 2, 2, 2,
	187, 199, 3, 2, 2, 2, 188, 198, 5, 20, 11, 2, 189, 198, 5, 26, 14, 2, 190,
	198, 5, 4, 3, 2, 191, 198, 5, 6, 4, 2, 192, 198, 5, 10, 6, 2, 193, 198,
	5, 12, 7, 2, 194, 198, 5, 14, 8, 2, 195, 198, 5, 28, 15, 2, 196, 198, 7,
	4, 2, 2, 197, 188, 3, 2, 2, 2, 197, 189, 3, 2, 2, 2, 197, 190, 3, 2, 2,
	2, 197, 191, 3, 2, 2, 2, 197, 192, 3, 2, 2, 2, 197, 193, 3, 2, 2, 2, 197,
	194, 3, 2, 2, 2, 197, 195, 3, 2, 2, 2, 197, 196, 3, 2, 2, 2, 198, 201,
	3, 2, 2, 2, 199, 197, 3, 2, 2, 2, 199, 200, 3, 2, 2, 2, 200, 217, 3, 2,
	2, 2, 201, 199, 3, 2, 2, 2, 202, 214, 7, 32, 2, 2, 203, 213, 5, 20, 11,
	2, 204, 213, 5, 26, 14, 2, 205, 213, 5, 4, 3, 2, 206, 213, 5, 6, 4, 2,
	207, 213, 5, 10, 6, 2, 208, 213, 5, 12, 7, 2, 209, 213, 5, 14, 8, 2, 210,
	213, 5, 28, 15, 2, 211, 213, 7, 4, 2, 2, 212, 203, 3, 2, 2, 2, 212, 204,
	3, 2, 2, 2, 212, 205, 3, 2, 2, 2, 212, 206, 3, 2, 2, 2, 212, 207, 3, 2,
	2, 2, 212, 208, 3, 2, 2, 2, 212, 209, 3, 2, 2, 2, 212, 210, 3, 2, 2, 2,
	212, 211, 3, 2, 2, 2, 213, 216, 3, 2, 2, 2, 214, 212, 3, 2, 2, 2, 214,
	215, 3, 2, 2, 2, 215, 218, 3, 2, 2, 2, 216, 214, 3, 2, 2, 2, 217, 202,
	3, 2, 2, 2, 217, 218, 3, 2, 2, 2, 218, 219, 3, 2, 2, 2, 219, 220, 7, 33,
	2, 2, 220, 11, 3, 2, 2, 2, 221, 222, 7, 29, 2, 2, 222, 226, 7, 25, 2, 2,
	223, 225, 7, 8, 2, 2, 224, 223, 3, 2, 2, 2, 225, 228, 3, 2, 2, 2, 226,
	224, 3, 2, 2, 2, 226, 227, 3, 2, 2, 2, 227, 246, 3, 2, 2, 2, 228, 226,
	3, 2, 2, 2, 229, 230, 7, 25, 2, 2, 230, 241, 5, 16, 9, 2, 231, 240, 5,
	20, 11, 2, 232, 240, 5, 4, 3, 2, 233, 240, 5, 6, 4, 2, 234, 240, 5, 10,
	6, 2, 235, 240, 5, 12, 7, 2, 236, 240, 5, 14, 8, 2, 237, 240, 5, 28, 15,
	2, 238, 240, 7, 4, 2, 2, 239, 231, 3, 2, 2, 2, 239, 232, 3, 2, 2, 2, 239,
	233, 3, 2, 2, 2, 239, 234, 3, 2, 2, 2, 239, 235, 3, 2, 2, 2, 239, 236,
	3, 2, 2, 2, 239, 237, 3, 2, 2, 2, 239, 238, 3, 2, 2, 2, 240, 243, 3, 2,
	2, 2, 241, 239, 3, 2, 2, 2, 241, 242, 3, 2, 2, 2, 242, 245, 3, 2, 2, 2,
	243, 241, 3, 2, 2, 2, 244, 229, 3, 2, 2, 2, 245, 248, 3, 2, 2, 2, 246,
	244, 3, 2, 2, 2, 246, 247, 3, 2, 2, 2, 247, 264, 3, 2, 2, 2, 248, 246,
	3, 2, 2, 2, 249, 261, 7, 52, 2, 2, 250, 260, 5, 20, 11, 2, 251, 260, 5,
	26, 14, 2, 252, 260, 5, 4, 3, 2, 253, 260, 5, 6, 4, 2, 254, 260, 5, 10,
	6, 2, 255, 260, 5, 12, 7, 2, 256, 260, 5, 14, 8, 2, 257, 260, 5, 28, 15,
	2, 258, 260, 7, 4, 2, 2, 259, 250, 3, 2, 2, 2, 259, 251, 3, 2, 2, 2, 259,
	252, 3, 2, 2, 2, 259, 253, 3, 2, 2, 2, 259, 254, 3, 2, 2, 2, 259, 255,
	3, 2, 2, 2, 259, 256, 3, 2, 2, 2, 259, 257, 3, 2, 2, 2, 259, 258, 3, 2,
	2, 2, 260, 263, 3, 2, 2, 2, 261, 259, 3, 2, 2, 2, 261, 262, 3, 2, 2, 2,
	262, 265, 3, 2, 2, 2, 263, 261, 3, 2, 2, 2, 264, 249, 3, 2, 2, 2, 264,
	265, 3, 2, 2, 2, 265, 266, 3, 2, 2, 2, 266, 267, 7, 51, 2, 2, 267, 13,
	3, 2, 2, 2, 268, 269, 7, 29, 2, 2, 269, 270, 7, 34, 2, 2, 270, 272, 5,
	16, 9, 2, 271, 273, 7, 6, 2, 2, 272, 271, 3, 2, 2, 2, 272, 273, 3, 2, 2,
	2, 273, 285, 3, 2, 2, 2, 274, 284, 5, 20, 11, 2, 275, 284, 5, 26, 14, 2,
	276, 284, 5, 4, 3, 2, 277, 284, 5, 6, 4, 2, 278, 284, 5, 10, 6, 2, 279,
	284, 5, 12, 7, 2, 280, 284, 5, 14, 8, 2, 281, 284, 5, 28, 15, 2, 282, 284,
	7, 4, 2, 2, 283, 274, 3, 2, 2, 2, 283, 275, 3, 2, 2, 2, 283, 276, 3, 2,
	2, 2, 283, 277, 3, 2, 2, 2, 283, 278, 3, 2, 2, 2, 283, 279, 3, 2, 2, 2,
	283, 280, 3, 2, 2, 2, 283, 281, 3, 2, 2, 2, 283, 282, 3, 2, 2, 2, 284,
	287, 3, 2, 2, 2, 285, 283, 3, 2, 2, 2, 285, 286, 3, 2, 2, 2, 286, 288,
	3, 2, 2, 2, 287, 285, 3, 2, 2, 2, 288, 289, 7, 30, 2, 2, 289, 15, 3, 2,
	2, 2, 290, 291, 8, 9, 1, 2, 291, 292, 5, 24, 13, 2, 292, 293, 9, 5, 2,
	2, 293, 294, 5, 24, 13, 2, 294, 300, 3, 2, 2, 2, 295, 296, 7, 168, 2, 2,
	296, 297, 5, 16, 9, 2, 297, 298, 7, 169, 2, 2, 298, 300, 3, 2, 2, 2, 299,
	290, 3, 2, 2, 2, 299, 295, 3, 2, 2, 2, 300, 306, 3, 2, 2, 2, 301, 302,
	12, 4, 2, 2, 302, 303, 9, 6, 2, 2, 303, 305, 5, 16, 9, 5, 304, 301, 3,
	2, 2, 2, 305, 308, 3, 2, 2, 2, 306, 304, 3, 2, 2, 2, 306, 307, 3, 2, 2,
	2, 307, 17, 3, 2, 2, 2, 308, 306, 3, 2, 2, 2, 309, 310, 7, 40, 2, 2, 310,
	322, 7, 11, 2, 2, 311, 321, 5, 20, 11, 2, 312, 321, 5, 26, 14, 2, 313,
	321, 5, 4, 3, 2, 314, 321, 5, 6, 4, 2, 315, 321, 5, 10, 6, 2, 316, 321,
	5, 12, 7, 2, 317, 321, 5, 14, 8, 2, 318, 321, 5, 28, 15, 2, 319, 321, 7,
	4, 2, 2, 320, 311, 3, 2, 2, 2, 320, 312, 3, 2, 2, 2, 320, 313, 3, 2, 2,
	2, 320, 314, 3, 2, 2, 2, 320, 315, 3, 2, 2, 2, 320, 316, 3, 2, 2, 2, 320,
	317, 3, 2, 2, 2, 320, 318, 3, 2, 2, 2, 320, 319, 3, 2, 2, 2, 321, 324,
	3, 2, 2, 2, 322, 320, 3, 2, 2, 2, 322, 323, 3, 2, 2, 2, 323, 325, 3, 2,
	2, 2, 324, 322, 3, 2, 2, 2, 325, 326, 7, 41, 2, 2, 326, 19, 3, 2, 2, 2,
	327, 369, 5, 22, 12, 2, 328, 334, 7, 10, 2, 2, 329, 335, 7, 171, 2, 2,
	330, 331, 7, 174, 2, 2, 331, 335, 7, 171, 2, 2, 332, 333, 7, 175, 2, 2,
	333, 335, 7, 171, 2, 2, 334, 329, 3, 2, 2, 2, 334, 330, 3, 2, 2, 2, 334,
	332, 3, 2, 2, 2, 335, 338, 3, 2, 2, 2, 336, 339, 5, 22, 12, 2, 337, 339,
	5, 24, 13, 2, 338, 336, 3, 2, 2, 2, 338, 337, 3, 2, 2, 2, 339, 369, 3,
	2, 2, 2, 340, 346, 7, 165, 2, 2, 341, 347, 7, 171, 2, 2, 342, 343, 7, 174,
	2, 2, 343, 347, 7, 171, 2, 2, 344, 345, 7, 175, 2, 2, 345, 347, 7, 171,
	2, 2, 346, 341, 3, 2, 2, 2, 346, 342, 3, 2, 2, 2, 346, 344, 3, 2, 2, 2,
	347, 350, 3, 2, 2, 2, 348, 351, 5, 22, 12, 2, 349, 351, 5, 24, 13, 2, 350,
	348, 3, 2, 2, 2, 350, 349, 3, 2, 2, 2, 351, 369, 3, 2, 2, 2, 352, 358,
	7, 166, 2, 2, 353, 359, 7, 171, 2, 2, 354, 355, 7, 174, 2, 2, 355, 359,
	7, 171, 2, 2, 356, 357, 7, 175, 2, 2, 357, 359, 7, 171, 2, 2, 358, 353,
	3, 2, 2, 2, 358, 354, 3, 2, 2, 2, 358, 356, 3, 2, 2, 2, 359, 362, 3, 2,
	2, 2, 360, 363, 5, 22, 12, 2, 361, 363, 5, 24, 13, 2, 362, 360, 3, 2, 2,
	2, 362, 361, 3, 2, 2, 2, 363, 369, 3, 2, 2, 2, 364, 365, 7, 29, 2, 2, 365,
	369, 7, 11, 2, 2, 366, 369, 7, 26, 2, 2, 367, 369, 7, 45, 2, 2, 368, 327,
	3, 2, 2, 2, 368, 328, 3, 2, 2, 2, 368, 340, 3, 2, 2, 2, 368, 352, 3, 2,
	2, 2, 368, 364, 3, 2, 2, 2, 368, 366, 3, 2, 2, 2, 368, 367, 3, 2, 2, 2,
	369, 21, 3, 2, 2, 2, 370, 371, 7, 53, 2, 2, 371, 373, 7, 168, 2, 2, 372,
	374, 5, 24, 13, 2, 373, 372, 3, 2, 2, 2, 373, 374, 3, 2, 2, 2, 374, 379,
	3, 2, 2, 2, 375, 376, 7, 170, 2, 2, 376, 378, 5, 24, 13, 2, 377, 375, 3,
	2, 2, 2, 378, 381, 3, 2, 2, 2, 379, 377, 3, 2, 2, 2, 379, 380, 3, 2, 2,
	2, 380, 382, 3, 2, 2, 2, 381, 379, 3, 2, 2, 2, 382, 390, 7, 169, 2, 2,
	383, 384, 7, 166, 2, 2, 384, 386, 7, 168, 2, 2, 385, 387, 5, 24, 13, 2,
	386, 385, 3, 2, 2, 2, 386, 387, 3, 2, 2, 2, 387, 388, 3, 2, 2, 2, 388,
	390, 7, 169, 2, 2, 389, 370, 3, 2, 2, 2, 389, 383, 3, 2, 2, 2, 390, 23,
	3, 2, 2, 2, 391, 392, 8, 13, 1, 2, 392, 407, 7, 10, 2, 2, 393, 407, 7,
	11, 2, 2, 394, 407, 7, 18, 2, 2, 395, 407, 7, 165, 2, 2, 396, 407, 5, 22,
	12, 2, 397, 407, 7, 166, 2, 2, 398, 399, 7, 27, 2, 2, 399, 400, 7, 165,
	2, 2, 400, 401, 7, 168, 2, 2, 401, 407, 7, 169, 2, 2, 402, 403, 7, 168,
	2, 2, 403, 404, 5, 24, 13, 2, 404, 405, 7, 169, 2, 2, 405, 407, 3, 2, 2,
	2, 406, 391, 3, 2, 2, 2, 406, 393, 3, 2, 2, 2, 406, 394, 3, 2, 2, 2, 406,
	395, 3, 2, 2, 2, 406, 396, 3, 2, 2, 2, 406, 397, 3, 2, 2, 2, 406, 398,
	3, 2, 2, 2, 406, 402, 3, 2, 2, 2, 407, 416, 3, 2, 2, 2, 408, 409, 12, 5,
	2, 2, 409, 410, 9, 7, 2, 2, 410, 415, 5, 24, 13, 6, 411, 412, 12, 4, 2,
	2, 412, 413, 9, 8, 2, 2, 413, 415, 5, 24, 13, 5, 414, 408, 3, 2, 2, 2,
	414, 411, 3, 2, 2, 2, 415, 418, 3, 2, 2, 2, 416, 414, 3, 2, 2, 2, 416,
	417, 3, 2, 2, 2, 417, 25, 3, 2, 2, 2, 418, 416, 3, 2, 2, 2, 419, 420, 7,
	35, 2, 2, 420, 421, 7, 165, 2, 2, 421, 27, 3, 2, 2, 2, 422, 426, 7, 5,
	2, 2, 423, 426, 7, 7, 2, 2, 424, 426, 7, 8, 2, 2, 425, 422, 3, 2, 2, 2,
	425, 423, 3, 2, 2, 2, 425, 424, 3, 2, 2, 2, 426, 29, 3, 2, 2, 2, 63, 41,
	43, 47, 52, 58, 64, 68, 72, 77, 83, 94, 96, 110, 112, 115, 119, 129, 140,
	142, 147, 161, 163, 167, 172, 179, 181, 186, 197, 199, 212, 214, 217, 226,
	239, 241, 246, 259, 261, 264, 272, 283, 285, 299, 306, 320, 322, 334, 338,
	346, 350, 358, 362, 368, 373, 379, 386, 389, 406, 414, 416, 425,
}
var literalNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "'('", "')'", "','", "'='", "'++'", "'--'", "'+'", "'-'",
	"'<'", "'>'", "'<='", "'>='", "'<>'", "'|'", "'.'", "'*'", "'/'", "'@'",
}
var symbolicNames = []string{
	"", "WS", "COMMENT", "DocLineTag", "DocLineInfoPre", "DocLineInfo", "COMENTARIO",
	"FIN", "VARIABLE", "StringLiteral", "StringLiteralDGU", "StringLiteralSGU",
	"StringLiteralDGP", "StringLiteralSGP", "StringLiteralDZG", "StringLiteralSZG",
	"DecimalLiteral", "FOR_", "ADD_", "CLEAR_", "NONE_", "WHERE_", "ENDFOR_",
	"CASE_", "EXIT_", "NEW_", "ENDNEW_", "DO_", "ENDDO_", "IF_", "ELSE_", "ENDIF_",
	"WHILE_", "PRINT_", "EACH_", "WHEN_", "DEFINED_", "BY_", "SUB_", "ENDSUB_",
	"AND_", "OR_", "LIKE_", "DELETE_", "EVENT_", "ENDEVENT_", "TO_", "STEP_",
	"IN_", "ENDCASE_", "OTHERWISE_", "FUNCIONES", "CALL_", "YMDHMSTOT_", "ADDMTH_",
	"ADDYR_", "AFTER_", "AGE_", "ASK_", "BROWSERID_", "BROWSERVERSION_", "CDOW_",
	"CMONTH_", "COLOR_", "COLS_", "CONCAT_", "CONFIRMED_", "CREATE_", "CTOD_",
	"CTOT_", "CURSOR_", "DFRCLOSE_", "DFRGDATE_", "DFRGNUM_", "DFRGTXT_", "DFRNEXT_",
	"DFROPEN_", "DFWCLOSE_", "DFWNEXT_", "DFWOPEN_", "DFWPDATE_", "DFWPNUM_",
	"DFWPTXT_", "DAY_", "DECRYPT64_", "DELETEFILE_", "DOW_", "DTOC_", "ENCRYPT64_",
	"EOM_", "FILEEXIST_", "FORMAT_", "GXGETMLI_", "GXMLINES_", "GETCOOKIE_",
	"GETDATASTORE_", "GETENCRYPTIONKEY_", "GETLOCATION_", "GETSOAPERR_", "GETSOAPERRMSG_",
	"HOUR_", "INT_", "ISNULL_", "LTRIM_", "LEN_", "LEVEL_", "LINK_", "LOADBITMAP_",
	"LOWER_", "MINUTE_", "MODIFIED_", "MONTH_", "NEWLINE_", "NOW_", "NULL_",
	"NULLVALUE_", "OLD_", "OPENDOCUMENT_", "PADL_", "PADR_", "PREVIOUS_", "PRINTDOCUMENT_",
	"RGB_", "RTRIM_", "RANDOM_", "READREGKEY_", "REMOTEADDR_", "ROUND_", "ROWS_",
	"RSEED_", "SECOND_", "SERVERDATE_", "SERVERNOW_", "SERVERTIME_", "SETCOOKIE_",
	"SHELL_", "SLEEP_", "SPACE_", "STR_", "STRREPLACE_", "STRSEARCH_", "STRSEARCHREV_",
	"SUBSTR_", "SYSDATE_", "SYSTIME_", "TADD_", "TDIFF_", "TIME_", "TOFORMATTEDSTRING_",
	"TODAY_", "TRIM_", "TRUNC_", "TTOC_", "UDF_", "UDP_", "UPPER_", "USERCLS_",
	"USERID_", "VAL_", "WRKST_", "WRITEREGKEY_", "YMDTOD_", "YEAR_", "ATRIBUTO",
	"ATRIBUTOVAR", "COMPARISON", "OpenParen", "CloseParen", "Comma", "Assign",
	"PlusPlus", "MinusMinus", "Plus", "Minus", "LessThan", "MoreThan", "LessThanEquals",
	"GreaterThanEquals", "NotEquals", "Bar", "POINT", "STAR", "SLASH", "AT",
}

var ruleNames = []string{
	"gxcode", "newBlock", "forBlock", "whereCondition", "ifBlock", "doCase",
	"doWhile", "condition", "subrutine", "statement", "funcion", "singleExpression",
	"printStatement", "docLine",
}

type GXParser struct {
	*antlr.BaseParser
}

// NewGXParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *GXParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewGXParser(input antlr.TokenStream) *GXParser {
	this := new(GXParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "GXParser.g4"

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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INewBlockContext)(nil)).Elem())
	var tst = make([]INewBlockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INewBlockContext)
		}
	}

	return tst
}

func (s *GxcodeContext) NewBlock(i int) INewBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INewBlockContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INewBlockContext)
}

func (s *GxcodeContext) AllForBlock() []IForBlockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IForBlockContext)(nil)).Elem())
	var tst = make([]IForBlockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IForBlockContext)
		}
	}

	return tst
}

func (s *GxcodeContext) ForBlock(i int) IForBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IForBlockContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IForBlockContext)
}

func (s *GxcodeContext) AllIfBlock() []IIfBlockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIfBlockContext)(nil)).Elem())
	var tst = make([]IIfBlockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIfBlockContext)
		}
	}

	return tst
}

func (s *GxcodeContext) IfBlock(i int) IIfBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfBlockContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIfBlockContext)
}

func (s *GxcodeContext) AllDoCase() []IDoCaseContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDoCaseContext)(nil)).Elem())
	var tst = make([]IDoCaseContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDoCaseContext)
		}
	}

	return tst
}

func (s *GxcodeContext) DoCase(i int) IDoCaseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDoCaseContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDoCaseContext)
}

func (s *GxcodeContext) AllDoWhile() []IDoWhileContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDoWhileContext)(nil)).Elem())
	var tst = make([]IDoWhileContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDoWhileContext)
		}
	}

	return tst
}

func (s *GxcodeContext) DoWhile(i int) IDoWhileContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDoWhileContext)(nil)).Elem(), i)

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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISubrutineContext)(nil)).Elem())
	var tst = make([]ISubrutineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISubrutineContext)
		}
	}

	return tst
}

func (s *GxcodeContext) Subrutine(i int) ISubrutineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubrutineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISubrutineContext)
}

func (s *GxcodeContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *GxcodeContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *GxcodeContext) AllPrintStatement() []IPrintStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPrintStatementContext)(nil)).Elem())
	var tst = make([]IPrintStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPrintStatementContext)
		}
	}

	return tst
}

func (s *GxcodeContext) PrintStatement(i int) IPrintStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrintStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPrintStatementContext)
}

func (s *GxcodeContext) AllDocLine() []IDocLineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDocLineContext)(nil)).Elem())
	var tst = make([]IDocLineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDocLineContext)
		}
	}

	return tst
}

func (s *GxcodeContext) DocLine(i int) IDocLineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDocLineContext)(nil)).Elem(), i)

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

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GXParserCOMMENT)|(1<<GXParserDocLineTag)|(1<<GXParserDocLineInfoPre)|(1<<GXParserDocLineInfo)|(1<<GXParserCOMENTARIO)|(1<<GXParserFIN)|(1<<GXParserVARIABLE)|(1<<GXParserFOR_)|(1<<GXParserEXIT_)|(1<<GXParserNEW_)|(1<<GXParserDO_)|(1<<GXParserIF_))) != 0) || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(GXParserPRINT_-33))|(1<<(GXParserSUB_-33))|(1<<(GXParserDELETE_-33))|(1<<(GXParserFUNCIONES-33)))) != 0) || _la == GXParserATRIBUTO || _la == GXParserATRIBUTOVAR {
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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *NewBlockContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

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

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GXParserVARIABLE)|(1<<GXParserEXIT_)|(1<<GXParserDO_))) != 0) || _la == GXParserDELETE_ || _la == GXParserFUNCIONES || _la == GXParserATRIBUTO || _la == GXParserATRIBUTOVAR {
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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *ForBlockContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ForBlockContext) AllPrintStatement() []IPrintStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPrintStatementContext)(nil)).Elem())
	var tst = make([]IPrintStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPrintStatementContext)
		}
	}

	return tst
}

func (s *ForBlockContext) PrintStatement(i int) IPrintStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrintStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPrintStatementContext)
}

func (s *ForBlockContext) AllNewBlock() []INewBlockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INewBlockContext)(nil)).Elem())
	var tst = make([]INewBlockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INewBlockContext)
		}
	}

	return tst
}

func (s *ForBlockContext) NewBlock(i int) INewBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INewBlockContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INewBlockContext)
}

func (s *ForBlockContext) AllForBlock() []IForBlockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IForBlockContext)(nil)).Elem())
	var tst = make([]IForBlockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IForBlockContext)
		}
	}

	return tst
}

func (s *ForBlockContext) ForBlock(i int) IForBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IForBlockContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IForBlockContext)
}

func (s *ForBlockContext) AllIfBlock() []IIfBlockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIfBlockContext)(nil)).Elem())
	var tst = make([]IIfBlockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIfBlockContext)
		}
	}

	return tst
}

func (s *ForBlockContext) IfBlock(i int) IIfBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfBlockContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIfBlockContext)
}

func (s *ForBlockContext) AllDoCase() []IDoCaseContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDoCaseContext)(nil)).Elem())
	var tst = make([]IDoCaseContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDoCaseContext)
		}
	}

	return tst
}

func (s *ForBlockContext) DoCase(i int) IDoCaseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDoCaseContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDoCaseContext)
}

func (s *ForBlockContext) AllDoWhile() []IDoWhileContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDoWhileContext)(nil)).Elem())
	var tst = make([]IDoWhileContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDoWhileContext)
		}
	}

	return tst
}

func (s *ForBlockContext) DoWhile(i int) IDoWhileContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDoWhileContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDoWhileContext)
}

func (s *ForBlockContext) AllDocLine() []IDocLineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDocLineContext)(nil)).Elem())
	var tst = make([]IDocLineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDocLineContext)
		}
	}

	return tst
}

func (s *ForBlockContext) DocLine(i int) IDocLineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDocLineContext)(nil)).Elem(), i)

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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IWhereConditionContext)(nil)).Elem())
	var tst = make([]IWhereConditionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IWhereConditionContext)
		}
	}

	return tst
}

func (s *ForBlockContext) WhereCondition(i int) IWhereConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhereConditionContext)(nil)).Elem(), i)

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

		for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GXParserCOMMENT)|(1<<GXParserDocLineTag)|(1<<GXParserDocLineInfoPre)|(1<<GXParserDocLineInfo)|(1<<GXParserCOMENTARIO)|(1<<GXParserVARIABLE)|(1<<GXParserFOR_)|(1<<GXParserEXIT_)|(1<<GXParserNEW_)|(1<<GXParserDO_)|(1<<GXParserIF_))) != 0) || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(GXParserPRINT_-33))|(1<<(GXParserDELETE_-33))|(1<<(GXParserFUNCIONES-33)))) != 0) || _la == GXParserATRIBUTO || _la == GXParserATRIBUTOVAR {
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

			for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GXParserCOMMENT)|(1<<GXParserDocLineTag)|(1<<GXParserDocLineInfoPre)|(1<<GXParserDocLineInfo)|(1<<GXParserCOMENTARIO)|(1<<GXParserVARIABLE)|(1<<GXParserFOR_)|(1<<GXParserEXIT_)|(1<<GXParserNEW_)|(1<<GXParserDO_)|(1<<GXParserIF_))) != 0) || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(GXParserPRINT_-33))|(1<<(GXParserDELETE_-33))|(1<<(GXParserFUNCIONES-33)))) != 0) || _la == GXParserATRIBUTO || _la == GXParserATRIBUTOVAR {
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

		for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GXParserCOMMENT)|(1<<GXParserDocLineTag)|(1<<GXParserDocLineInfoPre)|(1<<GXParserDocLineInfo)|(1<<GXParserCOMENTARIO)|(1<<GXParserVARIABLE)|(1<<GXParserFOR_)|(1<<GXParserEXIT_)|(1<<GXParserNEW_)|(1<<GXParserDO_)|(1<<GXParserIF_))) != 0) || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(GXParserPRINT_-33))|(1<<(GXParserDELETE_-33))|(1<<(GXParserFUNCIONES-33)))) != 0) || _la == GXParserATRIBUTO || _la == GXParserATRIBUTOVAR {
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

		for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GXParserCOMMENT)|(1<<GXParserDocLineTag)|(1<<GXParserDocLineInfoPre)|(1<<GXParserDocLineInfo)|(1<<GXParserCOMENTARIO)|(1<<GXParserVARIABLE)|(1<<GXParserFOR_)|(1<<GXParserEXIT_)|(1<<GXParserNEW_)|(1<<GXParserDO_)|(1<<GXParserIF_))) != 0) || _la == GXParserDELETE_ || _la == GXParserFUNCIONES || _la == GXParserATRIBUTO || _la == GXParserATRIBUTOVAR {
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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IConditionContext)(nil)).Elem())
	var tst = make([]IConditionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IConditionContext)
		}
	}

	return tst
}

func (s *WhereConditionContext) Condition(i int) IConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionContext)(nil)).Elem(), i)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *IfBlockContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *IfBlockContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *IfBlockContext) AllPrintStatement() []IPrintStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPrintStatementContext)(nil)).Elem())
	var tst = make([]IPrintStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPrintStatementContext)
		}
	}

	return tst
}

func (s *IfBlockContext) PrintStatement(i int) IPrintStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrintStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPrintStatementContext)
}

func (s *IfBlockContext) AllNewBlock() []INewBlockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INewBlockContext)(nil)).Elem())
	var tst = make([]INewBlockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INewBlockContext)
		}
	}

	return tst
}

func (s *IfBlockContext) NewBlock(i int) INewBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INewBlockContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INewBlockContext)
}

func (s *IfBlockContext) AllForBlock() []IForBlockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IForBlockContext)(nil)).Elem())
	var tst = make([]IForBlockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IForBlockContext)
		}
	}

	return tst
}

func (s *IfBlockContext) ForBlock(i int) IForBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IForBlockContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IForBlockContext)
}

func (s *IfBlockContext) AllIfBlock() []IIfBlockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIfBlockContext)(nil)).Elem())
	var tst = make([]IIfBlockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIfBlockContext)
		}
	}

	return tst
}

func (s *IfBlockContext) IfBlock(i int) IIfBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfBlockContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIfBlockContext)
}

func (s *IfBlockContext) AllDoCase() []IDoCaseContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDoCaseContext)(nil)).Elem())
	var tst = make([]IDoCaseContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDoCaseContext)
		}
	}

	return tst
}

func (s *IfBlockContext) DoCase(i int) IDoCaseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDoCaseContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDoCaseContext)
}

func (s *IfBlockContext) AllDoWhile() []IDoWhileContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDoWhileContext)(nil)).Elem())
	var tst = make([]IDoWhileContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDoWhileContext)
		}
	}

	return tst
}

func (s *IfBlockContext) DoWhile(i int) IDoWhileContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDoWhileContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDoWhileContext)
}

func (s *IfBlockContext) AllDocLine() []IDocLineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDocLineContext)(nil)).Elem())
	var tst = make([]IDocLineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDocLineContext)
		}
	}

	return tst
}

func (s *IfBlockContext) DocLine(i int) IDocLineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDocLineContext)(nil)).Elem(), i)

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

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GXParserCOMMENT)|(1<<GXParserDocLineTag)|(1<<GXParserDocLineInfoPre)|(1<<GXParserDocLineInfo)|(1<<GXParserCOMENTARIO)|(1<<GXParserVARIABLE)|(1<<GXParserFOR_)|(1<<GXParserEXIT_)|(1<<GXParserNEW_)|(1<<GXParserDO_)|(1<<GXParserIF_))) != 0) || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(GXParserPRINT_-33))|(1<<(GXParserDELETE_-33))|(1<<(GXParserFUNCIONES-33)))) != 0) || _la == GXParserATRIBUTO || _la == GXParserATRIBUTOVAR {
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

		for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GXParserCOMMENT)|(1<<GXParserDocLineTag)|(1<<GXParserDocLineInfoPre)|(1<<GXParserDocLineInfo)|(1<<GXParserCOMENTARIO)|(1<<GXParserVARIABLE)|(1<<GXParserFOR_)|(1<<GXParserEXIT_)|(1<<GXParserNEW_)|(1<<GXParserDO_)|(1<<GXParserIF_))) != 0) || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(GXParserPRINT_-33))|(1<<(GXParserDELETE_-33))|(1<<(GXParserFUNCIONES-33)))) != 0) || _la == GXParserATRIBUTO || _la == GXParserATRIBUTOVAR {
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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IConditionContext)(nil)).Elem())
	var tst = make([]IConditionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IConditionContext)
		}
	}

	return tst
}

func (s *DoCaseContext) Condition(i int) IConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *DoCaseContext) OTHERWISE_() antlr.TerminalNode {
	return s.GetToken(GXParserOTHERWISE_, 0)
}

func (s *DoCaseContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *DoCaseContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *DoCaseContext) AllNewBlock() []INewBlockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INewBlockContext)(nil)).Elem())
	var tst = make([]INewBlockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INewBlockContext)
		}
	}

	return tst
}

func (s *DoCaseContext) NewBlock(i int) INewBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INewBlockContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INewBlockContext)
}

func (s *DoCaseContext) AllForBlock() []IForBlockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IForBlockContext)(nil)).Elem())
	var tst = make([]IForBlockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IForBlockContext)
		}
	}

	return tst
}

func (s *DoCaseContext) ForBlock(i int) IForBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IForBlockContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IForBlockContext)
}

func (s *DoCaseContext) AllIfBlock() []IIfBlockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIfBlockContext)(nil)).Elem())
	var tst = make([]IIfBlockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIfBlockContext)
		}
	}

	return tst
}

func (s *DoCaseContext) IfBlock(i int) IIfBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfBlockContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIfBlockContext)
}

func (s *DoCaseContext) AllDoCase() []IDoCaseContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDoCaseContext)(nil)).Elem())
	var tst = make([]IDoCaseContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDoCaseContext)
		}
	}

	return tst
}

func (s *DoCaseContext) DoCase(i int) IDoCaseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDoCaseContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDoCaseContext)
}

func (s *DoCaseContext) AllDoWhile() []IDoWhileContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDoWhileContext)(nil)).Elem())
	var tst = make([]IDoWhileContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDoWhileContext)
		}
	}

	return tst
}

func (s *DoCaseContext) DoWhile(i int) IDoWhileContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDoWhileContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDoWhileContext)
}

func (s *DoCaseContext) AllDocLine() []IDocLineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDocLineContext)(nil)).Elem())
	var tst = make([]IDocLineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDocLineContext)
		}
	}

	return tst
}

func (s *DoCaseContext) DocLine(i int) IDocLineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDocLineContext)(nil)).Elem(), i)

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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPrintStatementContext)(nil)).Elem())
	var tst = make([]IPrintStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPrintStatementContext)
		}
	}

	return tst
}

func (s *DoCaseContext) PrintStatement(i int) IPrintStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrintStatementContext)(nil)).Elem(), i)

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

		for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GXParserCOMMENT)|(1<<GXParserDocLineTag)|(1<<GXParserDocLineInfoPre)|(1<<GXParserDocLineInfo)|(1<<GXParserCOMENTARIO)|(1<<GXParserVARIABLE)|(1<<GXParserFOR_)|(1<<GXParserEXIT_)|(1<<GXParserNEW_)|(1<<GXParserDO_)|(1<<GXParserIF_))) != 0) || _la == GXParserDELETE_ || _la == GXParserFUNCIONES || _la == GXParserATRIBUTO || _la == GXParserATRIBUTOVAR {
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

		for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GXParserCOMMENT)|(1<<GXParserDocLineTag)|(1<<GXParserDocLineInfoPre)|(1<<GXParserDocLineInfo)|(1<<GXParserCOMENTARIO)|(1<<GXParserVARIABLE)|(1<<GXParserFOR_)|(1<<GXParserEXIT_)|(1<<GXParserNEW_)|(1<<GXParserDO_)|(1<<GXParserIF_))) != 0) || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(GXParserPRINT_-33))|(1<<(GXParserDELETE_-33))|(1<<(GXParserFUNCIONES-33)))) != 0) || _la == GXParserATRIBUTO || _la == GXParserATRIBUTOVAR {
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *DoWhileContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *DoWhileContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *DoWhileContext) AllPrintStatement() []IPrintStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPrintStatementContext)(nil)).Elem())
	var tst = make([]IPrintStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPrintStatementContext)
		}
	}

	return tst
}

func (s *DoWhileContext) PrintStatement(i int) IPrintStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrintStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPrintStatementContext)
}

func (s *DoWhileContext) AllNewBlock() []INewBlockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INewBlockContext)(nil)).Elem())
	var tst = make([]INewBlockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INewBlockContext)
		}
	}

	return tst
}

func (s *DoWhileContext) NewBlock(i int) INewBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INewBlockContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INewBlockContext)
}

func (s *DoWhileContext) AllForBlock() []IForBlockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IForBlockContext)(nil)).Elem())
	var tst = make([]IForBlockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IForBlockContext)
		}
	}

	return tst
}

func (s *DoWhileContext) ForBlock(i int) IForBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IForBlockContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IForBlockContext)
}

func (s *DoWhileContext) AllIfBlock() []IIfBlockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIfBlockContext)(nil)).Elem())
	var tst = make([]IIfBlockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIfBlockContext)
		}
	}

	return tst
}

func (s *DoWhileContext) IfBlock(i int) IIfBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfBlockContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIfBlockContext)
}

func (s *DoWhileContext) AllDoCase() []IDoCaseContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDoCaseContext)(nil)).Elem())
	var tst = make([]IDoCaseContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDoCaseContext)
		}
	}

	return tst
}

func (s *DoWhileContext) DoCase(i int) IDoCaseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDoCaseContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDoCaseContext)
}

func (s *DoWhileContext) AllDoWhile() []IDoWhileContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDoWhileContext)(nil)).Elem())
	var tst = make([]IDoWhileContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDoWhileContext)
		}
	}

	return tst
}

func (s *DoWhileContext) DoWhile(i int) IDoWhileContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDoWhileContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDoWhileContext)
}

func (s *DoWhileContext) AllDocLine() []IDocLineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDocLineContext)(nil)).Elem())
	var tst = make([]IDocLineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDocLineContext)
		}
	}

	return tst
}

func (s *DoWhileContext) DocLine(i int) IDocLineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDocLineContext)(nil)).Elem(), i)

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

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GXParserCOMMENT)|(1<<GXParserDocLineTag)|(1<<GXParserDocLineInfoPre)|(1<<GXParserDocLineInfo)|(1<<GXParserCOMENTARIO)|(1<<GXParserVARIABLE)|(1<<GXParserFOR_)|(1<<GXParserEXIT_)|(1<<GXParserNEW_)|(1<<GXParserDO_)|(1<<GXParserIF_))) != 0) || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(GXParserPRINT_-33))|(1<<(GXParserDELETE_-33))|(1<<(GXParserFUNCIONES-33)))) != 0) || _la == GXParserATRIBUTO || _la == GXParserATRIBUTOVAR {
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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *ConditionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IConditionContext)(nil)).Elem())
	var tst = make([]IConditionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IConditionContext)
		}
	}

	return tst
}

func (s *ConditionContext) Condition(i int) IConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionContext)(nil)).Elem(), i)

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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *SubrutineContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *SubrutineContext) AllPrintStatement() []IPrintStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPrintStatementContext)(nil)).Elem())
	var tst = make([]IPrintStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPrintStatementContext)
		}
	}

	return tst
}

func (s *SubrutineContext) PrintStatement(i int) IPrintStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrintStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPrintStatementContext)
}

func (s *SubrutineContext) AllNewBlock() []INewBlockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INewBlockContext)(nil)).Elem())
	var tst = make([]INewBlockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INewBlockContext)
		}
	}

	return tst
}

func (s *SubrutineContext) NewBlock(i int) INewBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INewBlockContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INewBlockContext)
}

func (s *SubrutineContext) AllForBlock() []IForBlockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IForBlockContext)(nil)).Elem())
	var tst = make([]IForBlockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IForBlockContext)
		}
	}

	return tst
}

func (s *SubrutineContext) ForBlock(i int) IForBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IForBlockContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IForBlockContext)
}

func (s *SubrutineContext) AllIfBlock() []IIfBlockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIfBlockContext)(nil)).Elem())
	var tst = make([]IIfBlockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIfBlockContext)
		}
	}

	return tst
}

func (s *SubrutineContext) IfBlock(i int) IIfBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfBlockContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIfBlockContext)
}

func (s *SubrutineContext) AllDoCase() []IDoCaseContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDoCaseContext)(nil)).Elem())
	var tst = make([]IDoCaseContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDoCaseContext)
		}
	}

	return tst
}

func (s *SubrutineContext) DoCase(i int) IDoCaseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDoCaseContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDoCaseContext)
}

func (s *SubrutineContext) AllDoWhile() []IDoWhileContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDoWhileContext)(nil)).Elem())
	var tst = make([]IDoWhileContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDoWhileContext)
		}
	}

	return tst
}

func (s *SubrutineContext) DoWhile(i int) IDoWhileContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDoWhileContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDoWhileContext)
}

func (s *SubrutineContext) AllDocLine() []IDocLineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDocLineContext)(nil)).Elem())
	var tst = make([]IDocLineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDocLineContext)
		}
	}

	return tst
}

func (s *SubrutineContext) DocLine(i int) IDocLineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDocLineContext)(nil)).Elem(), i)

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

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GXParserCOMMENT)|(1<<GXParserDocLineTag)|(1<<GXParserDocLineInfoPre)|(1<<GXParserDocLineInfo)|(1<<GXParserCOMENTARIO)|(1<<GXParserVARIABLE)|(1<<GXParserFOR_)|(1<<GXParserEXIT_)|(1<<GXParserNEW_)|(1<<GXParserDO_)|(1<<GXParserIF_))) != 0) || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(GXParserPRINT_-33))|(1<<(GXParserDELETE_-33))|(1<<(GXParserFUNCIONES-33)))) != 0) || _la == GXParserATRIBUTO || _la == GXParserATRIBUTOVAR {
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncionContext)(nil)).Elem(), 0)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *FuncionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

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

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GXParserVARIABLE)|(1<<GXParserStringLiteral)|(1<<GXParserDecimalLiteral)|(1<<GXParserNEW_))) != 0) || _la == GXParserFUNCIONES || (((_la-163)&-(0x1f+1)) == 0 && ((1<<uint((_la-163)))&((1<<(GXParserATRIBUTO-163))|(1<<(GXParserATRIBUTOVAR-163))|(1<<(GXParserOpenParen-163)))) != 0) {
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

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GXParserVARIABLE)|(1<<GXParserStringLiteral)|(1<<GXParserDecimalLiteral)|(1<<GXParserNEW_))) != 0) || _la == GXParserFUNCIONES || (((_la-163)&-(0x1f+1)) == 0 && ((1<<uint((_la-163)))&((1<<(GXParserATRIBUTO-163))|(1<<(GXParserATRIBUTOVAR-163))|(1<<(GXParserOpenParen-163)))) != 0) {
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncionContext)(nil)).Elem(), 0)

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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *SingleExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

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
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *GXParser) SingleExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
