package ts3

import (
	"time"

	"github.com/sclevine/agouti"
)

type Config struct {
	URL      string
	Username string
	Password string
}

type AmountResult struct {
	PreviousMonthHeader string
	PreviousMonthAmount string
	CurrentMonthHeader  string
	CurrentMonthAmount  string
	NextMonthHeader     string
	NextMonthAmount     string
}

func Navigate(config *Config) (*AmountResult, error) {
	driver := agouti.ChromeDriver(
		agouti.ChromeOptions("args", []string{
			"--headless",
		}),
		agouti.Debug,
	)
	if err := driver.Start(); err != nil {
		return nil, err
	}

	page, err := driver.NewPage()
	if err != nil {
		return nil, err
	}
	if err = page.Navigate(config.URL); err != nil {
		return nil, err
	}
	loginTextInput := page.FindByID("login")
	passwordInput := page.FindByID("pass")
	loginTextInput.Fill(config.Username)
	passwordInput.Fill(config.Password)

	// ログイン
	if err = page.FindByXPath(`//img[@alt="ログイン"]`).Click(); err != nil {
		return nil, err
	}

	// うざいポップアップ (XPath = /html/body/script[4]) を抑止したい
	// 超暫定対応だが，待てばポップアップは消える
	time.Sleep(5 * time.Second)

	// 詳細 をクリック
	if err = page.FindByXPath(`//*[@id="MeisaiDetail"]/div[1]/a`).Click(); err != nil {
		return nil, err
	}
	time.Sleep(1 * time.Second)

	// 次へ をクリック
	if err = page.FindByXPath(`//*[@id="center"]/span/a`).Click(); err != nil {
		return nil, err
	}
	time.Sleep(1 * time.Second)

	// 前回(xx月xx日)ご請求分
	prevMonthHeader, err := page.FindByXPath(`//*[@id="wrap"]/div[2]/div/div/div[2]/div[1]/table/tbody/tr[1]/th`).Text()
	if err != nil {
		return nil, err
	}
	// 前回の金額
	prevMonthAmount, err := page.FindByXPath(`//*[@id="wrap"]/div[2]/div/div/div[2]/div[1]/table/tbody/tr[1]/td[1]`).Text()
	if err != nil {
		return nil, err
	}

	// 今回(xx月xx日)ご請求分
	currentMonthHeader, err := page.FindByXPath(`//*[@id="wrap"]/div[2]/div/div/div[2]/div[1]/table/tbody/tr[2]/th`).Text()
	if err != nil {
		return nil, err
	}
	// 今回の金額
	currentMonthAmount, err := page.FindByXPath(`//*[@id="wrap"]/div[2]/div/div/div[2]/div[1]/table/tbody/tr[2]/td[1]`).Text()
	if err != nil {
		return nil, err
	}
	// 次回(xx月xx日)ご請求分
	nextMonthHeader, err := page.FindByXPath(`//*[@id="wrap"]/div[2]/div/div/div[2]/div[1]/table/tbody/tr[3]/th`).Text()
	if err != nil {
		return nil, err
	}
	// 次回の金額
	nextMonthAmount, err := page.FindByXPath(`//*[@id="wrap"]/div[2]/div/div/div[2]/div[1]/table/tbody/tr[3]/td[1]`).Text()
	if err != nil {
		return nil, err
	}
	return &AmountResult{
		PreviousMonthHeader: prevMonthHeader,
		PreviousMonthAmount: prevMonthAmount,
		CurrentMonthHeader:  currentMonthHeader,
		CurrentMonthAmount:  currentMonthAmount,
		NextMonthHeader:     nextMonthHeader,
		NextMonthAmount:     nextMonthAmount,
	}, nil
}
