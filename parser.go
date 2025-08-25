package uaparser

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"sync"
	"sync/atomic"
	"time"

	"gopkg.in/yaml.v2"
)

type RegexesDefinitions struct {
	UA []*uaParser `yaml:"user_agent_parsers"`
	_  [4]byte     // padding for alignment
	sync.RWMutex
}

type Parser struct {
	/* atomic operation are done on the following unit64.
	 * These must be 64bit aligned. On 32bit architectures
	 * this is only guaranteed to be on the beginning of a struct */
	UserAgentMisses uint64

	cache *cache

	RegexesDefinitions
	UseSort         bool
	debugMode       bool
	missesThreshold uint64
	matchIdxNotOk   int
	logger          Logger
}

const (
	cMinMissesThreshold     = 100000
	cDefaultMissesThreshold = 500000
	cDefaultMatchIdxNotOk   = 20
	cDefaultSortOption      = false
)

func (parser *Parser) mustCompile() { // until we can use yaml.UnmarshalYAML with embedded pointer struct
	for _, p := range parser.UA {
		p.Reg = compileRegex(p.Flags, p.Expr)
		p.setDefaults()
	}
}

func New(data []byte, opts ...ParserOption) (*Parser, error) {
	parser, err := newFromBytes(data)
	if err != nil {
		return nil, err
	}
	parser.matchIdxNotOk = cDefaultMatchIdxNotOk
	parser.missesThreshold = cDefaultMissesThreshold
	parser.UseSort = cDefaultSortOption
	for _, opt := range opts {
		opt(parser)
	}
	return parser, nil
}

func NewFromFile(regexFile string, opts ...ParserOption) (*Parser, error) {
	data, err := os.ReadFile(regexFile)
	if nil != err {
		return nil, err
	}
	parser, err := newFromBytes(data)
	if err != nil {
		return nil, err
	}
	parser.matchIdxNotOk = cDefaultMatchIdxNotOk
	parser.missesThreshold = cDefaultMissesThreshold
	parser.UseSort = cDefaultSortOption
	for _, opt := range opts {
		opt(parser)
	}
	return parser, nil
}

func NewFromSaved(opts ...ParserOption) (*Parser, error) {
	parser, err := newFromBytes(DefinitionYaml)
	if err != nil {
		return nil, err
	}
	parser.matchIdxNotOk = cDefaultMatchIdxNotOk
	parser.missesThreshold = cDefaultMissesThreshold
	parser.UseSort = cDefaultSortOption
	for _, opt := range opts {
		opt(parser)
	}
	return parser, nil
}

func newFromBytes(data []byte) (*Parser, error) {
	parser := &Parser{
		cache:  newCache(),
		logger: NewNoOpLogger(), // Default to no-op logger
	}
	if err := yaml.Unmarshal(data, &parser.RegexesDefinitions); err != nil {
		return nil, err
	}

	parser.mustCompile()

	return parser, nil
}

func (parser *Parser) Parse(line string) *UserAgent {
	cli := new(UserAgent)
	parser.RLock()
	cli = parser.parseUserAgent(line)
	parser.RUnlock()
	if parser.UseSort {
		parser.checkAndSort()
	}
	return cli
}

func (parser *Parser) parseUserAgent(line string) *UserAgent {
	cachedUA, ok := parser.cache.userAgent.Get(line)
	if ok {
		return cachedUA.(*UserAgent)
	}
	ua := new(UserAgent)
	foundIdx := -1
	found := false
	for i, uaPattern := range parser.UA {
		uaPattern.Match(line, ua)
		if len(ua.Family) > 0 {
			found = true
			foundIdx = i
			atomic.AddUint64(&uaPattern.MatchesCount, 1)
			if parser.debugMode && parser.logger != nil {
				parser.logger.Infof("[match ua]\t%s\t%s\t[expr]\t%s", line, ua.Family, uaPattern.Expr)
			}
			break
		}
	}
	if !found {
		ua.Family = "Other"
	}
	if foundIdx > parser.matchIdxNotOk {
		atomic.AddUint64(&parser.UserAgentMisses, 1)
	}
	if parser.debugMode && parser.logger != nil {
		if !found {
			parser.logger.Infof("[not match ua]\t%s\n", line)
		}
	}
	parser.cache.userAgent.Add(line, ua)
	return ua
}

func (parser *Parser) checkAndSort() {
	parser.Lock()
	if atomic.LoadUint64(&parser.UserAgentMisses) >= parser.missesThreshold {
		if parser.debugMode && parser.logger != nil {
			parser.logger.Infof("%s\tSorting UserAgents slice\n", time.Now())
		}
		parser.UserAgentMisses = 0
		sort.Sort(UserAgentSorter(parser.UA))
	}
	parser.Unlock()
}

func compileRegex(flags, expr string) *regexp.Regexp {
	if flags == "" {
		return regexp.MustCompile(expr)
	} else {
		return regexp.MustCompile(fmt.Sprintf("(?%s)%s", flags, expr))
	}
}
