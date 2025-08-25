package uaparser

type ParserOption func(*Parser)

func WithMissesThreshold(threshold uint64) ParserOption {
	return func(parser *Parser) {
		if threshold > cMinMissesThreshold {
			parser.missesThreshold = threshold
		}
	}
}

func WithMatchIdxNotOk(topCnt int) ParserOption {
	return func(parser *Parser) {
		if topCnt >= 0 {
			parser.matchIdxNotOk = topCnt
		}
	}
}

func WithUseSort(useSort bool) ParserOption {
	return func(parser *Parser) {
		parser.UseSort = useSort
	}
}

func WithDebugMode(debugMode bool) ParserOption {
	return func(parser *Parser) {
		parser.debugMode = debugMode
	}
}

func WithLogger(logger Logger) ParserOption {
	return func(parser *Parser) {
		if logger != nil {
			parser.logger = logger
		}
	}
}
