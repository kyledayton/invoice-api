package pdf

import (
	"context"
	"errors"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

type RemoteChromeDevToolsRenderer struct {
	devTools *RemoteChromeDevToolsContext
	Options  ChromeDevToolsPDFRenderOptions
}

func NewRemoteChromeDevToolsRenderer(ctx *RemoteChromeDevToolsContext, options ChromeDevToolsPDFRenderOptions) *RemoteChromeDevToolsRenderer {
	return &RemoteChromeDevToolsRenderer{
		devTools: ctx,
		Options:  options,
	}
}

func (r *RemoteChromeDevToolsRenderer) RenderPDF(htmlContent []byte) (res []byte, err error) {
	prepareContent := chromedp.ActionFunc(func(c context.Context) error {
		treeParams, err := page.GetFrameTree().Do(c)
		if err != nil {
			return err
		}

		frame := treeParams.Frame
		if frame == nil {
			return errors.New("frame is not defined")
		}

		frameId := frame.ID

		setDocParams := page.SetDocumentContent(frameId, string(htmlContent))
		err = setDocParams.Do(c)
		if err == nil {
			return err
		}

		return nil
	})

	printPdf := chromedp.ActionFunc(func(c context.Context) error {
		pdf := page.PrintToPDF()
		pdf = pdf.WithDisplayHeaderFooter(r.Options.DisplayHeaderFooter)
		pdf = pdf.WithFooterTemplate(r.Options.FooterTemplate)
		pdf = pdf.WithGenerateDocumentOutline(r.Options.GenerateDocumentOutline)
		pdf = pdf.WithGenerateTaggedPDF(r.Options.GenerateTaggedPDF)
		pdf = pdf.WithHeaderTemplate(r.Options.HeaderTemplate)
		pdf = pdf.WithLandscape(r.Options.Landscape)
		pdf = pdf.WithMarginTop(r.Options.Margins.Top)
		pdf = pdf.WithMarginBottom(r.Options.Margins.Bottom)
		pdf = pdf.WithMarginLeft(r.Options.Margins.Left)
		pdf = pdf.WithMarginRight(r.Options.Margins.Right)
		pdf = pdf.WithPageRanges(r.Options.PageRanges)
		pdf = pdf.WithPaperWidth(r.Options.PaperSize.Width)
		pdf = pdf.WithPaperHeight(r.Options.PaperSize.Height)
		pdf = pdf.WithPreferCSSPageSize(r.Options.PreferCSSPageSize)
		pdf = pdf.WithPrintBackground(r.Options.PrintBackground)
		pdf = pdf.WithScale(r.Options.Scale)

		buf, _, err := pdf.Do(c)
		if err != nil {
			return err
		}

		res = buf
		return nil
	})

	r.devTools.WithContext(func(ctx context.Context) {
		err = chromedp.Run(ctx,
			chromedp.Navigate("about:blank"),
			prepareContent,
			printPdf,
		)
		if err != nil {
			return
		}
	})

	return
}

type ChromeDevToolsPDFRenderOptions struct {
	DisplayHeaderFooter     bool
	FooterTemplate          string
	GenerateDocumentOutline bool
	GenerateTaggedPDF       bool
	HeaderTemplate          string
	Landscape               bool
	Margins                 Margins
	PageRanges              string
	PaperSize               Size
	PreferCSSPageSize       bool
	PrintBackground         bool
	Scale                   float64
}

func DefaultChromeDevToolsPDFRenderOptions() ChromeDevToolsPDFRenderOptions {
	return ChromeDevToolsPDFRenderOptions{
		DisplayHeaderFooter:     false,
		FooterTemplate:          "",
		GenerateDocumentOutline: false,
		GenerateTaggedPDF:       false,
		HeaderTemplate:          "",
		Landscape:               false,
		Margins:                 UniformMargins(0.4),
		PageRanges:              "",
		PaperSize:               Size{8.5, 11},
		PreferCSSPageSize:       false,
		PrintBackground:         false,
		Scale:                   1,
	}
}

type Margins struct {
	Top    float64
	Bottom float64
	Left   float64
	Right  float64
}

func UniformMargins(margin float64) Margins {
	return Margins{margin, margin, margin, margin}
}

type Size struct {
	Width  float64
	Height float64
}
