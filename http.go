package main

import (
	"github.com/valyala/fasthttp"
	"github.com/valyala/fastjson"
	"crypto/tls"
	"strconv"
	"time"
	"fmt"
	"os"
)

func createTLSClient() *fasthttp.Client {
	return &fasthttp.Client {
		TLSConfig: 						&tls.Config{InsecureSkipVerify: true, SessionTicketsDisabled: true, DynamicRecordSizingDisabled: true, PreferServerCipherSuites: true},
		MaxConnsPerHost: 				threads + 250,
		DisablePathNormalizing: 		true,
		NoDefaultUserAgentHeader: 		true,
		DisableHeaderNamesNormalizing: 	true,
	}
}

func usernameAvailable(username string) bool {
	var response *fasthttp.Response = fasthttp.AcquireResponse()
	var request *fasthttp.Request = fasthttp.AcquireRequest()
	var xKhronos int64 = time.Now().Unix()

	defer fasthttp.ReleaseResponse(response)
	defer fasthttp.ReleaseRequest(request)

	request.Header.SetMethod("GET")
	request.SetRequestURI("http://api31-normal-q-useast2a.tiktokv.com/aweme/v1/user/uniqueid/?iid=7091502993825253162&device_id=7091503001088591406&ac=wifi&channel=beta&aid=1233&app_name=musical_ly&version_code=240216&version_name=24.2.16&device_platform=android&ab_version=24.2.16&ssmix=a&device_type=G011A&device_brand=google&language=en&os_api=25&os_version=7.1.2&openudid=96a943134babc0db&manifest_version_code=2022402160&resolution=900*1600&dpi=300&update_version_code=2022402160&_rticket=1651260864258&current_region=US&app_type=normal&sys_region=US&mcc_mnc=31004&timezone_name=America%2FChicago&residence=US&ts=1651260864&timezone_offset=-21600&build_number=24.2.16&region=US&uoo=0&app_language=en&carrier_region=US&locale=en&op_region=US&ac2=wifi&host_abi=x86&cdid=b65dbfb4-3dbf-436e-9b16-0863ccda5a93&id=" + username)

	request.Header.Set("X-Gorgon", generateXGorgon(xKhronos, convertBytes(generateSigHash("iid=7091502993825253162&device_id=7091503001088591406&ac=wifi&channel=beta&aid=1233&app_name=musical_ly&version_code=240216&version_name=24.2.16&device_platform=android&ab_version=24.2.16&ssmix=a&device_type=G011A&device_brand=google&language=en&os_api=25&os_version=7.1.2&openudid=96a943134babc0db&manifest_version_code=2022402160&resolution=900*1600&dpi=300&update_version_code=2022402160&_rticket=1651260864258&current_region=US&app_type=normal&sys_region=US&mcc_mnc=31004&timezone_name=America%2FChicago&residence=US&ts=1651260864&timezone_offset=-21600&build_number=24.2.16&region=US&uoo=0&app_language=en&carrier_region=US&locale=en&op_region=US&ac2=wifi&host_abi=x86&cdid=b65dbfb4-3dbf-436e-9b16-0863ccda5a93&id=" + username))))
	request.Header.Set("X-SS-REQ-TICKET", strconv.FormatInt(time.Now().Unix() * 1000, 10))
	request.Header.Set("X-Khronos", strconv.FormatInt(xKhronos, 10))
	
	request.Header.Set("User-Agent", "com.zhiliaoapp.musically/2022402160 (Linux; U; Android 7.1.2; en; G011A; Build/N2G48H;tt-ok/3.10.0.2)")
	request.Header.Set("Cache-Control", "no-cache")
	request.Header.Set("Connection", "keep-alive")
	request.Header.Set("Cookie", generateSession())

	httpClient.Do(request, response)

	if len(response.Body()) == 0 {
		rl++
		if rl > 500 {
			fmt.Printf("[%s] Rate limited after %s attempts\n", green("Akuma"), pink(formatNumber(attempts)))
			os.Exit(0)
		}
	}
	return fastjson.GetInt(response.Body(), "status_code") == 8196
}

func claimUsername(username string) bool {
	var response *fasthttp.Response = fasthttp.AcquireResponse()
	var request *fasthttp.Request = fasthttp.AcquireRequest()
	var xKhronos int64 = time.Now().Unix()

	defer fasthttp.ReleaseResponse(response)
	defer fasthttp.ReleaseRequest(request)

	request.Header.SetMethod("POST")
	request.SetRequestURI("http://api16-normal-useast5.us.tiktokv.com/passport/login_name/update/?iid=7091502993825253162&device_id=7091503001088591406&ac=wifi&channel=beta&aid=1233&app_name=musical_ly&version_code=240216&version_name=24.2.16&device_platform=android&ab_version=24.2.16&ssmix=a&device_type=G011A&device_brand=google&language=en&os_api=25&os_version=7.1.2&openudid=96a943134babc0db&manifest_version_code=2022402160&resolution=900*1600&dpi=300&update_version_code=2022402160&_rticket=1651260864258&current_region=US&app_type=normal&sys_region=US&mcc_mnc=31004&timezone_name=America%2FChicago&residence=US&ts=1651260864&timezone_offset=-21600&build_number=24.2.16&region=US&uoo=0&app_language=en&carrier_region=US&locale=en&op_region=US&ac2=wifi&host_abi=x86&cdid=b65dbfb4-3dbf-436e-9b16-0863ccda5a93&login_name=" + username)

	request.Header.Set("X-Gorgon", generateXGorgon(xKhronos, convertBytes(generateSigHash("iid=7091502993825253162&device_id=7091503001088591406&ac=wifi&channel=beta&aid=1233&app_name=musical_ly&version_code=240216&version_name=24.2.16&device_platform=android&ab_version=24.2.16&ssmix=a&device_type=G011A&device_brand=google&language=en&os_api=25&os_version=7.1.2&openudid=96a943134babc0db&manifest_version_code=2022402160&resolution=900*1600&dpi=300&update_version_code=2022402160&_rticket=1651260864258&current_region=US&app_type=normal&sys_region=US&mcc_mnc=31004&timezone_name=America%2FChicago&residence=US&ts=1651260864&timezone_offset=-21600&build_number=24.2.16&region=US&uoo=0&app_language=en&carrier_region=US&locale=en&op_region=US&ac2=wifi&host_abi=x86&cdid=b65dbfb4-3dbf-436e-9b16-0863ccda5a93&login_name=" + username))))
	request.Header.Set("X-SS-REQ-TICKET", strconv.FormatInt(time.Now().Unix() * 1000, 10))
	request.Header.Set("X-Khronos", strconv.FormatInt(xKhronos, 10))
	
	request.Header.Set("User-Agent", "com.zhiliaoapp.musically/2022402160 (Linux; U; Android 7.1.2; en; G011A; Build/N2G48H;tt-ok/3.10.0.2")
	request.Header.Set("passport-sdk-version", "19")
	request.Header.Set("Cache-Control", "no-cache")
	request.Header.Set("Connection", "keep-alive")
	request.Header.Set("Cookie", session)

	httpClient.Do(request, response)

	return fastjson.GetString(response.Body(), "message") == "success"
}

func usernameID(username string) string {
	var response *fasthttp.Response = fasthttp.AcquireResponse()
	var request *fasthttp.Request = fasthttp.AcquireRequest()
	var xKhronos int64 = time.Now().Unix()

	defer fasthttp.ReleaseResponse(response)
	defer fasthttp.ReleaseRequest(request)

	request.Header.SetMethod("GET")
	request.SetRequestURI("http://api31-normal-q-useast2a.tiktokv.com/aweme/v1/user/uniqueid/?iid=7091502993825253162&device_id=7091503001088591406&ac=wifi&channel=beta&aid=1233&app_name=musical_ly&version_code=240216&version_name=24.2.16&device_platform=android&ab_version=24.2.16&ssmix=a&device_type=G011A&device_brand=google&language=en&os_api=25&os_version=7.1.2&openudid=96a943134babc0db&manifest_version_code=2022402160&resolution=900*1600&dpi=300&update_version_code=2022402160&_rticket=1651260864258&current_region=US&app_type=normal&sys_region=US&mcc_mnc=31004&timezone_name=America%2FChicago&residence=US&ts=1651260864&timezone_offset=-21600&build_number=24.2.16&region=US&uoo=0&app_language=en&carrier_region=US&locale=en&op_region=US&ac2=wifi&host_abi=x86&cdid=b65dbfb4-3dbf-436e-9b16-0863ccda5a93&id=" + username)

	request.Header.Set("X-Gorgon", generateXGorgon(xKhronos, convertBytes(generateSigHash("iid=7091502993825253162&device_id=7091503001088591406&ac=wifi&channel=beta&aid=1233&app_name=musical_ly&version_code=240216&version_name=24.2.16&device_platform=android&ab_version=24.2.16&ssmix=a&device_type=G011A&device_brand=google&language=en&os_api=25&os_version=7.1.2&openudid=96a943134babc0db&manifest_version_code=2022402160&resolution=900*1600&dpi=300&update_version_code=2022402160&_rticket=1651260864258&current_region=US&app_type=normal&sys_region=US&mcc_mnc=31004&timezone_name=America%2FChicago&residence=US&ts=1651260864&timezone_offset=-21600&build_number=24.2.16&region=US&uoo=0&app_language=en&carrier_region=US&locale=en&op_region=US&ac2=wifi&host_abi=x86&cdid=b65dbfb4-3dbf-436e-9b16-0863ccda5a93&id=" + username))))
	request.Header.Set("X-SS-REQ-TICKET", strconv.FormatInt(time.Now().Unix() * 1000, 10))
	request.Header.Set("X-Khronos", strconv.FormatInt(xKhronos, 10))
	
	request.Header.Set("User-Agent", "com.zhiliaoapp.musically/2022402160 (Linux; U; Android 7.1.2; en; G011A; Build/N2G48H;tt-ok/3.10.0.2)")
	request.Header.Set("Cache-Control", "no-cache")
	request.Header.Set("Connection", "keep-alive")
	request.Header.Set("Cookie", generateSession())

	httpClient.Do(request, response)

	return fastjson.GetString(response.Body(), "uid")
}

func discordWebhook(webhook, data string) {
	var request *fasthttp.Request = fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(request)

	request.Header.SetMethod("POST")
	request.SetRequestURI(webhook)

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("USer-Agent", "Fatal/TikTik")
	request.SetBody([]byte(data))

	httpClient.Do(request, nil)
}