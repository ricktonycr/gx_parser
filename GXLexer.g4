lexer grammar GXLexer;

WS
   : [ \t\n\r]+ -> skip
   ;

COMMENT: '/*' .*? '*/';

DocLineTag: SLASH SLASH AT ATRIBUTO WS+ DoubleStringCharacter+ (Comma DoubleStringCharacter+)* Comma StringLiteral '\n';
DocLineInfoPre: SLASH SLASH SLASH ~[\r\n]* '\n';
DocLineInfo: SLASH SLASH MoreThan ~[\r\n]* '\n';
COMENTARIO: SLASH SLASH ~[\r\n]* '\n';
FIN: '//@ #region ***************** GENERATED SUBROUTINES (PUBLIC) ******************' .*? EOF ;


VARIABLE
    : '&' NameString
    | '&' NameString OpenParen (VARIABLE | DecimalLiteral)? (Comma (VARIABLE | DecimalLiteral))* CloseParen
    ;

StringLiteral
    : ('"' DoubleStringCharacter* '"'
    | '\'' SingleStringCharacter* '\'')
    ;

StringLiteralDGU
    : 'GU: ' DoubleStringCharacter+
    ;

StringLiteralSGU
    : 'GU: '  SingleStringCharacter+
    ;

StringLiteralDGP
    : 'GP: ' DoubleStringCharacter+
    ;

StringLiteralSGP
    : 'GP: '  SingleStringCharacter+
    ;

StringLiteralDZG
    : 'ZG1: ' DoubleStringCharacter+
    ;

StringLiteralSZG
    : 'ZG1: '  SingleStringCharacter+
    ;

DecimalLiteral
    : DecimalIntegerLiteral ('.' [0-9]*)?
    ;

fragment DecimalIntegerLiteral
    : '0'
    | [1-9] [0-9]*
    ;





// Palabras reservadas
FOR_       : F O R ;
ADD_       : A D D ;
CLEAR_     : C L E A R ;
NONE_      : N O N E ;
WHERE_     : W H E R E ;
ENDFOR_    : E N D F O R ;
CASE_      : C A S E ;
EXIT_      : E X I T ;
NEW_       : N E W ;
ENDNEW_    : E N D N E W ;
DO_        : D O ;
ENDDO_     : E N D D O ;
IF_        : I F ;
ELSE_      : E L S E ;
ENDIF_     : E N D I F ;
WHILE_     : W H I L E ;
PRINT_     : P R I N T ;
EACH_      : E A C H ;
WHEN_      : W H E N ;
DEFINED_   : D E F I N E D ;
BY_        : B Y ;
SUB_       : S U B ;
ENDSUB_    : E N D S U B ;
AND_       : (A N D | '.' A N D '.') ;
OR_        : (O R | '.' O R '.') ;
LIKE_      : L I K E ;
DELETE_    : D E L E T E ;
EVENT_     : E V E N T ;
ENDEVENT_  : E N D E V E N T ;
TO_        : T O ;
STEP_      : S T E P ;
IN_        : I N ;
ENDCASE_   : E N D C A S E;
OTHERWISE_ : O T H E R W I S E;

// Palabras reservadas para funciones conocidas
FUNCIONES
    : CALL_
    | YMDHMSTOT_
    | ADDMTH_
    | ADDYR_
    | AFTER_
    | AGE_
    | ASK_
    | BROWSERID_
    | BROWSERVERSION_
    | CDOW_
    | CMONTH_
    | COLOR_
    | COLS_
    | CONCAT_
    | CONFIRMED_
    | CREATE_
    | CTOD_
    | CTOT_
    | CURSOR_
    | DFRGDATE_
    | DFRGNUM_
    | DFRGTXT_
    | DFRNEXT_
    | DFROPEN_
    | DFWCLOSE_
    | DFWNEXT_
    | DFWOPEN_
    | DFWPDATE_
    | DFWPNUM_
    | DFWPTXT_
    | DAY_
    | DECRYPT64_
    | DELETEFILE_
    | DOW_
    | DTOC_
    | ENCRYPT64_
    | EOM_
    | FILEEXIST_
    | FORMAT_
    | GXGETMLI_
    | GXMLINES_
    | GETCOOKIE_
    | GETDATASTORE_
    | GETENCRYPTIONKEY_
    | GETLOCATION_
    | GETSOAPERR_
    | GETSOAPERRMSG_
    | HOUR_
    | INT_
    | ISNULL_
    | LTRIM_
    | LEN_
    | LEVEL_
    | LINK_
    | LOADBITMAP_
    | LOWER_
    | MINUTE_
    | MODIFIED_
    | MONTH_
    | NEWLINE_
    | NOW_
    | NULL_
    | NULLVALUE_
    | OLD_
    | OPENDOCUMENT_
    | PADL_
    | PADR_
    | PREVIOUS_
    | PRINTDOCUMENT_
    | RGB_
    | RTRIM_
    | RANDOM_
    | READREGKEY_
    | REMOTEADDR_
    | ROUND_
    | ROWS_
    | RSEED_
    | SECOND_
    | SERVERDATE_
    | SERVERNOW_
    | SERVERTIME_
    | SETCOOKIE_
    | SHELL_
    | SLEEP_
    | SPACE_
    | STR_
    | STRREPLACE_
    | STRSEARCH_
    | STRSEARCHREV_
    | SUBSTR_
    | SYSDATE_
    | SYSTIME_
    | TADD_
    | TDIFF_
    | TIME_
    | TOFORMATTEDSTRING_
    | TODAY_
    | TRIM_
    | TRUNC_
    | TTOC_
    | UDF_
    | UDP_
    | UPPER_
    | USERCLS_
    | USERID_
    | VAL_
    | WRKST_
    | WRITEREGKEY_
    | YMDTOD_
    | YEAR_
    ;

CALL_              : C A L L ;
YMDHMSTOT_         : Y M D H M S T O T ;
ADDMTH_            : A D D M T H ;
ADDYR_             : A D D Y R ;
AFTER_             : A F T E R ;
AGE_               : A G E ;
ASK_               : A S K ;
BROWSERID_         : B R O W S E R I D ;
BROWSERVERSION_    : B R O W S E R V E R S I O N ;
CDOW_              : C D O W ;
CMONTH_            : C M O N T H ;
COLOR_             : C O L O R ;
COLS_              : C O L S ;
CONCAT_            : C O N C A T ;
CONFIRMED_         : C O N F I R M E D ;
CREATE_            : C R E A T E ;
CTOD_              : C T O D ;
CTOT_              : C T O T ;
CURSOR_            : C U R S O R ;
DFRCLOSE_          : D F R C L O S E ;
DFRGDATE_          : D F R G D A T E ;
DFRGNUM_           : D F R G N U M ;
DFRGTXT_           : D F R G T X T ;
DFRNEXT_           : D F R N E X T ;
DFROPEN_           : D F R O P E N ;
DFWCLOSE_          : D F W C L O S E ;
DFWNEXT_           : D F W N E X T ;
DFWOPEN_           : D F W O P E N ;
DFWPDATE_          : D F W P D A T E ;
DFWPNUM_           : D F W P N U M ;
DFWPTXT_           : D F W P T X T ;
DAY_               : D A Y ;
DECRYPT64_         : D E C R Y P T '6' '4' ;
DELETEFILE_        : D E L E T E F I L E ;
DOW_               : D O W ;
DTOC_              : D T O C ;
ENCRYPT64_         : E N C R Y P T '6' '4' ;
EOM_               : E O M ;
FILEEXIST_         : F I L E E X I S T ;
FORMAT_            : F O R M A T ;
GXGETMLI_          : G X G E T M L I ;
GXMLINES_          : G X M L I N E S ;
GETCOOKIE_         : G E T C O O K I E ;
GETDATASTORE_      : G E T D A T A S T O R E ;
GETENCRYPTIONKEY_  : G E T E N C R Y P T I O N K E Y ;
GETLOCATION_       : G E T L O C A T  I O N ;
GETSOAPERR_        : G E T S O A P E R R ;
GETSOAPERRMSG_     : G E T S O A P E R R M S G ;
HOUR_              : H O U R ;
INT_               : I N T ;
ISNULL_            : I S N U L L ;
LTRIM_             : L T R I M ;
LEN_               : L E N ;
LEVEL_             : L E V E L ;
LINK_              : L I N K ;
LOADBITMAP_        : L O A D B I T M A P ;
LOWER_             : L O W E R ;
MINUTE_            : M I N U T E ;
MODIFIED_          : M O D I F I E D ;
MONTH_             : M O N T H ;
NEWLINE_           : N E W L I N E ;
NOW_               : N O W ;
NULL_              : N U L L ;
NULLVALUE_         : N U L L V A L U E ;
OLD_               : O L D ;
OPENDOCUMENT_      : O P E N D O C U M E N T ;
PADL_              : P A D L ;
PADR_              : P A D R ;
PREVIOUS_          : P R E V I O U S ;
PRINTDOCUMENT_     : P R I N T D O C U M E N T ;
RGB_               : R G B ;
RTRIM_             : R T R I M ;
RANDOM_            : R A N D O M ;
READREGKEY_        : R E A D R E G K E Y ;
REMOTEADDR_        : R E M O T E A D D R ;
ROUND_             : R O U N D ;
ROWS_              : R O W S ;
RSEED_             : R S E E D ;
SECOND_            : S E C O N D ;
SERVERDATE_        : S E R V E R D A T E ;
SERVERNOW_         : S E R V E R N O W ;
SERVERTIME_        : S E R V E R T I M E ;
SETCOOKIE_         : S E T C O O K I E ;
SHELL_             : S H E L L ;
SLEEP_             : S L E E P ;
SPACE_             : S P A C E ;
STR_               : S T R ;
STRREPLACE_        : S T R R E P L A C E ;
STRSEARCH_         : S T R S E A R C H ;
STRSEARCHREV_      : S T R S E A R C H R E V ;
SUBSTR_            : S U B S T R ;
SYSDATE_           : S Y S D A T E;
SYSTIME_           : S Y S T I M E ;
TADD_              : T A D D ;
TDIFF_             : T D I F F ;
TIME_              : T I M E ;
TOFORMATTEDSTRING_ : T O F O R M A T T E D S T R I N G ;
TODAY_             : T O D A Y ;
TRIM_              : T R I M ;
TRUNC_             : T R U N C ;
TTOC_              : T T O C ;
UDF_               : U D F ;
UDP_               : U D P ;
UPPER_             : U P P E R ;
USERCLS_           : U S E R C L S ;
USERID_            : U S E R I D ;
VAL_               : V A L ;
WRKST_             : W R K S T ;
WRITEREGKEY_       : W R I T E R E G K E Y;
YMDTOD_            : Y M D T O D ;
YEAR_              : Y E A R ;

fragment DoubleStringCharacter
    : ~["\r\n]
    ;
fragment SingleStringCharacter
    : ~['\r\n]
    ;

fragment NameString: [a-zA-Z_][a-zA-Z0-9_]*;


ATRIBUTO
    : [a-zA-Z][a-zA-Z0-9]*
    ;

ATRIBUTOVAR
    : VARIABLE '.' ATRIBUTO
    | VARIABLE '.' ATRIBUTO '.' ATRIBUTO
    | VARIABLE '.' ADD_
    | VARIABLE '.' CLEAR_
    ;


fragment A: [aA];
fragment B: [bB];
fragment C: [cC];
fragment D: [dD];
fragment E: [eE];
fragment F: [fF];
fragment G: [gG];
fragment H: [hH];
fragment I: [iI];
fragment J: [jJ];
fragment K: [kK];
fragment L: [lL];
fragment M: [mM];
fragment N: [nN];
fragment O: [oO];
fragment P: [pP];
fragment Q: [qQ];
fragment R: [rR];
fragment S: [sS];
fragment T: [tT];
fragment U: [uU];
fragment V: [vV];
fragment W: [wW];
fragment X: [xX];
fragment Y: [yY];
fragment Z: [zZ];


COMPARISON
    : LessThan
    | LessThanEquals
    | MoreThan
    | GreaterThanEquals
    | NotEquals
    ;

OpenParen         : '(';
CloseParen        : ')';
Comma             : ',';
Assign            : '=';
PlusPlus          : '++';
MinusMinus        : '--';
Plus              : '+';
Minus             : '-';
LessThan          : '<';
MoreThan          : '>';
LessThanEquals    : '<=';
GreaterThanEquals : '>=';
NotEquals         : '<>';
Bar               : '|';

POINT: '.';

STAR
	: '*'
	;

SLASH
	: '/'
	;

AT
	: '@'
	;

