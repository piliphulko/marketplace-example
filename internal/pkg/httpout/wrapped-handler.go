package httpout

import (
	"context"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
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
	HeaderResponseWriterMap map[string]string
	HTML                    *template.Template
}

func StartOptionsHTTP() *OptionsHTTP { return &OptionsHTTP{} }

func (optHTTP *OptionsHTTP) HeaderHTTPResponse(header map[string]string) *OptionsHTTP {
	optHTTP.HeaderResponseWriterMap = header
	return optHTTP
}

func (optHTTP OptionsHTTP) HeaderSet(w http.ResponseWriter) {
	for k, v := range optHTTP.HeaderResponseWriterMap {
		w.Header().Set(k, v)
	}
}

func (optHTTP *OptionsHTTP) WithHTML(temp *template.Template, nameHTML string) *OptionsHTTP {
	html := temp.Lookup(nameHTML)
	if html == nil {
		LogHTTP.Panic("")
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
				LogHTTP.Error(fmt.Sprintf("func: handlerTest | context dedline: %d sec", timeCtx))
				w.WriteHeader(http.StatusRequestTimeout)
			} else {
				LogHTTP.Error(context.Cause(ctx).Error())
				w.WriteHeader(http.StatusInternalServerError)
			}
			return
		case b := <-ch:
			optHTTP.HeaderSet(w)
			w.Write(b)
			return
		}
	}
}
