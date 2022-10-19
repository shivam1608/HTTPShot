package handler

import (
	"context"
	"log"
	"strconv"

	"github.com/chromedp/chromedp"
)

// Screenshot : Captures Screenshot of the given URL
func Screenshot(url string, buffer *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitReady("body"),
		chromedp.CaptureScreenshot(buffer),
	}
}

// FullScreenshot : Captures Full Screenshot of the given URL
func FullScreenshot(url string, buffer *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitReady("body"),
		chromedp.FullScreenshot(buffer, 100),
	}
}

// CaptureHTTPShot : Captures HTTPShot of the given URL
func CaptureHTTPShot(url string, width string, height string, fullscreen string) []byte {
	w, _ := strconv.ParseInt(width, 10, 64)
	h, _ := strconv.ParseInt(height, 10, 64)
	fs, _ := strconv.ParseBool(fullscreen)
	log.Println("URL: " + url)
	var buffer []byte
	ctx, cancel := chromedp.NewContext(
		context.Background(),
	)
	defer cancel()

	if fs {
		if err := chromedp.Run(ctx,
			chromedp.EmulateViewport(w, h),
			FullScreenshot(url, &buffer),
		); err != nil {
			log.Println(err)
		}

	} else {
		if err := chromedp.Run(ctx,
			chromedp.EmulateViewport(w, h),
			Screenshot(url, &buffer),
		); err != nil {
			log.Println(err)
		}
	}

	return buffer
}
