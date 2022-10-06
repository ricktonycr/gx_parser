/** Grammar by Genexus code */

// Grammar named as GX
//grammar GXParser;
parser grammar GXParser;

options { tokenVocab=GXLexer; }

gxcode
    : (newBlock | forBlock | ifBlock | doCase | doWhile | FIN | subrutine | statement | printStatement | docLine | COMMENT)+
    ;

newBlock
    : NEW_ comentario=COMENTARIO? statement* ENDNEW_ 
    ;

forBlock
    : doc=DocLineInfoPre? FOR_ EACH_ (indices+=ATRIBUTO ','?)* comentario=COMENTARIO? condiciones+=whereCondition* (DEFINED_ BY_ ATRIBUTO)? (statement | printStatement | newBlock | forBlock | ifBlock | doCase | doWhile | docLine | COMMENT)* (WHEN_ NONE_ (statement | printStatement | newBlock | forBlock | ifBlock | doCase | doWhile | docLine | COMMENT)*)? ENDFOR_
    | doc=DocLineInfoPre? FOR_ en=VARIABLE Assign desde=(DecimalLiteral|VARIABLE) TO_ hasta=(DecimalLiteral|VARIABLE) (STEP_ cada=(DecimalLiteral|VARIABLE))? (statement | printStatement | newBlock | forBlock | ifBlock | doCase | doWhile | docLine | COMMENT)* ENDFOR_
    | doc=DocLineInfoPre? FOR_ VARIABLE IN_ sdt=(VARIABLE | ATRIBUTOVAR) (statement | newBlock | forBlock | ifBlock | doCase | doWhile | docLine | COMMENT)* ENDFOR_
    ;

whereCondition
    : WHERE_ condicion=condition (COMENTARIO | COMMENT)?
    | WHERE_ condicion=condition WHEN_ condition (COMENTARIO | COMMENT)?
    ;

ifBlock
    : IF_ condicion=condition comentario=DocLineInfoPre? (statement | printStatement | newBlock | forBlock | ifBlock | doCase | doWhile | docLine | COMMENT)* (ELSE_ (statement | printStatement | newBlock | forBlock | ifBlock | doCase | doWhile | docLine | COMMENT)*)? ENDIF_ 
    ;

doCase
    : DO_ CASE_ COMENTARIO* (CASE_ condition (statement | newBlock | forBlock | ifBlock | doCase | doWhile | docLine | COMMENT)*)* (OTHERWISE_ (statement | printStatement | newBlock | forBlock | ifBlock | doCase | doWhile | docLine | COMMENT)*)? ENDCASE_
    ;

doWhile
    : DO_ WHILE_ condicion=condition comentario=DocLineInfoPre? (statement | printStatement | newBlock | forBlock | ifBlock | doCase | doWhile | docLine | COMMENT)* ENDDO_
    ;

condition
    : expresions+=singleExpression (COMPARISON | Assign) expresions+=singleExpression
    | condition (AND_ | OR_) condition
    | OpenParen condition CloseParen
    ;

subrutine
    : SUB_ nombre=StringLiteral (statement | printStatement | newBlock | forBlock | ifBlock | doCase | doWhile | docLine | COMMENT)* ENDSUB_
    ;

statement
    : funcion
    | variable=VARIABLE (Assign | Plus Assign | Minus Assign) (funcion | expresion=singleExpression)
    | ATRIBUTO (Assign | Plus Assign | Minus Assign) (funcion | singleExpression)
    | ATRIBUTOVAR (Assign | Plus Assign | Minus Assign) (funcion | singleExpression)
    | DO_ StringLiteral
    | EXIT_
    | DELETE_
    ;

funcion
    : FUNCIONES OpenParen (singleExpression?  (Comma singleExpression)* ) CloseParen
    | ATRIBUTOVAR OpenParen singleExpression? CloseParen
    ;

singleExpression
    : variable=VARIABLE
    | cadena=StringLiteral
    | decimal=DecimalLiteral
    | atributo=ATRIBUTO
    | funcion
    | ATRIBUTOVAR
    | NEW_ ATRIBUTO OpenParen CloseParen
    | singleExpression (STAR | SLASH) singleExpression
    | singleExpression (Plus | Minus) singleExpression
    | OpenParen singleExpression CloseParen
    ;

printStatement
    : PRINT_ ATRIBUTO
    ;

docLine
    : tag=DocLineTag
    | info=DocLineInfo
    | comment=COMENTARIO
    ;