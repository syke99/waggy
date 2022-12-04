package v2

import (
	"github.com/stretchr/testify/assert"
	"github.com/syke99/waggy/v2/internal/resources"
	"os"
	"testing"
)

func TestOverrideParentLogger(t *testing.T) {
	// Act
	o := OverrideParentLogger()

	// Assert
	assert.Equal(t, true, o())
}

func TestNewLogger(t *testing.T) {
	// Arrange
	level := Info
	testLogFile := resources.TestLogFile

	// Act
	l := NewLogger(level, testLogFile)

	// Assert
	assert.IsType(t, resources.TestLogFile, l.log)
	assert.Equal(t, Info.level(), l.logLevel)
}

func TestLogger_SetLogFile(t *testing.T) {
	// Arrange
	l := Logger{
		logLevel: "",
		key:      "",
		message:  "",
		err:      "",
		vals:     make(map[string]interface{}),
		log:      nil,
	}
	testLogFile := resources.TestLogFile

	// Act
	err := l.SetLogFile(testLogFile)

	// Assert
	assert.NoError(t, err)
	assert.IsType(t, resources.TestLogFile, l.log)
}

func TestLogger_SetLogFile_Error(t *testing.T) {
	// Arrange
	l := Logger{
		logLevel: "",
		key:      "",
		message:  "",
		err:      "",
		vals:     make(map[string]interface{}),
		log:      nil,
	}

	// Act
	err := l.SetLogFile(nil)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, "no log file provided", err.Error())
}

func TestLogger_Level(t *testing.T) {
	// Arrange
	l := Logger{
		logLevel: "",
		key:      "",
		message:  "",
		err:      "",
		vals:     make(map[string]interface{}),
		log:      nil,
	}
	testLogLevel := Info

	// Act
	l.Level(testLogLevel)

	// Assert
	assert.Equal(t, Info.level(), l.logLevel)
}

func TestLogger_Err(t *testing.T) {
	// Arrange
	l := Logger{
		logLevel: "",
		key:      "",
		message:  "",
		err:      "",
		vals:     make(map[string]interface{}),
		log:      nil,
	}
	testErr := resources.TestError

	// Act
	l.Err(testErr)

	// Assert
	assert.Equal(t, testErr.Error(), l.err)
}

func TestLogger_Err_Nil(t *testing.T) {
	// Arrange
	l := Logger{
		logLevel: "",
		key:      "",
		message:  "",
		err:      "",
		vals:     make(map[string]interface{}),
		log:      nil,
	}

	// Act
	l.Err(nil)

	// Assert
	assert.Equal(t, "", l.err)
}

func TestLogger_Msg(t *testing.T) {
	// Arrange
	l := Logger{
		logLevel: "",
		key:      "",
		message:  "",
		err:      "",
		vals:     make(map[string]interface{}),
		log:      nil,
	}
	testKey := resources.TestKey
	testMsg := resources.TestMessage

	// Act
	l.Msg(testKey, testMsg)

	// Assert
	assert.Equal(t, resources.TestKey, l.key)
	assert.Equal(t, resources.TestMessage, l.message)
}

func TestLogger_Val(t *testing.T) {
	// Arrange
	l := Logger{
		logLevel: "",
		key:      "",
		message:  "",
		err:      "",
		vals:     make(map[string]interface{}),
		log:      nil,
	}
	testKey := resources.TestKey
	testValue := resources.TestValue

	//Act
	l.Val(testKey, testValue)

	assert.Equal(t, 1, len(l.vals))
	for k, v := range l.vals {
		assert.Equal(t, k, resources.TestKey)
		assert.Equal(t, v, resources.TestValue)
	}
}

func TestLogger_Log(t *testing.T) {
	// Arrange
	l := Logger{
		logLevel: Info.level(),
		key:      "",
		message:  "",
		err:      "",
		vals:     make(map[string]interface{}),
		log:      os.Stdout,
	}

	// Act
	_, err := l.Log()

	// Assert
	assert.NoError(t, err)
}

func TestLogger_Log_Err(t *testing.T) {
	// Arrange
	l := Logger{
		logLevel: Info.level(),
		key:      "",
		message:  "",
		err:      "",
		vals:     make(map[string]interface{}),
		log:      nil,
	}

	// Act
	_, err := l.Log()

	// Assert
	assert.Error(t, err)
}
