// Command screenshot is a chromedp example demonstrating how to take a
// screenshot of a specific element and of the entire browser viewport.
package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strings"
	"time"

	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func RunWithTimeOut(ctx *context.Context, timeout time.Duration, tasks chromedp.Tasks) chromedp.ActionFunc {
	return func(ctx context.Context) error {
		timeoutContext, cancel := context.WithTimeout(ctx, timeout*time.Second)
		defer cancel()
		return tasks.Do(timeoutContext)
	}
}

func main() {
	// get flags
	pathPtr := flag.String("path", "http://localhost:9980/", "url")
	elementPtr := flag.String("element", ".main-body", "class to wait for and then capture")
	filePtr := flag.String("file", "screenshot.png", "save as filename")
	typePtr := flag.String("type", "element", "element or full")
	flag.Parse()

	// ignore unsigned certs
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("ignore-certificate-errors", "1"),
	)

	// create context
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()
	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()
	var buf []byte

	if *typePtr == `element` {
		// capture screenshot of an element
		if err := chromedp.Run(ctx, RunWithTimeOut(&ctx, 30, elementScreenshot(*pathPtr, *elementPtr, &buf))); err != nil {
			log.Fatal(err)
		}
		if err := ioutil.WriteFile(*filePtr, buf, 0644); err != nil {
			log.Fatal(err)
		}
	}
	if *typePtr == "full" {
		// capture entire browser viewport, returning png with quality=90
		if err := chromedp.Run(ctx, RunWithTimeOut(&ctx, 30, fullScreenshot(*pathPtr, *elementPtr, 90, &buf))); err != nil {
			log.Fatal(err)
		}
		if err := ioutil.WriteFile(*filePtr, buf, 0644); err != nil {
			log.Fatal(err)
		}
	}
}

// elementScreenshot takes a screenshot of a specific element.
func elementScreenshot(urlstr string, sel string, res *[]byte) chromedp.Tasks {
	if strings.Contains(urlstr, "https") {
		return chromedp.Tasks{
			chromedp.Navigate(urlstr),
			chromedp.WaitVisible(`input[type='text']`, chromedp.ByQuery),
			chromedp.SendKeys(`input[type='text']`, "admin", chromedp.ByQuery),
			chromedp.WaitVisible(`input[type='password']`, chromedp.ByQuery),
			chromedp.SendKeys(`input[type='password']`, "adminpw\t\t\n", chromedp.ByQuery),
			chromedp.WaitVisible(sel, chromedp.ByQuery),
			chromedp.Sleep(2 * time.Second),
			chromedp.Screenshot(sel, res, chromedp.NodeVisible, chromedp.ByQuery),
		}
	} else {
		return chromedp.Tasks{
			chromedp.Navigate(urlstr),
			chromedp.WaitVisible(sel, chromedp.ByQuery),
			chromedp.Sleep(2 * time.Second),
			chromedp.Screenshot(sel, res, chromedp.NodeVisible, chromedp.ByQuery),
		}
	}
}

// fullScreenshot takes a screenshot of the entire browser viewport.
//
// Liberally copied from puppeteer's source.
//
// Note: this will override the viewport emulation settings.
func fullScreenshot(urlstr string, sel string, quality int64, res *[]byte) chromedp.Tasks {

	if strings.Contains(urlstr, "https") {
		return chromedp.Tasks{
			chromedp.Navigate(urlstr),
			chromedp.WaitVisible(`input[type='text']`, chromedp.ByQuery),
			chromedp.SendKeys(`input[type='text']`, "admin", chromedp.ByQuery),
			chromedp.WaitVisible(`input[type='password']`, chromedp.ByQuery),
			chromedp.SendKeys(`input[type='password']`, "adminpw\t\t\n", chromedp.ByQuery),
			chromedp.WaitVisible(sel, chromedp.ByQuery),
			chromedp.Sleep(2 * time.Second),
			chromedp.ActionFunc(func(ctx context.Context) error {
				// get layout metrics
				_, _, contentSize, err := page.GetLayoutMetrics().Do(ctx)
				if err != nil {
					return err
				}
				width, height := int64(math.Ceil(contentSize.Width)), int64(math.Ceil(contentSize.Height))

				// force viewport emulation
				err = emulation.SetDeviceMetricsOverride(width, height, 1, false).
					WithScreenOrientation(&emulation.ScreenOrientation{
						Type:  emulation.OrientationTypePortraitPrimary,
						Angle: 0,
					}).
					Do(ctx)
				if err != nil {
					return err
				}

				// capture screenshot
				*res, err = page.CaptureScreenshot().
					WithQuality(quality).
					WithClip(&page.Viewport{
						X:      contentSize.X,
						Y:      contentSize.Y,
						Width:  contentSize.Width,
						Height: contentSize.Height,
						Scale:  1,
					}).Do(ctx)
				if err != nil {
					return err
				}
				return nil
			}),
		}
	} else {
		return chromedp.Tasks{
			chromedp.Navigate(urlstr),
			chromedp.WaitVisible(sel, chromedp.ByQuery),
			chromedp.Sleep(2 * time.Second),
			chromedp.ActionFunc(func(ctx context.Context) error {
				// get layout metrics
				_, _, contentSize, err := page.GetLayoutMetrics().Do(ctx)
				if err != nil {
					return err
				}
				width, height := int64(math.Ceil(contentSize.Width)), int64(math.Ceil(contentSize.Height))

				// force viewport emulation
				err = emulation.SetDeviceMetricsOverride(width, height, 1, false).
					WithScreenOrientation(&emulation.ScreenOrientation{
						Type:  emulation.OrientationTypePortraitPrimary,
						Angle: 0,
					}).
					Do(ctx)
				if err != nil {
					return err
				}

				// capture screenshot
				*res, err = page.CaptureScreenshot().
					WithQuality(quality).
					WithClip(&page.Viewport{
						X:      contentSize.X,
						Y:      contentSize.Y,
						Width:  contentSize.Width,
						Height: contentSize.Height,
						Scale:  1,
					}).Do(ctx)
				if err != nil {
					return err
				}
				return nil
			}),
		}
	}

}

func getValues() string {
	// Social struct which contains a
	// list of links
	type AnsibleOutput struct {
		Bp_integration_tests_token string `json:"bp_integration_tests_token"`
	}
	// Open our jsonFile
	jsonFile, err := os.Open("localhost")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened localhost")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result AnsibleOutput
	// var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)
	fmt.Println(result.Bp_integration_tests_token)
	return result.Bp_integration_tests_token
}
