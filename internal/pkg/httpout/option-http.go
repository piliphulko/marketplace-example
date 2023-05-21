package httpout

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"html/template"
	"io"
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
	HeaderResponseMap     map[string]string
	HeaderRequestMap      map[string]string
	HTML                  *template.Template
	OkRedirectUse         bool
	OkRedirectUseDataURL  bool
	OkRedirectPath        string
	ErrRedirectUse        bool
	ErrRedirectUseDataURL bool
	ErrRedirectPath       string
	PossibleErrorsClient  []error
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
				LogHTTP.Error(fmt.Sprintf("context dedline: %d", timeCtx))
				w.WriteHeader(http.StatusRequestTimeout)
				return
			} else {
				if optHTTP.ErrRedirectUseDataURL {
					errBackend, errFrontend, err := optHTTP.TakeBackendFrontendError(context.Cause(ctx))
					if err == ErrNoClientError {
						errFrontend = ErrSpiderMan
					} else if err == ErrReportedErrorNotList {
						LogHTTP.Info(ErrReportedErrorNotList.Error())
					}
					LogHTTP.Error(errBackend.Error())

					buf := bytes.Buffer{}
					if err := JSON.NewEncoder(&buf).Encode(RedirectAnswer{
						Ok:      true,
						ErrInfo: errFrontend.Error(),
					}); err != nil {
						LogHTTP.Error(err.Error())
						goto errorFinish
					}
					//RedirectPath := string(<-ch)
					//http.Redirect(w, r, RedirectPath+"?data="+url.QueryEscape(buf.String()), http.StatusMovedPermanently)
					http.Redirect(w, r, optHTTP.ErrRedirectPath+"?data="+url.QueryEscape(buf.String()), http.StatusMovedPermanently)
				} else if optHTTP.ErrRedirectUse {
					//RedirectPath := string(<-ch)
					//http.Redirect(w, r, RedirectPath, http.StatusMovedPermanently)
					http.Redirect(w, r, optHTTP.ErrRedirectPath, http.StatusMovedPermanently)
				}
			errorFinish:
				LogHTTP.Error(context.Cause(ctx).Error())
				w.WriteHeader(http.StatusInternalServerError)
			}
			return
		case b := <-ch:
			if optHTTP.OkRedirectUseDataURL {
				//RedirectPath := string(<-ch)
				//http.Redirect(w, r, RedirectPath+"?data="+url.QueryEscape(string(b)), http.StatusMovedPermanently)
				http.Redirect(w, r, optHTTP.OkRedirectPath+"?data="+url.QueryEscape(string(b)), http.StatusMovedPermanently)
				return
			} else if optHTTP.OkRedirectUse {
				//RedirectPath := string(<-ch)
				//http.Redirect(w, r, RedirectPath, http.StatusMovedPermanently)
				return
			}
			optHTTP.HeaderResponseSet(w)
			w.Write(b)
			return
		}
	}
}

func (optHTTP *OptionsHTTP) UseOkRedirect(pathRedirect string) *OptionsHTTP {
	optHTTP.OkRedirectUse = true
	optHTTP.OkRedirectPath = pathRedirect
	return optHTTP
}

func (optHTTP *OptionsHTTP) UseErrRedirect(pathRedirect string) *OptionsHTTP {
	optHTTP.ErrRedirectUse = true
	optHTTP.ErrRedirectPath = pathRedirect
	return optHTTP
}

func (optHTTP *OptionsHTTP) UseOkRedirectDataURL(pathRedirect string) *OptionsHTTP {
	optHTTP.OkRedirectUseDataURL = true
	optHTTP.OkRedirectPath = pathRedirect
	return optHTTP
}

func (optHTTP *OptionsHTTP) UseErrRedirectDataURL(pathRedirect string) *OptionsHTTP {
	optHTTP.ErrRedirectUseDataURL = true
	optHTTP.ErrRedirectPath = pathRedirect
	return optHTTP
}

type RedirectAnswer struct {
	Ok      bool
	OkInfo  string
	ErrInfo string
}

func withTimeoutSecond(t int) time.Duration {
	return time.Duration(t) * time.Second
}

func (optHTTP *OptionsHTTP) ReceptionRedirectURL() *OptionsHTTP { return optHTTP }

func ErrorIntoClient(err error, clientErr error) error {
	return fmt.Errorf("%v %w", clientErr, err)
}

func (optHTTP *OptionsHTTP) SetErrorClientList(errArray ...error) *OptionsHTTP {
	optHTTP.PossibleErrorsClient = errArray
	return optHTTP
}

func (optHTTP OptionsHTTP) TakeBackendFrontendError(err error) (error, error, error) {
	errBackend := errors.Unwrap(err)
	if errBackend == nil {
		return err, nil, ErrNoClientError
	}
	for _, errPossible := range optHTTP.PossibleErrorsClient {
		if errors.Is(err, errPossible) {
			return errBackend, errPossible, nil
		}
	}
	return err, fmt.Errorf(strings.Replace(err.Error(), errBackend.Error(), "", 1)), ErrReportedErrorNotList
}

func TakeRedirectAnswerFromURL(r *http.Request, redirectAnswer *RedirectAnswer) error {
	params := r.URL.Query()
	data := params.Get("data")
	if err := JSON.NewDecoder(strings.NewReader(data)).Decode(&redirectAnswer); err != nil {
		if err == io.EOF {
		} else {
			return err
		}
	}
	return nil
}
