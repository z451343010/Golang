package test

import "testing"

func TestStore(t *testing.T) {

	filePath := "F:\\Golang\\goProgramProject\\src\\unitTest\\unitTestExc\\monsterJson.txt"
	monster := &Monster{
		Name:  "zjune",
		Age:   24,
		Skill: "programing",
	}

	monster.Store(filePath)

}
