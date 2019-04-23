package logger

import (
	"github.com/Microsoft/KubeDevice-API/pkg/utils"
	"k8s.io/klog"
)

func Log(level int, format string, args ...interface{}) {
	if klog.V(klog.Level(level)) {
		klog.Infof(format, args...)
	}
}

func Error(format string, args ...interface{}) {
	klog.Errorf(format, args...)
}

func Warning(format string, args ...interface{}) {
	klog.Warningf(format, args...)
}

func SetLogger() {
	utils.Logf = Log
	utils.Errorf = Error
	utils.Warningf = Warning
}
