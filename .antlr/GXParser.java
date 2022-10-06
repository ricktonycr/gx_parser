// Generated from /home/rcruz/go/src/github.com/ricktonycr/gx_parser/GXParser.g4 by ANTLR 4.8
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class GXParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.8", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		WS=1, COMMENT=2, DocLineTag=3, DocLineInfoPre=4, DocLineInfo=5, COMENTARIO=6, 
		FIN=7, VARIABLE=8, StringLiteral=9, StringLiteralDGU=10, StringLiteralSGU=11, 
		StringLiteralDGP=12, StringLiteralSGP=13, StringLiteralDZG=14, StringLiteralSZG=15, 
		DecimalLiteral=16, FOR_=17, ADD_=18, CLEAR_=19, NONE_=20, WHERE_=21, ENDFOR_=22, 
		CASE_=23, EXIT_=24, NEW_=25, ENDNEW_=26, DO_=27, ENDDO_=28, IF_=29, ELSE_=30, 
		ENDIF_=31, WHILE_=32, PRINT_=33, EACH_=34, WHEN_=35, DEFINED_=36, BY_=37, 
		SUB_=38, ENDSUB_=39, AND_=40, OR_=41, LIKE_=42, DELETE_=43, EVENT_=44, 
		ENDEVENT_=45, TO_=46, STEP_=47, IN_=48, ENDCASE_=49, OTHERWISE_=50, FUNCIONES=51, 
		CALL_=52, YMDHMSTOT_=53, ADDMTH_=54, ADDYR_=55, AFTER_=56, AGE_=57, ASK_=58, 
		BROWSERID_=59, BROWSERVERSION_=60, CDOW_=61, CMONTH_=62, COLOR_=63, COLS_=64, 
		CONCAT_=65, CONFIRMED_=66, CREATE_=67, CTOD_=68, CTOT_=69, CURSOR_=70, 
		DFRCLOSE_=71, DFRGDATE_=72, DFRGNUM_=73, DFRGTXT_=74, DFRNEXT_=75, DFROPEN_=76, 
		DFWCLOSE_=77, DFWNEXT_=78, DFWOPEN_=79, DFWPDATE_=80, DFWPNUM_=81, DFWPTXT_=82, 
		DAY_=83, DECRYPT64_=84, DELETEFILE_=85, DOW_=86, DTOC_=87, ENCRYPT64_=88, 
		EOM_=89, FILEEXIST_=90, FORMAT_=91, GXGETMLI_=92, GXMLINES_=93, GETCOOKIE_=94, 
		GETDATASTORE_=95, GETENCRYPTIONKEY_=96, GETLOCATION_=97, GETSOAPERR_=98, 
		GETSOAPERRMSG_=99, HOUR_=100, INT_=101, ISNULL_=102, LTRIM_=103, LEN_=104, 
		LEVEL_=105, LINK_=106, LOADBITMAP_=107, LOWER_=108, MINUTE_=109, MODIFIED_=110, 
		MONTH_=111, NEWLINE_=112, NOW_=113, NULL_=114, NULLVALUE_=115, OLD_=116, 
		OPENDOCUMENT_=117, PADL_=118, PADR_=119, PREVIOUS_=120, PRINTDOCUMENT_=121, 
		RGB_=122, RTRIM_=123, RANDOM_=124, READREGKEY_=125, REMOTEADDR_=126, ROUND_=127, 
		ROWS_=128, RSEED_=129, SECOND_=130, SERVERDATE_=131, SERVERNOW_=132, SERVERTIME_=133, 
		SETCOOKIE_=134, SHELL_=135, SLEEP_=136, SPACE_=137, STR_=138, STRREPLACE_=139, 
		STRSEARCH_=140, STRSEARCHREV_=141, SUBSTR_=142, SYSDATE_=143, SYSTIME_=144, 
		TADD_=145, TDIFF_=146, TIME_=147, TOFORMATTEDSTRING_=148, TODAY_=149, 
		TRIM_=150, TRUNC_=151, TTOC_=152, UDF_=153, UDP_=154, UPPER_=155, USERCLS_=156, 
		USERID_=157, VAL_=158, WRKST_=159, WRITEREGKEY_=160, YMDTOD_=161, YEAR_=162, 
		ATRIBUTO=163, ATRIBUTOVAR=164, COMPARISON=165, OpenParen=166, CloseParen=167, 
		Comma=168, Assign=169, PlusPlus=170, MinusMinus=171, Plus=172, Minus=173, 
		LessThan=174, MoreThan=175, LessThanEquals=176, GreaterThanEquals=177, 
		NotEquals=178, Bar=179, POINT=180, STAR=181, SLASH=182, AT=183;
	public static final int
		RULE_gxcode = 0, RULE_newBlock = 1, RULE_forBlock = 2, RULE_whereCondition = 3, 
		RULE_ifBlock = 4, RULE_doCase = 5, RULE_doWhile = 6, RULE_condition = 7, 
		RULE_subrutine = 8, RULE_statement = 9, RULE_funcion = 10, RULE_singleExpression = 11, 
		RULE_printStatement = 12, RULE_docLine = 13;
	private static String[] makeRuleNames() {
		return new String[] {
			"gxcode", "newBlock", "forBlock", "whereCondition", "ifBlock", "doCase", 
			"doWhile", "condition", "subrutine", "statement", "funcion", "singleExpression", 
			"printStatement", "docLine"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, "'('", "')'", 
			"','", "'='", "'++'", "'--'", "'+'", "'-'", "'<'", "'>'", "'<='", "'>='", 
			"'<>'", "'|'", "'.'", "'*'", "'/'", "'@'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "WS", "COMMENT", "DocLineTag", "DocLineInfoPre", "DocLineInfo", 
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
			"DECRYPT64_", "DELETEFILE_", "DOW_", "DTOC_", "ENCRYPT64_", "EOM_", "FILEEXIST_", 
			"FORMAT_", "GXGETMLI_", "GXMLINES_", "GETCOOKIE_", "GETDATASTORE_", "GETENCRYPTIONKEY_", 
			"GETLOCATION_", "GETSOAPERR_", "GETSOAPERRMSG_", "HOUR_", "INT_", "ISNULL_", 
			"LTRIM_", "LEN_", "LEVEL_", "LINK_", "LOADBITMAP_", "LOWER_", "MINUTE_", 
			"MODIFIED_", "MONTH_", "NEWLINE_", "NOW_", "NULL_", "NULLVALUE_", "OLD_", 
			"OPENDOCUMENT_", "PADL_", "PADR_", "PREVIOUS_", "PRINTDOCUMENT_", "RGB_", 
			"RTRIM_", "RANDOM_", "READREGKEY_", "REMOTEADDR_", "ROUND_", "ROWS_", 
			"RSEED_", "SECOND_", "SERVERDATE_", "SERVERNOW_", "SERVERTIME_", "SETCOOKIE_", 
			"SHELL_", "SLEEP_", "SPACE_", "STR_", "STRREPLACE_", "STRSEARCH_", "STRSEARCHREV_", 
			"SUBSTR_", "SYSDATE_", "SYSTIME_", "TADD_", "TDIFF_", "TIME_", "TOFORMATTEDSTRING_", 
			"TODAY_", "TRIM_", "TRUNC_", "TTOC_", "UDF_", "UDP_", "UPPER_", "USERCLS_", 
			"USERID_", "VAL_", "WRKST_", "WRITEREGKEY_", "YMDTOD_", "YEAR_", "ATRIBUTO", 
			"ATRIBUTOVAR", "COMPARISON", "OpenParen", "CloseParen", "Comma", "Assign", 
			"PlusPlus", "MinusMinus", "Plus", "Minus", "LessThan", "MoreThan", "LessThanEquals", 
			"GreaterThanEquals", "NotEquals", "Bar", "POINT", "STAR", "SLASH", "AT"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "GXParser.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public GXParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class GxcodeContext extends ParserRuleContext {
		public List<NewBlockContext> newBlock() {
			return getRuleContexts(NewBlockContext.class);
		}
		public NewBlockContext newBlock(int i) {
			return getRuleContext(NewBlockContext.class,i);
		}
		public List<ForBlockContext> forBlock() {
			return getRuleContexts(ForBlockContext.class);
		}
		public ForBlockContext forBlock(int i) {
			return getRuleContext(ForBlockContext.class,i);
		}
		public List<IfBlockContext> ifBlock() {
			return getRuleContexts(IfBlockContext.class);
		}
		public IfBlockContext ifBlock(int i) {
			return getRuleContext(IfBlockContext.class,i);
		}
		public List<DoCaseContext> doCase() {
			return getRuleContexts(DoCaseContext.class);
		}
		public DoCaseContext doCase(int i) {
			return getRuleContext(DoCaseContext.class,i);
		}
		public List<DoWhileContext> doWhile() {
			return getRuleContexts(DoWhileContext.class);
		}
		public DoWhileContext doWhile(int i) {
			return getRuleContext(DoWhileContext.class,i);
		}
		public List<TerminalNode> FIN() { return getTokens(GXParser.FIN); }
		public TerminalNode FIN(int i) {
			return getToken(GXParser.FIN, i);
		}
		public List<SubrutineContext> subrutine() {
			return getRuleContexts(SubrutineContext.class);
		}
		public SubrutineContext subrutine(int i) {
			return getRuleContext(SubrutineContext.class,i);
		}
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public List<PrintStatementContext> printStatement() {
			return getRuleContexts(PrintStatementContext.class);
		}
		public PrintStatementContext printStatement(int i) {
			return getRuleContext(PrintStatementContext.class,i);
		}
		public List<DocLineContext> docLine() {
			return getRuleContexts(DocLineContext.class);
		}
		public DocLineContext docLine(int i) {
			return getRuleContext(DocLineContext.class,i);
		}
		public List<TerminalNode> COMMENT() { return getTokens(GXParser.COMMENT); }
		public TerminalNode COMMENT(int i) {
			return getToken(GXParser.COMMENT, i);
		}
		public GxcodeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_gxcode; }
	}

	public final GxcodeContext gxcode() throws RecognitionException {
		GxcodeContext _localctx = new GxcodeContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_gxcode);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(39); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				setState(39);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,0,_ctx) ) {
				case 1:
					{
					setState(28);
					newBlock();
					}
					break;
				case 2:
					{
					setState(29);
					forBlock();
					}
					break;
				case 3:
					{
					setState(30);
					ifBlock();
					}
					break;
				case 4:
					{
					setState(31);
					doCase();
					}
					break;
				case 5:
					{
					setState(32);
					doWhile();
					}
					break;
				case 6:
					{
					setState(33);
					match(FIN);
					}
					break;
				case 7:
					{
					setState(34);
					subrutine();
					}
					break;
				case 8:
					{
					setState(35);
					statement();
					}
					break;
				case 9:
					{
					setState(36);
					printStatement();
					}
					break;
				case 10:
					{
					setState(37);
					docLine();
					}
					break;
				case 11:
					{
					setState(38);
					match(COMMENT);
					}
					break;
				}
				}
				setState(41); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << COMMENT) | (1L << DocLineTag) | (1L << DocLineInfoPre) | (1L << DocLineInfo) | (1L << COMENTARIO) | (1L << FIN) | (1L << VARIABLE) | (1L << FOR_) | (1L << EXIT_) | (1L << NEW_) | (1L << DO_) | (1L << IF_) | (1L << PRINT_) | (1L << SUB_) | (1L << DELETE_) | (1L << FUNCIONES))) != 0) || _la==ATRIBUTO || _la==ATRIBUTOVAR );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class NewBlockContext extends ParserRuleContext {
		public Token comentario;
		public TerminalNode NEW_() { return getToken(GXParser.NEW_, 0); }
		public TerminalNode ENDNEW_() { return getToken(GXParser.ENDNEW_, 0); }
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public TerminalNode COMENTARIO() { return getToken(GXParser.COMENTARIO, 0); }
		public NewBlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_newBlock; }
	}

	public final NewBlockContext newBlock() throws RecognitionException {
		NewBlockContext _localctx = new NewBlockContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_newBlock);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(43);
			match(NEW_);
			setState(45);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==COMENTARIO) {
				{
				setState(44);
				((NewBlockContext)_localctx).comentario = match(COMENTARIO);
				}
			}

			setState(50);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << VARIABLE) | (1L << EXIT_) | (1L << DO_) | (1L << DELETE_) | (1L << FUNCIONES))) != 0) || _la==ATRIBUTO || _la==ATRIBUTOVAR) {
				{
				{
				setState(47);
				statement();
				}
				}
				setState(52);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(53);
			match(ENDNEW_);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ForBlockContext extends ParserRuleContext {
		public Token doc;
		public Token ATRIBUTO;
		public List<Token> indices = new ArrayList<Token>();
		public Token comentario;
		public WhereConditionContext whereCondition;
		public List<WhereConditionContext> condiciones = new ArrayList<WhereConditionContext>();
		public Token en;
		public Token desde;
		public Token hasta;
		public Token cada;
		public Token sdt;
		public TerminalNode FOR_() { return getToken(GXParser.FOR_, 0); }
		public TerminalNode EACH_() { return getToken(GXParser.EACH_, 0); }
		public TerminalNode ENDFOR_() { return getToken(GXParser.ENDFOR_, 0); }
		public TerminalNode DEFINED_() { return getToken(GXParser.DEFINED_, 0); }
		public TerminalNode BY_() { return getToken(GXParser.BY_, 0); }
		public List<TerminalNode> ATRIBUTO() { return getTokens(GXParser.ATRIBUTO); }
		public TerminalNode ATRIBUTO(int i) {
			return getToken(GXParser.ATRIBUTO, i);
		}
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public List<PrintStatementContext> printStatement() {
			return getRuleContexts(PrintStatementContext.class);
		}
		public PrintStatementContext printStatement(int i) {
			return getRuleContext(PrintStatementContext.class,i);
		}
		public List<NewBlockContext> newBlock() {
			return getRuleContexts(NewBlockContext.class);
		}
		public NewBlockContext newBlock(int i) {
			return getRuleContext(NewBlockContext.class,i);
		}
		public List<ForBlockContext> forBlock() {
			return getRuleContexts(ForBlockContext.class);
		}
		public ForBlockContext forBlock(int i) {
			return getRuleContext(ForBlockContext.class,i);
		}
		public List<IfBlockContext> ifBlock() {
			return getRuleContexts(IfBlockContext.class);
		}
		public IfBlockContext ifBlock(int i) {
			return getRuleContext(IfBlockContext.class,i);
		}
		public List<DoCaseContext> doCase() {
			return getRuleContexts(DoCaseContext.class);
		}
		public DoCaseContext doCase(int i) {
			return getRuleContext(DoCaseContext.class,i);
		}
		public List<DoWhileContext> doWhile() {
			return getRuleContexts(DoWhileContext.class);
		}
		public DoWhileContext doWhile(int i) {
			return getRuleContext(DoWhileContext.class,i);
		}
		public List<DocLineContext> docLine() {
			return getRuleContexts(DocLineContext.class);
		}
		public DocLineContext docLine(int i) {
			return getRuleContext(DocLineContext.class,i);
		}
		public List<TerminalNode> COMMENT() { return getTokens(GXParser.COMMENT); }
		public TerminalNode COMMENT(int i) {
			return getToken(GXParser.COMMENT, i);
		}
		public TerminalNode WHEN_() { return getToken(GXParser.WHEN_, 0); }
		public TerminalNode NONE_() { return getToken(GXParser.NONE_, 0); }
		public TerminalNode DocLineInfoPre() { return getToken(GXParser.DocLineInfoPre, 0); }
		public TerminalNode COMENTARIO() { return getToken(GXParser.COMENTARIO, 0); }
		public List<WhereConditionContext> whereCondition() {
			return getRuleContexts(WhereConditionContext.class);
		}
		public WhereConditionContext whereCondition(int i) {
			return getRuleContext(WhereConditionContext.class,i);
		}
		public List<TerminalNode> Comma() { return getTokens(GXParser.Comma); }
		public TerminalNode Comma(int i) {
			return getToken(GXParser.Comma, i);
		}
		public TerminalNode Assign() { return getToken(GXParser.Assign, 0); }
		public TerminalNode TO_() { return getToken(GXParser.TO_, 0); }
		public List<TerminalNode> VARIABLE() { return getTokens(GXParser.VARIABLE); }
		public TerminalNode VARIABLE(int i) {
			return getToken(GXParser.VARIABLE, i);
		}
		public List<TerminalNode> DecimalLiteral() { return getTokens(GXParser.DecimalLiteral); }
		public TerminalNode DecimalLiteral(int i) {
			return getToken(GXParser.DecimalLiteral, i);
		}
		public TerminalNode STEP_() { return getToken(GXParser.STEP_, 0); }
		public TerminalNode IN_() { return getToken(GXParser.IN_, 0); }
		public TerminalNode ATRIBUTOVAR() { return getToken(GXParser.ATRIBUTOVAR, 0); }
		public ForBlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_forBlock; }
	}

	public final ForBlockContext forBlock() throws RecognitionException {
		ForBlockContext _localctx = new ForBlockContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_forBlock);
		int _la;
		try {
			int _alt;
			setState(165);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,22,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(56);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==DocLineInfoPre) {
					{
					setState(55);
					((ForBlockContext)_localctx).doc = match(DocLineInfoPre);
					}
				}

				setState(58);
				match(FOR_);
				setState(59);
				match(EACH_);
				setState(66);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,6,_ctx);
				while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
					if ( _alt==1 ) {
						{
						{
						setState(60);
						((ForBlockContext)_localctx).ATRIBUTO = match(ATRIBUTO);
						((ForBlockContext)_localctx).indices.add(((ForBlockContext)_localctx).ATRIBUTO);
						setState(62);
						_errHandler.sync(this);
						_la = _input.LA(1);
						if (_la==Comma) {
							{
							setState(61);
							match(Comma);
							}
						}

						}
						} 
					}
					setState(68);
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,6,_ctx);
				}
				setState(70);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
				case 1:
					{
					setState(69);
					((ForBlockContext)_localctx).comentario = match(COMENTARIO);
					}
					break;
				}
				setState(75);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==WHERE_) {
					{
					{
					setState(72);
					((ForBlockContext)_localctx).whereCondition = whereCondition();
					((ForBlockContext)_localctx).condiciones.add(((ForBlockContext)_localctx).whereCondition);
					}
					}
					setState(77);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(81);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==DEFINED_) {
					{
					setState(78);
					match(DEFINED_);
					setState(79);
					match(BY_);
					setState(80);
					match(ATRIBUTO);
					}
				}

				setState(94);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << COMMENT) | (1L << DocLineTag) | (1L << DocLineInfoPre) | (1L << DocLineInfo) | (1L << COMENTARIO) | (1L << VARIABLE) | (1L << FOR_) | (1L << EXIT_) | (1L << NEW_) | (1L << DO_) | (1L << IF_) | (1L << PRINT_) | (1L << DELETE_) | (1L << FUNCIONES))) != 0) || _la==ATRIBUTO || _la==ATRIBUTOVAR) {
					{
					setState(92);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
					case 1:
						{
						setState(83);
						statement();
						}
						break;
					case 2:
						{
						setState(84);
						printStatement();
						}
						break;
					case 3:
						{
						setState(85);
						newBlock();
						}
						break;
					case 4:
						{
						setState(86);
						forBlock();
						}
						break;
					case 5:
						{
						setState(87);
						ifBlock();
						}
						break;
					case 6:
						{
						setState(88);
						doCase();
						}
						break;
					case 7:
						{
						setState(89);
						doWhile();
						}
						break;
					case 8:
						{
						setState(90);
						docLine();
						}
						break;
					case 9:
						{
						setState(91);
						match(COMMENT);
						}
						break;
					}
					}
					setState(96);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(113);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==WHEN_) {
					{
					setState(97);
					match(WHEN_);
					setState(98);
					match(NONE_);
					setState(110);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << COMMENT) | (1L << DocLineTag) | (1L << DocLineInfoPre) | (1L << DocLineInfo) | (1L << COMENTARIO) | (1L << VARIABLE) | (1L << FOR_) | (1L << EXIT_) | (1L << NEW_) | (1L << DO_) | (1L << IF_) | (1L << PRINT_) | (1L << DELETE_) | (1L << FUNCIONES))) != 0) || _la==ATRIBUTO || _la==ATRIBUTOVAR) {
						{
						setState(108);
						_errHandler.sync(this);
						switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
						case 1:
							{
							setState(99);
							statement();
							}
							break;
						case 2:
							{
							setState(100);
							printStatement();
							}
							break;
						case 3:
							{
							setState(101);
							newBlock();
							}
							break;
						case 4:
							{
							setState(102);
							forBlock();
							}
							break;
						case 5:
							{
							setState(103);
							ifBlock();
							}
							break;
						case 6:
							{
							setState(104);
							doCase();
							}
							break;
						case 7:
							{
							setState(105);
							doWhile();
							}
							break;
						case 8:
							{
							setState(106);
							docLine();
							}
							break;
						case 9:
							{
							setState(107);
							match(COMMENT);
							}
							break;
						}
						}
						setState(112);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					}
				}

				setState(115);
				match(ENDFOR_);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(117);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==DocLineInfoPre) {
					{
					setState(116);
					((ForBlockContext)_localctx).doc = match(DocLineInfoPre);
					}
				}

				setState(119);
				match(FOR_);
				setState(120);
				((ForBlockContext)_localctx).en = match(VARIABLE);
				setState(121);
				match(Assign);
				setState(122);
				((ForBlockContext)_localctx).desde = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==VARIABLE || _la==DecimalLiteral) ) {
					((ForBlockContext)_localctx).desde = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(123);
				match(TO_);
				setState(124);
				((ForBlockContext)_localctx).hasta = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==VARIABLE || _la==DecimalLiteral) ) {
					((ForBlockContext)_localctx).hasta = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(127);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==STEP_) {
					{
					setState(125);
					match(STEP_);
					setState(126);
					((ForBlockContext)_localctx).cada = _input.LT(1);
					_la = _input.LA(1);
					if ( !(_la==VARIABLE || _la==DecimalLiteral) ) {
						((ForBlockContext)_localctx).cada = (Token)_errHandler.recoverInline(this);
					}
					else {
						if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
						_errHandler.reportMatch(this);
						consume();
					}
					}
				}

				setState(140);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << COMMENT) | (1L << DocLineTag) | (1L << DocLineInfoPre) | (1L << DocLineInfo) | (1L << COMENTARIO) | (1L << VARIABLE) | (1L << FOR_) | (1L << EXIT_) | (1L << NEW_) | (1L << DO_) | (1L << IF_) | (1L << PRINT_) | (1L << DELETE_) | (1L << FUNCIONES))) != 0) || _la==ATRIBUTO || _la==ATRIBUTOVAR) {
					{
					setState(138);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,17,_ctx) ) {
					case 1:
						{
						setState(129);
						statement();
						}
						break;
					case 2:
						{
						setState(130);
						printStatement();
						}
						break;
					case 3:
						{
						setState(131);
						newBlock();
						}
						break;
					case 4:
						{
						setState(132);
						forBlock();
						}
						break;
					case 5:
						{
						setState(133);
						ifBlock();
						}
						break;
					case 6:
						{
						setState(134);
						doCase();
						}
						break;
					case 7:
						{
						setState(135);
						doWhile();
						}
						break;
					case 8:
						{
						setState(136);
						docLine();
						}
						break;
					case 9:
						{
						setState(137);
						match(COMMENT);
						}
						break;
					}
					}
					setState(142);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(143);
				match(ENDFOR_);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(145);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==DocLineInfoPre) {
					{
					setState(144);
					((ForBlockContext)_localctx).doc = match(DocLineInfoPre);
					}
				}

				setState(147);
				match(FOR_);
				setState(148);
				match(VARIABLE);
				setState(149);
				match(IN_);
				setState(150);
				((ForBlockContext)_localctx).sdt = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==VARIABLE || _la==ATRIBUTOVAR) ) {
					((ForBlockContext)_localctx).sdt = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(161);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << COMMENT) | (1L << DocLineTag) | (1L << DocLineInfoPre) | (1L << DocLineInfo) | (1L << COMENTARIO) | (1L << VARIABLE) | (1L << FOR_) | (1L << EXIT_) | (1L << NEW_) | (1L << DO_) | (1L << IF_) | (1L << DELETE_) | (1L << FUNCIONES))) != 0) || _la==ATRIBUTO || _la==ATRIBUTOVAR) {
					{
					setState(159);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,20,_ctx) ) {
					case 1:
						{
						setState(151);
						statement();
						}
						break;
					case 2:
						{
						setState(152);
						newBlock();
						}
						break;
					case 3:
						{
						setState(153);
						forBlock();
						}
						break;
					case 4:
						{
						setState(154);
						ifBlock();
						}
						break;
					case 5:
						{
						setState(155);
						doCase();
						}
						break;
					case 6:
						{
						setState(156);
						doWhile();
						}
						break;
					case 7:
						{
						setState(157);
						docLine();
						}
						break;
					case 8:
						{
						setState(158);
						match(COMMENT);
						}
						break;
					}
					}
					setState(163);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(164);
				match(ENDFOR_);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class WhereConditionContext extends ParserRuleContext {
		public ConditionContext condicion;
		public TerminalNode WHERE_() { return getToken(GXParser.WHERE_, 0); }
		public List<ConditionContext> condition() {
			return getRuleContexts(ConditionContext.class);
		}
		public ConditionContext condition(int i) {
			return getRuleContext(ConditionContext.class,i);
		}
		public TerminalNode COMENTARIO() { return getToken(GXParser.COMENTARIO, 0); }
		public TerminalNode COMMENT() { return getToken(GXParser.COMMENT, 0); }
		public TerminalNode WHEN_() { return getToken(GXParser.WHEN_, 0); }
		public WhereConditionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_whereCondition; }
	}

	public final WhereConditionContext whereCondition() throws RecognitionException {
		WhereConditionContext _localctx = new WhereConditionContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_whereCondition);
		int _la;
		try {
			setState(179);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,25,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(167);
				match(WHERE_);
				setState(168);
				((WhereConditionContext)_localctx).condicion = condition(0);
				setState(170);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,23,_ctx) ) {
				case 1:
					{
					setState(169);
					_la = _input.LA(1);
					if ( !(_la==COMMENT || _la==COMENTARIO) ) {
					_errHandler.recoverInline(this);
					}
					else {
						if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
						_errHandler.reportMatch(this);
						consume();
					}
					}
					break;
				}
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(172);
				match(WHERE_);
				setState(173);
				((WhereConditionContext)_localctx).condicion = condition(0);
				setState(174);
				match(WHEN_);
				setState(175);
				condition(0);
				setState(177);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,24,_ctx) ) {
				case 1:
					{
					setState(176);
					_la = _input.LA(1);
					if ( !(_la==COMMENT || _la==COMENTARIO) ) {
					_errHandler.recoverInline(this);
					}
					else {
						if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
						_errHandler.reportMatch(this);
						consume();
					}
					}
					break;
				}
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class IfBlockContext extends ParserRuleContext {
		public ConditionContext condicion;
		public Token comentario;
		public TerminalNode IF_() { return getToken(GXParser.IF_, 0); }
		public TerminalNode ENDIF_() { return getToken(GXParser.ENDIF_, 0); }
		public ConditionContext condition() {
			return getRuleContext(ConditionContext.class,0);
		}
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public List<PrintStatementContext> printStatement() {
			return getRuleContexts(PrintStatementContext.class);
		}
		public PrintStatementContext printStatement(int i) {
			return getRuleContext(PrintStatementContext.class,i);
		}
		public List<NewBlockContext> newBlock() {
			return getRuleContexts(NewBlockContext.class);
		}
		public NewBlockContext newBlock(int i) {
			return getRuleContext(NewBlockContext.class,i);
		}
		public List<ForBlockContext> forBlock() {
			return getRuleContexts(ForBlockContext.class);
		}
		public ForBlockContext forBlock(int i) {
			return getRuleContext(ForBlockContext.class,i);
		}
		public List<IfBlockContext> ifBlock() {
			return getRuleContexts(IfBlockContext.class);
		}
		public IfBlockContext ifBlock(int i) {
			return getRuleContext(IfBlockContext.class,i);
		}
		public List<DoCaseContext> doCase() {
			return getRuleContexts(DoCaseContext.class);
		}
		public DoCaseContext doCase(int i) {
			return getRuleContext(DoCaseContext.class,i);
		}
		public List<DoWhileContext> doWhile() {
			return getRuleContexts(DoWhileContext.class);
		}
		public DoWhileContext doWhile(int i) {
			return getRuleContext(DoWhileContext.class,i);
		}
		public List<DocLineContext> docLine() {
			return getRuleContexts(DocLineContext.class);
		}
		public DocLineContext docLine(int i) {
			return getRuleContext(DocLineContext.class,i);
		}
		public List<TerminalNode> COMMENT() { return getTokens(GXParser.COMMENT); }
		public TerminalNode COMMENT(int i) {
			return getToken(GXParser.COMMENT, i);
		}
		public TerminalNode ELSE_() { return getToken(GXParser.ELSE_, 0); }
		public TerminalNode DocLineInfoPre() { return getToken(GXParser.DocLineInfoPre, 0); }
		public IfBlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ifBlock; }
	}

	public final IfBlockContext ifBlock() throws RecognitionException {
		IfBlockContext _localctx = new IfBlockContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_ifBlock);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(181);
			match(IF_);
			setState(182);
			((IfBlockContext)_localctx).condicion = condition(0);
			setState(184);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,26,_ctx) ) {
			case 1:
				{
				setState(183);
				((IfBlockContext)_localctx).comentario = match(DocLineInfoPre);
				}
				break;
			}
			setState(197);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << COMMENT) | (1L << DocLineTag) | (1L << DocLineInfoPre) | (1L << DocLineInfo) | (1L << COMENTARIO) | (1L << VARIABLE) | (1L << FOR_) | (1L << EXIT_) | (1L << NEW_) | (1L << DO_) | (1L << IF_) | (1L << PRINT_) | (1L << DELETE_) | (1L << FUNCIONES))) != 0) || _la==ATRIBUTO || _la==ATRIBUTOVAR) {
				{
				setState(195);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,27,_ctx) ) {
				case 1:
					{
					setState(186);
					statement();
					}
					break;
				case 2:
					{
					setState(187);
					printStatement();
					}
					break;
				case 3:
					{
					setState(188);
					newBlock();
					}
					break;
				case 4:
					{
					setState(189);
					forBlock();
					}
					break;
				case 5:
					{
					setState(190);
					ifBlock();
					}
					break;
				case 6:
					{
					setState(191);
					doCase();
					}
					break;
				case 7:
					{
					setState(192);
					doWhile();
					}
					break;
				case 8:
					{
					setState(193);
					docLine();
					}
					break;
				case 9:
					{
					setState(194);
					match(COMMENT);
					}
					break;
				}
				}
				setState(199);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(215);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ELSE_) {
				{
				setState(200);
				match(ELSE_);
				setState(212);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << COMMENT) | (1L << DocLineTag) | (1L << DocLineInfoPre) | (1L << DocLineInfo) | (1L << COMENTARIO) | (1L << VARIABLE) | (1L << FOR_) | (1L << EXIT_) | (1L << NEW_) | (1L << DO_) | (1L << IF_) | (1L << PRINT_) | (1L << DELETE_) | (1L << FUNCIONES))) != 0) || _la==ATRIBUTO || _la==ATRIBUTOVAR) {
					{
					setState(210);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,29,_ctx) ) {
					case 1:
						{
						setState(201);
						statement();
						}
						break;
					case 2:
						{
						setState(202);
						printStatement();
						}
						break;
					case 3:
						{
						setState(203);
						newBlock();
						}
						break;
					case 4:
						{
						setState(204);
						forBlock();
						}
						break;
					case 5:
						{
						setState(205);
						ifBlock();
						}
						break;
					case 6:
						{
						setState(206);
						doCase();
						}
						break;
					case 7:
						{
						setState(207);
						doWhile();
						}
						break;
					case 8:
						{
						setState(208);
						docLine();
						}
						break;
					case 9:
						{
						setState(209);
						match(COMMENT);
						}
						break;
					}
					}
					setState(214);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(217);
			match(ENDIF_);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class DoCaseContext extends ParserRuleContext {
		public TerminalNode DO_() { return getToken(GXParser.DO_, 0); }
		public List<TerminalNode> CASE_() { return getTokens(GXParser.CASE_); }
		public TerminalNode CASE_(int i) {
			return getToken(GXParser.CASE_, i);
		}
		public TerminalNode ENDCASE_() { return getToken(GXParser.ENDCASE_, 0); }
		public List<TerminalNode> COMENTARIO() { return getTokens(GXParser.COMENTARIO); }
		public TerminalNode COMENTARIO(int i) {
			return getToken(GXParser.COMENTARIO, i);
		}
		public List<ConditionContext> condition() {
			return getRuleContexts(ConditionContext.class);
		}
		public ConditionContext condition(int i) {
			return getRuleContext(ConditionContext.class,i);
		}
		public TerminalNode OTHERWISE_() { return getToken(GXParser.OTHERWISE_, 0); }
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public List<NewBlockContext> newBlock() {
			return getRuleContexts(NewBlockContext.class);
		}
		public NewBlockContext newBlock(int i) {
			return getRuleContext(NewBlockContext.class,i);
		}
		public List<ForBlockContext> forBlock() {
			return getRuleContexts(ForBlockContext.class);
		}
		public ForBlockContext forBlock(int i) {
			return getRuleContext(ForBlockContext.class,i);
		}
		public List<IfBlockContext> ifBlock() {
			return getRuleContexts(IfBlockContext.class);
		}
		public IfBlockContext ifBlock(int i) {
			return getRuleContext(IfBlockContext.class,i);
		}
		public List<DoCaseContext> doCase() {
			return getRuleContexts(DoCaseContext.class);
		}
		public DoCaseContext doCase(int i) {
			return getRuleContext(DoCaseContext.class,i);
		}
		public List<DoWhileContext> doWhile() {
			return getRuleContexts(DoWhileContext.class);
		}
		public DoWhileContext doWhile(int i) {
			return getRuleContext(DoWhileContext.class,i);
		}
		public List<DocLineContext> docLine() {
			return getRuleContexts(DocLineContext.class);
		}
		public DocLineContext docLine(int i) {
			return getRuleContext(DocLineContext.class,i);
		}
		public List<TerminalNode> COMMENT() { return getTokens(GXParser.COMMENT); }
		public TerminalNode COMMENT(int i) {
			return getToken(GXParser.COMMENT, i);
		}
		public List<PrintStatementContext> printStatement() {
			return getRuleContexts(PrintStatementContext.class);
		}
		public PrintStatementContext printStatement(int i) {
			return getRuleContext(PrintStatementContext.class,i);
		}
		public DoCaseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_doCase; }
	}

	public final DoCaseContext doCase() throws RecognitionException {
		DoCaseContext _localctx = new DoCaseContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_doCase);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(219);
			match(DO_);
			setState(220);
			match(CASE_);
			setState(224);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMENTARIO) {
				{
				{
				setState(221);
				match(COMENTARIO);
				}
				}
				setState(226);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(244);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==CASE_) {
				{
				{
				setState(227);
				match(CASE_);
				setState(228);
				condition(0);
				setState(239);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << COMMENT) | (1L << DocLineTag) | (1L << DocLineInfoPre) | (1L << DocLineInfo) | (1L << COMENTARIO) | (1L << VARIABLE) | (1L << FOR_) | (1L << EXIT_) | (1L << NEW_) | (1L << DO_) | (1L << IF_) | (1L << DELETE_) | (1L << FUNCIONES))) != 0) || _la==ATRIBUTO || _la==ATRIBUTOVAR) {
					{
					setState(237);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,33,_ctx) ) {
					case 1:
						{
						setState(229);
						statement();
						}
						break;
					case 2:
						{
						setState(230);
						newBlock();
						}
						break;
					case 3:
						{
						setState(231);
						forBlock();
						}
						break;
					case 4:
						{
						setState(232);
						ifBlock();
						}
						break;
					case 5:
						{
						setState(233);
						doCase();
						}
						break;
					case 6:
						{
						setState(234);
						doWhile();
						}
						break;
					case 7:
						{
						setState(235);
						docLine();
						}
						break;
					case 8:
						{
						setState(236);
						match(COMMENT);
						}
						break;
					}
					}
					setState(241);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
				}
				setState(246);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(262);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==OTHERWISE_) {
				{
				setState(247);
				match(OTHERWISE_);
				setState(259);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << COMMENT) | (1L << DocLineTag) | (1L << DocLineInfoPre) | (1L << DocLineInfo) | (1L << COMENTARIO) | (1L << VARIABLE) | (1L << FOR_) | (1L << EXIT_) | (1L << NEW_) | (1L << DO_) | (1L << IF_) | (1L << PRINT_) | (1L << DELETE_) | (1L << FUNCIONES))) != 0) || _la==ATRIBUTO || _la==ATRIBUTOVAR) {
					{
					setState(257);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,36,_ctx) ) {
					case 1:
						{
						setState(248);
						statement();
						}
						break;
					case 2:
						{
						setState(249);
						printStatement();
						}
						break;
					case 3:
						{
						setState(250);
						newBlock();
						}
						break;
					case 4:
						{
						setState(251);
						forBlock();
						}
						break;
					case 5:
						{
						setState(252);
						ifBlock();
						}
						break;
					case 6:
						{
						setState(253);
						doCase();
						}
						break;
					case 7:
						{
						setState(254);
						doWhile();
						}
						break;
					case 8:
						{
						setState(255);
						docLine();
						}
						break;
					case 9:
						{
						setState(256);
						match(COMMENT);
						}
						break;
					}
					}
					setState(261);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(264);
			match(ENDCASE_);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class DoWhileContext extends ParserRuleContext {
		public ConditionContext condicion;
		public Token comentario;
		public TerminalNode DO_() { return getToken(GXParser.DO_, 0); }
		public TerminalNode WHILE_() { return getToken(GXParser.WHILE_, 0); }
		public TerminalNode ENDDO_() { return getToken(GXParser.ENDDO_, 0); }
		public ConditionContext condition() {
			return getRuleContext(ConditionContext.class,0);
		}
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public List<PrintStatementContext> printStatement() {
			return getRuleContexts(PrintStatementContext.class);
		}
		public PrintStatementContext printStatement(int i) {
			return getRuleContext(PrintStatementContext.class,i);
		}
		public List<NewBlockContext> newBlock() {
			return getRuleContexts(NewBlockContext.class);
		}
		public NewBlockContext newBlock(int i) {
			return getRuleContext(NewBlockContext.class,i);
		}
		public List<ForBlockContext> forBlock() {
			return getRuleContexts(ForBlockContext.class);
		}
		public ForBlockContext forBlock(int i) {
			return getRuleContext(ForBlockContext.class,i);
		}
		public List<IfBlockContext> ifBlock() {
			return getRuleContexts(IfBlockContext.class);
		}
		public IfBlockContext ifBlock(int i) {
			return getRuleContext(IfBlockContext.class,i);
		}
		public List<DoCaseContext> doCase() {
			return getRuleContexts(DoCaseContext.class);
		}
		public DoCaseContext doCase(int i) {
			return getRuleContext(DoCaseContext.class,i);
		}
		public List<DoWhileContext> doWhile() {
			return getRuleContexts(DoWhileContext.class);
		}
		public DoWhileContext doWhile(int i) {
			return getRuleContext(DoWhileContext.class,i);
		}
		public List<DocLineContext> docLine() {
			return getRuleContexts(DocLineContext.class);
		}
		public DocLineContext docLine(int i) {
			return getRuleContext(DocLineContext.class,i);
		}
		public List<TerminalNode> COMMENT() { return getTokens(GXParser.COMMENT); }
		public TerminalNode COMMENT(int i) {
			return getToken(GXParser.COMMENT, i);
		}
		public TerminalNode DocLineInfoPre() { return getToken(GXParser.DocLineInfoPre, 0); }
		public DoWhileContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_doWhile; }
	}

	public final DoWhileContext doWhile() throws RecognitionException {
		DoWhileContext _localctx = new DoWhileContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_doWhile);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(266);
			match(DO_);
			setState(267);
			match(WHILE_);
			setState(268);
			((DoWhileContext)_localctx).condicion = condition(0);
			setState(270);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,39,_ctx) ) {
			case 1:
				{
				setState(269);
				((DoWhileContext)_localctx).comentario = match(DocLineInfoPre);
				}
				break;
			}
			setState(283);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << COMMENT) | (1L << DocLineTag) | (1L << DocLineInfoPre) | (1L << DocLineInfo) | (1L << COMENTARIO) | (1L << VARIABLE) | (1L << FOR_) | (1L << EXIT_) | (1L << NEW_) | (1L << DO_) | (1L << IF_) | (1L << PRINT_) | (1L << DELETE_) | (1L << FUNCIONES))) != 0) || _la==ATRIBUTO || _la==ATRIBUTOVAR) {
				{
				setState(281);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,40,_ctx) ) {
				case 1:
					{
					setState(272);
					statement();
					}
					break;
				case 2:
					{
					setState(273);
					printStatement();
					}
					break;
				case 3:
					{
					setState(274);
					newBlock();
					}
					break;
				case 4:
					{
					setState(275);
					forBlock();
					}
					break;
				case 5:
					{
					setState(276);
					ifBlock();
					}
					break;
				case 6:
					{
					setState(277);
					doCase();
					}
					break;
				case 7:
					{
					setState(278);
					doWhile();
					}
					break;
				case 8:
					{
					setState(279);
					docLine();
					}
					break;
				case 9:
					{
					setState(280);
					match(COMMENT);
					}
					break;
				}
				}
				setState(285);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(286);
			match(ENDDO_);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ConditionContext extends ParserRuleContext {
		public SingleExpressionContext singleExpression;
		public List<SingleExpressionContext> expresions = new ArrayList<SingleExpressionContext>();
		public List<SingleExpressionContext> singleExpression() {
			return getRuleContexts(SingleExpressionContext.class);
		}
		public SingleExpressionContext singleExpression(int i) {
			return getRuleContext(SingleExpressionContext.class,i);
		}
		public TerminalNode COMPARISON() { return getToken(GXParser.COMPARISON, 0); }
		public TerminalNode Assign() { return getToken(GXParser.Assign, 0); }
		public TerminalNode OpenParen() { return getToken(GXParser.OpenParen, 0); }
		public List<ConditionContext> condition() {
			return getRuleContexts(ConditionContext.class);
		}
		public ConditionContext condition(int i) {
			return getRuleContext(ConditionContext.class,i);
		}
		public TerminalNode CloseParen() { return getToken(GXParser.CloseParen, 0); }
		public TerminalNode AND_() { return getToken(GXParser.AND_, 0); }
		public TerminalNode OR_() { return getToken(GXParser.OR_, 0); }
		public ConditionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_condition; }
	}

	public final ConditionContext condition() throws RecognitionException {
		return condition(0);
	}

	private ConditionContext condition(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ConditionContext _localctx = new ConditionContext(_ctx, _parentState);
		ConditionContext _prevctx = _localctx;
		int _startState = 14;
		enterRecursionRule(_localctx, 14, RULE_condition, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(297);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,42,_ctx) ) {
			case 1:
				{
				setState(289);
				((ConditionContext)_localctx).singleExpression = singleExpression(0);
				((ConditionContext)_localctx).expresions.add(((ConditionContext)_localctx).singleExpression);
				setState(290);
				_la = _input.LA(1);
				if ( !(_la==COMPARISON || _la==Assign) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(291);
				((ConditionContext)_localctx).singleExpression = singleExpression(0);
				((ConditionContext)_localctx).expresions.add(((ConditionContext)_localctx).singleExpression);
				}
				break;
			case 2:
				{
				setState(293);
				match(OpenParen);
				setState(294);
				condition(0);
				setState(295);
				match(CloseParen);
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(304);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,43,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ConditionContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_condition);
					setState(299);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(300);
					_la = _input.LA(1);
					if ( !(_la==AND_ || _la==OR_) ) {
					_errHandler.recoverInline(this);
					}
					else {
						if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
						_errHandler.reportMatch(this);
						consume();
					}
					setState(301);
					condition(3);
					}
					} 
				}
				setState(306);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,43,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class SubrutineContext extends ParserRuleContext {
		public Token nombre;
		public TerminalNode SUB_() { return getToken(GXParser.SUB_, 0); }
		public TerminalNode ENDSUB_() { return getToken(GXParser.ENDSUB_, 0); }
		public TerminalNode StringLiteral() { return getToken(GXParser.StringLiteral, 0); }
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public List<PrintStatementContext> printStatement() {
			return getRuleContexts(PrintStatementContext.class);
		}
		public PrintStatementContext printStatement(int i) {
			return getRuleContext(PrintStatementContext.class,i);
		}
		public List<NewBlockContext> newBlock() {
			return getRuleContexts(NewBlockContext.class);
		}
		public NewBlockContext newBlock(int i) {
			return getRuleContext(NewBlockContext.class,i);
		}
		public List<ForBlockContext> forBlock() {
			return getRuleContexts(ForBlockContext.class);
		}
		public ForBlockContext forBlock(int i) {
			return getRuleContext(ForBlockContext.class,i);
		}
		public List<IfBlockContext> ifBlock() {
			return getRuleContexts(IfBlockContext.class);
		}
		public IfBlockContext ifBlock(int i) {
			return getRuleContext(IfBlockContext.class,i);
		}
		public List<DoCaseContext> doCase() {
			return getRuleContexts(DoCaseContext.class);
		}
		public DoCaseContext doCase(int i) {
			return getRuleContext(DoCaseContext.class,i);
		}
		public List<DoWhileContext> doWhile() {
			return getRuleContexts(DoWhileContext.class);
		}
		public DoWhileContext doWhile(int i) {
			return getRuleContext(DoWhileContext.class,i);
		}
		public List<DocLineContext> docLine() {
			return getRuleContexts(DocLineContext.class);
		}
		public DocLineContext docLine(int i) {
			return getRuleContext(DocLineContext.class,i);
		}
		public List<TerminalNode> COMMENT() { return getTokens(GXParser.COMMENT); }
		public TerminalNode COMMENT(int i) {
			return getToken(GXParser.COMMENT, i);
		}
		public SubrutineContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_subrutine; }
	}

	public final SubrutineContext subrutine() throws RecognitionException {
		SubrutineContext _localctx = new SubrutineContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_subrutine);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(307);
			match(SUB_);
			setState(308);
			((SubrutineContext)_localctx).nombre = match(StringLiteral);
			setState(320);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << COMMENT) | (1L << DocLineTag) | (1L << DocLineInfoPre) | (1L << DocLineInfo) | (1L << COMENTARIO) | (1L << VARIABLE) | (1L << FOR_) | (1L << EXIT_) | (1L << NEW_) | (1L << DO_) | (1L << IF_) | (1L << PRINT_) | (1L << DELETE_) | (1L << FUNCIONES))) != 0) || _la==ATRIBUTO || _la==ATRIBUTOVAR) {
				{
				setState(318);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,44,_ctx) ) {
				case 1:
					{
					setState(309);
					statement();
					}
					break;
				case 2:
					{
					setState(310);
					printStatement();
					}
					break;
				case 3:
					{
					setState(311);
					newBlock();
					}
					break;
				case 4:
					{
					setState(312);
					forBlock();
					}
					break;
				case 5:
					{
					setState(313);
					ifBlock();
					}
					break;
				case 6:
					{
					setState(314);
					doCase();
					}
					break;
				case 7:
					{
					setState(315);
					doWhile();
					}
					break;
				case 8:
					{
					setState(316);
					docLine();
					}
					break;
				case 9:
					{
					setState(317);
					match(COMMENT);
					}
					break;
				}
				}
				setState(322);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(323);
			match(ENDSUB_);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class StatementContext extends ParserRuleContext {
		public Token variable;
		public SingleExpressionContext expresion;
		public FuncionContext funcion() {
			return getRuleContext(FuncionContext.class,0);
		}
		public TerminalNode VARIABLE() { return getToken(GXParser.VARIABLE, 0); }
		public TerminalNode Assign() { return getToken(GXParser.Assign, 0); }
		public TerminalNode Plus() { return getToken(GXParser.Plus, 0); }
		public TerminalNode Minus() { return getToken(GXParser.Minus, 0); }
		public SingleExpressionContext singleExpression() {
			return getRuleContext(SingleExpressionContext.class,0);
		}
		public TerminalNode ATRIBUTO() { return getToken(GXParser.ATRIBUTO, 0); }
		public TerminalNode ATRIBUTOVAR() { return getToken(GXParser.ATRIBUTOVAR, 0); }
		public TerminalNode DO_() { return getToken(GXParser.DO_, 0); }
		public TerminalNode StringLiteral() { return getToken(GXParser.StringLiteral, 0); }
		public TerminalNode EXIT_() { return getToken(GXParser.EXIT_, 0); }
		public TerminalNode DELETE_() { return getToken(GXParser.DELETE_, 0); }
		public StatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statement; }
	}

	public final StatementContext statement() throws RecognitionException {
		StatementContext _localctx = new StatementContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_statement);
		try {
			setState(366);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,52,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(325);
				funcion();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(326);
				((StatementContext)_localctx).variable = match(VARIABLE);
				setState(332);
				_errHandler.sync(this);
				switch (_input.LA(1)) {
				case Assign:
					{
					setState(327);
					match(Assign);
					}
					break;
				case Plus:
					{
					setState(328);
					match(Plus);
					setState(329);
					match(Assign);
					}
					break;
				case Minus:
					{
					setState(330);
					match(Minus);
					setState(331);
					match(Assign);
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(336);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,47,_ctx) ) {
				case 1:
					{
					setState(334);
					funcion();
					}
					break;
				case 2:
					{
					setState(335);
					((StatementContext)_localctx).expresion = singleExpression(0);
					}
					break;
				}
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(338);
				match(ATRIBUTO);
				setState(344);
				_errHandler.sync(this);
				switch (_input.LA(1)) {
				case Assign:
					{
					setState(339);
					match(Assign);
					}
					break;
				case Plus:
					{
					setState(340);
					match(Plus);
					setState(341);
					match(Assign);
					}
					break;
				case Minus:
					{
					setState(342);
					match(Minus);
					setState(343);
					match(Assign);
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(348);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,49,_ctx) ) {
				case 1:
					{
					setState(346);
					funcion();
					}
					break;
				case 2:
					{
					setState(347);
					singleExpression(0);
					}
					break;
				}
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(350);
				match(ATRIBUTOVAR);
				setState(356);
				_errHandler.sync(this);
				switch (_input.LA(1)) {
				case Assign:
					{
					setState(351);
					match(Assign);
					}
					break;
				case Plus:
					{
					setState(352);
					match(Plus);
					setState(353);
					match(Assign);
					}
					break;
				case Minus:
					{
					setState(354);
					match(Minus);
					setState(355);
					match(Assign);
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(360);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,51,_ctx) ) {
				case 1:
					{
					setState(358);
					funcion();
					}
					break;
				case 2:
					{
					setState(359);
					singleExpression(0);
					}
					break;
				}
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(362);
				match(DO_);
				setState(363);
				match(StringLiteral);
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(364);
				match(EXIT_);
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(365);
				match(DELETE_);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FuncionContext extends ParserRuleContext {
		public TerminalNode FUNCIONES() { return getToken(GXParser.FUNCIONES, 0); }
		public TerminalNode OpenParen() { return getToken(GXParser.OpenParen, 0); }
		public TerminalNode CloseParen() { return getToken(GXParser.CloseParen, 0); }
		public List<SingleExpressionContext> singleExpression() {
			return getRuleContexts(SingleExpressionContext.class);
		}
		public SingleExpressionContext singleExpression(int i) {
			return getRuleContext(SingleExpressionContext.class,i);
		}
		public List<TerminalNode> Comma() { return getTokens(GXParser.Comma); }
		public TerminalNode Comma(int i) {
			return getToken(GXParser.Comma, i);
		}
		public TerminalNode ATRIBUTOVAR() { return getToken(GXParser.ATRIBUTOVAR, 0); }
		public FuncionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_funcion; }
	}

	public final FuncionContext funcion() throws RecognitionException {
		FuncionContext _localctx = new FuncionContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_funcion);
		int _la;
		try {
			setState(387);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case FUNCIONES:
				enterOuterAlt(_localctx, 1);
				{
				setState(368);
				match(FUNCIONES);
				setState(369);
				match(OpenParen);
				{
				setState(371);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << VARIABLE) | (1L << StringLiteral) | (1L << DecimalLiteral) | (1L << NEW_) | (1L << FUNCIONES))) != 0) || ((((_la - 163)) & ~0x3f) == 0 && ((1L << (_la - 163)) & ((1L << (ATRIBUTO - 163)) | (1L << (ATRIBUTOVAR - 163)) | (1L << (OpenParen - 163)))) != 0)) {
					{
					setState(370);
					singleExpression(0);
					}
				}

				setState(377);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==Comma) {
					{
					{
					setState(373);
					match(Comma);
					setState(374);
					singleExpression(0);
					}
					}
					setState(379);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
				setState(380);
				match(CloseParen);
				}
				break;
			case ATRIBUTOVAR:
				enterOuterAlt(_localctx, 2);
				{
				setState(381);
				match(ATRIBUTOVAR);
				setState(382);
				match(OpenParen);
				setState(384);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << VARIABLE) | (1L << StringLiteral) | (1L << DecimalLiteral) | (1L << NEW_) | (1L << FUNCIONES))) != 0) || ((((_la - 163)) & ~0x3f) == 0 && ((1L << (_la - 163)) & ((1L << (ATRIBUTO - 163)) | (1L << (ATRIBUTOVAR - 163)) | (1L << (OpenParen - 163)))) != 0)) {
					{
					setState(383);
					singleExpression(0);
					}
				}

				setState(386);
				match(CloseParen);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class SingleExpressionContext extends ParserRuleContext {
		public Token variable;
		public Token cadena;
		public Token decimal;
		public Token atributo;
		public TerminalNode VARIABLE() { return getToken(GXParser.VARIABLE, 0); }
		public TerminalNode StringLiteral() { return getToken(GXParser.StringLiteral, 0); }
		public TerminalNode DecimalLiteral() { return getToken(GXParser.DecimalLiteral, 0); }
		public TerminalNode ATRIBUTO() { return getToken(GXParser.ATRIBUTO, 0); }
		public FuncionContext funcion() {
			return getRuleContext(FuncionContext.class,0);
		}
		public TerminalNode ATRIBUTOVAR() { return getToken(GXParser.ATRIBUTOVAR, 0); }
		public TerminalNode NEW_() { return getToken(GXParser.NEW_, 0); }
		public TerminalNode OpenParen() { return getToken(GXParser.OpenParen, 0); }
		public TerminalNode CloseParen() { return getToken(GXParser.CloseParen, 0); }
		public List<SingleExpressionContext> singleExpression() {
			return getRuleContexts(SingleExpressionContext.class);
		}
		public SingleExpressionContext singleExpression(int i) {
			return getRuleContext(SingleExpressionContext.class,i);
		}
		public TerminalNode STAR() { return getToken(GXParser.STAR, 0); }
		public TerminalNode SLASH() { return getToken(GXParser.SLASH, 0); }
		public TerminalNode Plus() { return getToken(GXParser.Plus, 0); }
		public TerminalNode Minus() { return getToken(GXParser.Minus, 0); }
		public SingleExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_singleExpression; }
	}

	public final SingleExpressionContext singleExpression() throws RecognitionException {
		return singleExpression(0);
	}

	private SingleExpressionContext singleExpression(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		SingleExpressionContext _localctx = new SingleExpressionContext(_ctx, _parentState);
		SingleExpressionContext _prevctx = _localctx;
		int _startState = 22;
		enterRecursionRule(_localctx, 22, RULE_singleExpression, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(404);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,57,_ctx) ) {
			case 1:
				{
				setState(390);
				((SingleExpressionContext)_localctx).variable = match(VARIABLE);
				}
				break;
			case 2:
				{
				setState(391);
				((SingleExpressionContext)_localctx).cadena = match(StringLiteral);
				}
				break;
			case 3:
				{
				setState(392);
				((SingleExpressionContext)_localctx).decimal = match(DecimalLiteral);
				}
				break;
			case 4:
				{
				setState(393);
				((SingleExpressionContext)_localctx).atributo = match(ATRIBUTO);
				}
				break;
			case 5:
				{
				setState(394);
				funcion();
				}
				break;
			case 6:
				{
				setState(395);
				match(ATRIBUTOVAR);
				}
				break;
			case 7:
				{
				setState(396);
				match(NEW_);
				setState(397);
				match(ATRIBUTO);
				setState(398);
				match(OpenParen);
				setState(399);
				match(CloseParen);
				}
				break;
			case 8:
				{
				setState(400);
				match(OpenParen);
				setState(401);
				singleExpression(0);
				setState(402);
				match(CloseParen);
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(414);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,59,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(412);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,58,_ctx) ) {
					case 1:
						{
						_localctx = new SingleExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_singleExpression);
						setState(406);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(407);
						_la = _input.LA(1);
						if ( !(_la==STAR || _la==SLASH) ) {
						_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(408);
						singleExpression(4);
						}
						break;
					case 2:
						{
						_localctx = new SingleExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_singleExpression);
						setState(409);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(410);
						_la = _input.LA(1);
						if ( !(_la==Plus || _la==Minus) ) {
						_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(411);
						singleExpression(3);
						}
						break;
					}
					} 
				}
				setState(416);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,59,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class PrintStatementContext extends ParserRuleContext {
		public TerminalNode PRINT_() { return getToken(GXParser.PRINT_, 0); }
		public TerminalNode ATRIBUTO() { return getToken(GXParser.ATRIBUTO, 0); }
		public PrintStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_printStatement; }
	}

	public final PrintStatementContext printStatement() throws RecognitionException {
		PrintStatementContext _localctx = new PrintStatementContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_printStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(417);
			match(PRINT_);
			setState(418);
			match(ATRIBUTO);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class DocLineContext extends ParserRuleContext {
		public Token tag;
		public Token info;
		public Token comment;
		public TerminalNode DocLineTag() { return getToken(GXParser.DocLineTag, 0); }
		public TerminalNode DocLineInfo() { return getToken(GXParser.DocLineInfo, 0); }
		public TerminalNode COMENTARIO() { return getToken(GXParser.COMENTARIO, 0); }
		public DocLineContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_docLine; }
	}

	public final DocLineContext docLine() throws RecognitionException {
		DocLineContext _localctx = new DocLineContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_docLine);
		try {
			setState(423);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case DocLineTag:
				enterOuterAlt(_localctx, 1);
				{
				setState(420);
				((DocLineContext)_localctx).tag = match(DocLineTag);
				}
				break;
			case DocLineInfo:
				enterOuterAlt(_localctx, 2);
				{
				setState(421);
				((DocLineContext)_localctx).info = match(DocLineInfo);
				}
				break;
			case COMENTARIO:
				enterOuterAlt(_localctx, 3);
				{
				setState(422);
				((DocLineContext)_localctx).comment = match(COMENTARIO);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 7:
			return condition_sempred((ConditionContext)_localctx, predIndex);
		case 11:
			return singleExpression_sempred((SingleExpressionContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean condition_sempred(ConditionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean singleExpression_sempred(SingleExpressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 1:
			return precpred(_ctx, 3);
		case 2:
			return precpred(_ctx, 2);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\u00b9\u01ac\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\3\2\3\2\3\2\3\2\3\2\3\2\3\2"+
		"\3\2\3\2\3\2\3\2\6\2*\n\2\r\2\16\2+\3\3\3\3\5\3\60\n\3\3\3\7\3\63\n\3"+
		"\f\3\16\3\66\13\3\3\3\3\3\3\4\5\4;\n\4\3\4\3\4\3\4\3\4\5\4A\n\4\7\4C\n"+
		"\4\f\4\16\4F\13\4\3\4\5\4I\n\4\3\4\7\4L\n\4\f\4\16\4O\13\4\3\4\3\4\3\4"+
		"\5\4T\n\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\7\4_\n\4\f\4\16\4b\13\4"+
		"\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\7\4o\n\4\f\4\16\4r\13\4\5"+
		"\4t\n\4\3\4\3\4\5\4x\n\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\5\4\u0082\n\4"+
		"\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\7\4\u008d\n\4\f\4\16\4\u0090\13\4"+
		"\3\4\3\4\5\4\u0094\n\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4"+
		"\7\4\u00a2\n\4\f\4\16\4\u00a5\13\4\3\4\5\4\u00a8\n\4\3\5\3\5\3\5\5\5\u00ad"+
		"\n\5\3\5\3\5\3\5\3\5\3\5\5\5\u00b4\n\5\5\5\u00b6\n\5\3\6\3\6\3\6\5\6\u00bb"+
		"\n\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\7\6\u00c6\n\6\f\6\16\6\u00c9"+
		"\13\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\7\6\u00d5\n\6\f\6\16\6\u00d8"+
		"\13\6\5\6\u00da\n\6\3\6\3\6\3\7\3\7\3\7\7\7\u00e1\n\7\f\7\16\7\u00e4\13"+
		"\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\7\7\u00f0\n\7\f\7\16\7\u00f3"+
		"\13\7\7\7\u00f5\n\7\f\7\16\7\u00f8\13\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3"+
		"\7\3\7\3\7\7\7\u0104\n\7\f\7\16\7\u0107\13\7\5\7\u0109\n\7\3\7\3\7\3\b"+
		"\3\b\3\b\3\b\5\b\u0111\n\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\7\b\u011c"+
		"\n\b\f\b\16\b\u011f\13\b\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\5"+
		"\t\u012c\n\t\3\t\3\t\3\t\7\t\u0131\n\t\f\t\16\t\u0134\13\t\3\n\3\n\3\n"+
		"\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\7\n\u0141\n\n\f\n\16\n\u0144\13\n\3\n"+
		"\3\n\3\13\3\13\3\13\3\13\3\13\3\13\3\13\5\13\u014f\n\13\3\13\3\13\5\13"+
		"\u0153\n\13\3\13\3\13\3\13\3\13\3\13\3\13\5\13\u015b\n\13\3\13\3\13\5"+
		"\13\u015f\n\13\3\13\3\13\3\13\3\13\3\13\3\13\5\13\u0167\n\13\3\13\3\13"+
		"\5\13\u016b\n\13\3\13\3\13\3\13\3\13\5\13\u0171\n\13\3\f\3\f\3\f\5\f\u0176"+
		"\n\f\3\f\3\f\7\f\u017a\n\f\f\f\16\f\u017d\13\f\3\f\3\f\3\f\3\f\5\f\u0183"+
		"\n\f\3\f\5\f\u0186\n\f\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r"+
		"\3\r\3\r\3\r\5\r\u0197\n\r\3\r\3\r\3\r\3\r\3\r\3\r\7\r\u019f\n\r\f\r\16"+
		"\r\u01a2\13\r\3\16\3\16\3\16\3\17\3\17\3\17\5\17\u01aa\n\17\3\17\2\4\20"+
		"\30\20\2\4\6\b\n\f\16\20\22\24\26\30\32\34\2\t\4\2\n\n\22\22\4\2\n\n\u00a6"+
		"\u00a6\4\2\4\4\b\b\4\2\u00a7\u00a7\u00ab\u00ab\3\2*+\3\2\u00b7\u00b8\3"+
		"\2\u00ae\u00af\2\u0237\2)\3\2\2\2\4-\3\2\2\2\6\u00a7\3\2\2\2\b\u00b5\3"+
		"\2\2\2\n\u00b7\3\2\2\2\f\u00dd\3\2\2\2\16\u010c\3\2\2\2\20\u012b\3\2\2"+
		"\2\22\u0135\3\2\2\2\24\u0170\3\2\2\2\26\u0185\3\2\2\2\30\u0196\3\2\2\2"+
		"\32\u01a3\3\2\2\2\34\u01a9\3\2\2\2\36*\5\4\3\2\37*\5\6\4\2 *\5\n\6\2!"+
		"*\5\f\7\2\"*\5\16\b\2#*\7\t\2\2$*\5\22\n\2%*\5\24\13\2&*\5\32\16\2\'*"+
		"\5\34\17\2(*\7\4\2\2)\36\3\2\2\2)\37\3\2\2\2) \3\2\2\2)!\3\2\2\2)\"\3"+
		"\2\2\2)#\3\2\2\2)$\3\2\2\2)%\3\2\2\2)&\3\2\2\2)\'\3\2\2\2)(\3\2\2\2*+"+
		"\3\2\2\2+)\3\2\2\2+,\3\2\2\2,\3\3\2\2\2-/\7\33\2\2.\60\7\b\2\2/.\3\2\2"+
		"\2/\60\3\2\2\2\60\64\3\2\2\2\61\63\5\24\13\2\62\61\3\2\2\2\63\66\3\2\2"+
		"\2\64\62\3\2\2\2\64\65\3\2\2\2\65\67\3\2\2\2\66\64\3\2\2\2\678\7\34\2"+
		"\28\5\3\2\2\29;\7\6\2\2:9\3\2\2\2:;\3\2\2\2;<\3\2\2\2<=\7\23\2\2=D\7$"+
		"\2\2>@\7\u00a5\2\2?A\7\u00aa\2\2@?\3\2\2\2@A\3\2\2\2AC\3\2\2\2B>\3\2\2"+
		"\2CF\3\2\2\2DB\3\2\2\2DE\3\2\2\2EH\3\2\2\2FD\3\2\2\2GI\7\b\2\2HG\3\2\2"+
		"\2HI\3\2\2\2IM\3\2\2\2JL\5\b\5\2KJ\3\2\2\2LO\3\2\2\2MK\3\2\2\2MN\3\2\2"+
		"\2NS\3\2\2\2OM\3\2\2\2PQ\7&\2\2QR\7\'\2\2RT\7\u00a5\2\2SP\3\2\2\2ST\3"+
		"\2\2\2T`\3\2\2\2U_\5\24\13\2V_\5\32\16\2W_\5\4\3\2X_\5\6\4\2Y_\5\n\6\2"+
		"Z_\5\f\7\2[_\5\16\b\2\\_\5\34\17\2]_\7\4\2\2^U\3\2\2\2^V\3\2\2\2^W\3\2"+
		"\2\2^X\3\2\2\2^Y\3\2\2\2^Z\3\2\2\2^[\3\2\2\2^\\\3\2\2\2^]\3\2\2\2_b\3"+
		"\2\2\2`^\3\2\2\2`a\3\2\2\2as\3\2\2\2b`\3\2\2\2cd\7%\2\2dp\7\26\2\2eo\5"+
		"\24\13\2fo\5\32\16\2go\5\4\3\2ho\5\6\4\2io\5\n\6\2jo\5\f\7\2ko\5\16\b"+
		"\2lo\5\34\17\2mo\7\4\2\2ne\3\2\2\2nf\3\2\2\2ng\3\2\2\2nh\3\2\2\2ni\3\2"+
		"\2\2nj\3\2\2\2nk\3\2\2\2nl\3\2\2\2nm\3\2\2\2or\3\2\2\2pn\3\2\2\2pq\3\2"+
		"\2\2qt\3\2\2\2rp\3\2\2\2sc\3\2\2\2st\3\2\2\2tu\3\2\2\2u\u00a8\7\30\2\2"+
		"vx\7\6\2\2wv\3\2\2\2wx\3\2\2\2xy\3\2\2\2yz\7\23\2\2z{\7\n\2\2{|\7\u00ab"+
		"\2\2|}\t\2\2\2}~\7\60\2\2~\u0081\t\2\2\2\177\u0080\7\61\2\2\u0080\u0082"+
		"\t\2\2\2\u0081\177\3\2\2\2\u0081\u0082\3\2\2\2\u0082\u008e\3\2\2\2\u0083"+
		"\u008d\5\24\13\2\u0084\u008d\5\32\16\2\u0085\u008d\5\4\3\2\u0086\u008d"+
		"\5\6\4\2\u0087\u008d\5\n\6\2\u0088\u008d\5\f\7\2\u0089\u008d\5\16\b\2"+
		"\u008a\u008d\5\34\17\2\u008b\u008d\7\4\2\2\u008c\u0083\3\2\2\2\u008c\u0084"+
		"\3\2\2\2\u008c\u0085\3\2\2\2\u008c\u0086\3\2\2\2\u008c\u0087\3\2\2\2\u008c"+
		"\u0088\3\2\2\2\u008c\u0089\3\2\2\2\u008c\u008a\3\2\2\2\u008c\u008b\3\2"+
		"\2\2\u008d\u0090\3\2\2\2\u008e\u008c\3\2\2\2\u008e\u008f\3\2\2\2\u008f"+
		"\u0091\3\2\2\2\u0090\u008e\3\2\2\2\u0091\u00a8\7\30\2\2\u0092\u0094\7"+
		"\6\2\2\u0093\u0092\3\2\2\2\u0093\u0094\3\2\2\2\u0094\u0095\3\2\2\2\u0095"+
		"\u0096\7\23\2\2\u0096\u0097\7\n\2\2\u0097\u0098\7\62\2\2\u0098\u00a3\t"+
		"\3\2\2\u0099\u00a2\5\24\13\2\u009a\u00a2\5\4\3\2\u009b\u00a2\5\6\4\2\u009c"+
		"\u00a2\5\n\6\2\u009d\u00a2\5\f\7\2\u009e\u00a2\5\16\b\2\u009f\u00a2\5"+
		"\34\17\2\u00a0\u00a2\7\4\2\2\u00a1\u0099\3\2\2\2\u00a1\u009a\3\2\2\2\u00a1"+
		"\u009b\3\2\2\2\u00a1\u009c\3\2\2\2\u00a1\u009d\3\2\2\2\u00a1\u009e\3\2"+
		"\2\2\u00a1\u009f\3\2\2\2\u00a1\u00a0\3\2\2\2\u00a2\u00a5\3\2\2\2\u00a3"+
		"\u00a1\3\2\2\2\u00a3\u00a4\3\2\2\2\u00a4\u00a6\3\2\2\2\u00a5\u00a3\3\2"+
		"\2\2\u00a6\u00a8\7\30\2\2\u00a7:\3\2\2\2\u00a7w\3\2\2\2\u00a7\u0093\3"+
		"\2\2\2\u00a8\7\3\2\2\2\u00a9\u00aa\7\27\2\2\u00aa\u00ac\5\20\t\2\u00ab"+
		"\u00ad\t\4\2\2\u00ac\u00ab\3\2\2\2\u00ac\u00ad\3\2\2\2\u00ad\u00b6\3\2"+
		"\2\2\u00ae\u00af\7\27\2\2\u00af\u00b0\5\20\t\2\u00b0\u00b1\7%\2\2\u00b1"+
		"\u00b3\5\20\t\2\u00b2\u00b4\t\4\2\2\u00b3\u00b2\3\2\2\2\u00b3\u00b4\3"+
		"\2\2\2\u00b4\u00b6\3\2\2\2\u00b5\u00a9\3\2\2\2\u00b5\u00ae\3\2\2\2\u00b6"+
		"\t\3\2\2\2\u00b7\u00b8\7\37\2\2\u00b8\u00ba\5\20\t\2\u00b9\u00bb\7\6\2"+
		"\2\u00ba\u00b9\3\2\2\2\u00ba\u00bb\3\2\2\2\u00bb\u00c7\3\2\2\2\u00bc\u00c6"+
		"\5\24\13\2\u00bd\u00c6\5\32\16\2\u00be\u00c6\5\4\3\2\u00bf\u00c6\5\6\4"+
		"\2\u00c0\u00c6\5\n\6\2\u00c1\u00c6\5\f\7\2\u00c2\u00c6\5\16\b\2\u00c3"+
		"\u00c6\5\34\17\2\u00c4\u00c6\7\4\2\2\u00c5\u00bc\3\2\2\2\u00c5\u00bd\3"+
		"\2\2\2\u00c5\u00be\3\2\2\2\u00c5\u00bf\3\2\2\2\u00c5\u00c0\3\2\2\2\u00c5"+
		"\u00c1\3\2\2\2\u00c5\u00c2\3\2\2\2\u00c5\u00c3\3\2\2\2\u00c5\u00c4\3\2"+
		"\2\2\u00c6\u00c9\3\2\2\2\u00c7\u00c5\3\2\2\2\u00c7\u00c8\3\2\2\2\u00c8"+
		"\u00d9\3\2\2\2\u00c9\u00c7\3\2\2\2\u00ca\u00d6\7 \2\2\u00cb\u00d5\5\24"+
		"\13\2\u00cc\u00d5\5\32\16\2\u00cd\u00d5\5\4\3\2\u00ce\u00d5\5\6\4\2\u00cf"+
		"\u00d5\5\n\6\2\u00d0\u00d5\5\f\7\2\u00d1\u00d5\5\16\b\2\u00d2\u00d5\5"+
		"\34\17\2\u00d3\u00d5\7\4\2\2\u00d4\u00cb\3\2\2\2\u00d4\u00cc\3\2\2\2\u00d4"+
		"\u00cd\3\2\2\2\u00d4\u00ce\3\2\2\2\u00d4\u00cf\3\2\2\2\u00d4\u00d0\3\2"+
		"\2\2\u00d4\u00d1\3\2\2\2\u00d4\u00d2\3\2\2\2\u00d4\u00d3\3\2\2\2\u00d5"+
		"\u00d8\3\2\2\2\u00d6\u00d4\3\2\2\2\u00d6\u00d7\3\2\2\2\u00d7\u00da\3\2"+
		"\2\2\u00d8\u00d6\3\2\2\2\u00d9\u00ca\3\2\2\2\u00d9\u00da\3\2\2\2\u00da"+
		"\u00db\3\2\2\2\u00db\u00dc\7!\2\2\u00dc\13\3\2\2\2\u00dd\u00de\7\35\2"+
		"\2\u00de\u00e2\7\31\2\2\u00df\u00e1\7\b\2\2\u00e0\u00df\3\2\2\2\u00e1"+
		"\u00e4\3\2\2\2\u00e2\u00e0\3\2\2\2\u00e2\u00e3\3\2\2\2\u00e3\u00f6\3\2"+
		"\2\2\u00e4\u00e2\3\2\2\2\u00e5\u00e6\7\31\2\2\u00e6\u00f1\5\20\t\2\u00e7"+
		"\u00f0\5\24\13\2\u00e8\u00f0\5\4\3\2\u00e9\u00f0\5\6\4\2\u00ea\u00f0\5"+
		"\n\6\2\u00eb\u00f0\5\f\7\2\u00ec\u00f0\5\16\b\2\u00ed\u00f0\5\34\17\2"+
		"\u00ee\u00f0\7\4\2\2\u00ef\u00e7\3\2\2\2\u00ef\u00e8\3\2\2\2\u00ef\u00e9"+
		"\3\2\2\2\u00ef\u00ea\3\2\2\2\u00ef\u00eb\3\2\2\2\u00ef\u00ec\3\2\2\2\u00ef"+
		"\u00ed\3\2\2\2\u00ef\u00ee\3\2\2\2\u00f0\u00f3\3\2\2\2\u00f1\u00ef\3\2"+
		"\2\2\u00f1\u00f2\3\2\2\2\u00f2\u00f5\3\2\2\2\u00f3\u00f1\3\2\2\2\u00f4"+
		"\u00e5\3\2\2\2\u00f5\u00f8\3\2\2\2\u00f6\u00f4\3\2\2\2\u00f6\u00f7\3\2"+
		"\2\2\u00f7\u0108\3\2\2\2\u00f8\u00f6\3\2\2\2\u00f9\u0105\7\64\2\2\u00fa"+
		"\u0104\5\24\13\2\u00fb\u0104\5\32\16\2\u00fc\u0104\5\4\3\2\u00fd\u0104"+
		"\5\6\4\2\u00fe\u0104\5\n\6\2\u00ff\u0104\5\f\7\2\u0100\u0104\5\16\b\2"+
		"\u0101\u0104\5\34\17\2\u0102\u0104\7\4\2\2\u0103\u00fa\3\2\2\2\u0103\u00fb"+
		"\3\2\2\2\u0103\u00fc\3\2\2\2\u0103\u00fd\3\2\2\2\u0103\u00fe\3\2\2\2\u0103"+
		"\u00ff\3\2\2\2\u0103\u0100\3\2\2\2\u0103\u0101\3\2\2\2\u0103\u0102\3\2"+
		"\2\2\u0104\u0107\3\2\2\2\u0105\u0103\3\2\2\2\u0105\u0106\3\2\2\2\u0106"+
		"\u0109\3\2\2\2\u0107\u0105\3\2\2\2\u0108\u00f9\3\2\2\2\u0108\u0109\3\2"+
		"\2\2\u0109\u010a\3\2\2\2\u010a\u010b\7\63\2\2\u010b\r\3\2\2\2\u010c\u010d"+
		"\7\35\2\2\u010d\u010e\7\"\2\2\u010e\u0110\5\20\t\2\u010f\u0111\7\6\2\2"+
		"\u0110\u010f\3\2\2\2\u0110\u0111\3\2\2\2\u0111\u011d\3\2\2\2\u0112\u011c"+
		"\5\24\13\2\u0113\u011c\5\32\16\2\u0114\u011c\5\4\3\2\u0115\u011c\5\6\4"+
		"\2\u0116\u011c\5\n\6\2\u0117\u011c\5\f\7\2\u0118\u011c\5\16\b\2\u0119"+
		"\u011c\5\34\17\2\u011a\u011c\7\4\2\2\u011b\u0112\3\2\2\2\u011b\u0113\3"+
		"\2\2\2\u011b\u0114\3\2\2\2\u011b\u0115\3\2\2\2\u011b\u0116\3\2\2\2\u011b"+
		"\u0117\3\2\2\2\u011b\u0118\3\2\2\2\u011b\u0119\3\2\2\2\u011b\u011a\3\2"+
		"\2\2\u011c\u011f\3\2\2\2\u011d\u011b\3\2\2\2\u011d\u011e\3\2\2\2\u011e"+
		"\u0120\3\2\2\2\u011f\u011d\3\2\2\2\u0120\u0121\7\36\2\2\u0121\17\3\2\2"+
		"\2\u0122\u0123\b\t\1\2\u0123\u0124\5\30\r\2\u0124\u0125\t\5\2\2\u0125"+
		"\u0126\5\30\r\2\u0126\u012c\3\2\2\2\u0127\u0128\7\u00a8\2\2\u0128\u0129"+
		"\5\20\t\2\u0129\u012a\7\u00a9\2\2\u012a\u012c\3\2\2\2\u012b\u0122\3\2"+
		"\2\2\u012b\u0127\3\2\2\2\u012c\u0132\3\2\2\2\u012d\u012e\f\4\2\2\u012e"+
		"\u012f\t\6\2\2\u012f\u0131\5\20\t\5\u0130\u012d\3\2\2\2\u0131\u0134\3"+
		"\2\2\2\u0132\u0130\3\2\2\2\u0132\u0133\3\2\2\2\u0133\21\3\2\2\2\u0134"+
		"\u0132\3\2\2\2\u0135\u0136\7(\2\2\u0136\u0142\7\13\2\2\u0137\u0141\5\24"+
		"\13\2\u0138\u0141\5\32\16\2\u0139\u0141\5\4\3\2\u013a\u0141\5\6\4\2\u013b"+
		"\u0141\5\n\6\2\u013c\u0141\5\f\7\2\u013d\u0141\5\16\b\2\u013e\u0141\5"+
		"\34\17\2\u013f\u0141\7\4\2\2\u0140\u0137\3\2\2\2\u0140\u0138\3\2\2\2\u0140"+
		"\u0139\3\2\2\2\u0140\u013a\3\2\2\2\u0140\u013b\3\2\2\2\u0140\u013c\3\2"+
		"\2\2\u0140\u013d\3\2\2\2\u0140\u013e\3\2\2\2\u0140\u013f\3\2\2\2\u0141"+
		"\u0144\3\2\2\2\u0142\u0140\3\2\2\2\u0142\u0143\3\2\2\2\u0143\u0145\3\2"+
		"\2\2\u0144\u0142\3\2\2\2\u0145\u0146\7)\2\2\u0146\23\3\2\2\2\u0147\u0171"+
		"\5\26\f\2\u0148\u014e\7\n\2\2\u0149\u014f\7\u00ab\2\2\u014a\u014b\7\u00ae"+
		"\2\2\u014b\u014f\7\u00ab\2\2\u014c\u014d\7\u00af\2\2\u014d\u014f\7\u00ab"+
		"\2\2\u014e\u0149\3\2\2\2\u014e\u014a\3\2\2\2\u014e\u014c\3\2\2\2\u014f"+
		"\u0152\3\2\2\2\u0150\u0153\5\26\f\2\u0151\u0153\5\30\r\2\u0152\u0150\3"+
		"\2\2\2\u0152\u0151\3\2\2\2\u0153\u0171\3\2\2\2\u0154\u015a\7\u00a5\2\2"+
		"\u0155\u015b\7\u00ab\2\2\u0156\u0157\7\u00ae\2\2\u0157\u015b\7\u00ab\2"+
		"\2\u0158\u0159\7\u00af\2\2\u0159\u015b\7\u00ab\2\2\u015a\u0155\3\2\2\2"+
		"\u015a\u0156\3\2\2\2\u015a\u0158\3\2\2\2\u015b\u015e\3\2\2\2\u015c\u015f"+
		"\5\26\f\2\u015d\u015f\5\30\r\2\u015e\u015c\3\2\2\2\u015e\u015d\3\2\2\2"+
		"\u015f\u0171\3\2\2\2\u0160\u0166\7\u00a6\2\2\u0161\u0167\7\u00ab\2\2\u0162"+
		"\u0163\7\u00ae\2\2\u0163\u0167\7\u00ab\2\2\u0164\u0165\7\u00af\2\2\u0165"+
		"\u0167\7\u00ab\2\2\u0166\u0161\3\2\2\2\u0166\u0162\3\2\2\2\u0166\u0164"+
		"\3\2\2\2\u0167\u016a\3\2\2\2\u0168\u016b\5\26\f\2\u0169\u016b\5\30\r\2"+
		"\u016a\u0168\3\2\2\2\u016a\u0169\3\2\2\2\u016b\u0171\3\2\2\2\u016c\u016d"+
		"\7\35\2\2\u016d\u0171\7\13\2\2\u016e\u0171\7\32\2\2\u016f\u0171\7-\2\2"+
		"\u0170\u0147\3\2\2\2\u0170\u0148\3\2\2\2\u0170\u0154\3\2\2\2\u0170\u0160"+
		"\3\2\2\2\u0170\u016c\3\2\2\2\u0170\u016e\3\2\2\2\u0170\u016f\3\2\2\2\u0171"+
		"\25\3\2\2\2\u0172\u0173\7\65\2\2\u0173\u0175\7\u00a8\2\2\u0174\u0176\5"+
		"\30\r\2\u0175\u0174\3\2\2\2\u0175\u0176\3\2\2\2\u0176\u017b\3\2\2\2\u0177"+
		"\u0178\7\u00aa\2\2\u0178\u017a\5\30\r\2\u0179\u0177\3\2\2\2\u017a\u017d"+
		"\3\2\2\2\u017b\u0179\3\2\2\2\u017b\u017c\3\2\2\2\u017c\u017e\3\2\2\2\u017d"+
		"\u017b\3\2\2\2\u017e\u0186\7\u00a9\2\2\u017f\u0180\7\u00a6\2\2\u0180\u0182"+
		"\7\u00a8\2\2\u0181\u0183\5\30\r\2\u0182\u0181\3\2\2\2\u0182\u0183\3\2"+
		"\2\2\u0183\u0184\3\2\2\2\u0184\u0186\7\u00a9\2\2\u0185\u0172\3\2\2\2\u0185"+
		"\u017f\3\2\2\2\u0186\27\3\2\2\2\u0187\u0188\b\r\1\2\u0188\u0197\7\n\2"+
		"\2\u0189\u0197\7\13\2\2\u018a\u0197\7\22\2\2\u018b\u0197\7\u00a5\2\2\u018c"+
		"\u0197\5\26\f\2\u018d\u0197\7\u00a6\2\2\u018e\u018f\7\33\2\2\u018f\u0190"+
		"\7\u00a5\2\2\u0190\u0191\7\u00a8\2\2\u0191\u0197\7\u00a9\2\2\u0192\u0193"+
		"\7\u00a8\2\2\u0193\u0194\5\30\r\2\u0194\u0195\7\u00a9\2\2\u0195\u0197"+
		"\3\2\2\2\u0196\u0187\3\2\2\2\u0196\u0189\3\2\2\2\u0196\u018a\3\2\2\2\u0196"+
		"\u018b\3\2\2\2\u0196\u018c\3\2\2\2\u0196\u018d\3\2\2\2\u0196\u018e\3\2"+
		"\2\2\u0196\u0192\3\2\2\2\u0197\u01a0\3\2\2\2\u0198\u0199\f\5\2\2\u0199"+
		"\u019a\t\7\2\2\u019a\u019f\5\30\r\6\u019b\u019c\f\4\2\2\u019c\u019d\t"+
		"\b\2\2\u019d\u019f\5\30\r\5\u019e\u0198\3\2\2\2\u019e\u019b\3\2\2\2\u019f"+
		"\u01a2\3\2\2\2\u01a0\u019e\3\2\2\2\u01a0\u01a1\3\2\2\2\u01a1\31\3\2\2"+
		"\2\u01a2\u01a0\3\2\2\2\u01a3\u01a4\7#\2\2\u01a4\u01a5\7\u00a5\2\2\u01a5"+
		"\33\3\2\2\2\u01a6\u01aa\7\5\2\2\u01a7\u01aa\7\7\2\2\u01a8\u01aa\7\b\2"+
		"\2\u01a9\u01a6\3\2\2\2\u01a9\u01a7\3\2\2\2\u01a9\u01a8\3\2\2\2\u01aa\35"+
		"\3\2\2\2?)+/\64:@DHMS^`npsw\u0081\u008c\u008e\u0093\u00a1\u00a3\u00a7"+
		"\u00ac\u00b3\u00b5\u00ba\u00c5\u00c7\u00d4\u00d6\u00d9\u00e2\u00ef\u00f1"+
		"\u00f6\u0103\u0105\u0108\u0110\u011b\u011d\u012b\u0132\u0140\u0142\u014e"+
		"\u0152\u015a\u015e\u0166\u016a\u0170\u0175\u017b\u0182\u0185\u0196\u019e"+
		"\u01a0\u01a9";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}