package sumlinkedlist

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

type LinkedList struct {
	Data string      `json:"data"` // struct tags are mostly used when encoding data to rename the field in json.
	Next *LinkedList `json:"next"` // We capitalize the fields so that they are exported outside of the struct and can be filled up when json decoding
}

func getReversedDatasInLinkedList(linkedlist LinkedList, result string) string {
	if linkedlist.Next == nil {
		return linkedlist.Data + result
	}

	return getReversedDatasInLinkedList(*linkedlist.Next, linkedlist.Data+result)
}

func getLinkedListFromReversedDatas(reversedData string) LinkedList {
	var result *LinkedList = nil

	for _, char := range reversedData {
		linkedList := LinkedList{Data: fmt.Sprintf("%c", char), Next: result}

		result = &linkedList
	}

	return *result
}

func sumOfLinkedList(linkedList1 LinkedList, linkedList2 LinkedList) LinkedList {
	reversedLinkedList1Data, err := strconv.Atoi(getReversedDatasInLinkedList(linkedList1, ""))

	if err != nil {
		log.Panicf("An error occurred while trying to convert to numbers %v", err.Error())
	}

	reversedLinkedList2Data, err := strconv.Atoi(getReversedDatasInLinkedList(linkedList2, ""))

	if err != nil {
		log.Panicf("An error occurred while trying to convert to numbers %v", err.Error())
	}

	reversedData := reversedLinkedList1Data + reversedLinkedList2Data

	return getLinkedListFromReversedDatas(fmt.Sprint(reversedData))
}

func Init() {
	var linkedList1 LinkedList
	var linkedList2 LinkedList
	var err error

	fmt.Println("=========== Program begins ==============")

	fmt.Println("Enter the first linked list")
	err = json.NewDecoder(os.Stdin).Decode(&linkedList1)

	if err != nil {
		fmt.Printf("An error occurred while reading the first linked list ===> %v\n", err.Error())

		return
	}

	fmt.Println("Enter the second linked list")
	err = json.NewDecoder(os.Stdin).Decode(&linkedList2)

	if err != nil {
		fmt.Printf("An error occurred while reading the second linked list ===> %v\n", err.Error())

		return
	}

	result := sumOfLinkedList(linkedList1, linkedList2)
	sum := getReversedDatasInLinkedList(result, "")

	fmt.Printf("The sum of the linked list is %v and the value is %#v\n", sum, result)
}
