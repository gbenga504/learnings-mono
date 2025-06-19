function bubbleSort(array) {
  let leastSortedPartitionIndex = Number.MAX_SAFE_INTEGER;
  let i = 0;

  while (i < array.length - 1 && i + 1 < leastSortedPartitionIndex) {
    if (array[i] > array[i + 1]) {
      let temp = array[i + 1];

      array[i + 1] = array[i];
      array[i] = temp;
    }

    if (
      i + 1 === array.length - 1 &&
      leastSortedPartitionIndex === Number.MAX_SAFE_INTEGER
    ) {
      leastSortedPartitionIndex = array.length - 1;
      i = 0;
    } else if (i + 2 === leastSortedPartitionIndex) {
      leastSortedPartitionIndex = leastSortedPartitionIndex - 1;
      i = 0;
    } else {
      i++;
    }
  }
}

function main() {
  try {
    const array = JSON.parse(process.argv[2]).slice(0);
    bubbleSort(array);

    console.log(array);
  } catch (error) {
    console.log("Error is:", error);
  }
}

main();
