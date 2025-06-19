function getIntersectingPointBetweenLinkedList(
  firstLinkedList,
  secondLinkedList
) {
  let rootForFirstLinkedList = firstLinkedList.head;

  while (rootForFirstLinkedList) {
    const dataForFirstLinkedList = rootForFirstLinkedList.data;
    let rootForSecondLinkedList = secondLinkedList.head;

    while (rootForSecondLinkedList) {
      //When we find the intersecting node, we return it
      if (rootForSecondLinkedList.data === dataForFirstLinkedList) {
        return dataForFirstLinkedList;
      }

      rootForSecondLinkedList = rootForSecondLinkedList.next;
    }

    rootForFirstLinkedList = rootForFirstLinkedList.next;
  }

  // Here we cannot find the intersecting node so we return null
  return null;
}

function main() {
  try {
    const firstLinkedList = JSON.parse(process.argv[2]);
    const secondLinkedList = JSON.parse(process.argv[3]);

    const result = getIntersectingPointBetweenLinkedList(
      firstLinkedList,
      secondLinkedList
    );

    console.log(result);
  } catch (error) {
    console.log("An error occurred:", error);
  }
}

main();
