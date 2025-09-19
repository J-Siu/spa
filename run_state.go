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
	"strconv"

	"github.com/J-Siu/basetype"
	"github.com/go-rod/rod"
)

// SPA Processor Run State
type RunState struct {
	*basetype.Base

	Elements          *rod.Elements
	ElementLast       *rod.Element
	ElementLastScroll *rod.Element
	ElementCountLast  int
	ListItemLast      IInfo
	Scroll            bool
	ScrollCount       int
}

func (s *RunState) New() *RunState {
	s.Base = new(basetype.Base)
	s.MyType = "RunState"
	s.Initialized = true

	s.Elements = nil
	s.ElementLast = nil
	s.ElementLastScroll = nil
	s.ElementCountLast = 0
	s.ListItemLast = nil
	s.Scroll = true
	s.ScrollCount = 0

	return s
}

// This should only be used at Trace level log
func (s *RunState) String() *string {
	var str string
	if s.Elements != nil {
		str += "esCount:" + strconv.Itoa(len(*s.Elements)) + "\n"
	}
	if s.ElementLast != nil {
		str += "eLast:" + string(s.ElementLast.Object.ObjectID) + "\n"
	}
	if s.ElementLastScroll != nil {
		str += "eLastScroll:" + string(s.ElementLastScroll.Object.ObjectID) + "\n"
	}
	if s.ListItemLast != nil {
		str += "ListItemLast(matched):" + strconv.FormatBool(s.ListItemLast.GetMatched()) + "\n"
	}
	str += "eCountLast:" + strconv.Itoa(s.ElementCountLast) + "\n"
	str += "Scroll:" + strconv.FormatBool(s.Scroll)
	return &str
}
