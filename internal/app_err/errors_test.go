package app_err

import (
	"errors"
	"fmt"
	"runtime"
	"testing"
)

func getCallerInfo() (string, int) {
	_, file, line, _ := runtime.Caller(1)
	return file, line
}

func TestNewErrorWithoutOriginalError(t *testing.T) {
	file, line := getCallerInfo()
	err := NewError("500", "Something went wrong")

	// 检查错误消息是否符合预期
	expectedErrMsg := fmt.Sprintf("[%s:%d] Something went wrong (Code: 500)", file, line+1)
	if err.Error() != expectedErrMsg {
		t.Errorf("Error message does not match. Expected: %s, Got: %s", expectedErrMsg, err.Error())
	}

	// 检查错误代码是否正确
	if err.Code != "500" {
		t.Errorf("Error code does not match. Expected: 500, Got: %s", err.Code)
	}

	// 检查原始错误是否为nil
	if err.Err != nil {
		t.Errorf("Expected original error to be nil, but got: %v", err.Err)
	}
}

func TestNewErrorWithOriginalError(t *testing.T) {
	originalErr := errors.New("Original error")
	file, line := getCallerInfo()
	errWithOriginal := NewError("500", "Something went wrong", originalErr)

	// 检查错误消息是否包含原始错误信息
	expectedErrMsgWithOriginal := fmt.Sprintf("[%s:%d] Something went wrong (Code: 500) - Original error", file, line+1)
	if errWithOriginal.Error() != expectedErrMsgWithOriginal {
		t.Errorf("Error message with original error does not match. Expected: %s, Got: %s", expectedErrMsgWithOriginal, errWithOriginal.Error())
	}

	// 检查原始错误是否与预期一致
	if errWithOriginal.Err != originalErr {
		t.Errorf("Original error does not match. Expected: %v, Got: %v", originalErr, errWithOriginal.Err)
	}
}

func TestErrorsIs(t *testing.T) {
	// 创建一个内部错误
	innerErr := errors.New("Inner error")

	t.Run("test errors.IS equal", func(t *testing.T) {
		// 创建一个AppError
		anotherErr := NewError("500", "Something went wrong", innerErr)
		if ok := errors.Is(anotherErr, innerErr); !ok {
			t.Errorf("Expected errors.Is ")
		}
	})

	t.Run("test errors.IS not equal", func(t *testing.T) {
		// 创建一个AppError
		anotherErr := NewError("500", "Something went wrong", errors.New("anotherErr"))
		if ok := errors.Is(anotherErr, innerErr); ok {
			t.Errorf("Expected errors.Is shoud not equal")
		}
	})

	t.Run("test errors.IS error chain", func(t *testing.T) {
		// 创建一个包装了内部错误的外部错误
		outerErr := fmt.Errorf("Outer error: %w", innerErr)

		// 创建一个AppError
		anotherErr := NewError("500", "Something went wrong", outerErr)
		if ok := errors.Is(anotherErr, innerErr); !ok {
			t.Errorf("Expected errors.Is shoud not equal")
		}
	})

}
