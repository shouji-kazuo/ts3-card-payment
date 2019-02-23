package ts3

import (
	"fmt"
	"os"

	"github.com/sclevine/agouti"
)

type Config struct {
	URL      string
	Username string
	Password string
}

var (
	driver *agouti.WebDriver
)

func init() {
	driver = agouti.ChromeDriver(
		agouti.ChromeOptions("args", []string{
			"--headless",
		}),
		agouti.Debug,
	)
	if err := driver.Start(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		if err2 := driver.Stop(); err2 != nil {
			fmt.Fprintln(os.Stderr, err2)
		}
		os.Exit(1)
	}
}

// Navigate is HOGE.
// ログイン→ https://my.ts3card.com/webapp/login/login.jsp
// username→ <input type="text" name="vo.loginId" maxlength="30" value="" id="login" style="ime-mode: disabled;" class="login-input">
// password→ <input type="password" name="vo.password" maxlength="30" value="" id="pass" class="pass-input">
func Navigate(config *Config) error {
	page, err := driver.NewPage()
	if err != nil {
		return err
	}
	if err = page.Navigate(config.URL); err != nil {
		return err
	}
	getSource, err := page.HTML()
	fmt.Println(getSource)
	return nil
}

// Close is 最後にかならず呼んでね.
func Close() error {
	return driver.Stop()
}
