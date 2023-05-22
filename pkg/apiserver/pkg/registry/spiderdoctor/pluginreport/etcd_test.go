package pluginreport

import (
	"encoding/json"
	"fmt"
	"github.com/spidernet-io/spiderdoctor/pkg/k8s/apis/system/v1beta1"
	"io"
	"os"
	"strings"
	"testing"
	"time"
)

func TestABC(t *testing.T) {
	fileDir := "/Users/icaruswu/gocode/src/spiderdoctor/pkg/apiserver/pkg/registry/spiderdoctor/pluginreport/Nethttp_test-agent_round1_spiderdoctor-control-plane_2023-06-22T07:09:20Z"
	file, err := os.Open(fileDir)
	if nil != err {
		panic(err)
	}
	readAll, err := io.ReadAll(file)
	if nil != err {
		panic(err)
	}

	pluginReportSpec := &v1beta1.PluginReportSpec{}
	err = json.Unmarshal(readAll, &pluginReportSpec)
	if nil != err {
		panic(err)
	}
	fmt.Println(pluginReportSpec)

	pluginReport := &v1beta1.PluginReport{}
	err = json.Unmarshal(readAll, &(pluginReport.Spec))
	if nil != err {
		panic(err)
	}
	fmt.Println(pluginReport)
}

func TestB(t *testing.T) {
	fileName := "Nethttp_test-agent_round1_spiderdoctor-control-plane_2023-06-22T07:09:20Z"
	split := strings.Split(fileName, "_")

	timeStr := split[len(split)-1]
	//fmt.Println(timeStr)
	times, err := time.Parse(time.RFC3339, timeStr)
	if nil != err {
		panic(err)
	}

	fmt.Println(times.Unix())

}
