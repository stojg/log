package log_test
import (
	"github.com/stojg/log"
	"testing"
)

func TestDebug(t *testing.T) {
	log.Debug("debug: should not show")
}

func TestDebugF(t *testing.T) {
	log.Debugf("debugf: %s", "should not show")
	log.ShowDebug = false
}

func TestDebugEnabled(t *testing.T) {
	log.ShowDebug = true
	log.Debug("debug: should show")
	log.ShowDebug = false
}

func TestDebugFEnabled(t *testing.T) {
	log.ShowDebug = true
	log.Debugf("debugf: %s", "should show")
	log.ShowDebug = false
}

func TestInfo(t *testing.T) {
	log.Info("info: asdsd")
}

func TestInfoF(t *testing.T) {
	log.Infof("infof: %s", "asdsd")
}

func TestError(t *testing.T) {
	log.Error("error: asdsd")
}

func TestErrorF(t *testing.T) {
	log.Errorf("errorf: %s", "asdsd")
}