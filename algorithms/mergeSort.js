function mergeSort(array, lowerBoundary, upperBoundary) {
  if (lowerBoundary >= upperBoundary) {
    return;
  }

  const pivot = Math.floor((upperBoundary + lowerBoundary) / 2);

  mergeSort(array, lowerBoundary, pivot);
  mergeSort(array, pivot + 1, upperBoundary);
  mergeHalves(array, lowerBoundary, upperBoundary);
}

function mergeHalves(array, lowerBoundary, upperBoundary) {
  let pivot = Math.floor((upperBoundary + lowerBoundary) / 2);

  let leftStart = lowerBoundary;
  let leftEnd = pivot;
  let rightStart = pivot + 1;
  let rightEnd = upperBoundary;
  let temp = [];

  while (leftStart <= leftEnd && rightStart <= rightEnd) {
    if (array[leftStart] < array[rightStart]) {
      temp.push(array[leftStart]);
      leftStart++;
    }

    if (array[rightStart] < array[leftStart]) {
      temp.push(array[rightStart]);
      rightStart++;
    }
  }

  const remainingContent =
    leftStart > leftEnd
      ? array.slice(rightStart, rightEnd + 1)
      : array.slice(leftStart, leftEnd + 1);

  temp = temp.concat(remainingContent);
  array.splice(lowerBoundary, upperBoundary - lowerBoundary + 1, ...temp);
}

function main() {
  try {
    let array = JSON.parse(process.argv[2]).slice(0);
    mergeSort(array, 0, array.length - 1);

    console.log(array);
  } catch (error) {
    console.log("Error occurred: ", error);
  }
}

main();
