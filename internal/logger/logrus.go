package log

import (
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

type logger struct {
	httpDebugLog *logrus.Logger
	mQDebugLog   *logrus.Logger
	mQErrorLog   *logrus.Logger
	httpErrorLog *logrus.Logger
	fatalLog     *logrus.Logger
	infoLog      *logrus.Logger
}

// NewErrorLogger crates a new logger to help track activities of the entire solution during the usage of the service
func NewLogrusLogger(httpDebugFile *os.File, mQDebugFile *os.File, mQErrorLgFile *os.File, httpErrorLgFile *os.File, fatalLgFile *os.File, infoLgFile *os.File) Logger {

	return &logger{
		httpDebugLog: setup(httpDebugFile, logrus.DebugLevel),
		mQDebugLog:   setup(mQDebugFile, logrus.InfoLevel),
		mQErrorLog:   setup(mQErrorLgFile, logrus.WarnLevel),
		httpErrorLog: setup(httpErrorLgFile, logrus.ErrorLevel),
		fatalLog:     setup(fatalLgFile, logrus.FatalLevel),
		infoLog:      setup(infoLgFile, logrus.FatalLevel)}
}

func setup(f *os.File, logLevel logrus.Level) *logrus.Logger {
	//if f != nil {
	lg := logrus.New()
	lg.SetOutput(f)
	lg.SetFormatter(&logrus.JSONFormatter{})
	lg.SetLevel(logLevel)

	return lg
	//}
	//return nil
}

// HttpDebug, writes the HttpDebug messages to the log file
func (l *logger) HttpDebug(r *http.Request, fn string, err error) {
	l.httpDebugLog.WithFields(
		logrus.Fields{
			"remote address": r.RemoteAddr,
			"method":         r.Method,
			"path":           r.URL.Path,
			"function":       fn,
			"error msg":      err,
		},
	).Info("debug")
}

// HttpError, writes the HttpError messages to the log file
func (l *logger) HttpError(r *http.Request, fn string, err error) {
	l.httpErrorLog.WithFields(
		logrus.Fields{
			"remote address": r.RemoteAddr,
			"method":         r.Method,
			"path":           r.URL.Path,
			"function":       fn,
			"error msg":      err,
		},
	).Info("debug")
}

// Fatal, writes the Fatal messages to the log file
func (l *logger) Fatal(fn string, err error) {
	l.fatalLog.WithFields(
		logrus.Fields{
			"function":  fn,
			"error msg": err,
		},
	).Fatal("debug")
}

//Info, writes app usage info messages to the log file
func (l *logger) Info(fn string) {
	l.infoLog.WithFields(
		logrus.Fields{
			"info": fn,
		},
	).Info("debug")
}

// MqDebug, writes the Message Queue info to the log file
func (l *logger) MqDebug(action string, msg interface{}) {
	l.mQDebugLog.WithFields(
		logrus.Fields{
			"action":  action,
			"message": msg,
		},
	).Info("debug")
}

// MqError, writes the Message Queue errors to the log file
func (l *logger) MqError(action string, msg interface{}, err error) {
	l.mQErrorLog.WithFields(
		logrus.Fields{
			"action":    action,
			"message":   msg,
			"error msg": err,
		},
	).Info("debug")
}
