function mergeList(array) {
  while (array.length > 1) {
    let newArray = array[0].concat(array[1]);
    array.splice(0, 2, newArray);

    mergeSort(array[0], 0, array[0].length - 1);
  }

  return array[0];
}

function mergeSort(array, lowerBoundary, upperBoundary) {
  if (lowerBoundary >= upperBoundary) {
    return;
  }

  let pivot = Math.floor((upperBoundary + lowerBoundary) / 2);

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
    if (array[leftStart] <= array[rightStart]) {
      temp.push(array[leftStart]);
      leftStart++;
    } else if (array[rightStart] <= array[leftStart]) {
      temp.push(array[rightStart]);
      rightStart++;
    }
  }

  let remnant =
    leftStart > leftEnd
      ? array.slice(rightStart, rightEnd + 1)
      : array.slice(leftStart, leftEnd + 1);

  temp = temp.concat(remnant);

  array.splice(lowerBoundary, upperBoundary - lowerBoundary + 1, ...temp);
}

function main() {
  try {
    const array = JSON.parse(process.argv[2]).slice(0);
    const result = mergeList(array);

    console.log(result);
  } catch (error) {
    console.log("Error occurred:", error);
  }
}

main();
