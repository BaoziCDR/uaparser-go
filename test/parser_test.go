package uaparser_test

import (
	"log"
	"strings"
	"testing"

	"github.com/BaoziCDR/uaparser-go"
)

var testParser *uaparser.Parser

func init() {
	var err error
	testParser, err = uaparser.NewFromSaved()
	if err != nil {
		log.Fatal(err)
	}
}

func TestReadsInternalYAML(t *testing.T) {
	_, err := uaparser.NewFromSaved() // should not error
	if err != nil {
		log.Fatal(err)
	}
}

func TestUaParse(t *testing.T) {
	tests := []struct {
		uaInput         string
		expectedOutput  string
		expectedVersion string
	}{
		{
			"Mozilla/5.0 (Linux; Android 9; V1901A Build/P00610; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/87.0.4280.141 Mobile Safari/537.36 VivoBrowser/18.6.0.0",
			"vivobrowser",
			"18.6.0.0",
		},
		{
			"Mozilla/5.0 (Linux; U; Android 8.1.0; zh-cn; PBAT00 Build/OPM1.171019.026) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/70.0.3538.80 Mobile Safari/537.36 HeyTapBrowser/10.7.21.2",
			"heytapbrowser",
			"10.7.21.2",
		},
		{
			"Mozilla/5.0 (Linux; Android 10; HarmonyOS; STK-AL00; HMSCore 6.13.0.302) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.88 HuaweiBrowser/14.0.7.302 Mobile Safari/537.36",
			"huaweibrowser",
			"14.0.7.302",
		},
		{
			"Mozilla/5.0 (Linux; U; Android 13; zh-cn; M2104K10AC Build/TP1A.220624.014) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/100.0.4896.127 Mobile Safari/537.36 XiaoMi/MiuiBrowser/17.6.70714 swan-mibrowser",
			"miuibrowser",
			"17.6.70714",
		},
		{
			"Mozilla/5.0 (Linux; Android 14; SAMSUNG SM-S9080) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/22.1 Chrome/111.0.5563.116 Mobile Safari/537.36",
			"samsungbrowser",
			"22.1",
		},
		{
			"Mozilla/5.0 (Linux; U; Android 14; zh-cn; SM-S9080 Build/UP1A.231005.007) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/117.0.0.0 MQQBrowser/14.8 Mobile Safari/537.36",
			"qqbrowser",
			"14.8",
		},
		{
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36 Edg/128.0.0.0",
			"edge",
			"128.0.0.0",
		},
		{
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:131.0) Gecko/20100101 Firefox/131.0",
			"firefox",
			"131.0",
		},
		{
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36 OPR/113.0.0.0",
			"opera",
			"113.0.0.0",
		},
		{
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.1 Safari/537.36",
			"chrome",
			"128.0.0.1",
		},
		{
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.1 Safari/605.1.15",
			"safari",
			"16.1", // 605.1.15 => 16.1
		},
	}

	for _, tt := range tests {
		t.Run(tt.expectedOutput, func(t *testing.T) {
			result := testParser.Parse(tt.uaInput)
			retName := strings.ReplaceAll(strings.ToLower(result.Family), " ", "")
			wantName := strings.ReplaceAll(strings.ToLower(tt.expectedOutput), " ", "")
			if retName != wantName || result.Version != tt.expectedVersion {
				t.Errorf("parsed(%s) = %s, %s; want %s, %s", tt.uaInput, result.Family, result.Version, tt.expectedOutput, tt.expectedVersion)
			}
		})
	}
}
