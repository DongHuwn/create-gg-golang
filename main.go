package main

import (
	"fmt"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/stealth"
	"strings"
	"time"
)

const REGISTER_GOOGLE_URL = "https://goo.gl/aqsW6D"
const FIRST_NAME_INPUT = "#firstName"
const LAST_NAME_INPUT = "#lastName"
const USER_NAME_INPUT = "#username"
const PASSWD_INPUT = "#passwd > div.aCsJod.oJeWuf > div > div.Xb9hP > input"
const CONFIRM_PASSWD_INPUT = "#confirm-passwd > div.aCsJod.oJeWuf > div > div.Xb9hP > input"

const RECOVERY_EMAIL_INPUT = "input[name='recoveryEmail']"
const DAY_INPUT = "input[name='day']"
const MONTH_INPUT = "select#month"
const YEAR_INPUT = "input[name='year']"
const GENDER_INPUT = "select#gender"
const BUTTON_NEXT_ACCOUNT_GOOGLE = "#accountDetailsNext > div > button"
const IS_LOGIN_GMAIL = "a[href^='https://accounts.google.com/SignOutOptions']"

func init() {
	launcher.NewBrowser().MustGet()
}

func main() {
	// test voi create account google
	wsURL := launcher.NewUserMode().MustLaunch()
	browser := rod.New().ControlURL(wsURL).MustConnect().NoDefaultDevice()
	defer browser.MustClose()
	//browser.EachEvent(func(e *proto.PageWindowOpen) {
	//
	//})

	page := stealth.MustPage(browser)
	//ken&redirect_uri=https%3A%2F%2Fwww.tiktok.com%2Flogin%2F&state={\"client_id\"%3A\"1096011445005-sdea0nf5jvj14eia93icpttv27cidkvk.apps.googleusercontent.com\"%2C\"network\"%3A\"google\"%2C\"display\"%3A\"popup\"%2C\"callback\"%3A\"_hellojs_9c2diz1s\"%2C\"state\"%3A\"\"%2C\"redirect_uri\"%3A\"https%3A%2F%2Fwww.tiktok.com%2Flogin%2F\"%2C\"scope\"%3A\"basic\"}&scope=openid profile&prompt=consent&flowName=GeneralOAuthFlow")
	//page := browser.MustPage(REGISTER_GOOGLE_URL)
	//page.MustEmulate(devices.IPhoneX)
	page.MustNavigate("https://www.tiktok.com/en/")
	wait := page.MustWaitNavigation()
	wait()
	page.MustWaitLoad()
	time.Sleep(time.Second * 5)
	page.Timeout(15 * time.Second).MustElement("#app > div.tiktok-12azhi0-DivHeaderContainer.e10win0d0 > div > div.tiktok-ba55d9-DivHeaderRightContainer.e13wiwn60 > button").CancelTimeout()
	page.MustElement("#app > div.tiktok-12azhi0-DivHeaderContainer.e10win0d0 > div > div.tiktok-ba55d9-DivHeaderRightContainer.e13wiwn60 > button").MustClick()

	page.MustWaitLoad()
	time.Sleep(time.Second * 5)
	page.Timeout(15 * time.Second).MustElement("#loginContainer > div > div > div:nth-child(5)").CancelTimeout()
	page.MustElement("#loginContainer > div > div > div:nth-child(5)").MustClick().MustWaitInvisible()
	newPage := page.MustWaitOpen()
	newTab := newPage()
	newTab.EvalOnNewDocument(stealth.JS)

	//page.MustWaitLoad()
	time.Sleep(time.Second * 1000000)

	//ele, error := page.Element("#loginContainer > div > div > div:nth-child(5)")
	//if error != nil {
	//	fmt.Println("Vao day")
	//	ele.MustClick().MustWaitInvisible()
	//	page.MustWaitLoad()
	//	time.Sleep(time.Second * 1000000)
	//}
	time.Sleep(time.Second * 1000000)
	//page.MustElement(FIRST_NAME_INPUT).MustInput("Bùi Doãn")
	//time.Sleep(time.Millisecond * 5)
	//page.MustElement(LAST_NAME_INPUT).MustInput("Hà")
	//time.Sleep(time.Millisecond * 5)
	//page.MustElement(USER_NAME_INPUT).MustInput("buidoanha24535612")
	//time.Sleep(time.Millisecond * 5)
	//page.MustElement(PASSWD_INPUT).MustInput("Buidoanha@1")
	//time.Sleep(time.Millisecond * 5)
	//page.MustElement(CONFIRM_PASSWD_INPUT).MustInput("Buidoanha@1")
	//page.MustElement(BUTTON_NEXT_ACCOUNT_GOOGLE).MustClick().MustWaitInvisible()
	//time.Sleep(time.Millisecond * 10)
	//page.MustWaitLoad()
	//page.Timeout(15 * time.Second).MustElement("img[src^='https://ssl.gstatic.com/accounts/signup/glif/account.svg']").CancelTimeout()
	//page.MustElement(RECOVERY_EMAIL_INPUT).MustInput("miaooo.sl.ys.lyn.hyeon.lb@gmail.com")
	//time.Sleep(time.Millisecond * 5)
	//page.MustElement(DAY_INPUT).MustInput("5")
	//time.Sleep(time.Millisecond * 5)
	//page.MustElement(MONTH_INPUT).Select([]string{`[value="2"]`}, false, rod.SelectorTypeCSSSector)
	//time.Sleep(time.Millisecond * 5)
	//page.MustElement(YEAR_INPUT).MustInput("1999")
	//time.Sleep(time.Millisecond * 5)
	//page.MustElement(GENDER_INPUT).Select([]string{`[value="1"]`}, false, rod.SelectorTypeCSSSector)
	//time.Sleep(time.Millisecond * 5)
	//page.MustElement(BUTTON_NEXT_ACCOUNT_GOOGLE).MustClick().MustWaitInvisible()
	//time.Sleep(time.Millisecond * 10)
	//page.MustWaitLoad()
	//page.Timeout(15 * time.Second).MustElement("img[src^='https://ssl.gstatic.com/accounts/signup/glif/privacy.svg']").CancelTimeout()
	//page.MustElement(BUTTON_NEXT_ACCOUNT_GOOGLE).MustScrollIntoView().MustClick().MustWaitInvisible()
	//page.MustWaitLoad()
	el, error := page.Element(IS_LOGIN_GMAIL)
	if error != nil {
		// print the username after successful login
		content := el.MustAttribute("aria-label")
		contentValue := *content
		if strings.Contains(contentValue, "buidoanhai2451242") {
			fmt.Println("Register success")
		} else {
			fmt.Println("Register failed")
		}
	}
	//if page.MustHas("img[src^='https://ssl.gstatic.com/accounts/signup/glif/account.svg']") {
	//	page.MustElementR(FIRST_NAME_INPUT, "First name")
	//	page.MustElement(FIRST_NAME_INPUT).MustInput("Bùi Doãn")
	//	page.MustElement(LAST_NAME_INPUT).MustInput("Hải")
	//	page.MustElement(USER_NAME_INPUT).MustInput("buidoanhai2451242")
	//	page.MustElement(PASSWD_INPUT).MustInput("Buidoanhai@1")
	//	page.MustElement(CONFIRM_PASSWD_INPUT).MustInput("Buidoanhai@1")
	//	page.MustElement(BUTTON_NEXT_ACCOUNT_GOOGLE).MustClick().MustWaitInvisible()
	//	time.Sleep(time.Millisecond * 10)
	//	page.MustWaitLoad()
	//}
	//if page.MustHas("img[src^='https://ssl.gstatic.com/accounts/signup/glif/personal.svg']") {
	//	page.MustElementR(RECOVERY_EMAIL_INPUT, "Recovery email")
	//	page.MustElement(RECOVERY_EMAIL_INPUT).MustInput("miaooo.sl.ys.lyn.hyeon.lb@gmail.com")
	//	page.MustElement(DAY_INPUT).MustInput("5")
	//	page.MustElement(MONTH_INPUT).Select([]string{`[value="2"]`}, false, rod.SelectorTypeCSSSector)
	//	page.MustElement(YEAR_INPUT).MustInput("1999")
	//	page.MustElement(GENDER_INPUT).Select([]string{`[value="1"]`}, false, rod.SelectorTypeCSSSector)
	//	page.MustElementR(BUTTON_NEXT_ACCOUNT_GOOGLE, "/^(Next|I agree)$/").MustClick().MustWaitInvisible()
	//	time.Sleep(time.Millisecond * 10)
	//	page.MustWaitLoad()
	//}
	//if page.MustHas("img[src^='https://ssl.gstatic.com/accounts/signup/glif/privacy.svg']") {
	//	page.MustElementR(BUTTON_NEXT_ACCOUNT_GOOGLE, "/^(Next|I agree)$/").MustClick().MustWaitInvisible()
	//	time.Sleep(time.Millisecond * 20)
	//	page.MustWaitLoad()
	//}
	//el, _ := page.Element(IS_LOGIN_GMAIL)
	//if page.MustHas(IS_LOGIN_GMAIL) {
	//	// print the username after successful login
	//	content := el.MustAttribute("aria-label")
	//	contentValue := *content
	//	if strings.Contains(contentValue, "buidoanhai2451242") {
	//		fmt.Println("Register success")
	//	} else {
	//		fmt.Println("Register failed")
	//	}
	//}
}

//func printReport(page *rod.Page) {
//	el := page.MustElement("#broken-image-dimensions.passed")
//	for _, row := range el.MustParents("table").First().MustElements("tr:nth-child(n+2)") {
//		cells := row.MustElements("td")
//		key := cells[0].MustProperty("textContent")
//		if strings.HasPrefix(key.String(), "User Agent") {
//			fmt.Printf("\t\t%s: %t\n\n", key, !strings.Contains(cells[1].MustProperty("textContent").String(), "HeadlessChrome/"))
//		} else if strings.HasPrefix(key.String(), "Hairline Feature") {
//			// Detects support for hidpi/retina hairlines, which are CSS borders with less than 1px in width, for being physically 1px on hidpi screens.
//			// Not all the machine suppports it.
//			continue
//		} else {
//			fmt.Printf("\t\t%s: %s\n\n", key, cells[1].MustProperty("textContent"))
//		}
//	}
//
//	page.MustScreenshot("")
//}
