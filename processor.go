/*
Copyright © 2025 John, Sing Dao, Siu <john.sd.siu@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package spa

import (
	"errors"
	"strconv"

	"github.com/J-Siu/basetype"
	"github.com/J-Siu/ezlog"
	"github.com/go-rod/rod"
)

type IProcessor interface {
	Run(loadUrl bool, checkURL bool, scrollMax int) IProcessor
}

// SPA Processor
type Processor struct {
	*basetype.Base

	Container       *rod.Element
	InfoList        []IInfo
	InfoListEnabled bool
	PageP           *rod.Page
	ScrollOffset    int // TODO: maybe not required
	ScrollSleep     int // TODO: maybe not required
	Url             string

	// IProcess                   func(loadUrl bool, checkURL bool, scrollMax int)
	V000_LoadPage              func(urlStr string, loadUrl bool, checkUrl bool)
	V010_ElementsContainer     func() *rod.Element
	V020_Elements              func(element *rod.Element) *rod.Elements
	V030_ElementInfo           func(element *rod.Element, index int) (infoP IInfo)
	V040_ElementMatch          func(element *rod.Element, index int, infoP IInfo) (matched bool, matchedStr string)
	V050_ElementProcessMatched func(element *rod.Element, index int, infoP IInfo)
	V060_ElementProcessUnmatch func(element *rod.Element, index int, infoP IInfo)
	V070_ElementProcess        func(element *rod.Element, index int, infoP IInfo)
	V080_ElementScrollable     func(element *rod.Element, index int, infoP IInfo) bool
	V090_ElementLoopEnd        func(element *rod.Element, index int, infoP IInfo)
	V100_ScrollLoopEnd         func(state *RunState)
}

func New(pageP *rod.Page, urlStr string, debug bool) *Processor {
	p := new(Processor)
	p.Base = new(basetype.Base)
	p.MyType = "SPA_Processor"

	p.Debug = debug
	p.Url = urlStr
	if pageP == nil {
		p.Err = errors.New("page/tab is nil")
	} else {
		p.PageP = pageP
		p.Initialized = true
	}

	p.initFunc()
	return p
}

func (p *Processor) Reset() *Processor {
	p.MyType = "SPA_Processor"
	p.Initialized = false

	p.Container = nil
	p.InfoList = nil
	p.Url = ""

	return p
}

func (p *Processor) initFunc() {
	p.V000_LoadPage = func(urlStr string, loadUrl bool, checkUrl bool) {
		prefix := p.MyType + ".V000_LoadPage" + "(base)"
		ezlog.Trace(prefix + ":Start")
		if !p.ChkErrInit(prefix) {
			if loadUrl {
				ezlog.Debug(prefix + ": urlStr: " + urlStr)
				p.Err = p.PageP.Navigate(urlStr)
			}
			p.ErrHandler(prefix)
		}
		p.PageP.MustWaitDOMStable()
		ezlog.Trace(prefix + ": End")
	}
	p.V010_ElementsContainer = func() *rod.Element {
		prefix := p.MyType + ".V010_ElementsContainer" + "(base)"
		ezlog.Trace(prefix + ": Done")
		return p.Container
	}
	p.V020_Elements = func(element *rod.Element) *rod.Elements {
		prefix := p.MyType + ".V020_Elements" + "(base)"
		ezlog.Trace(prefix + ": Done")
		return nil
	}
	p.V030_ElementInfo = func(element *rod.Element, index int) (infoP IInfo) {
		prefix := p.MyType + ".V030_ElementInfo" + "(base)"
		ezlog.Trace(prefix + ": Done")
		return nil
	}
	p.V040_ElementMatch = func(element *rod.Element, index int, infoP IInfo) (matched bool, matchedStr string) {
		prefix := p.MyType + ".V040_ElementMatch" + "(base)"
		ezlog.Trace(prefix + ": Done")
		return true, ""
	}
	p.V050_ElementProcessMatched = func(element *rod.Element, index int, infoP IInfo) {
		prefix := p.MyType + ".V050_ElementProcessMatched" + "(base)"
		ezlog.Trace(prefix + ": Done")
	}
	p.V060_ElementProcessUnmatch = func(element *rod.Element, index int, infoP IInfo) {
		prefix := p.MyType + ".V060_ElementProcessUnmatch" + "(base)"
		ezlog.Trace(prefix + ": Done")
	}
	p.V070_ElementProcess = func(element *rod.Element, index int, infoP IInfo) {
		prefix := p.MyType + ".V070_ElementProcess" + "(base)"
		ezlog.Trace(prefix + ": Done")
	}
	p.V080_ElementScrollable = func(element *rod.Element, index int, infoP IInfo) bool {
		prefix := p.MyType + ".V080_ElementScrollable" + "(base)"
		ezlog.Trace(prefix + ": Done")
		return true
	}
	p.V090_ElementLoopEnd = func(element *rod.Element, index int, infoP IInfo) {
		prefix := p.MyType + ".V090_ElementLoopEnd" + "(base)"
		ezlog.Trace(prefix + ": Done")
	}
	p.V100_ScrollLoopEnd = func(state *RunState) {
		prefix := p.MyType + ".V100_ScrollLoopEnd" + "(base)"
		ezlog.Trace(prefix + ": Done")
	}
}

func (p *Processor) Run(loadUrl bool, checkURL bool, scrollMax int) IProcessor {
	// prefix := p.MyType + ".Run"
	state := new(RunState).New()
	p.V000_LoadPage(p.Url, loadUrl, checkURL)
	p.Container = p.V010_ElementsContainer()
	// -- Scroll Loop
	for {
		breakLoop := !(state.Scroll && (state.ScrollCount <= scrollMax || scrollMax < 0))
		if ezlog.LogLevel() == ezlog.TraceLevel {
			var msg string
			msg += p.MyType + ": Run\n"
			msg += "breakLoop == " + "!(state.Scroll && (state.ScrollCount <= scrollMax || scrollMax < 0))" + "\n"
			msg += "breakLoop: " + strconv.FormatBool(breakLoop) + "\n"
			msg += "state.Scroll: " + strconv.FormatBool(state.Scroll)
			msg += "state.ScrollCount" + strconv.Itoa(state.ScrollCount)
			msg += "scrollMax" + strconv.Itoa(scrollMax)
			ezlog.TraceP(&msg)
		}
		if breakLoop {
			break
		}

		elementsCount := 0
		p.ElementScroll(state.ElementLast, p.ScrollOffset, p.ScrollSleep)
		state.ElementLastScroll = state.ElementLast
		state.Elements = p.V020_Elements(p.Container)
		if state.Elements == nil {
			state.Scroll = false
		} else {
			elementsCount = len(*state.Elements)
			ezlog.Debug(p.MyType + ": Run: elementCount: " + strconv.Itoa(elementsCount))
			// -- Element Loop
			for index := state.ElementCountLast; index < elementsCount; index++ {
				element := (*state.Elements)[index]
				info := p.V030_ElementInfo(element, index)
				matched, matchedStr := p.V040_ElementMatch(element, index, info)
				info.SetMatched(matched)
				info.SetMatchedStr(matchedStr)
				if matched {
					p.V050_ElementProcessMatched(element, index, info)
				} else {
					p.V060_ElementProcessUnmatch(element, index, info)
				}
				p.V070_ElementProcess(element, index, info)
				if p.InfoListEnabled {
					p.InfoList = append(p.InfoList, info)
				}
				if p.V080_ElementScrollable(element, index, info) {
					state.ElementLast = element
					state.ListItemLast = info
				}
				p.V090_ElementLoopEnd(element, index, info)
			}
			/*
				Detect end of scroll

				If elements are removed, use V100_ExitScroll to do custom override.
				As both of following checks can be flawed if elements are removed from page DOM.

				"if (elementsCount == elementsCountLast)":
					will be triggered, if number of elements removed
							= number of new elements added after scroll

				"if (ElementLastScroll == ElementLast)":
					will be triggered, if all new elements added after scroll are removed
			*/
			if state.ElementLastScroll != nil && state.ElementLast != nil {
				if state.ElementLastScroll.Object.ObjectID == state.ElementLast.Object.ObjectID {
					state.Scroll = false
				}
			} else if state.ElementLastScroll == state.ElementLast {
				state.Scroll = false
			}
			state.ElementCountLast = elementsCount
		}

		p.V100_ScrollLoopEnd(state)
		state.ScrollCount++

		ezlog.TraceP(state.String()) // Trace level log
	}
	return p
}

func (p *Processor) ElementScroll(element *rod.Element, scrollOffset int, scrollSleep int) IProcessor {
	prefix := p.MyType + ".ElementScroll"
	ezlog.Trace(prefix + ": Start")
	if element != nil {

		element.MustScrollIntoView()
		ezlog.Trace(prefix + ": MustWaitDOMStable: Start")
		p.PageP.MustWaitDOMStable()
		ezlog.Trace(prefix + ": MustWaitDOMStable: End")
	}
	ezlog.Trace(prefix + ": End")
	return p
}
