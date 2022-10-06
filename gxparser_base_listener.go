// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // GXParser

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BaseGXParserListener is a complete listener for a parse tree produced by GXParser.
type BaseGXParserListener struct{}

var _ GXParserListener = &BaseGXParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGXParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGXParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGXParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGXParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterGxcode is called when production gxcode is entered.
func (s *BaseGXParserListener) EnterGxcode(ctx *GxcodeContext) {}

// ExitGxcode is called when production gxcode is exited.
func (s *BaseGXParserListener) ExitGxcode(ctx *GxcodeContext) {}

// EnterNewBlock is called when production newBlock is entered.
func (s *BaseGXParserListener) EnterNewBlock(ctx *NewBlockContext) {}

// ExitNewBlock is called when production newBlock is exited.
func (s *BaseGXParserListener) ExitNewBlock(ctx *NewBlockContext) {}

// EnterForBlock is called when production forBlock is entered.
func (s *BaseGXParserListener) EnterForBlock(ctx *ForBlockContext) {}

// ExitForBlock is called when production forBlock is exited.
func (s *BaseGXParserListener) ExitForBlock(ctx *ForBlockContext) {}

// EnterWhereCondition is called when production whereCondition is entered.
func (s *BaseGXParserListener) EnterWhereCondition(ctx *WhereConditionContext) {}

// ExitWhereCondition is called when production whereCondition is exited.
func (s *BaseGXParserListener) ExitWhereCondition(ctx *WhereConditionContext) {}

// EnterIfBlock is called when production ifBlock is entered.
func (s *BaseGXParserListener) EnterIfBlock(ctx *IfBlockContext) {}

// ExitIfBlock is called when production ifBlock is exited.
func (s *BaseGXParserListener) ExitIfBlock(ctx *IfBlockContext) {}

// EnterDoCase is called when production doCase is entered.
func (s *BaseGXParserListener) EnterDoCase(ctx *DoCaseContext) {}

// ExitDoCase is called when production doCase is exited.
func (s *BaseGXParserListener) ExitDoCase(ctx *DoCaseContext) {}

// EnterDoWhile is called when production doWhile is entered.
func (s *BaseGXParserListener) EnterDoWhile(ctx *DoWhileContext) {}

// ExitDoWhile is called when production doWhile is exited.
func (s *BaseGXParserListener) ExitDoWhile(ctx *DoWhileContext) {}

// EnterCondition is called when production condition is entered.
func (s *BaseGXParserListener) EnterCondition(ctx *ConditionContext) {}

// ExitCondition is called when production condition is exited.
func (s *BaseGXParserListener) ExitCondition(ctx *ConditionContext) {}

// EnterSubrutine is called when production subrutine is entered.
func (s *BaseGXParserListener) EnterSubrutine(ctx *SubrutineContext) {}

// ExitSubrutine is called when production subrutine is exited.
func (s *BaseGXParserListener) ExitSubrutine(ctx *SubrutineContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseGXParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseGXParserListener) ExitStatement(ctx *StatementContext) {}

// EnterFuncion is called when production funcion is entered.
func (s *BaseGXParserListener) EnterFuncion(ctx *FuncionContext) {}

// ExitFuncion is called when production funcion is exited.
func (s *BaseGXParserListener) ExitFuncion(ctx *FuncionContext) {}

// EnterSingleExpression is called when production singleExpression is entered.
func (s *BaseGXParserListener) EnterSingleExpression(ctx *SingleExpressionContext) {}

// ExitSingleExpression is called when production singleExpression is exited.
func (s *BaseGXParserListener) ExitSingleExpression(ctx *SingleExpressionContext) {}

// EnterPrintStatement is called when production printStatement is entered.
func (s *BaseGXParserListener) EnterPrintStatement(ctx *PrintStatementContext) {}

// ExitPrintStatement is called when production printStatement is exited.
func (s *BaseGXParserListener) ExitPrintStatement(ctx *PrintStatementContext) {}

// EnterDocLine is called when production docLine is entered.
func (s *BaseGXParserListener) EnterDocLine(ctx *DocLineContext) {}

// ExitDocLine is called when production docLine is exited.
func (s *BaseGXParserListener) ExitDocLine(ctx *DocLineContext) {}
