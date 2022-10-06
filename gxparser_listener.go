// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // GXParser

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// GXParserListener is a complete listener for a parse tree produced by GXParser.
type GXParserListener interface {
	antlr.ParseTreeListener

	// EnterGxcode is called when entering the gxcode production.
	EnterGxcode(c *GxcodeContext)

	// EnterNewBlock is called when entering the newBlock production.
	EnterNewBlock(c *NewBlockContext)

	// EnterForBlock is called when entering the forBlock production.
	EnterForBlock(c *ForBlockContext)

	// EnterWhereCondition is called when entering the whereCondition production.
	EnterWhereCondition(c *WhereConditionContext)

	// EnterIfBlock is called when entering the ifBlock production.
	EnterIfBlock(c *IfBlockContext)

	// EnterDoCase is called when entering the doCase production.
	EnterDoCase(c *DoCaseContext)

	// EnterDoWhile is called when entering the doWhile production.
	EnterDoWhile(c *DoWhileContext)

	// EnterCondition is called when entering the condition production.
	EnterCondition(c *ConditionContext)

	// EnterSubrutine is called when entering the subrutine production.
	EnterSubrutine(c *SubrutineContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterFuncion is called when entering the funcion production.
	EnterFuncion(c *FuncionContext)

	// EnterSingleExpression is called when entering the singleExpression production.
	EnterSingleExpression(c *SingleExpressionContext)

	// EnterPrintStatement is called when entering the printStatement production.
	EnterPrintStatement(c *PrintStatementContext)

	// EnterDocLine is called when entering the docLine production.
	EnterDocLine(c *DocLineContext)

	// ExitGxcode is called when exiting the gxcode production.
	ExitGxcode(c *GxcodeContext)

	// ExitNewBlock is called when exiting the newBlock production.
	ExitNewBlock(c *NewBlockContext)

	// ExitForBlock is called when exiting the forBlock production.
	ExitForBlock(c *ForBlockContext)

	// ExitWhereCondition is called when exiting the whereCondition production.
	ExitWhereCondition(c *WhereConditionContext)

	// ExitIfBlock is called when exiting the ifBlock production.
	ExitIfBlock(c *IfBlockContext)

	// ExitDoCase is called when exiting the doCase production.
	ExitDoCase(c *DoCaseContext)

	// ExitDoWhile is called when exiting the doWhile production.
	ExitDoWhile(c *DoWhileContext)

	// ExitCondition is called when exiting the condition production.
	ExitCondition(c *ConditionContext)

	// ExitSubrutine is called when exiting the subrutine production.
	ExitSubrutine(c *SubrutineContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitFuncion is called when exiting the funcion production.
	ExitFuncion(c *FuncionContext)

	// ExitSingleExpression is called when exiting the singleExpression production.
	ExitSingleExpression(c *SingleExpressionContext)

	// ExitPrintStatement is called when exiting the printStatement production.
	ExitPrintStatement(c *PrintStatementContext)

	// ExitDocLine is called when exiting the docLine production.
	ExitDocLine(c *DocLineContext)
}
