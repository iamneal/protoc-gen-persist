package parser

import (
	"bufio"
	"fmt"
	"io"
)

type Parser struct {
	s *Scanner
}

func NewParser(r io.Reader) *Parser {
	buf := bufio.NewReader(r)
	s := NewScanner(buf)

	return &Parser{s: s}
}

func (p *Parser) Parse() (Query, error) {
	q, err := (&StartMode{}).Parse(p.s)
	if err != nil {
		return nil, fmt.Errorf("Failure parsing query.\n err: %v", err)
	}
	return q, nil
}

type Mode interface {
	Parse(s *Scanner) (Query, error)
}

type StartMode struct{}

func (m *StartMode) Parse(scanner *Scanner) (Query, error) {
	tkn := scanner.Peek(1)[0]
	switch tkn.tk {
	case SELECT:
		return NewSelectMode().Parse(scanner)
	case INSERT:
		return NewInsertMode().Parse(scanner)
	case UPDATE:
		return NewUpdateMode().Parse(scanner)
	case DELETE:
		return NewDeleteMode().Parse(scanner)
	case EOF:
		return nil, fmt.Errorf("no query found")
	case ILLEGAL:
		return nil, fmt.Errorf("scanned illegal token: %s", tkn.raw)
	default:
		return nil, fmt.Errorf("token kind %d, raw: '%s' not acceptable here", tkn.tk, tkn.raw)
	}
}

// we don't parse anything for select mode
type SelectMode struct {
	tkns []*Token
}

func NewSelectMode() *SelectMode {
	return &SelectMode{}
}
func (m *SelectMode) Parse(scanner *Scanner) (Query, error) {
	// we have scanned out the select word, so add it back
	fields := make([]*Token, 0)
	query := ""
	eater := NewEater(scanner)
	if eater.Eat(SELECT) {
		query += eater.TakeTokens()[0].raw
	} else {
		return nil, fmt.Errorf("told to parse a select query, but SELECT not found")
	}
	for {
		ch, stop := scanner.Read()
		if stop {
			break
		}
		// ignore everything except fields, which will start with @
		if ch == '@' {
			scanner.Unread()
			if eater.Eat(IDENT_FIELD, IDENT_DIRECTIVE) {
				token := eater.TakeTokens()[0]
				query += token.raw
				if token.tk == IDENT_FIELD {
					fields = append(fields, token)
				}
			}
			if err := eater.TakeErr(); err != nil {
				return nil, fmt.Errorf("could not parse field name for select query: %s", err)
			}
		} else {
			// just a regular character part of the query
			query += string(ch)
		}
	}
	return NewSelectQuery(query, fields), nil
}

type InsertMode struct{}

func NewInsertMode() *InsertMode {
	return &InsertMode{}
}

//INSERT INTO tablename <optional> (colname, colname) VALUES(...)
func (m *InsertMode) Parse(scanner *Scanner) (Query, error) {
	var table *Token
	var cols []*Token
	var values []*Token

	eater := NewEater(scanner)
	eater.Eat(INSERT)
	eater.Eat(INTO)
	if eater.Eat(IDENT_TABLE_OR_COL) {
		table = eater.Top()
	}
	// we have an optional array here
	if scanner.Peek(1)[0].tk == OPEN_PARAN {
		// declaring column ordering here
		cols, _ = eater.EatArrayOf(IDENT_TABLE_OR_COL)
	}
	eater.Eat(VALUES)
	values, _ = eater.EatArrayOf(
		IDENT_STRING,
		IDENT_FLOAT,
		IDENT_INT,
		IDENT_FIELD,
		IDENT_BOOL,
	)
	if !eater.Eat(EOF) {
		return nil, eater.TakeErr()
	}

	// validate columns
	if len(cols) != len(values) {
		return nil, fmt.Errorf(
			"columns len: %d, does not match values len: %d",
			len(cols),
			len(values),
		)
	}
	return &InsertQuery{
		tokens:    eater.TakeTokens(),
		cols:      cols,
		values:    values,
		tableName: table,
	}, nil
}

type UpdateMode struct{}

func NewUpdateMode() *UpdateMode {
	return &UpdateMode{}
}

// Normal Update sql query syntax does not really make a ton of sense here
// for now, just support:
// Update table <set loop> | <(col), VALUES (values)> PK(IDENT)
func (m *UpdateMode) Parse(scanner *Scanner) (Query, error) {
	var cols []*Token
	var values []*Token
	var table *Token
	var primaryKey = make(map[Token]*Token)

	eater := NewEater(scanner)
	eater.Eat(UPDATE)
	if eater.Eat(IDENT_TABLE_OR_COL) {
		table = eater.Top()
	}

	//either array of col names, or set loop
	switch scanner.Peek(1)[0].tk {
	case OPEN_PARAN:
		cs, _ := eater.EatArrayOf(IDENT_TABLE_OR_COL)
		cols = append(cols, cs...)
		eater.Eat(VALUES)
		vs, _ := eater.EatArrayOf(
			IDENT_STRING,
			IDENT_FLOAT,
			IDENT_INT,
			IDENT_FIELD,
			IDENT_BOOL,
		)
		values = append(values, vs...)
	case SET:
		eater.Eat(SET)
		groups, _ := eater.EatCommaSeperatedGroupOf(
			IDENT_TABLE_OR_COL,
			EQUAL_OP,
			Group(
				IDENT_STRING,
				IDENT_FLOAT,
				IDENT_INT,
				IDENT_FIELD,
				IDENT_BOOL,
			),
		)
		for _, g := range groups {
			if len(g) > 2 {
				cols = append(cols, g[0])
				values = append(values, g[2])
			}
		}
	}
	eater.Eat(PRIMARY_KEY)
	eater.Eat(OPEN_PARAN)
	// each index in groups will contain []{table/col, =, ident}
	groups, _ := eater.EatCommaSeperatedGroupOf(
		IDENT_TABLE_OR_COL,
		EQUAL_OP,
		Group(
			IDENT_STRING,
			IDENT_FLOAT,
			IDENT_INT,
			IDENT_FIELD,
			IDENT_BOOL,
		),
	)
	eater.Eat(CLOSE_PARAN)
	for _, g := range groups {
		if len(g) > 2 {
			primaryKey[*g[0]] = g[2]
		}
	}

	eater.Eat(EOF)
	if err := eater.TakeErr(); err != nil {
		return nil, err
	}
	return &UpdateQuery{
		tokens:    eater.TakeTokens(),
		cols:      cols,
		values:    values,
		tableName: table,
		pk:        primaryKey,
	}, nil
}

type DeleteMode struct {
}

func NewDeleteMode() *DeleteMode {
	return &DeleteMode{}
}

// supports two formats:
// delete key range: DELETE FROM table_name START(...) END(...) KIND(...)
// delete record pk: DELETE FROM table_name (...) VALUES (...) PRIMARY_KEY(...)
func (m *DeleteMode) Parse(scanner *Scanner) (Query, error) {
	var table *Token
	var kind *Token
	var start []*Token
	var end []*Token
	var values []*Token
	var usesKeyRange bool

	eater := NewEater(scanner)
	eater.Eat(DELETE)
	eater.Eat(FROM)
	if eater.Eat(IDENT_TABLE_OR_COL) {
		table = eater.Top()
	}
	switch scanner.Peek(1)[0].tk {
	case START: // is a key range query
		usesKeyRange = true
		eater.Eat(START)
		s, _ := eater.EatArrayOf(
			IDENT_STRING,
			IDENT_FLOAT,
			IDENT_INT,
			IDENT_FIELD,
			IDENT_BOOL,
		)
		start = append(start, s...)

		eater.Eat(END)
		e, _ := eater.EatArrayOf(
			IDENT_STRING,
			IDENT_FLOAT,
			IDENT_INT,
			IDENT_FIELD,
			IDENT_BOOL,
		)
		end = append(end, e...)

		eater.Eat(KIND)
		eater.Eat(OPEN_PARAN)
		if eater.Eat(
			CLOSED_OPEN_KIND,
			CLOSED_CLOSED_KIND,
			OPEN_OPEN_KIND,
			OPEN_CLOSED_KIND,
		) {
			kind = eater.Top()
		}
		eater.Eat(CLOSE_PARAN)
	case VALUES: // is a delete single record query
		eater.Eat(VALUES)
		vs, _ := eater.EatArrayOf(
			IDENT_STRING,
			IDENT_FLOAT,
			IDENT_INT,
			IDENT_FIELD,
			IDENT_BOOL,
		)
		values = append(values, vs...)
	}
	eater.Eat(EOF)

	if err := eater.TakeErr(); err != nil {
		return nil, err
	}

	return &DeleteQuery{
		tokens:       eater.TakeTokens(),
		start:        start,
		end:          end,
		kind:         kind,
		table:        table,
		usesKeyRange: usesKeyRange,
		values:       values,
	}, nil
}
