package test

import "testing"

func TestReStore(t *testing.T) {

	filePath := "F:\\Golang\\goProgramProject\\src\\unitTest\\unitTestExc\\monsterJson.txt"

	var monster2 Monster
	monster2.ReStore(filePath)

}
