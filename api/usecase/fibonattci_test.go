package usecase

import (
	"math/big"
	"testing"
)

func TestCalculate(t *testing.T) {
  //今回はRepositoryに依存していないのでmockなどは生成せずテストを行う
	usecase := fibonacciUsecase{}

	testCases := []struct {
		input    string
		expected *big.Int
		wantErr  bool
	}{
    //正常系
		{"1", big.NewInt(1), false},
		{"2", big.NewInt(1), false},
		{"5", big.NewInt(5), false},
		{"10", big.NewInt(55), false},
    //異常系
		{"-1", nil, true},
		{"abc", nil, true},
	}

	for _, tc := range testCases {
		result, err := usecase.Calculate(tc.input)

		if tc.wantErr {
			if err == nil {
        t.Errorf("エラーを出力しないといけないのに出力されていません。 INPUTされた内容:%s",tc.input)
			}
		} else {
			if err != nil {
        t.Errorf("予期せぬエラーを受け取りました INPUTされた内容:%s エラー:%v", tc.input, err)
			}
			if result.Cmp(tc.expected) != 0 {
        t.Errorf("期待した内容と異なります 予期した内容:%v 実際の結果:%v INPUTされた内容:%s", tc.expected, result, tc.input)
			}
		}
	}
}
