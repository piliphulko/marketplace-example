package httpout

import (
	"context"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
	"time"
)

func FillTempHTMLfromDir(pathDir string) {
	err := filepath.WalkDir(pathDir, func(path string, d fs.DirEntry, err error) error {
		if strings.Contains(path, ".html") {
			_, err := TempHTML.ParseFiles(path)
			if err != nil {
				LogHTTP.Panic(err.Error())
				log.Panic(err)
			}
		}
		return err
	})
	if err != nil {
		LogHTTP.Panic(err.Error())
		log.Panic(err)
	}
}

type OptionsHTTP struct {
	HeaderResponseMap map[string]string
	HeaderRequestMap  map[string]string
	HTML              *template.Template
	Redirect          bool
	RedirectOK        string
	RedirectERR       string
}

func StartOptionsHTTP() *OptionsHTTP { return &OptionsHTTP{} }

func (optHTTP *OptionsHTTP) HeaderHTTPResponse(header map[string]string) *OptionsHTTP {
	optHTTP.HeaderResponseMap = header
	return optHTTP
}

func (optHTTP *OptionsHTTP) HeaderHTTPRequest(header map[string]string) *OptionsHTTP {
	optHTTP.HeaderRequestMap = header
	return optHTTP
}

func (optHTTP OptionsHTTP) HeaderResponseSet(w http.ResponseWriter) {
	for k, v := range optHTTP.HeaderResponseMap {
		w.Header().Set(k, v)
	}
}

func (optHTTP OptionsHTTP) HeaderRequestAdd(r *http.Request) {
	for k, v := range optHTTP.HeaderRequestMap {
		r.Header.Add(k, v)
	}
}

func (optHTTP *OptionsHTTP) WithHTML(temp *template.Template, nameHTML string) *OptionsHTTP {
	html := temp.Lookup(nameHTML)
	if html == nil {
		LogHTTP.Panic("WithHTML")
		log.Panic()
	}
	optHTTP.HTML = html
	return optHTTP
}

type handlerLogics func(context.Context, context.CancelCauseFunc, *OptionsHTTP, *http.Request, chan []byte)

func (optHTTP *OptionsHTTP) handlerRun(ctx context.Context, timeCtx time.Duration, logics handlerLogics) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancelCtx := context.WithTimeout(ctx, timeCtx)
		defer cancelCtx()
		ctx, cancelCtxError := context.WithCancelCause(ctx)
		ch := make(chan []byte, 1)

		go logics(ctx, cancelCtxError, optHTTP, r, ch)

		select {
		case <-ctx.Done():
			if ctx.Err() == context.DeadlineExceeded {
				LogHTTP.Error(fmt.Sprintf("context dedline: %d sec", timeCtx))
				w.WriteHeader(http.StatusRequestTimeout)
			} else {
				LogHTTP.Error(context.Cause(ctx).Error())
				w.WriteHeader(http.StatusInternalServerError)
			}
			return
		case b := <-ch:
			if optHTTP.Redirect {
				optHTTP.RedirectOK = optHTTP.RedirectOK + "?data=" + url.QueryEscape(string(b))
				http.Redirect(w, r, optHTTP.RedirectOK, 301)
				return
			}
			optHTTP.HeaderResponseSet(w)
			w.Write(b)
			return
		}
	}
}

func (optHTTP *OptionsHTTP) SetPathRedirectOK(redirectOK string) *OptionsHTTP {
	optHTTP.RedirectOK = redirectOK
	return optHTTP
}

func (optHTTP *OptionsHTTP) SetPathRedirectERR(redirectERR string) *OptionsHTTP {
	optHTTP.RedirectOK = redirectERR
	return optHTTP
}

func (optHTTP *OptionsHTTP) ReceptionRedirectOK() *OptionsHTTP { return optHTTP }

type RedirectAnswer struct {
	Ok      bool
	OkInfo  string
	ErrInfo string
}

func withTimeoutSecond(t int) time.Duration {
	return time.Duration(t) * time.Second
}

func (optHTTP *OptionsHTTP) RedirectUse() *OptionsHTTP {
	optHTTP.Redirect = true
	return optHTTP
}
