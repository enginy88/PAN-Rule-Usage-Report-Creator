package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
)

var AppNameString string

type AppFlagStruct struct {
	PanoramaFlag bool
	VerboseFlag  bool
}

var appFlag *AppFlagStruct

var (
	LogErr    *log.Logger
	LogWarn   *log.Logger
	LogInfo   *log.Logger
	LogAlways *log.Logger
)

func init() {

	bi, ok := debug.ReadBuildInfo()
	if !ok {
		log.Fatalln("FATAL ERROR: Failed to read build info! Please build the binary with module support.")
	}
	AppNameString = bi.Path

	LogErr = log.New(os.Stderr, "("+AppNameString+") ERROR: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	LogWarn = log.New(os.Stdout, "("+AppNameString+") WARNING: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	LogInfo = log.New(os.Stdout, "("+AppNameString+") INFO: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	LogAlways = log.New(os.Stdout, "("+AppNameString+") ALWAYS: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)

}

func GetAppFlag() *AppFlagStruct {

	appFlagObiect := new(AppFlagStruct)
	appFlag = appFlagObiect

	parseAppFlag()

	return appFlag

}

func parseAppFlag() {

	verboseFlag := flag.Bool("verbose", false, "Can be set to introduce verbose logs. (Optional)  [Default: Unset]")
	panoramaFlag := flag.Bool("panorama", false, "Can be set to use Panorama DB for creating report. (Optional)  [Default: Unset]")

	flag.Parse()

	appFlag.PanoramaFlag = *panoramaFlag
	appFlag.VerboseFlag = *verboseFlag

}

func main() {

	start := time.Now()

	appFlag = GetAppFlag()

	if !appFlag.VerboseFlag {
		LogWarn.SetOutput(io.Discard)
		LogInfo.SetOutput(io.Discard)
		LogAlways.SetOutput(io.Discard)
	}

	uuids := parseInput()

	LogAlways.Println("HELLO MSG: Welcome to " + AppNameString + " v1.0 by EY!")

	processUUIDs(uuids)

	duration := fmt.Sprintf("%.1f", time.Since(start).Seconds())
	LogAlways.Println("BYE MSG: All done in " + duration + "s, bye!")

}

func parseInput() (uuids []string) {

	stdin, err := io.ReadAll(os.Stdin)
	if err != nil {
		LogErr.Fatalln("FATAL ERROR: Cannot read from stdin! (" + err.Error() + ")")
	}

	count := 0
	str := strings.ReplaceAll(string(stdin), "\r\n", "\n")
	for _, line := range strings.Split(str, "\n") {
		uuid := strings.TrimSpace(line)
		if len(uuid) > 0 {
			uuids = append(uuids, uuid)
			count++
		}
	}
	if count == 0 {
		LogErr.Fatalln("FATAL ERROR: Cannot find any input to process!")
	}

	LogInfo.Println("VERBOSE: Total Parsed Rule UUIDs: " + strconv.Itoa(count))

	return uuids
}

func processUUIDs(uuids []string) {

	if len(uuids) == 0 {
		LogErr.Fatalln("FATAL ERROR: UUID value is empty!")
	}

	LogAlways.Println("RESULT: Created Custom Report Config: ")

	reportCount := 0
	queryStr := ""
	for n, uuid := range uuids {
		if !isValidUUID(uuid) {
			LogErr.Fatalln("FATAL ERROR: Input is not in expected UUID format! (" + uuid + ")")
		}
		if n%36 == 0 {
			queryStr = "(rule_uuid eq '" + uuid + "')"
			reportCount++
		} else {
			queryStr = queryStr + " or (rule_uuid eq '" + uuid + "')"
		}
		if (n%36 == 35) || (n == len(uuids)-1) {
			createReport(reportCount, queryStr)
		}
	}

	LogInfo.Println("VERBOSE: Total Created Report Config: " + strconv.Itoa(reportCount))

}

func createReport(reportCount int, queryStr string) {

	typeStr := ""
	if appFlag.PanoramaFlag {
		typeStr += "panorama-traffic"
	} else {
		typeStr += "traffic"
	}

	setCommandPrefix := fmt.Sprint("set shared reports X-Custom-Usage-Report-" + strconv.Itoa(reportCount) + " ")

	fmt.Println(setCommandPrefix + "caption X-Custom-Usage-Report-" + strconv.Itoa(reportCount))

	fmt.Println(setCommandPrefix + "type " + typeStr + " aggregate-by [ rule rule_uuid action device_name vsys_name ]")
	fmt.Println(setCommandPrefix + "type " + typeStr + " values [ packets bytes repeatcnt ]")
	fmt.Println(setCommandPrefix + "type " + typeStr + " sortby repeatcnt")
	fmt.Println(setCommandPrefix + "period last-7-days")
	fmt.Println(setCommandPrefix + "topn 1000")
	fmt.Println(setCommandPrefix + "topm 1000")
	fmt.Println(setCommandPrefix + "query \"" + queryStr + "\"")

}

// xvalues returns the value of a byte as a hexadecimal digit or 255.
var xvalues = [256]byte{
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 255, 255, 255, 255, 255, 255,
	255, 10, 11, 12, 13, 14, 15, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 10, 11, 12, 13, 14, 15, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
}

// xtob converts hex characters x1 and x2 into a byte.
func xtob(x1, x2 byte) (byte, bool) {
	b1 := xvalues[x1]
	b2 := xvalues[x2]
	return (b1 << 4) | b2, b1 != 255 && b2 != 255
}

// validate returns an error if s is not a properly formatted UUID in one of the following format: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
func isValidUUID(s string) bool {

	if len(s) != 36 {
		return false
	}

	if s[8] != '-' || s[13] != '-' || s[18] != '-' || s[23] != '-' {
		return false
	}
	for _, x := range []int{0, 2, 4, 6, 9, 11, 14, 16, 19, 21, 24, 26, 28, 30, 32, 34} {
		if _, ok := xtob(s[x], s[x+1]); !ok {
			return false
		}
	}

	return true

}
