package cal

import (
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite struct {
	suite.Suite
	finalTest   int
	finalResult int
	testInput   [][]int
	testOutput  []int
}

func (s *Suite) SetupSuite() {
	s.testInput = [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
	}
	s.testOutput = []int{3, 5, 7}
	s.finalTest = 0
	for _, n := range s.testOutput {
		s.finalResult += n
	}
}

//AfterTest 用於測試完畢之後的檢查
func (s *Suite) AfterTest(_, _ string) {
	require.Equal(s.T(), s.finalResult, s.finalTest, "should be equal")
}

//TestStart 為測試程式進入點
func TestStart(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) TestAdd() {
	for i, in := range s.testInput {
		result := Add(in[0], in[1])
		s.finalTest += result
		require.Equal(s.T(), s.testOutput[i], result, "should be equal")
	}
}
