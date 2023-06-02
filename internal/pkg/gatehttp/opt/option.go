package opt

import (
	"bytes"
	"context"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"google.golang.org/grpc"
)

type OptionsHTTP struct {
	headerResponseMap            map[string]string
	headerRequestMap             map[string]string
	html                         *template.Template
	RedirectUseDataURL           bool
	RedirectHTTPPathOk           string
	RedirectHTTPPathMistake      string
	connectingToMicroservicesMap map[string]grpc.ClientConnInterface
}

// NewOptionsHTTP creates an object for setting HTTP request parameters
func NewOptionsHTTP() *OptionsHTTP { return &OptionsHTTP{} }

// SetHeaderResponse setting response headers
func (optHTTP *OptionsHTTP) SetHeaderResponse(headers map[string]string) *OptionsHTTP {
	optHTTP.headerResponseMap = headers
	return optHTTP
}

func (optHTTP OptionsHTTP) takeHeaderResponse(w http.ResponseWriter) {
	for k, v := range optHTTP.headerResponseMap {
		w.Header().Set(k, v)
	}
}

// SetHeaderRequest setting request headers
func (optHTTP *OptionsHTTP) SetHeaderRequest(headers map[string]string) *OptionsHTTP {
	optHTTP.headerRequestMap = headers
	return optHTTP
}

func (optHTTP OptionsHTTP) takeHeaderRequest(r *http.Request) {
	for k, v := range optHTTP.headerRequestMap {
		r.Header.Add(k, v)
	}
}

// WithHTML set template html
func (optHTTP *OptionsHTTP) WithHTML(temp *template.Template, nameHTML string) *OptionsHTTP {
	htmlCopy := temp.Lookup(nameHTML)
	if htmlCopy == nil {
		log.Fatalf("%s - template missing", nameHTML)
	}
	optHTTP.html = htmlCopy
	return optHTTP
}

// TakeHTML getting html template
func (optHTTP *OptionsHTTP) TakeHTML() *template.Template {
	return optHTTP.html
}

// ReceptionRedirectURL informative function that indicates that the handler is ready to receive the redirectAnswer type
func (optHTTP *OptionsHTTP) ReceptionRedirectURL() *OptionsHTTP { return optHTTP }

// SetConnectingToMicroservices connecting grpc services to HTTP handler
func (optHTTP *OptionsHTTP) SetConnectingToMicroservices(connections map[string]grpc.ClientConnInterface) *OptionsHTTP {
	for k, v := range connections {
		optHTTP.connectingToMicroservicesMap[k] = v
	}
	return optHTTP
}

// TakeConnGrpc getting grpc client
func (optHTTP OptionsHTTP) TakeConnGrpc(nameGrpc string) grpc.ClientConnInterface {
	return optHTTP.connectingToMicroservicesMap[nameGrpc]
}

// URLSendRedirectOk setting the reddirect path
// you need to change the path manually using the ChangePathOkRedirect method in your HandlerLogics
func (optHTTP *OptionsHTTP) URLSendRedirectOk(pathRedirect string) *OptionsHTTP {
	optHTTP.RedirectUseDataURL = true
	optHTTP.RedirectHTTPPathOk = pathRedirect
	return optHTTP
}

func (optHTTP *OptionsHTTP) ChangePathOkRedirect(pattern string, fill string) {
	optHTTP.RedirectHTTPPathOk = strings.ReplaceAll(optHTTP.RedirectHTTPPathOk, pattern, fill)
}

// URLSendRedirectMistake setting the reddirect path
// you need to change the path manually using the ChangePathMistakeRedirect method in your HandlerLogics
func (optHTTP *OptionsHTTP) URLSendRedirectMistake(pathRedirect string) *OptionsHTTP {
	optHTTP.RedirectUseDataURL = true
	optHTTP.RedirectHTTPPathMistake = pathRedirect
	return optHTTP
}

func (optHTTP *OptionsHTTP) ChangePathMistakeRedirect(pattern string, fill string) {
	optHTTP.RedirectHTTPPathOk = strings.ReplaceAll(optHTTP.RedirectHTTPPathOk, pattern, fill)
}

// HandlerLogics the handler logic function must match this type
// after the successful execution of the function, the response must be written to the channel to complete the work of the handler
// to terminate the work due to an error, you need to use the function context.CancelCauseFunc
type HandlerLogics func(context.Context, context.CancelCauseFunc, *OptionsHTTP, *http.Request, chan []byte)

type redirectAnswer struct {
	Ok          bool
	OkInfo      string
	MistakeInfo string
}

// WriteRedirectAnswerInfoOk writing information for a reddirect to io.Writer
// then the information should be written to the channel
func WriteRedirectAnswerInfoOk(writer io.Writer, okInfo string) error {
	if err := JSON.NewEncoder(writer).Encode(redirectAnswer{
		Ok:     true,
		OkInfo: okInfo,
	}); err != nil {
		return err
	}
	return nil
}

// WriteRedirectAnswerInfoErr writing error information for a reddirect to io.Writer
// then the recorded information should be used in the context.CancelCauseFunc function
func WriteRedirectAnswerInfoErr(writer io.Writer, mistakeInfo string) error {
	if err := JSON.NewEncoder(writer).Encode(redirectAnswer{
		Ok:          true,
		MistakeInfo: mistakeInfo,
	}); err != nil {
		return err
	}
	return nil
}

// TakeRedirectAnswerFromURL get redirectAnswer type from URL, if not then null
func TakeRedirectAnswerFromURL(r *http.Request) (*redirectAnswer, error) {
	var (
		params              = r.URL.Query()
		data                = params.Get("data")
		redirectAnswerValue = redirectAnswer{}
	)
	if err := JSON.NewDecoder(strings.NewReader(data)).Decode(&redirectAnswerValue); err != nil {
		if err == io.EOF {
		} else {
			return nil, err
		}
	}
	return &redirectAnswerValue, nil
}

// HandlerLogicsRun connecting logic to HTTP handler
func (optHTTP *OptionsHTTP) HandlerLogicsRun(ctx context.Context, timeCtx time.Duration, logicsHandler HandlerLogics) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, ctxCancel := context.WithTimeout(ctx, timeCtx)
		defer ctxCancel()
		optHTTP.takeHeaderRequest(r)
		ctx, ctxReturnErr := context.WithCancelCause(ctx)
		chanResult := make(chan []byte, 1)

		go logicsHandler(ctx, ctxReturnErr, optHTTP, r, chanResult)

		select {
		case <-ctx.Done():
			switch ctx.Err() {
			case context.DeadlineExceeded:
				LogZap.Error("timeout expired")
				w.WriteHeader(http.StatusRequestTimeout)
				return
			default:
				err := context.Cause(ctx)
				LogZap.Error(err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

		case result := <-chanResult:
			switch optHTTP.RedirectUseDataURL {
			case true:
				redirectAnswerValue := redirectAnswer{}
				if err := JSON.NewDecoder(bytes.NewReader(result)).Decode(&redirectAnswerValue); err != nil {
					if err == io.EOF {
						LogZap.Error("responseAnswer type expected")
						w.WriteHeader(http.StatusInternalServerError)
						return
					} else {
						LogZap.Error(err.Error())
						w.WriteHeader(http.StatusInternalServerError)
						return
					}
				}
				if redirectAnswerValue.OkInfo != "" {
					http.Redirect(w, r, optHTTP.RedirectHTTPPathOk+"?data="+url.QueryEscape(string(result)), http.StatusMovedPermanently)
					return
				} else {
					http.Redirect(w, r, optHTTP.RedirectHTTPPathMistake+"?data="+url.QueryEscape(string(result)), http.StatusMovedPermanently)
					return
				}
			default:
				optHTTP.takeHeaderResponse(w)
				w.Write(result)
				return
			}
		}
	}
}
