package uaparser

import (
	"regexp"
	"sync/atomic"
)

type UserAgentSorter []*uaParser

func (a UserAgentSorter) Len() int      { return len(a) }
func (a UserAgentSorter) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a UserAgentSorter) Less(i, j int) bool {
	return atomic.LoadUint64(&a[i].MatchesCount) > atomic.LoadUint64(&a[j].MatchesCount)
}

type uaParser struct {
	Reg                *regexp.Regexp
	Expr               string  `yaml:"regex"`
	Flags              string  `yaml:"regex_flag"`
	FamilyReplacement  string  `yaml:"family_replacement"`
	VersionReplacement string  `yaml:"version_replacement"`
	_                  [4]byte // padding for alignment
	MatchesCount       uint64
}

func (parser *uaParser) setDefaults() {
	if parser.FamilyReplacement == "" {
		parser.FamilyReplacement = "$1"
	}
	if parser.VersionReplacement == "" {
		parser.VersionReplacement = "$2"
	}
}

func (parser *uaParser) Match(line string, ua *UserAgent) {
	matches := parser.Reg.FindStringSubmatchIndex(line)
	if len(matches) > 0 {
		ua.Family = string(parser.Reg.ExpandString(nil, parser.FamilyReplacement, line, matches))
		ua.Version = string(parser.Reg.ExpandString(nil, parser.VersionReplacement, line, matches))
	}
}

type UserAgent struct {
	Family  string
	Version string
}

func (ua *UserAgent) ToString() string {
	var str string
	if ua.Family != "" {
		str += ua.Family
	}
	if ua.Version != "" {
		str += " " + ua.Version
	}
	return str
}
